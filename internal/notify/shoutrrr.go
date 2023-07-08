package notify

import (
	"github.com/containrrr/shoutrrr"
	"log"

	"github.com/aceberg/miniboard/internal/models"
)

// Notify - send message with shoutrrr
func Notify(panelName, host, state string, uptime models.Uptime) {
	var urls []string

	msg := "Host " + panelName + ":" + host + " " + state + "!"

	for _, urlName := range uptime.Panels[panelName].Notify {
		urls = append(urls, uptime.Notify[urlName])
		log.Println("INFO:", msg, "Sending notification to", urlName)
	}

	sender, err := shoutrrr.CreateSender(urls...)
	sender.Send("MiniBoard: "+msg, nil)
	if err != nil {
		log.Println("ERROR: Notification failed (shoutrrr):", err)
	}
}

// Test - send test notification
func Test(urlName string, uptime models.Uptime) {

	err := shoutrrr.Send(uptime.Notify[urlName], "MiniBoard: test notification")
	log.Println("INFO: Sending test notification to", urlName)
	if err != nil {
		log.Println("ERROR: Notification failed (shoutrrr):", err)
	}
}
