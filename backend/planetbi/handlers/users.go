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
				http.Error(w, "", http.StatusConflict)
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
}
