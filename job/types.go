package job

import "encoding/json"

type Info struct {
}

type RunInfo struct {
}

// Parameters is a []Parameter
type Parameters []Parameter

// Parameter is a struct containing parameter key and value pairs
type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type parameterPost struct {
	JSON struct {
		Parameters Parameters `json:"parameter"`
	} `json:"json"`
}

// GetPostBody converts a Parameters struct into []byte to be used for sending requests
// to Jenkins
func (ps Parameters) GetPostBody() []byte {
	reqBodyStruct := parameterPost{}
	for _, p:= range ps {
		reqBodyStruct.JSON.Parameters = append(reqBodyStruct.JSON.Parameters, p)
	}
	reqBodyBytes, _ := json.Marshal(reqBodyStruct)
	return reqBodyBytes
}
