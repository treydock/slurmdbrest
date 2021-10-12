/*
 * Slurm Rest API
 *
 * API to access and control Slurm DB.
 *
 * API version: dbv0.0.37
 * Contact: sales@schedmd.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmdbrest

import (
	"encoding/json"
)

// Dbv0037QosLimitsMaxTresMinutes Max TRES minutes settings
type Dbv0037QosLimitsMaxTresMinutes struct {
	Per *Dbv0037QosLimitsMaxTresMinutesPer `json:"per,omitempty"`
}

// NewDbv0037QosLimitsMaxTresMinutes instantiates a new Dbv0037QosLimitsMaxTresMinutes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037QosLimitsMaxTresMinutes() *Dbv0037QosLimitsMaxTresMinutes {
	this := Dbv0037QosLimitsMaxTresMinutes{}
	return &this
}

// NewDbv0037QosLimitsMaxTresMinutesWithDefaults instantiates a new Dbv0037QosLimitsMaxTresMinutes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037QosLimitsMaxTresMinutesWithDefaults() *Dbv0037QosLimitsMaxTresMinutes {
	this := Dbv0037QosLimitsMaxTresMinutes{}
	return &this
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *Dbv0037QosLimitsMaxTresMinutes) GetPer() Dbv0037QosLimitsMaxTresMinutesPer {
	if o == nil || o.Per == nil {
		var ret Dbv0037QosLimitsMaxTresMinutesPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037QosLimitsMaxTresMinutes) GetPerOk() (*Dbv0037QosLimitsMaxTresMinutesPer, bool) {
	if o == nil || o.Per == nil {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *Dbv0037QosLimitsMaxTresMinutes) HasPer() bool {
	if o != nil && o.Per != nil {
		return true
	}

	return false
}

// SetPer gets a reference to the given Dbv0037QosLimitsMaxTresMinutesPer and assigns it to the Per field.
func (o *Dbv0037QosLimitsMaxTresMinutes) SetPer(v Dbv0037QosLimitsMaxTresMinutesPer) {
	o.Per = &v
}

func (o Dbv0037QosLimitsMaxTresMinutes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Per != nil {
		toSerialize["per"] = o.Per
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037QosLimitsMaxTresMinutes struct {
	value *Dbv0037QosLimitsMaxTresMinutes
	isSet bool
}

func (v NullableDbv0037QosLimitsMaxTresMinutes) Get() *Dbv0037QosLimitsMaxTresMinutes {
	return v.value
}

func (v *NullableDbv0037QosLimitsMaxTresMinutes) Set(val *Dbv0037QosLimitsMaxTresMinutes) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037QosLimitsMaxTresMinutes) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037QosLimitsMaxTresMinutes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037QosLimitsMaxTresMinutes(val *Dbv0037QosLimitsMaxTresMinutes) *NullableDbv0037QosLimitsMaxTresMinutes {
	return &NullableDbv0037QosLimitsMaxTresMinutes{value: val, isSet: true}
}

func (v NullableDbv0037QosLimitsMaxTresMinutes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037QosLimitsMaxTresMinutes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


