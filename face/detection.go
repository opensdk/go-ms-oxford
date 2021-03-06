// Face - Detection
// Detects human faces in an image and returns face locations, face landmarks,
// and optional attributes including head-pose, gender, and age.
// Detection is an essential API that provides faceId to other APIs like Identification,
// Verification, and Find Similar.

// The supported input image formats includes JPEG, PNG, GIF(the first frame), BMP, and image file size should be no larger than 4MB.
// Input parameter face landmarks, head-poses, gender, and age are optional. You need to specify them in the URL to turn on the functions. Otherwise, only face rectangles will be returned.
// The detectable face size range is 36x36 to 4096x4096 pixels. The faces out of this range will not be detected.
// For each image, the maximum number of faces returned is 64 and the faces are ranked by face rectangle size in descending order.
// Some faces may not be detected for technical challenges, e.g. very large face angles (head-pose), large occlusion. Frontal and near-frontal faces have the best results.
// The attribute headPose's pitch value is reserved as 0.0
// The attributes gender and age are still experimental and may not be very accurate for now.
// The face ID will expire 24 hours after detection.
package face

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opensdk/go-ms-oxford"
	"github.com/opensdk/go-ms-oxford/util"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

// detect result
type DetectResult struct {
	oxford.Result
	FaceResults []Face
}

// face detect result
type Face struct {
	FaceId        string
	FaceRectangle FaceRectangle
	FaceLandmarks FaceLandmarks
	Attributes    Attributes
}

type FaceLandmarks struct {
	PupilLeft           util.Point
	PupilRight          util.Point
	NoseTip             util.Point
	MouthLeft           util.Point
	MouthRight          util.Point
	EyebrowLeftOuter    util.Point
	EyeLeftOuter        util.Point
	EyeLeftTop          util.Point
	EyeLeftBottom       util.Point
	EyeLeftInner        util.Point
	EyebrowRightInner   util.Point
	EyebrowRightOuter   util.Point
	EyeRightInner       util.Point
	EyeRightTop         util.Point
	EyeRightBottom      util.Point
	EyeRightOuter       util.Point
	NoseRootLeft        util.Point
	NoseRootRight       util.Point
	NoseLeftAlarTop     util.Point
	NoseRightAlarTop    util.Point
	NoseLeftAlarOutTip  util.Point
	NoseRightAlarOutTip util.Point
	UpperLipTop         util.Point
	UpperLipBottom      util.Point
	UnderLipTop         util.Point
	UnderLipBottom      util.Point
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
	CodeOK               = "Ok"               // success
	CodeBadArgument      = "BadArgument"      // JSON parsing error.
	CodeInvalidURL       = "InvalidURL"       // Invalid image URL.
	CodeInvalidImage     = "InvalidImage"     // Decoding error or unsupported image format.
	CodeInvalidImageSize = "InvalidImageSize" //Image size is too small (smaller than a detectable face size of 36x36 pixels) or too big (larger than 4MB file limit).
)

type FaceDetection struct {
	URL                    string //if set, will use url, otherwise will use Reader
	Reader                 io.Reader
	OcpApimSubscriptionKey string //if set, will use this, otherwise will use Config.OcpApimSubscriptionKey
}

// start to detect a photo
func (self FaceDetection) Do() (detectResult DetectResult, err error) {
	apiKey := oxford.Config.OcpApimSubscriptionKey
	if len(self.OcpApimSubscriptionKey) > 0 {
		apiKey = self.OcpApimSubscriptionKey
	}
	var resp *http.Response
	if len(self.URL) == 0 {
		if nil == self.Reader {
			err = errors.New("no content or url to detect")
			return
		}

		resp, err = postFile(oxford.Config.DetectionsURL, self.Reader, apiKey)

	} else {
		resp, err = postURL(oxford.Config.DetectionsURL, self.URL, apiKey)
	}

	if err != nil {
		return
	}

	respBody, err := util.HandleResponse(resp)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	detectResult.StatusCode = resp.StatusCode

	if resp.StatusCode != 200 {
		var bodyMap map[string]interface{}
		if err = json.Unmarshal(respBody, &bodyMap); err != nil {
			return
		}

		detectResult.Message = util.Convert2String(bodyMap["message"])
		detectResult.Code = util.Convert2String(bodyMap["message"])
		return
	}

	var faceResult []Face
	if err = json.Unmarshal(respBody, &faceResult); err != nil {
		return
	}

	detectResult.Success = true
	detectResult.Message = "success"
	detectResult.Code = CodeOK
	detectResult.FaceResults = faceResult

	return
}

func postURL(url, imageURL string, apiKey string) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", getURL(url), strings.NewReader(fmt.Sprintf(`{"url":"%v"}`, imageURL)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", oxford.ContentTypeJson)
	req.Header.Set("Ocp-Apim-Subscription-Key", apiKey)

	client := &http.Client{}
	return client.Do(req)
}

func postFile(url string, reader io.Reader, apiKey string) (resp *http.Response, err error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", "111.png")
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, reader)
	if err != nil {
		return
	}

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
