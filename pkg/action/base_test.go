package action_test

import (
	"os"
	"fmt"
	"testing"
	"github.com/romanyx/polluter"
	"github.com/LiamYabou/top100-ranking/pkg/variable"
	"github.com/LiamYabou/top100-ranking/pkg/test"
	"github.com/stretchr/testify/suite"
)

type actionSuite struct {
	suite.Suite
}

// Run before the tests in the suite are run.
func (a *actionSuite) SetupSuite() {
	// Initialize the DB
	msg, err := test.InitDB()
	if err != nil {
		a.T().Errorf("%s, error: %v", msg, err)
	}
	// Initialize the dbcleaner
	test.InitCleaner()
}

// Run before each test in the suite.
func (a *actionSuite) SetupTest() {
	test.Cleaner.Acquire("products", "categories")
	// Populate the data into the tables
	seedPath := fmt.Sprintf("%s/data.yml", variable.FixturesURI)
	seed, err := os.Open(seedPath)
	if err != nil {
		a.T().Errorf("Failed to opent the seed, error: %v", err)
	}
	defer seed.Close()
	poluter := polluter.New(polluter.PostgresEngine(test.PQconn))
	if err := poluter.Pollute(seed); err != nil {
		a.T().Errorf("Failed to pollute the seed, error: %v", err)
	}
}

// Run after each test in the suite.
func (a *actionSuite) TearDownTest() {
	test.Cleaner.Clean("products", "categories", "product_categories")
}

// Run after all the tests in the suite have been run.
func (a *actionSuite) TearDownSuite() {
	test.Finalize()
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(actionSuite))
}