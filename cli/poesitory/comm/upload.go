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
)

func UploadUrl(url string, data io.Reader) error {
	resp, err := http.Post(url, "application/gzip", data)
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

func CreateTarGz(basePath string, include []string) (io.Reader, error) {
	if _, err := os.Stat(basePath); err != nil {
		return nil, fmt.Errorf("unable to tar files - %v", err.Error())
	}

	var writer *bytes.Buffer = &bytes.Buffer{}

	gzw := gzip.NewWriter(writer)
	tw := tar.NewWriter(gzw)

	for _, filePath := range include {
		path := filepath.Join(basePath, filePath)
		fi, err := os.Stat(path)
		if err != nil {
			return nil, err
		}
		if !fi.Mode().IsRegular() {
			return nil, fmt.Errorf("path %s is not a file", path)
		}
		header, err := tar.FileInfoHeader(fi, fi.Name())
		if err != nil {
			return nil, err
		}
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		header.Name = filePath
		if err := tw.WriteHeader(header); err != nil {
			return nil, err
		}
		if _, err := io.Copy(tw, f); err != nil {
			return nil, err
		}
		f.Close()
	}

	tw.Close()
	gzw.Close()
	return writer, nil
}
