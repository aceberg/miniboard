package web

import (
	// "log"
	"strconv"
	"time"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/models"
	"github.com/aceberg/miniboard/internal/notify"
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
			onePanel.Timeout = AllLinks.Panels[panelName].Timeout

			onePanel.Hosts = make(map[int]models.Host)
			for key, host := range AllLinks.Panels[panelName].Hosts {
				oldState := host.State
				host.State = check.State(host)
				onePanel.Hosts[key] = host

				if AllLinks.Uptime.Enabled && exists && (oldState != host.State) {

					mon := models.MonData{}
					mon.Panel = panelName
					mon.Host = host.Name
					mon.Addr = host.Addr
					mon.Port = host.Port
					mon.Date = time.Now().Format("2006-01-02 15:04:05")
					mon.State = host.State
					UptimeMon = append(UptimeMon, mon)

					if host.State && (CountRetries[panelName+host.Name] >= AllLinks.Uptime.Panels[panelName].Retries) {
						notify.Notify(panelName, host.Name, "is up", AllLinks.Uptime)
						CountRetries[panelName+host.Name] = 0
					}
				}
				if AllLinks.Uptime.Enabled && exists && !host.State {
					if CountRetries[panelName+host.Name] == AllLinks.Uptime.Panels[panelName].Retries {
						retries := strconv.Itoa(CountRetries[panelName+host.Name])
						notify.Notify(panelName, host.Name, "is down"+" (retries: "+retries+")", AllLinks.Uptime)
					}
					CountRetries[panelName+host.Name] = CountRetries[panelName+host.Name] + 1
				}

			}
			AllLinks.Panels[panelName] = onePanel
		}

		timeout, err := strconv.Atoi(AllLinks.Panels[panelName].Timeout)
		if err != nil || timeout < 1 {
			timeout = 1
		}

		time.Sleep(time.Duration(timeout) * 60 * time.Second)
	}
}
