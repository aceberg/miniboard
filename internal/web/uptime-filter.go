package web

import (
	// "log"
	"net/http"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/models"
)

func filterUptime(r *http.Request) []models.MonData {
	var resultUptimeMon []models.MonData

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

	return resultUptimeMon
}
