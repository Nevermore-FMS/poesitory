package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/Nevermore-FMS/poesitory/backend/auth"
	"github.com/Nevermore-FMS/poesitory/backend/cdn"
	"github.com/Nevermore-FMS/poesitory/backend/database"
	"github.com/Nevermore-FMS/poesitory/backend/graph/generated"
	"github.com/Nevermore-FMS/poesitory/backend/graph/model"
	"github.com/Nevermore-FMS/poesitory/backend/identifier"
)

func (r *mutationResolver) CreatePlugin(ctx context.Context, name string, typeArg model.NevermorePluginType) (*model.MutatePluginPayload, error) {
	user := auth.UserForContext(ctx)
	if user == nil {
		return &model.MutatePluginPayload{
			Successful: false,
		}, auth.ErrNoPermissions
	}
	if ok, _ := regexp.MatchString("^[a-z-]+$", name); !ok {
		return nil, errors.New("invalid name (only lowercase letters and '-' allowed)")
	}
	id, err := database.CreatePlugin(name, typeArg, user.ID)
	if err != nil {
		return &model.MutatePluginPayload{
			Successful: false,
		}, err
	}
	return &model.MutatePluginPayload{
		Successful: true,
		PluginID:   id,
	}, nil
}

func (r *mutationResolver) UploadPluginVersion(ctx context.Context, id string, version string, channel string) (*model.UploadPayload, error) {
	plugin := database.GetPluginByID(id)
	if plugin == nil {
		return nil, errors.New("plugin does not exist")
	}
	user := auth.UserForContext(ctx)
	userAuthed := user != nil && plugin.OwnerID == user.ID
	tPl := auth.PluginForContext(ctx)
	pluginAuthed := tPl != nil && plugin.ID == tPl.ID
	if !pluginAuthed && !userAuthed {
		return nil, auth.ErrNoPermissions
	}

	if ok, _ := regexp.MatchString("^[A-Z-]+$", channel); !ok {
		return nil, errors.New("invalid channel (only uppercase letters and '-' allowed)")
	}

	semVer, err := identifier.ParseVersion(version)
	if err != nil {
		return nil, err
	}
	latestSemVer := database.GetLatestPluginVersionForPlugin(id, channel)
	if latestSemVer != nil {
		isNewVersion := false
		if semVer.Major > latestSemVer.Major {
			isNewVersion = true
		} else {
			if semVer.Minor > latestSemVer.Minor {
				isNewVersion = true
			} else {
				if semVer.Patch > latestSemVer.Patch {
					isNewVersion = true
				}
			}
		}
		if !isNewVersion {
			return nil, errors.New("version must be greater than the latest version in desired channel")
		}
	}

	uploadToken := cdn.AddExpectedPlugin(cdn.ExpectedPlugin{
		ID:         id,
		Name:       plugin.Name,
		PluginType: plugin.Type,
		Version:    semVer,
		Channel:    channel,
	})
	return &model.UploadPayload{
		URL: fmt.Sprintf("%s/api/upload/%s", auth.SelfUrl, uploadToken),
	}, nil
}

func (r *mutationResolver) CreateUploadToken(ctx context.Context, pluginID string) (*string, error) {
	plugin := database.GetPluginByID(pluginID)
	if plugin == nil {
		return nil, errors.New("plugin does not exist")
	}
	user := auth.UserForContext(ctx)
	if user == nil || plugin.OwnerID != user.ID {
		return nil, auth.ErrNoPermissions
	}
	return database.CreateUploadToken(pluginID)
}

func (r *mutationResolver) DeleteUploadToken(ctx context.Context, id string) (*model.MutatePluginPayload, error) {
	uploadToken := database.GetUploadTokenByID(id)
	if uploadToken == nil {
		return nil, errors.New("upload token does not exist")
	}
	plugin := database.GetPluginByID(uploadToken.PluginID)
	user := auth.UserForContext(ctx)
	if user == nil || plugin.OwnerID != user.ID {
		return nil, auth.ErrNoPermissions
	}
	database.DeleteUploadToken(id)
	return &model.MutatePluginPayload{
		Successful: true,
		PluginID:   plugin.ID,
	}, nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	return auth.UserForContext(ctx), nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return database.GetUserByID(id), nil
}

func (r *queryResolver) SearchPlugins(ctx context.Context, search *string, typeArg *model.NevermorePluginType, owner *string, page *int) (*model.NevermorePluginPage, error) {
	plugins, hasNext := database.GetPlugins(*search, typeArg, owner, *page)
	return &model.NevermorePluginPage{
		PageNum: *page,
		HasNext: hasNext,
		Plugins: plugins,
	}, nil
}

func (r *queryResolver) PluginVersion(ctx context.Context, versionIdentifier string) (*model.NevermorePluginVersion, error) {
	pluginIdentifier, err := identifier.ParseStringIdentifier(versionIdentifier)
	if err != nil {
		return nil, err
	}
	plugin := database.GetPluginByName(pluginIdentifier.Name)
	if plugin == nil {
		return nil, errors.New("plugin does not exist")
	}
	if pluginIdentifier.Version == nil {
		return database.GetLatestPluginVersionForPlugin(plugin.ID, pluginIdentifier.Channel), nil
	} else {
		return database.GetPluginVersionForPlugin(plugin.ID, pluginIdentifier.Channel, pluginIdentifier.Version.Major, pluginIdentifier.Version.Minor, pluginIdentifier.Version.Patch), nil
	}
}

func (r *queryResolver) Plugin(ctx context.Context, id *string, name *string) (*model.NevermorePlugin, error) {
	if id != nil {
		return database.GetPluginByID(*id), nil
	}
	if name != nil {
		return database.GetPluginByName(*name), nil
	}
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
