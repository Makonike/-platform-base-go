package aosdk

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const (
	RequestIdHeader       = "Request-Id"
	BoxRegKeyHeader       = "Box-Reg-Key"
	ApplicationJsonHeader = "application/json"
)

var (
	// client is a shared http Client.
	client HttpClient = &http.Client{}
)

// SetHttpClient sets custom http Client.
func SetHttpClient(httpClient HttpClient) {
	client = httpClient
}

// HttpClient interface has the method required to use a type as custom http client.
// The net/*http.Client type satisfies this interface.
type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

// ErrorResponse only used when an interface call failed.
type ErrorResponse struct {
	RequestId string `json:"requestId"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

func DoPost(action, requestId, regBoxKey string, queryMap map[string]string, mBody interface{}) ([]byte, error) {
	url := GetUrl(action, queryMap)
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(mBody)
	if err != nil {
		return nil, err
	}

	req, err := NewRequest(http.MethodPost, url, requestId, regBoxKey, body)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}

func NewRequest(method, action, requestId, boxRegKey string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, action, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", ApplicationJsonHeader)
	if requestId != "" {
		req.Header.Set(RequestIdHeader, requestId)
	}
	if boxRegKey != "" {
		req.Header.Set(BoxRegKeyHeader, boxRegKey)
	}

	return req, nil
}
