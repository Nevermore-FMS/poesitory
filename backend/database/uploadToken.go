package database

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"

	"github.com/Nevermore-FMS/poesitory/backend/graph/model"
	"github.com/huandu/go-sqlbuilder"
)

var uploadTokenStruct = sqlbuilder.NewStruct(new(model.UploadToken)).For(sqlbuilder.PostgreSQL)

func GetUploadTokensForPlugin(pluginID string) []*model.UploadToken {
	sb := uploadTokenStruct.SelectFrom("upload_tokens")
	sb.Where(sb.Equal("plugin_id", pluginID))

	q, args := sb.Build()
	rows, err := db.Query(q, args...)
	if err != nil {
		panic(err)
	}
	uploadTokens := make([]*model.UploadToken, 0)
	for rows.Next() {
		uploadToken := model.UploadToken{}
		rows.Scan(uploadTokenStruct.Addr(&uploadToken)...)
		uploadTokens = append(uploadTokens, &uploadToken)
	}
	return uploadTokens
}

func GetUploadTokenByID(id string) *model.UploadToken {
	sb := uploadTokenStruct.SelectFrom("upload_tokens")
	sb.Where(sb.Equal("id", id))

	q, args := sb.Build()
	row := db.QueryRow(q, args...)
	uploadToken := model.UploadToken{}
	err := row.Scan(uploadTokenStruct.Addr(&uploadToken)...)
	if err != nil {
		panic(err)
	}
	return &uploadToken
}

func CreateUploadToken(pluginID string) (*string, error) {
	b := make([]byte, 32)
	rand.Read(b)
	token := fmt.Sprintf("%x", b)
	h := sha256.New()
	h.Write([]byte(token))
	bs := h.Sum(nil)
	tokenHash := fmt.Sprintf("%x\n", bs)

	id := node.Generate().String()
	q, args := sqlbuilder.InsertInto("upload_tokens").
		Cols("id", "hash", "plugin_id").
		Values(id, tokenHash, pluginID).
		Build()

	_, err := db.Exec(q, args...)
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return nil, errors.New("plugin does not exist")
		}
		panic(err)
	}
	return &token, nil
}

func DeleteUploadToken(id string) {
	sb := sqlbuilder.DeleteFrom("upload_tokens")
	sb.Where(sb.Equal("id", id))

	q, args := sb.Build()
	db.Exec(q, args...)
}
