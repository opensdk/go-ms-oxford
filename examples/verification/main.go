package main

import (
	"github.com/opensdk/go-ms-oxford/face"
	"fmt"
	"os"
	"github.com/opensdk/go-ms-oxford"
)

func main() {
	oxford.Config.AnalyzesAge = true
	oxford.Config.AnalyzesFaceLandmarks = true
	oxford.Config.AnalyzesGender = true
	oxford.Config.AnalyzesHeadPose = true
	oxford.Config.OcpApimSubscriptionKey = "66672111946e4cc09abdfdaf66846dfc"

	v := face.FaceVerification{FaceIdOne: "177c18d5-a95b-4293-8e63-39a1b7dd0f5b", FaceIdTwo: "177c18d5-a95b-4293-8e63-39a1b7dd0f5b"}
	result, err := v.Do()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result.Success, result.StatusCode, result.Code, result.Message, result.IsIdentical, result.Confidence)
}
