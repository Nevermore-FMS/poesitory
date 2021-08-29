package comm

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func UploadUrl(url string, data []byte) error {
	resp, err := http.Post(url, "application/gzip", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		errorMsg, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(errorMsg))
	}
	return nil
}

func CreateTarGz(basePath string) ([]byte, error) {
	if _, err := os.Stat(basePath); err != nil {
		return nil, fmt.Errorf("unable to tar files - %v", err.Error())
	}

	var writer *bytes.Buffer = &bytes.Buffer{}

	gzw := gzip.NewWriter(writer)
	tw := tar.NewWriter(gzw)

	err := filepath.Walk(basePath, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fi.Mode().IsRegular() {
			return nil
		}

		header, err := tar.FileInfoHeader(fi, fi.Name())
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(strings.Replace(file, basePath, "", -1), string(filepath.Separator))

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		f, err := os.Open(file)
		if err != nil {
			return err
		}

		if _, err := io.Copy(tw, f); err != nil {
			return err
		}

		f.Close() //defering would cause each file close to wait until all operations have completed.

		return nil
	})
	if err != nil {
		return nil, err
	}
	tw.Close()
	gzw.Close()
	return writer.Bytes(), nil
}
