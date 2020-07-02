package export

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	lib "github.com/vomnes/go-library"
	libHTTP "github.com/vomnes/go-library/http"

	models "../../models"
)

func KMZ(w http.ResponseWriter, r *http.Request) {
	var body map[string]interface{}
	errCode, errStatus, err := lib.ReaderJSONToInterface(r.Body, &body)
	if err != nil {
		libHTTP.RespondWithError(w, errCode, errStatus)
		return
	}

	var newKML models.KML
	data, err := xml.MarshalIndent(newKML, " ", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" " + xml.Header + string(data))
}
