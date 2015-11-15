package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/upload", func(writer http.ResponseWriter, req *http.Request) {
		req.ParseMultipartForm(2 << 4)
		file, header, err := req.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(header.Filename)
		f, _ := os.OpenFile("/Users/du/Desktop/aa/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		defer f.Close()
		io.Copy(f, file)

	})
	http.ListenAndServe(":8080", nil)
}
