/*
 * Places API
 *
 * The Places API provides access to the different features of the Lab5e Places server, including capabilities such as pub/sub, geofencing and storage of tracking data. 
 *
 * API version: v0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goplaces
// SubscriptionOutput The output for the Subscription
type SubscriptionOutput struct {
	// The type of output
	Type string `json:"type"`
	Config SmsOutputAllOfConfig `json:"config"`
}