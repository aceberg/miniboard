package web

import (
	// "log"
	"net/http"

	"github.com/aceberg/miniboard/internal/models"
	"github.com/aceberg/miniboard/internal/yaml"
)

func panelsHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	oldkey := r.FormValue("oldkey")
	key := r.FormValue("key")
	name := r.FormValue("name")
	scan := r.FormValue("scan")

	if key != "" && name != "" {
		// log.Println("PANEL:", oldkey, key, name, scan)

		_, exists := AllLinks.Panels[key]

		panel := models.Panel{}
		panel.Name = name
		if scan == "on" {
			panel.Scan = true
		}
		panel.Hosts = AllLinks.Panels[oldkey].Hosts
		if panel.Hosts == nil {
			panel.Hosts = make(map[int]models.Host)
		}

		if !exists && key != oldkey {
			AllLinks.Panels[key] = panel
			delete(AllLinks.Panels, oldkey)

			yaml.Write(AppConfig.YamlPath, AllLinks)
		} else if exists && key == oldkey {
			AllLinks.Panels[key] = panel

			yaml.Write(AppConfig.YamlPath, AllLinks)
		}
	}

	guiData.Config = AppConfig
	guiData.CurrentTab = "Edit Panels"
	guiData.Links.Panels = AllLinks.Panels
	guiData.Links.Tabs = AllLinks.Tabs

	execTemplate(w, "panels", guiData)
}
