package commands
import (
	"github.com/olebedev/config"
	"path"
	"log"
	"os"
	"path/filepath"
)

func GetConfig() (*config.Config, error) {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	configPath := path.Join(dir, "config.yaml")
	log.Println(configPath)
	return config.ParseYamlFile(configPath)
}
