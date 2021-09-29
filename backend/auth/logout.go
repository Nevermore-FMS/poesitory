package auth

import (
	"net/http"
	"time"

	"github.com/Nevermore-FMS/poesitory/backend/database"
)

func LogoutHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if cookie, err := r.Cookie("token"); err == nil {
			user := database.GetUserByToken(cookie.Value)
			if user != nil {
				database.ExpireTokensForUser(user.ID)
			}
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   "",
				Expires: time.Unix(0, 0),
				Path:    "/",
			})
		}
		http.Redirect(w, r, SelfUrl.String(), http.StatusFound)
	})
}
