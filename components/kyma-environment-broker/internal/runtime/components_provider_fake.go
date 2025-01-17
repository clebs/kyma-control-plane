package runtime

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

const kymaComponentsYAMLFileName = "kyma-components.yaml"

type fakeHTTPDoer struct{}

func (f *fakeHTTPDoer) Do(req *http.Request) (*http.Response, error) {
	urlPathSplit := strings.Split(req.URL.Path, "/")
	if !isSupportedVersion(urlPathSplit[2]) {
		return &http.Response{
			Status:     "404 Not Found",
			StatusCode: 404,
			Body:       http.NoBody,
			Request:    req,
		}, nil
	}

	yamlFilePath := path.Join("testdata", kymaComponentsYAMLFileName)
	contents, err := os.ReadFile(yamlFilePath)
	if err != nil {
		return &http.Response{
			Status:     "500 Internal Server Error",
			StatusCode: 500,
			Body:       http.NoBody,
			Request:    req,
		}, nil
	}

	body := io.NopCloser(bytes.NewReader(contents))

	return &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       body,
		Request:    req,
	}, nil
}

func isSupportedVersion(version string) bool {
	return strings.HasPrefix(version, "PR-") ||
		strings.HasPrefix(version, "main-") ||
		strings.Split(version, ".")[0] == "2"
}

func NewFakeComponentsProvider() *ComponentsProvider {
	return &ComponentsProvider{
		httpClient: &fakeHTTPDoer{},
	}
}
