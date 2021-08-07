package commands

import (
	"fmt"
	xxhash "github.com/cespare/xxhash/v2"
	"github.com/abrar-hnxlabs/go-hnx/commands/core"
	"io/ioutil"
	"strings"
	"sync"
	"sync/atomic"
)

var hashMaps chan map[uint64][]string
var counter uint32 = 0
var mutex = &sync.Mutex{}
var total uint32
var pageSize uint32 = 50

func Duper(srcDir string, extension string) {
	fmt.Println("Duper.", srcDir)
	fileList := core.RecursiveListFiles(srcDir)
	total = uint32(len(fileList))
	
	workers := total / pageSize + 1;
	hashMaps = make(chan map[uint64][]string, workers)
	fmt.Printf("Spawing, %d go routines \n", workers)
	var wg sync.WaitGroup
	startIdx := 0
	var i uint32 = 0
	for ; i< workers; i++ {
		wg.Add(1)
		endIdx := startIdx+ int(pageSize)
		if endIdx > len(fileList) {
			endIdx = len(fileList) - int(i * pageSize)
			endIdx = startIdx + endIdx
		}
		go createHashMap(fileList[startIdx:endIdx], extension, &wg)
		startIdx += int(pageSize)
	}
	wg.Wait()
	close(hashMaps)
	masterDupMap := make(map[uint64][]string)
	fmt.Println("Printing Dupes...")

	for singlemap := range hashMaps {
		for key, entry := range singlemap {
			for _, file := range entry {
				if _, ok := masterDupMap[key]; ok {
					masterDupMap[key] = append(masterDupMap[key], file)
				} else {
					masterDupMap[key] = []string {file}
				}
			}
		}
	}
	//fmt.Println(masterDupMap)
	for _, val := range(masterDupMap) {
		if len(val) > 1 {
			for _, f := range (val) {
				fmt.Printf("%s,", f)
			}
			fmt.Printf("\n")
		}
	}
}

func createHashMap(fileList []string, extension string, wg *sync.WaitGroup) {
	dupeMap := make(map[uint64][]string)
	for _, file := range(fileList) {
		fileLower := strings.ToLower(file) 
		// skip file if extension filter present , and does not match the file extenstion
		if len(extension) > 0 && !strings.HasSuffix(fileLower, extension) {
			atomic.AddUint32(&counter, 1)
			continue
		}
		fileBytes, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("error while reading file: %s \n", file)
			continue
		}
		atomic.AddUint32(&counter, 1)
		if atomic.LoadUint32(&counter) % pageSize == 0 {
			fmt.Printf("(%d/%d) file: %s\n", atomic.LoadUint32(&counter), total, file)
		}
		digest := xxhash.Sum64(fileBytes)
		if _, ok := dupeMap[digest]; ok {
			dupeMap[digest] = append(dupeMap[digest], file)
		} else {
			dupeMap[digest] = []string { file}
		}
		// current += 1
	}

	hashMaps <- dupeMap
	wg.Done()
}
