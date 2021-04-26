package watch_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/skipikash/jenkins/mock"
	"github.com/skipikash/jenkins/watch"
)

func TestWatchJob(t *testing.T) {
	r := mock.Client
	status, err := watch.Job(r, mock.TestSuccessJobURL, time.Microsecond)
	if err != nil {
		t.Error(err)
	}
	if status != "SUCCESS" {
		t.Fail()
	}
}

func TestWatchJobWithFunc(t *testing.T) {
	r := mock.Client
	status, err := watch.JobWithFunc(r, mock.TestSuccessJobURL, time.Microsecond, func() error {
		var err error
		if true == true {
			t.Log("Hello Jenkins User")
		} else {
			err = fmt.Errorf("The world is ending")
		}
		return err
	})
	
	if err != nil {
		t.Error(err)
	}
	if status != "SUCCESS" {
		t.Fail()
	}
}
