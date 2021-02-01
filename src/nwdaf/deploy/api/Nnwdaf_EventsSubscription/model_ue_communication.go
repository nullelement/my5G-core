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

// UeCommunication struct for UeCommunication
type UeCommunication struct {
	CommDur map[string]interface{} `json:"commDur"`
	CommDurVariance *map[string]interface{} `json:"commDurVariance,omitempty"`
	PerioTime *map[string]interface{} `json:"perioTime,omitempty"`
	PerioTimeVariance *map[string]interface{} `json:"perioTimeVariance,omitempty"`
	Ts *map[string]interface{} `json:"ts,omitempty"`
	TsVariance *map[string]interface{} `json:"tsVariance,omitempty"`
	RecurringTime *map[string]interface{} `json:"recurringTime,omitempty"`
	TrafChar TrafficCharacterization `json:"trafChar"`
	Ratio *map[string]interface{} `json:"ratio,omitempty"`
	Confidence *map[string]interface{} `json:"confidence,omitempty"`
}

// NewUeCommunication instantiates a new UeCommunication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeCommunication(commDur map[string]interface{}, trafChar TrafficCharacterization, ) *UeCommunication {
	this := UeCommunication{}
	this.CommDur = commDur
	this.TrafChar = trafChar
	return &this
}

// NewUeCommunicationWithDefaults instantiates a new UeCommunication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeCommunicationWithDefaults() *UeCommunication {
	this := UeCommunication{}
	return &this
}

// GetCommDur returns the CommDur field value
func (o *UeCommunication) GetCommDur() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}

	return o.CommDur
}

// GetCommDurOk returns a tuple with the CommDur field value
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetCommDurOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CommDur, true
}

// SetCommDur sets field value
func (o *UeCommunication) SetCommDur(v map[string]interface{}) {
	o.CommDur = v
}

// GetCommDurVariance returns the CommDurVariance field value if set, zero value otherwise.
func (o *UeCommunication) GetCommDurVariance() map[string]interface{} {
	if o == nil || o.CommDurVariance == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.CommDurVariance
}

// GetCommDurVarianceOk returns a tuple with the CommDurVariance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetCommDurVarianceOk() (*map[string]interface{}, bool) {
	if o == nil || o.CommDurVariance == nil {
		return nil, false
	}
	return o.CommDurVariance, true
}

// HasCommDurVariance returns a boolean if a field has been set.
func (o *UeCommunication) HasCommDurVariance() bool {
	if o != nil && o.CommDurVariance != nil {
		return true
	}

	return false
}

// SetCommDurVariance gets a reference to the given map[string]interface{} and assigns it to the CommDurVariance field.
func (o *UeCommunication) SetCommDurVariance(v map[string]interface{}) {
	o.CommDurVariance = &v
}

// GetPerioTime returns the PerioTime field value if set, zero value otherwise.
func (o *UeCommunication) GetPerioTime() map[string]interface{} {
	if o == nil || o.PerioTime == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.PerioTime
}

// GetPerioTimeOk returns a tuple with the PerioTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetPerioTimeOk() (*map[string]interface{}, bool) {
	if o == nil || o.PerioTime == nil {
		return nil, false
	}
	return o.PerioTime, true
}

// HasPerioTime returns a boolean if a field has been set.
func (o *UeCommunication) HasPerioTime() bool {
	if o != nil && o.PerioTime != nil {
		return true
	}

	return false
}

// SetPerioTime gets a reference to the given map[string]interface{} and assigns it to the PerioTime field.
func (o *UeCommunication) SetPerioTime(v map[string]interface{}) {
	o.PerioTime = &v
}

// GetPerioTimeVariance returns the PerioTimeVariance field value if set, zero value otherwise.
func (o *UeCommunication) GetPerioTimeVariance() map[string]interface{} {
	if o == nil || o.PerioTimeVariance == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.PerioTimeVariance
}

// GetPerioTimeVarianceOk returns a tuple with the PerioTimeVariance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetPerioTimeVarianceOk() (*map[string]interface{}, bool) {
	if o == nil || o.PerioTimeVariance == nil {
		return nil, false
	}
	return o.PerioTimeVariance, true
}

// HasPerioTimeVariance returns a boolean if a field has been set.
func (o *UeCommunication) HasPerioTimeVariance() bool {
	if o != nil && o.PerioTimeVariance != nil {
		return true
	}

	return false
}

// SetPerioTimeVariance gets a reference to the given map[string]interface{} and assigns it to the PerioTimeVariance field.
func (o *UeCommunication) SetPerioTimeVariance(v map[string]interface{}) {
	o.PerioTimeVariance = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *UeCommunication) GetTs() map[string]interface{} {
	if o == nil || o.Ts == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetTsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Ts == nil {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *UeCommunication) HasTs() bool {
	if o != nil && o.Ts != nil {
		return true
	}

	return false
}

// SetTs gets a reference to the given map[string]interface{} and assigns it to the Ts field.
func (o *UeCommunication) SetTs(v map[string]interface{}) {
	o.Ts = &v
}

// GetTsVariance returns the TsVariance field value if set, zero value otherwise.
func (o *UeCommunication) GetTsVariance() map[string]interface{} {
	if o == nil || o.TsVariance == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.TsVariance
}

// GetTsVarianceOk returns a tuple with the TsVariance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetTsVarianceOk() (*map[string]interface{}, bool) {
	if o == nil || o.TsVariance == nil {
		return nil, false
	}
	return o.TsVariance, true
}

// HasTsVariance returns a boolean if a field has been set.
func (o *UeCommunication) HasTsVariance() bool {
	if o != nil && o.TsVariance != nil {
		return true
	}

	return false
}

// SetTsVariance gets a reference to the given map[string]interface{} and assigns it to the TsVariance field.
func (o *UeCommunication) SetTsVariance(v map[string]interface{}) {
	o.TsVariance = &v
}

// GetRecurringTime returns the RecurringTime field value if set, zero value otherwise.
func (o *UeCommunication) GetRecurringTime() map[string]interface{} {
	if o == nil || o.RecurringTime == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.RecurringTime
}

// GetRecurringTimeOk returns a tuple with the RecurringTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetRecurringTimeOk() (*map[string]interface{}, bool) {
	if o == nil || o.RecurringTime == nil {
		return nil, false
	}
	return o.RecurringTime, true
}

// HasRecurringTime returns a boolean if a field has been set.
func (o *UeCommunication) HasRecurringTime() bool {
	if o != nil && o.RecurringTime != nil {
		return true
	}

	return false
}

// SetRecurringTime gets a reference to the given map[string]interface{} and assigns it to the RecurringTime field.
func (o *UeCommunication) SetRecurringTime(v map[string]interface{}) {
	o.RecurringTime = &v
}

// GetTrafChar returns the TrafChar field value
func (o *UeCommunication) GetTrafChar() TrafficCharacterization {
	if o == nil  {
		var ret TrafficCharacterization
		return ret
	}

	return o.TrafChar
}

// GetTrafCharOk returns a tuple with the TrafChar field value
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetTrafCharOk() (*TrafficCharacterization, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TrafChar, true
}

// SetTrafChar sets field value
func (o *UeCommunication) SetTrafChar(v TrafficCharacterization) {
	o.TrafChar = v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *UeCommunication) GetRatio() map[string]interface{} {
	if o == nil || o.Ratio == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetRatioOk() (*map[string]interface{}, bool) {
	if o == nil || o.Ratio == nil {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *UeCommunication) HasRatio() bool {
	if o != nil && o.Ratio != nil {
		return true
	}

	return false
}

// SetRatio gets a reference to the given map[string]interface{} and assigns it to the Ratio field.
func (o *UeCommunication) SetRatio(v map[string]interface{}) {
	o.Ratio = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *UeCommunication) GetConfidence() map[string]interface{} {
	if o == nil || o.Confidence == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetConfidenceOk() (*map[string]interface{}, bool) {
	if o == nil || o.Confidence == nil {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *UeCommunication) HasConfidence() bool {
	if o != nil && o.Confidence != nil {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given map[string]interface{} and assigns it to the Confidence field.
func (o *UeCommunication) SetConfidence(v map[string]interface{}) {
	o.Confidence = &v
}

func (o UeCommunication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["commDur"] = o.CommDur
	}
	if o.CommDurVariance != nil {
		toSerialize["commDurVariance"] = o.CommDurVariance
	}
	if o.PerioTime != nil {
		toSerialize["perioTime"] = o.PerioTime
	}
	if o.PerioTimeVariance != nil {
		toSerialize["perioTimeVariance"] = o.PerioTimeVariance
	}
	if o.Ts != nil {
		toSerialize["ts"] = o.Ts
	}
	if o.TsVariance != nil {
		toSerialize["tsVariance"] = o.TsVariance
	}
	if o.RecurringTime != nil {
		toSerialize["recurringTime"] = o.RecurringTime
	}
	if true {
		toSerialize["trafChar"] = o.TrafChar
	}
	if o.Ratio != nil {
		toSerialize["ratio"] = o.Ratio
	}
	if o.Confidence != nil {
		toSerialize["confidence"] = o.Confidence
	}
	return json.Marshal(toSerialize)
}

type NullableUeCommunication struct {
	value *UeCommunication
	isSet bool
}

func (v NullableUeCommunication) Get() *UeCommunication {
	return v.value
}

func (v *NullableUeCommunication) Set(val *UeCommunication) {
	v.value = val
	v.isSet = true
}

func (v NullableUeCommunication) IsSet() bool {
	return v.isSet
}

func (v *NullableUeCommunication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeCommunication(val *UeCommunication) *NullableUeCommunication {
	return &NullableUeCommunication{value: val, isSet: true}
}

func (v NullableUeCommunication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeCommunication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

