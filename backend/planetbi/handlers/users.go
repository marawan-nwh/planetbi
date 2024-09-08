package handlers

import (
	"html"
	"log/slog"
	"net/http"
	"net/mail"
	"planetbi/config"
	"planetbi/db"
	"planetbi/emails"
	"strings"

	"database/sql"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"

	"time"
)

var Users = []handler{
	{
		URI:    "/signup",
		Method: "POST",
		Params: "name,email,password",
		Do: func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			name := html.EscapeString(strings.TrimSpace(r.FormValue("name")))
			email := html.EscapeString(strings.TrimSpace(r.FormValue("email")))
			password := r.FormValue("password")

			// validate password
			// don't add complexity requirements
			// https://learn.microsoft.com/en-us/windows/security/threat-protection/security-policy-settings/password-must-meet-complexity-requirements
			if len(password) < 8 {
				http.Error(w, "", 407)
				return
			}

			// validate email
			// https://stackoverflow.com/a/66624104, https://stackoverflow.com/a/13975255
			_, err := mail.ParseAddress(email)
			if err != nil || !emailReg.MatchString(email) {
				http.Error(w, "", 408)
				return
			}

			// check if user already exists
			var exists bool
			err = db.Pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
			if exists {
				// get userID
				var userID string
				err := db.Pool.QueryRow(ctx, "SELECT id FROM users WHERE email = $1", email).Scan(&userID)
				if err != nil {
					slog.Error(err.Error())
					http.Error(w, "", http.StatusInternalServerError)
					return
				}

				// send email to user
				base := "http://localhost:" + config.ClientPort
				if isProduction() {
					base = "https://" + config.Domain
				}

				// TODO: remove the forgot-password link
				err = emails.Send(email, "Did you sign up for "+config.AppName+"?", `
Hello `+strings.Split(name, " ")[0]+`,<br/><br/>

It looks like you have tried to sign up for `+config.AppName+` with this email address, but you are already registered. Use your email and password to <a target="_blank" href="`+base+`/users/signin">sign in</a>, or follow the steps <a target="_blank" href="`+base+`/users/forgot-password">here</a> if you forgot your password.<br/><br/>

If you didn’t request to sign up, you can ignore this email.<br/><br/>

Thanks,<br/>
`+config.AppName+` Team`)
				if err != nil {
					slog.Error(err.Error())
					http.Error(w, "", http.StatusInternalServerError)
					return
				}
				return
			}

			// encrypt password
			// cost: https://stackoverflow.com/a/50470009
			encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), config.HashCost)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			// create user
			userID := xid.New().String()
			verificationToken := random(32)
			verificationTokenExpiry := time.Now().Add(config.EmailVerificationTokenLifetime)
			_, err = db.Pool.Exec(ctx, `INSERT INTO users (id, name, email, password, email_verification_token, email_verification_token_expiry) VALUES ($1, $2, $3, $4, $5, $6)`, userID, name, email, encryptedPassword, verificationToken, verificationTokenExpiry)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			// send verification email
			base := "http://localhost:" + config.ClientPort
			if isProduction() {
				base = "https://" + config.Domain
			}
			err = emails.Send(email, "Your verification link", `
Hello `+strings.Split(name, " ")[0]+`,<br/><br/>

Follow this link to verify your email address:
<a target="_blank" href="`+base+`/users/verify?user_id=`+userID+`&token=`+verificationToken+`">Verify</a><br/><br/>

If you didn’t ask to verify this address, you can ignore this email.<br/><br/>

Thanks,<br/>
`+config.AppName+` Team`)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
		}},
	{
		URI:    "/resend-verification-email",
		Method: "POST",
		Params: "email",
		Do: func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			email := r.FormValue("email")

			// check if user already verified
			var verified bool
			var name string
			var userID string
			var currentVerificationTokenExpiry sql.NullTime
			err := db.Pool.QueryRow(ctx, "SELECT id, name, email_verification_token_expiry, email_verified FROM users WHERE email = $1", email).Scan(&userID, &name, &currentVerificationTokenExpiry, &verified)
			if err == db.ErrNoRows {
				// if email doesn't exists, respond with 200 and do nothing
				// this is ok because in normal flow, the user shouldn't reach this point
				http.Error(w, "", http.StatusOK)
				return
			}
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
			if verified {
				// we shouldn't respond with a different status here
				// it will let attackers know if the email is registered
				// the user shouldn't reach this point too
				http.Error(w, "", http.StatusOK)
				return
			}

			// check if a verification email was sent a few moments ago
			if currentVerificationTokenExpiry.Valid && time.Now().Before(currentVerificationTokenExpiry.Time) {
				http.Error(w, "", http.StatusCreated)
				return
			}

			verificationToken := random(32)
			verificationTokenExpiry := time.Now().Add(config.EmailVerificationTokenLifetime)

			_, err = db.Pool.Exec(ctx, "UPDATE users SET email_verification_token = $1, email_verification_token_expiry = $2 WHERE id = $3", verificationToken, verificationTokenExpiry, userID)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			// send verification email
			base := "http://localhost:" + config.ClientPort
			if isProduction() {
				base = "https://" + config.Domain
			}
			err = emails.Send(email, "Your verification link", `
Hello `+strings.Split(name, " ")[0]+`,<br/><br/>

Follow this link to verify your email address:
<a target="_blank" href="`+base+`/users/verify?user_id=`+userID+`&token=`+verificationToken+`">Verify</a><br/><br/>

If you didn’t ask to verify this address, you can ignore this email.<br/><br/>

Thanks,<br/>
`+config.AppName+` Team`)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

		}},
	{
		URI:    "/verify",
		Method: "POST",
		Params: "user_id,token",
		Do: func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			userID := r.FormValue("user_id")
			token := r.FormValue("token")

			// check if user already verified
			var verified bool
			err := db.Pool.QueryRow(ctx, "SELECT email_verified FROM users WHERE id = $1", userID).Scan(&verified)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
			if verified {
				http.Error(w, "", http.StatusOK)
				return
			}

			var verificationToken string
			var verificationTokenExpiry sql.NullTime
			err = db.Pool.QueryRow(ctx, "SELECT email_verification_token, email_verification_token_expiry FROM users WHERE id = $1", userID).Scan(&verificationToken, &verificationTokenExpiry)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
			if verificationTokenExpiry.Valid && time.Now().After(verificationTokenExpiry.Time) {
				http.Error(w, "", http.StatusUnauthorized)
				return
			}
			if verificationToken != token || verificationToken == "" {
				http.Error(w, "", http.StatusUnauthorized)
				return
			}

			_, err = db.Pool.Exec(ctx, "UPDATE users SET email_verified = true, email_verification_token = '', email_verification_token_expiry = null WHERE id = $1", userID)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
		}},
	{
		URI:    "/signin",
		Method: "POST",
		Params: "email,password",
		Do: func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			email := r.FormValue("email")
			password := r.FormValue("password")

			var userID string
			var name string
			var verified bool
			var encryptedPassword string
			err := db.Pool.QueryRow(ctx, "SELECT id, name, email_verified, password FROM users WHERE email = $1", email).Scan(&userID, &name, &verified, &encryptedPassword)
			if err == db.ErrNoRows {
				http.Error(w, "", http.StatusUnauthorized)
				return
			}
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			if !verified {
				http.Error(w, "", 406)
				return
			}

			// compare password with bcrypt
			err = bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusUnauthorized)
				return
			}

			session, err := CookieStore.Get(r, "session")
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
			session.ID = xid.New().String()
			session.Values["user_id"] = userID
			session.Values["name"] = name
			session.Values["authenticated"] = true
			session.Options.MaxAge = 86400 * 30 // 30 days
			session.Options.SameSite = http.SameSiteLaxMode
			session.Options.HttpOnly = true

			err = session.Save(r, w)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			// update recent_login_at
			// don't return error if failed
			_, err = db.Pool.Exec(ctx, "UPDATE users SET recent_login_at = $1 WHERE id = $2", time.Now(), userID)
			if err != nil {
				slog.Error(err.Error())
			}

			JSON(w, &struct {
				Name string `json:"name"`
			}{Name: name})
		}},
	{
		URI:    "/init-reset-password",
		Method: "POST",
		Params: "email",
		Do: func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			email := r.FormValue("email")

			// check if user already exists
			var exists bool
			err := db.Pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			if !exists {
				// respond with success even if user doesn't exist
				// with fake correlation token
				JSON(w, &struct {
					PasswordResetCorrelationToken string `json:"password_reset_correlation_token"`
				}{PasswordResetCorrelationToken: random(32)})
				return
			}

			// check expiry of previous reset token
			var currentPasswordResetTokenExpiry sql.NullTime
			var currentPasswordCorrelationToken string
			err = db.Pool.QueryRow(ctx, "SELECT password_reset_token_expiry, password_reset_correlation_token FROM users WHERE email = $1", email).Scan(&currentPasswordResetTokenExpiry, &currentPasswordCorrelationToken)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			// if token is still valid, return 200 with the current correlation token
			if currentPasswordResetTokenExpiry.Valid && time.Now().Before(currentPasswordResetTokenExpiry.Time) {
				JSON(w, &struct {
					PasswordResetCorrelationToken string `json:"password_reset_correlation_token"`
				}{PasswordResetCorrelationToken: currentPasswordCorrelationToken})
				return
			}

			passwordResetToken := random(20)
			passwordResetCorrelationToken := random(32)
			passwordResetTokenExpiry := time.Now().Add(config.PasswordResetTokenLifetime)

			_, err = db.Pool.Exec(ctx, "UPDATE users SET password_reset_token = $1, password_reset_correlation_token = $2, password_reset_token_expiry = $3 WHERE email = $4", passwordResetToken, passwordResetCorrelationToken, passwordResetTokenExpiry, email)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			// send email to user with password reset token
			err = emails.Send(email, "Your password reset token", `
Hello,<br/><br/>

We received a request to reset your password.<br/><br/>

Please use the following token to complete the password reset process:<br/>
<b>`+passwordResetToken+`</b><br/><br/>

If you didn’t ask to reset your password, you can ignore this email.<br/><br/>

Thanks,<br/>
`+config.AppName+` Team`)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			JSON(w, &struct {
				PasswordResetCorrelationToken string `json:"password_reset_correlation_token"`
			}{PasswordResetCorrelationToken: passwordResetCorrelationToken})
		}},
	{
		URI:    "/reset-password",
		Method: "POST",
		Params: "email,password_reset_correlation_token,password_reset_token,new_password",
		Do: func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			email := r.FormValue("email")

			passwordResetCorrelationToken := r.FormValue("password_reset_correlation_token")
			passwordResetToken := r.FormValue("password_reset_token")

			var _passwordToken string
			var _passwordCorrelationToken string
			var _passwordResetTokenExpiry sql.NullTime
			err := db.Pool.QueryRow(ctx, "SELECT password_reset_token, password_reset_correlation_token, password_reset_token_expiry FROM users WHERE email = $1", email).Scan(&_passwordToken, &_passwordCorrelationToken, &_passwordResetTokenExpiry)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusUnauthorized)
				return
			}

			if _passwordResetTokenExpiry.Valid && time.Now().After(_passwordResetTokenExpiry.Time) {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusUnauthorized)
				return
			}

			if _passwordToken != passwordResetToken || _passwordToken == "" || _passwordCorrelationToken != passwordResetCorrelationToken || _passwordCorrelationToken == "" {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusUnauthorized)
				return
			}

			newPassword := r.FormValue("new_password")

			// encrypt password
			// cost: https://stackoverflow.com/a/50470009
			encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), config.HashCost)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			// update password, and also verify email
			_, err = db.Pool.Exec(ctx, "UPDATE users SET email_verified = true, email_verification_token = '', email_verification_token_expiry = null, password_reset_token = '', password_reset_correlation_token = '', password_reset_token_expiry = null, password = $1 WHERE email = $2", encryptedPassword, email)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusUnauthorized)
				return
			}
		}},
}
