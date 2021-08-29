package database

import (
	"database/sql"
	"errors"
	"strings"

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
	err := row.Scan(pluginVersionStruct.Addr(&pv)...)
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
		rows.Scan(pluginVersionStruct.Addr(&pv)...)
		pluginVersions = append(pluginVersions, &pv)
	}

	return pluginVersions
}

func GetPluginVersionForPlugin(pluginID string, channel string, major, minor, patch int) *model.NevermorePluginVersion {
	sb := pluginVersionStruct.SelectFrom("plugin_versions")
	sb.Where(sb.Equal("plugin", pluginID))
	sb.Where(sb.Equal("channel", channel))
	sb.Where(sb.Equal("major", major))
	sb.Where(sb.Equal("minor", minor))
	sb.Where(sb.Equal("patch", patch))

	q, args := sb.Build()
	row := db.QueryRow(q, args...)
	pv := model.NevermorePluginVersion{}
	err := row.Scan(pluginVersionStruct.Addr(&pv)...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return &pv
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
	err := row.Scan(pluginVersionStruct.Addr(&pv)...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return &pv
}

func CreatePluginVersion(pluginID string, hash string, major, minor, patch int, channel string, readme string) (string, error) {
	id := node.Generate().String()
	q, args := sqlbuilder.InsertInto("plugin_versions").
		Cols("id", "plugin", "hash", "major", "minor", "patch", "channel", "readme").
		Values(id, pluginID, hash, major, minor, patch, channel, readme).
		Build()

	_, err := db.Exec(q, args...)
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return "", errors.New("plugin does not exist")
		}
		panic(err)
	}

	return id, nil
}
