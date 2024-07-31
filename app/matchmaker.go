package app

import (
	"fmt"
	"regexp"
	"strings"
)

// extractValue traverses the nested map using the provided keys and returns the final value.
// Returns false if any key is missing or the path is not valid.
func extractValue(rb map[string]interface{}, keys []string) (interface{}, bool) {
	currentMap := rb
	for i, key := range keys {
		value, ok := currentMap[key]
		if !ok {
			return nil, false // Key doesn't exist
		}
		if i == len(keys)-1 {
			return value, true // Final value found
		}
		if nestedMap, ok := value.(map[string]interface{}); ok {
			currentMap = nestedMap // Traverse to the nested map
		} else {
			return nil, false // Value is not a map, path is invalid
		}
	}
	return nil, false // This line is technically unreachable
}

// IsMatch checks if the value at the given rule key path matches the rule value.
func IsMatch(rb map[string]interface{}, rule Rule) bool {
	keys := strings.Split(rule.Key, RULE_KEY_SEPARATOR)
	value, ok := extractValue(rb, keys)
	if !ok {
		return false // Key path is invalid
	}
	return fmt.Sprint(value) == rule.Value // Check for value match
}

// IsMatchPattern checks if the value at the given rule key path matches the pattern specified in the rule value.
func IsMatchPattern(rb map[string]interface{}, rule Rule) bool {
	keys := strings.Split(rule.Key, RULE_KEY_SEPARATOR)
	value, ok := extractValue(rb, keys)
	if !ok {
		return false // Key path is invalid
	}
	pattern := regexp.MustCompile(rule.Value)
	return pattern.MatchString(fmt.Sprint(value)) // Check for pattern match
}
