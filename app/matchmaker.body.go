package app

// CheckBody evaluates a list of rules against a provided request body map.
// Returns false if any rule is not satisfied according to its type.
func CheckBody(scnBody []Rule, rBody map[string]interface{}) bool {
	for _, rule := range scnBody {
		switch rule.RuleType {
		case RULE_TYPE_MATCH:
			if !IsMatch(rBody, rule) {
				return false // Rule not matched
			}
		case RULE_TYPE_NOT_MATCH:
			if IsMatch(rBody, rule) {
				return false // Rule matched when it should not
			}
		case RULE_TYPE_PATTERN:
			if !IsMatchPattern(rBody, rule) {
				return false // Pattern not matched
			}
		case RULE_TYPE_NOT_PATTERN:
			if IsMatchPattern(rBody, rule) {
				return false // Pattern matched when it should not
			}
		}
	}
	return true // All rules satisfied
}
