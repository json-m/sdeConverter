package main

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/yaml.v3"
	"os"
)

func convertToBson(t Filetype, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var data map[string]interface{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	// convert all keys to strings recursively
	convertedData := convertKeysToStrings(data)

	// assert the type of convertedData
	convertedMap, ok := convertedData.(map[string]interface{})
	if !ok {
		return fmt.Errorf("failed to convert data to map[string]interface{}")
	}

	bsonData, err := bson.Marshal(convertedMap)
	if err != nil {
		return fmt.Errorf("error marshaling to BSON: %v", err)
	}

	// determine the appropriate ID field based on the file type
	var id string
	switch t {
	case Region:
		id = getStringValue(convertedMap, "regionID")
	case Constellation:
		id = getStringValue(convertedMap, "constellationID")
	case SolarSystem:
		id = getStringValue(convertedMap, "solarSystemID")
	case Landmarks:
		id = getStringValue(convertedMap, "landmarkNameID") //todo: unused
	default:
		return errors.New("unhandled filetype for bson conversion")
	}

	if id == "" {
		return fmt.Errorf("could not determine ID for file: %s", path)
	}

	return writeData(t, id, bsonData)
}
