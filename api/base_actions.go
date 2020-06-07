package api

type categoryRow struct {
	id       int
	name     string
	url      string
	path     string
	parentID int
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
