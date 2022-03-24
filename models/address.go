package models

type Address struct {
	Id      int64   `json:"id,omitempty"`
	UserId  int64   `json:"user_id,omitempty"`
	Address *string `json:"address,omitempty"`
}
