package web

import (
	"net/http"
	"strconv"

	"github.com/aceberg/miniboard/internal/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	guiData.Config = AppConfig

	reload := r.URL.Query().Get("reload")
	if reload == "yes" {

		reloadScans() // webgui.go

		http.Redirect(w, r, "/", 302)
	}

	tabStr := r.URL.Query().Get("tab")
	var tab int
	if tabStr == "" {
		tab = 0
	} else {
		tab, _ = strconv.Atoi(tabStr)
	}

	guiData.CurrentTab = AllLinks.Tabs[tab].Name
	guiData.Links = AllLinks

	if AllLinks.Tabs[tab].Refresh != "" && AllLinks.Tabs[tab].Refresh != "0" {
		guiData.Config.WebRefresh = AllLinks.Tabs[tab].Refresh
	}

	guiData.Panels = make(map[int]models.Panel)
	for i, name := range AllLinks.Tabs[tab].Panels {
		guiData.Panels[i] = AllLinks.Panels[name]
	}

	execTemplate(w, "index", guiData)
}
