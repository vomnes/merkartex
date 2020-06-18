package upload

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/kylelemons/godebug/pretty"
	libHTTP "github.com/vomnes/go-library/http"
	libPretty "github.com/vomnes/go-library/pretty"
)

type coord struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

type icon struct {
	Style    string `json:"style,omitempty"`
	Category string `json:"category,omitempty"`
}

type placemark struct {
	Name        string    `json:"name,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
	Icon        icon      `json:"icon,omitempty"`
	Location    coord     `json:"location,omitempty"`
	Description string    `json:"description,omitempty"`
	FeatureType string    `json:"featureType,omitempty"`
	ID          int       `json:"id"`
}

type placemarkList struct {
	Name       string      `json:"name,omitempty"`
	GeoCenter  coord       `json:"geoCenter,omitempty"`
	Placemarks []placemark `json:"placemarks,omitempty"`
}

func getCentralGeoCordinate(placemarks []placemark) coord {
	if len(placemarks) == 1 {
		return placemarks[0].Location
	}

	x, y, z := 0.0, 0.0, 0.0

	for _, placemark := range placemarks {
		latitude := placemark.Location.Latitude * math.Pi / 180
		longitude := placemark.Location.Longitude * math.Pi / 180

		x += math.Cos(latitude) * math.Cos(longitude)
		y += math.Cos(latitude) * math.Sin(longitude)
		z += math.Sin(latitude)
	}

	total := float64(len(placemarks))

	x = x / total
	y = y / total
	z = z / total

	centralLongitude := math.Atan2(y, x)
	centralSquareRoot := math.Sqrt(x*x + y*y)
	centralLatitude := math.Atan2(z, centralSquareRoot)

	return coord{
		Latitude:  centralLatitude * 180 / math.Pi,
		Longitude: centralLongitude * 180 / math.Pi,
	}
}

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
	// pretty.Print(jsonData.Document.ExtendedData)
	// pretty.Print(jsonData.Document.Placemarks)
	// for _, placemark := range jsonData.Document.Placemarks {
	// 	fmt.Println(placemark.Name, placemark.Point.Coordinates)
	// }
	var placemarkList placemarkList
	placemarkList.Name = jsonData.Document.Name
	for index, rawPlacemark := range jsonData.Document.Placemarks {
		var placemark placemark
		placemark.Name = rawPlacemark.Name
		placemark.UpdatedAt = rawPlacemark.TimeStamp.When
		placemark.Icon.Style = strings.ReplaceAll(rawPlacemark.StyleURL, "#placemark-", "")
		placemark.Icon.Category = rawPlacemark.ExtendedData.Icon
		placemark.ID = index
		tmpCoord := strings.Split(rawPlacemark.Point.Coordinates, ",")
		placemark.Location.Longitude, err = strconv.ParseFloat(tmpCoord[0], 64)
		if err != nil {
			libHTTP.RespondWithError(w, 500, "Coordinates invalids Longitude")
			fmt.Println(libPretty.Error(err.Error()))
			return
		}
		placemark.Location.Latitude, err = strconv.ParseFloat(tmpCoord[1], 64)
		if err != nil {
			libHTTP.RespondWithError(w, 500, "Coordinates invalids Latitude")
			fmt.Println(libPretty.Error(err.Error()))
			return
		}
		placemark.Description = rawPlacemark.Description
		if len(rawPlacemark.ExtendedData.FeatureTypes.Value) > 0 {
			placemark.FeatureType = rawPlacemark.ExtendedData.FeatureTypes.Value[0]
		}
		placemarkList.Placemarks = append(placemarkList.Placemarks, placemark)
	}
	placemarkList.GeoCenter = getCentralGeoCordinate(placemarkList.Placemarks)
	pretty.Print(placemarkList)
	libHTTP.RespondWithJSON(w, 200, placemarkList)
}
