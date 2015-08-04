package goblitline

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	POST_URL = "http://api.blitline.com/job"
)

func Post(body io.Reader) (*Response, error) {
	res, err := http.Post(POST_URL, "application/json", body)
	if err != nil {
		return nil, err
	}

	response := &Response{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return response, nil
}
