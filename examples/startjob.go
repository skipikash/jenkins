package examples

import (
	"fmt"

	"github.com/skipikash/jenkins"
	"github.com/skipikash/jenkins/job"
)

func startMyJob() {
	// create client
	// Note: to create JenkinsAPIToken view cloudbees documentation
	jc := jenkins.Client{
		JenkinsUsername: "username",
		JenkinsAPIToken: "apiToken",
	}

	// set parameters if job requires parameters, otherwise params := nil
	params := job.Parameters{
		job.Parameter{
			Name:  "MYPARAM",
			Value: "Hello World",
		},
	}

	myJobURL := "http://myjenkins:8080.mycompany.com/job/myJob"
	runningJobUrl, _ := job.Start(jc, params, myJobURL)

	fmt.Printf("%s is RUNNING!", runningJobUrl)
}
