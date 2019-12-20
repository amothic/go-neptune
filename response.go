package neptune

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	RequestID string `json:"requestId"`
	Status    struct {
		Message    string                 `json:"message"`
		Code       int                    `json:"code"`
		Attributes map[string]interface{} `json:"attributes"`
	} `json:"status"`
	Result struct {
		Data json.RawMessage        `json:"data"`
		Meta map[string]interface{} `json:"meta"`
	} `json:"result"`
}

func (r *Response) isError() bool {
	switch r.Status.Code {
	case StatusSuccess, StatusNoContent, StatusPartialContent:
		return false
	default:
		return true
	}
}

func (r *Response) Error() error {
	if r.isError() {
		return fmt.Errorf("gremlin: code=%d, message=%q", r.Status.Code, r.Status.Message)
	}
	return nil
}
