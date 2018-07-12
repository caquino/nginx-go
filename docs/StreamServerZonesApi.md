# \StreamServerZonesApi

All URIs are relative to *http://localhost/api/3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteStreamServerZoneStat**](StreamServerZonesApi.md#DeleteStreamServerZoneStat) | **Delete** /stream/server_zones/{streamServerZoneName} | Reset statistics for a stream server zone
[**GetStreamServerZone**](StreamServerZonesApi.md#GetStreamServerZone) | **Get** /stream/server_zones/{streamServerZoneName} | Return status of a stream server zone
[**GetStreamServerZones**](StreamServerZonesApi.md#GetStreamServerZones) | **Get** /stream/server_zones/ | Return status of all stream server zones


# **DeleteStreamServerZoneStat**
> DeleteStreamServerZoneStat(ctx, streamServerZoneName)
Reset statistics for a stream server zone

Resets statistics of accepted and discarded connections, sessions, received and sent bytes in a particular stream server zone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **streamServerZoneName** | **string**| The name of a stream server zone. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStreamServerZone**
> NginxStreamServerZone GetStreamServerZone(ctx, streamServerZoneName, optional)
Return status of a stream server zone

Returns status of a particular stream server zone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **streamServerZoneName** | **string**| The name of a stream server zone. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamServerZoneName** | **string**| The name of a stream server zone. | 
 **fields** | **string**| Limits which fields of the server zone will be output. | 

### Return type

[**NginxStreamServerZone**](NginxStreamServerZone.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStreamServerZones**
> NginxStreamServerZonesMap GetStreamServerZones(ctx, optional)
Return status of all stream server zones

Returns status information for each stream [server zone](https://nginx.org/en/docs/http/ngx_http_status_module.html#status_zone).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string**| Limits which fields of server zones will be output. If the “&lt;literal&gt;fields&lt;/literal&gt;” value is empty, then only server zone names are output. | 

### Return type

[**NginxStreamServerZonesMap**](NginxStreamServerZonesMap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

