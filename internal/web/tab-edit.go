package web

import (
	// "log"
	"net/http"
	"strconv"

	"github.com/aceberg/miniboard/internal/models"
	"github.com/aceberg/miniboard/internal/yaml"
)

func tabEditHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.CurrentTab = "Edit Tab"
	guiData.Links = AllLinks

	tabStr := r.FormValue("tab")

	if tabStr != "" {
		tab, _ := strconv.Atoi(tabStr)
		guiData.TabEdit = tab

		action := r.FormValue("action")

		if action != "" {
			if action == "addpan" { // Add panels
				panels := r.PostForm["panels"]

				i := len(AllLinks.Tabs[tab].Panels)
				for _, panelName := range panels {
					AllLinks.Tabs[tab].Panels[i] = panelName
					i = i + 1
				}
				assignPanelIDs(tab)

			} else if action == "deltab" { // Delete Tab

				delete(AllLinks.Tabs, tab)
				assignTabIDs() // assign-IDs.go

				http.Redirect(w, r, "/tabs/", 302)
			} else {
				panStr := r.FormValue("pan")
				pan, _ := strconv.Atoi(panStr)

				if action == "up" { // Move panel up
					panelName := AllLinks.Tabs[tab].Panels[pan]
					AllLinks.Tabs[tab].Panels[pan] = AllLinks.Tabs[tab].Panels[pan-1]
					AllLinks.Tabs[tab].Panels[pan-1] = panelName

				} else if action == "delpan" { // Delete Panel from tab

					delete(AllLinks.Tabs[tab].Panels, pan)
					assignPanelIDs(tab) // assign-IDs.go
				}
			}

			yaml.Write(AppConfig.YamlPath, AllLinks)
		}
	}

	execTemplate(w, "tab-edit", guiData)
}
