package web

import (
	"embed"

	"github.com/aceberg/miniboard/internal/models"
)

var (
	// AppConfig - config for Web Gui
	AppConfig models.Conf
	// AllLinks - all links
	AllLinks models.Links
	// UptimeMon - monitoring data
	UptimeMon []models.MonData
	// CountRetries - when to send notifications
	CountRetries map[string]int
)

// TemplHTML - html templates
//
//go:embed templates/*
var TemplHTML embed.FS

// TemplPath - path to html templates
const TemplPath = "templates/"

// // TemplPath - path to html templates
// const TemplPath = "../../internal/web/templates/"
