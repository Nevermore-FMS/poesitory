package comm

import (
	"context"
	"errors"

	"github.com/shurcooL/graphql"
)

type Plugin struct {
	Id   graphql.ID
	Name graphql.String
}

func GetPlugin(name string) (*Plugin, error) {
	var query struct {
		Plugin *Plugin `graphql:"plugin(name: $name)"`
	}
	variables := map[string]interface{}{
		"name": graphql.String(name),
	}
	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, err
	}
	if query.Plugin == nil {
		return nil, errors.New("unknown error (plugin not present in api response)")
	}
	return query.Plugin, nil
}
