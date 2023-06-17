package check

import (
	// "log"
	"net"
	"time"

	"github.com/aceberg/miniboard/internal/models"
)

// State - returns state of a service
func State(host models.Host) bool {

	if host.Port == "" {
		host.Port = "80"
	}

	timeout := 3 * time.Second
	target := host.Addr + ":" + host.Port

	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		// log.Println("ERROR:", err)
		return false
	}
	conn.Close()
	return true
}
