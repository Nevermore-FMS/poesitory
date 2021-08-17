package cdn

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Nevermore-FMS/poesitory/backend/graph/model"
	"github.com/google/uuid"
)

type ExpectedPlugin struct {
	ID         string
	Name       string
	PluginType model.NevermorePluginType
	Version    string
	Channel    string
}

type nevermoreJson struct {
	Name        string                    `json:"name"`
	Author      string                    `json:"author"`
	Email       string                    `json:"email"`
	Url         string                    `json:"url"`
	PluginType  model.NevermorePluginType `json:"pluginType"`
	Permissions []string                  `json:"permissions"`
}

type packageJson struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

var expectedPlugins = make(map[string]ExpectedPlugin)

func UploadHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPlugin, ok := expectedPlugins[strings.Split(r.URL.Path, "/")[3]]
		if !ok {
			http.NotFound(w, r)
			return
		}
		pluginUploadHandler(expectedPlugin).ServeHTTP(w, r)
	})
}

func AddExpectedPlugin(ep ExpectedPlugin) string {
	uuid := uuid.New().String()
	expectedPlugins[uuid] = ep

	time.AfterFunc(1*time.Minute, func() {
		delete(expectedPlugins, uuid)
	})

	return uuid
}

func pluginUploadHandler(expectedPlugin ExpectedPlugin) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Only POST requests are supported"))
			return
		}
		if r.Header.Get("Content-Type") != "application/gzip" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("Content must be of type \"application/gzip\""))
			return
		}

		gzf, err := gzip.NewReader(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		tarReader := tar.NewReader(gzf)

		var nevermoreJson nevermoreJson
		var packageJson packageJson

		for {
			header, err := tarReader.Next()

			if err != nil {
				if err == io.EOF {
					break
				}
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			if header.Name == "./nevermore.json" {
				json.NewDecoder(tarReader).Decode(&nevermoreJson)
			}

			if header.Name == "./package.json" {
				json.NewDecoder(tarReader).Decode(&packageJson)
			}
		}

		if packageJson.Name != expectedPlugin.Name {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Name in package.json does not match name on poesitory"))
			return
		}
		if nevermoreJson.Name != expectedPlugin.Name {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Name in nevermore.json does not match name on poesitory"))
			return
		}
		if packageJson.Version != expectedPlugin.Version {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Version in package.json does not match expected version"))
			return
		}
		if nevermoreJson.PluginType != expectedPlugin.PluginType {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Plugin type in nevermore.json does not match plugin type on poesitory"))
			return
		}

		//TODO Upload file and update DB

		w.Write([]byte("OK"))
	})
}
