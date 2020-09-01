package core

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

func RecursiveListFiles(root string) []string {
	fileList := make([] string, 0)
	stack := make([] string,0)
	rootStat , err := os.Stat(root)
	if err != nil {
		log.Fatal(err)
		return fileList
	}
	if rootStat.IsDir() {
		stack = append(stack, root)
		for len(stack) > 0 {
			currentLen := len(stack)
			current := stack[currentLen-1]
			stack = stack[:currentLen-1]
			contents, _ := ioutil.ReadDir(current)
			for _, v := range contents {
				oldname := path.Join(current, v.Name())
				if !v.IsDir() {
					fileList = append(fileList, oldname)
				} else {
					stack = append(stack, path.Join(current, v.Name()))
				}
			}
		}
	}
	return fileList
}