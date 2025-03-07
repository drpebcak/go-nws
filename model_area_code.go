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

// AreaCode - State/territory codes and marine area codes
type AreaCode struct {
	MarineAreaCode *MarineAreaCode
	StateTerritoryCode *StateTerritoryCode
}

// MarineAreaCodeAsAreaCode is a convenience function that returns MarineAreaCode wrapped in AreaCode
func MarineAreaCodeAsAreaCode(v *MarineAreaCode) AreaCode {
	return AreaCode{
		MarineAreaCode: v,
	}
}

// StateTerritoryCodeAsAreaCode is a convenience function that returns StateTerritoryCode wrapped in AreaCode
func StateTerritoryCodeAsAreaCode(v *StateTerritoryCode) AreaCode {
	return AreaCode{
		StateTerritoryCode: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AreaCode) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MarineAreaCode
	err = newStrictDecoder(data).Decode(&dst.MarineAreaCode)
	if err == nil {
		jsonMarineAreaCode, _ := json.Marshal(dst.MarineAreaCode)
		if string(jsonMarineAreaCode) == "{}" { // empty struct
			dst.MarineAreaCode = nil
		} else {
			if err = validator.Validate(dst.MarineAreaCode); err != nil {
				dst.MarineAreaCode = nil
			} else {
				match++
			}
		}
	} else {
		dst.MarineAreaCode = nil
	}

	// try to unmarshal data into StateTerritoryCode
	err = newStrictDecoder(data).Decode(&dst.StateTerritoryCode)
	if err == nil {
		jsonStateTerritoryCode, _ := json.Marshal(dst.StateTerritoryCode)
		if string(jsonStateTerritoryCode) == "{}" { // empty struct
			dst.StateTerritoryCode = nil
		} else {
			if err = validator.Validate(dst.StateTerritoryCode); err != nil {
				dst.StateTerritoryCode = nil
			} else {
				match++
			}
		}
	} else {
		dst.StateTerritoryCode = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MarineAreaCode = nil
		dst.StateTerritoryCode = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AreaCode)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AreaCode)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AreaCode) MarshalJSON() ([]byte, error) {
	if src.MarineAreaCode != nil {
		return json.Marshal(&src.MarineAreaCode)
	}

	if src.StateTerritoryCode != nil {
		return json.Marshal(&src.StateTerritoryCode)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AreaCode) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MarineAreaCode != nil {
		return obj.MarineAreaCode
	}

	if obj.StateTerritoryCode != nil {
		return obj.StateTerritoryCode
	}

	// all schemas are nil
	return nil
}

type NullableAreaCode struct {
	value *AreaCode
	isSet bool
}

func (v NullableAreaCode) Get() *AreaCode {
	return v.value
}

func (v *NullableAreaCode) Set(val *AreaCode) {
	v.value = val
	v.isSet = true
}

func (v NullableAreaCode) IsSet() bool {
	return v.isSet
}

func (v *NullableAreaCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAreaCode(val *AreaCode) *NullableAreaCode {
	return &NullableAreaCode{value: val, isSet: true}
}

func (v NullableAreaCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAreaCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


