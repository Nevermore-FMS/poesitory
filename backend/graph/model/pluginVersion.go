package model

type NevermorePluginVersion struct {
	ID         string  `json:"id"`
	PluginID   string  `db:"plugin"`
	ChannelStr string  `db:"channel"`
	Major      int     `db:"major"`
	Minor      int     `db:"minor"`
	Patch      int     `db:"patch"`
	Readme     *string `db:"readme" json:"readme"`
}
