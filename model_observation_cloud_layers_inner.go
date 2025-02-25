/*
weather.gov API

weather.gov API

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ObservationCloudLayersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObservationCloudLayersInner{}

// ObservationCloudLayersInner struct for ObservationCloudLayersInner
type ObservationCloudLayersInner struct {
	Base QuantitativeValue `json:"base"`
	Amount MetarSkyCoverage `json:"amount"`
}

type _ObservationCloudLayersInner ObservationCloudLayersInner

// NewObservationCloudLayersInner instantiates a new ObservationCloudLayersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObservationCloudLayersInner(base QuantitativeValue, amount MetarSkyCoverage) *ObservationCloudLayersInner {
	this := ObservationCloudLayersInner{}
	this.Base = base
	this.Amount = amount
	return &this
}

// NewObservationCloudLayersInnerWithDefaults instantiates a new ObservationCloudLayersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObservationCloudLayersInnerWithDefaults() *ObservationCloudLayersInner {
	this := ObservationCloudLayersInner{}
	return &this
}

// GetBase returns the Base field value
func (o *ObservationCloudLayersInner) GetBase() QuantitativeValue {
	if o == nil {
		var ret QuantitativeValue
		return ret
	}

	return o.Base
}

// GetBaseOk returns a tuple with the Base field value
// and a boolean to check if the value has been set.
func (o *ObservationCloudLayersInner) GetBaseOk() (*QuantitativeValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Base, true
}

// SetBase sets field value
func (o *ObservationCloudLayersInner) SetBase(v QuantitativeValue) {
	o.Base = v
}

// GetAmount returns the Amount field value
func (o *ObservationCloudLayersInner) GetAmount() MetarSkyCoverage {
	if o == nil {
		var ret MetarSkyCoverage
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ObservationCloudLayersInner) GetAmountOk() (*MetarSkyCoverage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ObservationCloudLayersInner) SetAmount(v MetarSkyCoverage) {
	o.Amount = v
}

func (o ObservationCloudLayersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObservationCloudLayersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["base"] = o.Base
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

func (o *ObservationCloudLayersInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"base",
		"amount",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varObservationCloudLayersInner := _ObservationCloudLayersInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varObservationCloudLayersInner)

	if err != nil {
		return err
	}

	*o = ObservationCloudLayersInner(varObservationCloudLayersInner)

	return err
}

type NullableObservationCloudLayersInner struct {
	value *ObservationCloudLayersInner
	isSet bool
}

func (v NullableObservationCloudLayersInner) Get() *ObservationCloudLayersInner {
	return v.value
}

func (v *NullableObservationCloudLayersInner) Set(val *ObservationCloudLayersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableObservationCloudLayersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableObservationCloudLayersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservationCloudLayersInner(val *ObservationCloudLayersInner) *NullableObservationCloudLayersInner {
	return &NullableObservationCloudLayersInner{value: val, isSet: true}
}

func (v NullableObservationCloudLayersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservationCloudLayersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


