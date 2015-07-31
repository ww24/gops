package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/kr/pretty"
	"github.com/mrjones/oauth"

	"github.com/ww24/gops/libs"
	"github.com/zenazn/goji/web"
)

func newConsumer() (consumer *oauth.Consumer) {
	consumer = oauth.NewConsumer(
		libs.Config["twitter"].(libs.Hash)["consumerKey"].(string),
		libs.Config["twitter"].(libs.Hash)["consumerSecret"].(string),
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authenticate",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		},
	)
	return
}

// Auth of Twitter
func Auth(c web.C, w http.ResponseWriter, r *http.Request) {
	defer libs.InternalServerError(w)

	consumer := newConsumer()
	token, authURL, err := consumer.GetRequestTokenAndUrl(libs.Config["twitter"].(libs.Hash)["callbackURL"].(string))
	if err != nil {
		panic(err)
	}

	tokenJSON, err := json.Marshal(token)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(tokenJSON))

	cookie := &http.Cookie{
		Name:     "token",
		Value:    url.QueryEscape(string(tokenJSON)),
		Expires:  time.Now().AddDate(0, 0, 7),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, authURL, 302)
}

// Callback of Twitter Authenticate
func Callback(c web.C, w http.ResponseWriter, r *http.Request) {
	defer libs.InternalServerError(w)

	consumer := newConsumer()

	tokenCookie, err := r.Cookie("token")
	if err != nil {
		panic(err)
	}

	tokenJSON, err := url.QueryUnescape(tokenCookie.Value)
	if err != nil {
		panic(err)
	}

	token := new(oauth.RequestToken)
	err = json.Unmarshal([]byte(tokenJSON), token)
	if err != nil {
		panic(err)
	}

	code := r.URL.Query()
	oauthVerifier, key1 := code["oauth_verifier"]
	_, key2 := code["oauth_token"]
	if key1 == false || key2 == false {
		http.Error(w, "Bad Request", 400)
		return
	}

	accessToken, err := consumer.AuthorizeToken(token, oauthVerifier[0])
	if err != nil {
		panic(err)
	}

	pretty.Println(accessToken)
	fmt.Println(accessToken.AdditionalData["screen_name"])

	http.Redirect(w, r, "/", 302)
}
