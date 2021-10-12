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

// Dbv0037JobStepTasks Task properties
type Dbv0037JobStepTasks struct {
	// Number of tasks in step
	Count *int32 `json:"count,omitempty"`
}

// NewDbv0037JobStepTasks instantiates a new Dbv0037JobStepTasks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobStepTasks() *Dbv0037JobStepTasks {
	this := Dbv0037JobStepTasks{}
	return &this
}

// NewDbv0037JobStepTasksWithDefaults instantiates a new Dbv0037JobStepTasks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobStepTasksWithDefaults() *Dbv0037JobStepTasks {
	this := Dbv0037JobStepTasks{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Dbv0037JobStepTasks) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepTasks) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Dbv0037JobStepTasks) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *Dbv0037JobStepTasks) SetCount(v int32) {
	o.Count = &v
}

func (o Dbv0037JobStepTasks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037JobStepTasks struct {
	value *Dbv0037JobStepTasks
	isSet bool
}

func (v NullableDbv0037JobStepTasks) Get() *Dbv0037JobStepTasks {
	return v.value
}

func (v *NullableDbv0037JobStepTasks) Set(val *Dbv0037JobStepTasks) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobStepTasks) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobStepTasks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobStepTasks(val *Dbv0037JobStepTasks) *NullableDbv0037JobStepTasks {
	return &NullableDbv0037JobStepTasks{value: val, isSet: true}
}

func (v NullableDbv0037JobStepTasks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobStepTasks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


