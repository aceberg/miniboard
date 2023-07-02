package notify

import (
	"github.com/containrrr/shoutrrr"
	"log"

	"github.com/aceberg/miniboard/internal/models"
)

func shout(message string, url string) {
	if url != "" {
		err := shoutrrr.Send(url, message)
		if err != nil {
			log.Println("ERROR: Notification failed (shoutrrr):", err)
		}
	}
}

// Notify - send message with shoutrrr
func Notify(panelName, host, state string, uptime models.Uptime) {

	msg := "Host " + panelName + ":" + host + " " + state + "!"

	for _, urlName := range uptime.Panels[panelName].Notify {

		log.Println("INFO:", msg, "Sending notification")
		shout("MiniBoard: "+msg, uptime.Notify[urlName])
	}
}
