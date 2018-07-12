# NginxHttpCacheBypass

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Responses** | **int32** | The total number of responses not looked up in the cache due to the &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_cache_bypass\&quot;&gt;proxy_cache_bypass&lt;/a&gt; and other “&lt;code&gt;*_cache_bypass&lt;/code&gt;” directives. | [optional] [default to null]
**Bytes** | **int32** | The total number of bytes read from the proxied server. | [optional] [default to null]
**ResponsesWritten** | **int32** | The total number of responses written to the cache. | [optional] [default to null]
**BytesWritten** | **int32** | The total number of bytes written to the cache. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


