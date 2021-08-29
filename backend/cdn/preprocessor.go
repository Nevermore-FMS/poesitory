package cdn

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/Nevermore-FMS/poesitory/backend/database"
	"github.com/Nevermore-FMS/poesitory/backend/graph/model"
	"github.com/Nevermore-FMS/poesitory/backend/identifier"
	"github.com/google/uuid"
)

type ExpectedPlugin struct {
	Uploaded   bool
	ID         string
	Name       string
	PluginType model.NevermorePluginType
	Version    identifier.PluginSemVer
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
		uuid := strings.Split(r.URL.Path, "/")[3]
		expectedPlugin, ok := expectedPlugins[uuid]
		if !ok {
			http.NotFound(w, r)
			return
		}
		pluginUploadHandler(expectedPlugin, uuid).ServeHTTP(w, r)
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

func pluginUploadHandler(expectedPlugin ExpectedPlugin, uuid string) http.Handler {
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

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		r.Body.Close()

		gzf, err := gzip.NewReader(bytes.NewReader(b))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		defer gzf.Close()

		tarReader := tar.NewReader(gzf)

		var nevermoreJson nevermoreJson
		var packageJson packageJson
		var readme string

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

			if strings.HasSuffix(header.Name, "nevermore.json") {
				json.NewDecoder(tarReader).Decode(&nevermoreJson)
			}

			if strings.HasSuffix(header.Name, "package.json") {
				json.NewDecoder(tarReader).Decode(&packageJson)
			}
			if strings.Contains(header.Name, "README") {
				buf := &strings.Builder{}
				_, err := io.Copy(buf, tarReader)
				if err != nil {
					panic(err)
				}
				readme = buf.String()
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
		semVer, err := identifier.ParseVersion(packageJson.Version)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Version parsing error: %s", err.Error())))
			return
		}
		if semVer != expectedPlugin.Version {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Version in package.json does not match expected version"))
			return
		}
		if nevermoreJson.PluginType != expectedPlugin.PluginType {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Plugin type in nevermore.json does not match plugin type on poesitory"))
			return
		}

		hash := Upload(b)
		_, err = database.CreatePluginVersion(expectedPlugin.ID, hash, expectedPlugin.Version.Major, expectedPlugin.Version.Minor, expectedPlugin.Version.Patch, expectedPlugin.Channel, readme)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			panic(err)
		}

		delete(expectedPlugins, uuid)
		w.Write([]byte("OK"))
	})
}
