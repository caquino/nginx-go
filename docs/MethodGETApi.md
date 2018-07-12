# \MethodGETApi

All URIs are relative to *http://localhost/api/3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAPIEndpoints**](MethodGETApi.md#GetAPIEndpoints) | **Get** / | Return list of root endpoints
[**GetConnections**](MethodGETApi.md#GetConnections) | **Get** /connections | Return client connections statistics
[**GetHttp**](MethodGETApi.md#GetHttp) | **Get** /http/ | Return list of HTTP-related endpoints
[**GetHttpCacheZone**](MethodGETApi.md#GetHttpCacheZone) | **Get** /http/caches/{httpCacheZoneName} | Return status of a cache
[**GetHttpCaches**](MethodGETApi.md#GetHttpCaches) | **Get** /http/caches/ | Return status of all caches
[**GetHttpKeyvalZoneKeysValues**](MethodGETApi.md#GetHttpKeyvalZoneKeysValues) | **Get** /http/keyvals/{httpKeyvalZoneName} | Return key-value pairs from an HTTP keyval zone
[**GetHttpKeyvalZones**](MethodGETApi.md#GetHttpKeyvalZones) | **Get** /http/keyvals/ | Return key-value pairs from all HTTP keyval zones
[**GetHttpRequests**](MethodGETApi.md#GetHttpRequests) | **Get** /http/requests | Return HTTP requests statistics
[**GetHttpServerZone**](MethodGETApi.md#GetHttpServerZone) | **Get** /http/server_zones/{httpServerZoneName} | Return status of an HTTP server zone
[**GetHttpServerZones**](MethodGETApi.md#GetHttpServerZones) | **Get** /http/server_zones/ | Return status of all HTTP server zones
[**GetHttpUpstreamName**](MethodGETApi.md#GetHttpUpstreamName) | **Get** /http/upstreams/{httpUpstreamName}/ | Return status of an HTTP upstream server group
[**GetHttpUpstreamPeer**](MethodGETApi.md#GetHttpUpstreamPeer) | **Get** /http/upstreams/{httpUpstreamName}/servers/{httpUpstreamServerId} | Return configuration of a server in an HTTP upstream server group
[**GetHttpUpstreamServers**](MethodGETApi.md#GetHttpUpstreamServers) | **Get** /http/upstreams/{httpUpstreamName}/servers/ | Return configuration of all servers in an HTTP upstream server group
[**GetHttpUpstreams**](MethodGETApi.md#GetHttpUpstreams) | **Get** /http/upstreams/ | Return status of all HTTP upstream server groups
[**GetNginx**](MethodGETApi.md#GetNginx) | **Get** /nginx | Return status of nginx running instance
[**GetProcesses**](MethodGETApi.md#GetProcesses) | **Get** /processes | Return nginx processes status
[**GetSlabZone**](MethodGETApi.md#GetSlabZone) | **Get** /slabs/{slabZoneName} | Return status of a slab
[**GetSlabs**](MethodGETApi.md#GetSlabs) | **Get** /slabs/ | Return status of all slabs
[**GetSsl**](MethodGETApi.md#GetSsl) | **Get** /ssl | Return SSL statistics
[**GetStreamKeyvalZoneKeysValues**](MethodGETApi.md#GetStreamKeyvalZoneKeysValues) | **Get** /stream/keyvals/{streamKeyvalZoneName} | Return key-value pairs from a stream keyval zone
[**GetStreamKeyvalZones**](MethodGETApi.md#GetStreamKeyvalZones) | **Get** /stream/keyvals/ | Return key-value pairs from all stream keyval zones
[**GetStreamServerZone**](MethodGETApi.md#GetStreamServerZone) | **Get** /stream/server_zones/{streamServerZoneName} | Return status of a stream server zone
[**GetStreamServerZones**](MethodGETApi.md#GetStreamServerZones) | **Get** /stream/server_zones/ | Return status of all stream server zones
[**GetStreamUpstream**](MethodGETApi.md#GetStreamUpstream) | **Get** /stream/upstreams/{streamUpstreamName}/ | Return status of a stream upstream server group
[**GetStreamUpstreamServer**](MethodGETApi.md#GetStreamUpstreamServer) | **Get** /stream/upstreams/{streamUpstreamName}/servers/{streamUpstreamServerId} | Return configuration of a server in a stream upstream server group
[**GetStreamUpstreamServers**](MethodGETApi.md#GetStreamUpstreamServers) | **Get** /stream/upstreams/{streamUpstreamName}/servers/ | Return configuration of all servers in a stream upstream server group
[**GetStreamUpstreams**](MethodGETApi.md#GetStreamUpstreams) | **Get** /stream/upstreams/ | Return status of all stream upstream server groups
[**GetStreamZoneSync**](MethodGETApi.md#GetStreamZoneSync) | **Get** /stream/zone_sync/ | Return sync status of a node


# **GetAPIEndpoints**
> ArrayOfStrings GetAPIEndpoints(ctx, )
Return list of root endpoints

Returns a list of root endpoints.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ArrayOfStrings**](ArrayOfStrings.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnections**
> NginxConnections GetConnections(ctx, optional)
Return client connections statistics

Returns statistics of client connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string**| Limits which fields of the connections statistics will be output. | 

### Return type

[**NginxConnections**](NginxConnections.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHttp**
> ArrayOfStrings GetHttp(ctx, )
Return list of HTTP-related endpoints

Returns a list of first level HTTP endpoints.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ArrayOfStrings**](ArrayOfStrings.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHttpCacheZone**
> NginxHttpCache GetHttpCacheZone(ctx, httpCacheZoneName, optional)
Return status of a cache

Returns status of a particular cache.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **httpCacheZoneName** | **string**| The name of the cache zone. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **httpCacheZoneName** | **string**| The name of the cache zone. | 
 **fields** | **string**| Limits which fields of the cache zone will be output. | 

### Return type

[**NginxHttpCache**](NginxHTTPCache.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHttpCaches**
> NginxHttpCachesMap GetHttpCaches(ctx, optional)
Return status of all caches

Returns status of each cache configured by [proxy_cache_path](https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_cache_path) and other “<literal>*_cache_path</literal>” directives.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string**| Limits which fields of cache zones will be output. If the “&lt;literal&gt;fields&lt;/literal&gt;” value is empty, then only names of cache zones are output. | 

### Return type

[**NginxHttpCachesMap**](NginxHTTPCachesMap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

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

# **GetHttpRequests**
> NginxHttpRequests GetHttpRequests(ctx, optional)
Return HTTP requests statistics

Returns status of client HTTP requests.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string**| Limits which fields of client HTTP requests statistics will be output. | 

### Return type

[**NginxHttpRequests**](NginxHTTPRequests.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

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

# **GetNginx**
> NginxObject GetNginx(ctx, optional)
Return status of nginx running instance

Returns nginx version, build name, address, number of configuration reloads, IDs of master and worker processes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string**| Limits which fields of nginx running instance will be output. | 

### Return type

[**NginxObject**](NginxObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProcesses**
> NginxProcesses GetProcesses(ctx, )
Return nginx processes status

Returns the number of abnormally terminated and respawned child processes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NginxProcesses**](NginxProcesses.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSlabZone**
> NginxSlabZone GetSlabZone(ctx, slabZoneName, optional)
Return status of a slab

Returns status of slabs for a particular shared memory zone with slab allocator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **slabZoneName** | **string**| The name of the shared memory zone with slab allocator. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slabZoneName** | **string**| The name of the shared memory zone with slab allocator. | 
 **fields** | **string**| Limits which fields of the slab zone will be output. | 

### Return type

[**NginxSlabZone**](NginxSlabZone.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSlabs**
> NginxSlabZoneMap GetSlabs(ctx, optional)
Return status of all slabs

Returns status of slabs for each shared memory zone with slab allocator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string**| Limits which fields of slab zones will be output. If the “&lt;literal&gt;fields&lt;/literal&gt;” value is empty, then only zone names are output. | 

### Return type

[**NginxSlabZoneMap**](NginxSlabZoneMap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSsl**
> NginxSslObject GetSsl(ctx, optional)
Return SSL statistics

Returns SSL statistics.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string**| Limits which fields of SSL statistics will be output. | 

### Return type

[**NginxSslObject**](NginxSSLObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

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

# **GetStreamZoneSync**
> NginxStreamZoneSync GetStreamZoneSync(ctx, )
Return sync status of a node

Returns synchronization status of a cluster node.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NginxStreamZoneSync**](NginxStreamZoneSync.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

