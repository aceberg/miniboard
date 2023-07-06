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
	guiData.Links = AllLinks

	filter := r.FormValue("filter")
	panel := r.FormValue("panel")
	host := r.FormValue("host")
	state := r.FormValue("state")

	if filter == "yes" {
		guiData.UptimeMon = filterUptimeMon(panel, host, state) // uptime-filter.go
	} else {
		guiData.UptimeMon = UptimeMon
	}

	sort.Slice(guiData.UptimeMon, func(i, j int) bool {
		return guiData.UptimeMon[i].Date > guiData.UptimeMon[j].Date
	})

	execTemplate(w, "uptime", guiData)
}
