package job_test

import (
	"fmt"
	"testing"

	"github.com/skipikash/jenkins/jenkins"
	"github.com/skipikash/jenkins/job"
)

func TestGetInfo(t *testing.T) {
	expected := &job.Info{Name: "pbrtool_container_develop_ALTDEV"}
	r := jenkins.MockRequester{DataFilePath: "../testdata/jobinfo.json"}
	actual, err := job.GetInfo(r, "")
	if err != nil || expected.Name != actual.Name {
		fmt.Println(err)
		t.Fail()
	}
}
