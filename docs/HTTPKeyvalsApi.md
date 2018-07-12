# \HTTPKeyvalsApi

All URIs are relative to *http://localhost/api/3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteHttpKeyvalZoneData**](HTTPKeyvalsApi.md#DeleteHttpKeyvalZoneData) | **Delete** /http/keyvals/{httpKeyvalZoneName} | Empty the HTTP keyval zone
[**GetHttpKeyvalZoneKeysValues**](HTTPKeyvalsApi.md#GetHttpKeyvalZoneKeysValues) | **Get** /http/keyvals/{httpKeyvalZoneName} | Return key-value pairs from an HTTP keyval zone
[**GetHttpKeyvalZones**](HTTPKeyvalsApi.md#GetHttpKeyvalZones) | **Get** /http/keyvals/ | Return key-value pairs from all HTTP keyval zones
[**PatchHttpKeyvalZoneKeyValue**](HTTPKeyvalsApi.md#PatchHttpKeyvalZoneKeyValue) | **Patch** /http/keyvals/{httpKeyvalZoneName} | Modify a key-value or delete a key
[**PostHttpKeyvalZoneData**](HTTPKeyvalsApi.md#PostHttpKeyvalZoneData) | **Post** /http/keyvals/{httpKeyvalZoneName} | Add a key-value pair to the HTTP keyval zone


# **DeleteHttpKeyvalZoneData**
> DeleteHttpKeyvalZoneData(ctx, httpKeyvalZoneName)
Empty the HTTP keyval zone

Deletes all key-value pairs from the HTTP keyval shared memory [zone](https://nginx.org/en/docs/http/ngx_http_keyval_module.html#keyval_zone).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **httpKeyvalZoneName** | **string**| The name of an HTTP keyval shared memory zone. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHttpKeyvalZoneKeysValues**
> NginxHttpKeyvalZone GetHttpKeyvalZoneKeysValues(ctx, httpKeyvalZoneName, optional)
Return key-value pairs from an HTTP keyval zone

Returns key-value pairs stored in a particular HTTP keyval shared memory [zone](https://nginx.org/en/docs/http/ngx_http_keyval_module.html#keyval_zone).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **httpKeyvalZoneName** | **string**| The name of an HTTP keyval shared memory zone. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **httpKeyvalZoneName** | **string**| The name of an HTTP keyval shared memory zone. | 
 **key** | **string**| Get a particular key-value pair from the HTTP keyval zone. | 

### Return type

[**NginxHttpKeyvalZone**](NginxHTTPKeyvalZone.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHttpKeyvalZones**
> NginxHttpKeyvalZonesMap GetHttpKeyvalZones(ctx, optional)
Return key-value pairs from all HTTP keyval zones

Returns key-value pairs for each HTTP keyval shared memory [zone](https://nginx.org/en/docs/http/ngx_http_keyval_module.html#keyval_zone).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string**| If the “&lt;literal&gt;fields&lt;/literal&gt;” value is empty, then only HTTP keyval zone names are output. | 

### Return type

[**NginxHttpKeyvalZonesMap**](NginxHTTPKeyvalZonesMap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchHttpKeyvalZoneKeyValue**
> PatchHttpKeyvalZoneKeyValue(ctx, httpKeyvalZoneName, httpKeyvalZoneKeyValue)
Modify a key-value or delete a key

Changes the value of the selected key in the key-value pair or deletes a key by setting the key value to <literal>null</literal>.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **httpKeyvalZoneName** | **string**| The name of an HTTP keyval shared memory zone. | 
  **httpKeyvalZoneKeyValue** | [**NginxHttpKeyvalZone**](NginxHttpKeyvalZone.md)| A new value for the key is specified in the JSON format. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostHttpKeyvalZoneData**
> PostHttpKeyvalZoneData(ctx, httpKeyvalZoneName, keyValue)
Add a key-value pair to the HTTP keyval zone

Adds a new key-value pair to the HTTP keyval shared memory [zone](https://nginx.org/en/docs/http/ngx_http_keyval_module.html#keyval_zone). Several key-value pairs can be entered if the HTTP keyval shared memory zone is empty.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **httpKeyvalZoneName** | **string**| The name of an HTTP keyval shared memory zone. | 
  **keyValue** | [**NginxHttpKeyvalZone**](NginxHttpKeyvalZone.md)| A key-value pair is specified in the JSON format. Several key-value pairs can be entered if the HTTP keyval shared memory zone is empty. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

