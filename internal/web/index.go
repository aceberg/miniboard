package web

import (
	// "log"
	"net/http"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	guiData.Config = AppConfig

	newPanels := []models.Panel{}
	onePanel := models.Panel{}

	for _, panel := range AllLinks.Panels {
		onePanel.Name = panel.Name
		for _, host := range panel.Hosts {
			host.State = check.State(host)
			onePanel.Hosts = append(onePanel.Hosts, host)
		}
		newPanels = append(newPanels, onePanel)
		onePanel = models.Panel{}
	}
	AllLinks.Panels = newPanels
	guiData.Links = AllLinks

	execTemplate(w, "index", guiData)
}
