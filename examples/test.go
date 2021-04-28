package examples

import (
	"testing"

	"github.com/skipikash/jenkins"
	"github.com/skipikash/jenkins/job"
	"github.com/skipikash/jenkins/mock"
)

func myFunc(jc *jenkins.Client, jobRunURL string) {
	job.GetRunInfo(jc, jobRunURL)
	job.Abort(jc, jobRunURL)
}

func TestMyFunc(t *testing.T) {
	myFunc(&mock.Client, mock.TestRunningJobURL)
}

// to add your own custom mock response
func TestSomethingElse(t *testing.T) {
	mock.AddResponse(&mock.Client, "GET", "/job/mycustomjobpath/api/json", mock.Response{
		StatusCode: 200,
		Bytes:      nil,
	})

	// now do something with mock.Client
}
