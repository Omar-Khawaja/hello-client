package hello

// PingOrder is used to send a request to the ping endpoint
func (c *Client) PingOrder() error {
	r, err := c.newRequest("GET", "/ping")
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
