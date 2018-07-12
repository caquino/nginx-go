# \ConnectionsApi

All URIs are relative to *http://localhost/api/3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConnections**](ConnectionsApi.md#DeleteConnections) | **Delete** /connections | Reset client connections statistics
[**GetConnections**](ConnectionsApi.md#GetConnections) | **Get** /connections | Return client connections statistics


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

