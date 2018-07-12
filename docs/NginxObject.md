# NginxObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | Version of nginx. | [optional] [default to null]
**Build** | **string** | Name of nginx build. | [optional] [default to null]
**Address** | **string** | The address of the server that accepted status request. | [optional] [default to null]
**Generation** | **int32** | The total number of configuration &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/control.html#reconfiguration\&quot;&gt;reloads&lt;/a&gt;. | [optional] [default to null]
**LoadTimestamp** | [**time.Time**](time.Time.md) | Time of the last reload of configuration, in milliseconds since Epoch. | [optional] [default to null]
**Timestamp** | [**time.Time**](time.Time.md) | Current time in milliseconds since Epoch. | [optional] [default to null]
**Pid** | **int32** | The ID of the worker process that handled status request. | [optional] [default to null]
**Ppid** | **int32** | The ID of the master process that started the &lt;a href&#x3D;\&quot;https://nginx.org/en/docs/http/ngx_http_status_module.html#pid\&quot;&gt;worker process&lt;/a&gt;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


