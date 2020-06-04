package api_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/LiamYabou/top100-ranking/preference"
	"github.com/LiamYabou/top100-ranking/test"
	"github.com/LiamYabou/top100-ranking/api"
)

func (a *actionSuite) TestFindProducts() {
	assert := assert.New(a.T())
	page := 1
	categoryId := 2
	opts := &preference.Options{
		DB: test.DBpool,
		RunTimeEnv: "test",
	}
	opts = preference.LoadOptions(preference.WithOptions(*opts))
	// # Find products
	// ## Standard procedure
	expected := test.CanndedJsonProducts
	actual := api.FindProducts(categoryId, page, opts)
	failedMsg := fmt.Sprintf("Failed, expected the result: %v, got the result: %v", expected, actual)
	assert.Equal(expected, actual, failedMsg)
	// ## Empty result
	categoryId = 1
	expected = `{"status":"success","data":null}`
	actual = api.FindProducts(categoryId, page, opts)
	failedMsg = fmt.Sprintf("Failed, expected the result: %v, got the result: %v", expected, actual)
	assert.Equal(expected, actual, failedMsg)
	// ## category = 0
	categoryId = 0
	expected = `{"status":"fail","data":{"category_id":"The category id is invaild, it should be greater than zero."}}`
	actual = api.FindProducts(categoryId, page, opts)
	failedMsg = fmt.Sprintf("Failed, expected the result: %v, got the result: %v", expected, actual)
	assert.Equal(expected, actual, failedMsg)
	// ## page = 0
	categoryId = 2
	page = 0
	expected = `{"status":"fail","data":{"page":"The page number is invaild, it should either be 1 or 2."}}`
	actual = api.FindProducts(categoryId, page, opts)
	failedMsg = fmt.Sprintf("Failed, expected the result: %v, got the result: %v", expected, actual)
	assert.Equal(expected, actual, failedMsg)
	// ## page > 2
	page = 3
	expected = `{"status":"fail","data":{"page":"The page number is invaild, it should either be 1 or 2."}}`
	actual = api.FindProducts(categoryId, page, opts)
	failedMsg = fmt.Sprintf("Failed, expected the result: %v, got the result: %v", expected, actual)
	assert.Equal(expected, actual, failedMsg)
} 