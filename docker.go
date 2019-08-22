package main

import (
	"log"
	"os/exec"
	"fmt"
	"path"
)

func run(args string) {
	dockerFile := path.Join("~", "confs-docker", "docker-compose.yml")
	cmd := exec.Command("/usr/bin/docker-compose","-f",dockerFile, args)
	output, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", output)
}

func UpdateAndRestartDocker() {
	run("pull plex")
	run("down")
	run("up -d")
}