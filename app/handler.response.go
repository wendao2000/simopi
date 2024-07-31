package app

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

func WriteResponse(w http.ResponseWriter, res *Response, err error) {
	if res == nil {
		body := "internal server error"
		if err != nil {
			body = err.Error()
		}

		res = &Response{
			Code: http.StatusInternalServerError,
			Body: body,
		}
	}

	var duration int
	switch res.Delay.DelayType {
	case DELAY_TYPE_RANGE:
		duration = res.Delay.MinDuration + rand.Intn(res.Delay.MaxDuration-res.Delay.MinDuration)
	case DELAY_TYPE_FIXED:
		duration = res.Delay.Duration
	}
	if duration > 0 {
		time.Sleep(time.Duration(duration) * time.Millisecond)
	}

	if len(res.Header) > 0 {
		for k, v := range res.Header {
			w.Header().Set(k, v.(string))
		}
	}

	if len(w.Header().Get("Content-Type")) == 0 {
		if IsValidJson(res.Body) {
			w.Header().Set("Content-Type", "application/json")
		} else {
			w.Header().Set("Content-Type", "text/plain")
		}
	}

	w.WriteHeader(res.Code)
	if rb, ok := res.Body.(map[string]interface{}); ok {
		_rb, _ := json.Marshal(rb)
		w.Write(_rb)
	} else if rb, ok := res.Body.([]byte); ok {
		w.Write(rb)
	} else if rb, ok := res.Body.(string); ok {
		w.Write([]byte(rb))
	}
}

func NewResponse(code int, msg interface{}) Response {
	if len(http.StatusText(code)) == 0 {
		code = http.StatusInternalServerError
	}

	return Response{
		Code: code,
		Body: msg,
	}
}
