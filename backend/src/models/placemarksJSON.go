package models

import "time"

type Coord struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

type Icon struct {
	Style    string `json:"style,omitempty"`
	Category string `json:"category,omitempty"`
}

type PlacemarkJSON struct {
	Name        string    `json:"name,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
	Icon        Icon      `json:"icon,omitempty"`
	Location    Coord     `json:"location,omitempty"`
	Description string    `json:"description,omitempty"`
	FeatureType string    `json:"featureType,omitempty"`
	ID          int       `json:"id"`
}

type FolderJSON struct {
	Name       string          `json:"name,omitempty"`
	Placemarks []PlacemarkJSON `json:"placemarks,omitempty"`
}

type PlacemarkList struct {
	Name       string          `json:"name,omitempty"`
	GeoCenter  Coord           `json:"geoCenter,omitempty"`
	Length     int             `json:"length,omitempty"`
	Placemarks []PlacemarkJSON `json:"placemarks,omitempty"`
	Folders    []FolderJSON    `json:"folders,omitempty"`
}
