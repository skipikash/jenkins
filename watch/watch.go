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
func JobWithInput(r jenkins.Requester, interval time.Duration, params job.Parameters, jobRunURL string) (string, error) {
	return JobWithFunc(r, jobRunURL, interval, func() error {
		if job.IsRequestingInput(r, jobRunURL) {
			return job.PassProceedInput(r, params, jobRunURL)
		}
		return nil
	})
}

// JobWithFunc checks a jenkins job every interval and runs f every interval until the job completes.
// returns status and err
func JobWithFunc(r jenkins.Requester, jobRunURL string, interval time.Duration, f func() error) (string, error) {
	jobStatus := "RUNNING"
	for running := true; running; running = (jobStatus == "RUNNING") {
		info, err := job.GetRunInfo(r, jobRunURL)
		if err != nil {
			return "", err
		}
		jobStatus = info.Result
		err = f()
		if err != nil {
			return jobStatus, err
		}
	}
	return jobStatus, nil
}
