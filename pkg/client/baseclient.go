package client

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	// DefaultScheme is default http scheme
	DefaultScheme = "http"
	// DefaultRegion is default region
	DefaultRegion = "cn-north-1"
	// DefaultSdkVersion will be used in business framework.
	DefaultSdkVersion = "2021-03-03"
)

// Client defines base client struct
type Client struct {
	Client      http.Client
	SdkVersion  string
	ServiceInfo *base.ServiceInfo
	headers     http.Header
}

// NewServiceInfo return base ServiceInfo.
func NewServiceInfo() *base.ServiceInfo {
	return &base.ServiceInfo{
		Timeout:     120 * time.Second, // TOP gateway timeout
		Scheme:      DefaultScheme,
		Credentials: base.Credentials{Service: "iam", Region: DefaultRegion},
	}
}

// NewBaseClient return a base Client
func NewBaseClient() *Client {
	transport := &http.Transport{
		MaxIdleConns:        1000,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     10 * time.Second,
	}

	c := &Client{
		Client:  http.Client{Transport: transport},
		headers: http.Header{},
	}

	header := http.Header{
		"Content-Type": []string{"application/json"},
		"Accept":       []string{"application/json"},
	}
	c.headers = header

	return c
}

// AddExtraHeaders add custom headers into request
func (client *Client) AddExtraHeaders(h map[string]string) {
	if client.headers == nil {
		client.headers = http.Header{}
	}

	// add extra headers
	for k, v := range h {
		client.headers.Add(k, v)
	}
}

// CommonHandler handle http request and response.
func (client *Client) CommonHandler(action string, query url.Values, body string, resp interface{}) (int, error) {
	respBody, statusCode, err := client.JSON(action, query, body)
	if err != nil {
		return statusCode, errors.WithMessage(err, "client send json")
	}
	data := make(map[string]interface{})
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return statusCode, errors.WithMessage(err, "json unmarshal http response body")
	}
	result := data["Result"]
	res, err := json.Marshal(result)
	if err != nil {
		return statusCode, errors.WithMessage(err, "json marshal")
	}
	if err := json.Unmarshal(res, resp); err != nil {
		return statusCode, errors.WithMessage(err, "json unmarshal response struct")
	}
	return statusCode, nil
}

// JSON send a post json request
func (client *Client) JSON(action string, query url.Values, body string) ([]byte, int, error) {
	timeout := client.ServiceInfo.Timeout

	if query == nil {
		query = url.Values{}
	}
	query.Add("Action", action)
	query.Add("Version", client.SdkVersion)

	u := url.URL{
		Scheme:   client.ServiceInfo.Scheme,
		Host:     client.ServiceInfo.Host,
		Path:     "/",
		RawQuery: query.Encode(),
	}

	req, err := http.NewRequest(strings.ToUpper(http.MethodPost), u.String(), strings.NewReader(body))
	if err != nil {
		return nil, 0, errors.WithMessage(err, "build request")
	}
	req.Header = client.headers.Clone()
	return client.makeRequest(action, req, body, timeout)
}

func (client *Client) makeRequest(action string, req *http.Request, reqBody string, timeout time.Duration) ([]byte, int, error) {
	req = client.ServiceInfo.Credentials.Sign(req)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req = req.WithContext(ctx)

	logrus.Infof("request url is %v\n header is %v\n req body is %v\n", req.URL.String(), req.Header, reqBody)

	resp, err := client.Client.Do(req)
	if err != nil {
		return nil, 0, errors.WithMessage(err, "do request")
	}
	defer func() {
		_ = resp.Body.Close() // we don't care if close successfully.
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, errors.WithMessage(err, "read response err")
	}

	logrus.Infof("raw response is %v\n response code is %v\n", string(body), resp.Status)

	return body, resp.StatusCode, getErrorFromBody(body)
}
