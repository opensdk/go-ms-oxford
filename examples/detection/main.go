package main
import (
	"github.com/opensdk/go-ms-oxford"
	"fmt"
)

func main() {
	oxford.Config.AnalyzesAge = true
	oxford.Config.AnalyzesFaceLandmarks = true
	oxford.Config.AnalyzesGender = true
	oxford.Config.AnalyzesHeadPose = true
	oxford.Config.OcpApimSubscriptionKey = ""




	detection := &oxford.FaceDetection{}
	fmt.Println(detection.Detect())
}
