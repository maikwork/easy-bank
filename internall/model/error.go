package model

type ErrorHTTP struct {
	Err  string `json:"err"`
	Desc string `json:"description"`
}
