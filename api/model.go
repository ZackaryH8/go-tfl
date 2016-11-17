package api

import "time"

// BikePoints is a slice of BikePoint
type BikePoints []*BikePoint

// BikePointsLocus is the response from the BikePointInLocus call
type BikePointsLocus struct {
	Type        string      `json:"$type"`
	CentrePoint []float64   `json:"centrePoint"`
	Places      []BikePoint `json:"places"`
}

// BikePoint describes a single bike point
type BikePoint struct {
	Type                 string       `json:"$type"`
	ID                   string       `json:"id"`
	URL                  string       `json:"url"`
	CommonName           string       `json:"commonName"`
	PlaceType            string       `json:"placeType"`
	Lat                  float64      `json:"lat"`
	Long                 float64      `json:"lon"`
	AdditionalProperties []Properties `json:"additionalProperties"`
}

// Properties contain all the fields that are not standard
type Properties struct {
	Type     string    `json:"$type"`
	Category string    `json:"catogory"`
	Key      string    `json:"key"`
	Val      string    `json:"value"`
	Modified time.Time `json:"modified"`
}
