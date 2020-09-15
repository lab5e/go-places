/*
 * Places API
 *
 * The Places API provides access to the different features of the Lab5e Places server, including capabilities such as pub/sub, geofencing and storage of tracking data. 
 *
 * API version: v0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goplaces
// SubscriptionTriggerType the model 'SubscriptionTriggerType'
type SubscriptionTriggerType string

// List of SubscriptionTriggerType
const (
	INSIDE SubscriptionTriggerType = "inside"
	OUTSIDE SubscriptionTriggerType = "outside"
	ENTERED SubscriptionTriggerType = "entered"
	EXITED SubscriptionTriggerType = "exited"
)
