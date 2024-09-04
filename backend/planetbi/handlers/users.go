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
			verificationToekn := random(32)
			_, err = db.Pool.Exec(ctx, `INSERT INTO users (id, name, email, password, email_verification_token) VALUES ($1, $2, $3, $4, $5)`, userID, name, email, encryptedPassword, verificationToekn)
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
<a target="_blank" href="`+base+`/users/verify?user_id=`+userID+`&token=`+verificationToekn+`">Verify</a><br/><br/>

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

			var verifiedToken string
			err = db.Pool.QueryRow(ctx, "SELECT email_verification_token FROM users WHERE id = $1", userID).Scan(&verifiedToken)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
			if verifiedToken != token {
				http.Error(w, "", http.StatusUnauthorized)
				return
			}

			_, err = db.Pool.Exec(ctx, "UPDATE users SET email_verified = true, email_verification_token = null WHERE id = $1", userID)
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
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			if !verified {
				http.Error(w, "", http.StatusForbidden)
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
}
