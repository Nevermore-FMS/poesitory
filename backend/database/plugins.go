package database

import (
	"crypto/sha256"
	"database/sql"
	"fmt"

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

func GetPlugins(search string, pluginType *model.NevermorePluginType, owner *string, page int) (_ []*model.NevermorePlugin, hasNext bool) {
	sb := pluginStruct.SelectFrom("plugins")

	if len(search) > 0 {
		sb.Where(fmt.Sprintf("name @@ %s", sb.Var(search)))
	}
	if pluginType != nil {
		sb.Where(sb.Equal("type", pluginType))
	}
	if owner != nil {
		sb.Where(sb.Equal("owner", owner))
	}

	q, args := sb.Build()
	rows, err := db.Query(q, args...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	plugins := make([]*model.NevermorePlugin, 0)
	for rows.Next() {
		plugin := model.NevermorePlugin{}
		rows.Scan(pluginStruct.Addr(&plugin)...)
		plugins = append(plugins, &plugin)
	}

	sb = sqlbuilder.NewSelectBuilder().From("plugins")
	sb.Select("count(*) > 0")
	sb.Limit(1).Offset(page * pluginsPerPage)
	q, args = sb.Build()
	row := db.QueryRow(q, args...)
	err = row.Scan(&hasNext)
	if err != nil {
		if err == sql.ErrNoRows {
			hasNext = false
		} else {
			panic(err)
		}
	}
	defer rows.Close()

	return plugins, hasNext
}

func GetPluginByToken(token string) *model.NevermorePlugin {
	h := sha256.New()
	h.Write([]byte(token))
	bs := h.Sum(nil)
	tokenHash := fmt.Sprintf("%x\n", bs)

	sb := pluginStruct.SelectFrom("plugins")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "upload_tokens", "plugins.id = upload_tokens.plugin_id")
	sb.Where(sb.Equal("upload_tokens.hash", tokenHash))
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
