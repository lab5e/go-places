/*
 * Places API
 *
 * The Places API provides access to the different features of the Lab5e Places server, including capabilities such as pub/sub, geofencing and storage of tracking data. 
 *
 * API version: v0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goplaces
// GeoJsonFeatureCollection GeoJSON Feature collection
type GeoJsonFeatureCollection struct {
	Type string `json:"type"`
	Features []GeoJsonFeature `json:"features"`
}