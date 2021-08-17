package model

type NevermorePlugin struct {
	ID      string              `db:"id" json:"id"`
	Name    string              `db:"name" json:"name"`
	Type    NevermorePluginType `db:"type" json:"type"`
	OwnerID string              `db:"owner"`
}

type MutatePluginPayload struct {
	Successful bool `json:"successful"`
	PluginID   string
}

type NevermorePluginChannel struct {
	PluginID string
	Name     string `json:"name"`
}

type NevermorePluginPage struct {
	PageNum int                `json:"pageNum"`
	HasNext bool               `json:"hasNext"`
	Plugins []*NevermorePlugin `json:"plugins"`
}

type UploadPayload struct {
	URL string `json:"url"`
}
