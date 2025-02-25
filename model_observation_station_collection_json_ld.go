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

// checks if the ObservationStationCollectionJsonLd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObservationStationCollectionJsonLd{}

// ObservationStationCollectionJsonLd struct for ObservationStationCollectionJsonLd
type ObservationStationCollectionJsonLd struct {
	Context *JsonLdContext `json:"@context,omitempty"`
	Graph []ObservationStation `json:"@graph,omitempty"`
	ObservationStations []string `json:"observationStations,omitempty"`
	Pagination *PaginationInfo `json:"pagination,omitempty"`
}

// NewObservationStationCollectionJsonLd instantiates a new ObservationStationCollectionJsonLd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObservationStationCollectionJsonLd() *ObservationStationCollectionJsonLd {
	this := ObservationStationCollectionJsonLd{}
	return &this
}

// NewObservationStationCollectionJsonLdWithDefaults instantiates a new ObservationStationCollectionJsonLd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObservationStationCollectionJsonLdWithDefaults() *ObservationStationCollectionJsonLd {
	this := ObservationStationCollectionJsonLd{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ObservationStationCollectionJsonLd) GetContext() JsonLdContext {
	if o == nil || IsNil(o.Context) {
		var ret JsonLdContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationStationCollectionJsonLd) GetContextOk() (*JsonLdContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ObservationStationCollectionJsonLd) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given JsonLdContext and assigns it to the Context field.
func (o *ObservationStationCollectionJsonLd) SetContext(v JsonLdContext) {
	o.Context = &v
}

// GetGraph returns the Graph field value if set, zero value otherwise.
func (o *ObservationStationCollectionJsonLd) GetGraph() []ObservationStation {
	if o == nil || IsNil(o.Graph) {
		var ret []ObservationStation
		return ret
	}
	return o.Graph
}

// GetGraphOk returns a tuple with the Graph field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationStationCollectionJsonLd) GetGraphOk() ([]ObservationStation, bool) {
	if o == nil || IsNil(o.Graph) {
		return nil, false
	}
	return o.Graph, true
}

// HasGraph returns a boolean if a field has been set.
func (o *ObservationStationCollectionJsonLd) HasGraph() bool {
	if o != nil && !IsNil(o.Graph) {
		return true
	}

	return false
}

// SetGraph gets a reference to the given []ObservationStation and assigns it to the Graph field.
func (o *ObservationStationCollectionJsonLd) SetGraph(v []ObservationStation) {
	o.Graph = v
}

// GetObservationStations returns the ObservationStations field value if set, zero value otherwise.
func (o *ObservationStationCollectionJsonLd) GetObservationStations() []string {
	if o == nil || IsNil(o.ObservationStations) {
		var ret []string
		return ret
	}
	return o.ObservationStations
}

// GetObservationStationsOk returns a tuple with the ObservationStations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationStationCollectionJsonLd) GetObservationStationsOk() ([]string, bool) {
	if o == nil || IsNil(o.ObservationStations) {
		return nil, false
	}
	return o.ObservationStations, true
}

// HasObservationStations returns a boolean if a field has been set.
func (o *ObservationStationCollectionJsonLd) HasObservationStations() bool {
	if o != nil && !IsNil(o.ObservationStations) {
		return true
	}

	return false
}

// SetObservationStations gets a reference to the given []string and assigns it to the ObservationStations field.
func (o *ObservationStationCollectionJsonLd) SetObservationStations(v []string) {
	o.ObservationStations = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ObservationStationCollectionJsonLd) GetPagination() PaginationInfo {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationInfo
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationStationCollectionJsonLd) GetPaginationOk() (*PaginationInfo, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ObservationStationCollectionJsonLd) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationInfo and assigns it to the Pagination field.
func (o *ObservationStationCollectionJsonLd) SetPagination(v PaginationInfo) {
	o.Pagination = &v
}

func (o ObservationStationCollectionJsonLd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObservationStationCollectionJsonLd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["@context"] = o.Context
	}
	if !IsNil(o.Graph) {
		toSerialize["@graph"] = o.Graph
	}
	if !IsNil(o.ObservationStations) {
		toSerialize["observationStations"] = o.ObservationStations
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableObservationStationCollectionJsonLd struct {
	value *ObservationStationCollectionJsonLd
	isSet bool
}

func (v NullableObservationStationCollectionJsonLd) Get() *ObservationStationCollectionJsonLd {
	return v.value
}

func (v *NullableObservationStationCollectionJsonLd) Set(val *ObservationStationCollectionJsonLd) {
	v.value = val
	v.isSet = true
}

func (v NullableObservationStationCollectionJsonLd) IsSet() bool {
	return v.isSet
}

func (v *NullableObservationStationCollectionJsonLd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservationStationCollectionJsonLd(val *ObservationStationCollectionJsonLd) *NullableObservationStationCollectionJsonLd {
	return &NullableObservationStationCollectionJsonLd{value: val, isSet: true}
}

func (v NullableObservationStationCollectionJsonLd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservationStationCollectionJsonLd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


