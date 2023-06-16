package web

import (
	"log"
	"net/http"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/conf"
)

// Gui - start web server
func Gui(confPath string) {

	AppConfig = conf.Get(confPath)
	AppConfig.ConfPath = confPath
	AppConfig.Icon = Icon
	log.Println("INFO: starting web gui with config", AppConfig.ConfPath)

	address := AppConfig.Host + ":" + AppConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)                  // index.go
	http.HandleFunc("/config/", configHandler)          // config.go
	http.HandleFunc("/config_save/", saveConfigHandler) // config.go
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}
