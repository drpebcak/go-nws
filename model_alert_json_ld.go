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

// checks if the AlertJsonLd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertJsonLd{}

// AlertJsonLd struct for AlertJsonLd
type AlertJsonLd struct {
	Graph []Alert `json:"@graph,omitempty"`
}

// NewAlertJsonLd instantiates a new AlertJsonLd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertJsonLd() *AlertJsonLd {
	this := AlertJsonLd{}
	return &this
}

// NewAlertJsonLdWithDefaults instantiates a new AlertJsonLd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertJsonLdWithDefaults() *AlertJsonLd {
	this := AlertJsonLd{}
	return &this
}

// GetGraph returns the Graph field value if set, zero value otherwise.
func (o *AlertJsonLd) GetGraph() []Alert {
	if o == nil || IsNil(o.Graph) {
		var ret []Alert
		return ret
	}
	return o.Graph
}

// GetGraphOk returns a tuple with the Graph field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertJsonLd) GetGraphOk() ([]Alert, bool) {
	if o == nil || IsNil(o.Graph) {
		return nil, false
	}
	return o.Graph, true
}

// HasGraph returns a boolean if a field has been set.
func (o *AlertJsonLd) HasGraph() bool {
	if o != nil && !IsNil(o.Graph) {
		return true
	}

	return false
}

// SetGraph gets a reference to the given []Alert and assigns it to the Graph field.
func (o *AlertJsonLd) SetGraph(v []Alert) {
	o.Graph = v
}

func (o AlertJsonLd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertJsonLd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Graph) {
		toSerialize["@graph"] = o.Graph
	}
	return toSerialize, nil
}

type NullableAlertJsonLd struct {
	value *AlertJsonLd
	isSet bool
}

func (v NullableAlertJsonLd) Get() *AlertJsonLd {
	return v.value
}

func (v *NullableAlertJsonLd) Set(val *AlertJsonLd) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertJsonLd) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertJsonLd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertJsonLd(val *AlertJsonLd) *NullableAlertJsonLd {
	return &NullableAlertJsonLd{value: val, isSet: true}
}

func (v NullableAlertJsonLd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertJsonLd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


