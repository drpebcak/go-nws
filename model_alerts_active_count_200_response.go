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

// checks if the AlertsActiveCount200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertsActiveCount200Response{}

// AlertsActiveCount200Response struct for AlertsActiveCount200Response
type AlertsActiveCount200Response struct {
	// The total number of active alerts
	Total *int32 `json:"total,omitempty"`
	// The total number of active alerts affecting land zones
	Land *int32 `json:"land,omitempty"`
	// The total number of active alerts affecting marine zones
	Marine *int32 `json:"marine,omitempty"`
	// Active alerts by marine region
	Regions *map[string]int32 `json:"regions,omitempty"`
	// Active alerts by area (state/territory)
	Areas *map[string]int32 `json:"areas,omitempty"`
	// Active alerts by NWS public zone or county code
	Zones *map[string]int32 `json:"zones,omitempty"`
}

// NewAlertsActiveCount200Response instantiates a new AlertsActiveCount200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertsActiveCount200Response() *AlertsActiveCount200Response {
	this := AlertsActiveCount200Response{}
	return &this
}

// NewAlertsActiveCount200ResponseWithDefaults instantiates a new AlertsActiveCount200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertsActiveCount200ResponseWithDefaults() *AlertsActiveCount200Response {
	this := AlertsActiveCount200Response{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *AlertsActiveCount200Response) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsActiveCount200Response) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *AlertsActiveCount200Response) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *AlertsActiveCount200Response) SetTotal(v int32) {
	o.Total = &v
}

// GetLand returns the Land field value if set, zero value otherwise.
func (o *AlertsActiveCount200Response) GetLand() int32 {
	if o == nil || IsNil(o.Land) {
		var ret int32
		return ret
	}
	return *o.Land
}

// GetLandOk returns a tuple with the Land field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsActiveCount200Response) GetLandOk() (*int32, bool) {
	if o == nil || IsNil(o.Land) {
		return nil, false
	}
	return o.Land, true
}

// HasLand returns a boolean if a field has been set.
func (o *AlertsActiveCount200Response) HasLand() bool {
	if o != nil && !IsNil(o.Land) {
		return true
	}

	return false
}

// SetLand gets a reference to the given int32 and assigns it to the Land field.
func (o *AlertsActiveCount200Response) SetLand(v int32) {
	o.Land = &v
}

// GetMarine returns the Marine field value if set, zero value otherwise.
func (o *AlertsActiveCount200Response) GetMarine() int32 {
	if o == nil || IsNil(o.Marine) {
		var ret int32
		return ret
	}
	return *o.Marine
}

// GetMarineOk returns a tuple with the Marine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsActiveCount200Response) GetMarineOk() (*int32, bool) {
	if o == nil || IsNil(o.Marine) {
		return nil, false
	}
	return o.Marine, true
}

// HasMarine returns a boolean if a field has been set.
func (o *AlertsActiveCount200Response) HasMarine() bool {
	if o != nil && !IsNil(o.Marine) {
		return true
	}

	return false
}

// SetMarine gets a reference to the given int32 and assigns it to the Marine field.
func (o *AlertsActiveCount200Response) SetMarine(v int32) {
	o.Marine = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *AlertsActiveCount200Response) GetRegions() map[string]int32 {
	if o == nil || IsNil(o.Regions) {
		var ret map[string]int32
		return ret
	}
	return *o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsActiveCount200Response) GetRegionsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *AlertsActiveCount200Response) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given map[string]int32 and assigns it to the Regions field.
func (o *AlertsActiveCount200Response) SetRegions(v map[string]int32) {
	o.Regions = &v
}

// GetAreas returns the Areas field value if set, zero value otherwise.
func (o *AlertsActiveCount200Response) GetAreas() map[string]int32 {
	if o == nil || IsNil(o.Areas) {
		var ret map[string]int32
		return ret
	}
	return *o.Areas
}

// GetAreasOk returns a tuple with the Areas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsActiveCount200Response) GetAreasOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.Areas) {
		return nil, false
	}
	return o.Areas, true
}

// HasAreas returns a boolean if a field has been set.
func (o *AlertsActiveCount200Response) HasAreas() bool {
	if o != nil && !IsNil(o.Areas) {
		return true
	}

	return false
}

// SetAreas gets a reference to the given map[string]int32 and assigns it to the Areas field.
func (o *AlertsActiveCount200Response) SetAreas(v map[string]int32) {
	o.Areas = &v
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *AlertsActiveCount200Response) GetZones() map[string]int32 {
	if o == nil || IsNil(o.Zones) {
		var ret map[string]int32
		return ret
	}
	return *o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsActiveCount200Response) GetZonesOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.Zones) {
		return nil, false
	}
	return o.Zones, true
}

// HasZones returns a boolean if a field has been set.
func (o *AlertsActiveCount200Response) HasZones() bool {
	if o != nil && !IsNil(o.Zones) {
		return true
	}

	return false
}

// SetZones gets a reference to the given map[string]int32 and assigns it to the Zones field.
func (o *AlertsActiveCount200Response) SetZones(v map[string]int32) {
	o.Zones = &v
}

func (o AlertsActiveCount200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertsActiveCount200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Land) {
		toSerialize["land"] = o.Land
	}
	if !IsNil(o.Marine) {
		toSerialize["marine"] = o.Marine
	}
	if !IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	if !IsNil(o.Areas) {
		toSerialize["areas"] = o.Areas
	}
	if !IsNil(o.Zones) {
		toSerialize["zones"] = o.Zones
	}
	return toSerialize, nil
}

type NullableAlertsActiveCount200Response struct {
	value *AlertsActiveCount200Response
	isSet bool
}

func (v NullableAlertsActiveCount200Response) Get() *AlertsActiveCount200Response {
	return v.value
}

func (v *NullableAlertsActiveCount200Response) Set(val *AlertsActiveCount200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertsActiveCount200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertsActiveCount200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertsActiveCount200Response(val *AlertsActiveCount200Response) *NullableAlertsActiveCount200Response {
	return &NullableAlertsActiveCount200Response{value: val, isSet: true}
}

func (v NullableAlertsActiveCount200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertsActiveCount200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


