// Face - Verification
// Analyzes two faces and determine whether they are from the same person.
// Remarks:
// Verification works well for frontal and near-frontal faces.
// For the scenarios that are sensitive to accuracy please use with own judgment.
package face

import (
	"encoding/json"
	"fmt"
	"github.com/opensdk/go-ms-oxford"
	"github.com/opensdk/go-ms-oxford/util"
	"net/http"
	"strings"
)

type VerificationResult struct {
	oxford.Result
	IsIdentical bool
	Confidence  float64
}

const VerificationRUL = `https://api.projectoxford.ai/face/v0/verifications`

type FaceVerification struct {
	FaceIdOne              string
	FaceIdTwo              string
	OcpApimSubscriptionKey string //if set, will use this, otherwise will use Config.OcpApimSubscriptionKey
}

func (self FaceVerification) Do() (result VerificationResult, err error) {
	apiKey := oxford.Config.OcpApimSubscriptionKey
	if len(self.OcpApimSubscriptionKey) > 0 {
		apiKey = self.OcpApimSubscriptionKey
	}

	req, err := http.NewRequest("POST", VerificationRUL, strings.NewReader(fmt.Sprintf(`{"faceId1":"%v", "faceId2":"%v"}`, self.FaceIdOne, self.FaceIdTwo)))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", oxford.ContentTypeJson)
	req.Header.Set("Ocp-Apim-Subscription-Key", apiKey)

	var resp *http.Response
	client := &http.Client{}
	resp, err = client.Do(req)

	respBody, err := util.HandleResponse(resp)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var bodyMap map[string]interface{}
		if err = json.Unmarshal(respBody, &bodyMap); err != nil {
			return
		}

		result.Message = util.Convert2String(bodyMap["message"])
		result.Code = util.Convert2String(bodyMap["message"])
		return
	}

	if err = json.Unmarshal(respBody, &result); err != nil {
		return
	}

	result.Success = true
	result.Message = "success"
	result.Code = CodeOK

	return
}
