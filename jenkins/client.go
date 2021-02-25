package jenkins

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"os"
)

// Requester defines an interface used for making requests to Jenkins
type Requester interface {
	Request(method, url string, body io.Reader) (*http.Response, error)
}

// Client is a default struct that implements the Requester interface that can be used
// when performing requests to Jenkins
type Client struct {
	JenkinsUsername string
	JenkinsAPIToken string
	HTTPClient      *http.Client
	Context         context.Context
}

// RequestJSON makes a request to Jenkins and decodes the json response into the responseStruct
func RequestJSON(r Requester, method, url string, body io.Reader, responseStruct interface{}) error {
	res, err := r.Request(method, url, body)
	if err != nil {
		return err
	}
	return json.NewDecoder(res.Body).Decode(responseStruct)
}

// RequestXML makes a request to Jenkins and decodes the xml response into the responseStruct
func RequestXML(r Requester, method, url string, body io.Reader, responseStruct interface{}) error {
	res, err := r.Request(method, url, body)
	if err != nil {
		return err
	}
	return xml.NewDecoder(res.Body).Decode(responseStruct)
}

//Request makes requests to your Jenkins instance and returns an *http.Response and an error
func (c Client) Request(method, url string, body io.Reader) (*http.Response, error) {
	if c.Context == nil {
		c.Context = context.Background()
	}
	if c.JenkinsUsername == "" {
		c.JenkinsUsername = os.Getenv("JENKINS_USERNAME")
		c.JenkinsAPIToken = os.Getenv("JENKINS_API_TOKEN")
	}
	req, err := http.NewRequestWithContext(c.Context, method, url, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.JenkinsUsername, c.JenkinsAPIToken)
	req.Header.Add("content-type", "application/json")
	return c.HTTPClient.Do(req)
}
