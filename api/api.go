package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type params struct {
	Command string            `json:"command"`
	Url     string            `json:"url"`
	Header  map[string]string `json:"header"`
	Body    string            `json:"body"`
	Params  map[string]string `json:"url_params"`
}

func Process(rawParams map[string]string) {

	pData, err := processRawParams(rawParams)
	if err != nil {
		fmt.Print(err)
		return
	}

	if pData.Command == "POST" {

		request := buildRequest(*pData)
		response, _ := HttpRequest(request)

		fmt.Printf("\n **Success!!** \n response : %v", string(response.Data))
		return
	} else {
		fmt.Printf("\n Command Not Found [%v] Error : %d", pData.Command, 404)
		return
	}
}

func processRawParams(rawParams map[string]string) (*params, error) {

	if param_data, ok := rawParams["raw_params"]; ok {

		var pData = params{}

		err := json.Unmarshal([]byte(param_data), &pData)
		if err != nil {
			fmt.Printf(" failed to unmarshal raw_params: '%s' \n  error: [%v]", param_data, err)
			return nil, err
		}
		return &pData, err
	}
	return nil, nil
}

func buildRequest(p params) APIRequest {
	var req APIRequest

	req.Url = p.Url
	req.Body = bytes.NewBuffer([]byte(p.Body))

	if len(p.Params) > 0 {
		postFormData := url.Values{}
		for key, value := range p.Params {
			postFormData.Add(key, value)
		}
		req.PostForm = postFormData.Encode()
	}

	if len(p.Header) > 0 {
		req.Headers = p.Header
	} else {
		req.Headers = map[string]string{}
		req.Headers["Content-Type"] = "application/json"
	}

	req.Method = http.MethodPost

	return req
}
