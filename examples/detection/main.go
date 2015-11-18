package main

import (
	"fmt"
	"github.com/opensdk/go-ms-oxford"
	"github.com/opensdk/go-ms-oxford/face"
	"os"
)

func main() {
	oxford.Config.AnalyzesAge = true
	oxford.Config.AnalyzesFaceLandmarks = true
	oxford.Config.AnalyzesGender = true
	oxford.Config.AnalyzesHeadPose = true
	oxford.Config.OcpApimSubscriptionKey = "66672111946e4cc09abdfdaf66846dfc"

	testURL()
	testImage()
}

func testURL() {
	detection := &face.FaceDetection{URL: `http://i3.letvimg.com/lc02_iscms/201511/15/12/59/b378f6b5910c4904af7521a5e35432c8.jpg`}
	result, err := detection.Do()
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Success)
	if result.Success {
		for _, v := range result.FaceResults {
			fmt.Println("faceId:", v.FaceId, "age:", v.Attributes.Age, "gender:", v.Attributes.Gender, "NoseRightAlarOutTip:", v.FaceLandmarks.NoseRightAlarOutTip)
			fmt.Println(v.FaceLandmarks.PupilLeft)
		}
	}
}

func testImage() {
	{
		f, err := os.Open("../images/1.jpg")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()

		detection := &face.FaceDetection{Reader: f}
		result, err := detection.Do()
		if err != nil {
			panic(err)
		}

		fmt.Println(result.Success)
		if result.Success {
			for _, v := range result.FaceResults {
				fmt.Println("faceIdL", v.FaceId, "age:", v.Attributes.Age, "gender:", v.Attributes.Gender, "NoseRightAlarOutTip:", v.FaceLandmarks.NoseRightAlarOutTip)
				fmt.Println(v.FaceLandmarks.PupilLeft)
			}
		} else {
			fmt.Println(result.Message)
		}
	}
}
