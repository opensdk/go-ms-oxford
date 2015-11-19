package util

import (
	"bytes"
	"fmt"
	"github.com/opensdk/go-ms-oxford"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

func HandleResponse(response *http.Response) (body []byte, err error) {
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	return
}

func Convert2String(obj interface{}) string {
	if obj == nil {
		return ""
	}

	if str, ok := obj.(string); ok {
		return str
	}

	return ""
}

type Point struct {
	X float64
	Y float64
}
