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

// checks if the RelativeLocationJsonLd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelativeLocationJsonLd{}

// RelativeLocationJsonLd struct for RelativeLocationJsonLd
type RelativeLocationJsonLd struct {
	City *string `json:"city,omitempty"`
	State *string `json:"state,omitempty"`
	Distance *QuantitativeValue `json:"distance,omitempty"`
	Bearing *QuantitativeValue `json:"bearing,omitempty"`
	// A geometry represented in Well-Known Text (WKT) format.
	Geometry NullableString `json:"geometry"`
}

type _RelativeLocationJsonLd RelativeLocationJsonLd

// NewRelativeLocationJsonLd instantiates a new RelativeLocationJsonLd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelativeLocationJsonLd(geometry NullableString) *RelativeLocationJsonLd {
	this := RelativeLocationJsonLd{}
	this.Geometry = geometry
	return &this
}

// NewRelativeLocationJsonLdWithDefaults instantiates a new RelativeLocationJsonLd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelativeLocationJsonLdWithDefaults() *RelativeLocationJsonLd {
	this := RelativeLocationJsonLd{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *RelativeLocationJsonLd) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelativeLocationJsonLd) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *RelativeLocationJsonLd) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *RelativeLocationJsonLd) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RelativeLocationJsonLd) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelativeLocationJsonLd) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RelativeLocationJsonLd) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *RelativeLocationJsonLd) SetState(v string) {
	o.State = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *RelativeLocationJsonLd) GetDistance() QuantitativeValue {
	if o == nil || IsNil(o.Distance) {
		var ret QuantitativeValue
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelativeLocationJsonLd) GetDistanceOk() (*QuantitativeValue, bool) {
	if o == nil || IsNil(o.Distance) {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *RelativeLocationJsonLd) HasDistance() bool {
	if o != nil && !IsNil(o.Distance) {
		return true
	}

	return false
}

// SetDistance gets a reference to the given QuantitativeValue and assigns it to the Distance field.
func (o *RelativeLocationJsonLd) SetDistance(v QuantitativeValue) {
	o.Distance = &v
}

// GetBearing returns the Bearing field value if set, zero value otherwise.
func (o *RelativeLocationJsonLd) GetBearing() QuantitativeValue {
	if o == nil || IsNil(o.Bearing) {
		var ret QuantitativeValue
		return ret
	}
	return *o.Bearing
}

// GetBearingOk returns a tuple with the Bearing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelativeLocationJsonLd) GetBearingOk() (*QuantitativeValue, bool) {
	if o == nil || IsNil(o.Bearing) {
		return nil, false
	}
	return o.Bearing, true
}

// HasBearing returns a boolean if a field has been set.
func (o *RelativeLocationJsonLd) HasBearing() bool {
	if o != nil && !IsNil(o.Bearing) {
		return true
	}

	return false
}

// SetBearing gets a reference to the given QuantitativeValue and assigns it to the Bearing field.
func (o *RelativeLocationJsonLd) SetBearing(v QuantitativeValue) {
	o.Bearing = &v
}

// GetGeometry returns the Geometry field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RelativeLocationJsonLd) GetGeometry() string {
	if o == nil || o.Geometry.Get() == nil {
		var ret string
		return ret
	}

	return *o.Geometry.Get()
}

// GetGeometryOk returns a tuple with the Geometry field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RelativeLocationJsonLd) GetGeometryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Geometry.Get(), o.Geometry.IsSet()
}

// SetGeometry sets field value
func (o *RelativeLocationJsonLd) SetGeometry(v string) {
	o.Geometry.Set(&v)
}

func (o RelativeLocationJsonLd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelativeLocationJsonLd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Distance) {
		toSerialize["distance"] = o.Distance
	}
	if !IsNil(o.Bearing) {
		toSerialize["bearing"] = o.Bearing
	}
	toSerialize["geometry"] = o.Geometry.Get()
	return toSerialize, nil
}

func (o *RelativeLocationJsonLd) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"geometry",
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

	varRelativeLocationJsonLd := _RelativeLocationJsonLd{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRelativeLocationJsonLd)

	if err != nil {
		return err
	}

	*o = RelativeLocationJsonLd(varRelativeLocationJsonLd)

	return err
}

type NullableRelativeLocationJsonLd struct {
	value *RelativeLocationJsonLd
	isSet bool
}

func (v NullableRelativeLocationJsonLd) Get() *RelativeLocationJsonLd {
	return v.value
}

func (v *NullableRelativeLocationJsonLd) Set(val *RelativeLocationJsonLd) {
	v.value = val
	v.isSet = true
}

func (v NullableRelativeLocationJsonLd) IsSet() bool {
	return v.isSet
}

func (v *NullableRelativeLocationJsonLd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelativeLocationJsonLd(val *RelativeLocationJsonLd) *NullableRelativeLocationJsonLd {
	return &NullableRelativeLocationJsonLd{value: val, isSet: true}
}

func (v NullableRelativeLocationJsonLd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelativeLocationJsonLd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


