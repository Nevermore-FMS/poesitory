package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/Nevermore-FMS/poesitory/backend/database"
	"github.com/Nevermore-FMS/poesitory/backend/graph/model"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}
var tokenCtxKey = &contextKey{"token"}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := strings.Split(r.Header.Get("Authorization"), " ")

		if len(auth) == 2 {
			if auth[0] == "User" {
				user := database.GetUserByToken(auth[1])

				if user != nil {
					ctx := context.WithValue(r.Context(), userCtxKey, user)
					ctx = context.WithValue(ctx, tokenCtxKey, &auth[1])
					r = r.WithContext(ctx)

				}
			}
		}

		next.ServeHTTP(w, r)
	})
}

func UserForContext(ctx context.Context) *model.User {
	raw, ok := ctx.Value(userCtxKey).(*model.User)
	if !ok {
		return nil
	}
	return raw
}

func TokenForContext(ctx context.Context) *string {
	raw, ok := ctx.Value(tokenCtxKey).(*string)
	if !ok {
		return nil
	}
	return raw
}
