package web

import (
	// "log"
	"net/http"
	"strconv"

	"github.com/aceberg/miniboard/internal/models"
	"github.com/aceberg/miniboard/internal/yaml"
)

func tabsHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	idStr := r.FormValue("id")
	name := r.FormValue("name")
	refresh := r.FormValue("refresh")
	panels := r.PostForm["panels"]

	upStr := r.FormValue("up")

	id := -1
	if upStr != "" {
		id, _ = strconv.Atoi(upStr)
		tab := AllLinks.Tabs[id-1]
		AllLinks.Tabs[id-1] = AllLinks.Tabs[id]
		AllLinks.Tabs[id] = tab

	} else if name != "" {

		tab := models.Tab{}
		tab.Name = name
		tab.Refresh = refresh

		if idStr == "" {
			id = len(AllLinks.Tabs)
			tab.Panels = make(map[int]string)

			for i, panel := range panels {
				tab.Panels[i] = panel
			}

		} else {
			id, _ = strconv.Atoi(idStr)
			tab.Panels = AllLinks.Tabs[id].Panels
		}

		AllLinks.Tabs[id] = tab

	}
	if id > -1 {
		yaml.Write(AppConfig.YamlPath, AllLinks)
	}

	guiData.Config = AppConfig
	guiData.CurrentTab = "Edit Tabs"
	guiData.Links = AllLinks

	execTemplate(w, "tabs", guiData)
}
