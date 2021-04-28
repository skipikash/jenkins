package mock

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/skipikash/jenkins"
)

type Response struct {
	Bytes      []byte
	StatusCode int
	Status     string
	Header     http.Header
}

// RequestMocker is a RoundTripper used for mocking Jenkins API requests
type RequestMocker struct {
	RequestMap map[string]Response
}

// RoundTrip returns a predefined *http.Response and a nil error
func (rm RequestMocker) RoundTrip(req *http.Request) (*http.Response, error) {
	// create key
	key := fmt.Sprintf(keyFmt, req.Method, strings.ReplaceAll(req.URL.Path, "//", "/"))

	// if key is not in map return 404 Not Found response
	if res, ok := rm.RequestMap[key]; ok {
		return &http.Response{
			Body:       ioutil.NopCloser(bytes.NewReader(res.Bytes)),
			Header:     res.Header,
			Status:     res.Status,
			StatusCode: res.StatusCode,
		}, nil
	} else {
		return &http.Response{
			StatusCode: 404,
			Status:     "404 Not Found",
		}, nil
	}
}

// MockTransport
var MockTransport = &RequestMocker{
	RequestMap: map[string]Response{
		pipelineInfoResponseKey:         pipelineInfoResponse,
		runningPipelineInfoResponseKey:  runningPipelineInfoResponse,
		successPipelineInfoResponseKey:  successPipelineInfoResponse,
		failedPipelineInfoResponseKey:   failedPipelineInfoResponse,
		abortedPipelineInfoResponseKey:  abortedPipelineInfoResponse,
		configXMLResponseKey:            configXMLResponse,
		buildResponseKey:                buildResponse,
		buildWithParamsResponseKey:      buildWithParamsResponse,
		queueLocationInfoResponseKey:    queueLocationInfoResponse,
		pendingInputPipelineResponseKey: pendingInputPipelineResponse,
		submitInputResponseKey:          submitInputResponse,
		abortResponseKey:                abortResponse,
	},
}

// MockClient is a Client used for mocking Jenkins API requests
var Client = jenkins.Client{
	JenkinsURL:      testJenkinsURL,
	JenkinsUsername: "mockUsername",
	JenkinsAPIToken: "mockAPIToken",
	HTTPClient: &http.Client{
		Transport: *MockTransport,
	},
}

// To add a mock response to the built in mock client pass in
func AddResponse(jc *jenkins.Client, method string, urlPath string, resp Response) {
	key := fmt.Sprintf(keyFmt, strings.ToUpper(method), urlPath)
	jc.HTTPClient.Transport.(RequestMocker).RequestMap[key] = resp
}
