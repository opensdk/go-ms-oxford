package config

type apiConfig struct {
	imageSizeMax      int
	detectableSizeMin int
	detectableSizeMax int
	findsimilarsUrl   string //查找相似度url
	detectionsUrl     string // 生成 faceId
	formats           []string
	Key               string
	AnalyzesFaceLandmarks         bool
	AnalyzesHeadPose         bool
	AnalyzesGender            bool
	AnalyzesAge               bool
}


var Config apiConfig

func init() {
	Config.formats = []string{"JPEG", "JPG", "PNG", "GIF", "BMP"}
	Config.findsimilarsUrl = `https://api.projectoxford.ai/face/v0/findsimilars`
	Config.detectableSizeMin = 36
	Config.detectableSizeMax = 4096
	Config.imageSizeMax = 4 * 1024 * 1024
}