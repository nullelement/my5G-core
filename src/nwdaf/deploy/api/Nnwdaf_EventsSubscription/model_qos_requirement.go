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

// QosRequirement struct for QosRequirement
type QosRequirement struct {
	Var5qi *map[string]interface{} `json:"5qi,omitempty"`
	GfbrUl *map[string]interface{} `json:"gfbrUl,omitempty"`
	GfbrDl *map[string]interface{} `json:"gfbrDl,omitempty"`
	ResType *map[string]interface{} `json:"resType,omitempty"`
	Pdb *map[string]interface{} `json:"pdb,omitempty"`
	Per *map[string]interface{} `json:"per,omitempty"`
}

// NewQosRequirement instantiates a new QosRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosRequirement() *QosRequirement {
	this := QosRequirement{}
	return &this
}

// NewQosRequirementWithDefaults instantiates a new QosRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosRequirementWithDefaults() *QosRequirement {
	this := QosRequirement{}
	return &this
}

// GetVar5qi returns the Var5qi field value if set, zero value otherwise.
func (o *QosRequirement) GetVar5qi() map[string]interface{} {
	if o == nil || o.Var5qi == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Var5qi
}

// GetVar5qiOk returns a tuple with the Var5qi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosRequirement) GetVar5qiOk() (*map[string]interface{}, bool) {
	if o == nil || o.Var5qi == nil {
		return nil, false
	}
	return o.Var5qi, true
}

// HasVar5qi returns a boolean if a field has been set.
func (o *QosRequirement) HasVar5qi() bool {
	if o != nil && o.Var5qi != nil {
		return true
	}

	return false
}

// SetVar5qi gets a reference to the given map[string]interface{} and assigns it to the Var5qi field.
func (o *QosRequirement) SetVar5qi(v map[string]interface{}) {
	o.Var5qi = &v
}

// GetGfbrUl returns the GfbrUl field value if set, zero value otherwise.
func (o *QosRequirement) GetGfbrUl() map[string]interface{} {
	if o == nil || o.GfbrUl == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.GfbrUl
}

// GetGfbrUlOk returns a tuple with the GfbrUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosRequirement) GetGfbrUlOk() (*map[string]interface{}, bool) {
	if o == nil || o.GfbrUl == nil {
		return nil, false
	}
	return o.GfbrUl, true
}

// HasGfbrUl returns a boolean if a field has been set.
func (o *QosRequirement) HasGfbrUl() bool {
	if o != nil && o.GfbrUl != nil {
		return true
	}

	return false
}

// SetGfbrUl gets a reference to the given map[string]interface{} and assigns it to the GfbrUl field.
func (o *QosRequirement) SetGfbrUl(v map[string]interface{}) {
	o.GfbrUl = &v
}

// GetGfbrDl returns the GfbrDl field value if set, zero value otherwise.
func (o *QosRequirement) GetGfbrDl() map[string]interface{} {
	if o == nil || o.GfbrDl == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.GfbrDl
}

// GetGfbrDlOk returns a tuple with the GfbrDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosRequirement) GetGfbrDlOk() (*map[string]interface{}, bool) {
	if o == nil || o.GfbrDl == nil {
		return nil, false
	}
	return o.GfbrDl, true
}

// HasGfbrDl returns a boolean if a field has been set.
func (o *QosRequirement) HasGfbrDl() bool {
	if o != nil && o.GfbrDl != nil {
		return true
	}

	return false
}

// SetGfbrDl gets a reference to the given map[string]interface{} and assigns it to the GfbrDl field.
func (o *QosRequirement) SetGfbrDl(v map[string]interface{}) {
	o.GfbrDl = &v
}

// GetResType returns the ResType field value if set, zero value otherwise.
func (o *QosRequirement) GetResType() map[string]interface{} {
	if o == nil || o.ResType == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ResType
}

// GetResTypeOk returns a tuple with the ResType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosRequirement) GetResTypeOk() (*map[string]interface{}, bool) {
	if o == nil || o.ResType == nil {
		return nil, false
	}
	return o.ResType, true
}

// HasResType returns a boolean if a field has been set.
func (o *QosRequirement) HasResType() bool {
	if o != nil && o.ResType != nil {
		return true
	}

	return false
}

// SetResType gets a reference to the given map[string]interface{} and assigns it to the ResType field.
func (o *QosRequirement) SetResType(v map[string]interface{}) {
	o.ResType = &v
}

// GetPdb returns the Pdb field value if set, zero value otherwise.
func (o *QosRequirement) GetPdb() map[string]interface{} {
	if o == nil || o.Pdb == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Pdb
}

// GetPdbOk returns a tuple with the Pdb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosRequirement) GetPdbOk() (*map[string]interface{}, bool) {
	if o == nil || o.Pdb == nil {
		return nil, false
	}
	return o.Pdb, true
}

// HasPdb returns a boolean if a field has been set.
func (o *QosRequirement) HasPdb() bool {
	if o != nil && o.Pdb != nil {
		return true
	}

	return false
}

// SetPdb gets a reference to the given map[string]interface{} and assigns it to the Pdb field.
func (o *QosRequirement) SetPdb(v map[string]interface{}) {
	o.Pdb = &v
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *QosRequirement) GetPer() map[string]interface{} {
	if o == nil || o.Per == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosRequirement) GetPerOk() (*map[string]interface{}, bool) {
	if o == nil || o.Per == nil {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *QosRequirement) HasPer() bool {
	if o != nil && o.Per != nil {
		return true
	}

	return false
}

// SetPer gets a reference to the given map[string]interface{} and assigns it to the Per field.
func (o *QosRequirement) SetPer(v map[string]interface{}) {
	o.Per = &v
}

func (o QosRequirement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Var5qi != nil {
		toSerialize["5qi"] = o.Var5qi
	}
	if o.GfbrUl != nil {
		toSerialize["gfbrUl"] = o.GfbrUl
	}
	if o.GfbrDl != nil {
		toSerialize["gfbrDl"] = o.GfbrDl
	}
	if o.ResType != nil {
		toSerialize["resType"] = o.ResType
	}
	if o.Pdb != nil {
		toSerialize["pdb"] = o.Pdb
	}
	if o.Per != nil {
		toSerialize["per"] = o.Per
	}
	return json.Marshal(toSerialize)
}

type NullableQosRequirement struct {
	value *QosRequirement
	isSet bool
}

func (v NullableQosRequirement) Get() *QosRequirement {
	return v.value
}

func (v *NullableQosRequirement) Set(val *QosRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableQosRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableQosRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosRequirement(val *QosRequirement) *NullableQosRequirement {
	return &NullableQosRequirement{value: val, isSet: true}
}

func (v NullableQosRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

