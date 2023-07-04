package web

import (
	// "log"
	"strconv"
	"time"

	// "github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/models"
	"github.com/aceberg/miniboard/internal/notify"
)

func scanUptime(panelName string, host models.Host, oldState bool) {
	_, exists := AllLinks.Uptime.Panels[panelName]

	if AllLinks.Uptime.Enabled && exists {
		if oldState != host.State {

			mon := models.MonData{}
			mon.Panel = panelName
			mon.Host = host.Name
			mon.Addr = host.Addr
			mon.Port = host.Port
			mon.Date = time.Now().Format("2006-01-02 15:04:05")
			mon.State = host.State
			UptimeMon = append(UptimeMon, mon)

			if host.State && (CountRetries[panelName+host.Name] > AllLinks.Uptime.Panels[panelName].Retries) {
				notify.Notify(panelName, host.Name, "is up", AllLinks.Uptime)
				CountRetries[panelName+host.Name] = 0
			}
		}
		if !host.State {
			if CountRetries[panelName+host.Name] == AllLinks.Uptime.Panels[panelName].Retries {
				retries := strconv.Itoa(CountRetries[panelName+host.Name])
				notify.Notify(panelName, host.Name, "is down"+" (retries: "+retries+")", AllLinks.Uptime)
			}
			CountRetries[panelName+host.Name] = CountRetries[panelName+host.Name] + 1
		}
	}
}
