package client

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// ErrorData is the message in response body
type ErrorData struct {
	Code    string            `json:"Code"`
	Message string            `json:"Message"`
	Data    map[string]string `json:"Data"`
}

func (data *ErrorData) Error() string {
	str, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return string(str)
}

func getErrorFromBody(body []byte) error {
	res := make(map[string]interface{})
	err := json.Unmarshal(body, &res)
	if err != nil {
		return errors.WithMessage(err, "json marshal error body")
	}

	meta, ok := res["ResponseMetadata"].(map[string]interface{})
	if !ok {
		return errors.New("response meta data not exist")
	}

	errMeta, ok := meta["Error"]
	if !ok {
		return nil
	}

	errMetadata, err := json.Marshal(errMeta)
	if err != nil {
		return errors.WithMessage(err, "json marshal error meta data")
	}

	newError := &ErrorData{}
	if err := json.Unmarshal(errMetadata, newError); err != nil {
		return errors.WithMessage(err, "json unmarshal error meta data")
	}
	return newError
}
