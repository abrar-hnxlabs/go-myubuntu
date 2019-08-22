package main
import (
	"github.com/olebedev/config"
	"path"
)

func GetConfig() (*config.Config, error) {
	configPath := path.Join("~","bin", "config.yaml")
	return config.ParseYamlFile(configPath)
}
