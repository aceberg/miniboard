package web

import (
	// "log"
	"net/http"

	"github.com/aceberg/miniboard/internal/models"
	// "github.com/aceberg/miniboard/internal/yaml"
)

func panelsHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.CurrentTab = "Edit Panels"
	guiData.Links.Panels = AllLinks.Panels
	guiData.Links.Tabs = AllLinks.Tabs

	execTemplate(w, "panels", guiData)
}
