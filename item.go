package hello

import (
	"net/url"
)

func (c *Client) AppendItem(item string) error {
	urlstring, err := url.Parse("/items")
	if err != nil {
		return err
	}
	q := urlstring.Query()
	q.Add("item", item)
	urlstring.RawQuery = q.Encode()
	r, err := c.newRequest("POST", urlstring.String())
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
