# \PositionApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePosition**](PositionApi.md#CreatePosition) | **Post** /collections/{collectionId}/trackers/{trackerId}/positions | Create a new position for the Tracker
[**GetPositions**](PositionApi.md#GetPositions) | **Get** /collections/{collectionId}/trackers/{trackerId}/positions | Get all Positions for Tracker



## CreatePosition

> PositionResponse CreatePosition(ctx, collectionId, trackerId, optional)

Create a new position for the Tracker

Create a new position for the Tracker

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64**| The ID of the Collection | 
**trackerId** | **int64**| The ID of the Tracker | 
 **optional** | ***CreatePositionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePositionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **newPositionRequest** | [**optional.Interface of NewPositionRequest**](NewPositionRequest.md)|  | 

### Return type

[**PositionResponse**](PositionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPositions

> []PositionResponse GetPositions(ctx, collectionId, trackerId, optional)

Get all Positions for Tracker

Get all Positions for Tracker

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64**| The ID of the Collection | 
**trackerId** | **int64**| The ID of the Tracker | 
 **optional** | ***GetPositionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPositionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int64**| Limit number of responses (default 255, max 2500) | [default to 255]
 **offset** | **optional.Int64**| Offset the response by the integer given. If provided, needs to be above 0. | [default to 1]

### Return type

[**[]PositionResponse**](PositionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

