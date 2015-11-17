package util

import (
	"bytes"
	"fmt"
	"github.com/opensdk/go-ms-oxford"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
	"io/ioutil"
)

func PostFile(url string, reader io.Reader, apiKey string) (resp *http.Response, err error) {
	// TODO not worked
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", "111.png")
	if err != nil {
		return nil, err
	}

	var size int64
	size, err = io.Copy(part, reader)
	if err != nil {
		return
	}

	fmt.Println("size:", size)

	if err = writer.Close(); err != nil {
		return
	}

	req, err := http.NewRequest("POST", getURL(url), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", oxford.ContentTypeOctetStream)
	req.Header.Set("Ocp-Apim-Subscription-Key", apiKey)
	client := &http.Client{Timeout: time.Second * 30}
	return client.Do(req)
}

func PostURL(url, imageURL string, apiKey string) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", getURL(url), strings.NewReader(fmt.Sprintf(`{"url":"%v"}`, imageURL)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", oxford.ContentTypeJson)
	req.Header.Set("Ocp-Apim-Subscription-Key", apiKey)

	client := &http.Client{}
	return client.Do(req)
}


func HandleResponse(response *http.Response) (body []byte, err error)  {
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	return
}

func Convert2String(obj interface{}) string {
	if obj == nil {
		return ""
	}

	if str, ok := obj.(string); ok {
		return str
	}

	return ""
}

type Point struct {
	X float64
	Y float64
}

func getURL(url string) string {
	if oxford.Config.AnalyzesAge {
		url += "analyzesAge=true"
	}

	if oxford.Config.AnalyzesFaceLandmarks {
		url += "&analyzesFaceLandmarks=true"
	}

	if oxford.Config.AnalyzesGender {
		url += "&analyzesGender=true"
	}

	if oxford.Config.AnalyzesHeadPose {
		url += "&analyzesHeadPose=true"
	}

	return url
}
