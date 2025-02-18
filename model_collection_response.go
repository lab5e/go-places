/*
 * Places API
 *
 * The Places API provides access to the different features of the Lab5e Places server, including capabilities such as pub/sub, geofencing and storage of tracking data. 
 *
 * API version: v0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goplaces
// CollectionResponse CollectionResponse is the response you get when interacting with the Collection API
type CollectionResponse struct {
	// The ID of the Collection
	Id int64 `json:"id"`
	// Name of the Collection
	Name string `json:"name"`
	// Description of the Collection
	Description string `json:"description"`
	// The Team ID of the Collection
	TeamId float32 `json:"teamId,omitempty"`
}
