package web

import (
	// "log"
	"strconv"
	"time"

	"github.com/aceberg/miniboard/internal/db"
)

func trimDBRoutine() {

	for {
		trimDB()

		time.Sleep(time.Duration(1) * time.Hour) // Every hour
	}
}

func trimDB() {

	days, _ := strconv.Atoi(AppConfig.DBTrimDays)

	nowStr := time.Now().Format("2006-01-02")  // This needed so all time is
	now, _ := time.Parse("2006-01-02", nowStr) // in one format
	nowMinus := now.Add(-time.Duration(days) * 24 * time.Hour)

	history := db.Select(AppConfig.DBPath)

	for _, hist := range history {
		date, _ := time.Parse("2006-01-02", hist.Date)

		if date.Before(nowMinus) {

			// log.Println("TRIM:", hist)
			db.Delete(AppConfig.DBPath, hist.ID)
		}
	}

	UptimeMon = db.Select(AppConfig.DBPath)
}
