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
	"bytes"
	"fmt"
)

// checks if the AlertCollectionGeoJson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertCollectionGeoJson{}

// AlertCollectionGeoJson struct for AlertCollectionGeoJson
type AlertCollectionGeoJson struct {
	Context *JsonLdContext `json:"@context,omitempty"`
	Type string `json:"type"`
	Features []AlertCollectionGeoJsonAllOfFeatures `json:"features"`
	// A title describing the alert collection
	Title *string `json:"title,omitempty"`
	// The last time a change occurred to this collection
	Updated *time.Time `json:"updated,omitempty"`
	Pagination *PaginationInfo `json:"pagination,omitempty"`
}

type _AlertCollectionGeoJson AlertCollectionGeoJson

// NewAlertCollectionGeoJson instantiates a new AlertCollectionGeoJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertCollectionGeoJson(type_ string, features []AlertCollectionGeoJsonAllOfFeatures) *AlertCollectionGeoJson {
	this := AlertCollectionGeoJson{}
	this.Type = type_
	this.Features = features
	return &this
}

// NewAlertCollectionGeoJsonWithDefaults instantiates a new AlertCollectionGeoJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertCollectionGeoJsonWithDefaults() *AlertCollectionGeoJson {
	this := AlertCollectionGeoJson{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *AlertCollectionGeoJson) GetContext() JsonLdContext {
	if o == nil || IsNil(o.Context) {
		var ret JsonLdContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertCollectionGeoJson) GetContextOk() (*JsonLdContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *AlertCollectionGeoJson) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given JsonLdContext and assigns it to the Context field.
func (o *AlertCollectionGeoJson) SetContext(v JsonLdContext) {
	o.Context = &v
}

// GetType returns the Type field value
func (o *AlertCollectionGeoJson) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AlertCollectionGeoJson) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AlertCollectionGeoJson) SetType(v string) {
	o.Type = v
}

// GetFeatures returns the Features field value
func (o *AlertCollectionGeoJson) GetFeatures() []AlertCollectionGeoJsonAllOfFeatures {
	if o == nil {
		var ret []AlertCollectionGeoJsonAllOfFeatures
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *AlertCollectionGeoJson) GetFeaturesOk() ([]AlertCollectionGeoJsonAllOfFeatures, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *AlertCollectionGeoJson) SetFeatures(v []AlertCollectionGeoJsonAllOfFeatures) {
	o.Features = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AlertCollectionGeoJson) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertCollectionGeoJson) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AlertCollectionGeoJson) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AlertCollectionGeoJson) SetTitle(v string) {
	o.Title = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *AlertCollectionGeoJson) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertCollectionGeoJson) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *AlertCollectionGeoJson) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *AlertCollectionGeoJson) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *AlertCollectionGeoJson) GetPagination() PaginationInfo {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationInfo
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertCollectionGeoJson) GetPaginationOk() (*PaginationInfo, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *AlertCollectionGeoJson) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationInfo and assigns it to the Pagination field.
func (o *AlertCollectionGeoJson) SetPagination(v PaginationInfo) {
	o.Pagination = &v
}

func (o AlertCollectionGeoJson) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertCollectionGeoJson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["@context"] = o.Context
	}
	toSerialize["type"] = o.Type
	toSerialize["features"] = o.Features
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

func (o *AlertCollectionGeoJson) UnmarshalJSON(data []byte) (err error) {
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

	varAlertCollectionGeoJson := _AlertCollectionGeoJson{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlertCollectionGeoJson)

	if err != nil {
		return err
	}

	*o = AlertCollectionGeoJson(varAlertCollectionGeoJson)

	return err
}

type NullableAlertCollectionGeoJson struct {
	value *AlertCollectionGeoJson
	isSet bool
}

func (v NullableAlertCollectionGeoJson) Get() *AlertCollectionGeoJson {
	return v.value
}

func (v *NullableAlertCollectionGeoJson) Set(val *AlertCollectionGeoJson) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertCollectionGeoJson) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertCollectionGeoJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertCollectionGeoJson(val *AlertCollectionGeoJson) *NullableAlertCollectionGeoJson {
	return &NullableAlertCollectionGeoJson{value: val, isSet: true}
}

func (v NullableAlertCollectionGeoJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertCollectionGeoJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


