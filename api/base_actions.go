package api

import "errors"

type categoryRow struct {
	Id       int `json:"id,omitempty"`
	Name     string `json:"name"`
	Url     string `json:"url"`
	Path     string `json:"path,omitempty"`
	ParentID int `json:"parent_id,omitempty"`
}

type productRow struct {
	ID         int `json:"id,omitempty"`
	Name       string `json:"name"`
	Rank       int `json:"rank"`
	Page       int `json:"page,omitempty"`
	CategoryID int `json:"category_id,omitempty"`
}

type content map[string]interface{}

type response struct {
	Status  string `json:"status"`
	Data    *content `json:"data"`
	Message string `json:"message,omitempty"`
	Code    int `json:"code,omitempty"`
}

var ErrNoRows = errors.New("no rows in result set")
