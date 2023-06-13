package main

import (
	"flag"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/web"
)

const confPath = "/data/miniboard/config.yaml"

func main() {
	confPtr := flag.String("c", confPath, "Path to config yaml file")
	flag.Parse()

	check.Path(*confPtr)

	web.Gui(*confPtr) // webgui.go
}
