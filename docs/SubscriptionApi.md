# \SubscriptionApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](SubscriptionApi.md#CreateSubscription) | **Post** /subscriptions | Create a new Subscription
[**DeleteSubscription**](SubscriptionApi.md#DeleteSubscription) | **Delete** /subscriptions/{subscriptionId} | Delete a Subscription
[**GetSubscription**](SubscriptionApi.md#GetSubscription) | **Get** /subscriptions/{subscriptionId} | Get a single Subscription
[**GetSubscriptions**](SubscriptionApi.md#GetSubscriptions) | **Get** /subscriptions | Get all Subscriptions
[**UpdateSubscription**](SubscriptionApi.md#UpdateSubscription) | **Put** /subscriptions/{subscriptionId} | Update a Subscription



## CreateSubscription

> SubscriptionResponse CreateSubscription(ctx, optional)

Create a new Subscription

Create a new Subscription

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newSubscriptionRequest** | [**optional.Interface of NewSubscriptionRequest**](NewSubscriptionRequest.md)|  | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscription

> DeleteSubscription(ctx, subscriptionId)

Delete a Subscription

Delete a Subscription

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **int64**| The ID of the Subscription | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscription

> SubscriptionResponse GetSubscription(ctx, subscriptionId)

Get a single Subscription

Get a single Subscription

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **int64**| The ID of the Subscription | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptions

> []SubscriptionResponse GetSubscriptions(ctx, optional)

Get all Subscriptions

Get all Subscriptions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSubscriptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| Limit number of responses (default 255, max 2500) | [default to 255]
 **offset** | **optional.Int64**| Offset the response by the integer given. If provided, needs to be above 0. | [default to 1]

### Return type

[**[]SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscription

> SubscriptionResponse UpdateSubscription(ctx, subscriptionId, optional)

Update a Subscription

Update a Subscription

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **int64**| The ID of the Subscription | 
 **optional** | ***UpdateSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editSubscriptionRequest** | [**optional.Interface of EditSubscriptionRequest**](EditSubscriptionRequest.md)|  | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

