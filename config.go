package oxford

type apiConfig struct {
	imageSizeMax           int
	detectableSizeMin      int
	detectableSizeMax      int
	findsimilarsURL        string //查找相似度url
	detectionsURL          string // 生成 faceId
	formats                []string
	OcpApimSubscriptionKey string
	AnalyzesFaceLandmarks  bool
	AnalyzesHeadPose       bool
	AnalyzesGender         bool
	AnalyzesAge            bool
}

var Config apiConfig

func init() {

	Config.formats = []string{"JPEG", "JPG", "PNG", "GIF", "BMP"}
	Config.detectionsURL = `https://api.projectoxford.ai/face/v0/detections?`
	Config.findsimilarsURL = `https://api.projectoxford.ai/face/v0/findsimilars?`
	Config.detectableSizeMin = 36
	Config.detectableSizeMax = 4096
	Config.imageSizeMax = 4 * 1024 * 1024
}

//ContentType
const ContentTypeJson = "application/json"
const ContentTypeOctetStream = "application/octet-stram"
