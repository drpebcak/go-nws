# CenterWeatherAdvisoryGeoJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Geometry** | [**NullableGeoJsonGeometry**](GeoJsonGeometry.md) |  | 
**Properties** | [**CenterWeatherAdvisory**](CenterWeatherAdvisory.md) |  | 

## Methods

### NewCenterWeatherAdvisoryGeoJson

`func NewCenterWeatherAdvisoryGeoJson(type_ string, geometry NullableGeoJsonGeometry, properties CenterWeatherAdvisory, ) *CenterWeatherAdvisoryGeoJson`

NewCenterWeatherAdvisoryGeoJson instantiates a new CenterWeatherAdvisoryGeoJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCenterWeatherAdvisoryGeoJsonWithDefaults

`func NewCenterWeatherAdvisoryGeoJsonWithDefaults() *CenterWeatherAdvisoryGeoJson`

NewCenterWeatherAdvisoryGeoJsonWithDefaults instantiates a new CenterWeatherAdvisoryGeoJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *CenterWeatherAdvisoryGeoJson) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CenterWeatherAdvisoryGeoJson) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CenterWeatherAdvisoryGeoJson) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *CenterWeatherAdvisoryGeoJson) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *CenterWeatherAdvisoryGeoJson) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CenterWeatherAdvisoryGeoJson) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CenterWeatherAdvisoryGeoJson) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CenterWeatherAdvisoryGeoJson) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CenterWeatherAdvisoryGeoJson) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CenterWeatherAdvisoryGeoJson) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CenterWeatherAdvisoryGeoJson) SetType(v string)`

SetType sets Type field to given value.


### GetGeometry

`func (o *CenterWeatherAdvisoryGeoJson) GetGeometry() GeoJsonGeometry`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *CenterWeatherAdvisoryGeoJson) GetGeometryOk() (*GeoJsonGeometry, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *CenterWeatherAdvisoryGeoJson) SetGeometry(v GeoJsonGeometry)`

SetGeometry sets Geometry field to given value.


### SetGeometryNil

`func (o *CenterWeatherAdvisoryGeoJson) SetGeometryNil(b bool)`

 SetGeometryNil sets the value for Geometry to be an explicit nil

### UnsetGeometry
`func (o *CenterWeatherAdvisoryGeoJson) UnsetGeometry()`

UnsetGeometry ensures that no value is present for Geometry, not even an explicit nil
### GetProperties

`func (o *CenterWeatherAdvisoryGeoJson) GetProperties() CenterWeatherAdvisory`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CenterWeatherAdvisoryGeoJson) GetPropertiesOk() (*CenterWeatherAdvisory, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CenterWeatherAdvisoryGeoJson) SetProperties(v CenterWeatherAdvisory)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


