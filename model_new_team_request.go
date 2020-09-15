/*
 * Places API
 *
 * The Places API provides access to the different features of the Lab5e Places server, including capabilities such as pub/sub, geofencing and storage of tracking data. 
 *
 * API version: v0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goplaces
// NewTeamRequest NewTeamRequest is a object to be used when creating a new Team
type NewTeamRequest struct {
	// The name of the new Team
	Name string `json:"name,omitempty"`
	// The description of the new Team
	Description string `json:"description,omitempty"`
}
