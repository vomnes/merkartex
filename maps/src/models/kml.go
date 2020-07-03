package models

import (
	"encoding/xml"
	"time"
)

// Style define a style placemarks
type Style struct {
	ID            string `xml:"id,attr"`
	IconStyleHref string `xml:"IconStyle>Icon>href"`
}

// ContentLanguage is the XML structure with the content and his language code
type ContentLanguage struct {
	Content  string `xml:",chardata"`
	LangCode string `xml:"code,attr"`
}

// ExtendedData describe the extended data
type ExtendedData struct {
	XMLName xml.Name
	XMLNS   string `xml:"xmlns:mwm,attr"`
	Name    struct {
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
	} `xml:"customName,omitempty"`
	LastModified string `xml:"lastModified,omitempty"`
	AccessRules  string `xml:"accessRules,omitempty"`
	Scale        int    `xml:"scale,omitempty"`
	Icon         string `xml:"icon,omitempty"`
	Visibility   int    `xml:"visibility,omitempty"`
}

// ExtendedDataMVM describe the extended data
type ExtendedDataMVM struct {
	XMLName xml.Name
	XMLNS   string `xml:"xmlns:mwm,attr"`
	Name    struct {
		Languages []ContentLanguage `xml:"mvm:lang,omitempty"`
	} `xml:"mwm:name,omitempty"`
	Annotation  struct{} `xml:"mwm:annotation,omitempty"`
	Description struct {
		Languages []ContentLanguage `xml:"mvm:lang,omitempty"`
	} `xml:"mwm:description,omitempty"`
	FeatureTypes struct {
		Value []string `xml:"mwm:value,omitempty"`
	} `xml:"mwm:featureTypes,omitempty"`
	CustomName struct {
		Languages []ContentLanguage `xml:",any,omitempty"`
	} `xml:"mwm:customName,omitempty"`
	LastModified string `xml:"mwm:lastModified,omitempty"`
	AccessRules  string `xml:"mwm:accessRules,omitempty"`
	Scale        int    `xml:"mwm:scale,omitempty"`
	Icon         string `xml:"mwm:icon,omitempty"`
	Visibility   int    `xml:"mwm:visibility,omitempty"`
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

// PlacemarkMVM describe a place on the map
type PlacemarkMVM struct {
	Name         string    `xml:"name"`
	Description  string    `xml:"description"`
	TimeStamp    time.Time `xml:"TimeStamp>when"`
	StyleURL     string    `xml:"styleUrl"`
	Coordinates  string    `xml:"Point>coordinates"`
	ExtendedData ExtendedDataMVM
}

// Folder describe a group of placemarks
type Folder struct {
	Name       string      `xml:"name"`
	Placemarks []Placemark `xml:"Placemark"`
}

// FolderMVM describe a group of placemarks
type FolderMVM struct {
	Name       string         `xml:"name"`
	Placemarks []PlacemarkMVM `xml:"Placemark"`
}

// Document is the structure with the content
type Document struct {
	Styles       []Style      `xml:"Style"`
	Name         string       `xml:"name"`
	Visibility   int          `xml:"visibility"`
	ExtendedData ExtendedData `xml:"ExtendedData"`
	Placemarks   []Placemark  `xml:"Placemark"`
	Folders      []Folder     `xml:"Folder"`
}

// DocumentMVM is the structure with the content
type DocumentMVM struct {
	Styles       []Style         `xml:"Style"`
	Name         string          `xml:"name"`
	Visibility   int             `xml:"visibility"`
	ExtendedData ExtendedDataMVM `xml:"ExtendedData"`
	Placemarks   []PlacemarkMVM  `xml:"Placemark"`
	Folders      []FolderMVM     `xml:"Folder"`
}

// KML is the Keyhole Markup Language structure
type KML struct {
	Document Document
	XMLNS    string `xml:"xmlns,attr"`
}
