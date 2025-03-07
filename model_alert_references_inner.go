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

// checks if the AlertReferencesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertReferencesInner{}

// AlertReferencesInner struct for AlertReferencesInner
type AlertReferencesInner struct {
	// An API link to the prior alert.
	Id *string `json:"@id,omitempty"`
	// The identifier of the alert message.
	Identifier *string `json:"identifier,omitempty"`
	// The sender of the prior alert.
	Sender *string `json:"sender,omitempty"`
	// The time the prior alert was sent.
	Sent *time.Time `json:"sent,omitempty"`
}

// NewAlertReferencesInner instantiates a new AlertReferencesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertReferencesInner() *AlertReferencesInner {
	this := AlertReferencesInner{}
	return &this
}

// NewAlertReferencesInnerWithDefaults instantiates a new AlertReferencesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertReferencesInnerWithDefaults() *AlertReferencesInner {
	this := AlertReferencesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlertReferencesInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertReferencesInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlertReferencesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AlertReferencesInner) SetId(v string) {
	o.Id = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *AlertReferencesInner) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertReferencesInner) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *AlertReferencesInner) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *AlertReferencesInner) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *AlertReferencesInner) GetSender() string {
	if o == nil || IsNil(o.Sender) {
		var ret string
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertReferencesInner) GetSenderOk() (*string, bool) {
	if o == nil || IsNil(o.Sender) {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *AlertReferencesInner) HasSender() bool {
	if o != nil && !IsNil(o.Sender) {
		return true
	}

	return false
}

// SetSender gets a reference to the given string and assigns it to the Sender field.
func (o *AlertReferencesInner) SetSender(v string) {
	o.Sender = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *AlertReferencesInner) GetSent() time.Time {
	if o == nil || IsNil(o.Sent) {
		var ret time.Time
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertReferencesInner) GetSentOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Sent) {
		return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *AlertReferencesInner) HasSent() bool {
	if o != nil && !IsNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given time.Time and assigns it to the Sent field.
func (o *AlertReferencesInner) SetSent(v time.Time) {
	o.Sent = &v
}

func (o AlertReferencesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertReferencesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["@id"] = o.Id
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Sender) {
		toSerialize["sender"] = o.Sender
	}
	if !IsNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	return toSerialize, nil
}

type NullableAlertReferencesInner struct {
	value *AlertReferencesInner
	isSet bool
}

func (v NullableAlertReferencesInner) Get() *AlertReferencesInner {
	return v.value
}

func (v *NullableAlertReferencesInner) Set(val *AlertReferencesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertReferencesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertReferencesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertReferencesInner(val *AlertReferencesInner) *NullableAlertReferencesInner {
	return &NullableAlertReferencesInner{value: val, isSet: true}
}

func (v NullableAlertReferencesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertReferencesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


