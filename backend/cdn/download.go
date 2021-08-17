package cdn

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

func GenDownloadUrl(hash string) string {
	presignedURL, err := client.PresignedGetObject(context.Background(), "pdl", fmt.Sprintf("%s.tar.gz", hash), 5*time.Minute, url.Values{})
	if err != nil {
		panic(err)
	}
	return presignedURL.String()
}
