# CenterWeatherAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**IssueTime** | Pointer to **time.Time** |  | [optional] 
**Cwsu** | Pointer to [**NWSCenterWeatherServiceUnitId**](NWSCenterWeatherServiceUnitId.md) |  | [optional] 
**Sequence** | Pointer to **int32** |  | [optional] 
**Start** | Pointer to **time.Time** |  | [optional] 
**End** | Pointer to **time.Time** |  | [optional] 
**ObservedProperty** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 

## Methods

### NewCenterWeatherAdvisory

`func NewCenterWeatherAdvisory() *CenterWeatherAdvisory`

NewCenterWeatherAdvisory instantiates a new CenterWeatherAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCenterWeatherAdvisoryWithDefaults

`func NewCenterWeatherAdvisoryWithDefaults() *CenterWeatherAdvisory`

NewCenterWeatherAdvisoryWithDefaults instantiates a new CenterWeatherAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CenterWeatherAdvisory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CenterWeatherAdvisory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CenterWeatherAdvisory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CenterWeatherAdvisory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueTime

`func (o *CenterWeatherAdvisory) GetIssueTime() time.Time`

GetIssueTime returns the IssueTime field if non-nil, zero value otherwise.

### GetIssueTimeOk

`func (o *CenterWeatherAdvisory) GetIssueTimeOk() (*time.Time, bool)`

GetIssueTimeOk returns a tuple with the IssueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTime

`func (o *CenterWeatherAdvisory) SetIssueTime(v time.Time)`

SetIssueTime sets IssueTime field to given value.

### HasIssueTime

`func (o *CenterWeatherAdvisory) HasIssueTime() bool`

HasIssueTime returns a boolean if a field has been set.

### GetCwsu

`func (o *CenterWeatherAdvisory) GetCwsu() NWSCenterWeatherServiceUnitId`

GetCwsu returns the Cwsu field if non-nil, zero value otherwise.

### GetCwsuOk

`func (o *CenterWeatherAdvisory) GetCwsuOk() (*NWSCenterWeatherServiceUnitId, bool)`

GetCwsuOk returns a tuple with the Cwsu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwsu

`func (o *CenterWeatherAdvisory) SetCwsu(v NWSCenterWeatherServiceUnitId)`

SetCwsu sets Cwsu field to given value.

### HasCwsu

`func (o *CenterWeatherAdvisory) HasCwsu() bool`

HasCwsu returns a boolean if a field has been set.

### GetSequence

`func (o *CenterWeatherAdvisory) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *CenterWeatherAdvisory) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *CenterWeatherAdvisory) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *CenterWeatherAdvisory) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetStart

`func (o *CenterWeatherAdvisory) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CenterWeatherAdvisory) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CenterWeatherAdvisory) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *CenterWeatherAdvisory) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *CenterWeatherAdvisory) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CenterWeatherAdvisory) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CenterWeatherAdvisory) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CenterWeatherAdvisory) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetObservedProperty

`func (o *CenterWeatherAdvisory) GetObservedProperty() string`

GetObservedProperty returns the ObservedProperty field if non-nil, zero value otherwise.

### GetObservedPropertyOk

`func (o *CenterWeatherAdvisory) GetObservedPropertyOk() (*string, bool)`

GetObservedPropertyOk returns a tuple with the ObservedProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedProperty

`func (o *CenterWeatherAdvisory) SetObservedProperty(v string)`

SetObservedProperty sets ObservedProperty field to given value.

### HasObservedProperty

`func (o *CenterWeatherAdvisory) HasObservedProperty() bool`

HasObservedProperty returns a boolean if a field has been set.

### GetText

`func (o *CenterWeatherAdvisory) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CenterWeatherAdvisory) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CenterWeatherAdvisory) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CenterWeatherAdvisory) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


