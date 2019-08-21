package main

import (
	"log"
	"os/exec"
	"fmt"
	"bytes"
)

func down() {
	cmd := exec.Command("/usr/bin/docker-compose","-f","/home/abrar/confs-docker/docker-compose.yml", "down")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", out.String())
}

func updatePlex(){
	cmd := exec.Command("/usr/bin/docker-compose","-f","/home/abrar/confs-docker/docker-compose.yml", "pull", "plex")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", out.String())
}

func up(){
	cmd := exec.Command("/usr/bin/docker-compose","-f","/home/abrar/confs-docker/docker-compose.yml", "up", "-d")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", out.String())
}

func UpdateAndRestartDocker() {
	updatePlex()
	down()
	up()	
}