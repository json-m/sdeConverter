package main

import (
	"compress/gzip"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func writeData(t Filetype, name string, data []byte) error {
	var outPath string
	switch t {
	// universe stuff
	case Region:
		outPath = fmt.Sprintf("output/universe/%s.bson", name)
	case Constellation:
		outPath = fmt.Sprintf("output/universe/%s.bson", name)
	case SolarSystem:
		outPath = fmt.Sprintf("output/universe/%s.bson", name)
	case Landmarks:
		outPath = fmt.Sprintf("output/universe/%s.bson", name)
	// item stuff
	default:
		return errors.New("unhandled filetype to write")
	}

	// ensure the directory exists
	err := os.MkdirAll(filepath.Dir(outPath), 0755)
	if err != nil {
		return err
	}

	// create file
	file, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// compress the data with gzip
	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	// write the data
	_, err = gzipWriter.Write(data)
	if err != nil {
		return err
	}

	return nil
}
