package jenkins

import (
	"context"
	"io"
	"net/http"
	"os"
	"encoding/json"
)



type Requester interface {
	Request(method, url string, body io.Reader) *http.Response, error
	HTTPClient()*http.Client
	Context()context.Context
}

type Client struct {
	JenkinsUsername string
	JenkinsAPIToken string
	Context context.Context
}

type (r Requester) ReqJSONUnmarshal(method, url string, body io.Reader, responseStruct interface{}) error{
	json.Unmar
	return nil
}

//Request makes requests to your Jenkins instance and returns an *http.Response and an error
func (c Client) Request(method, url string, body io.Reader) (*http.Response, error) {
	ctx := c.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.JenkinsUsername, c.JenkinsAPIToken)
	req.Header.Add("content-type", "application/json")
	returnc.HTTPClient().Do(req)
}

// HTTPClient returns a *http.Client that will be used for configuring the http.Client 
// that will be used to perform the requests to your Jenkins instance
func (c Client) HTTPClient()*http.Client{
	return *http.Client{}
}

// Context returns a context.Context that will be used when performing requests
func (c Client) Context() context.Context{
	return c.Context
}

