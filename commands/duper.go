package commands

import (
	"fmt"
	xxhash "github.com/cespare/xxhash/v2"
	"github.com/abrar-hnxlabs/go-hnx/commands/core"
	"io/ioutil"
)
func Duper(srcDir string) {
	fmt.Println("Duper.", srcDir)
	fileList := core.RecursiveListFiles(srcDir)
	dupeMap := make(map[uint64][]string)
	for _, file := range(fileList) {
		fileBytes, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("error while reading file: %s \n", file)
			continue
		}
		fmt.Printf("Reading file: %s\n", file)
		digest := xxhash.Sum64(fileBytes)
		if _, ok := dupeMap[digest]; ok {
			dupeMap[digest] = append(dupeMap[digest], file)
		} else {
			dupeMap[digest] = []string { file}
		}
	}
	
	for _, val := range(dupeMap) {
		if len(val) > 1 {
			fmt.Printf("Found binary dupes. %v \n", val)
		}
	}
}