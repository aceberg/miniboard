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

	if filter == "yes" {
		guiData.UptimeMon = filterUptime(r)
	} else {
		guiData.UptimeMon = UptimeMon
	}

	sort.SliceStable(guiData.UptimeMon, func(i, j int) bool {
        mi, mj := guiData.UptimeMon[i], guiData.UptimeMon[j]
        switch {
        case mi.Date != mj.Date:
            return mi.Date > mj.Date
        default:
            return mi.Time > mj.Time
        }
    })

	if AllLinks.Uptime.Show < 1 {
		AllLinks.Uptime.Show = 20
	}
	if len(guiData.UptimeMon) > AllLinks.Uptime.Show {
		guiData.UptimeMon = guiData.UptimeMon[0:AllLinks.Uptime.Show]
	}

	execTemplate(w, "uptime", guiData)
}
