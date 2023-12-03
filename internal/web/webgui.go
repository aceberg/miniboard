package web

import (
	"log"
	"net/http"

	"github.com/aceberg/miniboard/internal/auth"
	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/conf"
	// "github.com/aceberg/miniboard/internal/db"
	"github.com/aceberg/miniboard/internal/yaml"
)

// Gui - start web server
func Gui(confPath, yamlPath, dbPath, nodePath string) {

	AppConfig, authConf = conf.Get(confPath)
	AppConfig.ConfPath = confPath
	AppConfig.NodePath = nodePath
	AppConfig.YamlPath = yamlPath
	AppConfig.DBPath = dbPath
	AppConfig.Icon = Icon

	log.Println("INFO: starting web gui with config", AppConfig.ConfPath)

	go dbRoutine()     // db-routine.go
	reloadScans()      // webgui.go
	go trimDBRoutine() // db-trim.go

	address := AppConfig.Host + ":" + AppConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)         // index.go
	http.HandleFunc("/login/", loginHandler)   // login.go
	http.HandleFunc("/uptime/", uptimeHandler) // uptime.go

	http.HandleFunc("/auth_conf/", auth.Auth(authConfHandler, &authConf))     // auth.go
	http.HandleFunc("/auth_save/", auth.Auth(saveAuthHandler, &authConf))     // auth.go
	http.HandleFunc("/config/", auth.Auth(configHandler, &authConf))          // config.go
	http.HandleFunc("/config_save/", auth.Auth(saveConfigHandler, &authConf)) // config.go
	http.HandleFunc("/file/", auth.Auth(fileHandler, &authConf))              // file.go
	http.HandleFunc("/host/", auth.Auth(hostHandler, &authConf))              // host.go
	http.HandleFunc("/panels/", auth.Auth(panelsHandler, &authConf))          // panels.go
	http.HandleFunc("/panel_edit/", auth.Auth(panelEditHandler, &authConf))   // panel-edit.go
	http.HandleFunc("/reload/", auth.Auth(reloadHandler, &authConf))          // index.go
	http.HandleFunc("/tabs/", auth.Auth(tabsHandler, &authConf))              // tabs.go
	http.HandleFunc("/tab_edit/", auth.Auth(tabEditHandler, &authConf))       // tab-edit.go
	http.HandleFunc("/uptime_edit/", auth.Auth(uptimeEditHandler, &authConf)) // uptime-edit.go
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
