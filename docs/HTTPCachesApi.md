# \HTTPCachesApi

All URIs are relative to *http://localhost/api/3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteHttpCacheZoneStat**](HTTPCachesApi.md#DeleteHttpCacheZoneStat) | **Delete** /http/caches/{httpCacheZoneName} | Reset cache statistics
[**GetHttpCacheZone**](HTTPCachesApi.md#GetHttpCacheZone) | **Get** /http/caches/{httpCacheZoneName} | Return status of a cache
[**GetHttpCaches**](HTTPCachesApi.md#GetHttpCaches) | **Get** /http/caches/ | Return status of all caches


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

