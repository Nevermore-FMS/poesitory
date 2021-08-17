package auth

import (
	"errors"
	"log"
	"net/url"
	"os"
)

var ErrNoPermissions = errors.New("lacking permissions to perform this action")

type Token struct {
	AccessToken string `json:"accessToken"`
	ExpiresAt   int    `json:"expiresAt"`
}

var SelfUrl *url.URL

func init() {
	var err error
	SelfUrl, err = url.Parse(envFallback("POESITORY_BASE_URL", "http://localhost:8080"))
	if err != nil {
		log.Fatal(err)
	}
	initGithub()
}

func envFallback(env string, fallback string) string {
	result := os.Getenv(env)
	if len(result) == 0 {
		result = fallback
	}
	return result
}
