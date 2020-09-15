# TokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | The Token that can be used for programmatical access | 
**Created** | **int64** | The UNIX timestamp of the creation of the Token | 
**Resource** | **string** | Limits the allowed paths of the token towards the rest of the API. Only the direct path or subpath of this parameter will be allowed | 
**PermWrite** | **bool** | Determines whether the token is allowed to do POST, PUT and DELETE methods | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


