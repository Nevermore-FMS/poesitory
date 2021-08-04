package auth

import (
	"net/url"
	"os"
)

type Token struct {
	AccessToken string `json:"accessToken"`
	ExpiresAt   int    `json:"expiresAt"`
}

var SelfUrl *url.URL

func Init() {
	var err error
	SelfUrl, err = url.Parse(envFallback("POESITORY_BASE_URL", "http://localhost:8080"))
	if err != nil {
		panic(err)
	}
	initGithub()
}

func envFallback(env string, fallback string) string {
	result := os.Getenv("POESITORY_BASE_URL")
	if len(result) == 0 {
		result = fallback
	}
	return result
}
