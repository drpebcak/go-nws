/*
weather.gov API

weather.gov API

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)


// IconsSizeParameter struct for IconsSizeParameter
type IconsSizeParameter struct {
	Int32 *int32
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *IconsSizeParameter) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Int32
	err = json.Unmarshal(data, &dst.Int32);
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			return nil // data stored in dst.Int32, return on the first match
		}
	} else {
		dst.Int32 = nil
	}

	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(IconsSizeParameter)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IconsSizeParameter) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableIconsSizeParameter struct {
	value *IconsSizeParameter
	isSet bool
}

func (v NullableIconsSizeParameter) Get() *IconsSizeParameter {
	return v.value
}

func (v *NullableIconsSizeParameter) Set(val *IconsSizeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableIconsSizeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableIconsSizeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIconsSizeParameter(val *IconsSizeParameter) *NullableIconsSizeParameter {
	return &NullableIconsSizeParameter{value: val, isSet: true}
}

func (v NullableIconsSizeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIconsSizeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


