# \HTTPUpstreamsApi

All URIs are relative to *http://localhost/api/3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteHttpUpstreamServer**](HTTPUpstreamsApi.md#DeleteHttpUpstreamServer) | **Delete** /http/upstreams/{httpUpstreamName}/servers/{httpUpstreamServerId} | Remove a server from an HTTP upstream server group
[**DeleteHttpUpstreamStat**](HTTPUpstreamsApi.md#DeleteHttpUpstreamStat) | **Delete** /http/upstreams/{httpUpstreamName}/ | Reset statistics of an HTTP upstream server group
[**GetHttpUpstreamName**](HTTPUpstreamsApi.md#GetHttpUpstreamName) | **Get** /http/upstreams/{httpUpstreamName}/ | Return status of an HTTP upstream server group
[**GetHttpUpstreamPeer**](HTTPUpstreamsApi.md#GetHttpUpstreamPeer) | **Get** /http/upstreams/{httpUpstreamName}/servers/{httpUpstreamServerId} | Return configuration of a server in an HTTP upstream server group
[**GetHttpUpstreamServers**](HTTPUpstreamsApi.md#GetHttpUpstreamServers) | **Get** /http/upstreams/{httpUpstreamName}/servers/ | Return configuration of all servers in an HTTP upstream server group
[**GetHttpUpstreams**](HTTPUpstreamsApi.md#GetHttpUpstreams) | **Get** /http/upstreams/ | Return status of all HTTP upstream server groups
[**PatchHttpUpstreamPeer**](HTTPUpstreamsApi.md#PatchHttpUpstreamPeer) | **Patch** /http/upstreams/{httpUpstreamName}/servers/{httpUpstreamServerId} | Modify a server in an HTTP upstream server group
[**PostHttpUpstreamServer**](HTTPUpstreamsApi.md#PostHttpUpstreamServer) | **Post** /http/upstreams/{httpUpstreamName}/servers/ | Add a server to an HTTP upstream server group


# **DeleteHttpUpstreamServer**
> NginxHttpUpstreamConfServerMap DeleteHttpUpstreamServer(ctx, httpUpstreamName, httpUpstreamServerId)
Remove a server from an HTTP upstream server group

Removes a server from an HTTP upstream server group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **httpUpstreamName** | **string**| The name of the upstream server group. | 
  **httpUpstreamServerId** | **string**| The ID of the server. | 

### Return type

[**NginxHttpUpstreamConfServerMap**](NginxHTTPUpstreamConfServerMap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHttpUpstreamStat**
> DeleteHttpUpstreamStat(ctx, httpUpstreamName)
Reset statistics of an HTTP upstream server group

Resets the statistics for each upstream server in an upstream server group and queue statistics.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **httpUpstreamName** | **string**| The name of an HTTP upstream server group. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHttpUpstreamName**
> NginxHttpUpstream GetHttpUpstreamName(ctx, httpUpstreamName, optional)
Return status of an HTTP upstream server group

Returns status of a particular HTTP upstream server group and its servers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **httpUpstreamName** | **string**| The name of an HTTP upstream server group. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **httpUpstreamName** | **string**| The name of an HTTP upstream server group. | 
 **fields** | **string**| Limits which fields of the upstream server group will be output. | 

### Return type

[**NginxHttpUpstream**](NginxHTTPUpstream.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHttpUpstreamPeer**
> NginxHttpUpstreamConfServer GetHttpUpstreamPeer(ctx, httpUpstreamName, httpUpstreamServerId)
Return configuration of a server in an HTTP upstream server group

Returns configuration of a particular server in the HTTP upstream server group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **httpUpstreamName** | **string**| The name of the upstream server group. | 
  **httpUpstreamServerId** | **string**| The ID of the server. | 

### Return type

[**NginxHttpUpstreamConfServer**](NginxHTTPUpstreamConfServer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHttpUpstreamServers**
> NginxHttpUpstreamConfServerMap GetHttpUpstreamServers(ctx, httpUpstreamName)
Return configuration of all servers in an HTTP upstream server group

Returns configuration of each server in a particular HTTP upstream server group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **httpUpstreamName** | **string**| The name of an upstream server group. | 

### Return type

[**NginxHttpUpstreamConfServerMap**](NginxHTTPUpstreamConfServerMap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHttpUpstreams**
> NginxHttpUpstreamMap GetHttpUpstreams(ctx, optional)
Return status of all HTTP upstream server groups

Returns status of each HTTP upstream server group and its servers.

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

[**NginxHttpUpstreamMap**](NginxHTTPUpstreamMap.md)

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

