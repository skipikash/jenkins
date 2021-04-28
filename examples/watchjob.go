package examples

import (
	"fmt"
	"time"

	"github.com/skipikash/jenkins"
	"github.com/skipikash/jenkins/watch"
)

func watchMyJob() {
	// create client
	// Note: to create JenkinsAPIToken view cloudbees documentation
	jc := jenkins.Client{
		JenkinsUsername: "username",
		JenkinsAPIToken: "apiToken",
	}

	// URL of running job
	runningJobURL := "http://myjenkins:8080.mycompany.com/job/myJob/1"

	// watch a job until completion, polling it every minute
	status, _ := watch.Job(jc, runningJobURL, time.Minute)
	fmt.Printf("%s finished with a status of %s!", runningJobURL, status)
}
