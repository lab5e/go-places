/*
 * Places API
 *
 * The Places API provides access to the different features of the Lab5e Places server, including capabilities such as pub/sub, geofencing and storage of tracking data. 
 *
 * API version: v0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goplaces
// EditTrackerRequest struct for EditTrackerRequest
type EditTrackerRequest struct {
	// The updated Collection ID of the Tracker
	CollectionId int64 `json:"collectionId"`
	// The updated Tracker name
	Name string `json:"name"`
	// The updated Tracker description
	Description string `json:"description"`
}
