package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func GetConfig(w http.ResponseWriter, r *http.Request) {
	var (
		res Response
		err error
	)

	defer WriteResponse(w, &res, err)

	if r.Method != http.MethodGet {
		res = NewResponse(http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	q, _ := url.ParseQuery(r.URL.RawQuery)

	user := q.Get("user")
	if len(user) == 0 {
		res = NewResponse(http.StatusBadRequest, "user cannot be empty")
		return
	}

	// Get cache
	val, ok := cache[user]
	if !ok || len(val) == 0 {
		res = NewResponse(http.StatusOK, "{}")
		return
	}

	// Flatten response
	var resp []Simopi
	for _, v := range val {
		resp = append(resp, v)
	}

	res = NewResponse(http.StatusOK, map[string]interface{}{"data": resp})
}

func CreateConfig(w http.ResponseWriter, r *http.Request) {
	var (
		res Response
		err error
	)

	defer WriteResponse(w, &res, err)

	if r.Method != http.MethodPost {
		res = NewResponse(http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	defer r.Body.Close()

	var m Simopi
	b, _ := io.ReadAll(r.Body)
	if err = json.Unmarshal(b, &m); err != nil {
		res = NewResponse(http.StatusBadRequest, "invalid request body")
		return
	}

	if err = m.ValidateRequest(); err != nil {
		res = NewResponse(http.StatusBadRequest, err.Error())
		return
	}

	SetCache(m)

	res = NewResponse(http.StatusCreated, fmt.Sprintf("/%s/%s", m.User, m.Endpoint))
}

func DeleteConfig(w http.ResponseWriter, r *http.Request) {
	var (
		res Response
		err error
	)

	defer WriteResponse(w, &res, err)

	if r.Method != http.MethodDelete {
		res = NewResponse(http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	defer r.Body.Close()

	var c Cred
	b, _ := io.ReadAll(r.Body)
	if err = json.Unmarshal(b, &c); err != nil {
		res = NewResponse(http.StatusBadRequest, "invalid request body")
		return
	}

	DeleteCacheByCred(c)

	res = NewResponse(http.StatusNoContent, nil)
}

func MatchmakeConfig(w http.ResponseWriter, r *http.Request) {
	var (
		res Response
		err error
	)

	defer WriteResponse(w, &res, err)

	paths := strings.SplitN(strings.ToLower(r.URL.Path), "/", 3)[1:]
	if len(paths) < 2 {
		res = NewResponse(http.StatusNotFound, "not found")
		return
	}
	defer r.Body.Close()

	b, _ := io.ReadAll(r.Body)

	rb := make(map[string]interface{})
	if len(b) > 0 {
		if err := json.Unmarshal(b, &rb); err != nil {
			res = NewResponse(http.StatusBadRequest, "invalid request body")
			return
		}
	}

	cred := Cred{
		User:     paths[0],
		Endpoint: paths[1],
		Method:   r.Method,
	}

	m, ok := GetCacheByCred(cred)
	if !ok {
		res = NewResponse(http.StatusNotFound, "no mock found")
		return
	}

	if err = m.CheckSignature(b); err != nil {
		// TODO implement failed scenario
		res = NewResponse(http.StatusUnauthorized, "invalid signature")
		return
	}

	var found bool

	for _, scn := range m.Scenarios {
		header := CheckHeader(scn.Request.Header, r.Header)
		if !header {
			continue
		}

		body := CheckBody(scn.Request.Body, rb)
		if !body {
			continue
		}

		res = scn.Response
		found = true
		break
	}

	if !found { // no scenario found for this request
		res = NewResponse(http.StatusNotFound, "no scenario found")
	}
}
