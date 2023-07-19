package web

import (
	"net/http"
	"os"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/models"
	"github.com/aceberg/miniboard/internal/yaml"
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

		AllLinks = yaml.Read(AppConfig.YamlPath)
		assignAllIDs() // assign-IDs.go

		close(AppConfig.Quit)
		AppConfig.Quit = make(chan bool)

		go scanPorts(AppConfig.Quit) // scan.go

		http.Redirect(w, r, r.Header.Get("Referer"), 302)
	}

	execTemplate(w, "file", guiData)
}
