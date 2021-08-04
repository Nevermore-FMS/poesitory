package database

import (
	"database/sql"

	"github.com/Nevermore-FMS/poesitory/backend/graph/model"
	"github.com/huandu/go-sqlbuilder"
)

var pluginStruct = sqlbuilder.NewStruct(new(model.NevermorePlugin)).For(sqlbuilder.PostgreSQL)

func GetPluginByID(id string) *model.NevermorePlugin {
	sb := pluginStruct.SelectFrom("plugins")
	sb.Where(sb.Equal("id", id))

	q, args := sb.Build()
	row := db.QueryRow(q, args...)
	plugin := model.NevermorePlugin{}
	err := row.Scan(pluginStruct.Addr(&plugin)...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}
	return &plugin
}

func GetPluginByName(name string) *model.NevermorePlugin {
	sb := pluginStruct.SelectFrom("plugins")
	sb.Where(sb.Equal("name", name))

	q, args := sb.Build()
	row := db.QueryRow(q, args...)
	plugin := model.NevermorePlugin{}
	err := row.Scan(pluginStruct.Addr(&plugin)...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}
	return &plugin
}

var pluginsPerPage = 20

func GetPlugins(search string, page int) []*model.NevermorePlugin {
	sb := pluginStruct.SelectFrom("plugins")
	sb.OrderBy("name").Desc()
	sb.Limit(pluginsPerPage).Offset((page - 1) * pluginsPerPage)

	q, args := sb.Build()
	rows, err := db.Query(q, args...)
	if err != nil {
		panic(err)
	}
	plugins := make([]*model.NevermorePlugin, 0)
	for rows.Next() {
		plugin := model.NevermorePlugin{}
		rows.Scan(pluginStruct.Addr(&plugin)...)
		plugins = append(plugins, &plugin)
	}
	return plugins
}
