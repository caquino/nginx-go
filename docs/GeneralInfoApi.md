# \GeneralInfoApi

All URIs are relative to *http://localhost/api/3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAPIEndpoints**](GeneralInfoApi.md#GetAPIEndpoints) | **Get** / | Return list of root endpoints
[**GetNginx**](GeneralInfoApi.md#GetNginx) | **Get** /nginx | Return status of nginx running instance


# **GetAPIEndpoints**
> ArrayOfStrings GetAPIEndpoints(ctx, )
Return list of root endpoints

Returns a list of root endpoints.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ArrayOfStrings**](ArrayOfStrings.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNginx**
> NginxObject GetNginx(ctx, optional)
Return status of nginx running instance

Returns nginx version, build name, address, number of configuration reloads, IDs of master and worker processes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string**| Limits which fields of nginx running instance will be output. | 

### Return type

[**NginxObject**](NginxObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

