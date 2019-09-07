package main

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

func RunDockerInstance(commmand string){
	commmands := strings.Split(commmand, " ")
	log.Printf("%s\n", commmand)
	run(commmands...)
}