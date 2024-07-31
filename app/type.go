package app

type (
	Cred struct {
		User     string `json:"user"`
		Endpoint string `json:"endpoint"`
		Method   string `json:"method"` // [ GET | POST | PATCH | PUT | DELETE ]
	}

	Simopi struct {
		Cred
		Signature Signature  `json:"signature"` // signature
		Scenarios []Scenario `json:"scenarios"` // scenarios to check
	}

	Signature struct {
		Enabled   bool   `json:"enabled,omitempty"`    // whether to do signature checking
		HeaderKey string `json:"header_key,omitempty"` // header where the signature is
		Method    string `json:"method,omitempty"`     // [ MD5 | SHA1 | SHA256 | PKCS1v15 | PSS ]
		PublicKey string `json:"public_key,omitempty"` // for signature method PKCS1v15 and PSS
		KeyType   string `json:"key_type,omitempty"`   // [ PKCS1 - RSA PUBLIC KEY | PKIX - PUBLIC KEY ]
	}

	Scenario struct {
		Request  Request  `json:"request"` // a.k.a. rule
		Response Response `json:"response"`
	}

	Request struct {
		Header   []Rule      `json:"header,omitempty"` // request header
		Body     []Rule      `json:"body,omitempty"`   // request body
		BodyJson interface{} `json:"body_json"`
	}

	Rule struct {
		RuleType string `json:"rule_type"` // [ MATCH | NOT_MATCH | PATTERN | NOT_PATTERN ]
		Key      string `json:"key"`       // json key
		Value    string `json:"value"`     // json value
	}

	Response struct {
		Code   int                    `json:"code"`   // response http status code
		Delay  ResponseDelay          `json:"delay"`  // response delay
		Header map[string]interface{} `json:"header"` // response header
		Body   interface{}            `json:"body"`   // response body
	}

	ResponseDelay struct {
		DelayType   string `json:"delay_type"`             // [ FIXED | RANGE ]
		Duration    int    `json:"duration"`               // in ms, only available on FIXED delay type
		MinDuration int    `json:"min_duration,omitempty"` // in ms, only available on RANGE delay type
		MaxDuration int    `json:"max_duration,omitempty"` // in ms, only available on RANGE delay type
	}
)
