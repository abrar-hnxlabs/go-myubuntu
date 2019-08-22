package main
import (
	"github.com/olebedev/config"
	"path"
	"log"
	"os"
)

func GetConfig() (*config.Config, error) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("cannot get current working directory")
		return nil, err
	}
	configPath := path.Join(cwd, "config.yaml")
	return config.ParseYamlFile(configPath)
}
