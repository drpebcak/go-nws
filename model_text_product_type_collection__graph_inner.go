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

// checks if the TextProductTypeCollectionGraphInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextProductTypeCollectionGraphInner{}

// TextProductTypeCollectionGraphInner struct for TextProductTypeCollectionGraphInner
type TextProductTypeCollectionGraphInner struct {
	ProductCode string `json:"productCode"`
	ProductName string `json:"productName"`
}

type _TextProductTypeCollectionGraphInner TextProductTypeCollectionGraphInner

// NewTextProductTypeCollectionGraphInner instantiates a new TextProductTypeCollectionGraphInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextProductTypeCollectionGraphInner(productCode string, productName string) *TextProductTypeCollectionGraphInner {
	this := TextProductTypeCollectionGraphInner{}
	this.ProductCode = productCode
	this.ProductName = productName
	return &this
}

// NewTextProductTypeCollectionGraphInnerWithDefaults instantiates a new TextProductTypeCollectionGraphInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextProductTypeCollectionGraphInnerWithDefaults() *TextProductTypeCollectionGraphInner {
	this := TextProductTypeCollectionGraphInner{}
	return &this
}

// GetProductCode returns the ProductCode field value
func (o *TextProductTypeCollectionGraphInner) GetProductCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value
// and a boolean to check if the value has been set.
func (o *TextProductTypeCollectionGraphInner) GetProductCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductCode, true
}

// SetProductCode sets field value
func (o *TextProductTypeCollectionGraphInner) SetProductCode(v string) {
	o.ProductCode = v
}

// GetProductName returns the ProductName field value
func (o *TextProductTypeCollectionGraphInner) GetProductName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value
// and a boolean to check if the value has been set.
func (o *TextProductTypeCollectionGraphInner) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductName, true
}

// SetProductName sets field value
func (o *TextProductTypeCollectionGraphInner) SetProductName(v string) {
	o.ProductName = v
}

func (o TextProductTypeCollectionGraphInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextProductTypeCollectionGraphInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productCode"] = o.ProductCode
	toSerialize["productName"] = o.ProductName
	return toSerialize, nil
}

func (o *TextProductTypeCollectionGraphInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"productCode",
		"productName",
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

	varTextProductTypeCollectionGraphInner := _TextProductTypeCollectionGraphInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTextProductTypeCollectionGraphInner)

	if err != nil {
		return err
	}

	*o = TextProductTypeCollectionGraphInner(varTextProductTypeCollectionGraphInner)

	return err
}

type NullableTextProductTypeCollectionGraphInner struct {
	value *TextProductTypeCollectionGraphInner
	isSet bool
}

func (v NullableTextProductTypeCollectionGraphInner) Get() *TextProductTypeCollectionGraphInner {
	return v.value
}

func (v *NullableTextProductTypeCollectionGraphInner) Set(val *TextProductTypeCollectionGraphInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTextProductTypeCollectionGraphInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTextProductTypeCollectionGraphInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextProductTypeCollectionGraphInner(val *TextProductTypeCollectionGraphInner) *NullableTextProductTypeCollectionGraphInner {
	return &NullableTextProductTypeCollectionGraphInner{value: val, isSet: true}
}

func (v NullableTextProductTypeCollectionGraphInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextProductTypeCollectionGraphInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


