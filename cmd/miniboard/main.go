package main

import (
	"flag"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/web"
)

const confPath = "/data/miniboard/config.yaml"
const yamlPath = "/data/miniboard/board.yaml"

func main() {
	yamlPtr := flag.String("b", yamlPath, "Path to board yaml file")
	confPtr := flag.String("c", confPath, "Path to config yaml file")
	flag.Parse()

	check.Path(*confPtr)
	check.Path(*yamlPtr)

	web.Gui(*confPtr, *yamlPtr) // webgui.go
}
