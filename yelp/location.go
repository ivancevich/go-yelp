package yelp

import "net/url"

// Location is where the business is located.
type Location struct {
	Address1       string   `json:"address1"`
	Address2       string   `json:"address2"`
	Address3       string   `json:"address3"`
	City           string   `json:"city"`
	State          string   `json:"state"`
	ZipCode        string   `json:"zip_code"`
	Country        string   `json:"country"`
	DisplayAddress []string `json:"display_address"`
	CrossStreets   string   `json:"cross_streets"`
}

// Coordinates defines a location with Latitude and Longitude.
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// URLValues returns Coordinates as url.Values.
func (c Coordinates) URLValues() url.Values {
	vals := url.Values{}
	vals.Add("latitude", FloatString(c.Latitude))
	vals.Add("longitude", FloatString(c.Longitude))
	return vals
}

// Region defines an area of the businesses.
type Region struct {
	Center Coordinates `json:"center"`
}

// URLValues returns Region as url.Values.
func (r Region) URLValues() url.Values {
	return r.Center.URLValues()
}
