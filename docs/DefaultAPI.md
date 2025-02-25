# \DefaultAPI

All URIs are relative to *https://api.weather.gov*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertsActive**](DefaultAPI.md#AlertsActive) | **Get** /alerts/active | 
[**AlertsActiveArea**](DefaultAPI.md#AlertsActiveArea) | **Get** /alerts/active/area/{area} | 
[**AlertsActiveCount**](DefaultAPI.md#AlertsActiveCount) | **Get** /alerts/active/count | 
[**AlertsActiveRegion**](DefaultAPI.md#AlertsActiveRegion) | **Get** /alerts/active/region/{region} | 
[**AlertsActiveZone**](DefaultAPI.md#AlertsActiveZone) | **Get** /alerts/active/zone/{zoneId} | 
[**AlertsQuery**](DefaultAPI.md#AlertsQuery) | **Get** /alerts | 
[**AlertsSingle**](DefaultAPI.md#AlertsSingle) | **Get** /alerts/{id} | 
[**AlertsTypes**](DefaultAPI.md#AlertsTypes) | **Get** /alerts/types | 
[**Cwa**](DefaultAPI.md#Cwa) | **Get** /aviation/cwsus/{cwsuId}/cwas/{date}/{sequence} | 
[**Cwas**](DefaultAPI.md#Cwas) | **Get** /aviation/cwsus/{cwsuId}/cwas | 
[**Cwsu**](DefaultAPI.md#Cwsu) | **Get** /aviation/cwsus/{cwsuId} | 
[**Glossary**](DefaultAPI.md#Glossary) | **Get** /glossary | 
[**Gridpoint**](DefaultAPI.md#Gridpoint) | **Get** /gridpoints/{wfo}/{x},{y} | 
[**GridpointForecast**](DefaultAPI.md#GridpointForecast) | **Get** /gridpoints/{wfo}/{x},{y}/forecast | 
[**GridpointForecastHourly**](DefaultAPI.md#GridpointForecastHourly) | **Get** /gridpoints/{wfo}/{x},{y}/forecast/hourly | 
[**GridpointStations**](DefaultAPI.md#GridpointStations) | **Get** /gridpoints/{wfo}/{x},{y}/stations | 
[**Icons**](DefaultAPI.md#Icons) | **Get** /icons/{set}/{timeOfDay}/{first} | 
[**IconsDualCondition**](DefaultAPI.md#IconsDualCondition) | **Get** /icons/{set}/{timeOfDay}/{first}/{second} | 
[**IconsSummary**](DefaultAPI.md#IconsSummary) | **Get** /icons | 
[**LocationProducts**](DefaultAPI.md#LocationProducts) | **Get** /products/locations/{locationId}/types | 
[**ObsStation**](DefaultAPI.md#ObsStation) | **Get** /stations/{stationId} | 
[**ObsStations**](DefaultAPI.md#ObsStations) | **Get** /stations | 
[**Office**](DefaultAPI.md#Office) | **Get** /offices/{officeId} | 
[**OfficeHeadline**](DefaultAPI.md#OfficeHeadline) | **Get** /offices/{officeId}/headlines/{headlineId} | 
[**OfficeHeadlines**](DefaultAPI.md#OfficeHeadlines) | **Get** /offices/{officeId}/headlines | 
[**Point**](DefaultAPI.md#Point) | **Get** /points/{point} | 
[**PointStations**](DefaultAPI.md#PointStations) | **Get** /points/{point}/stations | 
[**Product**](DefaultAPI.md#Product) | **Get** /products/{productId} | 
[**ProductLocations**](DefaultAPI.md#ProductLocations) | **Get** /products/locations | 
[**ProductTypes**](DefaultAPI.md#ProductTypes) | **Get** /products/types | 
[**ProductsQuery**](DefaultAPI.md#ProductsQuery) | **Get** /products | 
[**ProductsType**](DefaultAPI.md#ProductsType) | **Get** /products/types/{typeId} | 
[**ProductsTypeLocation**](DefaultAPI.md#ProductsTypeLocation) | **Get** /products/types/{typeId}/locations/{locationId} | 
[**ProductsTypeLocations**](DefaultAPI.md#ProductsTypeLocations) | **Get** /products/types/{typeId}/locations | 
[**RadarProfiler**](DefaultAPI.md#RadarProfiler) | **Get** /radar/profilers/{stationId} | 
[**RadarQueue**](DefaultAPI.md#RadarQueue) | **Get** /radar/queues/{host} | 
[**RadarServer**](DefaultAPI.md#RadarServer) | **Get** /radar/servers/{id} | 
[**RadarServers**](DefaultAPI.md#RadarServers) | **Get** /radar/servers | 
[**RadarStation**](DefaultAPI.md#RadarStation) | **Get** /radar/stations/{stationId} | 
[**RadarStationAlarms**](DefaultAPI.md#RadarStationAlarms) | **Get** /radar/stations/{stationId}/alarms | 
[**RadarStations**](DefaultAPI.md#RadarStations) | **Get** /radar/stations | 
[**SatelliteThumbnails**](DefaultAPI.md#SatelliteThumbnails) | **Get** /thumbnails/satellite/{area} | 
[**Sigmet**](DefaultAPI.md#Sigmet) | **Get** /aviation/sigmets/{atsu}/{date}/{time} | 
[**SigmetQuery**](DefaultAPI.md#SigmetQuery) | **Get** /aviation/sigmets | 
[**SigmetsByATSU**](DefaultAPI.md#SigmetsByATSU) | **Get** /aviation/sigmets/{atsu} | 
[**SigmetsByATSUByDate**](DefaultAPI.md#SigmetsByATSUByDate) | **Get** /aviation/sigmets/{atsu}/{date} | 
[**StationObservationLatest**](DefaultAPI.md#StationObservationLatest) | **Get** /stations/{stationId}/observations/latest | 
[**StationObservationList**](DefaultAPI.md#StationObservationList) | **Get** /stations/{stationId}/observations | 
[**StationObservationTime**](DefaultAPI.md#StationObservationTime) | **Get** /stations/{stationId}/observations/{time} | 
[**Taf**](DefaultAPI.md#Taf) | **Get** /stations/{stationId}/tafs/{date}/{time} | 
[**Tafs**](DefaultAPI.md#Tafs) | **Get** /stations/{stationId}/tafs | 
[**Zone**](DefaultAPI.md#Zone) | **Get** /zones/{type}/{zoneId} | 
[**ZoneForecast**](DefaultAPI.md#ZoneForecast) | **Get** /zones/{type}/{zoneId}/forecast | 
[**ZoneList**](DefaultAPI.md#ZoneList) | **Get** /zones | 
[**ZoneListType**](DefaultAPI.md#ZoneListType) | **Get** /zones/{type} | 
[**ZoneObs**](DefaultAPI.md#ZoneObs) | **Get** /zones/forecast/{zoneId}/observations | 
[**ZoneStations**](DefaultAPI.md#ZoneStations) | **Get** /zones/forecast/{zoneId}/stations | 



## AlertsActive

> AlertCollectionGeoJson AlertsActive(ctx).Status(status).MessageType(messageType).Event(event).Code(code).Area(area).Point(point).Region(region).RegionType(regionType).Zone(zone).Urgency(urgency).Severity(severity).Certainty(certainty).Limit(limit).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	status := []string{"Status_example"} // []string | Status (actual, exercise, system, test, draft) (optional)
	messageType := []string{"MessageType_example"} // []string | Message type (alert, update, cancel) (optional)
	event := []string{"Inner_example"} // []string | Event name (optional)
	code := []string{"Inner_example"} // []string | Event code (optional)
	area := []openapiclient.AreaCode{openapiclient.AreaCode{MarineAreaCode: openapiclient.MarineAreaCode("AM")}} // []AreaCode | State/territory code or marine area code This parameter is incompatible with the following parameters: point, region, region_type, zone  (optional)
	point := "point_example" // string | Point (latitude,longitude) This parameter is incompatible with the following parameters: area, region, region_type, zone  (optional)
	region := []openapiclient.MarineRegionCode{openapiclient.MarineRegionCode("AL")} // []MarineRegionCode | Marine region code This parameter is incompatible with the following parameters: area, point, region_type, zone  (optional)
	regionType := "regionType_example" // string | Region type (land or marine) This parameter is incompatible with the following parameters: area, point, region, zone  (optional)
	zone := []string{"Inner_example"} // []string | Zone ID (forecast or county) This parameter is incompatible with the following parameters: area, point, region, region_type  (optional)
	urgency := []openapiclient.AlertUrgency{openapiclient.AlertUrgency("Immediate")} // []AlertUrgency | Urgency (immediate, expected, future, past, unknown) (optional)
	severity := []openapiclient.AlertSeverity{openapiclient.AlertSeverity("Extreme")} // []AlertSeverity | Severity (extreme, severe, moderate, minor, unknown) (optional)
	certainty := []openapiclient.AlertCertainty{openapiclient.AlertCertainty("Observed")} // []AlertCertainty | Certainty (observed, likely, possible, unlikely, unknown) (optional)
	limit := int32(56) // int32 | Limit (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AlertsActive(context.Background()).Status(status).MessageType(messageType).Event(event).Code(code).Area(area).Point(point).Region(region).RegionType(regionType).Zone(zone).Urgency(urgency).Severity(severity).Certainty(certainty).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AlertsActive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertsActive`: AlertCollectionGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AlertsActive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsActiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **[]string** | Status (actual, exercise, system, test, draft) | 
 **messageType** | **[]string** | Message type (alert, update, cancel) | 
 **event** | **[]string** | Event name | 
 **code** | **[]string** | Event code | 
 **area** | [**[]AreaCode**](AreaCode.md) | State/territory code or marine area code This parameter is incompatible with the following parameters: point, region, region_type, zone  | 
 **point** | **string** | Point (latitude,longitude) This parameter is incompatible with the following parameters: area, region, region_type, zone  | 
 **region** | [**[]MarineRegionCode**](MarineRegionCode.md) | Marine region code This parameter is incompatible with the following parameters: area, point, region_type, zone  | 
 **regionType** | **string** | Region type (land or marine) This parameter is incompatible with the following parameters: area, point, region, zone  | 
 **zone** | **[]string** | Zone ID (forecast or county) This parameter is incompatible with the following parameters: area, point, region, region_type  | 
 **urgency** | [**[]AlertUrgency**](AlertUrgency.md) | Urgency (immediate, expected, future, past, unknown) | 
 **severity** | [**[]AlertSeverity**](AlertSeverity.md) | Severity (extreme, severe, moderate, minor, unknown) | 
 **certainty** | [**[]AlertCertainty**](AlertCertainty.md) | Certainty (observed, likely, possible, unlikely, unknown) | 
 **limit** | **int32** | Limit | [default to 500]

### Return type

[**AlertCollectionGeoJson**](AlertCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/atom+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsActiveArea

> AlertCollectionGeoJson AlertsActiveArea(ctx, area).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	area := openapiclient.AreaCode{MarineAreaCode: openapiclient.MarineAreaCode("AM")} // AreaCode | State/area ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AlertsActiveArea(context.Background(), area).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AlertsActiveArea``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertsActiveArea`: AlertCollectionGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AlertsActiveArea`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**area** | [**AreaCode**](.md) | State/area ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsActiveAreaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertCollectionGeoJson**](AlertCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/atom+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsActiveCount

> AlertsActiveCount200Response AlertsActiveCount(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AlertsActiveCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AlertsActiveCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertsActiveCount`: AlertsActiveCount200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AlertsActiveCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsActiveCountRequest struct via the builder pattern


### Return type

[**AlertsActiveCount200Response**](AlertsActiveCount200Response.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsActiveRegion

> AlertCollectionGeoJson AlertsActiveRegion(ctx, region).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	region := openapiclient.MarineRegionCode("AL") // MarineRegionCode | Marine region ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AlertsActiveRegion(context.Background(), region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AlertsActiveRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertsActiveRegion`: AlertCollectionGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AlertsActiveRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | [**MarineRegionCode**](.md) | Marine region ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsActiveRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertCollectionGeoJson**](AlertCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/atom+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsActiveZone

> AlertCollectionGeoJson AlertsActiveZone(ctx, zoneId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	zoneId := "zoneId_example" // string | NWS public zone/county identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AlertsActiveZone(context.Background(), zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AlertsActiveZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertsActiveZone`: AlertCollectionGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AlertsActiveZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | NWS public zone/county identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsActiveZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertCollectionGeoJson**](AlertCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/atom+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsQuery

> AlertCollectionGeoJson AlertsQuery(ctx).Active(active).Start(start).End(end).Status(status).MessageType(messageType).Event(event).Code(code).Area(area).Point(point).Region(region).RegionType(regionType).Zone(zone).Urgency(urgency).Severity(severity).Certainty(certainty).Limit(limit).Cursor(cursor).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	active := true // bool | List only active alerts (use /alerts/active endpoints instead) (optional)
	start := time.Now() // time.Time | Start time (optional)
	end := time.Now() // time.Time | End time (optional)
	status := []string{"Status_example"} // []string | Status (actual, exercise, system, test, draft) (optional)
	messageType := []string{"MessageType_example"} // []string | Message type (alert, update, cancel) (optional)
	event := []string{"Inner_example"} // []string | Event name (optional)
	code := []string{"Inner_example"} // []string | Event code (optional)
	area := []openapiclient.AreaCode{openapiclient.AreaCode{MarineAreaCode: openapiclient.MarineAreaCode("AM")}} // []AreaCode | State/territory code or marine area code This parameter is incompatible with the following parameters: point, region, region_type, zone  (optional)
	point := "point_example" // string | Point (latitude,longitude) This parameter is incompatible with the following parameters: area, region, region_type, zone  (optional)
	region := []openapiclient.MarineRegionCode{openapiclient.MarineRegionCode("AL")} // []MarineRegionCode | Marine region code This parameter is incompatible with the following parameters: area, point, region_type, zone  (optional)
	regionType := "regionType_example" // string | Region type (land or marine) This parameter is incompatible with the following parameters: area, point, region, zone  (optional)
	zone := []string{"Inner_example"} // []string | Zone ID (forecast or county) This parameter is incompatible with the following parameters: area, point, region, region_type  (optional)
	urgency := []openapiclient.AlertUrgency{openapiclient.AlertUrgency("Immediate")} // []AlertUrgency | Urgency (immediate, expected, future, past, unknown) (optional)
	severity := []openapiclient.AlertSeverity{openapiclient.AlertSeverity("Extreme")} // []AlertSeverity | Severity (extreme, severe, moderate, minor, unknown) (optional)
	certainty := []openapiclient.AlertCertainty{openapiclient.AlertCertainty("Observed")} // []AlertCertainty | Certainty (observed, likely, possible, unlikely, unknown) (optional)
	limit := int32(56) // int32 | Limit (optional) (default to 500)
	cursor := "cursor_example" // string | Pagination cursor (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AlertsQuery(context.Background()).Active(active).Start(start).End(end).Status(status).MessageType(messageType).Event(event).Code(code).Area(area).Point(point).Region(region).RegionType(regionType).Zone(zone).Urgency(urgency).Severity(severity).Certainty(certainty).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AlertsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertsQuery`: AlertCollectionGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AlertsQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **bool** | List only active alerts (use /alerts/active endpoints instead) | 
 **start** | **time.Time** | Start time | 
 **end** | **time.Time** | End time | 
 **status** | **[]string** | Status (actual, exercise, system, test, draft) | 
 **messageType** | **[]string** | Message type (alert, update, cancel) | 
 **event** | **[]string** | Event name | 
 **code** | **[]string** | Event code | 
 **area** | [**[]AreaCode**](AreaCode.md) | State/territory code or marine area code This parameter is incompatible with the following parameters: point, region, region_type, zone  | 
 **point** | **string** | Point (latitude,longitude) This parameter is incompatible with the following parameters: area, region, region_type, zone  | 
 **region** | [**[]MarineRegionCode**](MarineRegionCode.md) | Marine region code This parameter is incompatible with the following parameters: area, point, region_type, zone  | 
 **regionType** | **string** | Region type (land or marine) This parameter is incompatible with the following parameters: area, point, region, zone  | 
 **zone** | **[]string** | Zone ID (forecast or county) This parameter is incompatible with the following parameters: area, point, region, region_type  | 
 **urgency** | [**[]AlertUrgency**](AlertUrgency.md) | Urgency (immediate, expected, future, past, unknown) | 
 **severity** | [**[]AlertSeverity**](AlertSeverity.md) | Severity (extreme, severe, moderate, minor, unknown) | 
 **certainty** | [**[]AlertCertainty**](AlertCertainty.md) | Certainty (observed, likely, possible, unlikely, unknown) | 
 **limit** | **int32** | Limit | [default to 500]
 **cursor** | **string** | Pagination cursor | 

### Return type

[**AlertCollectionGeoJson**](AlertCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/atom+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsSingle

> AlertGeoJson AlertsSingle(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	id := "id_example" // string | Alert identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AlertsSingle(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AlertsSingle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertsSingle`: AlertGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AlertsSingle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Alert identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsSingleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertGeoJson**](AlertGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/cap+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsTypes

> AlertsTypes200Response AlertsTypes(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AlertsTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AlertsTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertsTypes`: AlertsTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AlertsTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsTypesRequest struct via the builder pattern


### Return type

[**AlertsTypes200Response**](AlertsTypes200Response.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Cwa

> CenterWeatherAdvisoryGeoJson Cwa(ctx, cwsuId, date, sequence).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	cwsuId := openapiclient.NWSCenterWeatherServiceUnitId("ZAB") // NWSCenterWeatherServiceUnitId | NWS CWSU ID
	date := time.Now() // string | Date (YYYY-MM-DD format)
	sequence := int32(56) // int32 | Sequence number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Cwa(context.Background(), cwsuId, date, sequence).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Cwa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Cwa`: CenterWeatherAdvisoryGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Cwa`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cwsuId** | [**NWSCenterWeatherServiceUnitId**](.md) | NWS CWSU ID | 
**date** | **string** | Date (YYYY-MM-DD format) | 
**sequence** | **int32** | Sequence number | 

### Other Parameters

Other parameters are passed through a pointer to a apiCwaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CenterWeatherAdvisoryGeoJson**](CenterWeatherAdvisoryGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/vnd.noaa.uswx+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Cwas

> CenterWeatherAdvisoryCollectionGeoJson Cwas(ctx, cwsuId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	cwsuId := openapiclient.NWSCenterWeatherServiceUnitId("ZAB") // NWSCenterWeatherServiceUnitId | NWS CWSU ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Cwas(context.Background(), cwsuId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Cwas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Cwas`: CenterWeatherAdvisoryCollectionGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Cwas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cwsuId** | [**NWSCenterWeatherServiceUnitId**](.md) | NWS CWSU ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCwasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CenterWeatherAdvisoryCollectionGeoJson**](CenterWeatherAdvisoryCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Cwsu

> Office Cwsu(ctx, cwsuId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	cwsuId := openapiclient.NWSCenterWeatherServiceUnitId("ZAB") // NWSCenterWeatherServiceUnitId | NWS CWSU ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Cwsu(context.Background(), cwsuId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Cwsu``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Cwsu`: Office
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Cwsu`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cwsuId** | [**NWSCenterWeatherServiceUnitId**](.md) | NWS CWSU ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCwsuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Office**](Office.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Glossary

> Glossary200Response Glossary(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Glossary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Glossary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Glossary`: Glossary200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Glossary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGlossaryRequest struct via the builder pattern


### Return type

[**Glossary200Response**](Glossary200Response.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Gridpoint

> GridpointGeoJson Gridpoint(ctx, wfo, x, y).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	wfo := openapiclient.NWSForecastOfficeId("AKQ") // NWSForecastOfficeId | Forecast office ID
	x := int32(56) // int32 | Forecast grid X coordinate
	y := int32(56) // int32 | Forecast grid Y coordinate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Gridpoint(context.Background(), wfo, x, y).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Gridpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Gridpoint`: GridpointGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Gridpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wfo** | [**NWSForecastOfficeId**](.md) | Forecast office ID | 
**x** | **int32** | Forecast grid X coordinate | 
**y** | **int32** | Forecast grid Y coordinate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGridpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GridpointGeoJson**](GridpointGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridpointForecast

> GridpointForecastGeoJson GridpointForecast(ctx, wfo, x, y).FeatureFlags(featureFlags).Units(units).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	wfo := openapiclient.NWSForecastOfficeId("AKQ") // NWSForecastOfficeId | Forecast office ID
	x := int32(56) // int32 | Forecast grid X coordinate
	y := int32(56) // int32 | Forecast grid Y coordinate
	featureFlags := []string{"FeatureFlags_example"} // []string | Enable future and experimental features (see documentation for more info): * forecast_temperature_qv: Represent temperature as QuantitativeValue * forecast_wind_speed_qv: Represent wind speed as QuantitativeValue  (optional)
	units := openapiclient.GridpointForecastUnits("us") // GridpointForecastUnits | Use US customary or SI (metric) units in textual output (optional) (default to "us")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GridpointForecast(context.Background(), wfo, x, y).FeatureFlags(featureFlags).Units(units).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GridpointForecast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridpointForecast`: GridpointForecastGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GridpointForecast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wfo** | [**NWSForecastOfficeId**](.md) | Forecast office ID | 
**x** | **int32** | Forecast grid X coordinate | 
**y** | **int32** | Forecast grid Y coordinate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGridpointForecastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **featureFlags** | **[]string** | Enable future and experimental features (see documentation for more info): * forecast_temperature_qv: Represent temperature as QuantitativeValue * forecast_wind_speed_qv: Represent wind speed as QuantitativeValue  | 
 **units** | [**GridpointForecastUnits**](GridpointForecastUnits.md) | Use US customary or SI (metric) units in textual output | [default to &quot;us&quot;]

### Return type

[**GridpointForecastGeoJson**](GridpointForecastGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/vnd.noaa.dwml+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridpointForecastHourly

> GridpointForecastGeoJson GridpointForecastHourly(ctx, wfo, x, y).FeatureFlags(featureFlags).Units(units).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	wfo := openapiclient.NWSForecastOfficeId("AKQ") // NWSForecastOfficeId | Forecast office ID
	x := int32(56) // int32 | Forecast grid X coordinate
	y := int32(56) // int32 | Forecast grid Y coordinate
	featureFlags := []string{"FeatureFlags_example"} // []string | Enable future and experimental features (see documentation for more info): * forecast_temperature_qv: Represent temperature as QuantitativeValue * forecast_wind_speed_qv: Represent wind speed as QuantitativeValue  (optional)
	units := openapiclient.GridpointForecastUnits("us") // GridpointForecastUnits | Use US customary or SI (metric) units in textual output (optional) (default to "us")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GridpointForecastHourly(context.Background(), wfo, x, y).FeatureFlags(featureFlags).Units(units).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GridpointForecastHourly``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridpointForecastHourly`: GridpointForecastGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GridpointForecastHourly`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wfo** | [**NWSForecastOfficeId**](.md) | Forecast office ID | 
**x** | **int32** | Forecast grid X coordinate | 
**y** | **int32** | Forecast grid Y coordinate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGridpointForecastHourlyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **featureFlags** | **[]string** | Enable future and experimental features (see documentation for more info): * forecast_temperature_qv: Represent temperature as QuantitativeValue * forecast_wind_speed_qv: Represent wind speed as QuantitativeValue  | 
 **units** | [**GridpointForecastUnits**](GridpointForecastUnits.md) | Use US customary or SI (metric) units in textual output | [default to &quot;us&quot;]

### Return type

[**GridpointForecastGeoJson**](GridpointForecastGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/vnd.noaa.dwml+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridpointStations

> ObservationStationCollectionGeoJson GridpointStations(ctx, wfo, x, y).Limit(limit).Cursor(cursor).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	wfo := openapiclient.NWSForecastOfficeId("AKQ") // NWSForecastOfficeId | Forecast office ID
	x := int32(56) // int32 | Forecast grid X coordinate
	y := int32(56) // int32 | Forecast grid Y coordinate
	limit := int32(56) // int32 | Limit (optional) (default to 500)
	cursor := "cursor_example" // string | Pagination cursor (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GridpointStations(context.Background(), wfo, x, y).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GridpointStations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridpointStations`: ObservationStationCollectionGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GridpointStations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wfo** | [**NWSForecastOfficeId**](.md) | Forecast office ID | 
**x** | **int32** | Forecast grid X coordinate | 
**y** | **int32** | Forecast grid Y coordinate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGridpointStationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | Limit | [default to 500]
 **cursor** | **string** | Pagination cursor | 

### Return type

[**ObservationStationCollectionGeoJson**](ObservationStationCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Icons

> *os.File Icons(ctx, set, timeOfDay, first).Size(size).Fontsize(fontsize).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	set := "set_example" // string | .
	timeOfDay := "timeOfDay_example" // string | .
	first := "first_example" // string | .
	size := *openapiclient.NewIconsSizeParameter() // IconsSizeParameter | Font size (optional)
	fontsize := int32(56) // int32 | Font size (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Icons(context.Background(), set, timeOfDay, first).Size(size).Fontsize(fontsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Icons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Icons`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Icons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**set** | **string** | . | 
**timeOfDay** | **string** | . | 
**first** | **string** | . | 

### Other Parameters

Other parameters are passed through a pointer to a apiIconsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **size** | [**IconsSizeParameter**](IconsSizeParameter.md) | Font size | 
 **fontsize** | **int32** | Font size | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IconsDualCondition

> *os.File IconsDualCondition(ctx, set, timeOfDay, first, second).Size(size).Fontsize(fontsize).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	set := "set_example" // string | .
	timeOfDay := "timeOfDay_example" // string | .
	first := "first_example" // string | .
	second := "second_example" // string | .
	size := *openapiclient.NewIconsSizeParameter() // IconsSizeParameter | Font size (optional)
	fontsize := int32(56) // int32 | Font size (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.IconsDualCondition(context.Background(), set, timeOfDay, first, second).Size(size).Fontsize(fontsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.IconsDualCondition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IconsDualCondition`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.IconsDualCondition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**set** | **string** | . | 
**timeOfDay** | **string** | . | 
**first** | **string** | . | 
**second** | **string** | . | 

### Other Parameters

Other parameters are passed through a pointer to a apiIconsDualConditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **size** | [**IconsSizeParameter**](IconsSizeParameter.md) | Font size | 
 **fontsize** | **int32** | Font size | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IconsSummary

> IconsSummary200Response IconsSummary(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.IconsSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.IconsSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IconsSummary`: IconsSummary200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.IconsSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIconsSummaryRequest struct via the builder pattern


### Return type

[**IconsSummary200Response**](IconsSummary200Response.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocationProducts

> TextProductTypeCollection LocationProducts(ctx, locationId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	locationId := "locationId_example" // string | .

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.LocationProducts(context.Background(), locationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.LocationProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocationProducts`: TextProductTypeCollection
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.LocationProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | . | 

### Other Parameters

Other parameters are passed through a pointer to a apiLocationProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TextProductTypeCollection**](TextProductTypeCollection.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObsStation

> ObservationStationGeoJson ObsStation(ctx, stationId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	stationId := "stationId_example" // string | Observation station ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ObsStation(context.Background(), stationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ObsStation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObsStation`: ObservationStationGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ObsStation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stationId** | **string** | Observation station ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiObsStationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObservationStationGeoJson**](ObservationStationGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObsStations

> ObservationStationCollectionGeoJson ObsStations(ctx).Id(id).State(state).Limit(limit).Cursor(cursor).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	id := []string{"Inner_example"} // []string | Filter by observation station ID (optional)
	state := []openapiclient.AreaCode{openapiclient.AreaCode{MarineAreaCode: openapiclient.MarineAreaCode("AM")}} // []AreaCode | Filter by state/marine area code (optional)
	limit := int32(56) // int32 | Limit (optional) (default to 500)
	cursor := "cursor_example" // string | Pagination cursor (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ObsStations(context.Background()).Id(id).State(state).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ObsStations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObsStations`: ObservationStationCollectionGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ObsStations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiObsStationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | Filter by observation station ID | 
 **state** | [**[]AreaCode**](AreaCode.md) | Filter by state/marine area code | 
 **limit** | **int32** | Limit | [default to 500]
 **cursor** | **string** | Pagination cursor | 

### Return type

[**ObservationStationCollectionGeoJson**](ObservationStationCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Office

> Office Office(ctx, officeId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	officeId := openapiclient.NWSOfficeId{NWSForecastOfficeId: openapiclient.NWSForecastOfficeId("AKQ")} // NWSOfficeId | NWS office ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Office(context.Background(), officeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Office``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Office`: Office
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Office`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | [**NWSOfficeId**](.md) | NWS office ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Office**](Office.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficeHeadline

> OfficeHeadline OfficeHeadline(ctx, officeId, headlineId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	officeId := openapiclient.NWSOfficeId{NWSForecastOfficeId: openapiclient.NWSForecastOfficeId("AKQ")} // NWSOfficeId | NWS office ID
	headlineId := "headlineId_example" // string | Headline record ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.OfficeHeadline(context.Background(), officeId, headlineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OfficeHeadline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficeHeadline`: OfficeHeadline
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OfficeHeadline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | [**NWSOfficeId**](.md) | NWS office ID | 
**headlineId** | **string** | Headline record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficeHeadlineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OfficeHeadline**](OfficeHeadline.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficeHeadlines

> OfficeHeadlineCollection OfficeHeadlines(ctx, officeId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	officeId := openapiclient.NWSOfficeId{NWSForecastOfficeId: openapiclient.NWSForecastOfficeId("AKQ")} // NWSOfficeId | NWS office ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.OfficeHeadlines(context.Background(), officeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OfficeHeadlines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficeHeadlines`: OfficeHeadlineCollection
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OfficeHeadlines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | [**NWSOfficeId**](.md) | NWS office ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficeHeadlinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OfficeHeadlineCollection**](OfficeHeadlineCollection.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Point

> PointGeoJson Point(ctx, point).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	point := "point_example" // string | Point (latitude, longitude)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Point(context.Background(), point).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Point``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Point`: PointGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Point`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**point** | **string** | Point (latitude, longitude) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PointGeoJson**](PointGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PointStations

> ProblemDetail PointStations(ctx, point).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	point := "point_example" // string | Point (latitude, longitude)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PointStations(context.Background(), point).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PointStations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PointStations`: ProblemDetail
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PointStations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**point** | **string** | Point (latitude, longitude) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPointStationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProblemDetail**](ProblemDetail.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Product

> TextProduct Product(ctx, productId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	productId := "productId_example" // string | .

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Product(context.Background(), productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Product``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Product`: TextProduct
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Product`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | . | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TextProduct**](TextProduct.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductLocations

> TextProductLocationCollection ProductLocations(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ProductLocations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ProductLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductLocations`: TextProductLocationCollection
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ProductLocations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProductLocationsRequest struct via the builder pattern


### Return type

[**TextProductLocationCollection**](TextProductLocationCollection.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductTypes

> TextProductTypeCollection ProductTypes(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ProductTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ProductTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTypes`: TextProductTypeCollection
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ProductTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProductTypesRequest struct via the builder pattern


### Return type

[**TextProductTypeCollection**](TextProductTypeCollection.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsQuery

> TextProductCollection ProductsQuery(ctx).Location(location).Start(start).End(end).Office(office).Wmoid(wmoid).Type_(type_).Limit(limit).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	location := []string{"Inner_example"} // []string | Location id (optional)
	start := time.Now() // time.Time | Start time (optional)
	end := time.Now() // time.Time | End time (optional)
	office := []string{"Inner_example"} // []string | Issuing office (optional)
	wmoid := []string{"Inner_example"} // []string | WMO id code (optional)
	type_ := []string{"Inner_example"} // []string | Product code (optional)
	limit := int32(56) // int32 | Limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ProductsQuery(context.Background()).Location(location).Start(start).End(end).Office(office).Wmoid(wmoid).Type_(type_).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ProductsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsQuery`: TextProductCollection
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ProductsQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **[]string** | Location id | 
 **start** | **time.Time** | Start time | 
 **end** | **time.Time** | End time | 
 **office** | **[]string** | Issuing office | 
 **wmoid** | **[]string** | WMO id code | 
 **type_** | **[]string** | Product code | 
 **limit** | **int32** | Limit | 

### Return type

[**TextProductCollection**](TextProductCollection.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsType

> TextProductCollection ProductsType(ctx, typeId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	typeId := "typeId_example" // string | .

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ProductsType(context.Background(), typeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ProductsType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsType`: TextProductCollection
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ProductsType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeId** | **string** | . | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductsTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TextProductCollection**](TextProductCollection.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsTypeLocation

> TextProductCollection ProductsTypeLocation(ctx, typeId, locationId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	typeId := "typeId_example" // string | .
	locationId := "locationId_example" // string | .

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ProductsTypeLocation(context.Background(), typeId, locationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ProductsTypeLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsTypeLocation`: TextProductCollection
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ProductsTypeLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeId** | **string** | . | 
**locationId** | **string** | . | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductsTypeLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TextProductCollection**](TextProductCollection.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsTypeLocations

> TextProductLocationCollection ProductsTypeLocations(ctx, typeId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	typeId := "typeId_example" // string | .

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ProductsTypeLocations(context.Background(), typeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ProductsTypeLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsTypeLocations`: TextProductLocationCollection
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ProductsTypeLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeId** | **string** | . | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductsTypeLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TextProductLocationCollection**](TextProductLocationCollection.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadarProfiler

> interface{} RadarProfiler(ctx, stationId).Time(time).Interval(interval).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	stationId := "stationId_example" // string | Profiler station ID
	time := openapiclient.ISO8601Interval{String: new(string)} // ISO8601Interval | Time interval (optional)
	interval := "interval_example" // string | Averaging interval (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RadarProfiler(context.Background(), stationId).Time(time).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RadarProfiler``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RadarProfiler`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RadarProfiler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stationId** | **string** | Profiler station ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRadarProfilerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | [**ISO8601Interval**](ISO8601Interval.md) | Time interval | 
 **interval** | **string** | Averaging interval | 

### Return type

**interface{}**

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadarQueue

> interface{} RadarQueue(ctx, host).Limit(limit).Arrived(arrived).Created(created).Published(published).Station(station).Type_(type_).Feed(feed).Resolution(resolution).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	host := "host_example" // string | LDM host
	limit := int32(56) // int32 | Record limit (optional)
	arrived := openapiclient.ISO8601Interval{String: new(string)} // ISO8601Interval | Range for arrival time (optional)
	created := openapiclient.ISO8601Interval{String: new(string)} // ISO8601Interval | Range for creation time (optional)
	published := openapiclient.ISO8601Interval{String: new(string)} // ISO8601Interval | Range for publish time (optional)
	station := "station_example" // string | Station identifier (optional)
	type_ := "type__example" // string | Record type (optional)
	feed := "feed_example" // string | Originating product feed (optional)
	resolution := int32(56) // int32 | Resolution version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RadarQueue(context.Background(), host).Limit(limit).Arrived(arrived).Created(created).Published(published).Station(station).Type_(type_).Feed(feed).Resolution(resolution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RadarQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RadarQueue`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RadarQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**host** | **string** | LDM host | 

### Other Parameters

Other parameters are passed through a pointer to a apiRadarQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Record limit | 
 **arrived** | [**ISO8601Interval**](ISO8601Interval.md) | Range for arrival time | 
 **created** | [**ISO8601Interval**](ISO8601Interval.md) | Range for creation time | 
 **published** | [**ISO8601Interval**](ISO8601Interval.md) | Range for publish time | 
 **station** | **string** | Station identifier | 
 **type_** | **string** | Record type | 
 **feed** | **string** | Originating product feed | 
 **resolution** | **int32** | Resolution version | 

### Return type

**interface{}**

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadarServer

> interface{} RadarServer(ctx, id).ReportingHost(reportingHost).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	id := "id_example" // string | Server ID
	reportingHost := "reportingHost_example" // string | Show records from specific reporting host (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RadarServer(context.Background(), id).ReportingHost(reportingHost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RadarServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RadarServer`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RadarServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRadarServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportingHost** | **string** | Show records from specific reporting host | 

### Return type

**interface{}**

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadarServers

> interface{} RadarServers(ctx).ReportingHost(reportingHost).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	reportingHost := "reportingHost_example" // string | Show records from specific reporting host (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RadarServers(context.Background()).ReportingHost(reportingHost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RadarServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RadarServers`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RadarServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRadarServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportingHost** | **string** | Show records from specific reporting host | 

### Return type

**interface{}**

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadarStation

> interface{} RadarStation(ctx, stationId).ReportingHost(reportingHost).Host(host).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	stationId := "stationId_example" // string | Radar station ID
	reportingHost := "reportingHost_example" // string | Show RDA and latency info from specific reporting host (optional)
	host := "host_example" // string | Show latency info from specific LDM host (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RadarStation(context.Background(), stationId).ReportingHost(reportingHost).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RadarStation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RadarStation`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RadarStation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stationId** | **string** | Radar station ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRadarStationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportingHost** | **string** | Show RDA and latency info from specific reporting host | 
 **host** | **string** | Show latency info from specific LDM host | 

### Return type

**interface{}**

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadarStationAlarms

> interface{} RadarStationAlarms(ctx, stationId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	stationId := "stationId_example" // string | Radar station ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RadarStationAlarms(context.Background(), stationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RadarStationAlarms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RadarStationAlarms`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RadarStationAlarms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stationId** | **string** | Radar station ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRadarStationAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadarStations

> interface{} RadarStations(ctx).StationType(stationType).ReportingHost(reportingHost).Host(host).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	stationType := []string{"Inner_example"} // []string | Limit results to a specific station type or types (optional)
	reportingHost := "reportingHost_example" // string | Show RDA and latency info from specific reporting host (optional)
	host := "host_example" // string | Show latency info from specific LDM host (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RadarStations(context.Background()).StationType(stationType).ReportingHost(reportingHost).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RadarStations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RadarStations`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RadarStations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRadarStationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stationType** | **[]string** | Limit results to a specific station type or types | 
 **reportingHost** | **string** | Show RDA and latency info from specific reporting host | 
 **host** | **string** | Show latency info from specific LDM host | 

### Return type

**interface{}**

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SatelliteThumbnails

> *os.File SatelliteThumbnails(ctx, area).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	area := "area_example" // string | .

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SatelliteThumbnails(context.Background(), area).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SatelliteThumbnails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SatelliteThumbnails`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SatelliteThumbnails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**area** | **string** | . | 

### Other Parameters

Other parameters are passed through a pointer to a apiSatelliteThumbnailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/jpeg, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Sigmet

> SigmetGeoJson Sigmet(ctx, atsu, date, time).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	atsu := "atsu_example" // string | ATSU identifier
	date := time.Now() // string | Date (YYYY-MM-DD format)
	time := "time_example" // string | Time (HHMM format). This time is always specified in UTC (Zulu) time.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Sigmet(context.Background(), atsu, date, time).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Sigmet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Sigmet`: SigmetGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Sigmet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**atsu** | **string** | ATSU identifier | 
**date** | **string** | Date (YYYY-MM-DD format) | 
**time** | **string** | Time (HHMM format). This time is always specified in UTC (Zulu) time. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSigmetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SigmetGeoJson**](SigmetGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/vnd.noaa.uswx+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SigmetQuery

> SigmetCollectionGeoJson SigmetQuery(ctx).Start(start).End(end).Date(date).Atsu(atsu).Sequence(sequence).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	start := time.Now() // time.Time | Start time (optional)
	end := time.Now() // time.Time | End time (optional)
	date := time.Now() // string | Date (YYYY-MM-DD format) (optional)
	atsu := "atsu_example" // string | ATSU identifier (optional)
	sequence := "sequence_example" // string | SIGMET sequence number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SigmetQuery(context.Background()).Start(start).End(end).Date(date).Atsu(atsu).Sequence(sequence).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SigmetQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SigmetQuery`: SigmetCollectionGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SigmetQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSigmetQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **time.Time** | Start time | 
 **end** | **time.Time** | End time | 
 **date** | **string** | Date (YYYY-MM-DD format) | 
 **atsu** | **string** | ATSU identifier | 
 **sequence** | **string** | SIGMET sequence number | 

### Return type

[**SigmetCollectionGeoJson**](SigmetCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SigmetsByATSU

> SigmetCollectionGeoJson SigmetsByATSU(ctx, atsu).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	atsu := "atsu_example" // string | ATSU identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SigmetsByATSU(context.Background(), atsu).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SigmetsByATSU``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SigmetsByATSU`: SigmetCollectionGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SigmetsByATSU`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**atsu** | **string** | ATSU identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSigmetsByATSURequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SigmetCollectionGeoJson**](SigmetCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SigmetsByATSUByDate

> SigmetCollectionGeoJson SigmetsByATSUByDate(ctx, atsu, date).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	atsu := "atsu_example" // string | ATSU identifier
	date := time.Now() // string | Date (YYYY-MM-DD format)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SigmetsByATSUByDate(context.Background(), atsu, date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SigmetsByATSUByDate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SigmetsByATSUByDate`: SigmetCollectionGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SigmetsByATSUByDate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**atsu** | **string** | ATSU identifier | 
**date** | **string** | Date (YYYY-MM-DD format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSigmetsByATSUByDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SigmetCollectionGeoJson**](SigmetCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StationObservationLatest

> ObservationGeoJson StationObservationLatest(ctx, stationId).RequireQc(requireQc).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	stationId := "stationId_example" // string | Observation station ID
	requireQc := true // bool | Require QC (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.StationObservationLatest(context.Background(), stationId).RequireQc(requireQc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.StationObservationLatest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StationObservationLatest`: ObservationGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.StationObservationLatest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stationId** | **string** | Observation station ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStationObservationLatestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requireQc** | **bool** | Require QC | 

### Return type

[**ObservationGeoJson**](ObservationGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/vnd.noaa.uswx+xml, application/vnd.noaa.obs+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StationObservationList

> ObservationCollectionGeoJson StationObservationList(ctx, stationId).Start(start).End(end).Limit(limit).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	stationId := "stationId_example" // string | Observation station ID
	start := time.Now() // time.Time | Start time (optional)
	end := time.Now() // time.Time | End time (optional)
	limit := int32(56) // int32 | Limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.StationObservationList(context.Background(), stationId).Start(start).End(end).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.StationObservationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StationObservationList`: ObservationCollectionGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.StationObservationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stationId** | **string** | Observation station ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStationObservationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **time.Time** | Start time | 
 **end** | **time.Time** | End time | 
 **limit** | **int32** | Limit | 

### Return type

[**ObservationCollectionGeoJson**](ObservationCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StationObservationTime

> ObservationGeoJson StationObservationTime(ctx, stationId, time).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	stationId := "stationId_example" // string | Observation station ID
	time := time.Now() // time.Time | Timestamp of requested observation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.StationObservationTime(context.Background(), stationId, time).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.StationObservationTime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StationObservationTime`: ObservationGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.StationObservationTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stationId** | **string** | Observation station ID | 
**time** | **time.Time** | Timestamp of requested observation | 

### Other Parameters

Other parameters are passed through a pointer to a apiStationObservationTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ObservationGeoJson**](ObservationGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/vnd.noaa.uswx+xml, application/vnd.noaa.obs+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Taf

> interface{} Taf(ctx, stationId, date, time).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	stationId := "stationId_example" // string | Observation station ID
	date := time.Now() // string | Date (YYYY-MM-DD format)
	time := "time_example" // string | Time (HHMM format). This time is always specified in UTC (Zulu) time.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Taf(context.Background(), stationId, date, time).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Taf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Taf`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Taf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stationId** | **string** | Observation station ID | 
**date** | **string** | Date (YYYY-MM-DD format) | 
**time** | **string** | Time (HHMM format). This time is always specified in UTC (Zulu) time. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTafRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**interface{}**

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.wmo.iwxxm+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Tafs

> interface{} Tafs(ctx, stationId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	stationId := "stationId_example" // string | Observation station ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Tafs(context.Background(), stationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Tafs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Tafs`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Tafs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stationId** | **string** | Observation station ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTafsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Zone

> ZoneGeoJson Zone(ctx, type_, zoneId).Effective(effective).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	type_ := openapiclient.NWSZoneType("land") // NWSZoneType | Zone type
	zoneId := "zoneId_example" // string | NWS public zone/county identifier
	effective := time.Now() // time.Time | Effective date/time (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Zone(context.Background(), type_, zoneId).Effective(effective).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Zone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Zone`: ZoneGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Zone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**NWSZoneType**](.md) | Zone type | 
**zoneId** | **string** | NWS public zone/county identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **effective** | **time.Time** | Effective date/time | 

### Return type

[**ZoneGeoJson**](ZoneGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneForecast

> ZoneForecastGeoJson ZoneForecast(ctx, type_, zoneId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	type_ := "type__example" // string | Zone type
	zoneId := "zoneId_example" // string | NWS public zone/county identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ZoneForecast(context.Background(), type_, zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ZoneForecast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ZoneForecast`: ZoneForecastGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ZoneForecast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Zone type | 
**zoneId** | **string** | NWS public zone/county identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiZoneForecastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ZoneForecastGeoJson**](ZoneForecastGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneList

> ZoneCollectionGeoJson ZoneList(ctx).Id(id).Area(area).Region(region).Type_(type_).Point(point).IncludeGeometry(includeGeometry).Limit(limit).Effective(effective).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	id := []string{"Inner_example"} // []string | Zone ID (forecast or county) (optional)
	area := []openapiclient.AreaCode{openapiclient.AreaCode{MarineAreaCode: openapiclient.MarineAreaCode("AM")}} // []AreaCode | State/marine area code (optional)
	region := []openapiclient.RegionCode{openapiclient.RegionCode{LandRegionCode: openapiclient.LandRegionCode("AR")}} // []RegionCode | Region code (optional)
	type_ := []openapiclient.NWSZoneType{openapiclient.NWSZoneType("land")} // []NWSZoneType | Zone type (optional)
	point := "point_example" // string | Point (latitude,longitude) (optional)
	includeGeometry := true // bool | Include geometry in results (true/false) (optional)
	limit := int32(56) // int32 | Limit (optional)
	effective := time.Now() // time.Time | Effective date/time (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ZoneList(context.Background()).Id(id).Area(area).Region(region).Type_(type_).Point(point).IncludeGeometry(includeGeometry).Limit(limit).Effective(effective).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ZoneList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ZoneList`: ZoneCollectionGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ZoneList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZoneListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | Zone ID (forecast or county) | 
 **area** | [**[]AreaCode**](AreaCode.md) | State/marine area code | 
 **region** | [**[]RegionCode**](RegionCode.md) | Region code | 
 **type_** | [**[]NWSZoneType**](NWSZoneType.md) | Zone type | 
 **point** | **string** | Point (latitude,longitude) | 
 **includeGeometry** | **bool** | Include geometry in results (true/false) | 
 **limit** | **int32** | Limit | 
 **effective** | **time.Time** | Effective date/time | 

### Return type

[**ZoneCollectionGeoJson**](ZoneCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneListType

> ZoneCollectionGeoJson ZoneListType(ctx, type_).Id(id).Area(area).Region(region).Type_2(type_2).Point(point).IncludeGeometry(includeGeometry).Limit(limit).Effective(effective).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	type_ := openapiclient.NWSZoneType("land") // NWSZoneType | Zone type
	id := []string{"Inner_example"} // []string | Zone ID (forecast or county) (optional)
	area := []openapiclient.AreaCode{openapiclient.AreaCode{MarineAreaCode: openapiclient.MarineAreaCode("AM")}} // []AreaCode | State/marine area code (optional)
	region := []openapiclient.RegionCode{openapiclient.RegionCode{LandRegionCode: openapiclient.LandRegionCode("AR")}} // []RegionCode | Region code (optional)
	type_2 := []openapiclient.NWSZoneType{openapiclient.NWSZoneType("land")} // []NWSZoneType | Zone type (optional)
	point := "point_example" // string | Point (latitude,longitude) (optional)
	includeGeometry := true // bool | Include geometry in results (true/false) (optional)
	limit := int32(56) // int32 | Limit (optional)
	effective := time.Now() // time.Time | Effective date/time (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ZoneListType(context.Background(), type_).Id(id).Area(area).Region(region).Type_2(type_2).Point(point).IncludeGeometry(includeGeometry).Limit(limit).Effective(effective).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ZoneListType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ZoneListType`: ZoneCollectionGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ZoneListType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**NWSZoneType**](.md) | Zone type | 

### Other Parameters

Other parameters are passed through a pointer to a apiZoneListTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **[]string** | Zone ID (forecast or county) | 
 **area** | [**[]AreaCode**](AreaCode.md) | State/marine area code | 
 **region** | [**[]RegionCode**](RegionCode.md) | Region code | 
 **type_2** | [**[]NWSZoneType**](NWSZoneType.md) | Zone type | 
 **point** | **string** | Point (latitude,longitude) | 
 **includeGeometry** | **bool** | Include geometry in results (true/false) | 
 **limit** | **int32** | Limit | 
 **effective** | **time.Time** | Effective date/time | 

### Return type

[**ZoneCollectionGeoJson**](ZoneCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneObs

> ObservationCollectionGeoJson ZoneObs(ctx, zoneId).Start(start).End(end).Limit(limit).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	zoneId := "zoneId_example" // string | NWS public zone/county identifier
	start := time.Now() // time.Time | Start date/time (optional)
	end := time.Now() // time.Time | End date/time (optional)
	limit := int32(56) // int32 | Limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ZoneObs(context.Background(), zoneId).Start(start).End(end).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ZoneObs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ZoneObs`: ObservationCollectionGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ZoneObs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | NWS public zone/county identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiZoneObsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **time.Time** | Start date/time | 
 **end** | **time.Time** | End date/time | 
 **limit** | **int32** | Limit | 

### Return type

[**ObservationCollectionGeoJson**](ObservationCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneStations

> ObservationStationCollectionGeoJson ZoneStations(ctx, zoneId).Limit(limit).Cursor(cursor).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drpebcak/go-nws"
)

func main() {
	zoneId := "zoneId_example" // string | NWS public zone/county identifier
	limit := int32(56) // int32 | Limit (optional) (default to 500)
	cursor := "cursor_example" // string | Pagination cursor (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ZoneStations(context.Background(), zoneId).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ZoneStations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ZoneStations`: ObservationStationCollectionGeoJson
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ZoneStations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | NWS public zone/county identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiZoneStationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit | [default to 500]
 **cursor** | **string** | Pagination cursor | 

### Return type

[**ObservationStationCollectionGeoJson**](ObservationStationCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

