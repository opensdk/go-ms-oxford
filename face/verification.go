// Face - Verification
// Analyzes two faces and determine whether they are from the same person.
// Remarks:
// Verification works well for frontal and near-frontal faces.
// For the scenarios that are sensitive to accuracy please use with own judgment.
package face
import (
	"github.com/opensdk/go-ms-oxford"
	"net/http"
	"strings"
	"fmt"
	"github.com/opensdk/go-ms-oxford/util"
	"encoding/json"
)

type VerificationResult struct {
	oxford.Result
	VerificationResult VerificationResult
}

const VerificationRUL = `https://api.projectoxford.ai/face/v0/verifications`

type FaceVerification struct {
	FaceIdOne              string
	FaceIdTwo              string
	OcpApimSubscriptionKey string //if set, will use this, otherwise will use Config.OcpApimSubscriptionKey
}

func (self FaceVerification) Do() (result VerificationResult, err error) {
	var resp *http.Response
	resp, err = http.Post(VerificationRUL, "application/json", strings.NewReader(fmt.Sprint(`{"faceid1":"%v", "faceid2":"%v"}`, self.FaceIdOne, self.FaceIdTwo)))

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

	var verificationResult VerificationResult
	if err = json.Unmarshal(respBody, &verificationResult); err != nil {
		return
	}

	result.Success = true
	result.Message = "success"
	result.Code = CodeOK
	result.VerificationResult = verificationResult

	return
}