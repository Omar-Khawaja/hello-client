package hello

// PongOrder is used to send a request to the pong endpoint
func (c *Client) PongOrder() error {
	r, err := c.newRequest("GET", "/pong")
	if err != nil {
		return err
	}

	req, err := r.toHTTP()
	if err != nil {
		return err
	}

	_, err = c.httpClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}
