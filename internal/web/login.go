package web

import (
	"log"
	"net/http"

	"github.com/aceberg/miniboard/internal/auth"
	"github.com/aceberg/miniboard/internal/models"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	guiData.Config = AppConfig

	username := r.FormValue("username")
	password := r.FormValue("password")
	logout := r.FormValue("logout")

	if username == authConf.User && auth.MatchPasswords(authConf.Password, password) {

		log.Println("INFO: user '"+username+"' logged in. Session expire time", authConf.Expire)

		auth.StartSession(w, r)
	}
	if logout == "yes" {

		auth.LogOut(w, r)
	}

	execTemplate(w, "login", guiData)
}
