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

// GridpointForecastUnits Denotes the units used in the textual portions of the forecast.
type GridpointForecastUnits string

// List of GridpointForecastUnits
const (
	US GridpointForecastUnits = "us"
	SI GridpointForecastUnits = "si"
)

// All allowed values of GridpointForecastUnits enum
var AllowedGridpointForecastUnitsEnumValues = []GridpointForecastUnits{
	"us",
	"si",
}

func (v *GridpointForecastUnits) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GridpointForecastUnits(value)
	for _, existing := range AllowedGridpointForecastUnitsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GridpointForecastUnits", value)
}

// NewGridpointForecastUnitsFromValue returns a pointer to a valid GridpointForecastUnits
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGridpointForecastUnitsFromValue(v string) (*GridpointForecastUnits, error) {
	ev := GridpointForecastUnits(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GridpointForecastUnits: valid values are %v", v, AllowedGridpointForecastUnitsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GridpointForecastUnits) IsValid() bool {
	for _, existing := range AllowedGridpointForecastUnitsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GridpointForecastUnits value
func (v GridpointForecastUnits) Ptr() *GridpointForecastUnits {
	return &v
}

type NullableGridpointForecastUnits struct {
	value *GridpointForecastUnits
	isSet bool
}

func (v NullableGridpointForecastUnits) Get() *GridpointForecastUnits {
	return v.value
}

func (v *NullableGridpointForecastUnits) Set(val *GridpointForecastUnits) {
	v.value = val
	v.isSet = true
}

func (v NullableGridpointForecastUnits) IsSet() bool {
	return v.isSet
}

func (v *NullableGridpointForecastUnits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGridpointForecastUnits(val *GridpointForecastUnits) *NullableGridpointForecastUnits {
	return &NullableGridpointForecastUnits{value: val, isSet: true}
}

func (v NullableGridpointForecastUnits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGridpointForecastUnits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

