package main

import (
	"fmt"
	"strconv"
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
