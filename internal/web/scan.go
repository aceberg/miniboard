package web

import (
	"log"
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

func scanPorts(quit chan bool) {
	alreadyScanning := make(map[string]string)

	for {
		select {
		case <-quit:
			return
		default:
			for name := range AllLinks.Panels {
				_, exists := alreadyScanning[name]
				if !exists {
					go scanPanel(name, quit)
					alreadyScanning[name] = name
				}
			}

			time.Sleep(60 * time.Second)
		}
	}
}

func scanPanel(panelName string, quit chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			panel, exists := AllLinks.Panels[panelName]
			if !exists {
				return
			}

			if panel.Scan {

				hosts := make(map[int]models.Host)

				for key, host := range panel.Hosts {
					oldState := host.State
					host.State = check.State(host)
					hosts[key] = host

					scanUptime(panelName, host, oldState) // scan-uptime.go
				}
				panel.Hosts = hosts
				Mu.Lock()
				AllLinks.Panels[panelName] = panel
				Mu.Unlock()
			}

			timeout, err := strconv.Atoi(panel.Timeout)
			if err != nil || timeout < 1 {
				timeout = 1
			}

			log.Println("Scaned panel", panelName)

			time.Sleep(time.Duration(timeout) * 60 * time.Second)
		}
	}
}
