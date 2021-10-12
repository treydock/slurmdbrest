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

// Dbv0037JobStepStepTresRequested TRES requested for job
type Dbv0037JobStepStepTresRequested struct {
	// TRES list of attributes
	Average *[]map[string]interface{} `json:"average,omitempty"`
	// TRES list of attributes
	Max *[]map[string]interface{} `json:"max,omitempty"`
	// TRES list of attributes
	Min *[]map[string]interface{} `json:"min,omitempty"`
	// TRES list of attributes
	Total *[]map[string]interface{} `json:"total,omitempty"`
}

// NewDbv0037JobStepStepTresRequested instantiates a new Dbv0037JobStepStepTresRequested object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobStepStepTresRequested() *Dbv0037JobStepStepTresRequested {
	this := Dbv0037JobStepStepTresRequested{}
	return &this
}

// NewDbv0037JobStepStepTresRequestedWithDefaults instantiates a new Dbv0037JobStepStepTresRequested object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobStepStepTresRequestedWithDefaults() *Dbv0037JobStepStepTresRequested {
	this := Dbv0037JobStepStepTresRequested{}
	return &this
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *Dbv0037JobStepStepTresRequested) GetAverage() []map[string]interface{} {
	if o == nil || o.Average == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepStepTresRequested) GetAverageOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Average == nil {
		return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *Dbv0037JobStepStepTresRequested) HasAverage() bool {
	if o != nil && o.Average != nil {
		return true
	}

	return false
}

// SetAverage gets a reference to the given []map[string]interface{} and assigns it to the Average field.
func (o *Dbv0037JobStepStepTresRequested) SetAverage(v []map[string]interface{}) {
	o.Average = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *Dbv0037JobStepStepTresRequested) GetMax() []map[string]interface{} {
	if o == nil || o.Max == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepStepTresRequested) GetMaxOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *Dbv0037JobStepStepTresRequested) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given []map[string]interface{} and assigns it to the Max field.
func (o *Dbv0037JobStepStepTresRequested) SetMax(v []map[string]interface{}) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *Dbv0037JobStepStepTresRequested) GetMin() []map[string]interface{} {
	if o == nil || o.Min == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepStepTresRequested) GetMinOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *Dbv0037JobStepStepTresRequested) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given []map[string]interface{} and assigns it to the Min field.
func (o *Dbv0037JobStepStepTresRequested) SetMin(v []map[string]interface{}) {
	o.Min = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *Dbv0037JobStepStepTresRequested) GetTotal() []map[string]interface{} {
	if o == nil || o.Total == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepStepTresRequested) GetTotalOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *Dbv0037JobStepStepTresRequested) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given []map[string]interface{} and assigns it to the Total field.
func (o *Dbv0037JobStepStepTresRequested) SetTotal(v []map[string]interface{}) {
	o.Total = &v
}

func (o Dbv0037JobStepStepTresRequested) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Average != nil {
		toSerialize["average"] = o.Average
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037JobStepStepTresRequested struct {
	value *Dbv0037JobStepStepTresRequested
	isSet bool
}

func (v NullableDbv0037JobStepStepTresRequested) Get() *Dbv0037JobStepStepTresRequested {
	return v.value
}

func (v *NullableDbv0037JobStepStepTresRequested) Set(val *Dbv0037JobStepStepTresRequested) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobStepStepTresRequested) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobStepStepTresRequested) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobStepStepTresRequested(val *Dbv0037JobStepStepTresRequested) *NullableDbv0037JobStepStepTresRequested {
	return &NullableDbv0037JobStepStepTresRequested{value: val, isSet: true}
}

func (v NullableDbv0037JobStepStepTresRequested) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobStepStepTresRequested) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


