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

// NsiIdInfo Represents the S-NSSAI and the optionally associated Network Slice Instance(s).
type NsiIdInfo struct {
	Snssai map[string]interface{} `json:"snssai"`
	NsiIds *[]map[string]interface{} `json:"nsiIds,omitempty"`
}

// NewNsiIdInfo instantiates a new NsiIdInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsiIdInfo(snssai map[string]interface{}, ) *NsiIdInfo {
	this := NsiIdInfo{}
	this.Snssai = snssai
	return &this
}

// NewNsiIdInfoWithDefaults instantiates a new NsiIdInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsiIdInfoWithDefaults() *NsiIdInfo {
	this := NsiIdInfo{}
	return &this
}

// GetSnssai returns the Snssai field value
func (o *NsiIdInfo) GetSnssai() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}

	return o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value
// and a boolean to check if the value has been set.
func (o *NsiIdInfo) GetSnssaiOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Snssai, true
}

// SetSnssai sets field value
func (o *NsiIdInfo) SetSnssai(v map[string]interface{}) {
	o.Snssai = v
}

// GetNsiIds returns the NsiIds field value if set, zero value otherwise.
func (o *NsiIdInfo) GetNsiIds() []map[string]interface{} {
	if o == nil || o.NsiIds == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.NsiIds
}

// GetNsiIdsOk returns a tuple with the NsiIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsiIdInfo) GetNsiIdsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.NsiIds == nil {
		return nil, false
	}
	return o.NsiIds, true
}

// HasNsiIds returns a boolean if a field has been set.
func (o *NsiIdInfo) HasNsiIds() bool {
	if o != nil && o.NsiIds != nil {
		return true
	}

	return false
}

// SetNsiIds gets a reference to the given []map[string]interface{} and assigns it to the NsiIds field.
func (o *NsiIdInfo) SetNsiIds(v []map[string]interface{}) {
	o.NsiIds = &v
}

func (o NsiIdInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["snssai"] = o.Snssai
	}
	if o.NsiIds != nil {
		toSerialize["nsiIds"] = o.NsiIds
	}
	return json.Marshal(toSerialize)
}

type NullableNsiIdInfo struct {
	value *NsiIdInfo
	isSet bool
}

func (v NullableNsiIdInfo) Get() *NsiIdInfo {
	return v.value
}

func (v *NullableNsiIdInfo) Set(val *NsiIdInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNsiIdInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNsiIdInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsiIdInfo(val *NsiIdInfo) *NullableNsiIdInfo {
	return &NullableNsiIdInfo{value: val, isSet: true}
}

func (v NullableNsiIdInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsiIdInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


