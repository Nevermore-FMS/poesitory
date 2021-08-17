package cdn

import (
	"context"
	"log"
	"net/url"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var client *minio.Client

func init() {
	endpoint, secure := decodeUrl(envFallback("DEV_CDN_URI", os.Getenv("POESITORY_CDN_URI")))
	accessKeyID := envFallback("DEV_CDN_USER", "poesitory")
	secretAccessKey := envFallback("DEV_CDN_KEY", os.Getenv("POESITORY_SECRET"))

	var err error
	client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: secure,
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

func decodeUrl(urlStr string) (endpoint string, secure bool) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		log.Fatal(err)
	}
	secure = parsedURL.Scheme == "https"
	endpoint = parsedURL.Host + parsedURL.Path
	return
}

func envFallback(env string, fallback string) string {
	result := os.Getenv(env)
	if len(result) == 0 {
		result = fallback
	}
	return result
}
