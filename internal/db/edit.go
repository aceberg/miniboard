package db

import (
	"fmt"

	"github.com/aceberg/miniboard/internal/models"
)

// Create - create table if not exists
func Create(path string) {

	sqlStatement := `CREATE TABLE IF NOT EXISTS records (
		"ID"		INTEGER PRIMARY KEY,
		"PANEL"		TEXT,
		"HOST"		TEXT,
		"ADDR"		TEXT,
		"PORT"		TEXT,
		"DATE"		TEXT,
		"TIME"		TEXT,
		"COLOR"		TEXT,
		"STATE"		INTEGER,
		"NOTIFY"	TEXT
	);`
	exec(path, sqlStatement)
}

// Insert - insert one rec into DB
func Insert(path string, rec models.MonData) {
	var intState int

	sqlStatement := `INSERT INTO records (PANEL, HOST, ADDR, PORT, DATE, TIME, COLOR, STATE, NOTIFY) 
	VALUES ('%s','%s','%s','%s','%s','%s','%s','%d','%s');`

	rec.Panel = quoteStr(rec.Panel)
	rec.Host = quoteStr(rec.Host)

	if rec.State {
		intState = 1
	}

	sqlStatement = fmt.Sprintf(sqlStatement, rec.Panel, rec.Host, rec.Addr, rec.Port, rec.Date, rec.Time, rec.Color, intState, rec.Notify)

	exec(path, sqlStatement)
}

// Delete - delete one record
func Delete(path string, id int) {

	sqlStatement := `DELETE FROM records WHERE ID='%d';`

	sqlStatement = fmt.Sprintf(sqlStatement, id)

	exec(path, sqlStatement)
}

// Clear - delete all records from table
func Clear(path string) {
	sqlStatement := `DELETE FROM records;`
	exec(path, sqlStatement)
}
