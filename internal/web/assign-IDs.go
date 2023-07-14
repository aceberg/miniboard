package web

import (
	"sort"

	"github.com/aceberg/miniboard/internal/models"
	"github.com/aceberg/miniboard/internal/yaml"
)

func assignHostIDs(panel string) {
	var newPanel models.Panel
	var keys []int
	for k := range AllLinks.Panels[panel].Hosts {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	newPanel.Name = AllLinks.Panels[panel].Name
	newPanel.Scan = AllLinks.Panels[panel].Scan
	newPanel.Timeout = AllLinks.Panels[panel].Timeout
	newPanel.Hosts = make(map[int]models.Host)

	i := 0
	for _, k := range keys {
		newPanel.Hosts[i] = AllLinks.Panels[panel].Hosts[k]
		i = i + 1
	}
	AllLinks.Panels[panel] = newPanel
}

func assignTabIDs() {
	var keys []int
	for k := range AllLinks.Tabs {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	newTabs := make(map[int]models.Tab)
	i := 0
	for _, k := range keys {
		newTabs[i] = AllLinks.Tabs[k]
		i = i + 1
	}
	AllLinks.Tabs = newTabs
}

func assignPanelIDs(tab int) {
	var newTab models.Tab
	var keys []int
	for k := range AllLinks.Tabs[tab].Panels {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	newTab.Name = AllLinks.Tabs[tab].Name
	newTab.Refresh = AllLinks.Tabs[tab].Refresh
	newTab.Panels = make(map[int]string)
	i := 0
	for _, k := range keys {
		newTab.Panels[i] = AllLinks.Tabs[tab].Panels[k]
		i = i + 1
	}
	AllLinks.Tabs[tab] = newTab
}

func assignAllIDs() {

	assignTabIDs()

	for tabID := range AllLinks.Tabs {
		assignPanelIDs(tabID)
	}

	for panelID := range AllLinks.Panels {
		assignHostIDs(panelID)
	}

	if AllLinks.Panels == nil {
		AllLinks.Panels = make(map[string]models.Panel)
	}

	yaml.Write(AppConfig.YamlPath, AllLinks)
}
