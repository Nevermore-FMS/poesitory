package local

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type NevermorePluginType string

const (
	NevermorePluginTypeGeneric             NevermorePluginType = "GENERIC"
	NevermorePluginTypeGame                NevermorePluginType = "GAME"
	NevermorePluginTypeNetworkConfigurator NevermorePluginType = "NETWORK_CONFIGURATOR"
)

type NevermoreJson struct {
	Name        string              `json:"name"`
	Author      string              `json:"author"`
	Email       string              `json:"email"`
	Url         string              `json:"url"`
	PluginType  NevermorePluginType `json:"pluginType"`
	Permissions []string            `json:"permissions"`
}

type PackageJson struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func ReadNevermoreJson(basePath string) (nevermoreJson *NevermoreJson, err error) {
	nevermoreJson = &NevermoreJson{}
	data, err := ioutil.ReadFile(filepath.Join(basePath, "./nevermore.json"))
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, nevermoreJson)
	return nevermoreJson, nil
}

func ReadPackageJson(basePath string) (packageJson *PackageJson, err error) {
	packageJson = &PackageJson{}
	data, err := ioutil.ReadFile(filepath.Join(basePath, "./package.json"))
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, packageJson)
	return packageJson, nil
}
