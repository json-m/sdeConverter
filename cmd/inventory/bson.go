package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"sdeConverter/pkg/types"
)

func convertToBson(data types.I) error {

	// convert all keys to strings recursively
	convertedData := convertKeysToStrings(data)

	//// assert the type of convertedData
	//convertedMap, ok := convertedData.(map[string]interface{})
	//if !ok {
	//	return fmt.Errorf("failed to convert data to map[string]interface{}")
	//}

	bsonData, err := bson.Marshal(convertedData)
	if err != nil {
		return fmt.Errorf("error marshaling to BSON: %v", err)
	}

	if data.Id == 0 {
		return fmt.Errorf("could not determine ID for item: %d %s", data.Id, data.Name)
	}

	return writeData(data.Id, bsonData)
}