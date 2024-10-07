package redmine

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
)

type uploadResponse struct {
	Upload Upload `json:"upload"`
}

type Upload struct {
	Token       string `json:"token"`
	Filename    string `json:"filename"`
	ContentType string `json:"content_type"`
}

func (c *Client) Upload(filename string, userName ...string) (*Upload, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.endpoint+"/uploads.json?key="+c.apikey, bytes.NewBuffer(content))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	if len(userName) > 0 {
		req.Header.Set("X-Redmine-Switch-User", userName[0])	
	}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r uploadResponse
	if res.StatusCode != 201 {
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return &r.Upload, nil
}
