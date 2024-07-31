package app

import (
	"net/http"
	"regexp"
)

// CheckHeader evaluates a list of rules against a provided HTTP header.
// Returns false if any rule is not satisfied according to its type.
func CheckHeader(scnHeader []Rule, rHeader http.Header) bool {
	for _, rule := range scnHeader {
		headerValue := rHeader.Get(rule.Key)
		switch rule.RuleType {
		case RULE_TYPE_MATCH:
			if rule.Value != headerValue {
				return false // Header value does not match the rule value
			}
		case RULE_TYPE_NOT_MATCH:
			if rule.Value == headerValue {
				return false // Header value matches when it should not
			}
		case RULE_TYPE_PATTERN:
			if !regexp.MustCompile(rule.Value).MatchString(headerValue) {
				return false // Header value does not match the pattern
			}
		case RULE_TYPE_NOT_PATTERN:
			if regexp.MustCompile(rule.Value).MatchString(headerValue) {
				return false // Header value matches the pattern when it should not
			}
		}
	}
	return true // All rules satisfied
}
