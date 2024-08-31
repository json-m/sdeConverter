package main

import (
	"context"
	"fmt"
	"github.com/saracen/fastzip"
	"io"
	"net/http"
	"os"
)

func downloadSDE() error {
	url := "https://eve-static-data-export.s3-eu-west-1.amazonaws.com/tranquility/universe.zip"

	// download zip file from url to current directory
	out, err := os.Create("universe.zip")
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status not OK: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// unpackSDE unpacks a zip to the current directory
func unpackSDE() error {
	e, err := fastzip.NewExtractor("universe.zip", "./universe")
	if err != nil {
		return err
	}
	defer e.Close()

	// Extract archive files
	if err = e.Extract(context.TODO()); err != nil {
		return err
	}
	return nil
}