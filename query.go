package neptune

func (c *Client) Query(query string) (*Response, error) {
	req := NewRequest(query)
	res, err := c.post(req)
	if err != nil {
		return nil, err
	}
	var target Response
	err = c.decodeJSON(res, &target)
	if err != nil {
		return nil, err
	}
	return &target, target.Error()
}
