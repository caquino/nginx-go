# \StreamUpstreamsApi

All URIs are relative to *http://localhost/api/3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteStreamUpstreamServer**](StreamUpstreamsApi.md#DeleteStreamUpstreamServer) | **Delete** /stream/upstreams/{streamUpstreamName}/servers/{streamUpstreamServerId} | Remove a server from a stream upstream server group
[**DeleteStreamUpstreamStat**](StreamUpstreamsApi.md#DeleteStreamUpstreamStat) | **Delete** /stream/upstreams/{streamUpstreamName}/ | Reset statistics of a stream upstream server group
[**GetStreamUpstream**](StreamUpstreamsApi.md#GetStreamUpstream) | **Get** /stream/upstreams/{streamUpstreamName}/ | Return status of a stream upstream server group
[**GetStreamUpstreamServer**](StreamUpstreamsApi.md#GetStreamUpstreamServer) | **Get** /stream/upstreams/{streamUpstreamName}/servers/{streamUpstreamServerId} | Return configuration of a server in a stream upstream server group
[**GetStreamUpstreamServers**](StreamUpstreamsApi.md#GetStreamUpstreamServers) | **Get** /stream/upstreams/{streamUpstreamName}/servers/ | Return configuration of all servers in a stream upstream server group
[**GetStreamUpstreams**](StreamUpstreamsApi.md#GetStreamUpstreams) | **Get** /stream/upstreams/ | Return status of all stream upstream server groups
[**PatchStreamUpstreamServer**](StreamUpstreamsApi.md#PatchStreamUpstreamServer) | **Patch** /stream/upstreams/{streamUpstreamName}/servers/{streamUpstreamServerId} | Modify a server in a stream upstream server group
[**PostStreamUpstreamServer**](StreamUpstreamsApi.md#PostStreamUpstreamServer) | **Post** /stream/upstreams/{streamUpstreamName}/servers/ | Add a server to a stream upstream server group


# **DeleteStreamUpstreamServer**
> NginxStreamUpstreamConfServerMap DeleteStreamUpstreamServer(ctx, streamUpstreamName, streamUpstreamServerId)
Remove a server from a stream upstream server group

Removes a server from a stream server group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **streamUpstreamName** | **string**| The name of the upstream server group. | 
  **streamUpstreamServerId** | **string**| The ID of the server. | 

### Return type

[**NginxStreamUpstreamConfServerMap**](NginxStreamUpstreamConfServerMap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStreamUpstreamStat**
> DeleteStreamUpstreamStat(ctx, streamUpstreamName)
Reset statistics of a stream upstream server group

Resets the statistics for each upstream server in an upstream server group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **streamUpstreamName** | **string**| The name of a stream upstream server group. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStreamUpstream**
> NginxStreamUpstream GetStreamUpstream(ctx, streamUpstreamName, optional)
Return status of a stream upstream server group

Returns status of a particular stream upstream server group and its servers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **streamUpstreamName** | **string**| The name of a stream upstream server group. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamUpstreamName** | **string**| The name of a stream upstream server group. | 
 **fields** | **string**| Limits which fields of the upstream server group will be output. | 

### Return type

[**NginxStreamUpstream**](NginxStreamUpstream.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStreamUpstreamServer**
> NginxStreamUpstreamConfServer GetStreamUpstreamServer(ctx, streamUpstreamName, streamUpstreamServerId)
Return configuration of a server in a stream upstream server group

Returns configuration of a particular server in the stream upstream server group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **streamUpstreamName** | **string**| The name of the upstream server group. | 
  **streamUpstreamServerId** | **string**| The ID of the server. | 

### Return type

[**NginxStreamUpstreamConfServer**](NginxStreamUpstreamConfServer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStreamUpstreamServers**
> NginxStreamUpstreamConfServerMap GetStreamUpstreamServers(ctx, streamUpstreamName)
Return configuration of all servers in a stream upstream server group

Returns configuration of each server in a particular stream upstream server group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **streamUpstreamName** | **string**| The name of an upstream server group. | 

### Return type

[**NginxStreamUpstreamConfServerMap**](NginxStreamUpstreamConfServerMap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStreamUpstreams**
> NginxStreamUpstreamMap GetStreamUpstreams(ctx, optional)
Return status of all stream upstream server groups

Returns status of each stream upstream server group and its servers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string**| Limits which fields of upstream server groups will be output. If the “&lt;literal&gt;fields&lt;/literal&gt;” value is empty, only names of upstreams are output. | 

### Return type

[**NginxStreamUpstreamMap**](NginxStreamUpstreamMap.md)

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

