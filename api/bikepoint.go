package api

import (
	"fmt"
	"log"
	"net/url"
)

// Contants
const (
	EndpointBikePoint = "BikePoint"
)

// BikePointGetAll Gets all bike point locations
//
// /BikePoint
func (c *Client) BikePointGetAll() (BikePoints, error) {
	var (
		response BikePoints
		url      = fmt.Sprintf(
			"%s?%s",
			c.getEndpointURL(EndpointBikePoint),
			c.getAuthArgs(),
		)
	)

	err := c.getJSON(url, &response)
	return response, err
}

// BikePointGetByID Gets the bike point with the given id.
//
// /Bikepoint/{id}
func (c *Client) BikePointGetByID(id string) (*BikePoint, error) {
	var (
		response BikePoint
		url      = fmt.Sprintf(
			"%s/%s?%s",
			c.getEndpointURL(EndpointBikePoint),
			id,
			c.getAuthArgs(),
		)
	)

	err := c.getJSON(url, &response)
	return &response, err
}

// BikePointInRectangle Gets the bike points that lie within the bounding box
// defined by the lat/lon of its north-west and south-east corners.
//
// /BikePoint?swLat={swLat}&swLon={swLon}&neLat={neLat}&neLon={neLon}
func (c *Client) BikePointInRectangle(swLat, swLon, neLat, neLon float64) (
	BikePoints, error) {
	var (
		response BikePoints
		url      = fmt.Sprintf(
			"%s?swLat=%f&swLon=%f&neLat=%f&neLon=%f&%s",
			c.getEndpointURL(EndpointBikePoint),
			swLat,
			swLon,
			neLat,
			neLon,
			c.getAuthArgs(),
		)
	)

	err := c.getJSON(url, &response)
	return response, err
}

// BikePointInLocus Gets the bike points that lie within the locus defined by
// the lat/lon of its centre and a radius in metres.
//
// /BikePoint?lat={lat}&lon={lon}&radius={radius}
func (c *Client) BikePointInLocus(lat, lon float64, radius int) (BikePointsLocus, error) {
	var (
		response BikePointsLocus
		url      = fmt.Sprintf(
			"%s?lat=%f&lon=%f&radius=%d&%s",
			c.getEndpointURL(EndpointBikePoint),
			lat,
			lon,
			radius,
			c.getAuthArgs(),
		)
	)

	err := c.getJSON(url, &response)
	return response, err
}

// BikePointSearch Search for bike stations by their name, a bike points name
// contains location information
//
// /BikePoint/Search?query={query}
func (c *Client) BikePointSearch(query string) (BikePoints, error) {
	var (
		response BikePoints
		url      = fmt.Sprintf(
			"%s/Search?query=%s&%s",
			c.getEndpointURL(EndpointBikePoint),
			url.QueryEscape(query),
			c.getAuthArgs(),
		)
	)

	log.Println(url)
	err := c.getJSON(url, &response)
	return response, err
}
