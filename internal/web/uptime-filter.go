package web

import (
	// "log"

	"github.com/aceberg/miniboard/internal/models"
)

func filterUptimeMon(panel, host, state string) []models.MonData {
	var resultUptimeMon []models.MonData

	if panel != "" {
		for _, mon := range UptimeMon {
			if panel == mon.Panel {
				resultUptimeMon = append(resultUptimeMon, mon)
			}
		}

		if host != "" {
			newUptimeMon := []models.MonData{}

			for _, mon := range resultUptimeMon {
				if host == mon.Host {
					newUptimeMon = append(newUptimeMon, mon)
				}
			}
			resultUptimeMon = newUptimeMon
		}
	}

	if state == "on" {
		for _, mon := range UptimeMon {
			if mon.State {
				resultUptimeMon = append(resultUptimeMon, mon)
			}
		}
	}

	if state == "off" {
		for _, mon := range UptimeMon {
			if !mon.State {
				resultUptimeMon = append(resultUptimeMon, mon)
			}
		}
	}

	return resultUptimeMon
}
