package main

import (
	"fmt"

	"github.com/skipikash/jenkins/jenkins"

	"github.com/skipikash/jenkins/job"
)

func main() {
	c := jenkins.Client{}
	params := job.Parameters{
		job.Parameter{Name: "FirstName", Value: "Steve"},
	}
	job.Start(c, params, "")
	fmt.Println(c)
}
