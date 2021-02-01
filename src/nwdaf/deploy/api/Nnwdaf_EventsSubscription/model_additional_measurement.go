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

// AdditionalMeasurement struct for AdditionalMeasurement
type AdditionalMeasurement struct {
	UnexpLoc *map[string]interface{} `json:"unexpLoc,omitempty"`
	UnexpFlowTeps *[]IpEthFlowDescription `json:"unexpFlowTeps,omitempty"`
	UnexpWakes *[]map[string]interface{} `json:"unexpWakes,omitempty"`
	DdosAttack *AddressList `json:"ddosAttack,omitempty"`
	WrgDest *AddressList `json:"wrgDest,omitempty"`
	Circums *[]CircumstanceDescription `json:"circums,omitempty"`
}

// NewAdditionalMeasurement instantiates a new AdditionalMeasurement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalMeasurement() *AdditionalMeasurement {
	this := AdditionalMeasurement{}
	return &this
}

// NewAdditionalMeasurementWithDefaults instantiates a new AdditionalMeasurement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalMeasurementWithDefaults() *AdditionalMeasurement {
	this := AdditionalMeasurement{}
	return &this
}

// GetUnexpLoc returns the UnexpLoc field value if set, zero value otherwise.
func (o *AdditionalMeasurement) GetUnexpLoc() map[string]interface{} {
	if o == nil || o.UnexpLoc == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.UnexpLoc
}

// GetUnexpLocOk returns a tuple with the UnexpLoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalMeasurement) GetUnexpLocOk() (*map[string]interface{}, bool) {
	if o == nil || o.UnexpLoc == nil {
		return nil, false
	}
	return o.UnexpLoc, true
}

// HasUnexpLoc returns a boolean if a field has been set.
func (o *AdditionalMeasurement) HasUnexpLoc() bool {
	if o != nil && o.UnexpLoc != nil {
		return true
	}

	return false
}

// SetUnexpLoc gets a reference to the given map[string]interface{} and assigns it to the UnexpLoc field.
func (o *AdditionalMeasurement) SetUnexpLoc(v map[string]interface{}) {
	o.UnexpLoc = &v
}

// GetUnexpFlowTeps returns the UnexpFlowTeps field value if set, zero value otherwise.
func (o *AdditionalMeasurement) GetUnexpFlowTeps() []IpEthFlowDescription {
	if o == nil || o.UnexpFlowTeps == nil {
		var ret []IpEthFlowDescription
		return ret
	}
	return *o.UnexpFlowTeps
}

// GetUnexpFlowTepsOk returns a tuple with the UnexpFlowTeps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalMeasurement) GetUnexpFlowTepsOk() (*[]IpEthFlowDescription, bool) {
	if o == nil || o.UnexpFlowTeps == nil {
		return nil, false
	}
	return o.UnexpFlowTeps, true
}

// HasUnexpFlowTeps returns a boolean if a field has been set.
func (o *AdditionalMeasurement) HasUnexpFlowTeps() bool {
	if o != nil && o.UnexpFlowTeps != nil {
		return true
	}

	return false
}

// SetUnexpFlowTeps gets a reference to the given []IpEthFlowDescription and assigns it to the UnexpFlowTeps field.
func (o *AdditionalMeasurement) SetUnexpFlowTeps(v []IpEthFlowDescription) {
	o.UnexpFlowTeps = &v
}

// GetUnexpWakes returns the UnexpWakes field value if set, zero value otherwise.
func (o *AdditionalMeasurement) GetUnexpWakes() []map[string]interface{} {
	if o == nil || o.UnexpWakes == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.UnexpWakes
}

// GetUnexpWakesOk returns a tuple with the UnexpWakes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalMeasurement) GetUnexpWakesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.UnexpWakes == nil {
		return nil, false
	}
	return o.UnexpWakes, true
}

// HasUnexpWakes returns a boolean if a field has been set.
func (o *AdditionalMeasurement) HasUnexpWakes() bool {
	if o != nil && o.UnexpWakes != nil {
		return true
	}

	return false
}

// SetUnexpWakes gets a reference to the given []map[string]interface{} and assigns it to the UnexpWakes field.
func (o *AdditionalMeasurement) SetUnexpWakes(v []map[string]interface{}) {
	o.UnexpWakes = &v
}

// GetDdosAttack returns the DdosAttack field value if set, zero value otherwise.
func (o *AdditionalMeasurement) GetDdosAttack() AddressList {
	if o == nil || o.DdosAttack == nil {
		var ret AddressList
		return ret
	}
	return *o.DdosAttack
}

// GetDdosAttackOk returns a tuple with the DdosAttack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalMeasurement) GetDdosAttackOk() (*AddressList, bool) {
	if o == nil || o.DdosAttack == nil {
		return nil, false
	}
	return o.DdosAttack, true
}

// HasDdosAttack returns a boolean if a field has been set.
func (o *AdditionalMeasurement) HasDdosAttack() bool {
	if o != nil && o.DdosAttack != nil {
		return true
	}

	return false
}

// SetDdosAttack gets a reference to the given AddressList and assigns it to the DdosAttack field.
func (o *AdditionalMeasurement) SetDdosAttack(v AddressList) {
	o.DdosAttack = &v
}

// GetWrgDest returns the WrgDest field value if set, zero value otherwise.
func (o *AdditionalMeasurement) GetWrgDest() AddressList {
	if o == nil || o.WrgDest == nil {
		var ret AddressList
		return ret
	}
	return *o.WrgDest
}

// GetWrgDestOk returns a tuple with the WrgDest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalMeasurement) GetWrgDestOk() (*AddressList, bool) {
	if o == nil || o.WrgDest == nil {
		return nil, false
	}
	return o.WrgDest, true
}

// HasWrgDest returns a boolean if a field has been set.
func (o *AdditionalMeasurement) HasWrgDest() bool {
	if o != nil && o.WrgDest != nil {
		return true
	}

	return false
}

// SetWrgDest gets a reference to the given AddressList and assigns it to the WrgDest field.
func (o *AdditionalMeasurement) SetWrgDest(v AddressList) {
	o.WrgDest = &v
}

// GetCircums returns the Circums field value if set, zero value otherwise.
func (o *AdditionalMeasurement) GetCircums() []CircumstanceDescription {
	if o == nil || o.Circums == nil {
		var ret []CircumstanceDescription
		return ret
	}
	return *o.Circums
}

// GetCircumsOk returns a tuple with the Circums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalMeasurement) GetCircumsOk() (*[]CircumstanceDescription, bool) {
	if o == nil || o.Circums == nil {
		return nil, false
	}
	return o.Circums, true
}

// HasCircums returns a boolean if a field has been set.
func (o *AdditionalMeasurement) HasCircums() bool {
	if o != nil && o.Circums != nil {
		return true
	}

	return false
}

// SetCircums gets a reference to the given []CircumstanceDescription and assigns it to the Circums field.
func (o *AdditionalMeasurement) SetCircums(v []CircumstanceDescription) {
	o.Circums = &v
}

func (o AdditionalMeasurement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnexpLoc != nil {
		toSerialize["unexpLoc"] = o.UnexpLoc
	}
	if o.UnexpFlowTeps != nil {
		toSerialize["unexpFlowTeps"] = o.UnexpFlowTeps
	}
	if o.UnexpWakes != nil {
		toSerialize["unexpWakes"] = o.UnexpWakes
	}
	if o.DdosAttack != nil {
		toSerialize["ddosAttack"] = o.DdosAttack
	}
	if o.WrgDest != nil {
		toSerialize["wrgDest"] = o.WrgDest
	}
	if o.Circums != nil {
		toSerialize["circums"] = o.Circums
	}
	return json.Marshal(toSerialize)
}

type NullableAdditionalMeasurement struct {
	value *AdditionalMeasurement
	isSet bool
}

func (v NullableAdditionalMeasurement) Get() *AdditionalMeasurement {
	return v.value
}

func (v *NullableAdditionalMeasurement) Set(val *AdditionalMeasurement) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalMeasurement) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalMeasurement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalMeasurement(val *AdditionalMeasurement) *NullableAdditionalMeasurement {
	return &NullableAdditionalMeasurement{value: val, isSet: true}
}

func (v NullableAdditionalMeasurement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalMeasurement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


