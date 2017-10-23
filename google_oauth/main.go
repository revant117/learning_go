package main

import(
	"fmt"
	"net/http"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
)

const htmlIndex = `<html><body>
<a href="/GoogleLogin">Log in with Google</a>
</body></html>`

var (
	OauthConfig = &oauth2.Config{
		RedirectURL: "http://localhost:3000/GoogleCallback",
		ClientID: "225199385571-uopedgpk47lr87r7uff9cbntmr1vlkr7.apps.googleusercontent.com",
		ClientSecret: "r2Pj1vmveFFzKGIIu_cmjmCM",
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/drive.file",
		},
		Endpoint: google.Endpoint,
	}
	oauthStateString = "random" //must be random for each request that is sent to google
)


func main(){
	http.HandleFunc("/", handleMain)
    http.HandleFunc("/GoogleLogin", handleGoogleLogin)
    http.HandleFunc("/GoogleCallback", handleGoogleCallback)
	fmt.Println(http.ListenAndServe(":3000", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlIndex)
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := OauthConfig.AuthCodeURL(oauthStateString)
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
	token, err := OauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
        fmt.Println("Code exchange failed with '%s'\n", err)
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("Code exchange failed with '%s'\n", err)
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
	}
	fmt.Println(token)
	defer response.Body.Close()
    contents, err := ioutil.ReadAll(response.Body)
    fmt.Fprintf(w, "Content: %s\n", contents)
}