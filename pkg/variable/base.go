package variable

// The place that you can share the various variables accross the whole project.

import (
	"os"
)

var (
	Env =  os.Getenv("ENV")
	TestDBURL  = os.Getenv("TEST_DB_DSN")
	FixturesURI = os.Getenv("FIXTURES_URI")
)
