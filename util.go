package oxford

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"fmt"
)

func postFile(url string, reader io.Reader, apiKey string) (resp *http.Response, err error) {
	// TODO not worked
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", "1.jpg")
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", getURL(url), &body)
	if err != nil {
		return nil, err
	}

	setHeader(req, apiKey)
	client := &http.Client{}
	return client.Do(req)
}

func postURL(url, imageURL string, apiKey string) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", getURL(url), strings.NewReader(fmt.Sprintf(`{"url":"%v"}`, imageURL)))
	if err != nil {
		return nil, err
	}

	setHeader(req, apiKey)

	client := &http.Client{}
	return client.Do(req)
}

func setHeader(request *http.Request, apiKey string) {
	request.Header.Set("Content-Type", ContentTypeJson)
	request.Header.Set("Ocp-Apim-Subscription-Key", apiKey)
}

func convert2String(obj interface{}) string {
	if obj == nil {
		return ""
	}

	if str, ok := obj.(string); ok {
		return str
	}

	return ""
}

type point struct {
	X float64
	Y float64
}

func getURL(url string) string {
	if Config.AnalyzesAge {
		url += "analyzesAge=true"
	}

	if Config.AnalyzesFaceLandmarks {
		url += "&analyzesFaceLandmarks=true"
	}

	if Config.AnalyzesGender {
		url += "&analyzesGender=true"
	}

	if Config.AnalyzesHeadPose {
		url += "&analyzesHeadPose=true"
	}

	return url
}