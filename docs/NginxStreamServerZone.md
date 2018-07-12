# NginxStreamServerZone

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Processing** | **int32** | The number of client connections that are currently being processed. | [optional] [default to null]
**Connections** | **int32** | The total number of connections accepted from clients. | [optional] [default to null]
**Sessions** | [***NginxStreamServerZoneSessions**](NginxStreamServerZone_sessions.md) |  | [optional] [default to null]
**Discarded** | **int32** | The total number of connections completed without creating a session. | [optional] [default to null]
**Received** | **int32** | The total number of bytes received from clients. | [optional] [default to null]
**Sent** | **int32** | The total number of bytes sent to clients. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


