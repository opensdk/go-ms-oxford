package oxford

import (
	"fmt"
	"io/ioutil"
	"io"
)

type FaceDetection struct {
	ContentType string
	URL         string
	Reader      io.Reader
}

func (self FaceDetection)Detect() error {
	if len(self.URL) == 0 {
//		if nil == self.Reader {
//			return errors.New("nothing to detect")
//		}

		self.ContentType = ContentTypeOctetStream
	} else {
		self.ContentType = ContentTypeJson
	}

	resp, err := post(Config.detectionsURL, self.ContentType, self.Reader)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(resp.Status)
	fmt.Println(string(output))


	return nil
}

