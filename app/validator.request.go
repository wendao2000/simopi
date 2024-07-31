package app

import (
	"fmt"
	"regexp"
)

func (m *Simopi) ValidateRequestHeader(i int) error {
	for j, hd := range m.Scenarios[i].Request.Header {
		if len(hd.RuleType) == 0 {
			return fmt.Errorf("scenario[%d].request.header[%d].rule_type - cannot be empty", i, j)
		}
		switch hd.RuleType {
		case RULE_TYPE_MATCH, RULE_TYPE_NOT_MATCH, RULE_TYPE_PATTERN, RULE_TYPE_NOT_PATTERN:
		default:
			return fmt.Errorf("scenario[%d].request.header[%d].rule_type - invalid type [ MATCH | NOT_MATCH | PATTERN | NOT_PATTERN ]", i, j)
		}

		if len(hd.Key) == 0 {
			return fmt.Errorf("scenario[%d].request.header[%d].key - cannot be empty", i, j)
		}

		if hd.RuleType == RULE_TYPE_PATTERN {
			if _, err := regexp.Compile(hd.Value); err != nil {
				return fmt.Errorf("scenario[%d].request.header[%d].value - not a valid regexp pattern", i, j)
			}
		}
	}

	return nil
}

func (m *Simopi) ValidateRequestBody(i int) error {
	bodyRules := m.Scenarios[i].Request.Body

	if m.Scenarios[i].Request.BodyJson != nil {
		switch val := m.Scenarios[i].Request.BodyJson.(type) {
		case map[string]interface{}:
			bodyRules = append(bodyRules, MapToRule(val)...)
		default:
			return fmt.Errorf("scenario[%d].request.body_json - invalid format", i)
		}
	}

	for j, bd := range bodyRules {
		if len(bd.RuleType) == 0 {
			return fmt.Errorf("scenario[%d].request.body[%d].rule_type - cannot be empty", i, j)
		}
		switch bd.RuleType {
		case RULE_TYPE_MATCH, RULE_TYPE_NOT_MATCH, RULE_TYPE_PATTERN, RULE_TYPE_NOT_PATTERN:
		default:
			return fmt.Errorf("scenario[%d].request.header[%d].rule_type - invalid type [ MATCH | NOT_MATCH | PATTERN | NOT_PATTERN ]", i, j)
		}

		if len(bd.Key) == 0 {
			return fmt.Errorf("scenario[%d].request.body[%d].key - cannot be empty", i, j)
		}
		if !keyRule.MatchString(bd.Key) {
			return fmt.Errorf("scenario[%d].request.body[%d].key - may only contain alphanumeric and -._", i, j)
		}

		if bd.RuleType == RULE_TYPE_PATTERN {
			if _, err := regexp.Compile(bd.Value); err != nil {
				return fmt.Errorf("scenario[%d].request.body[%d].value - not a valid regexp pattern", i, j)
			}
		}
	}

	m.Scenarios[i].Request.Body = bodyRules
	m.Scenarios[i].Request.BodyJson = nil

	return nil
}
