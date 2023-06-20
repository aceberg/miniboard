package web

import (
	"log"
	"net/http"
	"strconv"

	"github.com/aceberg/miniboard/internal/models"
	// "github.com/aceberg/miniboard/internal/yaml"
)

func tabEditHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.CurrentTab = "Edit Tab"
	guiData.Links = AllLinks

	idStr := r.FormValue("id")

	if idStr != "" {
		id, _ := strconv.Atoi(idStr)
		guiData.TabEdit = id

		log.Println("ID:", id)
	}

	execTemplate(w, "tab-edit", guiData)
}
