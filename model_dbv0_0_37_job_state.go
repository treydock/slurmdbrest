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

// Dbv0037JobState State properties of job
type Dbv0037JobState struct {
	// Current state of job
	Current *string `json:"current,omitempty"`
	// Last reason job didn't run
	Reason *string `json:"reason,omitempty"`
}

// NewDbv0037JobState instantiates a new Dbv0037JobState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobState() *Dbv0037JobState {
	this := Dbv0037JobState{}
	return &this
}

// NewDbv0037JobStateWithDefaults instantiates a new Dbv0037JobState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobStateWithDefaults() *Dbv0037JobState {
	this := Dbv0037JobState{}
	return &this
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *Dbv0037JobState) GetCurrent() string {
	if o == nil || o.Current == nil {
		var ret string
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobState) GetCurrentOk() (*string, bool) {
	if o == nil || o.Current == nil {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *Dbv0037JobState) HasCurrent() bool {
	if o != nil && o.Current != nil {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given string and assigns it to the Current field.
func (o *Dbv0037JobState) SetCurrent(v string) {
	o.Current = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *Dbv0037JobState) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobState) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *Dbv0037JobState) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *Dbv0037JobState) SetReason(v string) {
	o.Reason = &v
}

func (o Dbv0037JobState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Current != nil {
		toSerialize["current"] = o.Current
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037JobState struct {
	value *Dbv0037JobState
	isSet bool
}

func (v NullableDbv0037JobState) Get() *Dbv0037JobState {
	return v.value
}

func (v *NullableDbv0037JobState) Set(val *Dbv0037JobState) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobState) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobState(val *Dbv0037JobState) *NullableDbv0037JobState {
	return &NullableDbv0037JobState{value: val, isSet: true}
}

func (v NullableDbv0037JobState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


