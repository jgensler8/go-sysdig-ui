# \DefaultApi

All URIs are relative to *https://app.sysdigcloud.com/ui*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboard**](DefaultApi.md#CreateDashboard) | **Post** /dashboards | 
[**GetDashboards**](DefaultApi.md#GetDashboards) | **Get** /dashboards | 


# **CreateDashboard**
> Dashboards CreateDashboard($body)



Creates a new Dashboard


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Dashboard**](Dashboard.md)| Dashboard | 

### Return type

[**Dashboards**](Dashboards.md)

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

