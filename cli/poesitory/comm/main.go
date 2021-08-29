package comm

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/shurcooL/graphql"
)

var client *graphql.Client
var AuthorizationHeader string

type headerTransport struct {
	underlyingTransport http.RoundTripper
}

func (t *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", AuthorizationHeader)
	return t.underlyingTransport.RoundTrip(req)
}

func init() {
	var httpClient *http.Client = &http.Client{}
	headerTransport := &headerTransport{underlyingTransport: http.DefaultTransport}
	if os.Getenv("POESITORY_DEV_INSECURE") == "true" {
		log.Println("WARNING: Running in DEV INSECURE mode")
		http.DefaultTransport = &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		}
		headerTransport.underlyingTransport = http.DefaultTransport
	}
	httpClient.Transport = headerTransport
	client = graphql.NewClient(envFallback("POESITORY_ENDPOINT", "https://poesitory.nathankutzan.info/api/graphql"), httpClient)
}

func envFallback(env, fallback string) string {
	val, ok := os.LookupEnv(env)
	if ok {
		return val
	}
	return fallback
}
