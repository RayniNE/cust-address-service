package models

type User struct {
	Id        int64      `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Lastname  string     `json:"lastname,omitempty"`
	Addresses []*Address `json:"addresses,omitempty"`
}
