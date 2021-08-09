package database

import (
	"database/sql"

	"github.com/Nevermore-FMS/poesitory/backend/graph/model"
	"github.com/huandu/go-sqlbuilder"
)

var pluginVersionStruct = sqlbuilder.NewStruct(new(model.NevermorePluginVersion)).For(sqlbuilder.PostgreSQL)

func GetPluginVersion(id string) *model.NevermorePluginVersion {
	sb := pluginVersionStruct.SelectFrom("plugin_versions")
	sb.Where(sb.Equal("id", id))

	q, args := sb.Build()
	row := db.QueryRow(q, args...)
	pv := model.NevermorePluginVersion{}
	err := row.Scan(pluginStruct.Addr(&pv)...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return &pv
}

func GetPluginVersionsForPlugin(pluginID string, channel string) []*model.NevermorePluginVersion {
	sb := pluginVersionStruct.SelectFrom("plugin_versions")
	sb.Where(sb.Equal("plugin", pluginID))
	sb.Where(sb.Equal("channel", channel))
	sb.Desc().OrderBy("major", "minor", "patch")
	sb.Limit(50)

	q, args := sb.Build()
	rows, err := db.Query(q, args...)
	if err != nil {
		panic(err)
	}

	pluginVersions := make([]*model.NevermorePluginVersion, 0)
	for rows.Next() {
		pv := model.NevermorePluginVersion{}
		rows.Scan(pluginStruct.Addr(&pv)...)
		pluginVersions = append(pluginVersions, &pv)
	}

	return pluginVersions
}

func GetLatestPluginVersionForPlugin(pluginID string, channel string) *model.NevermorePluginVersion {
	sb := pluginVersionStruct.SelectFrom("plugin_versions")
	sb.Where(sb.Equal("plugin", pluginID))
	sb.Where(sb.Equal("channel", channel))
	sb.Desc().OrderBy("major", "minor", "patch")
	sb.Limit(1)

	q, args := sb.Build()
	row := db.QueryRow(q, args...)
	pv := model.NevermorePluginVersion{}
	err := row.Scan(pluginStruct.Addr(&pv)...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return &pv
}
