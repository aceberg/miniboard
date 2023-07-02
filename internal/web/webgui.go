package web

import (
	"log"
	"net/http"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/conf"
	"github.com/aceberg/miniboard/internal/yaml"
)

// Gui - start web server
func Gui(confPath, yamlPath, nodePath string) {

	AppConfig = conf.Get(confPath)
	AppConfig.ConfPath = confPath
	AppConfig.NodePath = nodePath
	AppConfig.YamlPath = yamlPath
	AppConfig.Icon = Icon
	log.Println("INFO: starting web gui with config", AppConfig.ConfPath)

	AllLinks = yaml.Read(AppConfig.YamlPath)
	// log.Println("ALL:", AllLinks)
	CountRetries = make(map[string]int)
	assignAllIDs() // assign-IDs.go
	go scanPorts() // scan.go

	address := AppConfig.Host + ":" + AppConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)                  // index.go
	http.HandleFunc("/config/", configHandler)          // config.go
	http.HandleFunc("/config_save/", saveConfigHandler) // config.go
	http.HandleFunc("/host/", hostHandler)              // host.go
	http.HandleFunc("/panels/", panelsHandler)          // panels.go
	http.HandleFunc("/panel_edit/", panelEditHandler)   // panel-edit.go
	http.HandleFunc("/tabs/", tabsHandler)              // tabs.go
	http.HandleFunc("/tab_edit/", tabEditHandler)       // tab-edit.go
	http.HandleFunc("/uptime/", uptimeHandler)          // uptime.go
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}
