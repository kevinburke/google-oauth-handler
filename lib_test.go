package google_oauth_handler_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	google "github.com/kevinburke/google-oauth-handler"
)

func ExampleAuthenticator_URL() {
	cfg := google.Config{
		SecretKey: key,
		BaseURL:   "https://example.com",
		ClientID:  "customdomain.apps.googleusercontent.com",
		Secret:    "W-secretkey",
		Scopes: []string{
			"email",
			"https://www.googleapis.com/auth/gmail.send",
		},
	}
	auth := google.NewAuthenticator(cfg)
	r, _ := http.NewRequest("GET", "/", nil)
	fmt.Println(auth.URL(r)) // "https://accounts.google.com/o/oauth2/..."
}

func TestServeLogin(t *testing.T) {
	l := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Custom Login")
	}
	a := google.NewAuthenticator(google.Config{ServeLogin: http.HandlerFunc(l)})
	h := a.Handle(func(w http.ResponseWriter, r *http.Request, auth *google.Auth) {
		io.WriteString(w, "inside authenticated function")
	})
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	h.ServeHTTP(w, r)
	if w.Body.String() != "Custom Login" {
		t.Errorf("bad body, want %q, got %q", "Custom Login", w.Body.String())
	}
}
