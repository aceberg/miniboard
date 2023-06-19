package web

import (
	// "log"
	"time"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/models"
)

func scanPorts(timeout int) {
	for {
		for name, panel := range AllLinks.Panels {

			AllLinks.Panels[name] = scanPanel(panel)
		}

		time.Sleep(time.Duration(timeout) * time.Second)
	}
}

func scanPanel(panel models.Panel) models.Panel {

	if !panel.Scan {
		return panel
	}
	onePanel := models.Panel{}
	onePanel.Name = panel.Name
	onePanel.Scan = panel.Scan

	onePanel.Hosts = make(map[int]models.Host)
	for key, host := range panel.Hosts {
		host.State = check.State(host)
		onePanel.Hosts[key] = host
	}

	return onePanel
}
