# \HTTPServerZonesApi

All URIs are relative to *http://localhost/api/3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteHttpServerZoneStat**](HTTPServerZonesApi.md#DeleteHttpServerZoneStat) | **Delete** /http/server_zones/{httpServerZoneName} | Reset statistics for an HTTP server zone
[**GetHttpServerZone**](HTTPServerZonesApi.md#GetHttpServerZone) | **Get** /http/server_zones/{httpServerZoneName} | Return status of an HTTP server zone
[**GetHttpServerZones**](HTTPServerZonesApi.md#GetHttpServerZones) | **Get** /http/server_zones/ | Return status of all HTTP server zones


# **DeleteHttpServerZoneStat**
> DeleteHttpServerZoneStat(ctx, httpServerZoneName)
Reset statistics for an HTTP server zone

Resets statistics of accepted and discarded requests, responses, received and sent bytes in a particular HTTP server zone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **httpServerZoneName** | **string**| The name of an HTTP server zone. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHttpServerZone**
> NginxHttpServerZone GetHttpServerZone(ctx, httpServerZoneName, optional)
Return status of an HTTP server zone

Returns status of a particular HTTP server zone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **httpServerZoneName** | **string**| The name of an HTTP server zone. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **httpServerZoneName** | **string**| The name of an HTTP server zone. | 
 **fields** | **string**| Limits which fields of the server zone will be output. | 

### Return type

[**NginxHttpServerZone**](NginxHTTPServerZone.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHttpServerZones**
> NginxHttpServerZonesMap GetHttpServerZones(ctx, optional)
Return status of all HTTP server zones

Returns status information for each HTTP [server zone](https://nginx.org/en/docs/http/ngx_http_status_module.html#status_zone).

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

[**NginxHttpServerZonesMap**](NginxHTTPServerZonesMap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

