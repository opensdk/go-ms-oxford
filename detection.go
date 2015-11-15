package oxford

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// detect result
type DetectResult struct {
	ResponseCode int
	Success      bool
	Code         string
	Message      string
	FaceResults  []FaceResult
}

// face detect result
type FaceResult struct {
	FaceId        string
	FaceRectangle FaceRectangle
	FaceLandmarks FaceLandmarks
	Attributes    Attributes
}

type FaceLandmarks struct {
	PupilLeft           point
	PupilRight          point
	NoseTip             point
	MouthLeft           point
	MouthRight          point
	EyebrowLeftOuter    point
	EyeLeftOuter        point
	EyeLeftTop          point
	EyeLeftBottom       point
	EyeLeftInner        point
	EyebrowRightInner   point
	EyebrowRightOuter   point
	EyeRightInner       point
	EyeRightTop         point
	EyeRightBottom      point
	EyeRightOuter       point
	NoseRootLeft        point
	NoseRootRight       point
	NoseLeftAlarTop     point
	NoseRightAlarTop    point
	NoseLeftAlarOutTip  point
	NoseRightAlarOutTip point
	UpperLipTop         point
	UpperLipBottom      point
	UnderLipTop         point
	UnderLipBottom      point
}

type FaceRectangle struct {
	Top    int
	Left   int
	Width  int
	Height int
}

type Attributes struct {
	HeadPose HeadPose
	Gender   string
	Age      float64
}

type HeadPose struct {
	Pitch float64
	Roll  float64
	Yaw   float64
}

const (
	CodeOK = "Ok"// success
	CodeBadArgument = "BadArgument"      // JSON parsing error.
	CodeInvalidURL = "InvalidURL"       // Invalid image URL.
	CodeInvalidImage = "InvalidImage"     // Decoding error or unsupported image format.
	CodeInvalidImageSize = "InvalidImageSize" //Image size is too small (smaller than a detectable face size of 36x36 pixels) or too big (larger than 4MB file limit).
)

type FaceDetection struct {
	ContentType            string
	URL                    string //if set, will use url, otherwise will use Reader
	Reader                 io.Reader
	OcpApimSubscriptionKey string //if set, will use this, otherwise will use Config.OcpApimSubscriptionKey
}

// start to detect a photo
func (self FaceDetection) Detect() (detectResult DetectResult, err error) {
	apiKey := Config.OcpApimSubscriptionKey
	if len(self.OcpApimSubscriptionKey) > 0 {
		apiKey = self.OcpApimSubscriptionKey
	}
	var resp *http.Response
	if len(self.URL) == 0 {
		err = errors.New("reader not support yet, please use url")
		return

		if nil == self.Reader {
			err = errors.New("no content or url to detect")
			return
		}

		self.ContentType = ContentTypeOctetStream
		resp, err = postFile(Config.detectionsURL, self.Reader, apiKey)

	} else {
		self.ContentType = ContentTypeJson
		resp, err = postURL(Config.detectionsURL, self.URL, apiKey)
	}

	defer resp.Body.Close()
	var respBody []byte
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	detectResult.ResponseCode = resp.StatusCode

	if resp.StatusCode != 200 {
		var bodyMap map[string]interface{}
		if err = json.Unmarshal(respBody, &bodyMap); err != nil {
			return
		}

		detectResult.Message = convert2String(bodyMap["message"])
		detectResult.Code = convert2String(bodyMap["message"])
		return
	}

	var faceResult []FaceResult
	if err = json.Unmarshal(respBody, &faceResult); err != nil {
		return
	}

	detectResult.Success = true
	detectResult.Message = "success"
	detectResult.Code = CodeOK
	detectResult.FaceResults = faceResult

	return
}
