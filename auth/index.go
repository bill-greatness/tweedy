package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"tweet/tweedy/config"
	"tweet/tweedy/models"

	"github.com/dghubble/gologin/v2/twitter"
	"github.com/dghubble/sessions"
)

// session information
const (
	sessionName     = "tweedy-app"
	sessionSecret   = "randomSecret"
	sessionUserKey  = "bill_greatness"
	sessionUserName = "Bill Greatness"
)

var sessionStore = sessions.NewCookieStore([]byte(sessionSecret), nil)

// create serveMux with app routes.
func New(info *config.Config) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/profile", profileHandler)
	mux.HandleFunc("/logout", logoutHandler)
	mux.HandleFunc("/get-tweets", models.GetTimelineTweets)
	mux.HandleFunc("/get-my-tweets", models.GetUserTweets)
	mux.Handle("/twitter/login", twitter.LoginHandler(config.TwitterClient, nil))
	mux.Handle("/twitter/callback", twitter.CallbackHandler(config.TwitterClient, issueSession(), nil))
	return mux
}

// profile handler
func profileHandler(w http.ResponseWriter, req *http.Request) {
	session, err := sessionStore.Get(req, sessionName)

	if err != nil {
		page, _ := ioutil.ReadFile("pages/home.html")
		fmt.Fprintf(w, string(page))
		return
	}

	fmt.Fprintf(w, `
	<p> You are Logged in %s!</p>
	<a href="/get-tweets"> Get Tweets </a> 
	<form action="/logout" method="post"> <input type="submit" value="Logout">
	</form>`, session.Values[sessionUserName])
}

// issue sessions after login.
func issueSession() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		twitterUser, err := twitter.UserFromContext(ctx)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		session := sessionStore.New(sessionName)
		// store session information from a successful twitter login.
		session.Values[sessionUserKey] = twitterUser.ID
		session.Values[sessionUserName] = twitterUser.ScreenName
		session.Save(w)

		// redirect user to profile handler.
		http.Redirect(w, req, "/profile", http.StatusFound)
	}
	return http.HandlerFunc(fn)
}

func logoutHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		sessionStore.Destroy(w, sessionName)
	}
	http.Redirect(w, req, "/", http.StatusFound)
}

//OAuth2 Client ID: MHBOeXh6dkF6UWFZRXBEeDRhbmQ6MTpjaQ
//OAuth2 Client Secret: E9Xf1eK5JbSiIw0mviIPtjT3TFYVC7jEqaXCabY8SBtMbdsZff
