package examples

import (
	"fmt"

	"github.com/skipikash/jenkins"
	"github.com/skipikash/jenkins/job"
)

func getJenkinsJobInfo(){

}

func getJenkinsJobRunInfo(){

}

func getJenkinsJobConfig(){

}

func startJenkinsJob(){
	c := jenkins.Client{}
	params := job.Parameters{
		job.Parameter{Name: "SomeParameter", Value: "SomeValue"},
	}
	url := "http://myjenkinsinstance.com/job/my_develop_job"
	runningJob, err := job.Start(c, params, url)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("Started %s", runningJob)
}