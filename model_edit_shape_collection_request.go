/*
 * Places API
 *
 * The Places API provides access to the different features of the Lab5e Places server, including capabilities such as pub/sub, geofencing and storage of tracking data. 
 *
 * API version: v0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goplaces
// EditShapeCollectionRequest struct for EditShapeCollectionRequest
type EditShapeCollectionRequest struct {
	// The updated Team ID of the ShapeCollection
	TeamId int64 `json:"teamId"`
	// The updated name of the ShapeCollection
	Name string `json:"name"`
	// The updated description of the ShapeCollection
	Description string `json:"description"`
}
