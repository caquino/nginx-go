# \StreamKeyvalsApi

All URIs are relative to *http://localhost/api/3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteStreamKeyvalZoneData**](StreamKeyvalsApi.md#DeleteStreamKeyvalZoneData) | **Delete** /stream/keyvals/{streamKeyvalZoneName} | Empty the stream keyval zone
[**GetStreamKeyvalZoneKeysValues**](StreamKeyvalsApi.md#GetStreamKeyvalZoneKeysValues) | **Get** /stream/keyvals/{streamKeyvalZoneName} | Return key-value pairs from a stream keyval zone
[**GetStreamKeyvalZones**](StreamKeyvalsApi.md#GetStreamKeyvalZones) | **Get** /stream/keyvals/ | Return key-value pairs from all stream keyval zones
[**PatchStreamKeyvalZoneKeyValue**](StreamKeyvalsApi.md#PatchStreamKeyvalZoneKeyValue) | **Patch** /stream/keyvals/{streamKeyvalZoneName} | Modify a key-value or delete a key
[**PostStreamKeyvalZoneData**](StreamKeyvalsApi.md#PostStreamKeyvalZoneData) | **Post** /stream/keyvals/{streamKeyvalZoneName} | Add a key-value pair to the stream keyval zone


# **DeleteStreamKeyvalZoneData**
> DeleteStreamKeyvalZoneData(ctx, streamKeyvalZoneName)
Empty the stream keyval zone

Deletes all key-value pairs from the stream keyval shared memory [zone](https://nginx.org/en/docs/stream/ngx_stream_keyval_module.html#keyval_zone).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **streamKeyvalZoneName** | **string**| The name of a stream keyval shared memory zone. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStreamKeyvalZoneKeysValues**
> NginxStreamKeyvalZone GetStreamKeyvalZoneKeysValues(ctx, streamKeyvalZoneName, optional)
Return key-value pairs from a stream keyval zone

Returns key-value pairs stored in a particular stream keyval shared memory [zone](https://nginx.org/en/docs/stream/ngx_stream_keyval_module.html#keyval_zone).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **streamKeyvalZoneName** | **string**| The name of a stream keyval shared memory zone. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamKeyvalZoneName** | **string**| The name of a stream keyval shared memory zone. | 
 **key** | **string**| Get a particular key-value pair from the stream keyval zone. | 

### Return type

[**NginxStreamKeyvalZone**](NginxStreamKeyvalZone.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStreamKeyvalZones**
> NginxStreamKeyvalZonesMap GetStreamKeyvalZones(ctx, optional)
Return key-value pairs from all stream keyval zones

Returns key-value pairs for each stream keyval shared memory [zone](https://nginx.org/en/docs/stream/ngx_stream_keyval_module.html#keyval_zone).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string**| If the “&lt;literal&gt;fields&lt;/literal&gt;” value is empty, then only stream keyval zone names are output. | 

### Return type

[**NginxStreamKeyvalZonesMap**](NginxStreamKeyvalZonesMap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchStreamKeyvalZoneKeyValue**
> PatchStreamKeyvalZoneKeyValue(ctx, streamKeyvalZoneName, streamKeyvalZoneKeyValue)
Modify a key-value or delete a key

Changes the value of the selected key in the key-value pair or deletes a key by setting the key value to <literal>null</literal>.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **streamKeyvalZoneName** | **string**| The name of a stream keyval shared memory zone. | 
  **streamKeyvalZoneKeyValue** | [**NginxStreamKeyvalZone**](NginxStreamKeyvalZone.md)| A new value for the key is specified in the JSON format. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostStreamKeyvalZoneData**
> PostStreamKeyvalZoneData(ctx, streamKeyvalZoneName, keyValue)
Add a key-value pair to the stream keyval zone

Adds a new key-value pair to the stream keyval shared memory [zone](https://nginx.org/en/docs/stream/ngx_stream_keyval_module.html#keyval_zone). Several key-value pairs can be entered if the stream keyval shared memory zone is empty.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **streamKeyvalZoneName** | **string**| The name of a stream keyval shared memory zone. | 
  **keyValue** | [**NginxStreamKeyvalZone**](NginxStreamKeyvalZone.md)| A key-value pair is specified in the JSON format. Several key-value pairs can be entered if the stream keyval shared memory zone is empty. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

