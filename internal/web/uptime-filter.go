package web

import (
	// "log"
	"net/http"
	"sort"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/models"
)

func uptimeFilterHandler(w http.ResponseWriter, r *http.Request) {
	var resultUptimeMon []models.MonData
	var guiData models.GuiData
	guiData.Config = AppConfig
	guiData.CurrentTab = "Uptime Monitor"
	guiData.Links = AllLinks

	panel := r.FormValue("panel")
	host := r.FormValue("host")
	state := r.FormValue("state")
	date := r.FormValue("date")
	addr := r.FormValue("addr")
	port := r.FormValue("port")
	notify := r.FormValue("notify")

	for _, mon := range UptimeMon {

		if panel == mon.Panel && (host == "" || (host != "" && host == mon.Host)) {
			resultUptimeMon = append(resultUptimeMon, mon)
		}

		if (mon.State && state == "on") || (!mon.State && state == "off") {
			resultUptimeMon = append(resultUptimeMon, mon)
		}

		if mon.Date == date || mon.Addr == addr || mon.Port == port {
			resultUptimeMon = append(resultUptimeMon, mon)
		}

		if check.InSlice(mon.Notify, notify) {
			resultUptimeMon = append(resultUptimeMon, mon)
		}
	}

	guiData.UptimeMon = resultUptimeMon

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
