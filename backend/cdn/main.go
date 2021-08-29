package cdn

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var client *minio.Client

func init() {
	endpoint := envFallback("DEV_CDN_URI", os.Getenv("POESITORY_CDN_URI"))
	secure := os.Getenv("POESITORY_HTTPS") == "true"
	accessKeyID := envFallback("DEV_CDN_USER", "poesitory")
	secretAccessKey := envFallback("DEV_CDN_KEY", os.Getenv("POESITORY_SECRET"))

	var transport http.RoundTripper
	if os.Getenv("POESITORY_DEV_INSECURE") == "true" {
		log.Println("WARNING: Running in DEV INSECURE mode")
		transport = &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		}
	}

	var err error
	client, err = minio.New(endpoint, &minio.Options{
		Creds:     credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure:    secure,
		Transport: transport,
	})
	if err != nil {
		log.Fatal(err)
	}

	initBuckets()
}

func initBuckets() {
	found, err := client.BucketExists(context.Background(), "pdl")
	if err != nil {
		log.Fatal(err)
	}
	if !found {
		err = client.MakeBucket(context.Background(), "pdl", minio.MakeBucketOptions{})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func envFallback(env string, fallback string) string {
	result := os.Getenv(env)
	if len(result) == 0 {
		result = fallback
	}
	return result
}
