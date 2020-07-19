package commands

import (
	"log"
	"os/exec"
	"fmt"
	"path"
	"strings"
)

func run(args ...string) {
	dockerFile := path.Join("/","mnt", "external","projects","docker", "docker-compose.yml")
	options := []string{ "-f", dockerFile}
	for _,v := range args {
		options = append(options, v)
	}
	cmd := exec.Command("/usr/bin/docker-compose", options...)
	output, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", output)
}

func UpdateAndRestartDocker() {
	run("pull","plex")
	run("down")
	run("up","-d")
}

func RunDockerInstance(commmand string){
	commmands := strings.Split(commmand, " ")
	log.Printf("%s\n", commmand)
	run(commmands...)
}