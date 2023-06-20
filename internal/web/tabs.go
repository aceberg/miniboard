package web

import (
	"log"
	"net/http"
	"strconv"

	"github.com/aceberg/miniboard/internal/models"
	// "github.com/aceberg/miniboard/internal/yaml"
)

func tabsHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	idStr := r.FormValue("id")
	name := r.FormValue("name")
	panels := r.PostForm["panels"]

	delStr := r.URL.Query().Get("del")
	upStr := r.FormValue("up")

	id := -1
	if delStr != "" {
		id, _ = strconv.Atoi(delStr)
		delete(AllLinks.Tabs, id)
		// assign ids

	} else if name != "" {

		if idStr == "" {
			id = len(AllLinks.Tabs)
		} else {
			id, _ = strconv.Atoi(idStr)
		}
		tab := models.Tab{}
		tab.Name = name
		tab.Panels = panels
		AllLinks.Tabs[id] = tab

	} else if upStr != "" {
		id, _ = strconv.Atoi(upStr)
		tab := AllLinks.Tabs[id-1]
		AllLinks.Tabs[id-1] = AllLinks.Tabs[id]
		AllLinks.Tabs[id] = tab
	}
	if id > -1 {
		log.Println("WRITE")
		// yaml.Write(AppConfig.YamlPath, AllLinks)
	}

	guiData.Config = AppConfig
	guiData.CurrentTab = "Edit Tabs"
	guiData.Links = AllLinks

	execTemplate(w, "tabs", guiData)
}
