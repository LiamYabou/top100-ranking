package test

import (
	"fmt"
	"context"
	"gopkg.in/khaiql/dbcleaner.v2"
	"gopkg.in/khaiql/dbcleaner.v2/engine"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/LiamYabou/top100-ranking/variable"
	"github.com/LiamYabou/top100-pkg/db"
)

func InitDB() (msg string, err error) {
	DBpool, err = db.OpenTest()
	if err != nil {
		return "Failed to connect the DB", err
	}
	PQconn, err = db.OpenPQtest()
	return "", err
}

func InitCleaner() {
	Cleaner = dbcleaner.New()
	psql := engine.NewPostgresEngine(variable.TestDBURL)
	Cleaner.SetEngine(psql)
}

// InitTable is used to truncated the table, and restart the identity of the table.
func InitTable(name string, db *pgxpool.Pool) error {
	stmt := fmt.Sprintf("truncate table %s restart identity cascade", name)
	_, err := db.Exec(context.Background(), stmt)
	return err
}