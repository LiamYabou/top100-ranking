package action

import (
	"fmt"
	"context"
	"encoding/json"
	"github.com/LiamYabou/top100-ranking/pkg/logger"
	"github.com/LiamYabou/top100-ranking/pkg/app"
	"github.com/LiamYabou/top100-ranking/pkg/preference"
)

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
	Data    *content `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Code    int `json:"code,omitempty"`
}

func FindProducts(categoryId int, page int, opts *preference.Options) string {
	// args validation
	if categoryId == 0 {
		response := &response{
			Status: "fail",
			Data: &content{"category_id": "The category id is invaild, it should be greater than zero."},
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			logger.Error("Failed to marshal an object.", err)
		}
		return string(jsonResponse)
	}
	if page == 0 || page > 2 {
		response := &response{
			Status: "fail",
			Data: &content{"page": "The page number is invaild, it should either be 1 or 2."},
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			logger.Error("Failed to marshal an object.", err)
		}
		return string(jsonResponse)
	}
	defer app.Finalize()
	stmt := fmt.Sprintf("select name, rank from products where category_id = %d and page = %d order by rank asc", categoryId, page)
	rows, err := opts.DB.Query(context.Background(), stmt)
	if err != nil {
		response := &response{
			Status: "error",
			Message: fmt.Sprintf("Failed to query on DB. Error: %s", err),
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			logger.Error("Failed to marshal an object.", err)
		}
		return string(jsonResponse)
	}
	defer rows.Close()
	set := make([]*productRow, 0)
	for rows.Next() {
		row := &productRow{}
		err = rows.Scan(&row.Name, &row.Rank)
		if err != nil {
			response := &response{
				Status: "error",
				Message: fmt.Sprintf("Failed to assign a value by the Scan. Error: %s", err),
			}
			jsonResponse, err := json.Marshal(response)
			if err != nil {
				logger.Error("Failed to marshal an object.", err)
			}
			return string(jsonResponse)
		}
		set = append(set, row)
	}
	// Get any error encountered during iteration.
	if err := rows.Err(); err != nil {
		response := &response{
			Status: "error",
			Message: fmt.Sprintf("The errors were encountered during the iteration. Error: %s", err),
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			logger.Error("Failed to marshal an object.", err)
		}
		return string(jsonResponse)
	}
	response := &response{
		Status: "success",
		Data: &content{"products": set},
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Error("Failed to marshal an object.", err)
	}
	return string(jsonResponse)
}