package export

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"log"
	"net/http"
	"strconv"

	lib "github.com/vomnes/go-library"
	libHTTP "github.com/vomnes/go-library/http"

	"merkartex/shared/models"
)

// KML ...
type KML struct {
	Document models.DocumentMVM
	XMLNS    string `xml:"xmlns,attr"`
}

func formatPlacemarkKML(placemark models.PlacemarkJSON) models.PlacemarkMVM {
	newPlacemark := models.PlacemarkMVM{
		Name:        placemark.Name,
		Description: placemark.Description,
		TimeStamp:   placemark.UpdatedAt,
		Coordinates: strconv.FormatFloat(placemark.Location.Longitude, 'f', 6, 64) + "," + strconv.FormatFloat(placemark.Location.Latitude, 'f', 6, 64),
		StyleURL:    "#placemark-" + placemark.Icon.Style,
		ExtendedData: models.ExtendedDataMVM{
			XMLNS:      "https://maps.me",
			Visibility: 1,
			Scale:      8,
		},
	}
	if placemark.Icon.Category != "" && placemark.Icon.Category != "None" {
		newPlacemark.ExtendedData.Icon = placemark.Icon.Category
	}
	newPlacemark.ExtendedData.Name.Languages = append(
		newPlacemark.ExtendedData.Name.Languages,
		models.ContentLanguage{
			Content:  placemark.Name,
			LangCode: "default",
		})
	if placemark.FeatureType != "" {
		newPlacemark.ExtendedData.FeatureTypes.Value = append(newPlacemark.ExtendedData.FeatureTypes.Value, placemark.FeatureType)
	}
	return newPlacemark
}

func KMZ(w http.ResponseWriter, r *http.Request) {
	var placemarksJSON models.PlacemarkList
	errCode, errStatus, err := lib.ReaderJSONToInterface(r.Body, &placemarksJSON)
	if err != nil {
		libHTTP.RespondWithError(w, errCode, errStatus)
		return
	}
	var newKML KML
	newKML.XMLNS = "http://earth.google.com/kml/2.2"
	newKML.Document.ExtendedData.XMLNS = "https://maps.me"
	for _, color := range models.PlacemarkColors {
		newColor := models.Style{
			ID:            "placemark-" + color,
			IconStyleHref: "http://maps.me/placemarks/placemark-" + color + ".png",
		}
		newKML.Document.Styles = append(newKML.Document.Styles, newColor)
	}
	newKML.Document.Name = placemarksJSON.Name
	newKML.Document.Visibility = 1
	newKML.Document.ExtendedData.Name.Languages =
		append(
			newKML.Document.ExtendedData.Name.Languages,
			models.ContentLanguage{
				Content:  placemarksJSON.Name,
				LangCode: "default",
			},
		)
	for _, placemark := range placemarksJSON.Placemarks {
		newPlacemark := formatPlacemarkKML(placemark)
		newKML.Document.Placemarks = append(newKML.Document.Placemarks, newPlacemark)
	}
	for _, folder := range placemarksJSON.Folders {
		for _, placemark := range folder.Placemarks {
			newPlacemark := formatPlacemarkKML(placemark)
			newKML.Document.Placemarks = append(newKML.Document.Placemarks, newPlacemark)
		}
	}
	data, err := xml.MarshalIndent(newKML, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	writerZip := zip.NewWriter(buf)
	zippedF, err := writerZip.Create(placemarksJSON.Name + ".kml")
	if err != nil {
		libHTTP.RespondWithError(w, 500, err.Error())
		return
	}
	_, err = zippedF.Write([]byte(xml.Header + string(data)))
	if err != nil {
		libHTTP.RespondWithError(w, 500, err.Error())
		return
	}
	err = writerZip.Close()
	if err != nil {
		libHTTP.RespondWithError(w, 500, err.Error())
		return
	}
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("Content-Type", "application/vnd.google-earth.kmz")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+placemarksJSON.Name+".kmz\"")
	w.WriteHeader(200)
	w.Write(buf.Bytes())
}
