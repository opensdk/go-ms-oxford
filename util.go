package oxford
import (
	"net/http"
	"io"
	"fmt"
	"mime/multipart"
	"bytes"
	"os"
)

func post(url string, bodyType string, reader io.Reader) (resp *http.Response, err error) {
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


	fmt.Println(url)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", "/Users/du/Desktop/1.jpg")
	if err != nil {
		return nil, err
	}

	file, err := os.Open("/Users/du/Desktop/1.jpg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", url, &body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", bodyType)
	req.Header.Set("Ocp-Apim-Subscription-Key", Config.OcpApimSubscriptionKey)
	client := &http.Client{}
	return client.Do(req)
}