# NginxHttpUpstreamPeer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the server. | [optional] [default to null]
**Server** | **string** | An  &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#server\&quot;&gt;address&lt;/a&gt; of the server. | [optional] [default to null]
**Service** | **string** | The &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#service\&quot;&gt;service&lt;/a&gt; parameter value of the &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#server\&quot;&gt;server&lt;/a&gt; directive. | [optional] [default to null]
**Name** | **string** | The name of the server specified in the &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#server\&quot;&gt;server&lt;/a&gt; directive. | [optional] [default to null]
**Backup** | **bool** | A boolean value indicating whether the server is a &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#backup\&quot;&gt;backup&lt;/a&gt; server. | [optional] [default to null]
**Weight** | **int32** | &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#weight\&quot;&gt;Weight&lt;/a&gt; of the server. | [optional] [default to null]
**State** | **string** | Current state, which may be one of “&lt;code&gt;up&lt;/code&gt;”, “&lt;code&gt;draining&lt;/code&gt;”, “&lt;code&gt;down&lt;/code&gt;”, “&lt;code&gt;unavail&lt;/code&gt;”, “&lt;code&gt;checking&lt;/code&gt;”, and “&lt;code&gt;unhealthy&lt;/code&gt;”. | [optional] [default to null]
**Active** | **int32** | The current number of active connections. | [optional] [default to null]
**MaxConns** | **int32** | The &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#max_conns\&quot;&gt;max_conns&lt;/a&gt; limit for the server. | [optional] [default to null]
**Requests** | **int32** | The total number of client requests forwarded to this server. | [optional] [default to null]
**Responses** | [***NginxHttpUpstreamPeerResponses**](NginxHTTPUpstreamPeer_responses.md) |  | [optional] [default to null]
**Sent** | **int32** | The total number of bytes sent to this server. | [optional] [default to null]
**Received** | **int32** | The total number of bytes received from this server. | [optional] [default to null]
**Fails** | **int32** | The total number of unsuccessful attempts to communicate with the server. | [optional] [default to null]
**Unavail** | **int32** | How many times the server became unavailable for client requests (state “&lt;code&gt;unavail&lt;/code&gt;”) due to the number of unsuccessful attempts reaching the &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#max_fails\&quot;&gt;max_fails&lt;/a&gt; threshold. | [optional] [default to null]
**HealthChecks** | [***NginxHttpUpstreamPeerHealthChecks**](NginxHTTPUpstreamPeer_health_checks.md) |  | [optional] [default to null]
**Downtime** | **int32** | Total time the server was in the “&lt;code&gt;unavail&lt;/code&gt;”, “&lt;code&gt;checking&lt;/code&gt;”, and “&lt;code&gt;unhealthy&lt;/code&gt;” states. | [optional] [default to null]
**Downstart** | [**time.Time**](time.Time.md) | The time (in milliseconds since Epoch) when the server became “&lt;code&gt;unavail&lt;/code&gt;”, “&lt;code&gt;checking&lt;/code&gt;”, or “&lt;code&gt;unhealthy&lt;/code&gt;”. | [optional] [default to null]
**Selected** | [**time.Time**](time.Time.md) | The time (in milliseconds since Epoch) when the server was last selected to process a request. | [optional] [default to null]
**HeaderTime** | **int32** | The average time to get the &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#var_upstream_header_time\&quot;&gt;response header&lt;/a&gt; from the server. | [optional] [default to null]
**ResponseTime** | **int32** | The average time to get the &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#var_upstream_response_time\&quot;&gt;full response&lt;/a&gt; from the server. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


