package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
)

const htmlIndex = ` <html>
						<body>
							<a href="/GoogleLogin">Log in with Google</a>
						</body>
					</html> `

var googleOauthConfig = &oauth2.Config{
	ClientID:     "514396512421-5pfh8uqu0vsncffbicaj2v4occ8sg9qk.apps.googleusercontent.com", // your_client_id
	ClientSecret: "5ho8gqlROi36cLNP0k6Dzsis",                                                 // your_client_secret
	RedirectURL:  "http://localhost:8000/GoogleCallback",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
	},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://accounts.google.com/o/oauth2/auth",
		TokenURL: "https://accounts.google.com/o/oauth2/token",
	},
}

const oauthStateString = "random"

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/GoogleLogin", handleGoogleLogin)
	http.HandleFunc("/GoogleCallback", handleGoogleCallback)
	fmt.Println(http.ListenAndServe(":8000", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlIndex)
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	fmt.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("Code exchange failed with %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	contents, err := ioutil.ReadAll(response.Body)
	fmt.Printf("Google Accounts Info: %s", string(contents))

	defer response.Body.Close()

	fmt.Fprintf(w, "Content: %s\n", contents)
}
