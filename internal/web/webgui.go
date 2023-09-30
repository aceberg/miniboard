package web

import (
	"log"
	"net/http"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/conf"
	"github.com/aceberg/miniboard/internal/yaml"
)

// Gui - start web server
func Gui(confPath, yamlPath, dbPath, nodePath string) {

	AppConfig = conf.Get(confPath)
	AppConfig.ConfPath = confPath
	AppConfig.NodePath = nodePath
	AppConfig.YamlPath = yamlPath
	AppConfig.DBPath = dbPath
	AppConfig.Icon = Icon

	log.Println("INFO: starting web gui with config", AppConfig.ConfPath)

	go dbRoutine() // db-routine.go
	reloadScans()  // webgui.go

	address := AppConfig.Host + ":" + AppConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)                  // index.go
	http.HandleFunc("/config/", configHandler)          // config.go
	http.HandleFunc("/config_save/", saveConfigHandler) // config.go
	http.HandleFunc("/file/", fileHandler)              // file.go
	http.HandleFunc("/host/", hostHandler)              // host.go
	http.HandleFunc("/panels/", panelsHandler)          // panels.go
	http.HandleFunc("/panel_edit/", panelEditHandler)   // panel-edit.go
	http.HandleFunc("/tabs/", tabsHandler)              // tabs.go
	http.HandleFunc("/tab_edit/", tabEditHandler)       // tab-edit.go
	http.HandleFunc("/uptime/", uptimeHandler)          // uptime.go
	http.HandleFunc("/uptime_edit/", uptimeEditHandler) // uptime-edit.go
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}

func reloadScans() {
	AllLinks = yaml.Read(AppConfig.YamlPath)
	assignAllIDs() // assign-IDs.go

	if AppConfig.Quit != nil {
		close(AppConfig.Quit)
	}
	AppConfig.Quit = make(chan bool)

	go scanPorts(AppConfig.Quit) // scan.go
}
