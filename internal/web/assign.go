package web

import (
	"github.com/aceberg/miniboard/internal/models"
	// "github.com/aceberg/miniboard/internal/yaml"
)

func assignIDs(panel string) {
	var newPanel models.Panel

	newPanel.Name = AllLinks.Panels[panel].Name
	newPanel.Scan = AllLinks.Panels[panel].Scan
	newPanel.Hosts = make(map[int]models.Host)

	i := 0
	for _, host := range AllLinks.Panels[panel].Hosts {
		newPanel.Hosts[i] = host
		i = i + 1
	}
	AllLinks.Panels[panel] = newPanel
}

// func assignAllIDs() {

// }
