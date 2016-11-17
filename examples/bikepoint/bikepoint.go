package main

import (
	"log"
	"os"

	"github.com/augier/go-tfl/api"
)

func main() {
	id := os.Getenv("API_ID")
	key := os.Getenv("API_KEY")

	client, err := api.NewClient(api.DefaultOptions(id, key))
	if err != nil {
		panic(err)
	}

	names, err := getAllNames(client)
	if err != nil {
		panic(err)
	}
	log.Printf("The names of all the bikepoints are %v", names)

	latLong, err := getLatLongByID(client, "BikePoints_760")
	if err != nil {
		panic(err)
	}
	log.Printf("The LatLong of this bikepoint are: %v", latLong)

	pointCount, err := getPointsInRegctangle(
		client,
		51.458164,
		-0.206002,
		51.521113,
		-0.078869)
	if err != nil {
		panic(err)
	}
	log.Printf("The number of bikepoints found: %d", pointCount)

	// Craven street: 51.508103, -0.126021
	pointCount, err = getPointsInLocus(
		client,
		51.508103,
		-0.126021,
		1000,
	)
	if err != nil {
		panic(err)
	}
	log.Printf("The number of bikepoints found: %d", pointCount)

	names, err = getBikepointMatching(client, "Bethnal Green")
	if err != nil {
		panic(err)
	}

	log.Printf("The points matching bethnal green are %v", names)
}

// Get all the bikepoints and return a slice of their names
func getAllNames(client *api.Client) ([]string, error) {
	var names []string

	bikePoints, err := client.BikePointGetAll()
	if err != nil {
		return names, err
	}

	for _, bikePoint := range bikePoints {
		names = append(names, bikePoint.CommonName)
	}

	return names, nil
}

// Get a single bikepoint using its ID and return a map of its latlong
func getLatLongByID(client *api.Client, id string) (map[string]float64, error) {
	var latLong = make(map[string]float64)

	bikePoint, err := client.BikePointGetByID(id)
	if err != nil {
		return latLong, err
	}

	latLong["lat"] = bikePoint.Lat
	latLong["long"] = bikePoint.Long

	return latLong, nil
}

func getPointsInRegctangle(client *api.Client, swLat, swLon, neLat, neLon float64) (
	int, error) {
	bikePoints, err := client.BikePointInRectangle(swLat, swLon, neLat, neLon)
	if err != nil {
		return 0, err
	}

	return len(bikePoints), nil
}

func getPointsInLocus(client *api.Client, lat, lon float64, radius int) (
	int, error) {
	bikePoints, err := client.BikePointInLocus(lat, lon, radius)
	if err != nil {
		return 0, err
	}

	return len(bikePoints.Places), nil
}

func getBikepointMatching(client *api.Client, query string) ([]string, error) {
	var names []string
	bikePoints, err := client.BikePointSearch(query)
	if err != nil {
		return names, err
	}

	for _, bikePoint := range bikePoints {
		names = append(names, bikePoint.CommonName)
	}

	return names, nil
}
