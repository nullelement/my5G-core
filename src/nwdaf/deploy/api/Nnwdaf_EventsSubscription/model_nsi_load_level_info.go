/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.2.0-alpha.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// NsiLoadLevelInfo Represents the slice instance and the load level information.
type NsiLoadLevelInfo struct {
	// Load level information of the network slice instance.
	LoadLevelInformation int32 `json:"loadLevelInformation"`
	Snssai map[string]interface{} `json:"snssai"`
	NsiId *map[string]interface{} `json:"nsiId,omitempty"`
}

// NewNsiLoadLevelInfo instantiates a new NsiLoadLevelInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsiLoadLevelInfo(loadLevelInformation int32, snssai map[string]interface{}, ) *NsiLoadLevelInfo {
	this := NsiLoadLevelInfo{}
	this.LoadLevelInformation = loadLevelInformation
	this.Snssai = snssai
	return &this
}

// NewNsiLoadLevelInfoWithDefaults instantiates a new NsiLoadLevelInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsiLoadLevelInfoWithDefaults() *NsiLoadLevelInfo {
	this := NsiLoadLevelInfo{}
	return &this
}

// GetLoadLevelInformation returns the LoadLevelInformation field value
func (o *NsiLoadLevelInfo) GetLoadLevelInformation() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.LoadLevelInformation
}

// GetLoadLevelInformationOk returns a tuple with the LoadLevelInformation field value
// and a boolean to check if the value has been set.
func (o *NsiLoadLevelInfo) GetLoadLevelInformationOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LoadLevelInformation, true
}

// SetLoadLevelInformation sets field value
func (o *NsiLoadLevelInfo) SetLoadLevelInformation(v int32) {
	o.LoadLevelInformation = v
}

// GetSnssai returns the Snssai field value
func (o *NsiLoadLevelInfo) GetSnssai() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}

	return o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value
// and a boolean to check if the value has been set.
func (o *NsiLoadLevelInfo) GetSnssaiOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Snssai, true
}

// SetSnssai sets field value
func (o *NsiLoadLevelInfo) SetSnssai(v map[string]interface{}) {
	o.Snssai = v
}

// GetNsiId returns the NsiId field value if set, zero value otherwise.
func (o *NsiLoadLevelInfo) GetNsiId() map[string]interface{} {
	if o == nil || o.NsiId == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.NsiId
}

// GetNsiIdOk returns a tuple with the NsiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsiLoadLevelInfo) GetNsiIdOk() (*map[string]interface{}, bool) {
	if o == nil || o.NsiId == nil {
		return nil, false
	}
	return o.NsiId, true
}

// HasNsiId returns a boolean if a field has been set.
func (o *NsiLoadLevelInfo) HasNsiId() bool {
	if o != nil && o.NsiId != nil {
		return true
	}

	return false
}

// SetNsiId gets a reference to the given map[string]interface{} and assigns it to the NsiId field.
func (o *NsiLoadLevelInfo) SetNsiId(v map[string]interface{}) {
	o.NsiId = &v
}

func (o NsiLoadLevelInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["loadLevelInformation"] = o.LoadLevelInformation
	}
	if true {
		toSerialize["snssai"] = o.Snssai
	}
	if o.NsiId != nil {
		toSerialize["nsiId"] = o.NsiId
	}
	return json.Marshal(toSerialize)
}

type NullableNsiLoadLevelInfo struct {
	value *NsiLoadLevelInfo
	isSet bool
}

func (v NullableNsiLoadLevelInfo) Get() *NsiLoadLevelInfo {
	return v.value
}

func (v *NullableNsiLoadLevelInfo) Set(val *NsiLoadLevelInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNsiLoadLevelInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNsiLoadLevelInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsiLoadLevelInfo(val *NsiLoadLevelInfo) *NullableNsiLoadLevelInfo {
	return &NullableNsiLoadLevelInfo{value: val, isSet: true}
}

func (v NullableNsiLoadLevelInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsiLoadLevelInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

