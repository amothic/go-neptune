package neptune

import (
	"encoding/json"
	"fmt"
)

const (
	StatusSuccess                  = 200
	StatusNoContent                = 204
	StatusPartialContent           = 206
	StatusUnauthorized             = 401
	StatusAuthenticate             = 407
	StatusClientSerializationError = 497
	StatusMalformedRequest         = 498
	StatusInvalidRequestArguments  = 499
	StatusServerError              = 500
	StatusScriptEvaluationError    = 597
	StatusServerTimeout            = 598
	StatusServerSerializationError = 599
)

type Response struct {
	RequestID string         `json:"requestId"`
	Status    ResponseStatus `json:"status"`
	Result    ResponseResult `json:"result"`
}

type ResponseStatus struct {
	Message    string                 `json:"message"`
	Code       int                    `json:"code"`
	Attributes map[string]interface{} `json:"attributes"`
}

type ResponseResult struct {
	Data json.RawMessage        `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

func (r *Response) detectError() error {
	switch r.Status.Code {
	case StatusSuccess, StatusNoContent, StatusPartialContent:
		return nil
	case StatusUnauthorized:
		return fmt.Errorf("unauthorized")
	case StatusAuthenticate:
		return fmt.Errorf("authenticate")
	case StatusClientSerializationError:
		return fmt.Errorf("client serialization error")
	case StatusMalformedRequest:
		return fmt.Errorf("malformed request")
	case StatusInvalidRequestArguments:
		return fmt.Errorf("invalid request arguments")
	case StatusServerError:
		return fmt.Errorf("server error")
	case StatusScriptEvaluationError:
		return fmt.Errorf("script evaluation error")
	case StatusServerTimeout:
		return fmt.Errorf("server timeout")
	case StatusServerSerializationError:
		return fmt.Errorf("serverserver serialization error")
	default:
		return fmt.Errorf("unknown error")
	}
}
