package job

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
)

// Info is a struct containing information about a jenkins job
type Info struct {
	Class   string `json:"_class"`
	Actions []struct {
		Class string `json:"_class,omitempty"`
	} `json:"actions"`
	Description       string      `json:"description"`
	DisplayName       string      `json:"displayName"`
	DisplayNameOrNull interface{} `json:"displayNameOrNull"`
	FullDisplayName   string      `json:"fullDisplayName"`
	FullName          string      `json:"fullName"`
	Name              string      `json:"name"`
	URL               string      `json:"url"`
	Buildable         bool        `json:"buildable"`
	Builds            []struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"builds"`
	Color      string `json:"color"`
	FirstBuild struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"firstBuild"`
	HealthReport []struct {
		Description   string `json:"description"`
		IconClassName string `json:"iconClassName"`
		IconURL       string `json:"iconUrl"`
		Score         int    `json:"score"`
	} `json:"healthReport"`
	InQueue          bool `json:"inQueue"`
	KeepDependencies bool `json:"keepDependencies"`
	LastBuild        struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"lastBuild"`
	LastCompletedBuild struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"lastCompletedBuild"`
	LastFailedBuild struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"lastFailedBuild"`
	LastStableBuild struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"lastStableBuild"`
	LastSuccessfulBuild struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"lastSuccessfulBuild"`
	LastUnstableBuild     interface{} `json:"lastUnstableBuild"`
	LastUnsuccessfulBuild struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"lastUnsuccessfulBuild"`
	NextBuildNumber int `json:"nextBuildNumber"`
	Property        []struct {
		Class                string `json:"_class"`
		ParameterDefinitions []struct {
			Class                 string `json:"_class"`
			DefaultParameterValue struct {
				Class string `json:"_class"`
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"defaultParameterValue"`
			Description string `json:"description"`
			Name        string `json:"name"`
			Type        string `json:"type"`
		} `json:"parameterDefinitions,omitempty"`
		Instance struct {
			Class string `json:"_class"`
		} `json:"instance,omitempty"`
	} `json:"property"`
	QueueItem       interface{} `json:"queueItem"`
	ConcurrentBuild bool        `json:"concurrentBuild"`
	ResumeBlocked   bool        `json:"resumeBlocked"`
}

// RunInfo is a struct containing information about a specific jenkins job run
type RunInfo struct {
	Class   string `json:"_class"`
	Actions []struct {
		Class  string `json:"_class,omitempty"`
		Causes []struct {
			Class            string `json:"_class"`
			ShortDescription string `json:"shortDescription"`
			UpstreamBuild    int    `json:"upstreamBuild"`
			UpstreamProject  string `json:"upstreamProject"`
			UpstreamURL      string `json:"upstreamUrl"`
		} `json:"causes,omitempty"`
		Parameters []struct {
			Class string `json:"_class"`
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"parameters,omitempty"`
		BlockedDurationMillis   int     `json:"blockedDurationMillis,omitempty"`
		BlockedTimeMillis       int     `json:"blockedTimeMillis,omitempty"`
		BuildableDurationMillis int     `json:"buildableDurationMillis,omitempty"`
		BuildableTimeMillis     int     `json:"buildableTimeMillis,omitempty"`
		BuildingDurationMillis  int     `json:"buildingDurationMillis,omitempty"`
		ExecutingTimeMillis     int     `json:"executingTimeMillis,omitempty"`
		ExecutorUtilization     float64 `json:"executorUtilization,omitempty"`
		SubTaskCount            int     `json:"subTaskCount,omitempty"`
		WaitingDurationMillis   int     `json:"waitingDurationMillis,omitempty"`
		WaitingTimeMillis       int     `json:"waitingTimeMillis,omitempty"`
		BuildsByBranchName      struct {
			Develop struct {
				Class       string      `json:"_class"`
				BuildNumber int         `json:"buildNumber"`
				BuildResult interface{} `json:"buildResult"`
				Marked      struct {
					SHA1   string `json:"SHA1"`
					Branch []struct {
						SHA1 string `json:"SHA1"`
						Name string `json:"name"`
					} `json:"branch"`
				} `json:"marked"`
				Revision struct {
					SHA1   string `json:"SHA1"`
					Branch []struct {
						SHA1 string `json:"SHA1"`
						Name string `json:"name"`
					} `json:"branch"`
				} `json:"revision"`
			} `json:"develop"`
		} `json:"buildsByBranchName,omitempty"`
		LastBuiltRevision struct {
			SHA1   string `json:"SHA1"`
			Branch []struct {
				SHA1 string `json:"SHA1"`
				Name string `json:"name"`
			} `json:"branch"`
		} `json:"lastBuiltRevision,omitempty"`
		RemoteUrls []string `json:"remoteUrls,omitempty"`
		ScmName    string   `json:"scmName,omitempty"`
	} `json:"actions"`
	Artifacts         []interface{} `json:"artifacts"`
	Building          bool          `json:"building"`
	Description       string        `json:"description"`
	DisplayName       string        `json:"displayName"`
	Duration          int           `json:"duration"`
	EstimatedDuration int           `json:"estimatedDuration"`
	Executor          interface{}   `json:"executor"`
	FullDisplayName   string        `json:"fullDisplayName"`
	ID                string        `json:"id"`
	KeepLog           bool          `json:"keepLog"`
	Number            int           `json:"number"`
	QueueID           int           `json:"queueId"`
	Result            string        `json:"result"`
	Timestamp         int64         `json:"timestamp"`
	URL               string        `json:"url"`
	ChangeSets        []struct {
		Class string `json:"_class"`
		Items []struct {
			Class         string   `json:"_class"`
			AffectedPaths []string `json:"affectedPaths"`
			CommitID      string   `json:"commitId"`
			Timestamp     int64    `json:"timestamp"`
			Author        struct {
				AbsoluteURL string `json:"absoluteUrl"`
				FullName    string `json:"fullName"`
			} `json:"author"`
			AuthorEmail string `json:"authorEmail"`
			Comment     string `json:"comment"`
			Date        string `json:"date"`
			ID          string `json:"id"`
			Msg         string `json:"msg"`
			Paths       []struct {
				EditType string `json:"editType"`
				File     string `json:"file"`
			} `json:"paths"`
		} `json:"items"`
		Kind string `json:"kind"`
	} `json:"changeSets"`
	Culprits []struct {
		AbsoluteURL string `json:"absoluteUrl"`
		FullName    string `json:"fullName"`
	} `json:"culprits"`
	NextBuild struct {
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"nextBuild"`
	PreviousBuild struct {
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"previousBuild"`
}

// Config is a struct used for storing configuration about a jenkins job
type Config struct {
	XMLName          xml.Name `xml:"flow-definition"`
	Text             string   `xml:",chardata"`
	Plugin           string   `xml:"plugin,attr"`
	Description      string   `xml:"description"`
	KeepDependencies string   `xml:"keepDependencies"`
	Properties       struct {
		Text                                  string `xml:",chardata"`
		ComSonyericssonRebuildRebuildSettings struct {
			Text            string `xml:",chardata"`
			Plugin          string `xml:"plugin,attr"`
			AutoRebuild     string `xml:"autoRebuild"`
			RebuildDisabled string `xml:"rebuildDisabled"`
		} `xml:"com.sonyericsson.rebuild.RebuildSettings"`
		HudsonModelParametersDefinitionProperty struct {
			Text                 string `xml:",chardata"`
			ParameterDefinitions struct {
				Text                                 string `xml:",chardata"`
				HudsonModelStringParameterDefinition []struct {
					Text         string `xml:",chardata"`
					Name         string `xml:"name"`
					Description  string `xml:"description"`
					DefaultValue string `xml:"defaultValue"`
					Trim         string `xml:"trim"`
				} `xml:"hudson.model.StringParameterDefinition"`
			} `xml:"parameterDefinitions"`
		} `xml:"hudson.model.ParametersDefinitionProperty"`
		EnvInjectJobProperty struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
			Info   struct {
				Text               string `xml:",chardata"`
				PropertiesContent  string `xml:"propertiesContent"`
				SecureGroovyScript struct {
					Text    string `xml:",chardata"`
					Plugin  string `xml:"plugin,attr"`
					Script  string `xml:"script"`
					Sandbox string `xml:"sandbox"`
				} `xml:"secureGroovyScript"`
				LoadFilesFromMaster string `xml:"loadFilesFromMaster"`
			} `xml:"info"`
			On                         string `xml:"on"`
			KeepJenkinsSystemVariables string `xml:"keepJenkinsSystemVariables"`
			KeepBuildVariables         string `xml:"keepBuildVariables"`
			OverrideBuildParameters    string `xml:"overrideBuildParameters"`
		} `xml:"EnvInjectJobProperty"`
		ComCloudbeesHudsonPluginsModelingImplJobTemplateJobPropertyImpl struct {
			Text     string `xml:",chardata"`
			Plugin   string `xml:"plugin,attr"`
			Instance struct {
				Text   string `xml:",chardata"`
				Model  string `xml:"model"`
				Values struct {
					Text  string `xml:",chardata"`
					Class string `xml:"class,attr"`
					Entry []struct {
						Text                                                           string   `xml:",chardata"`
						String                                                         []string `xml:"string"`
						ComCloudbeesHudsonPluginsModelingImplAuxiliaryAuxInstanceArray string   `xml:"com.cloudbees.hudson.plugins.modeling.impl.auxiliary.AuxInstance-array"`
						ComCloudbeesHudsonPluginsModelingImplAuxiliaryAuxInstance      struct {
							Text   string `xml:",chardata"`
							Model  string `xml:"model"`
							Values struct {
								Text  string `xml:",chardata"`
								Class string `xml:"class,attr"`
								Entry []struct {
									Text    string   `xml:",chardata"`
									String  []string `xml:"string"`
									Boolean string   `xml:"boolean"`
								} `xml:"entry"`
							} `xml:"values"`
						} `xml:"com.cloudbees.hudson.plugins.modeling.impl.auxiliary.AuxInstance"`
						Boolean string `xml:"boolean"`
					} `xml:"entry"`
				} `xml:"values"`
			} `xml:"instance"`
		} `xml:"com.cloudbees.hudson.plugins.modeling.impl.jobTemplate.JobPropertyImpl"`
	} `xml:"properties"`
	Definition struct {
		Text   string `xml:",chardata"`
		Class  string `xml:"class,attr"`
		Plugin string `xml:"plugin,attr"`
		Scm    struct {
			Text              string `xml:",chardata"`
			Class             string `xml:"class,attr"`
			Plugin            string `xml:"plugin,attr"`
			ConfigVersion     string `xml:"configVersion"`
			UserRemoteConfigs struct {
				Text                             string `xml:",chardata"`
				HudsonPluginsGitUserRemoteConfig struct {
					Text          string `xml:",chardata"`
					URL           string `xml:"url"`
					CredentialsID string `xml:"credentialsId"`
				} `xml:"hudson.plugins.git.UserRemoteConfig"`
			} `xml:"userRemoteConfigs"`
			Branches struct {
				Text                       string `xml:",chardata"`
				HudsonPluginsGitBranchSpec struct {
					Text string `xml:",chardata"`
					Name string `xml:"name"`
				} `xml:"hudson.plugins.git.BranchSpec"`
			} `xml:"branches"`
			DoGenerateSubmoduleConfigurations string `xml:"doGenerateSubmoduleConfigurations"`
			SubmoduleCfg                      struct {
				Text  string `xml:",chardata"`
				Class string `xml:"class,attr"`
			} `xml:"submoduleCfg"`
			Extensions string `xml:"extensions"`
		} `xml:"scm"`
		ScriptPath  string `xml:"scriptPath"`
		Lightweight string `xml:"lightweight"`
	} `xml:"definition"`
	Triggers string `xml:"triggers"`
	Disabled string `xml:"disabled"`
}

// Parameters is a []Parameter
type Parameters []Parameter

// Parameter is a struct containing parameter key and value pairs
type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// GetPostBody converts a Parameters struct into []byte to be used for sending requests
// to Jenkins
func (ps Parameters) GetPostBody() io.Reader {
	reqBodyStruct := struct {
		JSON struct {
			Parameters Parameters `json:"parameter"`
		} `json:"json"`
	}{}

	for _, p := range ps {
		reqBodyStruct.JSON.Parameters = append(reqBodyStruct.JSON.Parameters, p)
	}
	reqBodyBytes, _ := json.Marshal(reqBodyStruct)
	return bytes.NewReader(reqBodyBytes)
}
