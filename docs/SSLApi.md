# \SSLApi

All URIs are relative to *http://localhost/api/3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSslStat**](SSLApi.md#DeleteSslStat) | **Delete** /ssl | Reset SSL statistics
[**GetSsl**](SSLApi.md#GetSsl) | **Get** /ssl | Return SSL statistics


# **DeleteSslStat**
> DeleteSslStat(ctx, )
Reset SSL statistics

Resets counters of SSL handshakes and session reuses.

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

