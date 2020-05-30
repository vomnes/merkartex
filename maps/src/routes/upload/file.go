package upload

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"net/http"

	libHTTP "github.com/vomnes/go-library/http"
)

func File(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("map")
	if err != nil {
		libHTTP.RespondWithError(w, 400, "Invalid file")
		return
	}
	defer file.Close()
	zipReader, err := zip.NewReader(file, handler.Size)
	if err != nil {
		libHTTP.RespondWithError(w, 400, "Not a zip file")
		return
	}
	if len(zipReader.File) != 1 {
		libHTTP.RespondWithError(w, 400, "KMZ file content is invalid - 1")
		return
	}
	newFile, err := zipReader.File[0].Open()
	if err != nil {
		libHTTP.RespondWithError(w, 400, "KMZ file content is invalid - 2")
		return
	}
	defer newFile.Close()
	newFileContent, err := ioutil.ReadAll(newFile)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(newFileContent), newFileSize)
	w.WriteHeader(200)
	w.Write(newFileContent)
}
