package main
import "github.com/olebedev/config"

func GetConfig() (*config.Config, error) {
	return config.ParseYamlFile("config.yaml")
}
