package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Nevermore-FMS/poesitory/backend/auth"
	"github.com/Nevermore-FMS/poesitory/backend/database"
	"github.com/Nevermore-FMS/poesitory/backend/graph/generated"
	"github.com/Nevermore-FMS/poesitory/backend/graph/model"
)

func (r *mutationResolver) CreatePlugin(ctx context.Context, name string) (*model.MutatePluginPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UploadPluginVersion(ctx context.Context, id string, version string, channel string) (*model.UploadPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	return auth.UserForContext(ctx), nil
}

func (r *queryResolver) SearchPlugins(ctx context.Context, search *string, page *int) (*model.NevermorePluginPage, error) {
	plugins := database.GetPlugins(*search, *page)
	return &model.NevermorePluginPage{
		PageNum: *page,
		HasNext: false, // TODO
		Plugins: plugins,
	}, nil
}

func (r *queryResolver) PluginVersion(ctx context.Context, identifier string) (*model.NevermorePluginVersion, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
