/*
 * Places API
 *
 * The Places API provides access to the different features of the Lab5e Places server, including capabilities such as pub/sub, geofencing and storage of tracking data. 
 *
 * API version: v0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goplaces
// UnauthorizedResponse struct for UnauthorizedResponse
type UnauthorizedResponse struct {
	// Title representation of the error type
	Title string `json:"title"`
	// Number representation of the error type
	Status int32 `json:"status"`
	// Link to information about the error type
	Type string `json:"type"`
}
