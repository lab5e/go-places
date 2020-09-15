# NewPositionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int64** | The UNIX timestamp of the position. If omitted, server timestamp will be used. | [optional] 
**Lat** | **float32** | Latitude of the position | 
**Lng** | **float32** | Longitude of the position | 
**Alt** | **float32** | Altitude of the position | [optional] 
**Heading** | **float32** | Heading in degrees | [optional] 
**Speed** | **float32** | Speed in knots | [optional] 
**Precision** | **float32** | Normalized precision between 0..1 | [optional] [default to 1]
**Payload** | **string** | Arbritary payload connected to the position. Max size of 10 KB. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


