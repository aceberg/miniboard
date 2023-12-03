package web

import (
	// "log"
	// "runtime"
	"sync"
	"time"

	"github.com/aceberg/miniboard/internal/db"
	"github.com/aceberg/miniboard/internal/models"
)

var newUptimeMutex sync.Mutex

func dbRoutine() {
	var mon models.MonData
	var tmpUptimeMon []models.MonData

	timeout := 1

	db.Create(AppConfig.DBPath)
	UptimeMon = db.Select(AppConfig.DBPath)

	for {
		newUptimeMutex.Lock()
		tmpUptimeMon = NewUptimeMon
		NewUptimeMon = []models.MonData{}
		newUptimeMutex.Unlock()

		for _, mon = range tmpUptimeMon {
			db.Insert(AppConfig.DBPath, mon)
		}

		// // Analyzing goroutine leaks
		// var stats runtime.MemStats
		// runtime.ReadMemStats(&stats)
		// log.Printf("Number of Goroutines: %d\n", runtime.NumGoroutine())

		time.Sleep(time.Duration(timeout) * 60 * time.Second)
	}
}
