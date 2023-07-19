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
	// MuScan - mutex for AllLinks.Panels
	MuScan sync.Mutex
)

func scanPorts(quit chan bool) {
	alreadyScanning := make(map[string]string)

	for {
		select {
		case <-quit:
			return
		default:
			MuScan.Lock()
			panels := AllLinks.Panels
			MuScan.Unlock()
			for name := range panels {
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
			MuScan.Lock()
			panel, exists := AllLinks.Panels[panelName]
			MuScan.Unlock()
			if !exists {
				return
			}

			if panel.Scan {

				hosts := make(map[int]models.Host)

				for key, host := range panel.Hosts {
					oldState := host.State
					host.State = check.State(host)
					hosts[key] = host

					if AllLinks.Uptime.Enabled {
						scanUptime(panelName, host, oldState) // scan-uptime.go
					}
				}
				panel.Hosts = hosts
				MuScan.Lock()
				AllLinks.Panels[panelName] = panel
				MuScan.Unlock()
			}

			timeout, err := strconv.Atoi(panel.Timeout)
			if err != nil || timeout < 1 {
				timeout = 1
			}

			time.Sleep(time.Duration(timeout) * 60 * time.Second)
		}
	}
}
