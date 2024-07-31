package app

import (
	"errors"
	"regexp"
)

var (
	keyRule = regexp.MustCompile(`^[\w\d.-_]+$`)
)

func (m *Simopi) ValidateRequest() error {
	var err error
	if err = m.ValidateUser(); err != nil {
		return err
	}

	if err = m.ValidateEndpoint(); err != nil {
		return err
	}

	if err = m.ValidateMethod(); err != nil {
		return err
	}

	if len(m.Scenarios) == 0 {
		return errors.New("scenarios cannot be empty")
	}

	for i := range m.Scenarios {
		if err = m.ValidateRequestHeader(i); err != nil {
			return err
		}

		if err = m.ValidateRequestBody(i); err != nil {
			return err
		}

		if err = m.ValidateResponseCode(i); err != nil {
			return err
		}

		if err = m.ValidateResponseDelay(i); err != nil {
			return err
		}

		if err = m.ValidateResponseHeader(i); err != nil {
			return err
		}

		if err = m.ValidateResponseBody(i); err != nil {
			return err
		}
	}

	return nil
}
