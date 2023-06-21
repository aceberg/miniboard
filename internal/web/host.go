package web

import (
	// "log"
	"net/http"
	"strconv"

	"github.com/aceberg/miniboard/internal/models"
	"github.com/aceberg/miniboard/internal/yaml"
)

func hostHandler(w http.ResponseWriter, r *http.Request) {

	panel := r.URL.Query().Get("panel")
	del := r.URL.Query().Get("del")
	up := r.URL.Query().Get("up")
	key := r.FormValue("panel")

	if del != "" && panel != "" {
		// log.Println("DEL:", panel, del)

		id, _ := strconv.Atoi(del)
		delete(AllLinks.Panels[panel].Hosts, id)
		assignHostIDs(panel) // assign-IDs.go

	} else if up != "" && panel != "" {
		// log.Println("UP:", panel, up)
		id, _ := strconv.Atoi(up)
		host := AllLinks.Panels[panel].Hosts[id-1]
		AllLinks.Panels[panel].Hosts[id-1] = AllLinks.Panels[panel].Hosts[id]
		AllLinks.Panels[panel].Hosts[id] = host

	} else {
		var id int
		idStr := r.FormValue("id")
		if idStr == "" {
			id = len(AllLinks.Panels[key].Hosts)
		} else {
			id, _ = strconv.Atoi(idStr)
		}

		host := models.Host{}
		host.Name = r.FormValue("name")
		host.Addr = r.FormValue("addr")
		host.Port = r.FormValue("port")
		host.URL = r.FormValue("url")
		host.Icon = r.FormValue("icon")

		AllLinks.Panels[key].Hosts[id] = host
	}

	yaml.Write(AppConfig.YamlPath, AllLinks)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
