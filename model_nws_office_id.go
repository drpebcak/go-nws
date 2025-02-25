/*
weather.gov API

weather.gov API

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// NWSOfficeId - struct for NWSOfficeId
type NWSOfficeId struct {
	NWSForecastOfficeId *NWSForecastOfficeId
	NWSNationalHQId *NWSNationalHQId
	NWSRegionalHQId *NWSRegionalHQId
}

// NWSForecastOfficeIdAsNWSOfficeId is a convenience function that returns NWSForecastOfficeId wrapped in NWSOfficeId
func NWSForecastOfficeIdAsNWSOfficeId(v *NWSForecastOfficeId) NWSOfficeId {
	return NWSOfficeId{
		NWSForecastOfficeId: v,
	}
}

// NWSNationalHQIdAsNWSOfficeId is a convenience function that returns NWSNationalHQId wrapped in NWSOfficeId
func NWSNationalHQIdAsNWSOfficeId(v *NWSNationalHQId) NWSOfficeId {
	return NWSOfficeId{
		NWSNationalHQId: v,
	}
}

// NWSRegionalHQIdAsNWSOfficeId is a convenience function that returns NWSRegionalHQId wrapped in NWSOfficeId
func NWSRegionalHQIdAsNWSOfficeId(v *NWSRegionalHQId) NWSOfficeId {
	return NWSOfficeId{
		NWSRegionalHQId: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *NWSOfficeId) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NWSForecastOfficeId
	err = newStrictDecoder(data).Decode(&dst.NWSForecastOfficeId)
	if err == nil {
		jsonNWSForecastOfficeId, _ := json.Marshal(dst.NWSForecastOfficeId)
		if string(jsonNWSForecastOfficeId) == "{}" { // empty struct
			dst.NWSForecastOfficeId = nil
		} else {
			if err = validator.Validate(dst.NWSForecastOfficeId); err != nil {
				dst.NWSForecastOfficeId = nil
			} else {
				match++
			}
		}
	} else {
		dst.NWSForecastOfficeId = nil
	}

	// try to unmarshal data into NWSNationalHQId
	err = newStrictDecoder(data).Decode(&dst.NWSNationalHQId)
	if err == nil {
		jsonNWSNationalHQId, _ := json.Marshal(dst.NWSNationalHQId)
		if string(jsonNWSNationalHQId) == "{}" { // empty struct
			dst.NWSNationalHQId = nil
		} else {
			if err = validator.Validate(dst.NWSNationalHQId); err != nil {
				dst.NWSNationalHQId = nil
			} else {
				match++
			}
		}
	} else {
		dst.NWSNationalHQId = nil
	}

	// try to unmarshal data into NWSRegionalHQId
	err = newStrictDecoder(data).Decode(&dst.NWSRegionalHQId)
	if err == nil {
		jsonNWSRegionalHQId, _ := json.Marshal(dst.NWSRegionalHQId)
		if string(jsonNWSRegionalHQId) == "{}" { // empty struct
			dst.NWSRegionalHQId = nil
		} else {
			if err = validator.Validate(dst.NWSRegionalHQId); err != nil {
				dst.NWSRegionalHQId = nil
			} else {
				match++
			}
		}
	} else {
		dst.NWSRegionalHQId = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.NWSForecastOfficeId = nil
		dst.NWSNationalHQId = nil
		dst.NWSRegionalHQId = nil

		return fmt.Errorf("data matches more than one schema in oneOf(NWSOfficeId)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(NWSOfficeId)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NWSOfficeId) MarshalJSON() ([]byte, error) {
	if src.NWSForecastOfficeId != nil {
		return json.Marshal(&src.NWSForecastOfficeId)
	}

	if src.NWSNationalHQId != nil {
		return json.Marshal(&src.NWSNationalHQId)
	}

	if src.NWSRegionalHQId != nil {
		return json.Marshal(&src.NWSRegionalHQId)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NWSOfficeId) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.NWSForecastOfficeId != nil {
		return obj.NWSForecastOfficeId
	}

	if obj.NWSNationalHQId != nil {
		return obj.NWSNationalHQId
	}

	if obj.NWSRegionalHQId != nil {
		return obj.NWSRegionalHQId
	}

	// all schemas are nil
	return nil
}

type NullableNWSOfficeId struct {
	value *NWSOfficeId
	isSet bool
}

func (v NullableNWSOfficeId) Get() *NWSOfficeId {
	return v.value
}

func (v *NullableNWSOfficeId) Set(val *NWSOfficeId) {
	v.value = val
	v.isSet = true
}

func (v NullableNWSOfficeId) IsSet() bool {
	return v.isSet
}

func (v *NullableNWSOfficeId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNWSOfficeId(val *NWSOfficeId) *NullableNWSOfficeId {
	return &NullableNWSOfficeId{value: val, isSet: true}
}

func (v NullableNWSOfficeId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNWSOfficeId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


