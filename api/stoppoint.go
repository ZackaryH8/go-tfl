package api

import (
	"fmt"
)

// Constants
const (
	EndpointStopPoint = "StopPoint"
)

// StopPointGetCategories Gets the list of available StopPoint additional information categories
//
// /StopPoint/Meta/Categories
func (c *Client) StopPointGetCategories() (BikePoints, error) {
	var (
		response BikePoints
		uri      = fmt.Sprintf(
			"%s/Meta/Categories?%s",
			c.getEndpointURL(EndpointStopPoint),
			c.getAuthArgs(),
		)
	)

	err := c.getJSON(uri, &response)
	return response, err
}

// StopPointGetTypes Gets the list of available StopPoint types
//
// /StopPoint/Meta/StopTypes
func (c *Client) StopPointGetTypes() (BikePoints, error) {
	var (
		response BikePoints
		uri      = fmt.Sprintf(
			"%s/Meta/StopTypes?%s",
			c.getEndpointURL(EndpointStopPoint),
			c.getAuthArgs(),
		)
	)

	err := c.getJSON(uri, &response)
	return response, err
}

// StopPointGetModes Gets the list of available StopPoint modes
//
// /StopPoint/Meta/Modes
func (c *Client) StopPointGetModes() (BikePoints, error) {
	var (
		response BikePoints
		uri      = fmt.Sprintf(
			"%s/Meta/Modes?%s",
			c.getEndpointURL(EndpointStopPoint),
			c.getAuthArgs(),
		)
	)

	err := c.getJSON(uri, &response)
	return response, err
}

// StopPointGetByID Gets a list of StopPoints corresponding to the given list of stop ids
//
// /StopPoint/{id}
func (c *Client) StopPointGetByID(ids []string, includeCrowdingData bool) (*StopPoint, error) {
	var (
		response StopPoint
		uri      = fmt.Sprintf(
			"%s/%s?includeCrowdingData=%t&%s",
			c.getEndpointURL(EndpointStopPoint),
			c.arrayToCSV(ids),
			includeCrowdingData,
			c.getAuthArgs(),
		)
	)

	err := c.getJSON(uri, &response)
	return &response, err
}
