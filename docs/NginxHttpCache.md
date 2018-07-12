# NginxHttpCache

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **int32** | The current size of the cache. | [optional] [default to null]
**MaxSize** | **int32** | The limit on the maximum size of the cache specified in the configuration. | [optional] [default to null]
**Cold** | **bool** | A boolean value indicating whether the “cache loader” process is still loading data from disk into the cache. | [optional] [default to null]
**Hit** | [***NginxHttpCacheHit**](NginxHTTPCache_hit.md) |  | [optional] [default to null]
**Stale** | [***NginxHttpCacheStale**](NginxHTTPCache_stale.md) |  | [optional] [default to null]
**Updating** | [***NginxHttpCacheUpdating**](NginxHTTPCache_updating.md) |  | [optional] [default to null]
**Revalidated** | [***NginxHttpCacheRevalidated**](NginxHTTPCache_revalidated.md) |  | [optional] [default to null]
**Miss** | [***NginxHttpCacheMiss**](NginxHTTPCache_miss.md) |  | [optional] [default to null]
**Expired** | [***NginxHttpCacheExpired**](NginxHTTPCache_expired.md) |  | [optional] [default to null]
**Bypass** | [***NginxHttpCacheBypass**](NginxHTTPCache_bypass.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


