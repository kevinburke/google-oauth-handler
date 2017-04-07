package google_oauth_handler_test

import (
	"encoding/hex"
	"fmt"
	"net/http"

	google "github.com/kevinburke/google-oauth-handler"
	"golang.org/x/oauth2"
)

var key *[32]byte

func init() {
	secretKeyBytes, _ := hex.DecodeString("982a732cc3d72d13678dee2609cf55d736711ff1f293f95cab41bd45e5d77870")
	key = new([32]byte)
	copy(key[:], secretKeyBytes)
}

func Example() {
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
	http.Handle("/", auth.Handle(func(w http.ResponseWriter, r *http.Request, token *oauth2.Token) {
		fmt.Fprintf(w, "<html><body><h1>Hello World</h1><p>Token: %s</p></body></html>", token.AccessToken)
	}))
}
