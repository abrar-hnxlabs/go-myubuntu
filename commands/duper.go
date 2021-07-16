package commands

import (
	"fmt"
	xxhash "github.com/cespare/xxhash/v2"
	"github.com/abrar-hnxlabs/go-hnx/commands/core"
	"io/ioutil"
	// "regexp"
	"strings"
)

func Duper(srcDir string, extension string) {
	fmt.Println("Duper.", srcDir)
	fileList := core.RecursiveListFiles(srcDir)
	total := len(fileList)
	current := 1
	dupeMap := make(map[uint64][]string)
	for _, file := range(fileList) {
		fileLower := strings.ToLower(file) 
		// skip file if extension filter present , and does not match the file extenstion
		if len(extension) > 0 && !strings.HasSuffix(fileLower, extension) { 
			current +=1
			continue
		}
		fileBytes, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("error while reading file: %s \n", file)
			continue
		}
		fmt.Printf("(%d/%d) Reading file: %s\n", current, total, file)
		digest := xxhash.Sum64(fileBytes)
		if _, ok := dupeMap[digest]; ok {
			dupeMap[digest] = append(dupeMap[digest], file)
		} else {
			dupeMap[digest] = []string { file}
		}
		current += 1
	}
	
	for _, val := range(dupeMap) {
		if len(val) > 1 {
			fmt.Printf("Found binary dupes. %v \n", val)
		}
	}
}

/*
func allowedToDelete(filename string) bool {
	pattern := regexp.MustCompile("^photos/[0-9]{4}_[0-9]{2}/")
	return pattern.MatchString(filename)
}
*/