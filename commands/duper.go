package commands

import (
	"fmt"
	xxhash "github.com/cespare/xxhash/v2"
	"github.com/abrar-hnxlabs/go-hnx/commands/core"
	"io/ioutil"
	"regexp"
)

func Duper(srcDir string) {
	fmt.Println("Duper.", srcDir)
	fileList := core.RecursiveListFiles(srcDir)
	total := len(fileList)
	current := 1
	dupeMap := make(map[uint64][]string)
	for _, file := range(fileList) {
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
	
	deleteList := make([]string, 0)
	for _, val := range(dupeMap) {
		if len(val) > 1 {
			// fmt.Printf("Found binary dupes. %v \n", val)
			tempList := make([]string, 0)
			for _, file := range(val) {
				if allowedToDelete(file) {
					tempList = append(tempList, file)
				}
			}
			
			// if temp list is same as the dupe files
			// this means all file can be deled so remove one file and then add to master delete list
			if len(tempList) == len(val) {
				tempList = tempList[1:]
			}

			for _, deletes := range(tempList) {
				deleteList = append(deleteList, deletes)
			}
		}
	}

	for _, f := range(deleteList) {
		fmt.Println(f)
	}
}


func allowedToDelete(filename string) bool {
	pattern := regexp.MustCompile("^photos/[0-9]{4}_[0-9]{2}/")
	return pattern.MatchString(filename)
}