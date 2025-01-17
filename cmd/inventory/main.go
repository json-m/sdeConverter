package main

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"runtime"
	"sdeConverter/pkg/types"
	"sync"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Println("downloading SDE.. might take a sec")
	if dErr := downloadSDE(); dErr != nil {
		log.Fatalln(dErr)
	}
	log.Println("extracting SDE..")
	if eErr := unpackSDE(); eErr != nil {
		log.Fatalln(eErr)
	}
	log.Println("loading fsd/*")

	// read in fsd/types.yaml as a map of item ids
	var fsdData map[int]types.TypeInfo
	f, err := os.Open("fsd/types.yaml")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	decoder := yaml.NewDecoder(reader)
	if err := decoder.Decode(&fsdData); err != nil {
		log.Fatalln(err)
		return
	}
	log.Printf("loaded %d items from fsd/types.yaml\n", len(fsdData))

	// Process items concurrently
	numWorkers := runtime.NumCPU()
	itemChan := make(chan types.I, numWorkers)
	var wg sync.WaitGroup

	// Counter for processed items
	var processedItems int64

	// Start status reporting goroutine
	stopReporting := make(chan struct{})
	go reportStatus(&processedItems, stopReporting)

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(itemChan, &wg, &processedItems)
	}

	// Send items to be processed
	for i, item := range fsdData {
		itemData := types.I{
			Id:    i,
			Group: item.GroupID,
			Name:  item.Name["en"],
		}
		itemChan <- itemData
	}

	close(itemChan)
	wg.Wait()
	close(stopReporting)

	fmt.Print("\n")
}
