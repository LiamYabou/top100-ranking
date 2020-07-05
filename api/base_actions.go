package api

import (
	"errors"
	"database/sql"
	"encoding/json"
)

type categoryRow struct {
	Id       int `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Url     string `json:"url,omitempty"`
	Path     string `json:"path,omitempty"`
	ParentID int `json:"parent_id,omitempty"`
}

type productRow struct {
	ID         int `json:"id,omitempty"`
	Name       string `json:"name"`
	Rank       int `json:"rank"`
	Page       int `json:"page,omitempty"`
	CategoryID int `json:"category_id,omitempty"`
	ImagePath NullString `json:"image_path,omitempty"`
}

// NullString is an alias for sql.NullString data type
type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

type content map[string]interface{}

type response struct {
	Status  string `json:"status"`
	Data    *content `json:"data"`
	Message string `json:"message,omitempty"`
	Code    int `json:"code,omitempty"`
}

var ErrNoRows = errors.New("no rows in result set")
