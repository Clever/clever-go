# \EventsApi

All URIs are relative to *https://api.clever.com/v1.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](EventsApi.md#GetEvent) | **Get** /events/{id} | 
[**GetEvents**](EventsApi.md#GetEvents) | **Get** /events | 
[**GetEventsForSchool**](EventsApi.md#GetEventsForSchool) | **Get** /schools/{id}/events | 
[**GetEventsForSchoolAdmin**](EventsApi.md#GetEventsForSchoolAdmin) | **Get** /school_admins/{id}/events | 
[**GetEventsForSection**](EventsApi.md#GetEventsForSection) | **Get** /sections/{id}/events | 
[**GetEventsForStudent**](EventsApi.md#GetEventsForStudent) | **Get** /students/{id}/events | 
[**GetEventsForTeacher**](EventsApi.md#GetEventsForTeacher) | **Get** /teachers/{id}/events | 


# **GetEvent**
> EventResponse GetEvent($id)



Returns the specific event


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEvents**
> EventsResponse GetEvents($limit, $startingAfter, $endingBefore)



Returns a list of events


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**EventsResponse**](EventsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventsForSchool**
> EventsResponse GetEventsForSchool($id, $limit, $startingAfter, $endingBefore)



Returns a list of events for a school


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**EventsResponse**](EventsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventsForSchoolAdmin**
> EventsResponse GetEventsForSchoolAdmin($id, $limit, $startingAfter, $endingBefore)



Returns a list of events for a school admin


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**EventsResponse**](EventsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventsForSection**
> EventsResponse GetEventsForSection($id, $limit, $startingAfter, $endingBefore)



Returns a list of events for a section


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**EventsResponse**](EventsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventsForStudent**
> EventsResponse GetEventsForStudent($id, $limit, $startingAfter, $endingBefore)



Returns a list of events for a student


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**EventsResponse**](EventsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventsForTeacher**
> EventsResponse GetEventsForTeacher($id, $limit, $startingAfter, $endingBefore)



Returns a list of events for a teacher


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**EventsResponse**](EventsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

