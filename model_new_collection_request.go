/*
 * Places API
 *
 * The Places API provides access to the different features of the Lab5e Places server, including capabilities such as pub/sub, geofencing and storage of tracking data. 
 *
 * API version: v0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goplaces
// NewCollectionRequest A NewCollectionRequest is used when creating a new Collection
type NewCollectionRequest struct {
	// The name of the new Collection
	Name string `json:"name,omitempty"`
	// The description of the new Collection
	Description string `json:"description,omitempty"`
	// The Team ID of the new Collection
	TeamId int64 `json:"teamId"`
}
