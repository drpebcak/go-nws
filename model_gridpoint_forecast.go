/*
weather.gov API

weather.gov API

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the GridpointForecast type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GridpointForecast{}

// GridpointForecast A multi-day forecast for a 2.5km grid square.
type GridpointForecast struct {
	Context *JsonLdContext `json:"@context,omitempty"`
	// A geometry represented in Well-Known Text (WKT) format.
	Geometry NullableString `json:"geometry,omitempty"`
	Units *GridpointForecastUnits `json:"units,omitempty"`
	// The internal generator class used to create the forecast text (used for NWS debugging).
	ForecastGenerator *string `json:"forecastGenerator,omitempty"`
	// The time this forecast data was generated.
	GeneratedAt *time.Time `json:"generatedAt,omitempty"`
	// The last update time of the data this forecast was generated from.
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	ValidTimes *ISO8601Interval `json:"validTimes,omitempty"`
	Elevation *QuantitativeValue `json:"elevation,omitempty"`
	// An array of forecast periods.
	Periods []GridpointForecastPeriod `json:"periods,omitempty"`
}

// NewGridpointForecast instantiates a new GridpointForecast object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGridpointForecast() *GridpointForecast {
	this := GridpointForecast{}
	var units GridpointForecastUnits = US
	this.Units = &units
	return &this
}

// NewGridpointForecastWithDefaults instantiates a new GridpointForecast object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridpointForecastWithDefaults() *GridpointForecast {
	this := GridpointForecast{}
	var units GridpointForecastUnits = US
	this.Units = &units
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *GridpointForecast) GetContext() JsonLdContext {
	if o == nil || IsNil(o.Context) {
		var ret JsonLdContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecast) GetContextOk() (*JsonLdContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *GridpointForecast) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given JsonLdContext and assigns it to the Context field.
func (o *GridpointForecast) SetContext(v JsonLdContext) {
	o.Context = &v
}

// GetGeometry returns the Geometry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GridpointForecast) GetGeometry() string {
	if o == nil || IsNil(o.Geometry.Get()) {
		var ret string
		return ret
	}
	return *o.Geometry.Get()
}

// GetGeometryOk returns a tuple with the Geometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GridpointForecast) GetGeometryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Geometry.Get(), o.Geometry.IsSet()
}

// HasGeometry returns a boolean if a field has been set.
func (o *GridpointForecast) HasGeometry() bool {
	if o != nil && o.Geometry.IsSet() {
		return true
	}

	return false
}

// SetGeometry gets a reference to the given NullableString and assigns it to the Geometry field.
func (o *GridpointForecast) SetGeometry(v string) {
	o.Geometry.Set(&v)
}
// SetGeometryNil sets the value for Geometry to be an explicit nil
func (o *GridpointForecast) SetGeometryNil() {
	o.Geometry.Set(nil)
}

// UnsetGeometry ensures that no value is present for Geometry, not even an explicit nil
func (o *GridpointForecast) UnsetGeometry() {
	o.Geometry.Unset()
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *GridpointForecast) GetUnits() GridpointForecastUnits {
	if o == nil || IsNil(o.Units) {
		var ret GridpointForecastUnits
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecast) GetUnitsOk() (*GridpointForecastUnits, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *GridpointForecast) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given GridpointForecastUnits and assigns it to the Units field.
func (o *GridpointForecast) SetUnits(v GridpointForecastUnits) {
	o.Units = &v
}

// GetForecastGenerator returns the ForecastGenerator field value if set, zero value otherwise.
func (o *GridpointForecast) GetForecastGenerator() string {
	if o == nil || IsNil(o.ForecastGenerator) {
		var ret string
		return ret
	}
	return *o.ForecastGenerator
}

// GetForecastGeneratorOk returns a tuple with the ForecastGenerator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecast) GetForecastGeneratorOk() (*string, bool) {
	if o == nil || IsNil(o.ForecastGenerator) {
		return nil, false
	}
	return o.ForecastGenerator, true
}

// HasForecastGenerator returns a boolean if a field has been set.
func (o *GridpointForecast) HasForecastGenerator() bool {
	if o != nil && !IsNil(o.ForecastGenerator) {
		return true
	}

	return false
}

// SetForecastGenerator gets a reference to the given string and assigns it to the ForecastGenerator field.
func (o *GridpointForecast) SetForecastGenerator(v string) {
	o.ForecastGenerator = &v
}

// GetGeneratedAt returns the GeneratedAt field value if set, zero value otherwise.
func (o *GridpointForecast) GetGeneratedAt() time.Time {
	if o == nil || IsNil(o.GeneratedAt) {
		var ret time.Time
		return ret
	}
	return *o.GeneratedAt
}

// GetGeneratedAtOk returns a tuple with the GeneratedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecast) GetGeneratedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.GeneratedAt) {
		return nil, false
	}
	return o.GeneratedAt, true
}

// HasGeneratedAt returns a boolean if a field has been set.
func (o *GridpointForecast) HasGeneratedAt() bool {
	if o != nil && !IsNil(o.GeneratedAt) {
		return true
	}

	return false
}

// SetGeneratedAt gets a reference to the given time.Time and assigns it to the GeneratedAt field.
func (o *GridpointForecast) SetGeneratedAt(v time.Time) {
	o.GeneratedAt = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GridpointForecast) GetUpdateTime() time.Time {
	if o == nil || IsNil(o.UpdateTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecast) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GridpointForecast) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *GridpointForecast) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

// GetValidTimes returns the ValidTimes field value if set, zero value otherwise.
func (o *GridpointForecast) GetValidTimes() ISO8601Interval {
	if o == nil || IsNil(o.ValidTimes) {
		var ret ISO8601Interval
		return ret
	}
	return *o.ValidTimes
}

// GetValidTimesOk returns a tuple with the ValidTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecast) GetValidTimesOk() (*ISO8601Interval, bool) {
	if o == nil || IsNil(o.ValidTimes) {
		return nil, false
	}
	return o.ValidTimes, true
}

// HasValidTimes returns a boolean if a field has been set.
func (o *GridpointForecast) HasValidTimes() bool {
	if o != nil && !IsNil(o.ValidTimes) {
		return true
	}

	return false
}

// SetValidTimes gets a reference to the given ISO8601Interval and assigns it to the ValidTimes field.
func (o *GridpointForecast) SetValidTimes(v ISO8601Interval) {
	o.ValidTimes = &v
}

// GetElevation returns the Elevation field value if set, zero value otherwise.
func (o *GridpointForecast) GetElevation() QuantitativeValue {
	if o == nil || IsNil(o.Elevation) {
		var ret QuantitativeValue
		return ret
	}
	return *o.Elevation
}

// GetElevationOk returns a tuple with the Elevation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecast) GetElevationOk() (*QuantitativeValue, bool) {
	if o == nil || IsNil(o.Elevation) {
		return nil, false
	}
	return o.Elevation, true
}

// HasElevation returns a boolean if a field has been set.
func (o *GridpointForecast) HasElevation() bool {
	if o != nil && !IsNil(o.Elevation) {
		return true
	}

	return false
}

// SetElevation gets a reference to the given QuantitativeValue and assigns it to the Elevation field.
func (o *GridpointForecast) SetElevation(v QuantitativeValue) {
	o.Elevation = &v
}

// GetPeriods returns the Periods field value if set, zero value otherwise.
func (o *GridpointForecast) GetPeriods() []GridpointForecastPeriod {
	if o == nil || IsNil(o.Periods) {
		var ret []GridpointForecastPeriod
		return ret
	}
	return o.Periods
}

// GetPeriodsOk returns a tuple with the Periods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecast) GetPeriodsOk() ([]GridpointForecastPeriod, bool) {
	if o == nil || IsNil(o.Periods) {
		return nil, false
	}
	return o.Periods, true
}

// HasPeriods returns a boolean if a field has been set.
func (o *GridpointForecast) HasPeriods() bool {
	if o != nil && !IsNil(o.Periods) {
		return true
	}

	return false
}

// SetPeriods gets a reference to the given []GridpointForecastPeriod and assigns it to the Periods field.
func (o *GridpointForecast) SetPeriods(v []GridpointForecastPeriod) {
	o.Periods = v
}

func (o GridpointForecast) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GridpointForecast) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["@context"] = o.Context
	}
	if o.Geometry.IsSet() {
		toSerialize["geometry"] = o.Geometry.Get()
	}
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	if !IsNil(o.ForecastGenerator) {
		toSerialize["forecastGenerator"] = o.ForecastGenerator
	}
	if !IsNil(o.GeneratedAt) {
		toSerialize["generatedAt"] = o.GeneratedAt
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !IsNil(o.ValidTimes) {
		toSerialize["validTimes"] = o.ValidTimes
	}
	if !IsNil(o.Elevation) {
		toSerialize["elevation"] = o.Elevation
	}
	if !IsNil(o.Periods) {
		toSerialize["periods"] = o.Periods
	}
	return toSerialize, nil
}

type NullableGridpointForecast struct {
	value *GridpointForecast
	isSet bool
}

func (v NullableGridpointForecast) Get() *GridpointForecast {
	return v.value
}

func (v *NullableGridpointForecast) Set(val *GridpointForecast) {
	v.value = val
	v.isSet = true
}

func (v NullableGridpointForecast) IsSet() bool {
	return v.isSet
}

func (v *NullableGridpointForecast) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGridpointForecast(val *GridpointForecast) *NullableGridpointForecast {
	return &NullableGridpointForecast{value: val, isSet: true}
}

func (v NullableGridpointForecast) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGridpointForecast) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


