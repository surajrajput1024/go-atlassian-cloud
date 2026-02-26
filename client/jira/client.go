package jira

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	atlassian "github.com/surajsinghrajput/go-atlassian-cloud/client"
)

type Client struct {
	c *atlassian.Client
}

func New(c *atlassian.Client) *Client {
	return &Client{c: c}
}

func (j *Client) path(segments ...string) string {
	base := j.c.RestAPIURL()
	for _, s := range segments {
		base = strings.TrimSuffix(base, "/") + "/" + strings.TrimPrefix(s, "/")
	}
	return base
}

func (j *Client) getJSON(path string, out interface{}) error {
	resp, err := j.c.Get(path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return j.apiErr(resp)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, out)
}

func (j *Client) post(path string, body interface{}, out interface{}) error {
	return j.doWithBody(http.MethodPost, path, body, out)
}

func (j *Client) put(path string, body interface{}, out interface{}) error {
	return j.doWithBody(http.MethodPut, path, body, out)
}

func (j *Client) doWithBody(method, path string, body interface{}, out interface{}) error {
	var reqBody []byte
	if body != nil {
		var err error
		reqBody, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}
	var bodyReader io.Reader
	if len(reqBody) > 0 {
		bodyReader = bytes.NewReader(reqBody)
	}
	req, err := http.NewRequest(method, path, bodyReader)
	if err != nil {
		return err
	}
	resp, err := j.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return j.apiErrFromBody(resp.StatusCode, respBody)
	}
	if out != nil && len(respBody) > 0 {
		return json.Unmarshal(respBody, out)
	}
	return nil
}

func (j *Client) delete(path string) error {
	req, err := http.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	resp, err := j.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	return j.apiErr(resp)
}

func (j *Client) apiErr(resp *http.Response) error {
	body, _ := io.ReadAll(resp.Body)
	return j.apiErrFromBody(resp.StatusCode, body)
}

func (j *Client) apiErrFromBody(code int, body []byte) error {
	ae := &atlassian.APIError{StatusCode: code, Body: body}
	var parsed struct {
		ErrorMessages []string          `json:"errorMessages"`
		Errors        map[string]string `json:"errors"`
	}
	_ = json.Unmarshal(body, &parsed)
	ae.ErrorMessages = parsed.ErrorMessages
	ae.Errors = parsed.Errors
	return ae
}
