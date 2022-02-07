package client

import (
	"net/http"
	"net/url"
)

// NewFakeClient return a fake client
func NewFakeClient() Client {
	return &fakeClient{}
}

type fakeClient struct{}

// CommonHandler handle fake request and response.
func (client *fakeClient) CommonHandler(action string, query url.Values, body string, resp interface{}) (int, error) {
	return int(http.StatusOK), nil
}
