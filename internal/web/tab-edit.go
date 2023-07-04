package web

import (
	// "log"
	"net/http"
	"strconv"

	"github.com/aceberg/miniboard/internal/models"
)

func tabEditHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.CurrentTab = "Edit Tab"

	tabStr := r.FormValue("tab")
	action := r.FormValue("action")

	tab, _ := strconv.Atoi(tabStr)
	guiData.TabEdit = tab

	if action == "addpan" { // Add panels
		panels := r.PostForm["panels"]

		i := len(AllLinks.Tabs[tab].Panels)
		for _, panelName := range panels {
			AllLinks.Tabs[tab].Panels[i] = panelName
			i = i + 1
		}
	}

	if action == "deltab" { // Delete Tab
		delete(AllLinks.Tabs, tab)
		assignAllIDs() // assign-IDs.go

		http.Redirect(w, r, "/tabs/", 302)
	}

	panStr := r.FormValue("pan")
	pan, _ := strconv.Atoi(panStr)

	if action == "up" { // Move panel up
		panelName := AllLinks.Tabs[tab].Panels[pan]
		AllLinks.Tabs[tab].Panels[pan] = AllLinks.Tabs[tab].Panels[pan-1]
		AllLinks.Tabs[tab].Panels[pan-1] = panelName
	}

	if action == "delpan" { // Delete Panel from tab
		delete(AllLinks.Tabs[tab].Panels, pan)
	}

	if action != "" { // If changes assighn IDs and write to yaml
		assignAllIDs() // assign-IDs.go
		http.Redirect(w, r, "/tab_edit/?tab="+tabStr, 302)
	}

	guiData.Links = AllLinks

	execTemplate(w, "tab-edit", guiData)
}
