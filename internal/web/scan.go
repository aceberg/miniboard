package web

import (
	"log"
	"time"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/models"
)

func scanPorts() {
	for name := range AllLinks.Panels {

		go scanPanel(name)
	}
}

func scanPanel(panelName string) {
	for {
		if AllLinks.Panels[panelName].Scan {

			_, exists := AllLinks.Uptime.Panels[panelName]

			onePanel := models.Panel{}
			onePanel.Name = AllLinks.Panels[panelName].Name
			onePanel.Scan = AllLinks.Panels[panelName].Scan

			onePanel.Hosts = make(map[int]models.Host)
			for key, host := range AllLinks.Panels[panelName].Hosts {
				oldState := host.State
				host.State = check.State(host)
				onePanel.Hosts[key] = host

				if exists && (oldState != host.State) {
					log.Println("INFO: online status changed, host =", host.Name, "online =", host.State)

					mon := models.MonData{}
					mon.Panel = panelName
					mon.Host = host.Name
					mon.Addr = host.Addr
					mon.Port = host.Port
					mon.Date = time.Now().Format("2006-01-02 15:04:05")
					mon.State = host.State
					UptimeMon = append(UptimeMon, mon)

					if host.State {
						log.Println("INFO: host", panelName, ":", host.Name, "is up. Sending notification")
					}
				}
				if exists && !host.State {
					if CountRetries[panelName+host.Name] == AllLinks.Uptime.Panels[panelName].Retries {
						log.Println("INFO: host", panelName, ":", host.Name, "is down for", CountRetries[panelName+host.Name], "retries. Sending notification")
					}
					CountRetries[panelName+host.Name] = CountRetries[panelName+host.Name] + 1
				}

			}
			AllLinks.Panels[panelName] = onePanel
		}
		// if AllLinks.Panels[panelName].Timeout == "" {
		// 	timeout := 60
		// } else {
		// 	// get int timeout
		// }
		timeout := 60
		time.Sleep(time.Duration(timeout) * time.Second)
	}
}
