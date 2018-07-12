# \MethodPATCHApi

All URIs are relative to *http://localhost/api/3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PatchHttpKeyvalZoneKeyValue**](MethodPATCHApi.md#PatchHttpKeyvalZoneKeyValue) | **Patch** /http/keyvals/{httpKeyvalZoneName} | Modify a key-value or delete a key
[**PatchHttpUpstreamPeer**](MethodPATCHApi.md#PatchHttpUpstreamPeer) | **Patch** /http/upstreams/{httpUpstreamName}/servers/{httpUpstreamServerId} | Modify a server in an HTTP upstream server group
[**PatchStreamKeyvalZoneKeyValue**](MethodPATCHApi.md#PatchStreamKeyvalZoneKeyValue) | **Patch** /stream/keyvals/{streamKeyvalZoneName} | Modify a key-value or delete a key
[**PatchStreamUpstreamServer**](MethodPATCHApi.md#PatchStreamUpstreamServer) | **Patch** /stream/upstreams/{streamUpstreamName}/servers/{streamUpstreamServerId} | Modify a server in a stream upstream server group


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

# **PatchHttpUpstreamPeer**
> NginxHttpUpstreamConfServer PatchHttpUpstreamPeer(ctx, httpUpstreamName, httpUpstreamServerId, patchHttpUpstreamServer)
Modify a server in an HTTP upstream server group

Modifies settings of a particular server in an HTTP upstream server group. Server parameters are specified in the JSON format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **httpUpstreamName** | **string**| The name of the upstream server group. | 
  **httpUpstreamServerId** | **string**| The ID of the server. | 
  **patchHttpUpstreamServer** | [**NginxHttpUpstreamConfServer**](NginxHttpUpstreamConfServer.md)| Server parameters, specified in the JSON format. The “*ID*”, “*backup*”, and “*service*” parameters cannot be changed. | 

### Return type

[**NginxHttpUpstreamConfServer**](NginxHTTPUpstreamConfServer.md)

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

# **PatchStreamUpstreamServer**
> NginxStreamUpstreamConfServer PatchStreamUpstreamServer(ctx, streamUpstreamName, streamUpstreamServerId, patchStreamUpstreamServer)
Modify a server in a stream upstream server group

Modifies settings of a particular server in a stream upstream server group. Server parameters are specified in the JSON format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **streamUpstreamName** | **string**| The name of the upstream server group. | 
  **streamUpstreamServerId** | **string**| The ID of the server. | 
  **patchStreamUpstreamServer** | [**NginxStreamUpstreamConfServer**](NginxStreamUpstreamConfServer.md)| Server parameters, specified in the JSON format. The “*ID*”, “*backup*”, and “*service*” parameters cannot be changed. | 

### Return type

[**NginxStreamUpstreamConfServer**](NginxStreamUpstreamConfServer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

