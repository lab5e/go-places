# \TrackerApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTracker**](TrackerApi.md#CreateTracker) | **Post** /collections/{collectionId}/trackers | Create a new Tracker
[**DeleteTracker**](TrackerApi.md#DeleteTracker) | **Delete** /collections/{collectionId}/trackers/{trackerId} | Delete a Tracker
[**GetTracker**](TrackerApi.md#GetTracker) | **Get** /collections/{collectionId}/trackers/{trackerId} | Get a single Tracker
[**GetTrackers**](TrackerApi.md#GetTrackers) | **Get** /collections/{collectionId}/trackers | Get all Trackers
[**UpdateTracker**](TrackerApi.md#UpdateTracker) | **Put** /collections/{collectionId}/trackers/{trackerId} | Update a Tracker



## CreateTracker

> TrackerResponse CreateTracker(ctx, collectionId, optional)

Create a new Tracker

Create a new Tracker

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64**| The ID of the Collection | 
 **optional** | ***CreateTrackerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTrackerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newTrackerRequest** | [**optional.Interface of NewTrackerRequest**](NewTrackerRequest.md)|  | 

### Return type

[**TrackerResponse**](TrackerResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTracker

> DeleteTracker(ctx, collectionId, trackerId)

Delete a Tracker

Delete a Tracker

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64**| The ID of the Collection | 
**trackerId** | **int64**| The ID of the Tracker | 

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


## GetTracker

> TrackerResponse GetTracker(ctx, collectionId, trackerId)

Get a single Tracker

Get a single Tracker

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64**| The ID of the Collection | 
**trackerId** | **int64**| The ID of the Tracker | 

### Return type

[**TrackerResponse**](TrackerResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrackers

> []TrackerResponse GetTrackers(ctx, collectionId, optional)

Get all Trackers

Get all Trackers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64**| The ID of the Collection | 
 **optional** | ***GetTrackersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTrackersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int64**| Limit number of responses (default 255, max 2500) | [default to 255]
 **offset** | **optional.Int64**| Offset the response by the integer given. If provided, needs to be above 0. | [default to 1]

### Return type

[**[]TrackerResponse**](TrackerResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTracker

> TrackerResponse UpdateTracker(ctx, collectionId, trackerId, optional)

Update a Tracker

Update a Tracker

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64**| The ID of the Collection | 
**trackerId** | **int64**| The ID of the Tracker | 
 **optional** | ***UpdateTrackerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTrackerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **editTrackerRequest** | [**optional.Interface of EditTrackerRequest**](EditTrackerRequest.md)|  | 

### Return type

[**TrackerResponse**](TrackerResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

