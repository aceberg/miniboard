package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/models"
)

// Panel - returns panel formed from Docker API
func Panel(panelName string) models.Panel {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	check.IfError(err)
	defer cli.Close()

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	check.IfError(err)

	panel := models.Panel{}
	panel.Name = panelName
	panel.Hosts = make(map[int]models.Host)

	i := 0
	for _, container := range containers {
		name := container.Names[0]
		name = name[1:]
		ports := container.Ports

		for _, p := range ports {
			IP := p.IP
			port := p.PublicPort
			proto := p.Type

			if IP != "" && proto == "tcp" {
				host := models.Host{}
				host.Name = name
				host.Addr = IP
				host.Port = fmt.Sprintf("%d", port)

				panel.Hosts[i] = host
				i++
			}
		}
	}

	return panel
}
