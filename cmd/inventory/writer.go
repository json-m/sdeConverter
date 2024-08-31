package main

import (
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

	return os.WriteFile(outPath, data, 0644)
}
