package web

import (
	// "log"
	"time"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/models"
)

func scanPorts(timeout int) {
	for {
		newPanels := []models.Panel{}

		for _, panel := range AllLinks.Panels {

			newPanels = append(newPanels, scanPanel(panel))
		}
		AllLinks.Panels = newPanels

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

	for _, host := range panel.Hosts {
		host.State = check.State(host)
		onePanel.Hosts = append(onePanel.Hosts, host)
	}

	return onePanel
}
