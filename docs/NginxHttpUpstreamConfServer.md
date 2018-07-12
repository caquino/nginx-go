# NginxHttpUpstreamConfServer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the HTTP upstream server. The ID is assigned automatically and cannot be changed. | [optional] [default to null]
**Server** | **string** | Same as the &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#address\&quot;&gt;address&lt;/a&gt; parameter of the HTTP upstream server. When adding a server, it is possible to specify it as a domain name. In this case, changes of the IP addresses that correspond to a domain name will be monitored and automatically applied to the upstream configuration without the need of restarting nginx. This requires the &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_core_module.html#resolver\&quot;&gt;resolver&lt;/a&gt; directive in the “&lt;code&gt;http&lt;/code&gt;” block. See also the &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#resolve\&quot;&gt;resolve&lt;/a&gt; parameter of the HTTP upstream server. | [optional] [default to null]
**Service** | **string** | Same as the &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#service\&quot;&gt;service&lt;/a&gt; parameter of the HTTP upstream server. This parameter cannot be changed. | [optional] [default to null]
**Weight** | **int32** | Same as the &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#weight\&quot;&gt;weight&lt;/a&gt; parameter of the HTTP upstream server. | [optional] [default to null]
**MaxConns** | **int32** | Same as the &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#max_conns\&quot;&gt;max_conns&lt;/a&gt; parameter of the HTTP upstream server. | [optional] [default to null]
**MaxFails** | **int32** | Same as the &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#max_fails\&quot;&gt;max_fails&lt;/a&gt; parameter of the HTTP upstream server. | [optional] [default to null]
**FailTimeout** | **string** | Same as the &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#fail_timeout\&quot;&gt;fail_timeout&lt;/a&gt; parameter of the HTTP upstream server. | [optional] [default to null]
**SlowStart** | **string** | Same as the &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#slow_start\&quot;&gt;slow_start&lt;/a&gt; parameter of the HTTP upstream server. | [optional] [default to null]
**Route** | **string** | Same as the &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#route\&quot;&gt;route&lt;/a&gt; parameter of the HTTP upstream server. | [optional] [default to null]
**Backup** | **bool** | When &lt;code&gt;true&lt;/code&gt;, adds a &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#backup\&quot;&gt;backup&lt;/a&gt; server. This parameter cannot be changed. | [optional] [default to null]
**Down** | **bool** | Same as the &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#down\&quot;&gt;down&lt;/a&gt; parameter of the HTTP upstream server. | [optional] [default to null]
**Drain** | **bool** | Same as the &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_upstream_module.html#drain\&quot;&gt;drain&lt;/a&gt; parameter of the HTTP upstream server. | [optional] [default to null]
**Parent** | **string** | Parent server ID of the resolved server. The ID is assigned automatically and cannot be changed. | [optional] [default to null]
**Host** | **string** | Hostname of the resolved server. The hostname is assigned automatically and cannot be changed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


