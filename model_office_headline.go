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

// checks if the OfficeHeadline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfficeHeadline{}

// OfficeHeadline struct for OfficeHeadline
type OfficeHeadline struct {
	Context *JsonLdContext `json:"@context,omitempty"`
	Id *string `json:"@id,omitempty"`
	Id *string `json:"id,omitempty"`
	Office *string `json:"office,omitempty"`
	Important *bool `json:"important,omitempty"`
	IssuanceTime *time.Time `json:"issuanceTime,omitempty"`
	Link *string `json:"link,omitempty"`
	Name *string `json:"name,omitempty"`
	Title *string `json:"title,omitempty"`
	Summary NullableString `json:"summary,omitempty"`
	Content *string `json:"content,omitempty"`
}

// NewOfficeHeadline instantiates a new OfficeHeadline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfficeHeadline() *OfficeHeadline {
	this := OfficeHeadline{}
	return &this
}

// NewOfficeHeadlineWithDefaults instantiates a new OfficeHeadline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfficeHeadlineWithDefaults() *OfficeHeadline {
	this := OfficeHeadline{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *OfficeHeadline) GetContext() JsonLdContext {
	if o == nil || IsNil(o.Context) {
		var ret JsonLdContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeHeadline) GetContextOk() (*JsonLdContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *OfficeHeadline) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given JsonLdContext and assigns it to the Context field.
func (o *OfficeHeadline) SetContext(v JsonLdContext) {
	o.Context = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OfficeHeadline) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeHeadline) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OfficeHeadline) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OfficeHeadline) SetId(v string) {
	o.Id = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OfficeHeadline) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeHeadline) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OfficeHeadline) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OfficeHeadline) SetId(v string) {
	o.Id = &v
}

// GetOffice returns the Office field value if set, zero value otherwise.
func (o *OfficeHeadline) GetOffice() string {
	if o == nil || IsNil(o.Office) {
		var ret string
		return ret
	}
	return *o.Office
}

// GetOfficeOk returns a tuple with the Office field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeHeadline) GetOfficeOk() (*string, bool) {
	if o == nil || IsNil(o.Office) {
		return nil, false
	}
	return o.Office, true
}

// HasOffice returns a boolean if a field has been set.
func (o *OfficeHeadline) HasOffice() bool {
	if o != nil && !IsNil(o.Office) {
		return true
	}

	return false
}

// SetOffice gets a reference to the given string and assigns it to the Office field.
func (o *OfficeHeadline) SetOffice(v string) {
	o.Office = &v
}

// GetImportant returns the Important field value if set, zero value otherwise.
func (o *OfficeHeadline) GetImportant() bool {
	if o == nil || IsNil(o.Important) {
		var ret bool
		return ret
	}
	return *o.Important
}

// GetImportantOk returns a tuple with the Important field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeHeadline) GetImportantOk() (*bool, bool) {
	if o == nil || IsNil(o.Important) {
		return nil, false
	}
	return o.Important, true
}

// HasImportant returns a boolean if a field has been set.
func (o *OfficeHeadline) HasImportant() bool {
	if o != nil && !IsNil(o.Important) {
		return true
	}

	return false
}

// SetImportant gets a reference to the given bool and assigns it to the Important field.
func (o *OfficeHeadline) SetImportant(v bool) {
	o.Important = &v
}

// GetIssuanceTime returns the IssuanceTime field value if set, zero value otherwise.
func (o *OfficeHeadline) GetIssuanceTime() time.Time {
	if o == nil || IsNil(o.IssuanceTime) {
		var ret time.Time
		return ret
	}
	return *o.IssuanceTime
}

// GetIssuanceTimeOk returns a tuple with the IssuanceTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeHeadline) GetIssuanceTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.IssuanceTime) {
		return nil, false
	}
	return o.IssuanceTime, true
}

// HasIssuanceTime returns a boolean if a field has been set.
func (o *OfficeHeadline) HasIssuanceTime() bool {
	if o != nil && !IsNil(o.IssuanceTime) {
		return true
	}

	return false
}

// SetIssuanceTime gets a reference to the given time.Time and assigns it to the IssuanceTime field.
func (o *OfficeHeadline) SetIssuanceTime(v time.Time) {
	o.IssuanceTime = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *OfficeHeadline) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeHeadline) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *OfficeHeadline) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *OfficeHeadline) SetLink(v string) {
	o.Link = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OfficeHeadline) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeHeadline) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OfficeHeadline) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OfficeHeadline) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *OfficeHeadline) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeHeadline) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *OfficeHeadline) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *OfficeHeadline) SetTitle(v string) {
	o.Title = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OfficeHeadline) GetSummary() string {
	if o == nil || IsNil(o.Summary.Get()) {
		var ret string
		return ret
	}
	return *o.Summary.Get()
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OfficeHeadline) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Summary.Get(), o.Summary.IsSet()
}

// HasSummary returns a boolean if a field has been set.
func (o *OfficeHeadline) HasSummary() bool {
	if o != nil && o.Summary.IsSet() {
		return true
	}

	return false
}

// SetSummary gets a reference to the given NullableString and assigns it to the Summary field.
func (o *OfficeHeadline) SetSummary(v string) {
	o.Summary.Set(&v)
}
// SetSummaryNil sets the value for Summary to be an explicit nil
func (o *OfficeHeadline) SetSummaryNil() {
	o.Summary.Set(nil)
}

// UnsetSummary ensures that no value is present for Summary, not even an explicit nil
func (o *OfficeHeadline) UnsetSummary() {
	o.Summary.Unset()
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *OfficeHeadline) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeHeadline) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *OfficeHeadline) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *OfficeHeadline) SetContent(v string) {
	o.Content = &v
}

func (o OfficeHeadline) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfficeHeadline) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["@context"] = o.Context
	}
	if !IsNil(o.Id) {
		toSerialize["@id"] = o.Id
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Office) {
		toSerialize["office"] = o.Office
	}
	if !IsNil(o.Important) {
		toSerialize["important"] = o.Important
	}
	if !IsNil(o.IssuanceTime) {
		toSerialize["issuanceTime"] = o.IssuanceTime
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if o.Summary.IsSet() {
		toSerialize["summary"] = o.Summary.Get()
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

type NullableOfficeHeadline struct {
	value *OfficeHeadline
	isSet bool
}

func (v NullableOfficeHeadline) Get() *OfficeHeadline {
	return v.value
}

func (v *NullableOfficeHeadline) Set(val *OfficeHeadline) {
	v.value = val
	v.isSet = true
}

func (v NullableOfficeHeadline) IsSet() bool {
	return v.isSet
}

func (v *NullableOfficeHeadline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfficeHeadline(val *OfficeHeadline) *NullableOfficeHeadline {
	return &NullableOfficeHeadline{value: val, isSet: true}
}

func (v NullableOfficeHeadline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfficeHeadline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


