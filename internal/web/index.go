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

	guiData.Links = AllLinks

	execTemplate(w, "index", guiData)
}
