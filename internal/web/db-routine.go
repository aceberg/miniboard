package web

import (
	// "log"
	"time"

	"github.com/aceberg/miniboard/internal/db"
	"github.com/aceberg/miniboard/internal/models"
)

func dbRoutine() {
	var newUptimeMon []models.MonData
	var mon models.MonData

	db.Create(AppConfig.DBPath)
	UptimeMon = db.Select(AppConfig.DBPath)

	for {
		newUptimeMon = db.Select(AppConfig.DBPath)

		for _, mon = range UptimeMon {

			if !inSlice(newUptimeMon, mon) {
				// log.Println("NEW RECORD TO DB:", mon)
				db.Insert(AppConfig.DBPath, mon)
			}
		}

		time.Sleep(time.Duration(3*60) * time.Second)
	}
}

func inSlice(monSlice []models.MonData, elem models.MonData) bool {
	var mon models.MonData

	elem.ID = 0

	for _, mon = range monSlice {
		mon.ID = 0
		if elem == mon {
			return true
		}
	}

	return false
}
