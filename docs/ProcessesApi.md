# \ProcessesApi

All URIs are relative to *http://localhost/api/3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcesses**](ProcessesApi.md#DeleteProcesses) | **Delete** /processes | Reset nginx processes statistics
[**GetProcesses**](ProcessesApi.md#GetProcesses) | **Get** /processes | Return nginx processes status


# **DeleteProcesses**
> DeleteProcesses(ctx, )
Reset nginx processes statistics

Resets counters of abnormally terminated and respawned child processes.

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

# **GetProcesses**
> NginxProcesses GetProcesses(ctx, )
Return nginx processes status

Returns the number of abnormally terminated and respawned child processes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NginxProcesses**](NginxProcesses.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

