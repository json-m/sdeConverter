package main

import (
	"compress/gzip"
	"fmt"
	"os"
	"path/filepath"
)

func writeData(id int, data []byte) error {
	outPath := fmt.Sprintf("output/inventory/%d.bson", id)

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
