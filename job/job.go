package job

import "github.com/skipikash/jenkins/jenkins"

// GetInfo returns information about a job. 
// Note: To return information for a specific job run use GetRunInfo()
func GetInfo(r jenkins.Requester, jobURL string) {
	url := jobURL + "/api/json"
	err := jenkins.RequestJSON(r, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	return
}

// GetRunInfo asdf
func GetRunInfo(c jenkins.Requester, jobRunURL string) {
	return
}

// GetConfig asdf
func GetConfig(c jenkins.Requester, jobURL string) {
	return
}

// Start asdf
func Start(c jenkins.Requester, params Parameters, jobURL string) {
	return
}
