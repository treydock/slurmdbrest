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

// Dbv0037JobStepStepTres TRES usage
type Dbv0037JobStepStepTres struct {
	Requested *Dbv0037JobStepStepTresRequested `json:"requested,omitempty"`
	Consumed *Dbv0037JobStepStepTresRequested `json:"consumed,omitempty"`
	// TRES list of attributes
	Allocated *[]map[string]interface{} `json:"allocated,omitempty"`
}

// NewDbv0037JobStepStepTres instantiates a new Dbv0037JobStepStepTres object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobStepStepTres() *Dbv0037JobStepStepTres {
	this := Dbv0037JobStepStepTres{}
	return &this
}

// NewDbv0037JobStepStepTresWithDefaults instantiates a new Dbv0037JobStepStepTres object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobStepStepTresWithDefaults() *Dbv0037JobStepStepTres {
	this := Dbv0037JobStepStepTres{}
	return &this
}

// GetRequested returns the Requested field value if set, zero value otherwise.
func (o *Dbv0037JobStepStepTres) GetRequested() Dbv0037JobStepStepTresRequested {
	if o == nil || o.Requested == nil {
		var ret Dbv0037JobStepStepTresRequested
		return ret
	}
	return *o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepStepTres) GetRequestedOk() (*Dbv0037JobStepStepTresRequested, bool) {
	if o == nil || o.Requested == nil {
		return nil, false
	}
	return o.Requested, true
}

// HasRequested returns a boolean if a field has been set.
func (o *Dbv0037JobStepStepTres) HasRequested() bool {
	if o != nil && o.Requested != nil {
		return true
	}

	return false
}

// SetRequested gets a reference to the given Dbv0037JobStepStepTresRequested and assigns it to the Requested field.
func (o *Dbv0037JobStepStepTres) SetRequested(v Dbv0037JobStepStepTresRequested) {
	o.Requested = &v
}

// GetConsumed returns the Consumed field value if set, zero value otherwise.
func (o *Dbv0037JobStepStepTres) GetConsumed() Dbv0037JobStepStepTresRequested {
	if o == nil || o.Consumed == nil {
		var ret Dbv0037JobStepStepTresRequested
		return ret
	}
	return *o.Consumed
}

// GetConsumedOk returns a tuple with the Consumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepStepTres) GetConsumedOk() (*Dbv0037JobStepStepTresRequested, bool) {
	if o == nil || o.Consumed == nil {
		return nil, false
	}
	return o.Consumed, true
}

// HasConsumed returns a boolean if a field has been set.
func (o *Dbv0037JobStepStepTres) HasConsumed() bool {
	if o != nil && o.Consumed != nil {
		return true
	}

	return false
}

// SetConsumed gets a reference to the given Dbv0037JobStepStepTresRequested and assigns it to the Consumed field.
func (o *Dbv0037JobStepStepTres) SetConsumed(v Dbv0037JobStepStepTresRequested) {
	o.Consumed = &v
}

// GetAllocated returns the Allocated field value if set, zero value otherwise.
func (o *Dbv0037JobStepStepTres) GetAllocated() []map[string]interface{} {
	if o == nil || o.Allocated == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Allocated
}

// GetAllocatedOk returns a tuple with the Allocated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepStepTres) GetAllocatedOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Allocated == nil {
		return nil, false
	}
	return o.Allocated, true
}

// HasAllocated returns a boolean if a field has been set.
func (o *Dbv0037JobStepStepTres) HasAllocated() bool {
	if o != nil && o.Allocated != nil {
		return true
	}

	return false
}

// SetAllocated gets a reference to the given []map[string]interface{} and assigns it to the Allocated field.
func (o *Dbv0037JobStepStepTres) SetAllocated(v []map[string]interface{}) {
	o.Allocated = &v
}

func (o Dbv0037JobStepStepTres) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Requested != nil {
		toSerialize["requested"] = o.Requested
	}
	if o.Consumed != nil {
		toSerialize["consumed"] = o.Consumed
	}
	if o.Allocated != nil {
		toSerialize["allocated"] = o.Allocated
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037JobStepStepTres struct {
	value *Dbv0037JobStepStepTres
	isSet bool
}

func (v NullableDbv0037JobStepStepTres) Get() *Dbv0037JobStepStepTres {
	return v.value
}

func (v *NullableDbv0037JobStepStepTres) Set(val *Dbv0037JobStepStepTres) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobStepStepTres) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobStepStepTres) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobStepStepTres(val *Dbv0037JobStepStepTres) *NullableDbv0037JobStepStepTres {
	return &NullableDbv0037JobStepStepTres{value: val, isSet: true}
}

func (v NullableDbv0037JobStepStepTres) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobStepStepTres) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


