package cdn

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
)

func Upload(f []byte) string {
	hash := sha256.New()
	if _, err := io.Copy(hash, bytes.NewReader(f)); err != nil {
		panic(err)
	}
	sum := fmt.Sprintf("%x", hash.Sum(nil))

	client.PutObject(context.Background(), "pdl", fmt.Sprintf("%s.tar.gz", sum), bytes.NewReader(f), int64(len(f)), minio.PutObjectOptions{ContentType: "application/gzip"})

	return sum
}
