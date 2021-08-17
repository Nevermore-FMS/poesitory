package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Nevermore-FMS/poesitory/backend/database"
	"github.com/Nevermore-FMS/poesitory/backend/graph/generated"
	"github.com/Nevermore-FMS/poesitory/backend/graph/model"
)

func (r *mutatePluginPayloadResolver) Plugin(ctx context.Context, obj *model.MutatePluginPayload) (*model.NevermorePlugin, error) {
	return database.GetPluginByID(obj.PluginID), nil
}

// MutatePluginPayload returns generated.MutatePluginPayloadResolver implementation.
func (r *Resolver) MutatePluginPayload() generated.MutatePluginPayloadResolver {
	return &mutatePluginPayloadResolver{r}
}

type mutatePluginPayloadResolver struct{ *Resolver }
