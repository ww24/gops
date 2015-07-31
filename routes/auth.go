package routes

import (
	"fmt"
	"net/http"

	"github.com/ww24/gops/libs"
	"github.com/zenazn/goji/web"
	"golang.org/x/oauth2"
)

// Auth of Twitter
func Auth(c web.C, w http.ResponseWriter, r *http.Request) {

	config := &oauth2.Config{
		ClientID:     libs.Config["twitter"].(libs.Hash)["consumerKey"].(string),
		ClientSecret: libs.Config["twitter"].(libs.Hash)["consumerSecret"].(string),
		RedirectURL:  libs.Config["twitter"].(libs.Hash)["callbackURL"].(string),
		Endpoint: oauth2.Endpoint{
			AuthURL: "https://api.twitter.com/oauth/authenticate",
		},
	}

	fmt.Println(config)
}
