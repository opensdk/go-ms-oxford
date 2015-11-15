package oxford
import "image"

// 结果
type FaceResult struct {
	FaceId        string
	FaceRectangle FaceRectangle
	FaceLandmarks FaceLandmarks
	Attributes    Attributes
}

type FaceLandmarks struct {
	PupilLeft  image.Point
	PupilRight image.Point
	NoseTip    image.Point
}

type FaceRectangle struct {
	Top    string
	Left   string
	Width  string
	Height string
}

type Attributes struct {
	HeadPose HeadPose
	Gender   string
	Age      int
}

type HeadPose struct {
	Pitch int
	Roll  float64
	Yaw   float64
}