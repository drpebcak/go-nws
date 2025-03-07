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

// checks if the ObservationGeoJson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObservationGeoJson{}

// ObservationGeoJson struct for ObservationGeoJson
type ObservationGeoJson struct {
	Context *JsonLdContext `json:"@context,omitempty"`
	Id *string `json:"id,omitempty"`
	Type string `json:"type"`
	Geometry NullableGeoJsonGeometry `json:"geometry"`
	Properties Observation `json:"properties"`
}

type _ObservationGeoJson ObservationGeoJson

// NewObservationGeoJson instantiates a new ObservationGeoJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObservationGeoJson(type_ string, geometry NullableGeoJsonGeometry, properties Observation) *ObservationGeoJson {
	this := ObservationGeoJson{}
	this.Type = type_
	this.Geometry = geometry
	this.Properties = properties
	return &this
}

// NewObservationGeoJsonWithDefaults instantiates a new ObservationGeoJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObservationGeoJsonWithDefaults() *ObservationGeoJson {
	this := ObservationGeoJson{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ObservationGeoJson) GetContext() JsonLdContext {
	if o == nil || IsNil(o.Context) {
		var ret JsonLdContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationGeoJson) GetContextOk() (*JsonLdContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ObservationGeoJson) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given JsonLdContext and assigns it to the Context field.
func (o *ObservationGeoJson) SetContext(v JsonLdContext) {
	o.Context = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ObservationGeoJson) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationGeoJson) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ObservationGeoJson) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ObservationGeoJson) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value
func (o *ObservationGeoJson) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservationGeoJson) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ObservationGeoJson) SetType(v string) {
	o.Type = v
}

// GetGeometry returns the Geometry field value
// If the value is explicit nil, the zero value for GeoJsonGeometry will be returned
func (o *ObservationGeoJson) GetGeometry() GeoJsonGeometry {
	if o == nil || o.Geometry.Get() == nil {
		var ret GeoJsonGeometry
		return ret
	}

	return *o.Geometry.Get()
}

// GetGeometryOk returns a tuple with the Geometry field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObservationGeoJson) GetGeometryOk() (*GeoJsonGeometry, bool) {
	if o == nil {
		return nil, false
	}
	return o.Geometry.Get(), o.Geometry.IsSet()
}

// SetGeometry sets field value
func (o *ObservationGeoJson) SetGeometry(v GeoJsonGeometry) {
	o.Geometry.Set(&v)
}

// GetProperties returns the Properties field value
func (o *ObservationGeoJson) GetProperties() Observation {
	if o == nil {
		var ret Observation
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *ObservationGeoJson) GetPropertiesOk() (*Observation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *ObservationGeoJson) SetProperties(v Observation) {
	o.Properties = v
}

func (o ObservationGeoJson) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObservationGeoJson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["@context"] = o.Context
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type
	toSerialize["geometry"] = o.Geometry.Get()
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

func (o *ObservationGeoJson) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"geometry",
		"properties",
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

	varObservationGeoJson := _ObservationGeoJson{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varObservationGeoJson)

	if err != nil {
		return err
	}

	*o = ObservationGeoJson(varObservationGeoJson)

	return err
}

type NullableObservationGeoJson struct {
	value *ObservationGeoJson
	isSet bool
}

func (v NullableObservationGeoJson) Get() *ObservationGeoJson {
	return v.value
}

func (v *NullableObservationGeoJson) Set(val *ObservationGeoJson) {
	v.value = val
	v.isSet = true
}

func (v NullableObservationGeoJson) IsSet() bool {
	return v.isSet
}

func (v *NullableObservationGeoJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservationGeoJson(val *ObservationGeoJson) *NullableObservationGeoJson {
	return &NullableObservationGeoJson{value: val, isSet: true}
}

func (v NullableObservationGeoJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservationGeoJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


