package model

import "time"

type UploadToken struct {
	ID            string    `db:"id" json:"id"`
	CreatedAtTime time.Time `db:"created_at"`
	PluginID      string    `db:"plugin_id"`
}

func (s *UploadToken) CreatedAt() int64 {
	return s.CreatedAtTime.Unix()
}
