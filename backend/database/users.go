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

	sb := sqlbuilder.NewSelectBuilder().Select("user_id").From("user_tokens")
	sb.Where(sb.Equal("hash", tokenHash))
	sb.Where(sb.GreaterThan("expires_at", time.Now()))
	q, args := sb.Build()

	row := db.QueryRow(q, args...)
	var userId string
	err := row.Scan(&userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}
	return GetUserByID(userId)
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
