# \CollectionApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCollection**](CollectionApi.md#CreateCollection) | **Post** /collections | Create a new Collection
[**DeleteCollection**](CollectionApi.md#DeleteCollection) | **Delete** /collections/{collectionId} | Delete a Collection
[**GetCollection**](CollectionApi.md#GetCollection) | **Get** /collections/{collectionId} | Get a single Collection
[**GetCollections**](CollectionApi.md#GetCollections) | **Get** /collections | Get all Collections
[**GetToken**](CollectionApi.md#GetToken) | **Get** /tokens/{tokenId} | Get a single Collection
[**UpdateCollection**](CollectionApi.md#UpdateCollection) | **Put** /collections/{collectionId} | Update a Collection



## CreateCollection

> CollectionResponse CreateCollection(ctx, optional)

Create a new Collection

Create a new Collection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newCollectionRequest** | [**optional.Interface of NewCollectionRequest**](NewCollectionRequest.md)|  | 

### Return type

[**CollectionResponse**](CollectionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCollection

> DeleteCollection(ctx, collectionId)

Delete a Collection

Delete a Collection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64**| The ID of the Collection | 

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


## GetCollection

> CollectionResponse GetCollection(ctx, collectionId)

Get a single Collection

Get a single Collection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64**| The ID of the Collection | 

### Return type

[**CollectionResponse**](CollectionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollections

> []CollectionResponse GetCollections(ctx, optional)

Get all Collections

Get all Collections

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetCollectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCollectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| Limit number of responses (default 255, max 2500) | [default to 255]
 **offset** | **optional.Int64**| Offset the response by the integer given. If provided, needs to be above 0. | [default to 1]

### Return type

[**[]CollectionResponse**](CollectionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToken

> TokenResponse GetToken(ctx, tokenId)

Get a single Collection

Get a single Collection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string**| The Token value | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCollection

> CollectionResponse UpdateCollection(ctx, collectionId, optional)

Update a Collection

Update a new Collection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64**| The ID of the Collection | 
 **optional** | ***UpdateCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editCollectionRequest** | [**optional.Interface of EditCollectionRequest**](EditCollectionRequest.md)|  | 

### Return type

[**CollectionResponse**](CollectionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

