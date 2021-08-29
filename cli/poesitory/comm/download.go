package comm

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadUrl(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func ExtractTarGz(gzipStream io.Reader, basePath string) error {
	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		return err
	}
	defer uncompressedStream.Close()

	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(filepath.Join(basePath, header.Name), 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			err := os.MkdirAll(filepath.Dir(filepath.Join(basePath, header.Name)), 0755)
			if err != nil {
				return err
			}
			outFile, err := os.Create(filepath.Join(basePath, header.Name))
			if err != nil {
				return err
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return err
			}
			outFile.Close()

		default:
			log.Fatalf(
				"ExtractTarGz: uknown type: %b in %s",
				header.Typeflag,
				header.Name)
		}

	}
	return nil
}
