/*
 * Places API
 *
 * The Places API provides access to the different features of the Lab5e Places server, including capabilities such as pub/sub, geofencing and storage of tracking data. 
 *
 * API version: v0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goplaces
// TokenResponse TokenResponse is the response you get when interacting with the Token API
type TokenResponse struct {
	// The Token that can be used for programmatical access
	Token string `json:"token"`
	// The UNIX timestamp of the creation of the Token
	Created int64 `json:"created"`
	// Limits the allowed paths of the token towards the rest of the API. Only the direct path or subpath of this parameter will be allowed
	Resource string `json:"resource"`
	// Determines whether the token is allowed to do POST, PUT and DELETE methods
	PermWrite bool `json:"permWrite"`
}