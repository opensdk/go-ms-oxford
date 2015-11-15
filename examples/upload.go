package main
import (
	"github.com/opensdk/go-ms-oxford"
	"encoding/json"
	"fmt"
)

func main() {
	str := `[{"faceId":"a64fa436-3845-4542-b4cd-f2fa289f3fb9","faceRectangle":{"width":64,"height":64,"left":173,"top":38},"faceLandmarks":{"pupilLeft":{"x":188.8,"y":59.5},"pupilRight":{"x":213.5,"y":51.5},"noseTip":{"x":202.6,"y":71.6},"mouthLeft":{"x":195.2,"y":86.9},"mouthRight":{"x":221.7,"y":81.4},"eyebrowLeftOuter":{"x":176.8,"y":58.1},"eyebrowLeftInner":{"x":192.2,"y":49.9},"eyeLeftOuter":{"x":184.4,"y":61.4},"eyeLeftTop":{"x":187.3,"y":58.1},"eyeLeftBottom":{"x":188.5,"y":61.4},"eyeLeftInner":{"x":192.4,"y":59.1},"eyebrowRightInner":{"x":202.0,"y":46.6},"eyebrowRightOuter":{"x":221.8,"y":41.9},"eyeRightInner":{"x":209.8,"y":53.7},"eyeRightTop":{"x":212.8,"y":49.8},"eyeRightBottom":{"x":214.1,"y":53.1},"eyeRightOuter":{"x":217.9,"y":50.5},"noseRootLeft":{"x":197.6,"y":58.1},"noseRootRight":{"x":203.4,"y":56.1},"noseLeftAlarTop":{"x":197.8,"y":67.8},"noseRightAlarTop":{"x":208.5,"y":65.2},"noseLeftAlarOutTip":{"x":196.3,"y":75.2},"noseRightAlarOutTip":{"x":213.1,"y":70.5},"upperLipTop":{"x":205.9,"y":81.5},"upperLipBottom":{"x":206.9,"y":84.1},"underLipTop":{"x":207.0,"y":85.6},"underLipBottom":{"x":208.1,"y":88.9}},"attributes":{"age":28.0,"gender":"male","headPose":{"roll":-15.6,"yaw":-12.4,"pitch":0.0}}},{"faceId":"34e8dda0-f4d0-4a78-bd90-a19177a8cc23","faceRectangle":{"width":30,"height":30,"left":91,"top":0},"faceLandmarks":{"pupilLeft":{"x":99.6,"y":8.0},"pupilRight":{"x":111.0,"y":6.3},"noseTip":{"x":102.9,"y":14.5},"mouthLeft":{"x":100.9,"y":22.4},"mouthRight":{"x":111.7,"y":20.9},"eyebrowLeftOuter":{"x":94.7,"y":5.1},"eyebrowLeftInner":{"x":99.7,"y":4.1},"eyeLeftOuter":{"x":97.6,"y":8.1},"eyeLeftTop":{"x":99.3,"y":7.4},"eyeLeftBottom":{"x":99.5,"y":8.8},"eyeLeftInner":{"x":101.2,"y":7.8},"eyebrowRightInner":{"x":106.1,"y":2.9},"eyebrowRightOuter":{"x":114.9,"y":2.6},"eyeRightInner":{"x":109.0,"y":6.5},"eyeRightTop":{"x":110.8,"y":5.5},"eyeRightBottom":{"x":111.1,"y":7.1},"eyeRightOuter":{"x":112.9,"y":6.0},"noseRootLeft":{"x":102.1,"y":7.7},"noseRootRight":{"x":106.4,"y":7.1},"noseLeftAlarTop":{"x":100.9,"y":12.6},"noseRightAlarTop":{"x":107.0,"y":12.1},"noseLeftAlarOutTip":{"x":100.0,"y":15.1},"noseRightAlarOutTip":{"x":109.3,"y":14.4},"upperLipTop":{"x":104.5,"y":19.8},"upperLipBottom":{"x":104.7,"y":21.4},"underLipTop":{"x":105.4,"y":22.8},"underLipBottom":{"x":105.9,"y":24.7}},"attributes":{"age":62.0,"gender":"male","headPose":{"roll":-4.7,"yaw":-29.5,"pitch":0.0}}},{"faceId":"c65fbf13-60af-4480-b513-82c346d1f1ad","faceRectangle":{"width":28,"height":28,"left":293,"top":6},"faceLandmarks":{"pupilLeft":{"x":300.2,"y":13.1},"pupilRight":{"x":312.0,"y":13.0},"noseTip":{"x":308.2,"y":19.2},"mouthLeft":{"x":300.6,"y":26.1},"mouthRight":{"x":309.5,"y":26.0},"eyebrowLeftOuter":{"x":296.1,"y":10.0},"eyebrowLeftInner":{"x":304.0,"y":10.1},"eyeLeftOuter":{"x":298.1,"y":13.3},"eyeLeftTop":{"x":300.3,"y":12.7},"eyeLeftBottom":{"x":300.3,"y":14.0},"eyeLeftInner":{"x":302.3,"y":13.6},"eyebrowRightInner":{"x":310.5,"y":10.5},"eyebrowRightOuter":{"x":315.2,"y":9.9},"eyeRightInner":{"x":310.1,"y":13.4},"eyeRightTop":{"x":311.9,"y":12.6},"eyeRightBottom":{"x":312.0,"y":13.8},"eyeRightOuter":{"x":313.8,"y":13.3},"noseRootLeft":{"x":304.8,"y":13.6},"noseRootRight":{"x":308.8,"y":13.5},"noseLeftAlarTop":{"x":304.8,"y":17.8},"noseRightAlarTop":{"x":309.5,"y":17.3},"noseLeftAlarOutTip":{"x":303.7,"y":20.2},"noseRightAlarOutTip":{"x":310.4,"y":20.0},"upperLipTop":{"x":306.7,"y":24.3},"upperLipBottom":{"x":306.6,"y":25.7},"underLipTop":{"x":306.3,"y":26.6},"underLipBottom":{"x":306.1,"y":28.5}},"attributes":{"age":80.0,"gender":"female","headPose":{"roll":0.0,"yaw":19.2,"pitch":0.0}}},{"faceId":"e9f97aab-a34d-4189-ac9c-e3ab2d7c1784","faceRectangle":{"width":26,"height":27,"left":0,"top":18},"faceLandmarks":{"pupilLeft":{"x":5.6,"y":23.2},"pupilRight":{"x":15.4,"y":24.4},"noseTip":{"x":12.9,"y":32.9},"mouthLeft":{"x":6.1,"y":37.3},"mouthRight":{"x":12.5,"y":37.9},"eyebrowLeftOuter":{"x":2.9,"y":20.1},"eyebrowLeftInner":{"x":11.5,"y":21.5},"eyeLeftOuter":{"x":4.7,"y":23.1},"eyeLeftTop":{"x":6.9,"y":22.5},"eyeLeftBottom":{"x":6.5,"y":24.4},"eyeLeftInner":{"x":8.4,"y":23.8},"eyebrowRightInner":{"x":14.8,"y":22.1},"eyebrowRightOuter":{"x":20.0,"y":21.2},"eyeRightInner":{"x":14.8,"y":24.6},"eyeRightTop":{"x":16.4,"y":23.4},"eyeRightBottom":{"x":16.3,"y":25.4},"eyeRightOuter":{"x":17.8,"y":24.4},"noseRootLeft":{"x":10.6,"y":24.6},"noseRootRight":{"x":13.7,"y":25.0},"noseLeftAlarTop":{"x":10.1,"y":29.6},"noseRightAlarTop":{"x":14.1,"y":29.8},"noseLeftAlarOutTip":{"x":8.8,"y":32.7},"noseRightAlarOutTip":{"x":14.3,"y":32.9},"upperLipTop":{"x":11.3,"y":36.7},"upperLipBottom":{"x":10.9,"y":38.0},"underLipTop":{"x":10.5,"y":38.2},"underLipBottom":{"x":10.2,"y":39.5}},"attributes":{"age":24.0,"gender":"female","headPose":{"roll":5.0,"yaw":28.7,"pitch":0.0}}}]`
	var faceResults []oxford.FaceResult
	if err := json.Unmarshal([]byte(str), &faceResults); err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range faceResults {
		fmt.Println("faceIdL", v.FaceId, "age:", v.Attributes.Age, "gender:", v.Attributes.Gender)
		fmt.Println(v.FaceLandmarks.PupilLeft)
	}

}