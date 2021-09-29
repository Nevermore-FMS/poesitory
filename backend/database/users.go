package database

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Nevermore-FMS/poesitory/backend/graph/model"
	"github.com/huandu/go-sqlbuilder"
)

var userStruct = sqlbuilder.NewStruct(new(model.User)).For(sqlbuilder.PostgreSQL)

var ErrUserDoesNotExist = errors.New("user does not exist")
var ErrUserAlreadyExists = errors.New("user already exists")

func GetUserByID(id string) *model.User {
	sb := userStruct.SelectFrom("users")
	sb.Where(sb.Equal("id", id))

	q, args := sb.Build()
	row := db.QueryRow(q, args...)
	user := model.User{}
	err := row.Scan(userStruct.Addr(&user)...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}
	return &user
}

func GetUserByEmail(email string) *model.User {
	sb := userStruct.SelectFrom("users")
	sb.Where(sb.Equal("email", email))

	q, args := sb.Build()
	row := db.QueryRow(q, args...)
	user := model.User{}
	err := row.Scan(userStruct.Addr(&user)...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}
	return &user
}

func GetUserByToken(token string) *model.User {
	h := sha256.New()
	h.Write([]byte(token))
	bs := h.Sum(nil)
	tokenHash := fmt.Sprintf("%x\n", bs)

	sb := userStruct.SelectFrom("users")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "user_tokens", "users.id = user_tokens.user_id")
	sb.Where(sb.Equal("user_tokens.hash", tokenHash))
	sb.Where(sb.GreaterThan("user_tokens.expires_at", time.Now()))
	q, args := sb.Build()

	row := db.QueryRow(q, args...)
	user := model.User{}
	err := row.Scan(userStruct.Addr(&user)...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}
	return &user
}

func CreateUser(username string, email string) (string, error) {
	id := node.Generate().String()
	q, args := sqlbuilder.InsertInto("users").
		Cols("id", "username", "email").
		Values(id, username, email).
		Build()

	_, err := db.Exec(q, args...)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return "", ErrUserAlreadyExists
		}
		panic(err)
	}
	return id, nil
}

func CreateTokenForUser(id string) (string, error) {
	b := make([]byte, 32)
	rand.Read(b)
	token := fmt.Sprintf("%x", b)
	h := sha256.New()
	h.Write([]byte(token))
	bs := h.Sum(nil)
	tokenHash := fmt.Sprintf("%x\n", bs)

	q, args := sqlbuilder.InsertInto("user_tokens").
		Cols("hash", "user_id", "expires_at").
		Values(tokenHash, id, time.Now().Add(24*time.Hour)).
		Build()

	_, err := db.Exec(q, args...)
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint \"user_tokens_user_id_fkey\"") {
			return "", ErrUserDoesNotExist
		}
		panic(err)
	}
	return token, nil
}

func ExpireTokensForUser(id string) {
	sb := sqlbuilder.NewUpdateBuilder().Update("user_tokens")
	sb.Set(sb.Equal("expires_at", time.Now()))
	sb.Where(sb.GreaterThan("expires_at", time.Now()))
	sb.Where(sb.Equal("user_id", id))

	q, args := sb.Build()
	_, err := db.Exec(q, args...)
	if err != nil {
		panic(err)
	}
}
