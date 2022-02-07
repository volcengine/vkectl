package client

import "net/url"

type Client interface {
	CommonHandler(action string, query url.Values, body string, resp interface{}) (int, error)
}
