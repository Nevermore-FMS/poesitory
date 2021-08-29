package comm

import (
	"context"
	"errors"

	"github.com/shurcooL/graphql"
)

type PluginVersion struct {
	Id          graphql.ID
	Plugin      *Plugin
	DownloadUrl *graphql.String
}

func GetPluginVersion(identifier string) (*PluginVersion, error) {
	var query struct {
		PluginVersion *PluginVersion `graphql:"pluginVersion(versionIdentifier: $identifier)"`
	}
	variables := map[string]interface{}{
		"identifier": graphql.String(identifier),
	}
	err := client.Query(context.Background(), &query, variables)
	return query.PluginVersion, err
}

func UploadPluginVersion(id, version, channel string) (string, error) {
	var mutation struct {
		UploadPluginVersion *struct {
			Url string
		} `graphql:"uploadPluginVersion(id: $id, version: $version, channel: $channel)"`
	}
	variables := map[string]interface{}{
		"id":      graphql.ID(id),
		"version": graphql.String(version),
		"channel": graphql.String(channel),
	}
	err := client.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		return "", err
	}
	if mutation.UploadPluginVersion == nil {
		return "", errors.New("unknown error (uploadPluginVersion not present in api response)")
	}
	return mutation.UploadPluginVersion.Url, nil
}
