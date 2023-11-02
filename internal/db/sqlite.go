package db

import (
	"sync"

	"github.com/jmoiron/sqlx"

	// Import module for SQLite DB
	_ "modernc.org/sqlite"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/models"
)

var mu sync.Mutex

func connect(path string) *sqlx.DB {
	dbx, err := sqlx.Connect("sqlite", path)
	check.IfError(err)

	return dbx
}

func exec(path string, sqlStatement string) {

	mu.Lock()
	dbx := connect(path)
	_, err := dbx.Exec(sqlStatement)
	mu.Unlock()

	check.IfError(err)
}

// Select - select all from DB
func Select(path string) []models.MonData {
	var recs []models.MonData

	mu.Lock()
	dbx := connect(path)
	err := dbx.Select(&recs, "SELECT * FROM records ORDER BY ID DESC")
	mu.Unlock()

	check.IfError(err)

	return recs
}
