package web

import (
	// "log"
	"net/http"

	"github.com/aceberg/miniboard/internal/models"
	"github.com/aceberg/miniboard/internal/yaml"
)

func tabHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	key := r.FormValue("key")
	name := r.FormValue("name")
	panels := r.PostForm["panels"]

	remove := r.URL.Query().Get("remove")

	if key != "" && name != "" {
		_, exists := AllLinks.Tabs[key]

		if !exists {
			tab := models.Tab{}
			tab.Name = name
			tab.Panels = panels
			AllLinks.Tabs[key] = tab

			yaml.Write(AppConfig.YamlPath, AllLinks)
		}
	} else if remove != "" {
		delete(AllLinks.Tabs, remove)
		yaml.Write(AppConfig.YamlPath, AllLinks)
	}

	guiData.Config = AppConfig
	guiData.CurrentTab = "Edit Tabs"
	guiData.Links.Panels = AllLinks.Panels
	guiData.Links.Tabs = AllLinks.Tabs

	execTemplate(w, "tabs", guiData)
}
