package job_test

import (
	"testing"

	"github.com/skipikash/jenkins/job"
	"github.com/skipikash/jenkins/mock"
)

func TestGetInfo(t *testing.T) {
	r := mock.Client
	actual, err := job.GetInfo(r, mock.TestJobURL)
	if err != nil {
		t.Error(err)
	}
	if actual.DisplayName != "test_pipeline" {
		t.Fail()
	}
}

func TestGetRunInfo(t *testing.T) {
	r := mock.Client

	t.Run("Test for Running Job", func(t *testing.T) {
		actual, err := job.GetRunInfo(r, mock.TestRunningJobURL)
		if err != nil {
			t.Error(err)
		}
		if actual.Result != "RUNNING" {
			t.Fail()
		}
	})

	t.Run("Test for Successful Job", func(t *testing.T) {
		actual, err := job.GetRunInfo(r, mock.TestSuccessJobURL)
		if err != nil {
			t.Error(err)
		}
		if actual.Result != "SUCCESS" {
			t.Fail()
		}
	})

	t.Run("Test for Failed Job", func(t *testing.T) {
		actual, err := job.GetRunInfo(r, mock.TestFailureJobURL)
		if err != nil {
			t.Error(err)
		}
		if actual.Result != "FAILURE" {
			t.Fail()
		}
	})

	t.Run("Test for Aborted Job", func(t *testing.T) {
		actual, err := job.GetRunInfo(r, mock.TestAbortedJobURL)
		if err != nil {
			t.Error(err)
		}
		if actual.Result != "ABORTED" {
			t.Fail()
		}
	})

}

func TestGetConfig(t *testing.T) {
	r := mock.Client
	actual := job.Config{}
	err := job.GetConfig(r, mock.TestJobURL, &actual)
	if err != nil {
		t.Error(err)
	}
	if actual.Disabled != "false" {
		t.Fail()
	}
}

func TestStart(t *testing.T) {
	r := mock.Client

	t.Run("Test for Parameterized Job", func(t *testing.T) {
		params := job.Parameters{
			job.Parameter{
				Name:  "GREETING",
				Value: "Hello World",
			},
		}
		runningJob, err := job.Start(r, params, mock.TestJobURL)
		if err != nil {
			t.Error(err)
		}
		if runningJob != mock.TestRunningJobURL {
			t.Fail()
		}
	})

	t.Run("Test for no Parameters", func(t *testing.T) {
		runningJob, err := job.Start(r, nil, mock.TestJobURL)
		if err != nil {
			t.Error(err)
		}
		if runningJob != mock.TestRunningJobURL {
			t.Fail()
		}
	})
}

func TestAbort(t *testing.T) {
	if err := job.Abort(mock.Client, mock.TestRunningJobURL); err != nil {
		t.Error(err)
	}
}

func TestGetInputRequest(t *testing.T) {
	r := mock.Client
	inputActions, err := job.GetInputRequest(r, mock.TestPendingInputJobURL)
	if err != nil {
		t.Error(err)
	}
	if inputActions[0].Message != "Do you want to proceed?" {
		t.Fail()
	}
}

func TestIsRequestingInput(t *testing.T) {
	r := mock.Client
	isRequesting := job.IsRequestingInput(r, mock.TestPendingInputJobURL)
	if !isRequesting {
		t.Fail()
	}
}

func TestPassProceedInput(t *testing.T) {
	r := mock.Client
	err := job.PassProceedInput(r, nil, mock.TestPendingInputJobURL)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
