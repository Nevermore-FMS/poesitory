package auth

import (
	"errors"
	"fmt"
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
	selfUri := envFallback("POESITORY_BASE_URI", "localhost:8080")
	var err error
	if os.Getenv("POESITORY_HTTPS") == "true" {
		SelfUrl, err = url.Parse(fmt.Sprintf("https://%s", selfUri))
	} else {
		SelfUrl, err = url.Parse(fmt.Sprintf("https://%s", selfUri))
	}
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
