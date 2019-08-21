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
  WithOption(cli.NewOption("out", "output file name").WithType(cli.TypeString)).
  WithOption(cli.NewOption("p", "passwordto decrypt file").WithType(cli.TypeString)).
  WithAction(func(args []string, options map[string]string ) int{
    filename := options["in"]
    password := options["p"]
    output := options["out"]
    password32chars := password + strings.Repeat("X",32-len(password))
    Decryptfile(filename, output,password32chars)
    return 0
  })

  docker := cli.NewCommand("docker", "update plex restart docker").
    WithAction(func(args []string, options map[string]string ) int {
      UpdateAndRestartDocker()
      return 0
    })


  app := cli.New("myubuntu tool").
    WithCommand(dns).
    WithCommand(encrypt).
    WithCommand(decrypt).
    WithCommand(docker)

  os.Exit(app.Run(os.Args, os.Stdout))
}