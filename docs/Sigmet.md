# Sigmet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**IssueTime** | Pointer to **time.Time** |  | [optional] 
**Fir** | Pointer to **NullableString** |  | [optional] 
**Atsu** | Pointer to **string** | ATSU Identifier | [optional] 
**Sequence** | Pointer to **NullableString** |  | [optional] 
**Phenomenon** | Pointer to **NullableString** |  | [optional] 
**Start** | Pointer to **time.Time** |  | [optional] 
**End** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSigmet

`func NewSigmet() *Sigmet`

NewSigmet instantiates a new Sigmet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigmetWithDefaults

`func NewSigmetWithDefaults() *Sigmet`

NewSigmetWithDefaults instantiates a new Sigmet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Sigmet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Sigmet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Sigmet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Sigmet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueTime

`func (o *Sigmet) GetIssueTime() time.Time`

GetIssueTime returns the IssueTime field if non-nil, zero value otherwise.

### GetIssueTimeOk

`func (o *Sigmet) GetIssueTimeOk() (*time.Time, bool)`

GetIssueTimeOk returns a tuple with the IssueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTime

`func (o *Sigmet) SetIssueTime(v time.Time)`

SetIssueTime sets IssueTime field to given value.

### HasIssueTime

`func (o *Sigmet) HasIssueTime() bool`

HasIssueTime returns a boolean if a field has been set.

### GetFir

`func (o *Sigmet) GetFir() string`

GetFir returns the Fir field if non-nil, zero value otherwise.

### GetFirOk

`func (o *Sigmet) GetFirOk() (*string, bool)`

GetFirOk returns a tuple with the Fir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFir

`func (o *Sigmet) SetFir(v string)`

SetFir sets Fir field to given value.

### HasFir

`func (o *Sigmet) HasFir() bool`

HasFir returns a boolean if a field has been set.

### SetFirNil

`func (o *Sigmet) SetFirNil(b bool)`

 SetFirNil sets the value for Fir to be an explicit nil

### UnsetFir
`func (o *Sigmet) UnsetFir()`

UnsetFir ensures that no value is present for Fir, not even an explicit nil
### GetAtsu

`func (o *Sigmet) GetAtsu() string`

GetAtsu returns the Atsu field if non-nil, zero value otherwise.

### GetAtsuOk

`func (o *Sigmet) GetAtsuOk() (*string, bool)`

GetAtsuOk returns a tuple with the Atsu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtsu

`func (o *Sigmet) SetAtsu(v string)`

SetAtsu sets Atsu field to given value.

### HasAtsu

`func (o *Sigmet) HasAtsu() bool`

HasAtsu returns a boolean if a field has been set.

### GetSequence

`func (o *Sigmet) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *Sigmet) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *Sigmet) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *Sigmet) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### SetSequenceNil

`func (o *Sigmet) SetSequenceNil(b bool)`

 SetSequenceNil sets the value for Sequence to be an explicit nil

### UnsetSequence
`func (o *Sigmet) UnsetSequence()`

UnsetSequence ensures that no value is present for Sequence, not even an explicit nil
### GetPhenomenon

`func (o *Sigmet) GetPhenomenon() string`

GetPhenomenon returns the Phenomenon field if non-nil, zero value otherwise.

### GetPhenomenonOk

`func (o *Sigmet) GetPhenomenonOk() (*string, bool)`

GetPhenomenonOk returns a tuple with the Phenomenon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhenomenon

`func (o *Sigmet) SetPhenomenon(v string)`

SetPhenomenon sets Phenomenon field to given value.

### HasPhenomenon

`func (o *Sigmet) HasPhenomenon() bool`

HasPhenomenon returns a boolean if a field has been set.

### SetPhenomenonNil

`func (o *Sigmet) SetPhenomenonNil(b bool)`

 SetPhenomenonNil sets the value for Phenomenon to be an explicit nil

### UnsetPhenomenon
`func (o *Sigmet) UnsetPhenomenon()`

UnsetPhenomenon ensures that no value is present for Phenomenon, not even an explicit nil
### GetStart

`func (o *Sigmet) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Sigmet) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Sigmet) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *Sigmet) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *Sigmet) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Sigmet) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Sigmet) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Sigmet) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


