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

// checks if the ObservationCollectionGeoJson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObservationCollectionGeoJson{}

// ObservationCollectionGeoJson struct for ObservationCollectionGeoJson
type ObservationCollectionGeoJson struct {
	Context *JsonLdContext `json:"@context,omitempty"`
	Type string `json:"type"`
	Features []ObservationCollectionGeoJsonAllOfFeatures `json:"features"`
}

type _ObservationCollectionGeoJson ObservationCollectionGeoJson

// NewObservationCollectionGeoJson instantiates a new ObservationCollectionGeoJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObservationCollectionGeoJson(type_ string, features []ObservationCollectionGeoJsonAllOfFeatures) *ObservationCollectionGeoJson {
	this := ObservationCollectionGeoJson{}
	this.Type = type_
	this.Features = features
	return &this
}

// NewObservationCollectionGeoJsonWithDefaults instantiates a new ObservationCollectionGeoJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObservationCollectionGeoJsonWithDefaults() *ObservationCollectionGeoJson {
	this := ObservationCollectionGeoJson{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ObservationCollectionGeoJson) GetContext() JsonLdContext {
	if o == nil || IsNil(o.Context) {
		var ret JsonLdContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationCollectionGeoJson) GetContextOk() (*JsonLdContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ObservationCollectionGeoJson) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given JsonLdContext and assigns it to the Context field.
func (o *ObservationCollectionGeoJson) SetContext(v JsonLdContext) {
	o.Context = &v
}

// GetType returns the Type field value
func (o *ObservationCollectionGeoJson) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservationCollectionGeoJson) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ObservationCollectionGeoJson) SetType(v string) {
	o.Type = v
}

// GetFeatures returns the Features field value
func (o *ObservationCollectionGeoJson) GetFeatures() []ObservationCollectionGeoJsonAllOfFeatures {
	if o == nil {
		var ret []ObservationCollectionGeoJsonAllOfFeatures
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *ObservationCollectionGeoJson) GetFeaturesOk() ([]ObservationCollectionGeoJsonAllOfFeatures, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *ObservationCollectionGeoJson) SetFeatures(v []ObservationCollectionGeoJsonAllOfFeatures) {
	o.Features = v
}

func (o ObservationCollectionGeoJson) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObservationCollectionGeoJson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["@context"] = o.Context
	}
	toSerialize["type"] = o.Type
	toSerialize["features"] = o.Features
	return toSerialize, nil
}

func (o *ObservationCollectionGeoJson) UnmarshalJSON(data []byte) (err error) {
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

	varObservationCollectionGeoJson := _ObservationCollectionGeoJson{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varObservationCollectionGeoJson)

	if err != nil {
		return err
	}

	*o = ObservationCollectionGeoJson(varObservationCollectionGeoJson)

	return err
}

type NullableObservationCollectionGeoJson struct {
	value *ObservationCollectionGeoJson
	isSet bool
}

func (v NullableObservationCollectionGeoJson) Get() *ObservationCollectionGeoJson {
	return v.value
}

func (v *NullableObservationCollectionGeoJson) Set(val *ObservationCollectionGeoJson) {
	v.value = val
	v.isSet = true
}

func (v NullableObservationCollectionGeoJson) IsSet() bool {
	return v.isSet
}

func (v *NullableObservationCollectionGeoJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservationCollectionGeoJson(val *ObservationCollectionGeoJson) *NullableObservationCollectionGeoJson {
	return &NullableObservationCollectionGeoJson{value: val, isSet: true}
}

func (v NullableObservationCollectionGeoJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservationCollectionGeoJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


