# \DefaultApi

All URIs are relative to *https://app.sysdigcloud.com/ui*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboard**](DefaultApi.md#CreateDashboard) | **Post** /dashboards | 
[**DeleteDashboard**](DefaultApi.md#DeleteDashboard) | **Delete** /dashboards/{id} | 
[**GetDashboard**](DefaultApi.md#GetDashboard) | **Get** /dashboards/{id} | 
[**GetDashboards**](DefaultApi.md#GetDashboards) | **Get** /dashboards | 
[**UpdateDashboard**](DefaultApi.md#UpdateDashboard) | **Put** /dashboards/{id} | 


# **CreateDashboard**
> DashboardInput CreateDashboard($dashboardInput)



Creates a new Dashboard


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboardInput** | [**DashboardInput**](DashboardInput.md)| DashboardInput | 

### Return type

[**DashboardInput**](DashboardInput.md)

### Authorization

[bearerHeader](../README.md#bearerHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDashboard**
> Dashboards DeleteDashboard($id)



Deletes a new Dashboard


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| DashboardID | 

### Return type

[**Dashboards**](Dashboards.md)

### Authorization

[bearerHeader](../README.md#bearerHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboard**
> DashboardInput GetDashboard($id)



Get Dashboard


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| DashboardID | 

### Return type

[**DashboardInput**](DashboardInput.md)

### Authorization

[bearerHeader](../README.md#bearerHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboards**
> Dashboards GetDashboards()



Get Dashboards


### Parameters
This endpoint does not need any parameter.

### Return type

[**Dashboards**](Dashboards.md)

### Authorization

[bearerHeader](../README.md#bearerHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDashboard**
> DashboardInput UpdateDashboard($dashboardInput, $id)



Updates a new Dashboard


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboardInput** | [**DashboardInput**](DashboardInput.md)| DashboardInput | 
 **id** | **string**| DashboardID | 

### Return type

[**DashboardInput**](DashboardInput.md)

### Authorization

[bearerHeader](../README.md#bearerHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

