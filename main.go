package main

import (
	"os"
	"strings"
	"strconv"
	"github.com/abrar-hnxlabs/go-hnx/commands"
	"github.com/teris-io/cli"
)

func main() {
	dns := cli.NewCommand("dns", "update google domains").
		WithAction(func(args []string, options map[string]string) int {
			result, _ := commands.UpdateDns()
			return result
		})

	encrypt := cli.NewCommand("encrypt", "Encrypt a given file").
		WithOption(cli.NewOption("in", "input file name").WithType(cli.TypeString)).
		WithOption(cli.NewOption("p", "password to encrypt file.").WithType(cli.TypeString)).
		WithAction(func(args []string, options map[string]string) int {
			filename := options["in"]
			password := options["p"]
			password32chars := password + strings.Repeat("X", 32-len(password))
			commands.EncryptFile(filename, password32chars)
			return 0
		})

	decrypt := cli.NewCommand("decrypt", "Decrypt a given file").
		WithOption(cli.NewOption("in", "input file name").WithType(cli.TypeString)).
		WithOption(cli.NewOption("p", "passwordto decrypt file").WithType(cli.TypeString)).
		WithAction(func(args []string, options map[string]string) int {
			filename := options["in"]
			password := options["p"]
			password32chars := password + strings.Repeat("X", 32-len(password))
			commands.Decryptfile(filename, password32chars)
			return 0
		})

	docker := cli.NewCommand("docker", "update plex restart docker").
		WithAction(func(args []string, options map[string]string) int {
			commands.UpdateAndRestartDocker()
			return 0
		})

	dc := cli.NewCommand("dc", "proxy docker-compose commands").
		WithOption(cli.NewOption("c", "docker command to run").WithType(cli.TypeString)).
		WithAction(func(args []string, options map[string]string) int {
			commands.RunDockerInstance(options["c"])
			return 0
		})

	slugs := cli.NewCommand("sluggify", "sluggify the filenames.").
		WithOption(cli.NewOption("r", "root folder to start scan from.").WithType(cli.TypeString)).
		WithAction(func(args []string, options map[string]string) int {
			commands.SlugifyFiles(options["r"])
			return 0
		})

	duper := cli.NewCommand("duper", "find files for duper.").
		WithOption(cli.NewOption("r", "root folder to scan dupes").WithType(cli.TypeString)).
		WithOption(cli.NewOption("e", "file extension to filter").WithType(cli.TypeString)).
		WithAction(func(args []string, options map[string]string) int {
			commands.Duper(options["r"], options["e"])
			return 0
		})

	getip := cli.NewCommand("ip", "Get outside world ip.").
		WithAction(func(args []string, options map[string]string) int {
			commands.GetIp()
			return 0
		})

	getplexurl := cli.NewCommand("plex", "Get plex download url.").
		WithAction(func(args []string, options map[string]string) int {
			commands.Plexurl()
			return 0
		})
	
	renderCanvas := cli.NewCommand("passport", "Render passport size photos.").
		WithOption(cli.NewOption("ppi", "pixel per inch for image, Default: 150").WithType(cli.TypeInt)).
		WithOption(cli.NewOption("i", "source image, has to be 1x1 aspect ratio").WithType(cli.TypeString)).
		WithAction(func(args []string, options map[string]string) int {
		ppi := 150
		if len(options["ppi"]) > 0 {
			ppi, _ = strconv.Atoi(options["ppi"])
		}
		inputfile := options["i"]
		commands.RenderCanvas(inputfile, float64(ppi))
		return 0
	})

	app := cli.New("Version: 1.1.2").
		WithCommand(dns).
		WithCommand(encrypt).
		WithCommand(decrypt).
		WithCommand(docker).
		WithCommand(dc).
		WithCommand(slugs).
		WithCommand(duper).
		WithCommand(getip).
		WithCommand(getplexurl).
		WithCommand(renderCanvas)

	os.Exit(app.Run(os.Args, os.Stdout))
}
