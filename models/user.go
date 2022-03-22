package models

type User struct {
	Name      string     `json:"name,omitempty"`
	Lastname  string     `json:"lastname,omitempty"`
	Addresses []*Address `json:"addresses,omitempty"`
}
