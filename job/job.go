package job

import (
	"fmt"

	"github.com/skipikash/jenkins"
)

var (
	apiJSON      = "%s/api/json"
	configXML    = "%s/config.xml"
	apiBuild     = "%s/build"
	apiPassInput = "%s/Input"
)

// GetInfo returns information about a job.
// Note: To return information for a specific job run use GetRunInfo()
func GetInfo(r jenkins.Requester, jobURL string) (*Info, error) {
	jobInfo := Info{}
	url := fmt.Sprintf(apiJSON, jobURL)
	err := jenkins.RequestJSON(r, "GET", url, nil, &jobInfo)
	if err != nil {
		return nil, err
	}
	return &jobInfo, nil
}

// GetRunInfo returns information for a specific job run
func GetRunInfo(r jenkins.Requester, jobRunURL string) (*RunInfo, error) {
	jobRunInfo := RunInfo{}
	url := fmt.Sprintf(apiJSON, jobRunURL)
	err := jenkins.RequestJSON(r, "GET", url, nil, &jobRunInfo)
	if err != nil {
		return nil, err
	}
	return &jobRunInfo, nil
}

// GetConfig returns configuration for a given job
func GetConfig(r jenkins.Requester, jobURL string, configStruct interface{}) error {
	url := fmt.Sprintf(configXML, jobURL)
	return jenkins.RequestXML(r, "GET", url, nil, configStruct)
}

// Start starts a job and returns the run URL
func Start(r jenkins.Requester, params Parameters, jobURL string) (string, error) {
	url := fmt.Sprintf(apiBuild, jobURL)
	err := jenkins.RequestJSON(r, "POST", url, params.GetPostBody(), params)
	if err != nil {
		return "", err
	}
	jobRunURL := ""
	// TODO: get jobURL from queue
	return jobRunURL, err
}


// IsRequestingInput checks to see if a running job is requesting input
func IsRequestingInput(r jenkins.Requester, jobRunURL string) bool {
	// check to see if job is pendingInput
	// if so, post input
	return false
}


// PassInput passes input to a running job waiting for input
func PassInput(r jenkins.Requester, jobRunURL string, input []byte) error {
	// post input
	return nil
}

