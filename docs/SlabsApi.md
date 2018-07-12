# \SlabsApi

All URIs are relative to *http://localhost/api/3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSlabZoneStats**](SlabsApi.md#DeleteSlabZoneStats) | **Delete** /slabs/{slabZoneName} | Reset slab statistics
[**GetSlabZone**](SlabsApi.md#GetSlabZone) | **Get** /slabs/{slabZoneName} | Return status of a slab
[**GetSlabs**](SlabsApi.md#GetSlabs) | **Get** /slabs/ | Return status of all slabs


# **DeleteSlabZoneStats**
> DeleteSlabZoneStats(ctx, slabZoneName)
Reset slab statistics

Resets the “<literal>reqs</literal>” and “<literal>fails</literal>” metrics for each memory slot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **slabZoneName** | **string**| The name of the shared memory zone with slab allocator. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSlabZone**
> NginxSlabZone GetSlabZone(ctx, slabZoneName, optional)
Return status of a slab

Returns status of slabs for a particular shared memory zone with slab allocator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **slabZoneName** | **string**| The name of the shared memory zone with slab allocator. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slabZoneName** | **string**| The name of the shared memory zone with slab allocator. | 
 **fields** | **string**| Limits which fields of the slab zone will be output. | 

### Return type

[**NginxSlabZone**](NginxSlabZone.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSlabs**
> NginxSlabZoneMap GetSlabs(ctx, optional)
Return status of all slabs

Returns status of slabs for each shared memory zone with slab allocator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string**| Limits which fields of slab zones will be output. If the “&lt;literal&gt;fields&lt;/literal&gt;” value is empty, then only zone names are output. | 

### Return type

[**NginxSlabZoneMap**](NginxSlabZoneMap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

