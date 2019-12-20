package neptune

type Request struct {
	Gremlin string `json:"gremlin"`
}

func NewRequest(query string) *Request {
	return &Request{
		Gremlin: query,
	}
}
