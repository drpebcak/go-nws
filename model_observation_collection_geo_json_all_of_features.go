/*
weather.gov API

weather.gov API

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ObservationCollectionGeoJsonAllOfFeatures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObservationCollectionGeoJsonAllOfFeatures{}

// ObservationCollectionGeoJsonAllOfFeatures struct for ObservationCollectionGeoJsonAllOfFeatures
type ObservationCollectionGeoJsonAllOfFeatures struct {
	Properties *Observation `json:"properties,omitempty"`
}

// NewObservationCollectionGeoJsonAllOfFeatures instantiates a new ObservationCollectionGeoJsonAllOfFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObservationCollectionGeoJsonAllOfFeatures() *ObservationCollectionGeoJsonAllOfFeatures {
	this := ObservationCollectionGeoJsonAllOfFeatures{}
	return &this
}

// NewObservationCollectionGeoJsonAllOfFeaturesWithDefaults instantiates a new ObservationCollectionGeoJsonAllOfFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObservationCollectionGeoJsonAllOfFeaturesWithDefaults() *ObservationCollectionGeoJsonAllOfFeatures {
	this := ObservationCollectionGeoJsonAllOfFeatures{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ObservationCollectionGeoJsonAllOfFeatures) GetProperties() Observation {
	if o == nil || IsNil(o.Properties) {
		var ret Observation
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationCollectionGeoJsonAllOfFeatures) GetPropertiesOk() (*Observation, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ObservationCollectionGeoJsonAllOfFeatures) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given Observation and assigns it to the Properties field.
func (o *ObservationCollectionGeoJsonAllOfFeatures) SetProperties(v Observation) {
	o.Properties = &v
}

func (o ObservationCollectionGeoJsonAllOfFeatures) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObservationCollectionGeoJsonAllOfFeatures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableObservationCollectionGeoJsonAllOfFeatures struct {
	value *ObservationCollectionGeoJsonAllOfFeatures
	isSet bool
}

func (v NullableObservationCollectionGeoJsonAllOfFeatures) Get() *ObservationCollectionGeoJsonAllOfFeatures {
	return v.value
}

func (v *NullableObservationCollectionGeoJsonAllOfFeatures) Set(val *ObservationCollectionGeoJsonAllOfFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableObservationCollectionGeoJsonAllOfFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableObservationCollectionGeoJsonAllOfFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservationCollectionGeoJsonAllOfFeatures(val *ObservationCollectionGeoJsonAllOfFeatures) *NullableObservationCollectionGeoJsonAllOfFeatures {
	return &NullableObservationCollectionGeoJsonAllOfFeatures{value: val, isSet: true}
}

func (v NullableObservationCollectionGeoJsonAllOfFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservationCollectionGeoJsonAllOfFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


