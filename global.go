package oxford

type apiConfig struct {
	ImageSizeMax           int
	DetectableSizeMin      int
	DetectableSizeMax      int
	FindsimilarsURL        string //查找相似度url
	DetectionsURL          string // 生成 faceId
	Formats                []string
	OcpApimSubscriptionKey string
	AnalyzesFaceLandmarks  bool
	AnalyzesHeadPose       bool
	AnalyzesGender         bool
	AnalyzesAge            bool
}

var Config apiConfig

func init() {
	Config.Formats = []string{"JPEG", "JPG", "PNG", "GIF", "BMP"}
	Config.DetectionsURL = `https://api.projectoxford.ai/face/v0/detections?`
	Config.FindsimilarsURL = `https://api.projectoxford.ai/face/v0/findsimilars?`
	Config.DetectableSizeMin = 36
	Config.DetectableSizeMax = 4096
	Config.ImageSizeMax = 4 * 1024 * 1024
}

//ContentType
const ContentTypeJson = "application/json"
const ContentTypeOctetStream = "application/octet-stream"

type Result struct {
	StatusCode int
	Success    bool
	Code       string
	Message    string
}