package web

import (
	"net/http"
	"strconv"

	"github.com/aceberg/miniboard/internal/models"
	"github.com/aceberg/miniboard/internal/notify"
	"github.com/aceberg/miniboard/internal/yaml"
)

func uptimeEditHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.CurrentTab = "Edit Uptime"

	edit := r.FormValue("edit")
	enable := r.FormValue("enable")
	delnotify := r.FormValue("delnotify")
	delpanel := r.FormValue("delpanel")
	nname := r.FormValue("nname")
	link := r.FormValue("link")
	newpanel := r.FormValue("newpanel")
	show := r.FormValue("show")
	notif := r.FormValue("notify")

	if enable == "yes" {
		AllLinks.Uptime.Enabled = true
	}
	if enable == "no" {
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
	if show != "" {
		AllLinks.Uptime.Show, _ = strconv.Atoi(show)
	}
	if notif != "" {
		notify.Test(notif, AllLinks.Uptime)
	}

	guiData.Links = AllLinks
	if edit == "yes" {
		yaml.Write(AppConfig.YamlPath, AllLinks)
		http.Redirect(w, r, "/uptime_edit/", 302)
	}

	execTemplate(w, "uptime-edit", guiData)
}
