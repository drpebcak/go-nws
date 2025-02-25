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

// checks if the AlertAtomFeed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertAtomFeed{}

// AlertAtomFeed An alert feed in Atom format
type AlertAtomFeed struct {
	Id *string `json:"id,omitempty"`
	Generator *string `json:"generator,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Author *AlertAtomFeedAuthor `json:"author,omitempty"`
	Title *string `json:"title,omitempty"`
	Entry []AlertAtomEntry `json:"entry,omitempty"`
}

// NewAlertAtomFeed instantiates a new AlertAtomFeed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertAtomFeed() *AlertAtomFeed {
	this := AlertAtomFeed{}
	return &this
}

// NewAlertAtomFeedWithDefaults instantiates a new AlertAtomFeed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertAtomFeedWithDefaults() *AlertAtomFeed {
	this := AlertAtomFeed{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlertAtomFeed) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomFeed) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlertAtomFeed) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AlertAtomFeed) SetId(v string) {
	o.Id = &v
}

// GetGenerator returns the Generator field value if set, zero value otherwise.
func (o *AlertAtomFeed) GetGenerator() string {
	if o == nil || IsNil(o.Generator) {
		var ret string
		return ret
	}
	return *o.Generator
}

// GetGeneratorOk returns a tuple with the Generator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomFeed) GetGeneratorOk() (*string, bool) {
	if o == nil || IsNil(o.Generator) {
		return nil, false
	}
	return o.Generator, true
}

// HasGenerator returns a boolean if a field has been set.
func (o *AlertAtomFeed) HasGenerator() bool {
	if o != nil && !IsNil(o.Generator) {
		return true
	}

	return false
}

// SetGenerator gets a reference to the given string and assigns it to the Generator field.
func (o *AlertAtomFeed) SetGenerator(v string) {
	o.Generator = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *AlertAtomFeed) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomFeed) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *AlertAtomFeed) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *AlertAtomFeed) SetUpdated(v string) {
	o.Updated = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *AlertAtomFeed) GetAuthor() AlertAtomFeedAuthor {
	if o == nil || IsNil(o.Author) {
		var ret AlertAtomFeedAuthor
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomFeed) GetAuthorOk() (*AlertAtomFeedAuthor, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *AlertAtomFeed) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given AlertAtomFeedAuthor and assigns it to the Author field.
func (o *AlertAtomFeed) SetAuthor(v AlertAtomFeedAuthor) {
	o.Author = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AlertAtomFeed) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomFeed) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AlertAtomFeed) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AlertAtomFeed) SetTitle(v string) {
	o.Title = &v
}

// GetEntry returns the Entry field value if set, zero value otherwise.
func (o *AlertAtomFeed) GetEntry() []AlertAtomEntry {
	if o == nil || IsNil(o.Entry) {
		var ret []AlertAtomEntry
		return ret
	}
	return o.Entry
}

// GetEntryOk returns a tuple with the Entry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomFeed) GetEntryOk() ([]AlertAtomEntry, bool) {
	if o == nil || IsNil(o.Entry) {
		return nil, false
	}
	return o.Entry, true
}

// HasEntry returns a boolean if a field has been set.
func (o *AlertAtomFeed) HasEntry() bool {
	if o != nil && !IsNil(o.Entry) {
		return true
	}

	return false
}

// SetEntry gets a reference to the given []AlertAtomEntry and assigns it to the Entry field.
func (o *AlertAtomFeed) SetEntry(v []AlertAtomEntry) {
	o.Entry = v
}

func (o AlertAtomFeed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertAtomFeed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Generator) {
		toSerialize["generator"] = o.Generator
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Entry) {
		toSerialize["entry"] = o.Entry
	}
	return toSerialize, nil
}

type NullableAlertAtomFeed struct {
	value *AlertAtomFeed
	isSet bool
}

func (v NullableAlertAtomFeed) Get() *AlertAtomFeed {
	return v.value
}

func (v *NullableAlertAtomFeed) Set(val *AlertAtomFeed) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertAtomFeed) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertAtomFeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertAtomFeed(val *AlertAtomFeed) *NullableAlertAtomFeed {
	return &NullableAlertAtomFeed{value: val, isSet: true}
}

func (v NullableAlertAtomFeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertAtomFeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


