package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/Nevermore-FMS/poesitory/backend/database"
)

var githubClientId = os.Getenv("GITHUB_CLIENT_ID")
var githubClientSecret = os.Getenv("GITHUB_CLIENT_SECRET")

var githubRedirectUri url.URL

func initGithub() {
	githubRedirectUri = *SelfUrl
	githubRedirectUri.Path = "/api/github/callback"
}

func GithubLoginHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := url.URL{
			Scheme: "https",
			Host:   "github.com",
			Path:   "/login/oauth/authorize",
		}
		query := url.Query()
		query.Add("client_id", githubClientId)
		query.Add("redirect_uri", githubRedirectUri.String())
		query.Add("scope", "user:email")
		url.RawQuery = query.Encode()
		http.Redirect(w, r, url.String(), http.StatusFound)
	})
}

func GithubCallbackHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		redirectError := func(err string) {
			url := *SelfUrl
			url.Path = "/login"
			q := url.Query()
			q.Add("error", err)
			url.RawQuery = q.Encode()
			http.Redirect(w, r, url.String(), http.StatusFound)
		}
		qerr := r.URL.Query().Get("error")
		if len(qerr) > 1 {
			redirectError("Github login cancelled")
			return
		}
		code := r.URL.Query().Get("code")
		if len(code) < 1 {
			redirectError("Reponse did not contain code")
			return
		}
		githubToken := GetGithubToken(r.URL.Query().Get("code"))
		if githubToken == nil {
			redirectError("Unable to obtain github token")
			return
		}
		email := GetGithubEmail(*githubToken)
		if email == nil {
			redirectError("Unable to obtain github email")
			return
		}
		username := GetGithubUsername(*githubToken)
		if username == nil {
			redirectError("Unable to obtain github username")
			return
		}

		user := database.GetUserByEmail(*email)
		if user == nil {
			userID, err := database.CreateUser(*username, *email)
			if err != nil {
				if err == database.ErrUserAlreadyExists {
					redirectError("User exists with a different email address")
					return
				}
				panic(err)
			}
			user = database.GetUserByID(userID)
		}

		userToken, err := database.CreateTokenForUser(user.ID)
		if err != nil {
			panic(err)
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   userToken,
			Expires: time.Now().Add(24 * time.Hour),
			Path:    "/",
		})

		url := *SelfUrl
		url.Path = "/home"
		http.Redirect(w, r, url.String(), http.StatusFound)
	})
}

func GetGithubToken(code string) *string {
	resp, err := http.PostForm("https://github.com/login/oauth/access_token", url.Values{
		"client_id":     {githubClientId},
		"client_secret": {githubClientSecret},
		"code":          {code},
		"redirect_uri":  {githubRedirectUri.String()},
	})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	bodyString := string(bodyBytes)
	params, err := url.ParseQuery(bodyString)
	if err != nil {
		panic(err)
	}
	token := params.Get("access_token")
	return &token
}

type GithubEmail struct {
	Email    string
	Verified bool
	Primary  bool
}

func GetGithubEmail(token string) *string {
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/user/emails", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil
	}
	var emails []GithubEmail
	err = json.NewDecoder(resp.Body).Decode(&emails)
	if err != nil {
		panic(err)
	}
	for _, email := range emails {
		if email.Primary && email.Verified {
			return &email.Email
		}
	}
	return nil
}

func GetGithubUsername(token string) *string {
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/user", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil
	}
	var userMap map[string]json.RawMessage
	err = json.NewDecoder(resp.Body).Decode(&userMap)
	if err != nil {
		panic(err)
	}
	var username string
	err = json.Unmarshal(userMap["login"], &username)
	if err != nil {
		panic(err)
	}
	return &username
}
