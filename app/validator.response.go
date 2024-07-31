package app

import (
	"fmt"
	"net/http"
)

func (m *Simopi) ValidateResponseCode(i int) error {
	if len(http.StatusText(m.Scenarios[i].Response.Code)) == 0 {
		return fmt.Errorf("scenario[%d].response.code - invalid http status code", i)
	}

	return nil
}

func (m *Simopi) ValidateResponseDelay(i int) error {
	delay := m.Scenarios[i].Response.Delay

	if delay.DelayType == "" {
		return nil
	}

	switch delay.DelayType {
	case DELAY_TYPE_FIXED:
		if delay.Duration == 0 {
			return fmt.Errorf("scenario[%d].response.delay.duration - haven't been set (remove delay object if unused)", i)
		}
		if delay.Duration < 0 {
			return fmt.Errorf("scenario[%d].response.delay.duration - cannot be below 0", i)
		}
	case DELAY_TYPE_RANGE:
		if delay.MaxDuration < delay.MinDuration {
			return fmt.Errorf("scenario[%d].response.delay - max duration cannot be lower than min duration", i)
		}
		if delay.MinDuration < 0 {
			return fmt.Errorf("scenario[%d].response.delay.min_duration - cannot be below 0", i)
		}
	default:
		return fmt.Errorf("scenario[%d].response.delay.delay_type - invalid type [ FIXED | RANGE ]", i)
	}

	return nil
}

func (m *Simopi) ValidateResponseHeader(i int) error {
	for k, v := range m.Scenarios[i].Response.Header {
		if _, ok := v.(string); !ok {
			return fmt.Errorf("scenario[%d].response.header[%s].value - must be string", i, k)
		}
	}

	return nil
}

func (m *Simopi) ValidateResponseBody(i int) error {
	if !IsValidJson(m.Scenarios[i].Response.Body) {
		return fmt.Errorf("scenario[%d].response.body - not a valid json", i)
	}

	return nil
}
