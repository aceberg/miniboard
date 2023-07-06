package web

import (
	// "log"
	"strconv"
	"sync"
	"time"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/models"
)

var (
	// Mu - mutex to prevent concurrent map writes
	Mu sync.Mutex
)

func scanPorts() {
	for name := range AllLinks.Panels {

		go scanPanel(name)
	}
}

func scanPanel(panelName string) {
	for {
		if AllLinks.Panels[panelName].Scan {

			onePanel := models.Panel{}
			onePanel.Name = AllLinks.Panels[panelName].Name
			onePanel.Scan = AllLinks.Panels[panelName].Scan
			onePanel.Timeout = AllLinks.Panels[panelName].Timeout

			onePanel.Hosts = make(map[int]models.Host)
			for key, host := range AllLinks.Panels[panelName].Hosts {
				oldState := host.State
				host.State = check.State(host)
				onePanel.Hosts[key] = host

				scanUptime(panelName, host, oldState) // scan-uptime.go
			}
			Mu.Lock()
			AllLinks.Panels[panelName] = onePanel
			Mu.Unlock()
		}

		timeout, err := strconv.Atoi(AllLinks.Panels[panelName].Timeout)
		if err != nil || timeout < 1 {
			timeout = 1
		}

		time.Sleep(time.Duration(timeout) * 60 * time.Second)
	}
}
