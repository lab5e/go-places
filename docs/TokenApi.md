# \TokenApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](TokenApi.md#CreateToken) | **Post** /tokens | Create a new Token
[**DeleteToken**](TokenApi.md#DeleteToken) | **Delete** /tokens/{tokenId} | Delete a Token
[**GetTokens**](TokenApi.md#GetTokens) | **Get** /tokens | Get all Tokens
[**UpdateToken**](TokenApi.md#UpdateToken) | **Put** /tokens/{tokenId} | Update a Token



## CreateToken

> TokenResponse CreateToken(ctx, optional)

Create a new Token

Create a new Token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTokenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newTokenRequest** | [**optional.Interface of NewTokenRequest**](NewTokenRequest.md)|  | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteToken

> DeleteToken(ctx, tokenId)

Delete a Token

Delete a Token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string**| The Token value | 

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


## GetTokens

> []TokenResponse GetTokens(ctx, optional)

Get all Tokens

Get all Tokens

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTokensOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| Limit number of responses (default 255, max 2500) | [default to 255]
 **offset** | **optional.Int64**| Offset the response by the integer given. If provided, needs to be above 0. | [default to 1]

### Return type

[**[]TokenResponse**](TokenResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateToken

> TokenResponse UpdateToken(ctx, tokenId, optional)

Update a Token

Update a Token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string**| The Token value | 
 **optional** | ***UpdateTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTokenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editTokenRequest** | [**optional.Interface of EditTokenRequest**](EditTokenRequest.md)|  | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

