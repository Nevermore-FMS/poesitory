package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/Nevermore-FMS/poesitory/backend/database"
)

var gitlabClientId = os.Getenv("GITLAB_CLIENT_ID")
var gitlabClientSecret = os.Getenv("GITLAB_CLIENT_SECRET")

var gitlabRedirectUri url.URL

func initGitlab() {
	gitlabRedirectUri = *SelfUrl
	gitlabRedirectUri.Path = "/api/gitlab/callback"
}

func GitlabLoginHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := url.URL{
			Scheme: "https",
			Host:   "gitlab.com",
			Path:   "/oauth/authorize",
		}
		query := url.Query()
		query.Add("client_id", gitlabClientId)
		query.Add("redirect_uri", gitlabRedirectUri.String())
		query.Add("response_type", "code")
		query.Add("scope", "read_user")
		url.RawQuery = query.Encode()
		http.Redirect(w, r, url.String(), http.StatusFound)
	})
}

func GitlabCallbackHandler() http.Handler {
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
			redirectError("Gitlab login cancelled")
			return
		}
		code := r.URL.Query().Get("code")
		if len(code) < 1 {
			redirectError("Reponse did not contain code")
			return
		}
		gitlabToken := GetGitlabToken(r.URL.Query().Get("code"))
		if gitlabToken == nil {
			redirectError("Unable to obtain gitlab token")
			return
		}
		gitlabUser := GetGitlabUser(*gitlabToken)
		if gitlabUser == nil {
			redirectError("Unable to obtain gitlab user")
			return
		}

		user := database.GetUserByEmail(gitlabUser.Email)
		if user == nil {
			userID, err := database.CreateUser(gitlabUser.Username, gitlabUser.Email)
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

func GetGitlabToken(code string) *string {
	resp, err := http.PostForm("https://gitlab.com/oauth/token", url.Values{
		"client_id":     {gitlabClientId},
		"client_secret": {gitlabClientSecret},
		"code":          {code},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {gitlabRedirectUri.String()},
	})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var details struct {
		AccessToken string `json:"access_token"`
	}
	err = json.NewDecoder(resp.Body).Decode(&details)
	if err != nil {
		panic(err)
	}
	return &details.AccessToken
}

type GitlabUser struct {
	Username string
	Email    string
}

func GetGitlabUser(token string) *GitlabUser {
	req, err := http.NewRequest(http.MethodGet, "https://gitlab.com/api/v4/user", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil
	}
	var user GitlabUser
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	return &user
}
