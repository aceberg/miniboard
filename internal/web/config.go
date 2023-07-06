package web

import (
	"net/http"

	"github.com/aceberg/miniboard/internal/conf"
	"github.com/aceberg/miniboard/internal/models"
)

func configHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Links = AllLinks
	guiData.CurrentTab = "Config"

	guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "flatly", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "pulse", "quartz", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "yeti", "zephyr"}

	execTemplate(w, "config", guiData)
}

func saveConfigHandler(w http.ResponseWriter, r *http.Request) {

	AppConfig.Host = r.FormValue("host")
	AppConfig.Port = r.FormValue("port")
	AppConfig.Theme = r.FormValue("theme")
	AppConfig.Color = r.FormValue("color")
	AppConfig.ColorOn = r.FormValue("coloron")
	AppConfig.ColorOff = r.FormValue("coloroff")
	AppConfig.BtnWidth = r.FormValue("btnwidth")
	AppConfig.WebRefresh = r.FormValue("refresh")

	if AppConfig.WebRefresh == "0" {
		AppConfig.WebRefresh = "Do not use zero!"
	}

	conf.Write(AppConfig)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
