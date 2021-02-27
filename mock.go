package jenkins

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type MockRequester struct {
	DataFilePath string
}

func (mr MockRequester) Request(method, url string, body io.Reader) (*http.Response, error) {
	dataFile, err := os.Open(mr.DataFilePath)
	if err != nil {
		err = fmt.Errorf("Failed to open file %s with the error: %s", mr.DataFilePath, err.Error())
		panic(err.Error())
	}
	response := &http.Response{
		Body:       dataFile,
		StatusCode: 200,
	}
	return response, nil
}
