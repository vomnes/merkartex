package upload

import (
	"encoding/xml"
	"time"
)

// Style define a style placemarks
type Style struct {
	StyleID   string `xml:"id,attr"`
	IconStyle struct {
		Icon struct {
			Href string `xml:"href"`
		}
	}
}

// ContentLanguage is the XML structure with the content and his language code
type ContentLanguage struct {
	Content  string `xml:",chardata"`
	LangCode string `xml:"code,attr"`
}

// ExtendedData describe the extended data
type ExtendedData struct {
	XMLName xml.Name
	// ExtendedDataType string `xml:"xmlns mwm, ExtendedData"`
	Name struct {
		Languages []ContentLanguage `xml:",any"`
	} `xml:"name"`
	Annotation  struct{} `xml:"annotation"`
	Description struct {
		Languages []ContentLanguage `xml:",any"`
	} `xml:"description"`
	FeatureTypes struct {
		Value []string `xml:"value"`
	} `xml:"featureTypes"`
	CustomName struct {
		Languages []ContentLanguage `xml:",any"`
	} `xml:"customName"`
	LastModified string `xml:"lastModified"`
	AccessRules  string `xml:"accessRules"`
	Scale        int    `xml:"scale"`
	Visibility   int    `xml:"visibility"`
}

// Placemark describe a place on the map
type Placemark struct {
	Name        string `xml:"name"`
	Description string `xml:"description"`
	TimeStamp   struct {
		When time.Time `xml:"when"`
	}
	StyleURL string `xml:"styleUrl"`
	Point    struct {
		Coordinates string `xml:"coordinates"`
	}
	ExtendedData ExtendedData
}

// Document is the structure with the content
type Document struct {
	Styles       []Style      `xml:"Style"`
	Name         string       `xml:"name"`
	Visibility   int          `xml:"visibility"`
	ExtendedData ExtendedData `xml:"ExtendedData"`
	Placemarks   []Placemark  `xml:"Placemark"`
}

// KML is the Keyhole Markup Language structure
type KML struct {
	Document Document
}
