package upload

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kylelemons/godebug/pretty"
	libHTTP "github.com/vomnes/go-library/http"
	libPretty "github.com/vomnes/go-library/pretty"
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
	var jsonData KML
	err = xml.Unmarshal(newFileContent, &jsonData)
	if err != nil {
		libHTTP.RespondWithError(w, 500, "Invalid XML content")
		fmt.Println(libPretty.Error(err.Error()))
		return
	}
	pretty.Print(jsonData.Document)
	w.WriteHeader(200)
	w.Write(newFileContent)
}
