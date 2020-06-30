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

type folder struct {
	Name       string      `json:"name,omitempty"`
	Placemarks []placemark `json:"placemarks,omitempty"`
}

type placemarkList struct {
	Name       string      `json:"name,omitempty"`
	GeoCenter  coord       `json:"geoCenter,omitempty"`
	Placemarks []placemark `json:"placemarks,omitempty"`
	Folders    []folder    `json:"folders,omitempty"`
}

func getCentralGeoCordinate(coordinates []coord) coord {
	if len(coordinates) == 1 {
		return coordinates[0]
	}

	x, y, z := 0.0, 0.0, 0.0

	for _, coordinate := range coordinates {
		latitude := coordinate.Latitude * math.Pi / 180
		longitude := coordinate.Longitude * math.Pi / 180

		x += math.Cos(latitude) * math.Cos(longitude)
		y += math.Cos(latitude) * math.Sin(longitude)
		z += math.Sin(latitude)
	}

	total := float64(len(coordinates))

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

func getKMLFile(zipReader zip.Reader) (*zip.File, int, string) {
	if len(zipReader.File) < 1 {
		return nil, 400, "KMZ file content is invalid"
	}
	var fileIndex int
	fileIndex = -1
	for i, file := range zipReader.File {
		if strings.HasSuffix(file.Name, ".kml") {
			fileIndex = i
		}
	}
	if fileIndex == -1 {
		return nil, 400, "No KML file found in the KMZ"
	}
	return zipReader.File[fileIndex], 0, ""
}

func formatPlacemark(rawPlacemark Placemark, index int) (placemark, int, string) {
	var newPlacemark placemark
	newPlacemark.Name = extractName(rawPlacemark)
	newPlacemark.UpdatedAt = rawPlacemark.TimeStamp.When
	newPlacemark.ID = index
	newPlacemark.Description = rawPlacemark.Description
	if len(rawPlacemark.ExtendedData.FeatureTypes.Value) > 0 {
		newPlacemark.FeatureType = rawPlacemark.ExtendedData.FeatureTypes.Value[0]
	}
	return newPlacemark, 0, ""
}

var placemarkColors = []string{
	"red",
	"pink",
	"purple",
	"deeppurple",
	"blue",
	"lightblue",
	"cyan",
	"teal",
	"green",
	"lime",
	"yellow",
	"orange",
	"deeporange",
	"brown",
	"gray",
	"bluegray",
}

func extractName(rawPlacemark Placemark) string {
	name := rawPlacemark.Name
	nameLanguages := rawPlacemark.ExtendedData.Name.Languages
	if len(nameLanguages) >= 1 {
		for _, nameLanguage := range nameLanguages {
			if nameLanguage.LangCode == "en" {
				return nameLanguage.Content
			}
		}
	}
	return name
}

func extractIcon(rawPlacemark Placemark, flag int) icon {
	var newIcon icon
	if strings.Contains(rawPlacemark.StyleURL, "#placemark-") {
		newIcon.Style = strings.ReplaceAll(rawPlacemark.StyleURL, "#placemark-", "")
		newIcon.Category = rawPlacemark.ExtendedData.Icon
	} else {
		if flag == -1 {
			newIcon.Style = "red"
		} else {
			newIcon.Style = placemarkColors[flag%16]
		}
	}
	return newIcon
}

func extractLocation(rawPlacemark Placemark) (coord, int, string) {
	var extractedLoc coord
	var err error
	rawCoordinate := strings.TrimSpace(rawPlacemark.Point.Coordinates)
	tmpCoord := strings.Split(rawCoordinate, ",")
	if len(tmpCoord) < 2 {
		return coord{}, 400, "Coordinates invalid"
	}
	extractedLoc.Longitude, err = strconv.ParseFloat(tmpCoord[0], 64)
	if err != nil {
		return coord{}, 400, "Coordinates invalids Longitude"
	}
	extractedLoc.Latitude, err = strconv.ParseFloat(tmpCoord[1], 64)
	if err != nil {
		return coord{}, 400, "Coordinates invalids Latitude"
	}
	return extractedLoc, 0, ""
}

func formatOutputJSON(jsonKML KML) (placemarkList, int, string) {
	var listLocations []coord
	var output placemarkList
	output.Name = jsonKML.Document.Name
	var index int
	for i, placemarkFolder := range jsonKML.Document.Folders {
		var newFolder folder
		newFolder.Name = placemarkFolder.Name
		for _, placemarkFolderItem := range placemarkFolder.Placemarks {
			placemark, errCode, errStatus := formatPlacemark(placemarkFolderItem, index)
			if errCode != 0 {
				return placemarkList{}, errCode, errStatus
			}
			placemark.Location, errCode, errStatus = extractLocation(placemarkFolderItem)
			if errCode != 0 {
				return placemarkList{}, errCode, placemark.Name + " " + errStatus
			}
			placemark.Icon = extractIcon(placemarkFolderItem, i)
			newFolder.Placemarks = append(newFolder.Placemarks, placemark)
			listLocations = append(listLocations, placemark.Location)
			index++
		}
		output.Folders = append(output.Folders, newFolder)
	}
	for _, rawPlacemark := range jsonKML.Document.Placemarks {
		placemark, errCode, errStatus := formatPlacemark(rawPlacemark, index)
		if errCode != 0 {
			return placemarkList{}, errCode, errStatus
		}
		placemark.Location, errCode, errStatus = extractLocation(rawPlacemark)
		if errCode != 0 {
			return placemarkList{}, errCode, placemark.Name + " " + errStatus
		}
		placemark.Icon = extractIcon(rawPlacemark, -1)
		output.Placemarks = append(output.Placemarks, placemark)
		listLocations = append(listLocations, placemark.Location)
		index++
	}
	output.GeoCenter = getCentralGeoCordinate(listLocations)
	return output, 0, ""
}

func File(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("map")
	if err != nil {
		libHTTP.RespondWithError(w, 400, "Invalid file "+err.Error())
		return
	}
	defer file.Close()
	zipReader, err := zip.NewReader(file, handler.Size)
	if err != nil {
		libHTTP.RespondWithError(w, 400, "Not a zip file")
		return
	}
	kml, errCode, errStatus := getKMLFile(*zipReader)
	if errCode != 0 {
		libHTTP.RespondWithError(w, errCode, errStatus)
		return
	}
	newFile, err := kml.Open()
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
	output, errCode, errStatus := formatOutputJSON(jsonData)
	if errCode != 0 {
		libHTTP.RespondWithError(w, errCode, errStatus)
		return
	}
	libHTTP.RespondWithJSON(w, 200, output)
}
