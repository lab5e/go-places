# \ShapeCollectionApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateShapeCollection**](ShapeCollectionApi.md#CreateShapeCollection) | **Post** /shapecollections | Create a new ShapeCollection
[**CreateShapeCollectionFeature**](ShapeCollectionApi.md#CreateShapeCollectionFeature) | **Post** /shapecollections/{shapeCollectionId}/geojson | Create new GeoJSON Feature
[**DeleteShape**](ShapeCollectionApi.md#DeleteShape) | **Delete** /shapecollections/{shapeCollectionId}/shapes/{shapeId} | Delete a Shape
[**DeleteShapeCollection**](ShapeCollectionApi.md#DeleteShapeCollection) | **Delete** /shapecollections/{shapeCollectionId} | Delete a ShapeCollection
[**GetShape**](ShapeCollectionApi.md#GetShape) | **Get** /shapecollections/{shapeCollectionId}/shapes/{shapeId} | Get Shape
[**GetShapeCollection**](ShapeCollectionApi.md#GetShapeCollection) | **Get** /shapecollections/{shapeCollectionId} | Get a single ShapeCollection
[**GetShapeCollectionFeatureCollection**](ShapeCollectionApi.md#GetShapeCollectionFeatureCollection) | **Get** /shapecollections/{shapeCollectionId}/geojson | Get a GeoJSON FeatureCollection for a ShapeCollection
[**GetShapeCollectionShapes**](ShapeCollectionApi.md#GetShapeCollectionShapes) | **Get** /shapecollections/{shapeCollectionId}/shapes | Get Shapes for a ShapeCollection
[**GetShapeCollections**](ShapeCollectionApi.md#GetShapeCollections) | **Get** /shapecollections | Get all ShapeCollections
[**GetShapeFeature**](ShapeCollectionApi.md#GetShapeFeature) | **Get** /shapecollections/{shapeCollectionId}/shapes/{shapeId}/geojson | Get Shape GeoJSON Feature
[**ReplaceShapeCollectionFeatureCollection**](ShapeCollectionApi.md#ReplaceShapeCollectionFeatureCollection) | **Put** /shapecollections/{shapeCollectionId}/geojson | Replace GeoJSON FeatureCollection
[**UpdateShapeCollection**](ShapeCollectionApi.md#UpdateShapeCollection) | **Put** /shapecollections/{shapeCollectionId} | Edit a ShapeCollection
[**UpdateShapeCollectionFeature**](ShapeCollectionApi.md#UpdateShapeCollectionFeature) | **Put** /shapecollections/{shapeCollectionId}/shapes/{shapeId}/geojson | Update Shape GeoJSON Feature



## CreateShapeCollection

> ShapeCollectionResponse CreateShapeCollection(ctx, optional)

Create a new ShapeCollection

Create a new ShapeCollection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateShapeCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateShapeCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newShapeCollectionRequest** | [**optional.Interface of NewShapeCollectionRequest**](NewShapeCollectionRequest.md)|  | 

### Return type

[**ShapeCollectionResponse**](ShapeCollectionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShapeCollectionFeature

> GeoJsonFeature CreateShapeCollectionFeature(ctx, shapeCollectionId, optional)

Create new GeoJSON Feature

Create new GeoJSON Feature

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shapeCollectionId** | **int64**| The ID of the ShapeCollection | 
 **optional** | ***CreateShapeCollectionFeatureOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateShapeCollectionFeatureOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **geoJsonFeature** | [**optional.Interface of GeoJsonFeature**](GeoJsonFeature.md)|  | 

### Return type

[**GeoJsonFeature**](GeoJSONFeature.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteShape

> DeleteShape(ctx, shapeCollectionId, shapeId)

Delete a Shape

Delete a Shape

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shapeCollectionId** | **int64**| The ID of the ShapeCollection | 
**shapeId** | **int64**| The ID of the Shape | 

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


## DeleteShapeCollection

> DeleteShapeCollection(ctx, shapeCollectionId)

Delete a ShapeCollection

Delete a ShapeCollection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shapeCollectionId** | **int64**| The ID of the ShapeCollection | 

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


## GetShape

> ShapeResponse GetShape(ctx, shapeCollectionId, shapeId)

Get Shape

Get Shape

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shapeCollectionId** | **int64**| The ID of the ShapeCollection | 
**shapeId** | **int64**| The ID of the Shape | 

### Return type

[**ShapeResponse**](ShapeResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShapeCollection

> ShapeCollectionResponse GetShapeCollection(ctx, shapeCollectionId)

Get a single ShapeCollection

Get a single ShapeCollection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shapeCollectionId** | **int64**| The ID of the ShapeCollection | 

### Return type

[**ShapeCollectionResponse**](ShapeCollectionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShapeCollectionFeatureCollection

> map[string]map[string]interface{} GetShapeCollectionFeatureCollection(ctx, shapeCollectionId, optional)

Get a GeoJSON FeatureCollection for a ShapeCollection

Get a GeoJSON FeatureCollection for a ShapeCollection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shapeCollectionId** | **int64**| The ID of the ShapeCollection | 
 **optional** | ***GetShapeCollectionFeatureCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetShapeCollectionFeatureCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int64**| Limit number of responses (default 255, max 2500) | [default to 255]
 **offset** | **optional.Int64**| Offset the response by the integer given. If provided, needs to be above 0. | [default to 1]

### Return type

**map[string]map[string]interface{}**

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShapeCollectionShapes

> []ShapeResponse GetShapeCollectionShapes(ctx, shapeCollectionId)

Get Shapes for a ShapeCollection

Get Shapes for a ShapeCollection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shapeCollectionId** | **int64**| The ID of the ShapeCollection | 

### Return type

[**[]ShapeResponse**](ShapeResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShapeCollections

> []ShapeCollectionResponse GetShapeCollections(ctx, optional)

Get all ShapeCollections

Get all ShapeCollections

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetShapeCollectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetShapeCollectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| Limit number of responses (default 255, max 2500) | [default to 255]
 **offset** | **optional.Int64**| Offset the response by the integer given. If provided, needs to be above 0. | [default to 1]

### Return type

[**[]ShapeCollectionResponse**](ShapeCollectionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShapeFeature

> GeoJsonFeature GetShapeFeature(ctx, shapeCollectionId, shapeId)

Get Shape GeoJSON Feature

Get Shape GeoJSON Feature

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shapeCollectionId** | **int64**| The ID of the ShapeCollection | 
**shapeId** | **int64**| The ID of the Shape | 

### Return type

[**GeoJsonFeature**](GeoJSONFeature.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceShapeCollectionFeatureCollection

> map[string]map[string]interface{} ReplaceShapeCollectionFeatureCollection(ctx, shapeCollectionId, optional)

Replace GeoJSON FeatureCollection

Replace GeoJSON FeatureCollection. Note that this will replace all shapes and generate new IDs for all shapes. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shapeCollectionId** | **int64**| The ID of the ShapeCollection | 
 **optional** | ***ReplaceShapeCollectionFeatureCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReplaceShapeCollectionFeatureCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | [**optional.Interface of map[string]map[string]interface{}**](map[string]interface{}.md)|  | 

### Return type

**map[string]map[string]interface{}**

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShapeCollection

> ShapeCollectionResponse UpdateShapeCollection(ctx, shapeCollectionId, optional)

Edit a ShapeCollection

Edit a ShapeCollection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shapeCollectionId** | **int64**| The ID of the ShapeCollection | 
 **optional** | ***UpdateShapeCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateShapeCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editShapeCollectionRequest** | [**optional.Interface of EditShapeCollectionRequest**](EditShapeCollectionRequest.md)|  | 

### Return type

[**ShapeCollectionResponse**](ShapeCollectionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShapeCollectionFeature

> map[string]map[string]interface{} UpdateShapeCollectionFeature(ctx, shapeCollectionId, shapeId, optional)

Update Shape GeoJSON Feature

Update Shape GeoJSON Feature

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shapeCollectionId** | **int64**| The ID of the ShapeCollection | 
**shapeId** | **int64**| The ID of the Shape | 
 **optional** | ***UpdateShapeCollectionFeatureOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateShapeCollectionFeatureOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **geoJsonFeature** | [**optional.Interface of GeoJsonFeature**](GeoJsonFeature.md)|  | 

### Return type

**map[string]map[string]interface{}**

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

