package app_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/LiamYabou/top100-ranking/preference"
	"github.com/LiamYabou/top100-ranking/test"
	"github.com/LiamYabou/top100-ranking/app"
)


func (a *actionSuite) TestFindProducts() {
	expect := test.CanndedJsonProducts
	categoryId := 2
	page := 1
	opts := preference.LoadOptions(preference.WithDB(test.DBpool))
	actual := app.FindProducts(categoryId, page, opts)
	failedMsg := fmt.Sprintf("Failed, expected the result: %v, got the result: %v", expect, actual)
	assert.Equal(a.T(), expect, actual, failedMsg)
} 