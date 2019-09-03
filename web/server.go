package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

func NewWebServer() {
	http.HandleFunc("/addjob", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	r.ParseMultipartForm(10 << 20)

	files := r.MultipartForm.File["myFile[]"]
	jobid := ""
	for _, handler := range files {
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		// Create a temporary file within our temp-images directory that follows
		// a particular naming pattern
		suff := path.Ext(handler.Filename)
		pattern := "job-*" + suff
		tempFile, err := ioutil.TempFile("./uploads", pattern)
		if err != nil {
			fmt.Println("io util:", err)
		}
		defer tempFile.Close()

		file, err := handler.Open()
		if err != nil {
			fmt.Println(err)
		}

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		// write this byte array to our temporary file
		tempFile.Write(fileBytes)
		if suff == ".go" {
			jobid = strings.TrimSuffix(path.Base(tempFile.Name()), suff)
			// return that we have successfully uploaded our file!
			fmt.Fprintf(w, "Successfully Uploaded File As: %s \n", jobid)

			go func() {
				if loader, err := NewLoader(); err != nil {
					fmt.Println(err)
				} else {
					err = loader.CompileAndRun(tempFile.Name())
					if err != nil {
						fmt.Println("load:", err)
					}
					fmt.Println(tempFile.Name())
				}
			}()
		}
	}

	return

}
