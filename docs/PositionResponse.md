# PositionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the Position | 
**TrackerId** | **int64** | The ID of the Tracker the position belongs to | 
**Timestamp** | **int64** | The UNIX timestamp of the position. If omitted, server timestamp will be used. | 
**Lat** | **float32** | Latitude of the position | 
**Lng** | **float32** | Longitude of the position | 
**Alt** | **float32** | Altitude of the position | 
**Heading** | **float32** | Heading in degrees | 
**Speed** | **float32** | Speed in knots | 
**Precision** | **float32** | Normalized precision between 0..1 | [default to 1]
**Payload** | **string** | Arbritary payload connected to the position. Max size of 10 KB. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


