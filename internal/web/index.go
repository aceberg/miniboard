package web

import (
	"net/http"

	"github.com/aceberg/miniboard/internal/models"
	"github.com/aceberg/miniboard/internal/yaml"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	guiData.Config = AppConfig

	reload := r.URL.Query().Get("reload")
	if reload == "yes" {
		AllLinks = yaml.Read(AppConfig.YamlPath)
	}

	tab := r.URL.Query().Get("tab")
	if tab == "" {
		for key := range AllLinks.Tabs {
			tab = key
			break
		}
	}

	guiData.CurrentTab = AllLinks.Tabs[tab].Name
	guiData.Links.Tabs = AllLinks.Tabs
	guiData.Links.Panels = make(map[string]models.Panel)

	for _, name := range AllLinks.Tabs[tab].Panels {
		guiData.Links.Panels[name] = AllLinks.Panels[name]
	}

	execTemplate(w, "index", guiData)
}
