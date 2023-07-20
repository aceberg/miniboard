package web

import (
	"net/http"
	"os"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/models"
)

func fileHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	guiData.Config = AppConfig
	guiData.CurrentTab = "Edit board file"
	guiData.Links = AllLinks

	text := r.FormValue("text")

	file, err := os.ReadFile(AppConfig.YamlPath)
	check.IfError(err)
	guiData.Version = string(file)

	if text != "" {
		err := os.WriteFile(AppConfig.YamlPath, []byte(text), 0644)
		check.IfError(err)

		reloadScans() // webgui.go

		http.Redirect(w, r, r.Header.Get("Referer"), 302)
	}

	execTemplate(w, "file", guiData)
}
