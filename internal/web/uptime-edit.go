package web

import (
	"net/http"
	"strconv"

	"github.com/aceberg/miniboard/internal/models"
	"github.com/aceberg/miniboard/internal/yaml"
)

func uptimeEditHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.CurrentTab = "Edit Uptime"
	guiData.Links = AllLinks

	edit := r.FormValue("edit")
	enable := r.FormValue("enable")
	delnotify := r.FormValue("delnotify")
	delpanel := r.FormValue("delpanel")
	nname := r.FormValue("nname")
	link := r.FormValue("link")
	newpanel := r.FormValue("newpanel")

	if enable == "yes" {
		AllLinks.Uptime.Enabled = true
	} else if enable == "no" {
		AllLinks.Uptime.Enabled = false
	}
	if delnotify != "" {
		delete(AllLinks.Uptime.Notify, delnotify)
	}
	if delpanel != "" {
		delete(AllLinks.Uptime.Panels, delpanel)
	}
	if nname != "" && link != "" {
		AllLinks.Uptime.Notify[nname] = link
	}
	if newpanel != "" {
		retries := r.FormValue("retries")
		notify := r.PostForm["notify"]

		var newPan models.MonPanel
		newPan.Retries, _ = strconv.Atoi(retries)
		newPan.Notify = notify
		AllLinks.Uptime.Panels[newpanel] = newPan
	}

	if edit == "yes" {
		yaml.Write(AppConfig.YamlPath, AllLinks)
		http.Redirect(w, r, "/uptime_edit/", 302)
	}

	execTemplate(w, "uptime-edit", guiData)
}
