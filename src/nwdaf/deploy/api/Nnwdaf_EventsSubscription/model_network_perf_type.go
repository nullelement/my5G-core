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
	"fmt"
)

// NetworkPerfType Possible values are   - GNB_ACTIVE_RATIO: Indicates that the network performance requirement is gNodeB active (i.e. up and running) rate. Indicates the ratio of gNB active (i.e. up and running) number to the total number of gNB   - GNB_COMPUTING_USAGE: Indicates gNodeB computing resource usage.   - GNB_MEMORY_USAGE: Indicates gNodeB memory usage.   - GNB_DISK_USAGE: Indicates gNodeB disk usage.   - NUM_OF_UE: Indicates number of UEs.   - SESS_SUCC_RATIO: Indicates ratio of successful setup of PDU sessions to total PDU session setup attempts.   - SESS_SUCC_RATIO: Indicates Ratio of successful handovers to the total handover attempts.  
type NetworkPerfType struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NetworkPerfType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(NetworkPerfType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NetworkPerfType) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNetworkPerfType struct {
	value *NetworkPerfType
	isSet bool
}

func (v NullableNetworkPerfType) Get() *NetworkPerfType {
	return v.value
}

func (v *NullableNetworkPerfType) Set(val *NetworkPerfType) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkPerfType) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkPerfType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkPerfType(val *NetworkPerfType) *NullableNetworkPerfType {
	return &NullableNetworkPerfType{value: val, isSet: true}
}

func (v NullableNetworkPerfType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkPerfType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


