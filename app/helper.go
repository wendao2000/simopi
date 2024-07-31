package app

import (
	"encoding/json"
	"fmt"
)

// IsValidJson used to determine whether a string is a valid json format.
func IsValidJson(val interface{}) bool {
	if val == nil {
		return true
	}

	checkJsonBytes := func(data []byte) bool {
		var temp interface{}
		return json.Unmarshal(data, &temp) == nil
	}

	switch v := val.(type) {
	case string:
		return checkJsonBytes([]byte(v))
	case []byte:
		return checkJsonBytes(v)
	default:
		// Marshal the input directly if it's not a string or byte array
		b, err := json.Marshal(val)
		if err != nil {
			return false // Error in marshalling
		}
		// Check if the marshaled JSON starts and ends with curly brackets
		return len(b) > 0 && b[0] == '{' && b[len(b)-1] == '}'
	}
}

// MapToRule takes a map and converts it into a slice of Rule objects.
func MapToRule(data map[string]interface{}) (resp []Rule) {
	// Flatten the nested JSON structure into a slice of Rule objects
	FlattenJSON("", data, &resp)
	return
}

// FlattenJSON recursively flattens a nested JSON structure.
func FlattenJSON(prefix string, v interface{}, result *[]Rule) {
	switch value := v.(type) {
	case map[string]interface{}:
		// If it's a map, iterate through its keys.
		for k, v := range value {
			// Construct the new prefix for nested keys
			newPrefix := k
			if prefix != "" {
				newPrefix = prefix + "." + k
			}
			// Recursively flatten the nested structure
			FlattenJSON(newPrefix, v, result)
		}
	case []interface{}:
		// If it's a slice, iterate through its elements.
		for i, elem := range value {
			// Construct the new prefix for array elements
			newPrefix := fmt.Sprintf("%s[%d]", prefix, i)
			// Recursively flatten the nested structure
			FlattenJSON(newPrefix, elem, result)
		}
	default:
		// Otherwise, it's a primitive value, add it to the result slice.
		*result = append(*result, Rule{
			RuleType: RULE_TYPE_MATCH,
			Key:      prefix,
			Value:    fmt.Sprintf("%v", v),
		})
	}
}
