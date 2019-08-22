package main
import (
	"github.com/olebedev/config"
	"path"
	"log"
)

func GetConfig() (*config.Config, error) {
	configPath := path.Join("/","home","abrar","bin", "config.yaml")
	log.Println(configPath)
	return config.ParseYamlFile(configPath)
}
