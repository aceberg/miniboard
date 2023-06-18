package web

import (
	"log"
	"net/http"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/conf"
	"github.com/aceberg/miniboard/internal/yaml"
)

// Gui - start web server
func Gui(confPath, yamlPath string) {

	AppConfig = conf.Get(confPath)
	AppConfig.ConfPath = confPath
	AppConfig.YamlPath = yamlPath
	AppConfig.Icon = Icon
	log.Println("INFO: starting web gui with config", AppConfig.ConfPath)

	AllLinks = yaml.Read(AppConfig.YamlPath)
	// log.Println("ALL:", AllLinks)
	go scanPorts(60) // scan.go

	address := AppConfig.Host + ":" + AppConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)                  // index.go
	http.HandleFunc("/config/", configHandler)          // config.go
	http.HandleFunc("/config_save/", saveConfigHandler) // config.go
	http.HandleFunc("/tabs/", tabHandler)               // tabs.go
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}
