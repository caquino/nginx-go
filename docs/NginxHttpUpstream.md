# NginxHttpUpstream

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Peers** | [***NginxHttpUpstreamPeerMap**](NginxHTTPUpstreamPeerMap.md) |  | [optional] [default to null]
**Keepalive** | **int32** | The current number of idle &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#keepalive\&quot;&gt;keepalive&lt;/a&gt; connections. | [optional] [default to null]
**Zombies** | **int32** | The current number of servers removed from the group but still processing active client requests. | [optional] [default to null]
**Zone** | **string** | The name of the shared memory &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#zone\&quot;&gt;zone&lt;/a&gt; that keeps the group’s configuration and run-time state. | [optional] [default to null]
**Queue** | [***NginxHttpUpstreamQueue**](NginxHTTPUpstream_queue.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


