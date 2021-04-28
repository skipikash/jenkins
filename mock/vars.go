package mock

import (
	"fmt"
	"net/http"
)

// Test variables for MockClient and unit testing
var (
	keyFmt = "%s:%s"

	testJenkinsURL      = "http://mockjenkins:8080"
	testJobName         = "test_pipeline"
	jobPath             = "%s/job/%s"
	runPath             = "%s/job/%s/%d/"
	jobJSONPath         = "/job/%s/api/json"
	runJSONPath         = "/job/%s/%d/api/json"
	configXMLPath       = "/job/%s/config.xml"
	buildPath           = "/job/%s/build"
	buildWithParamsPath = "/job/%s/buildWithParameters"
	pendingInputPath    = "/job/%s/%d/wfapi/pendingInputActions"
	submitInputPath     = "/job/%s/%d/wfapi/inputSubmit"
	runningLocationPath = "/queue/item/61/api/json"
	runningLocationURL  = testJenkinsURL + "/queue/item/61"

	testJobURLPath                  = fmt.Sprintf(jobJSONPath, testJobName)
	testRunningPipelineURLPath      = fmt.Sprintf(runJSONPath, testJobName, 4)
	testSuccessPipelineURLPath      = fmt.Sprintf(runJSONPath, testJobName, 5)
	testFailurePipelineURLPath      = fmt.Sprintf(runJSONPath, testJobName, 6)
	testAbortedPipelineURLPath      = fmt.Sprintf(runJSONPath, testJobName, 7)
	testBuildPath                   = fmt.Sprintf(buildPath, testJobName)
	testBuildWithParamsPath         = fmt.Sprintf(buildWithParamsPath, testJobName)
	testPendingInputPipelineURLPath = fmt.Sprintf(pendingInputPath, testJobName, 8)
	testSubmitInputURLPath          = fmt.Sprintf(submitInputPath, testJobName, 8)
	testJobConfigXMLPath            = fmt.Sprintf(configXMLPath, testJobName)

	TestJobURL             = fmt.Sprintf(jobPath, testJenkinsURL, testJobName)
	TestRunningJobURL      = fmt.Sprintf(runPath, testJenkinsURL, testJobName, 4)
	TestSuccessJobURL      = fmt.Sprintf(runPath, testJenkinsURL, testJobName, 5)
	TestFailureJobURL      = fmt.Sprintf(runPath, testJenkinsURL, testJobName, 6)
	TestAbortedJobURL      = fmt.Sprintf(runPath, testJenkinsURL, testJobName, 7)
	TestPendingInputJobURL = fmt.Sprintf(runPath, testJenkinsURL, testJobName, 8)
)

// Set Test Responses
var (
	pipelineInfoResponseKey = fmt.Sprintf(keyFmt, "GET", testJobURLPath)
	pipelineInfoResponse    = Response{
		StatusCode: 200,
		Status:     "200 OK",
		Bytes:      []byte(`{"_class":"org.jenkinsci.plugins.workflow.job.WorkflowJob","actions":[{},{},{},{},{},{},{"_class":"org.jenkinsci.plugins.displayurlapi.actions.JobDisplayAction"},{},{},{"_class":"com.cloudbees.plugins.credentials.ViewCredentialsAction"}],"description":"","displayName":"test_pipeline","displayNameOrNull":null,"fullDisplayName":"test_pipeline","fullName":"test_pipeline","name":"test_pipeline","url":"http://mockjenkins:8080/job/test_pipeline/","buildable":true,"builds":[{"_class":"org.jenkinsci.plugins.workflow.job.WorkflowRun","number":3,"url":"http://mockjenkins:8080/job/test_pipeline/3/"},{"_class":"org.jenkinsci.plugins.workflow.job.WorkflowRun","number":2,"url":"http://mockjenkins:8080/job/test_pipeline/2/"},{"_class":"org.jenkinsci.plugins.workflow.job.WorkflowRun","number":1,"url":"http://mockjenkins:8080/job/test_pipeline/1/"}],"color":"red","firstBuild":{"_class":"org.jenkinsci.plugins.workflow.job.WorkflowRun","number":1,"url":"http://mockjenkins:8080/job/test_pipeline/1/"},"healthReport":[{"description":"Build stability: 1 out of the last 3 builds failed.","iconClassName":"icon-health-60to79","iconUrl":"health-60to79.png","score":66}],"inQueue":false,"keepDependencies":false,"lastBuild":{"_class":"org.jenkinsci.plugins.workflow.job.WorkflowRun","number":3,"url":"http://mockjenkins:8080/job/test_pipeline/3/"},"lastCompletedBuild":{"_class":"org.jenkinsci.plugins.workflow.job.WorkflowRun","number":3,"url":"http://mockjenkins:8080/job/test_pipeline/3/"},"lastFailedBuild":{"_class":"org.jenkinsci.plugins.workflow.job.WorkflowRun","number":3,"url":"http://mockjenkins:8080/job/test_pipeline/3/"},"lastStableBuild":{"_class":"org.jenkinsci.plugins.workflow.job.WorkflowRun","number":2,"url":"http://mockjenkins:8080/job/test_pipeline/2/"},"lastSuccessfulBuild":{"_class":"org.jenkinsci.plugins.workflow.job.WorkflowRun","number":2,"url":"http://mockjenkins:8080/job/test_pipeline/2/"},"lastUnstableBuild":null,"lastUnsuccessfulBuild":{"_class":"org.jenkinsci.plugins.workflow.job.WorkflowRun","number":3,"url":"http://mockjenkins:8080/job/test_pipeline/3/"},"nextBuildNumber":4,"property":[],"queueItem":null,"concurrentBuild":true,"resumeBlocked":false}`),
	}

	runningPipelineInfoResponseKey = fmt.Sprintf(keyFmt, "GET", testRunningPipelineURLPath)
	runningPipelineInfoResponse    = Response{
		StatusCode: 200,
		Status:     "200 OK",
		Bytes:      []byte(`{"_class":"org.jenkinsci.plugins.workflow.job.WorkflowRun","actions":[{"_class":"hudson.model.ParametersAction","parameters":[{"_class":"hudson.model.StringParameterValue","name":"Greeting","value":"Hello There"}]},{"_class":"hudson.model.CauseAction","causes":[{"_class":"hudson.model.Cause$UserIdCause","shortDescription":"Started by user admin","userId":"admin","userName":"admin"}]},{},{},{},{},{"_class":"org.jenkinsci.plugins.displayurlapi.actions.RunDisplayAction"},{"_class":"org.jenkinsci.plugins.pipeline.modeldefinition.actions.RestartDeclarativePipelineAction"},{},{"_class":"org.jenkinsci.plugins.workflow.job.views.FlowGraphAction"},{},{},{}],"artifacts":[],"building":false,"description":null,"displayName":"#5","duration":955,"estimatedDuration":1560,"executor":null,"fullDisplayName":"test_pipeline #5","id":"5","keepLog":false,"number":5,"queueId":8,"result":"RUNNING","timestamp":1618492532059,"url":"http://mockjenkins:8080/job/test_pipeline/5/","changeSets":[],"culprits":[],"nextBuild":{"number":6,"url":"http://mockjenkins:8080/job/test_pipeline/6/"},"previousBuild":{"number":4,"url":"http://mockjenkins:8080/job/test_pipeline/4/"}}`),
	}

	successPipelineInfoResponseKey = fmt.Sprintf(keyFmt, "GET", testSuccessPipelineURLPath)
	successPipelineInfoResponse    = Response{
		StatusCode: 200,
		Status:     "200 OK",
		Bytes:      []byte(`{"_class":"org.jenkinsci.plugins.workflow.job.WorkflowRun","actions":[{"_class":"hudson.model.ParametersAction","parameters":[{"_class":"hudson.model.StringParameterValue","name":"Greeting","value":"Hello There"}]},{"_class":"hudson.model.CauseAction","causes":[{"_class":"hudson.model.Cause$UserIdCause","shortDescription":"Started by user admin","userId":"admin","userName":"admin"}]},{},{},{},{},{"_class":"org.jenkinsci.plugins.displayurlapi.actions.RunDisplayAction"},{"_class":"org.jenkinsci.plugins.pipeline.modeldefinition.actions.RestartDeclarativePipelineAction"},{},{"_class":"org.jenkinsci.plugins.workflow.job.views.FlowGraphAction"},{},{},{}],"artifacts":[],"building":false,"description":null,"displayName":"#4","duration":2912,"estimatedDuration":1560,"executor":null,"fullDisplayName":"test_pipeline #4","id":"4","keepLog":false,"number":4,"queueId":6,"result":"SUCCESS","timestamp":1618492494269,"url":"http://mockjenkins:8080/job/test_pipeline/4/","changeSets":[],"culprits":[],"nextBuild":{"number":5,"url":"http://mockjenkins:8080/job/test_pipeline/5/"},"previousBuild":{"number":3,"url":"http://mockjenkins:8080/job/test_pipeline/3/"}}`),
	}

	failedPipelineInfoResponseKey = fmt.Sprintf(keyFmt, "GET", testFailurePipelineURLPath)
	failedPipelineInfoResponse    = Response{
		StatusCode: 200,
		Status:     "200 OK",
		Bytes:      []byte(`{"_class":"org.jenkinsci.plugins.workflow.job.WorkflowRun","actions":[{"_class":"hudson.model.ParametersAction","parameters":[{"_class":"hudson.model.StringParameterValue","name":"Greeting","value":"Hello There"}]},{"_class":"hudson.model.CauseAction","causes":[{"_class":"hudson.model.Cause$UserIdCause","shortDescription":"Started by user admin","userId":"admin","userName":"admin"}]},{},{},{"_class":"org.jenkinsci.plugins.displayurlapi.actions.RunDisplayAction"},{"_class":"org.jenkinsci.plugins.pipeline.modeldefinition.actions.RestartDeclarativePipelineAction"},{},{"_class":"org.jenkinsci.plugins.workflow.job.views.FlowGraphAction"},{},{},{}],"artifacts":[],"building":false,"description":null,"displayName":"#6","duration":18,"estimatedDuration":1560,"executor":null,"fullDisplayName":"test_pipeline #6","id":"6","keepLog":false,"number":6,"queueId":10,"result":"FAILURE","timestamp":1618492635439,"url":"http://mockjenkins:8080/job/test_pipeline/6/","changeSets":[],"culprits":[],"nextBuild":null,"previousBuild":{"number":5,"url":"http://mockjenkins:8080/job/test_pipeline/5/"}}`),
	}

	abortedPipelineInfoResponseKey = fmt.Sprintf(keyFmt, "GET", testAbortedPipelineURLPath)
	abortedPipelineInfoResponse    = Response{
		StatusCode: 200,
		Status:     "200 OK",
		Bytes:      []byte(`{"_class":"org.jenkinsci.plugins.workflow.job.WorkflowRun","actions":[{"_class":"hudson.model.ParametersAction","parameters":[{"_class":"hudson.model.StringParameterValue","name":"Greeting","value":"Hello There"}]},{"_class":"hudson.model.CauseAction","causes":[{"_class":"hudson.model.Cause$UserIdCause","shortDescription":"Started by user admin","userId":"admin","userName":"admin"}]},{},{},{"_class":"org.jenkinsci.plugins.displayurlapi.actions.RunDisplayAction"},{"_class":"org.jenkinsci.plugins.pipeline.modeldefinition.actions.RestartDeclarativePipelineAction"},{},{"_class":"org.jenkinsci.plugins.workflow.job.views.FlowGraphAction"},{},{},{}],"artifacts":[],"building":false,"description":null,"displayName":"#7","duration":18,"estimatedDuration":1560,"executor":null,"fullDisplayName":"test_pipeline #7","id":"7","keepLog":false,"number":7,"queueId":10,"result":"ABORTED","timestamp":1618492635439,"url":"http://mockjenkins:8080/job/test_pipeline/7/","changeSets":[],"culprits":[],"nextBuild":null,"previousBuild":{"number":6,"url":"http://mockjenkins:8080/job/test_pipeline/6/"}}`),
	}

	configXMLResponseKey = fmt.Sprintf(keyFmt, "GET", testJobConfigXMLPath)
	configXMLResponse    = Response{
		StatusCode: 200,
		Status:     "200 OK",
		Bytes:      []byte(`<flow-definition plugin="workflow-job@2.40"><actions><org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction plugin="pipeline-model-definition@1.8.4"/><org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction plugin="pipeline-model-definition@1.8.4"><jobProperties/><triggers/><parameters/><options/></org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction></actions><description/><keepDependencies>false</keepDependencies><properties><hudson.model.ParametersDefinitionProperty><parameterDefinitions><hudson.model.StringParameterDefinition><name>Greeting</name><description>parameter used for defining the greeting</description><defaultValue>Hello There</defaultValue><trim>false</trim></hudson.model.StringParameterDefinition></parameterDefinitions></hudson.model.ParametersDefinitionProperty></properties><definition class="org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition" plugin="workflow-cps@2.90"><script>pipeline { agent any stages { stage('Hello') { steps { echo 'Hello World' } } } }</script><sandbox>true</sandbox></definition><triggers/><disabled>false</disabled></flow-definition>`),
	}

	buildResponseKey = fmt.Sprintf(keyFmt, "POST", testBuildPath)
	buildResponse    = Response{
		StatusCode: 201,
		Status:     "201 Created",
		Header:     http.Header{"Location": []string{runningLocationURL}},
	}

	buildWithParamsResponseKey = fmt.Sprintf(keyFmt, "POST", testBuildWithParamsPath)
	buildWithParamsResponse    = Response{
		StatusCode: 201,
		Status:     "201 Created",
		Header:     http.Header{"Location": []string{runningLocationURL}},
	}

	queueLocationInfoResponseKey = fmt.Sprintf(keyFmt, "GET", runningLocationPath)
	queueLocationInfoResponse    = Response{
		StatusCode: 200,
		Status:     "200 OK",
		Bytes:      []byte(`{"_class":"hudson.model.Queue$LeftItem","actions":[{"_class":"hudson.model.ParametersAction","parameters":[{"_class":"hudson.model.StringParameterValue","name":"Greeting","value":"Hello There"}]},{"_class":"hudson.model.CauseAction","causes":[{"_class":"hudson.model.Cause$UserIdCause","shortDescription":"Started by user admin","userId":"admin","userName":"admin"}]}],"blocked":false,"buildable":false,"id":4,"inQueueSince":1619211753302,"params":"\u000aGreeting=Hello There","stuck":false,"task":{"_class":"org.jenkinsci.plugins.workflow.job.WorkflowJob","name":"test_pipeline","url":"http://mockjenkins:8080/job/test_pipeline/","color":"blue_anime"},"url":"queue/item/63/","why":null,"cancelled":false,"executable":{"_class":"org.jenkinsci.plugins.workflow.job.WorkflowRun","number":4,"url":"http://mockjenkins:8080/job/test_pipeline/4/"}}`),
	}

	pendingInputPipelineResponseKey = fmt.Sprintf(keyFmt, "GET", testPendingInputPipelineURLPath)
	pendingInputPipelineResponse    = Response{
		StatusCode: 200,
		Status:     "200 OK",
		Bytes:      []byte(`[{"id":"33d8fefabfe95e71576a8641d21689ff","proceedText":"Proceed","message":"Do you want to proceed?","inputs":[],"proceedUrl":"/job/test_pipeline/8/wfapi/inputSubmit?inputId=33d8fefabfe95e71576a8641d21689ff","abortUrl":"/job/test_pipeline/8/input/33d8fefabfe95e71576a8641d21689ff/abort","redirectApprovalUrl":"/job/test_pipeline/8/input/"}]`),
	}

	submitInputResponseKey = fmt.Sprintf(keyFmt, "POST", testSubmitInputURLPath)
	submitInputResponse    = Response{
		StatusCode: 200,
		Status:     "200 OK",
	}
)
