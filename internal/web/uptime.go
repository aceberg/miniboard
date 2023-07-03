package web

import (
	// "log"
	"net/http"
	"sort"

	"github.com/aceberg/miniboard/internal/models"
)

func uptimeHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	guiData.Config = AppConfig
	guiData.CurrentTab = "Uptime Monitor"
	guiData.Links.Tabs = AllLinks.Tabs

	panel := r.FormValue("panel")
	state := r.FormValue("state")

	if panel != "" {
		for _, mon := range UptimeMon {
			if panel == mon.Panel {
				guiData.UptimeMon = append(guiData.UptimeMon, mon)
			}
		}

		host := r.FormValue("host")
		if host != "" {
			newUptimeMon := []models.MonData{}

			for _, mon := range guiData.UptimeMon {
				if host == mon.Host {
					newUptimeMon = append(newUptimeMon, mon)
				}
			}
			guiData.UptimeMon = newUptimeMon
		}
	} else if state == "on" {
		for _, mon := range UptimeMon {
			if mon.State {
				guiData.UptimeMon = append(guiData.UptimeMon, mon)
			}
		}
	} else if state == "off" {
		for _, mon := range UptimeMon {
			if !mon.State {
				guiData.UptimeMon = append(guiData.UptimeMon, mon)
			}
		}
	} else {
		guiData.UptimeMon = UptimeMon
	}

	sort.Slice(guiData.UptimeMon, func(i, j int) bool {
		return guiData.UptimeMon[i].Date > guiData.UptimeMon[j].Date
	})

	execTemplate(w, "uptime", guiData)
}
