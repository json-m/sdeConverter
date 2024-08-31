package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"sync/atomic"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type Filetype int

const (
	Region = iota
	Constellation
	SolarSystem
	Landmarks
)

type FileInfo struct {
	Type Filetype
	Path string
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
	log.Println("converting universe/*")

	fileChan := make(chan FileInfo, 100)
	errChan := make(chan error, 100)

	var wg sync.WaitGroup

	// Add variables for status updates
	var processedCount int64
	var currentFile atomic.Value
	currentFile.Store("")

	// Start status updater
	stopStatusUpdates := make(chan struct{})
	go statusUpdater(&processedCount, &currentFile, stopStatusUpdates)

	// start workers
	numWorkers := runtime.NumCPU()
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(fileChan, errChan, &wg, &processedCount, &currentFile)
	}

	// start error handler
	go func() {
		for err := range errChan {
			log.Println("Error:", err)
		}
	}()

	// walk universe dir
	err := filepath.Walk("universe/", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		switch file := info.Name(); file {
		case "region.yaml":
			fileChan <- FileInfo{Type: Region, Path: path}
		case "constellation.yaml":
			fileChan <- FileInfo{Type: Constellation, Path: path}
		case "solarsystem.yaml":
			fileChan <- FileInfo{Type: SolarSystem, Path: path}
		case "landmarks.yaml":
			//fileChan <- FileInfo{Type: Landmarks, Path: path} // todo: do i care about these for now?
		default:
			log.Println("unknown file type in ", path)
		}

		return nil
	})
	if err != nil {
		log.Println(err)
	}

	// close everything
	close(fileChan)
	wg.Wait()
	close(errChan)
	close(stopStatusUpdates)

	fmt.Printf("\nProcessed %d files\n", atomic.LoadInt64(&processedCount))
}
