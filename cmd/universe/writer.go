package main

import (
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

	return os.WriteFile(outPath, data, 0644)
}
