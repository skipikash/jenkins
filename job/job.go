package job

import (
	"fmt"
	"time"

	"github.com/skipikash/jenkins"
)

var (
	apiJSON                = "%s/api/json"
	apiConfigXML           = "%s/config.xml"
	apiBuild               = "%s/build"
	apiBuildWithParams     = "%s/buildWithParameters?%s"
	apiPassInput           = "%s/Input"
	apiPendingInputActions = "%s/wfapi/pendingInputActions"
	apiSubmitInput         = "%s/wfapi/inputSubmit?inputId=%s&proceed=%s&%s"
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
// Note: Depending on your plugins installed in your jenkins instance you will need
// to define a configStruct such that the config.xml response can be unmarshalled. It is
// recommended to use an online "XML to go struct" generator
func GetConfig(r jenkins.Requester, jobURL string, configStruct interface{}) error {
	url := fmt.Sprintf(apiConfigXML, jobURL)
	return jenkins.RequestXML(r, "GET", url, nil, configStruct)
}

// Start starts a job, pauses so the jenkins queue populates data and returns the run URL
func Start(r jenkins.Requester, params Parameters, jobURL string) (string, error) {
	var url string
	if params == nil || len(params) == 0 {
		// url for no parameters
		url = fmt.Sprintf(apiBuild, jobURL)
	} else {
		// url for parameterized job
		url = fmt.Sprintf(apiBuildWithParams, jobURL, params.GetURLBuildArgs())
	}

	// send request to start job
	resp, err := r.Request("POST", url, nil)
	if err != nil {
		return "", err
	}

	// wait for queue to populate location data
	time.Sleep(10 * time.Second)

	// get runnning job URL
	locationResp := Location{}
	url = fmt.Sprintf(apiJSON, resp.Header.Get("Location"))
	err = jenkins.RequestJSON(r, "GET", url, nil, &locationResp)
	if err != nil {
		return "", err
	}
	jobRunURL := locationResp.Executable.URL

	// return running job url
	return jobRunURL, err
}

// GetInputRequest gets input requests when job is pending input
func GetInputRequest(r jenkins.Requester, jobRunURL string) (PendingInputActions, error) {
	url := fmt.Sprintf(apiPendingInputActions, jobRunURL)
	inputActions := PendingInputActions{}
	err := jenkins.RequestJSON(r, "GET", url, nil, &inputActions)
	if err != nil || len(inputActions) == 0 {
		return inputActions, fmt.Errorf("Job is not waiting for input")
	}
	return inputActions, nil
}

// IsRequestingInput checks to see if a running job is requesting input
func IsRequestingInput(r jenkins.Requester, jobRunURL string) bool {
	if inputActions, err := GetInputRequest(r, jobRunURL); err != nil {
		return false
	} else {
		return len(inputActions) > 0
	}
}

// PassProceedInput passes input to a running job waiting for input
func PassProceedInput(r jenkins.Requester, params Parameters, jobRunURL string) error {

	inputActions, err := GetInputRequest(r, jobRunURL)
	if err != nil {
		return err
	}

	// populate url
	inputID := inputActions[0].ID
	proceedText := inputActions[0].ProceedText
	jsonInput := "json={}"
	if params != nil && len(params) != 0 {
		jsonInput = params.GetURLEncodedJSON()
	}

	// send input
	url := fmt.Sprintf(apiSubmitInput, jobRunURL, inputID, proceedText, jsonInput)
	resp, err := r.Request("POST", url, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return fmt.Errorf("%s status returned from passing proceed input", resp.Status)
	}

	return nil
}
