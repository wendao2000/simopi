package app

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
)

var (
	userRule     = regexp.MustCompile(`^[\w\d]+$`)
	endpointRule = regexp.MustCompile(`^[\w\d-_\/]+$`)
)

func (m *Simopi) ValidateUser() error {
	if len(m.User) == 0 {
		return errors.New("user cannot be empty")
	}

	if len(m.User) > 16 {
		return errors.New("user cannot be longer than 16 characters")
	}

	if !userRule.MatchString(m.User) {
		return errors.New("user may only contain alphanumeric")
	}

	return nil
}

func (m *Simopi) ValidateEndpoint() error {
	m.Endpoint = strings.TrimPrefix(m.Endpoint, "/")
	if len(m.Endpoint) == 0 {
		return errors.New("endpoint - cannot be empty")
	}
	if !endpointRule.MatchString(m.Endpoint) {
		return errors.New("endpoint - may only contain alphaumeric, and -_/")
	}

	return nil
}

func (m *Simopi) ValidateMethod() error {
	if len(m.Method) == 0 {
		return errors.New("method cannot be empty")
	}

	switch m.Method {
	case http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete:
	default:
		return errors.New("invalid method, must be either [ GET | POST | PATCH | PUT | DELETE ]")
	}

	return nil
}
