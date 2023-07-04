package web

import (
	// "log"

	"github.com/aceberg/miniboard/internal/models"
)

func filterUptimeMon(panel, host, state string) []models.MonData {
	var resultUptimeMon []models.MonData

	if panel != "" {
		for _, mon := range UptimeMon {
			if panel == mon.Panel && (host == "" || (host != "" && host == mon.Host)) {
				resultUptimeMon = append(resultUptimeMon, mon)
			}
		}
	}

	if state != "" {
		for _, mon := range UptimeMon {
			if (mon.State && state == "on") || (!mon.State && state == "off") {
				resultUptimeMon = append(resultUptimeMon, mon)
			}
		}
	}

	return resultUptimeMon
}
