/*
 * Places API
 *
 * The Places API provides access to the different features of the Lab5e Places server, including capabilities such as pub/sub, geofencing and storage of tracking data. 
 *
 * API version: v0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goplaces
// SubscriptionResponse SubscriptionResponse is the response you get when interacting with the Subscription API
type SubscriptionResponse struct {
	// The ID of the Subscription
	Id int64 `json:"id"`
	// The ID of the Team for the Subscription
	TeamId int64 `json:"teamId"`
	// The name of the Subscription
	Name string `json:"name"`
	// The description of the Subscription
	Description string `json:"description"`
	// Boolean showing whether the subscription is active
	Active bool `json:"active"`
	Output SubscriptionOutput `json:"output"`
	TriggerCriteria SubscriptionResponseTriggerCriteria `json:"triggerCriteria"`
	// The ID of the ShapeCollection containing shapes to subscribe to
	ShapeCollectionId int64 `json:"shapeCollectionId"`
	Trackable SubscriptionResponseTrackable `json:"trackable"`
}
