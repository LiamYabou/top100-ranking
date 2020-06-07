package api

import (
	"fmt"
	"context"
	"encoding/json"
	"github.com/LiamYabou/top100-pkg/logger"
	"github.com/LiamYabou/top100-ranking/preference"
)

func encode(obj *response) (jsonResonse string) {
	jsonObj, err := json.Marshal(obj)
	if err != nil {
		logger.Error("Failed to marshal an object.", err)
	}
	jsonResonse = string(jsonObj)
	return jsonResonse
}

func FindRankings(categoryId int, page int, opts *preference.Options) string {
	// args validation
	if categoryId == 0 {
		response := &response{
			Status: "fail",
			Data: &content{"category_id": "The category id is invaild, it should be greater than zero."},
		}
		return encode(response)
	}
	if page == 0 || page > 2 {
		response := &response{
			Status: "fail",
			Data: &content{"page": "The page number is invaild, it should either be 1 or 2."},
		}
		return encode(response)
	}
	// Fetch products from DB
	stmt := fmt.Sprintf("select name, rank from products where category_id = %d and page = %d order by rank asc", categoryId, page)
	rows, err := opts.DB.Query(context.Background(), stmt)
	if err != nil {
		response := &response{
			Status: "error",
			Message: fmt.Sprintf("Failed to query on DB with the statement: %s. Error: %s", stmt, err),
		}
		return encode(response)
	}
	defer rows.Close()
	productSet := make([]*productRow, 0)
	for rows.Next() {
		row := &productRow{}
		err = rows.Scan(&row.Name, &row.Rank)
		if err != nil {
			response := &response{
				Status: "error",
				Message: fmt.Sprintf("Failed to assign a value by the Scan. Error: %s", err),
			}
			return encode(response)
		}
		productSet = append(productSet, row)
	}
	// Get any error encountered during iteration.
	if err := rows.Err(); err != nil {
		response := &response{
			Status: "error",
			Message: fmt.Sprintf("The errors were encountered during the iteration. Error: %s", err),
		}
		return encode(response)
	}
	// Fetch categories from DB
	stmt = fmt.Sprintf("select name, url from categories where parent_id = %d order by path asc", categoryId)
	rows, err = opts.DB.Query(context.Background(), stmt)
	if err != nil {
		response := &response{
			Status: "error",
			Message: fmt.Sprintf("Failed to query on DB with the statement: %s. Error: %s", stmt, err),
		}
		return encode(response)
	}
	defer rows.Close()
	categorySet := make([]*categoryRow, 0)
	for rows.Next() {
		row := &categoryRow{}
		err = rows.Scan(&row.Name, &row.Url)
		if err != nil {
			response := &response{
				Status: "error",
				Message: fmt.Sprintf("Failed to assign a value by the Scan. Error: %s", err),
			}
			return encode(response)
		}
		categorySet = append(categorySet, row)
	}
	// Get any error encountered during iteration.
	if err := rows.Err(); err != nil {
		response := &response{
			Status: "error",
			Message: fmt.Sprintf("The errors were encountered during the iteration. Error: %s", err),
		}
		return encode(response)
	}
	var selectedCategoryEntry string
	stmt = fmt.Sprintf("select name from categories where id = %d", categoryId)
	err = opts.DB.QueryRow(context.Background(), stmt).Scan(&selectedCategoryEntry)
	if err != nil && err.Error() == ErrNoRows.Error() {
		response := &response{
			Status: "success",
			Data: nil,
		}
		return encode(response)
	} else if err != nil {
		response := &response{
			Status: "error",
			Message: fmt.Sprintf("Failed to query on DB with the statement: %s. Error: %s", stmt, err),
		}
		return encode(response)
	}
	var data *content
	data = &content{
		"products": productSet,
		"categories": categorySet,
		"selected_category_entry": selectedCategoryEntry,
	}
	response := &response{
		Status: "success",
		Data: data,
	}
	return encode(response)
}
