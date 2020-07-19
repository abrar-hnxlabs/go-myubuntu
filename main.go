package main
import (
  "github.com/teris-io/cli"
  "os"
  "log"
  "strings"
)

func main() {
  log.Println("Starting App")
  dns := cli.NewCommand("dns", "update google domains").
    WithAction(func(args []string, options map[string]string ) int {
      result, _ := UpdateDns()
      return result
    })

  encrypt := cli.NewCommand("encrypt", "Encrypt a given file").
    WithOption(cli.NewOption("in", "input file name").WithType(cli.TypeString)).
    WithOption(cli.NewOption("p", "password to encrypt file.").WithType(cli.TypeString)).
    WithAction(func(args []string, options map[string]string ) int{
      filename := options["in"]
      password := options["p"]
      password32chars := password + strings.Repeat("X",32-len(password))
      EncryptFile(filename, password32chars)
      return 0
    })

  decrypt := cli.NewCommand("decrypt", "Decrypt a given file").
  WithOption(cli.NewOption("in", "input file name").WithType(cli.TypeString)).
  WithOption(cli.NewOption("p", "passwordto decrypt file").WithType(cli.TypeString)).
  WithAction(func(args []string, options map[string]string ) int{
    filename := options["in"]
    password := options["p"]
    password32chars := password + strings.Repeat("X",32-len(password))
    Decryptfile(filename,password32chars)
    return 0
  })

  docker := cli.NewCommand("docker", "update plex restart docker").
    WithAction(func(args []string, options map[string]string ) int {
      UpdateAndRestartDocker()
      return 0
    })

  dc := cli.NewCommand("dc", "proxy docker-compose commands").
  WithOption(cli.NewOption("c", "docker command to run").WithType(cli.TypeString)).
  WithAction(func(args []string, options map[string]string ) int {
    RunDockerInstance(options["c"])
    return 0
  })


  slugs := cli.NewCommand("sluggify", "sluggify the filenames.").
  WithOption(cli.NewOption("r", "root folder to start scan from.").WithType(cli.TypeString)).
  WithAction(func(args []string, options map[string]string ) int {
    SlugifyFiles(options["r"])
    return 0
  })


  app := cli.New("Version: 1.1.0").
    WithCommand(dns).
    WithCommand(encrypt).
    WithCommand(decrypt).
    WithCommand(docker).
    WithCommand(dc).
    WithCommand(slugs)

  os.Exit(app.Run(os.Args, os.Stdout))
}