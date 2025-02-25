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

// checks if the GeoJSONPolygon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeoJSONPolygon{}

// GeoJSONPolygon struct for GeoJSONPolygon
type GeoJSONPolygon struct {
	Type string `json:"type"`
	// A GeoJSON polygon. Please refer to IETF RFC 7946 for information on the GeoJSON format.
	Coordinates [][][]float32 `json:"coordinates"`
	// A GeoJSON bounding box. Please refer to IETF RFC 7946 for information on the GeoJSON format.
	Bbox []float32 `json:"bbox,omitempty"`
}

type _GeoJSONPolygon GeoJSONPolygon

// NewGeoJSONPolygon instantiates a new GeoJSONPolygon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoJSONPolygon(type_ string, coordinates [][][]float32) *GeoJSONPolygon {
	this := GeoJSONPolygon{}
	this.Type = type_
	this.Coordinates = coordinates
	return &this
}

// NewGeoJSONPolygonWithDefaults instantiates a new GeoJSONPolygon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoJSONPolygonWithDefaults() *GeoJSONPolygon {
	this := GeoJSONPolygon{}
	return &this
}

// GetType returns the Type field value
func (o *GeoJSONPolygon) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GeoJSONPolygon) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GeoJSONPolygon) SetType(v string) {
	o.Type = v
}

// GetCoordinates returns the Coordinates field value
func (o *GeoJSONPolygon) GetCoordinates() [][][]float32 {
	if o == nil {
		var ret [][][]float32
		return ret
	}

	return o.Coordinates
}

// GetCoordinatesOk returns a tuple with the Coordinates field value
// and a boolean to check if the value has been set.
func (o *GeoJSONPolygon) GetCoordinatesOk() ([][][]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Coordinates, true
}

// SetCoordinates sets field value
func (o *GeoJSONPolygon) SetCoordinates(v [][][]float32) {
	o.Coordinates = v
}

// GetBbox returns the Bbox field value if set, zero value otherwise.
func (o *GeoJSONPolygon) GetBbox() []float32 {
	if o == nil || IsNil(o.Bbox) {
		var ret []float32
		return ret
	}
	return o.Bbox
}

// GetBboxOk returns a tuple with the Bbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoJSONPolygon) GetBboxOk() ([]float32, bool) {
	if o == nil || IsNil(o.Bbox) {
		return nil, false
	}
	return o.Bbox, true
}

// HasBbox returns a boolean if a field has been set.
func (o *GeoJSONPolygon) HasBbox() bool {
	if o != nil && !IsNil(o.Bbox) {
		return true
	}

	return false
}

// SetBbox gets a reference to the given []float32 and assigns it to the Bbox field.
func (o *GeoJSONPolygon) SetBbox(v []float32) {
	o.Bbox = v
}

func (o GeoJSONPolygon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeoJSONPolygon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["coordinates"] = o.Coordinates
	if !IsNil(o.Bbox) {
		toSerialize["bbox"] = o.Bbox
	}
	return toSerialize, nil
}

func (o *GeoJSONPolygon) UnmarshalJSON(data []byte) (err error) {
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

	varGeoJSONPolygon := _GeoJSONPolygon{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGeoJSONPolygon)

	if err != nil {
		return err
	}

	*o = GeoJSONPolygon(varGeoJSONPolygon)

	return err
}

type NullableGeoJSONPolygon struct {
	value *GeoJSONPolygon
	isSet bool
}

func (v NullableGeoJSONPolygon) Get() *GeoJSONPolygon {
	return v.value
}

func (v *NullableGeoJSONPolygon) Set(val *GeoJSONPolygon) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoJSONPolygon) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoJSONPolygon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoJSONPolygon(val *GeoJSONPolygon) *NullableGeoJSONPolygon {
	return &NullableGeoJSONPolygon{value: val, isSet: true}
}

func (v NullableGeoJSONPolygon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoJSONPolygon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


