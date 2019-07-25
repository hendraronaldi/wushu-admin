package model

type Course struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Materials interface{} `json:"materials"`
}
