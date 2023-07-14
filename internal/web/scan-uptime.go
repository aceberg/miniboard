package web

import (
	// "log"
	"fmt"
	"sync"
	"time"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/models"
	"github.com/aceberg/miniboard/internal/notify"
)

var (
	// RetriesSyncMap - count retries to send notifications
	RetriesSyncMap sync.Map
)

func appendUptimeMon(panelName string, host models.Host, notif bool) {
	mon := models.MonData{}
	mon.Panel = panelName
	mon.Host = host.Name
	mon.Addr = host.Addr
	mon.Port = host.Port
	t := time.Now()
	mon.Date = t.Format("2006-01-02 15:04:05")
	mon.Color = check.Color(t.Format("2006-01-02"))
	mon.State = host.State
	if notif {
		mon.Notify = AllLinks.Uptime.Panels[panelName].Notify
	}
	UptimeMon = append(UptimeMon, mon)
}

func scanUptime(panelName string, host models.Host, oldState bool) {
	var retries int
	var notifEnabled, notif bool

	panel, exists := AllLinks.Uptime.Panels[panelName]

	if AllLinks.Uptime.Enabled && exists {
		if len(panel.Notify) > 0 {
			retriesAny, ok := RetriesSyncMap.LoadOrStore(panelName+host.Name, 0)
			if ok {
				retries = retriesAny.(int)
			}
			notifEnabled = true
		}

		if oldState != host.State {

			if notifEnabled && host.State && (retries > panel.Retries) {
				notify.Notify(panelName, host.Name, "is up", AllLinks.Uptime)
				RetriesSyncMap.Store(panelName+host.Name, 0)
				notif = true
			}
			appendUptimeMon(panelName, host, notif)
		}
		if notifEnabled && !host.State {
			if retries == panel.Retries {
				msg := fmt.Sprintf("is down (retries: %d)", panel.Retries)
				notify.Notify(panelName, host.Name, msg, AllLinks.Uptime)

				appendUptimeMon(panelName, host, true)
			}
			RetriesSyncMap.Store(panelName+host.Name, retries+1)
		}
	}
}
