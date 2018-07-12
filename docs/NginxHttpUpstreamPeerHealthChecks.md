# NginxHttpUpstreamPeerHealthChecks

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checks** | **int32** | The total number of &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_hc_module.html#health_check\&quot;&gt;health check&lt;/a&gt; requests made. | [optional] [default to null]
**Fails** | **int32** | The number of failed health checks. | [optional] [default to null]
**Unhealthy** | **int32** | How many times the server became unhealthy (state “&lt;code&gt;unhealthy&lt;/code&gt;”). | [optional] [default to null]
**LastPassed** | **bool** | Boolean indicating if the last health check request was successful and passed &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_hc_module.html#match\&quot;&gt;tests&lt;/a&gt;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


