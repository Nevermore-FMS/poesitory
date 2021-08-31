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
		return nil, errors.New("plugin does not exist on Poesitory")
	}
	return query.Plugin, nil
}
