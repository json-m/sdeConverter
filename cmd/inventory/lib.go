package main

import (
	"fmt"
	"log"
	"sdeConverter/pkg/types"
	"sync"
	"sync/atomic"
	"time"
)

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

func worker(itemChan <-chan types.I, wg *sync.WaitGroup, processedItems *int64) {
	defer wg.Done()
	for item := range itemChan {
		if bErr := convertToBson(item); bErr != nil {
			log.Printf("Error converting item %d to BSON: %v\n", item.Id, bErr)
		}
		atomic.AddInt64(processedItems, 1)
	}
}

func reportStatus(processedItems *int64, stop <-chan struct{}) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	var lastProcessed int64
	startTime := time.Now()
	_ = lastProcessed

	for {
		select {
		case <-ticker.C:
			currentProcessed := atomic.LoadInt64(processedItems)
			elapsedTime := time.Since(startTime).Seconds()
			averagePerSecond := float64(currentProcessed) / elapsedTime

			fmt.Printf("\rProcessed: %d, Avg: %.2f items/sec",
				currentProcessed, averagePerSecond)

			lastProcessed = currentProcessed
		case <-stop:
			return
		}
	}
}
