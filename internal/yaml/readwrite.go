package yaml

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/models"
)

// Read - read .yaml file to struct
func Read(path string) models.Links {

	file, err := os.ReadFile(path)
	check.IfError(err)

	var allLinks models.Links
	err = yaml.Unmarshal(file, &allLinks)
	check.IfError(err)

	return allLinks
}

// Write - write struct to  .yaml file
func Write(path string, allLinks models.Links) {

	yamlData, err := yaml.Marshal(&allLinks)
	check.IfError(err)

	err = os.WriteFile(path, yamlData, 0644)
	check.IfError(err)

	log.Println("INFO: writing new Links file to", path, "\n", string(yamlData))
}
