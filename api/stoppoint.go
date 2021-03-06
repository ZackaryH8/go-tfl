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

// StopPointGetRouteSectionByID Get the route sections for all the lines that service the given stop point id
//
// /StopPoint/{id}/Route
func (c *Client) StopPointGetRouteSectionByID(id string, serviceTypes []string) (*StopPoint, error) {
	var (
		response StopPoint
		uri      = fmt.Sprintf(
			"%s/%s/Route?%s&%s",
			c.getEndpointURL(EndpointStopPoint),
			id,
			c.arrayToCSV(serviceTypes),
			c.getAuthArgs(),
		)
	)

	err := c.getJSON(uri, &response)
	return &response, err
}

// StopPointGetInRadius Gets a list of StopPoints within {radius} by the specified criteria
//
// /StopPoint/{id}
func (c *Client) StopPointGetInRadius(stopTypes []string, radius int, useStopPointHierarchy bool, modes []string, catergories []string, returnLines bool, latitude, longitude float64) (*StopPoint, error) {
	var (
		response StopPoint
		uri      = fmt.Sprintf(
			"%s?stopTypes=%s&radius=%v&useStopPointHierarchy=%t&modes=%s&catergories=%s&returnLines=%t&latitude=%v&longitude=%v&%s",
			c.getEndpointURL(EndpointStopPoint),
			c.arrayToCSV(stopTypes),
			radius,
			useStopPointHierarchy,
			c.arrayToCSV(modes),
			c.arrayToCSV(catergories),
			returnLines,
			latitude,
			longitude,
			c.getAuthArgs(),
		)
	)

	err := c.getJSON(uri, &response)
	return &response, err
}

// StopPointGetBySMSCode Gets a StopPoint for a given sms code.
//
// /StopPoint/Sms/{smsID}
func (c *Client) StopPointGetBySMSCode(id string) (*StopPoint, error) {
	var (
		response StopPoint
		uri      = fmt.Sprintf(
			"%s/Sms/%s?%s",
			c.getEndpointURL(EndpointStopPoint),
			id,
			c.getAuthArgs(),
		)
	)

	err := c.getJSON(uri, &response)
	return &response, err
}

// StopPointGetTaxiRanksByID Gets a list of taxi ranks corresponding to the given stop point id
//
// /StopPoint/{id}/TaxiRanks
func (c *Client) StopPointGetTaxiRanksByID(id string) (*StopPoint, error) {
	var (
		response StopPoint
		uri      = fmt.Sprintf(
			"%s/%s/TaxiRanks?%s",
			c.getEndpointURL(EndpointStopPoint),
			id,
			c.getAuthArgs(),
		)
	)

	err := c.getJSON(uri, &response)
	return &response, err
}

// StopPointGetCarParksByID Get car parks corresponding to the given stop point id
//
// /StopPoint/{id}/CarParks
func (c *Client) StopPointGetCarParksByID(id string) (*StopPoint, error) {
	var (
		response StopPoint
		uri      = fmt.Sprintf(
			"%s/%s/CarParks?%s",
			c.getEndpointURL(EndpointStopPoint),
			id,
			c.getAuthArgs(),
		)
	)

	err := c.getJSON(uri, &response)
	return &response, err
}
