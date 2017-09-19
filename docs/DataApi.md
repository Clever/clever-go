# \DataApi

All URIs are relative to *https://api.clever.com/v1.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContact**](DataApi.md#GetContact) | **Get** /contacts/{id} | 
[**GetContacts**](DataApi.md#GetContacts) | **Get** /contacts | 
[**GetContactsForStudent**](DataApi.md#GetContactsForStudent) | **Get** /students/{id}/contacts | 
[**GetDistrict**](DataApi.md#GetDistrict) | **Get** /districts/{id} | 
[**GetDistrictAdmin**](DataApi.md#GetDistrictAdmin) | **Get** /district_admins/{id} | 
[**GetDistrictAdmins**](DataApi.md#GetDistrictAdmins) | **Get** /district_admins | 
[**GetDistrictForSchool**](DataApi.md#GetDistrictForSchool) | **Get** /schools/{id}/district | 
[**GetDistrictForSection**](DataApi.md#GetDistrictForSection) | **Get** /sections/{id}/district | 
[**GetDistrictForStudent**](DataApi.md#GetDistrictForStudent) | **Get** /students/{id}/district | 
[**GetDistrictForStudentContact**](DataApi.md#GetDistrictForStudentContact) | **Get** /contacts/{id}/district | 
[**GetDistrictForTeacher**](DataApi.md#GetDistrictForTeacher) | **Get** /teachers/{id}/district | 
[**GetDistrictStatus**](DataApi.md#GetDistrictStatus) | **Get** /districts/{id}/status | 
[**GetDistricts**](DataApi.md#GetDistricts) | **Get** /districts | 
[**GetGradeLevelsForTeacher**](DataApi.md#GetGradeLevelsForTeacher) | **Get** /teachers/{id}/grade_levels | 
[**GetSchool**](DataApi.md#GetSchool) | **Get** /schools/{id} | 
[**GetSchoolAdmin**](DataApi.md#GetSchoolAdmin) | **Get** /school_admins/{id} | 
[**GetSchoolAdmins**](DataApi.md#GetSchoolAdmins) | **Get** /school_admins | 
[**GetSchoolForSection**](DataApi.md#GetSchoolForSection) | **Get** /sections/{id}/school | 
[**GetSchoolForStudent**](DataApi.md#GetSchoolForStudent) | **Get** /students/{id}/school | 
[**GetSchoolForTeacher**](DataApi.md#GetSchoolForTeacher) | **Get** /teachers/{id}/school | 
[**GetSchools**](DataApi.md#GetSchools) | **Get** /schools | 
[**GetSchoolsForSchoolAdmin**](DataApi.md#GetSchoolsForSchoolAdmin) | **Get** /school_admins/{id}/schools | 
[**GetSection**](DataApi.md#GetSection) | **Get** /sections/{id} | 
[**GetSections**](DataApi.md#GetSections) | **Get** /sections | 
[**GetSectionsForSchool**](DataApi.md#GetSectionsForSchool) | **Get** /schools/{id}/sections | 
[**GetSectionsForStudent**](DataApi.md#GetSectionsForStudent) | **Get** /students/{id}/sections | 
[**GetSectionsForTeacher**](DataApi.md#GetSectionsForTeacher) | **Get** /teachers/{id}/sections | 
[**GetStudent**](DataApi.md#GetStudent) | **Get** /students/{id} | 
[**GetStudentForContact**](DataApi.md#GetStudentForContact) | **Get** /contacts/{id}/student | 
[**GetStudents**](DataApi.md#GetStudents) | **Get** /students | 
[**GetStudentsForSchool**](DataApi.md#GetStudentsForSchool) | **Get** /schools/{id}/students | 
[**GetStudentsForSection**](DataApi.md#GetStudentsForSection) | **Get** /sections/{id}/students | 
[**GetStudentsForTeacher**](DataApi.md#GetStudentsForTeacher) | **Get** /teachers/{id}/students | 
[**GetTeacher**](DataApi.md#GetTeacher) | **Get** /teachers/{id} | 
[**GetTeacherForSection**](DataApi.md#GetTeacherForSection) | **Get** /sections/{id}/teacher | 
[**GetTeachers**](DataApi.md#GetTeachers) | **Get** /teachers | 
[**GetTeachersForSchool**](DataApi.md#GetTeachersForSchool) | **Get** /schools/{id}/teachers | 
[**GetTeachersForSection**](DataApi.md#GetTeachersForSection) | **Get** /sections/{id}/teachers | 
[**GetTeachersForStudent**](DataApi.md#GetTeachersForStudent) | **Get** /students/{id}/teachers | 


# **GetContact**
> StudentContactResponse GetContact($id)



Returns a specific student contact


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**StudentContactResponse**](StudentContactResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContacts**
> StudentContactsResponse GetContacts($limit, $startingAfter, $endingBefore)



Returns a list of student contacts


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**StudentContactsResponse**](StudentContactsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactsForStudent**
> StudentContactsForStudentResponse GetContactsForStudent($id, $limit)



Returns the contacts for a student


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **limit** | **int32**|  | [optional] 

### Return type

[**StudentContactsForStudentResponse**](StudentContactsForStudentResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDistrict**
> DistrictResponse GetDistrict($id)



Returns a specific district


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**DistrictResponse**](DistrictResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDistrictAdmin**
> DistrictAdminResponse GetDistrictAdmin($id)



Returns a specific district admin


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**DistrictAdminResponse**](DistrictAdminResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDistrictAdmins**
> DistrictAdminsResponse GetDistrictAdmins($startingAfter, $endingBefore)



Returns a list of district admins


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**DistrictAdminsResponse**](DistrictAdminsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDistrictForSchool**
> DistrictResponse GetDistrictForSchool($id)



Returns the district for a school


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**DistrictResponse**](DistrictResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDistrictForSection**
> DistrictResponse GetDistrictForSection($id)



Returns the district for a section


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**DistrictResponse**](DistrictResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDistrictForStudent**
> DistrictResponse GetDistrictForStudent($id)



Returns the district for a student


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**DistrictResponse**](DistrictResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDistrictForStudentContact**
> DistrictResponse GetDistrictForStudentContact($id)



Returns the district for a student contact


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**DistrictResponse**](DistrictResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDistrictForTeacher**
> DistrictResponse GetDistrictForTeacher($id)



Returns the district for a teacher


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**DistrictResponse**](DistrictResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDistrictStatus**
> DistrictStatusResponses GetDistrictStatus($id)



Returns the status of the district


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**DistrictStatusResponses**](DistrictStatusResponses.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDistricts**
> DistrictsResponse GetDistricts()



Returns a list of districts


### Parameters
This endpoint does not need any parameter.

### Return type

[**DistrictsResponse**](DistrictsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGradeLevelsForTeacher**
> GradeLevelsResponse GetGradeLevelsForTeacher($id)



Returns the grade levels for sections a teacher teaches


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**GradeLevelsResponse**](GradeLevelsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSchool**
> SchoolResponse GetSchool($id)



Returns a specific school


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**SchoolResponse**](SchoolResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSchoolAdmin**
> SchoolAdminResponse GetSchoolAdmin($id)



Returns a specific school admin


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**SchoolAdminResponse**](SchoolAdminResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSchoolAdmins**
> SchoolAdminsResponse GetSchoolAdmins($limit, $startingAfter, $endingBefore)



Returns a list of school admins


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**SchoolAdminsResponse**](SchoolAdminsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSchoolForSection**
> SchoolResponse GetSchoolForSection($id)



Returns the school for a section


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**SchoolResponse**](SchoolResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSchoolForStudent**
> SchoolResponse GetSchoolForStudent($id)



Returns the primary school for a student


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**SchoolResponse**](SchoolResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSchoolForTeacher**
> SchoolResponse GetSchoolForTeacher($id)



Retrieves school info for a teacher.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**SchoolResponse**](SchoolResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSchools**
> SchoolsResponse GetSchools($limit, $startingAfter, $endingBefore)



Returns a list of schools


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**SchoolsResponse**](SchoolsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSchoolsForSchoolAdmin**
> SchoolsResponse GetSchoolsForSchoolAdmin($id, $limit, $startingAfter, $endingBefore)



Returns the schools for a school admin


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**SchoolsResponse**](SchoolsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSection**
> SectionResponse GetSection($id)



Returns a specific section


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**SectionResponse**](SectionResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSections**
> SectionsResponse GetSections($limit, $startingAfter, $endingBefore)



Returns a list of sections


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**SectionsResponse**](SectionsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSectionsForSchool**
> SectionsResponse GetSectionsForSchool($id, $limit, $startingAfter, $endingBefore)



Returns the sections for a school


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**SectionsResponse**](SectionsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSectionsForStudent**
> SectionsResponse GetSectionsForStudent($id, $limit, $startingAfter, $endingBefore)



Returns the sections for a student


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**SectionsResponse**](SectionsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSectionsForTeacher**
> SectionsResponse GetSectionsForTeacher($id, $limit, $startingAfter, $endingBefore)



Returns the sections for a teacher


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**SectionsResponse**](SectionsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStudent**
> StudentResponse GetStudent($id)



Returns a specific student


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**StudentResponse**](StudentResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStudentForContact**
> StudentResponse GetStudentForContact($id)



Returns the student for a student contact


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**StudentResponse**](StudentResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStudents**
> StudentsResponse GetStudents($limit, $startingAfter, $endingBefore)



Returns a list of students


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**StudentsResponse**](StudentsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStudentsForSchool**
> StudentsResponse GetStudentsForSchool($id, $limit, $startingAfter, $endingBefore)



Returns the students for a school


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**StudentsResponse**](StudentsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStudentsForSection**
> StudentsResponse GetStudentsForSection($id, $limit, $startingAfter, $endingBefore)



Returns the students for a section


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**StudentsResponse**](StudentsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStudentsForTeacher**
> StudentsResponse GetStudentsForTeacher($id, $limit, $startingAfter, $endingBefore)



Returns the students for a teacher


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**StudentsResponse**](StudentsResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeacher**
> TeacherResponse GetTeacher($id)



Returns a specific teacher


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**TeacherResponse**](TeacherResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeacherForSection**
> TeacherResponse GetTeacherForSection($id)



Returns the primary teacher for a section


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**TeacherResponse**](TeacherResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeachers**
> TeachersResponse GetTeachers($limit, $startingAfter, $endingBefore)



Returns a list of teachers


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**TeachersResponse**](TeachersResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeachersForSchool**
> TeachersResponse GetTeachersForSchool($id, $limit, $startingAfter, $endingBefore)



Returns the teachers for a school


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**TeachersResponse**](TeachersResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeachersForSection**
> TeachersResponse GetTeachersForSection($id, $limit, $startingAfter, $endingBefore)



Returns the teachers for a section


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**TeachersResponse**](TeachersResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeachersForStudent**
> TeachersResponse GetTeachersForStudent($id, $limit, $startingAfter, $endingBefore)



Returns the teachers for a student


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **limit** | **int32**|  | [optional] 
 **startingAfter** | **string**|  | [optional] 
 **endingBefore** | **string**|  | [optional] 

### Return type

[**TeachersResponse**](TeachersResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

