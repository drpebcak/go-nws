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

// checks if the ZoneCollectionGeoJson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneCollectionGeoJson{}

// ZoneCollectionGeoJson struct for ZoneCollectionGeoJson
type ZoneCollectionGeoJson struct {
	Context *JsonLdContext `json:"@context,omitempty"`
	Type string `json:"type"`
	Features []ZoneCollectionGeoJsonAllOfFeatures `json:"features"`
}

type _ZoneCollectionGeoJson ZoneCollectionGeoJson

// NewZoneCollectionGeoJson instantiates a new ZoneCollectionGeoJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneCollectionGeoJson(type_ string, features []ZoneCollectionGeoJsonAllOfFeatures) *ZoneCollectionGeoJson {
	this := ZoneCollectionGeoJson{}
	this.Type = type_
	this.Features = features
	return &this
}

// NewZoneCollectionGeoJsonWithDefaults instantiates a new ZoneCollectionGeoJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneCollectionGeoJsonWithDefaults() *ZoneCollectionGeoJson {
	this := ZoneCollectionGeoJson{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ZoneCollectionGeoJson) GetContext() JsonLdContext {
	if o == nil || IsNil(o.Context) {
		var ret JsonLdContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneCollectionGeoJson) GetContextOk() (*JsonLdContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ZoneCollectionGeoJson) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given JsonLdContext and assigns it to the Context field.
func (o *ZoneCollectionGeoJson) SetContext(v JsonLdContext) {
	o.Context = &v
}

// GetType returns the Type field value
func (o *ZoneCollectionGeoJson) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ZoneCollectionGeoJson) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ZoneCollectionGeoJson) SetType(v string) {
	o.Type = v
}

// GetFeatures returns the Features field value
func (o *ZoneCollectionGeoJson) GetFeatures() []ZoneCollectionGeoJsonAllOfFeatures {
	if o == nil {
		var ret []ZoneCollectionGeoJsonAllOfFeatures
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *ZoneCollectionGeoJson) GetFeaturesOk() ([]ZoneCollectionGeoJsonAllOfFeatures, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *ZoneCollectionGeoJson) SetFeatures(v []ZoneCollectionGeoJsonAllOfFeatures) {
	o.Features = v
}

func (o ZoneCollectionGeoJson) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZoneCollectionGeoJson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["@context"] = o.Context
	}
	toSerialize["type"] = o.Type
	toSerialize["features"] = o.Features
	return toSerialize, nil
}

func (o *ZoneCollectionGeoJson) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"features",
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

	varZoneCollectionGeoJson := _ZoneCollectionGeoJson{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varZoneCollectionGeoJson)

	if err != nil {
		return err
	}

	*o = ZoneCollectionGeoJson(varZoneCollectionGeoJson)

	return err
}

type NullableZoneCollectionGeoJson struct {
	value *ZoneCollectionGeoJson
	isSet bool
}

func (v NullableZoneCollectionGeoJson) Get() *ZoneCollectionGeoJson {
	return v.value
}

func (v *NullableZoneCollectionGeoJson) Set(val *ZoneCollectionGeoJson) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneCollectionGeoJson) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneCollectionGeoJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneCollectionGeoJson(val *ZoneCollectionGeoJson) *NullableZoneCollectionGeoJson {
	return &NullableZoneCollectionGeoJson{value: val, isSet: true}
}

func (v NullableZoneCollectionGeoJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneCollectionGeoJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


