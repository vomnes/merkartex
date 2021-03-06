package upload

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	lib "github.com/vomnes/go-library"
	libHTTP "github.com/vomnes/go-library/http"
	libPretty "github.com/vomnes/go-library/pretty"

	"merkartex/shared/models"
)

func getCentralGeoCordinate(coordinates []models.Coord) models.Coord {
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

	return models.Coord{
		Latitude:  centralLatitude * 180 / math.Pi,
		Longitude: centralLongitude * 180 / math.Pi,
	}
}

func getKMZFromGoogle(link string) (string, io.ReaderAt, int64, int, string) {
	u, err := url.Parse(link)
	if err != nil {
		return "", *new(io.ReaderAt), 0, 400, "Google My Maps URL not well formated"
	}
	q := u.Query()
	if len(q["mid"]) <= 0 {
		return "", *new(io.ReaderAt), 0, 400, "Google My Maps URL doesn't contains the required informations"
	}
	mid := q["mid"][0]
	client := http.DefaultClient
	respHTTP, err := http.NewRequest(
		"GET",
		"https://www.google.com/maps/d/kml?mid="+mid,
		nil,
	)
	if err != nil {
		return "", *new(io.ReaderAt), 0, 500, err.Error()
	}
	resp, err := client.Do(respHTTP)
	if err != nil {
		return "", *new(io.ReaderAt), 0, 400, err.Error()
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", *new(io.ReaderAt), 0, 400, err.Error()
	}
	return "google-my-maps", bytes.NewReader(bodyBytes), int64(len(bodyBytes)), 0, ""
}

func getKMZ(r *http.Request) (string, io.ReaderAt, int64, int, string) {
	var body map[string]string
	file, handler, err := r.FormFile("map")
	if err != nil {
		errCode, _, err := lib.ReaderJSONToInterface(r.Body, &body)
		if err != nil {
			return "", *new(io.ReaderAt), 0, errCode, "Invalid data send"
		}
		if body["google-my-maps-link"] == "" {
			return "", *new(io.ReaderAt), 0, 400, "No Google Maps Link"
		}
		return getKMZFromGoogle(body["google-my-maps-link"])
	}
	return "kmz-file", file, handler.Size, 0, ""
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

func formatPlacemark(rawPlacemark models.Placemark, index int) (models.PlacemarkJSON, int, string) {
	var newPlacemark models.PlacemarkJSON
	newPlacemark.Name = extractName(rawPlacemark)
	newPlacemark.UpdatedAt = rawPlacemark.TimeStamp.When
	newPlacemark.ID = index
	newPlacemark.Description = rawPlacemark.Description
	if len(rawPlacemark.ExtendedData.FeatureTypes.Value) > 0 {
		newPlacemark.FeatureType = rawPlacemark.ExtendedData.FeatureTypes.Value[0]
	}
	return newPlacemark, 0, ""
}

func extractName(rawPlacemark models.Placemark) string {
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

func extractIcon(rawPlacemark models.Placemark, flag int) models.Icon {
	var newIcon models.Icon
	if strings.Contains(rawPlacemark.StyleURL, "#placemark-") {
		newIcon.Style = strings.ReplaceAll(rawPlacemark.StyleURL, "#placemark-", "")
		newIcon.Category = rawPlacemark.ExtendedData.Icon
	} else {
		if flag == -1 {
			newIcon.Style = "red"
		} else {
			newIcon.Style = models.PlacemarkColors[flag%16]
		}
	}
	return newIcon
}

func extractLocation(rawPlacemark models.Placemark) (models.Coord, int, string) {
	var extractedLoc models.Coord
	var err error
	rawCoordinate := strings.TrimSpace(rawPlacemark.Point.Coordinates)
	tmpCoord := strings.Split(rawCoordinate, ",")
	if len(tmpCoord) < 2 {
		return models.Coord{}, 400, "Coordinates invalid"
	}
	extractedLoc.Longitude, err = strconv.ParseFloat(tmpCoord[0], 64)
	if err != nil {
		return models.Coord{}, 400, "Coordinates invalids Longitude"
	}
	extractedLoc.Latitude, err = strconv.ParseFloat(tmpCoord[1], 64)
	if err != nil {
		return models.Coord{}, 400, "Coordinates invalids Latitude"
	}
	return extractedLoc, 0, ""
}

func formatOutputJSON(jsonKML models.KML) (models.PlacemarkList, int, string) {
	var listLocations []models.Coord
	var output models.PlacemarkList
	output.Name = jsonKML.Document.Name
	var index int
	for i, placemarkFolder := range jsonKML.Document.Folders {
		var newFolder models.FolderJSON
		newFolder.Name = placemarkFolder.Name
		for _, placemarkFolderItem := range placemarkFolder.Placemarks {
			placemark, errCode, errStatus := formatPlacemark(placemarkFolderItem, index)
			if errCode != 0 {
				return models.PlacemarkList{}, errCode, errStatus
			}
			placemark.Location, errCode, errStatus = extractLocation(placemarkFolderItem)
			if errCode != 0 {
				fmt.Println(libPretty.Error(placemark.Name + " " + errStatus))
				continue
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
			return models.PlacemarkList{}, errCode, errStatus
		}
		placemark.Location, errCode, errStatus = extractLocation(rawPlacemark)
		if errCode != 0 {
			fmt.Println(libPretty.Error(placemark.Name + " >> " + errStatus))
			continue
		}
		placemark.Icon = extractIcon(rawPlacemark, -1)
		output.Placemarks = append(output.Placemarks, placemark)
		listLocations = append(listLocations, placemark.Location)
		index++
	}
	output.GeoCenter = getCentralGeoCordinate(listLocations)
	output.Length = index
	return output, 0, ""
}

func File(w http.ResponseWriter, r *http.Request) {
	kmzKind, file, fileSize, errCode, errStatus := getKMZ(r)
	if errCode != 0 && errStatus != "" {
		libHTTP.RespondWithError(w, errCode, errStatus)
		return
	}
	zipReader, err := zip.NewReader(file, fileSize)
	if err != nil {
		if kmzKind == "google-my-maps" {
			libHTTP.RespondWithError(w, 400, "Google My Maps Link doesn't return a valid KMZ file")
			return
		}
		libHTTP.RespondWithError(w, 400, "Not a good KMZ file")
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
	var jsonData models.KML
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
