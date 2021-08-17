package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/Nevermore-FMS/poesitory/backend/auth"
	"github.com/Nevermore-FMS/poesitory/backend/cdn"
	"github.com/Nevermore-FMS/poesitory/backend/database"
	"github.com/Nevermore-FMS/poesitory/backend/graph/generated"
	"github.com/Nevermore-FMS/poesitory/backend/graph/model"
	"github.com/Nevermore-FMS/poesitory/backend/identifier"
)

func (r *nevermorePluginResolver) Owner(ctx context.Context, obj *model.NevermorePlugin) (*model.User, error) {
	return database.GetUserByID(obj.OwnerID), nil
}

func (r *nevermorePluginResolver) LatestFullIdentifier(ctx context.Context, obj *model.NevermorePlugin) (*string, error) {
	latestVersion := database.GetLatestPluginVersionForPlugin(obj.ID, "STABLE")
	if latestVersion == nil {
		return nil, nil
	}
	pluginVersion := identifier.PluginSemVer{
		Major: latestVersion.Major,
		Minor: latestVersion.Minor,
		Patch: latestVersion.Patch,
	}
	idfr := identifier.ConstructIdentifierString(identifier.PluginVersionIdentifier{
		Name:    obj.Name,
		Version: &pluginVersion,
		Channel: "STABLE",
	}, nil)
	return &idfr, nil
}

func (r *nevermorePluginResolver) LatestVersion(ctx context.Context, obj *model.NevermorePlugin) (*model.NevermorePluginVersion, error) {
	return database.GetLatestPluginVersionForPlugin(obj.ID, "STABLE"), nil
}

func (r *nevermorePluginResolver) Channels(ctx context.Context, obj *model.NevermorePlugin) ([]*model.NevermorePluginChannel, error) {
	return database.GetChannelsForPlugin(obj.ID), nil
}

func (r *nevermorePluginResolver) UploadTokens(ctx context.Context, obj *model.NevermorePlugin) ([]*model.UploadToken, error) {
	user := auth.UserForContext(ctx)
	if user == nil || obj.OwnerID != user.ID {
		return nil, auth.ErrNoPermissions
	}
	return database.GetUploadTokensForPlugin(obj.ID), nil
}

func (r *nevermorePluginChannelResolver) Plugin(ctx context.Context, obj *model.NevermorePluginChannel) (*model.NevermorePlugin, error) {
	return database.GetPluginByID(obj.PluginID), nil
}

func (r *nevermorePluginChannelResolver) Versions(ctx context.Context, obj *model.NevermorePluginChannel) ([]*model.NevermorePluginVersion, error) {
	return database.GetPluginVersionsForPlugin(obj.PluginID, obj.Name), nil
}

func (r *nevermorePluginVersionResolver) Plugin(ctx context.Context, obj *model.NevermorePluginVersion) (*model.NevermorePlugin, error) {
	return database.GetPluginByID(obj.PluginID), nil
}

func (r *nevermorePluginVersionResolver) Channel(ctx context.Context, obj *model.NevermorePluginVersion) (*model.NevermorePluginChannel, error) {
	return &model.NevermorePluginChannel{
		PluginID: obj.PluginID,
		Name:     obj.ChannelStr,
	}, nil
}

func (r *nevermorePluginVersionResolver) ShortIdentifier(ctx context.Context, obj *model.NevermorePluginVersion) (string, error) {
	plugin := database.GetPluginByID(obj.PluginID)
	if plugin == nil {
		return "", errors.New("plugin does not exist")
	}
	idfr := identifier.PluginVersionIdentifier{
		Name:    plugin.Name,
		Channel: obj.ChannelStr,
		Version: &identifier.PluginSemVer{
			Major: obj.Major,
			Minor: obj.Minor,
			Patch: obj.Patch,
		},
	}
	latestPluginVersion := database.GetLatestPluginVersionForPlugin(obj.PluginID, obj.ChannelStr)
	return identifier.ConstructIdentifierString(idfr, &identifier.PluginSemVer{
		Major: latestPluginVersion.Major,
		Minor: latestPluginVersion.Minor,
		Patch: latestPluginVersion.Patch,
	}), nil
}

func (r *nevermorePluginVersionResolver) FullIdentifier(ctx context.Context, obj *model.NevermorePluginVersion) (string, error) {
	plugin := database.GetPluginByID(obj.PluginID)
	if plugin == nil {
		return "", errors.New("plugin does not exist")
	}
	idfr := identifier.PluginVersionIdentifier{
		Name:    plugin.Name,
		Channel: obj.ChannelStr,
		Version: &identifier.PluginSemVer{
			Major: obj.Major,
			Minor: obj.Minor,
			Patch: obj.Patch,
		},
	}
	return identifier.ConstructIdentifierString(idfr, nil), nil
}

func (r *nevermorePluginVersionResolver) DownloadURL(ctx context.Context, obj *model.NevermorePluginVersion) (string, error) {
	return cdn.GenDownloadUrl(obj.Hash), nil
}

// NevermorePlugin returns generated.NevermorePluginResolver implementation.
func (r *Resolver) NevermorePlugin() generated.NevermorePluginResolver {
	return &nevermorePluginResolver{r}
}

// NevermorePluginChannel returns generated.NevermorePluginChannelResolver implementation.
func (r *Resolver) NevermorePluginChannel() generated.NevermorePluginChannelResolver {
	return &nevermorePluginChannelResolver{r}
}

// NevermorePluginVersion returns generated.NevermorePluginVersionResolver implementation.
func (r *Resolver) NevermorePluginVersion() generated.NevermorePluginVersionResolver {
	return &nevermorePluginVersionResolver{r}
}

type nevermorePluginResolver struct{ *Resolver }
type nevermorePluginChannelResolver struct{ *Resolver }
type nevermorePluginVersionResolver struct{ *Resolver }
