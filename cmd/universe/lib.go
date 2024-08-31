package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func worker(fileChan <-chan FileInfo, errChan chan<- error, wg *sync.WaitGroup, processedCount *int64, currentFile *atomic.Value) {
	defer wg.Done()
	for file := range fileChan {
		currentFile.Store(file.Path)
		if err := convertToBson(file.Type, file.Path); err != nil {
			errChan <- fmt.Errorf("error converting %s: %v", file.Path, err)
		}
		atomic.AddInt64(processedCount, 1)
	}
}

func statusUpdater(processedCount *int64, currentFile *atomic.Value, stop <-chan struct{}) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			count := atomic.LoadInt64(processedCount)
			file := currentFile.Load().(string)
			fmt.Printf("\rProcessed %d files. Current file: %s", count, file)
		case <-stop:
			return
		}
	}
}

func convertKeysToStrings(v interface{}) interface{} {
	switch x := v.(type) {
	case map[interface{}]interface{}:
		m := make(map[string]interface{}, len(x))
		for k, v := range x {
			m[fmt.Sprint(k)] = convertKeysToStrings(v)
		}
		return m
	case map[string]interface{}:
		m := make(map[string]interface{}, len(x))
		for k, v := range x {
			m[k] = convertKeysToStrings(v)
		}
		return m
	case []interface{}:
		for i, e := range x {
			x[i] = convertKeysToStrings(e)
		}
	}
	return v
}

func getStringValue(m map[string]interface{}, key string) string {
	if val, ok := m[key]; ok {
		switch v := val.(type) {
		case int:
			return strconv.Itoa(v)
		case int64:
			return strconv.FormatInt(v, 10)
		case float64:
			return strconv.FormatFloat(v, 'f', -1, 64)
		case string:
			return v
		default:
			return fmt.Sprintf("%v", v)
		}
	}
	return ""
}
