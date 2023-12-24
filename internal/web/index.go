package web

import (
	"net/http"
	"strconv"

	"github.com/aceberg/miniboard/internal/auth"
	"github.com/aceberg/miniboard/internal/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	guiData.Config = AppConfig

	tabStr := r.URL.Query().Get("tab")
	var tab int
	if tabStr == "" {
		tab = 0
	} else {
		tab, _ = strconv.Atoi(tabStr)
	}

	AppConfig.LoggedIn = auth.IsLoggedIn(w, r)
	if AppConfig.Auth && !AppConfig.LoggedIn && AllLinks.Tabs[tab].Auth {
		http.Redirect(w, r, "/login/", 302)
	}

	guiData.CurrentTab = AllLinks.Tabs[tab].Name
	guiData.Links = AllLinks
	guiData.Horiz = AllLinks.Tabs[tab].Horiz

	if AllLinks.Tabs[tab].Refresh != "" && AllLinks.Tabs[tab].Refresh != "0" {
		guiData.Config.WebRefresh = AllLinks.Tabs[tab].Refresh
	}

	guiData.Panels = make(map[int]models.Panel)
	for i, name := range AllLinks.Tabs[tab].Panels {
		guiData.Panels[i] = AllLinks.Panels[name]
	}

	execTemplate(w, "index", guiData)
}

func reloadHandler(w http.ResponseWriter, r *http.Request) {

	reloadScans() // webgui.go

	http.Redirect(w, r, "/", 302)
}
