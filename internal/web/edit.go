package web

import (
	"log"
	"net/http"

	"github.com/aceberg/miniboard/internal/models"
)

func editHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig

	panelName := r.URL.Query().Get("panel")
	new := r.URL.Query().Get("new")

	guiData.Links = models.Links{}
	for _, panel := range AllLinks.Panels {
		if panel.Name == panelName {
			if new == "host" {
				panel.Hosts = append(panel.Hosts, models.Host{})
			}
			guiData.Links.Panels = append(guiData.Links.Panels, panel)
		}
	}

	execTemplate(w, "edit", guiData)
}

func saveEditHandler(w http.ResponseWriter, r *http.Request) {
	var host models.Host
	// var newPanels []models.Panel

	panelName := r.FormValue("panel")
	host.Name = r.FormValue("name")
	host.Addr = r.FormValue("addr")
	host.Port = r.FormValue("port")
	host.Icon = r.FormValue("icon")

	log.Println("PANEL:", panelName)
	log.Println("HOST:", host)

	// for _, panel := range AllLinks.Panels {
	// 	if panel.Name == panelName {
	// 		for _,
	// 	} else {
	// 		newPanels = append(newPanels, panel)
	// 	}
	// }

	// AllLinks.Panels = newPanels

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
