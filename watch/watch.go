package watch

import (
	"time"

	"github.com/skipikash/jenkins"
	"github.com/skipikash/jenkins/job"
)

// Job checks a jenkins job every interval until it completes and returns the final status
func Job(r jenkins.Requester, jobRunURL string, interval time.Duration) (string, error) {
	return JobWithFunc(r, jobRunURL, interval, nil)
}

// JobWithInput checks a jenkins job every interval until it completes and passes input to a job if it is requesting input
func JobWithInput(r jenkins.Requester, jobRunURL string, interval time.Duration, input []byte) (string, error) {
	return JobWithFunc(r, jobRunURL, interval, func(){
		if job.IsRequestingInput(r, jobRunURL) {
			job.PassInput(r, jobRunURL, input)
		}
	})
}

// JobWithFunc checks a jenkins job every interval and runs f every interval until the job completes
func JobWithFunc(r jenkins.Requester, jobRunURL string, interval time.Duration, f func()) (string, error) {
	jobStatus := "RUNNING"
	for running := true; running; running = (jobStatus == "RUNNING") {
		info, err := job.GetRunInfo(r, jobRunURL)
		if err != nil {
			return "", err
		}
		jobStatus = info.Result
		f()
	}
	return jobStatus, nil
}