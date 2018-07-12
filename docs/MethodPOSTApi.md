# \MethodPOSTApi

All URIs are relative to *http://localhost/api/3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostHttpKeyvalZoneData**](MethodPOSTApi.md#PostHttpKeyvalZoneData) | **Post** /http/keyvals/{httpKeyvalZoneName} | Add a key-value pair to the HTTP keyval zone
[**PostHttpUpstreamServer**](MethodPOSTApi.md#PostHttpUpstreamServer) | **Post** /http/upstreams/{httpUpstreamName}/servers/ | Add a server to an HTTP upstream server group
[**PostStreamKeyvalZoneData**](MethodPOSTApi.md#PostStreamKeyvalZoneData) | **Post** /stream/keyvals/{streamKeyvalZoneName} | Add a key-value pair to the stream keyval zone
[**PostStreamUpstreamServer**](MethodPOSTApi.md#PostStreamUpstreamServer) | **Post** /stream/upstreams/{streamUpstreamName}/servers/ | Add a server to a stream upstream server group


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

# **PostHttpUpstreamServer**
> NginxHttpUpstreamConfServer PostHttpUpstreamServer(ctx, httpUpstreamName, postHttpUpstreamServer)
Add a server to an HTTP upstream server group

Adds a new server to an HTTP upstream server group. Server parameters are specified in the JSON format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **httpUpstreamName** | **string**| The name of an upstream server group. | 
  **postHttpUpstreamServer** | [**NginxHttpUpstreamConfServer**](NginxHttpUpstreamConfServer.md)| Address of a new server and other optional parameters in the JSON format. The “*ID*”, “*backup*”, and “*service*” parameters cannot be changed. | 

### Return type

[**NginxHttpUpstreamConfServer**](NginxHTTPUpstreamConfServer.md)

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

# **PostStreamUpstreamServer**
> NginxStreamUpstreamConfServer PostStreamUpstreamServer(ctx, streamUpstreamName, postStreamUpstreamServer)
Add a server to a stream upstream server group

Adds a new server to a stream upstream server group. Server parameters are specified in the JSON format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **streamUpstreamName** | **string**| The name of an upstream server group. | 
  **postStreamUpstreamServer** | [**NginxStreamUpstreamConfServer**](NginxStreamUpstreamConfServer.md)| Address of a new server and other optional parameters in the JSON format. The “*ID*”, “*backup*”, and “*service*” parameters cannot be changed. | 

### Return type

[**NginxStreamUpstreamConfServer**](NginxStreamUpstreamConfServer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

