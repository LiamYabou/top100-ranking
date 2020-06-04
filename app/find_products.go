package app

import (
	"fmt"
	"context"
	"encoding/json"
	"github.com/LiamYabou/top100-pkg/logger"
	"github.com/LiamYabou/top100-ranking/preference"
)

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
	defer Finalize()
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