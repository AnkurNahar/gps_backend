package models

type Preferences struct {
	SortBy            string            `json:"sort_by"`
	HiddenDeviceIds   []string          `json:"hidden_device_ids"`
	UserDeviceIcons   map[string]string `json:"user_device_icons"`
}