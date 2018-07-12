# \MethodDELETEApi

All URIs are relative to *http://localhost/api/3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConnections**](MethodDELETEApi.md#DeleteConnections) | **Delete** /connections | Reset client connections statistics
[**DeleteHttpCacheZoneStat**](MethodDELETEApi.md#DeleteHttpCacheZoneStat) | **Delete** /http/caches/{httpCacheZoneName} | Reset cache statistics
[**DeleteHttpKeyvalZoneData**](MethodDELETEApi.md#DeleteHttpKeyvalZoneData) | **Delete** /http/keyvals/{httpKeyvalZoneName} | Empty the HTTP keyval zone
[**DeleteHttpRequests**](MethodDELETEApi.md#DeleteHttpRequests) | **Delete** /http/requests | Reset HTTP requests statistics
[**DeleteHttpServerZoneStat**](MethodDELETEApi.md#DeleteHttpServerZoneStat) | **Delete** /http/server_zones/{httpServerZoneName} | Reset statistics for an HTTP server zone
[**DeleteHttpUpstreamServer**](MethodDELETEApi.md#DeleteHttpUpstreamServer) | **Delete** /http/upstreams/{httpUpstreamName}/servers/{httpUpstreamServerId} | Remove a server from an HTTP upstream server group
[**DeleteHttpUpstreamStat**](MethodDELETEApi.md#DeleteHttpUpstreamStat) | **Delete** /http/upstreams/{httpUpstreamName}/ | Reset statistics of an HTTP upstream server group
[**DeleteProcesses**](MethodDELETEApi.md#DeleteProcesses) | **Delete** /processes | Reset nginx processes statistics
[**DeleteSlabZoneStats**](MethodDELETEApi.md#DeleteSlabZoneStats) | **Delete** /slabs/{slabZoneName} | Reset slab statistics
[**DeleteSslStat**](MethodDELETEApi.md#DeleteSslStat) | **Delete** /ssl | Reset SSL statistics
[**DeleteStreamKeyvalZoneData**](MethodDELETEApi.md#DeleteStreamKeyvalZoneData) | **Delete** /stream/keyvals/{streamKeyvalZoneName} | Empty the stream keyval zone
[**DeleteStreamServerZoneStat**](MethodDELETEApi.md#DeleteStreamServerZoneStat) | **Delete** /stream/server_zones/{streamServerZoneName} | Reset statistics for a stream server zone
[**DeleteStreamUpstreamServer**](MethodDELETEApi.md#DeleteStreamUpstreamServer) | **Delete** /stream/upstreams/{streamUpstreamName}/servers/{streamUpstreamServerId} | Remove a server from a stream upstream server group
[**DeleteStreamUpstreamStat**](MethodDELETEApi.md#DeleteStreamUpstreamStat) | **Delete** /stream/upstreams/{streamUpstreamName}/ | Reset statistics of a stream upstream server group


# **DeleteConnections**
> DeleteConnections(ctx, )
Reset client connections statistics

Resets statistics of accepted and dropped client connections.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHttpCacheZoneStat**
> DeleteHttpCacheZoneStat(ctx, httpCacheZoneName)
Reset cache statistics

Resets statistics of cache hits/misses in a particular cache zone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **httpCacheZoneName** | **string**| The name of the cache zone. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **DeleteHttpRequests**
> DeleteHttpRequests(ctx, )
Reset HTTP requests statistics

Resets the number of total client HTTP requests.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **DeleteProcesses**
> DeleteProcesses(ctx, )
Reset nginx processes statistics

Resets counters of abnormally terminated and respawned child processes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSlabZoneStats**
> DeleteSlabZoneStats(ctx, slabZoneName)
Reset slab statistics

Resets the “<literal>reqs</literal>” and “<literal>fails</literal>” metrics for each memory slot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **slabZoneName** | **string**| The name of the shared memory zone with slab allocator. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSslStat**
> DeleteSslStat(ctx, )
Reset SSL statistics

Resets counters of SSL handshakes and session reuses.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

