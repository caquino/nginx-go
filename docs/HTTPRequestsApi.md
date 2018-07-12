# \HTTPRequestsApi

All URIs are relative to *http://localhost/api/3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteHttpRequests**](HTTPRequestsApi.md#DeleteHttpRequests) | **Delete** /http/requests | Reset HTTP requests statistics
[**GetHttpRequests**](HTTPRequestsApi.md#GetHttpRequests) | **Get** /http/requests | Return HTTP requests statistics


# **DeleteHttpRequests**
> DeleteHttpRequests(ctx, )
Reset HTTP requests statistics

Resets the number of total client HTTP requests.

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

# **GetHttpRequests**
> NginxHttpRequests GetHttpRequests(ctx, optional)
Return HTTP requests statistics

Returns status of client HTTP requests.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string**| Limits which fields of client HTTP requests statistics will be output. | 

### Return type

[**NginxHttpRequests**](NginxHTTPRequests.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

