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

	guiData.UptimeMon = UptimeMon

	sort.Slice(guiData.UptimeMon, func(i, j int) bool {
		return guiData.UptimeMon[i].Time > guiData.UptimeMon[j].Time
	})

	if AllLinks.Uptime.Show < 1 {
		AllLinks.Uptime.Show = 20
	}
	if len(guiData.UptimeMon) > AllLinks.Uptime.Show {
		guiData.UptimeMon = guiData.UptimeMon[0:AllLinks.Uptime.Show]
	}

	execTemplate(w, "uptime", guiData)
}
