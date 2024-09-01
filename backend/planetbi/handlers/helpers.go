package handlers

import (
	"crypto/rand"
	"errors"
	"github.com/gorilla/sessions"
	"net/http"
	"os"
	"regexp"

	jsoniter "github.com/json-iterator/go"
)

var CookieStore *sessions.CookieStore
var emailReg *regexp.Regexp

func init() {
	CookieStore = sessions.NewCookieStore([]byte("RFBP8tfYi7ttKIZGJhCJPtjfLnhaMRQe"))
	emailReg = regexp.MustCompile(`\S+@\S+\.\S+`)
}

type authConfig struct {
	Redirect bool
}

// authenticate checks if the user is authenticated
// If the user is not authenticated, it redirects to the signin page
// If the user is authenticated, it returns the userID
// TODO: Check revocation of session
func authenticate(w http.ResponseWriter, r *http.Request, config ...authConfig) (string, error) {
	// CookieStore.Get always returns a session
	// and returns an error if the session exists but could not be decoded
	session, err := CookieStore.Get(r, "session")

	// return nil error only in this case
	if session.Values["authenticated"] == true && session.Values["user_id"] != nil && session.Values["user_id"] != "" {
		return session.Values["user_id"].(string), nil
	}

	// nil error doesn't mean that the user is authenticated
	if err == nil {
		err = errors.New("session not found")
	}

	// don't redirect if config.Redirect is false
	if len(config) > 0 && !config[0].Redirect {
		return "", err
	}

	http.Redirect(w, r, "/signin", http.StatusTemporaryRedirect)
	return "", err
}

func random(strenght int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, strenght)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func isProduction() bool {
	// https://cloud.google.com/appengine/docs/standard/go/runtime
	return os.Getenv("GOOGLE_CLOUD_PROJECT") == "planetbi"
}

// response 200 with json
func JSON(w http.ResponseWriter, data interface{}, statuses ...int) {
	jsonObj, err := jsoniter.Marshal(data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if len(statuses) > 0 {
		w.WriteHeader(statuses[0])
	}
	w.Write(jsonObj)
}
