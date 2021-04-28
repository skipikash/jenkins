package examples

import "github.com/skipikash/jenkins"

// create client
// Note: to get your JenkinsAPIToken view cloudbees documentation
var jc = jenkins.Client{
	JenkinsUsername: "username",
	JenkinsAPIToken: "apiToken",
}
