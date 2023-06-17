package web

import (
	"net/http"

	"github.com/aceberg/miniboard/internal/models"
)

func editHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig

	panel := r.URL.Query().Get("panel")
	guiData.Themes = append(guiData.Themes, panel)

	execTemplate(w, "edit", guiData)
}
