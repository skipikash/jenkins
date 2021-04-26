package jenkins

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
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
	JenkinsURL      string
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
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(&responseStruct)
}

// RequestXML makes a request to Jenkins and decodes the xml response into the responseStruct
func RequestXML(r Requester, method, url string, body io.Reader, responseStruct interface{}) error {
	res, err := r.Request(method, url, body)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return xml.NewDecoder(res.Body).Decode(responseStruct)
}

//Request makes requests to your Jenkins instance and returns an *http.Response and an error
func (c Client) Request(method, url string, body io.Reader) (*http.Response, error) {
	// Set default values for a Client
	if c.Context == nil {
		c.Context = context.Background()
	}
	if c.JenkinsUsername == "" {
		c.JenkinsUsername = os.Getenv("JENKINS_USERNAME")
	}
	if c.JenkinsAPIToken == "" {
		c.JenkinsAPIToken = os.Getenv("JENKINS_API_TOKEN")
	}
	if c.JenkinsURL == "" {
		c.JenkinsURL = os.Getenv("JENKINS_URL")
	}

	// Create http.Request
	req, err := http.NewRequestWithContext(c.Context, method, url, body)
	if err != nil {
		return nil, err
	}
	req.Close = true

	// Set Basic Auth
	req.SetBasicAuth(c.JenkinsUsername, c.JenkinsAPIToken)

	// Set Headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	//Perform request and return http.Response
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return res, fmt.Errorf("%s returned from %s %s", res.Status, method, url)
	}
	return res, err
}
