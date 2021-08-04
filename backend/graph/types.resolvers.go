package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Nevermore-FMS/poesitory/backend/graph/generated"
	"github.com/Nevermore-FMS/poesitory/backend/graph/model"
)

func (r *nevermorePluginResolver) Owner(ctx context.Context, obj *model.NevermorePlugin) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *nevermorePluginResolver) LatestShortIdentifier(ctx context.Context, obj *model.NevermorePlugin) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *nevermorePluginResolver) LatestFullIdentifier(ctx context.Context, obj *model.NevermorePlugin) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *nevermorePluginResolver) LatestVersion(ctx context.Context, obj *model.NevermorePlugin) (*model.NevermorePluginVersion, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *nevermorePluginResolver) Channels(ctx context.Context, obj *model.NevermorePlugin) ([]*model.NevermorePluginChannel, error) {
	panic(fmt.Errorf("not implemented"))
}

// NevermorePlugin returns generated.NevermorePluginResolver implementation.
func (r *Resolver) NevermorePlugin() generated.NevermorePluginResolver {
	return &nevermorePluginResolver{r}
}

type nevermorePluginResolver struct{ *Resolver }
