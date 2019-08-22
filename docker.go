package main

import (
	"log"
	"os/exec"
	"fmt"
	"path"
)

func run(args ...string) {
	dockerFile := path.Join("/","home", "abrar","confs-docker", "docker-compose.yml")
	options := []string{ "-f", dockerFile}
	for i :=0 ; i <len(args); i++ {
		options = append(options, args[i])
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