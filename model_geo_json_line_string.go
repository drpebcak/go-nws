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

// checks if the GeoJSONLineString type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeoJSONLineString{}

// GeoJSONLineString struct for GeoJSONLineString
type GeoJSONLineString struct {
	Type string `json:"type"`
	// A GeoJSON line string. Please refer to IETF RFC 7946 for information on the GeoJSON format.
	Coordinates [][]float32 `json:"coordinates"`
	// A GeoJSON bounding box. Please refer to IETF RFC 7946 for information on the GeoJSON format.
	Bbox []float32 `json:"bbox,omitempty"`
}

type _GeoJSONLineString GeoJSONLineString

// NewGeoJSONLineString instantiates a new GeoJSONLineString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoJSONLineString(type_ string, coordinates [][]float32) *GeoJSONLineString {
	this := GeoJSONLineString{}
	this.Type = type_
	this.Coordinates = coordinates
	return &this
}

// NewGeoJSONLineStringWithDefaults instantiates a new GeoJSONLineString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoJSONLineStringWithDefaults() *GeoJSONLineString {
	this := GeoJSONLineString{}
	return &this
}

// GetType returns the Type field value
func (o *GeoJSONLineString) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GeoJSONLineString) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GeoJSONLineString) SetType(v string) {
	o.Type = v
}

// GetCoordinates returns the Coordinates field value
func (o *GeoJSONLineString) GetCoordinates() [][]float32 {
	if o == nil {
		var ret [][]float32
		return ret
	}

	return o.Coordinates
}

// GetCoordinatesOk returns a tuple with the Coordinates field value
// and a boolean to check if the value has been set.
func (o *GeoJSONLineString) GetCoordinatesOk() ([][]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Coordinates, true
}

// SetCoordinates sets field value
func (o *GeoJSONLineString) SetCoordinates(v [][]float32) {
	o.Coordinates = v
}

// GetBbox returns the Bbox field value if set, zero value otherwise.
func (o *GeoJSONLineString) GetBbox() []float32 {
	if o == nil || IsNil(o.Bbox) {
		var ret []float32
		return ret
	}
	return o.Bbox
}

// GetBboxOk returns a tuple with the Bbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoJSONLineString) GetBboxOk() ([]float32, bool) {
	if o == nil || IsNil(o.Bbox) {
		return nil, false
	}
	return o.Bbox, true
}

// HasBbox returns a boolean if a field has been set.
func (o *GeoJSONLineString) HasBbox() bool {
	if o != nil && !IsNil(o.Bbox) {
		return true
	}

	return false
}

// SetBbox gets a reference to the given []float32 and assigns it to the Bbox field.
func (o *GeoJSONLineString) SetBbox(v []float32) {
	o.Bbox = v
}

func (o GeoJSONLineString) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeoJSONLineString) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["coordinates"] = o.Coordinates
	if !IsNil(o.Bbox) {
		toSerialize["bbox"] = o.Bbox
	}
	return toSerialize, nil
}

func (o *GeoJSONLineString) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"coordinates",
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

	varGeoJSONLineString := _GeoJSONLineString{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGeoJSONLineString)

	if err != nil {
		return err
	}

	*o = GeoJSONLineString(varGeoJSONLineString)

	return err
}

type NullableGeoJSONLineString struct {
	value *GeoJSONLineString
	isSet bool
}

func (v NullableGeoJSONLineString) Get() *GeoJSONLineString {
	return v.value
}

func (v *NullableGeoJSONLineString) Set(val *GeoJSONLineString) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoJSONLineString) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoJSONLineString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoJSONLineString(val *GeoJSONLineString) *NullableGeoJSONLineString {
	return &NullableGeoJSONLineString{value: val, isSet: true}
}

func (v NullableGeoJSONLineString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoJSONLineString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


