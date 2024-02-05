package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

// --------------------------------------------------------------------------
// Abstraction for HTTP request
type APIRequest struct {
	Method   string
	Url      string
	Headers  map[string]string
	Body     io.Reader
	PostForm string
}

// --------------------------------------------------------------------------
// Abstraction for HTTP response
type APIResponse struct {
	Data []byte
	URL  string
}

func HttpRequest(request APIRequest) (*APIResponse, error) {
	var err error
	var req *http.Request
	response := new(APIResponse)

	if len(request.PostForm) > 0 {
		request.Url += "?" + request.PostForm
	}

	if request.Method == http.MethodPost {
		req, err = http.NewRequest(http.MethodPost, request.Url, request.Body)
		if err != nil {
			fmt.Printf("Failed to build request '%s': [%v]", request.Url, err)
			return response, err
		}
	}
	if req != nil {
		if len(request.Headers) > 0 {
			for hkey, hval := range request.Headers {
				req.Header.Set(hkey, hval)
			}
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("Http Request Failed '%s': [%v]", request.Url, err)
			return response, err
		}
		defer resp.Body.Close()

		response.URL = request.Url
		if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
			msg := fmt.Sprintf("Got invalid response status code: '%d' for request: '%s'", resp.StatusCode, request.Url)
			return response, errors.New(msg)
		}
		response.Data, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Failed to read response '%s': [%v]", response.URL, err)
			return response, err
		}
		return response, err
	}
	return response, err
}
