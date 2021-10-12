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

// Dbv0037AssociationMaxTresPer Max TRES per settings
type Dbv0037AssociationMaxTresPer struct {
	// TRES list of attributes
	Job *[]map[string]interface{} `json:"job,omitempty"`
	// TRES list of attributes
	Node *[]map[string]interface{} `json:"node,omitempty"`
}

// NewDbv0037AssociationMaxTresPer instantiates a new Dbv0037AssociationMaxTresPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037AssociationMaxTresPer() *Dbv0037AssociationMaxTresPer {
	this := Dbv0037AssociationMaxTresPer{}
	return &this
}

// NewDbv0037AssociationMaxTresPerWithDefaults instantiates a new Dbv0037AssociationMaxTresPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037AssociationMaxTresPerWithDefaults() *Dbv0037AssociationMaxTresPer {
	this := Dbv0037AssociationMaxTresPer{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *Dbv0037AssociationMaxTresPer) GetJob() []map[string]interface{} {
	if o == nil || o.Job == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037AssociationMaxTresPer) GetJobOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Job == nil {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *Dbv0037AssociationMaxTresPer) HasJob() bool {
	if o != nil && o.Job != nil {
		return true
	}

	return false
}

// SetJob gets a reference to the given []map[string]interface{} and assigns it to the Job field.
func (o *Dbv0037AssociationMaxTresPer) SetJob(v []map[string]interface{}) {
	o.Job = &v
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *Dbv0037AssociationMaxTresPer) GetNode() []map[string]interface{} {
	if o == nil || o.Node == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037AssociationMaxTresPer) GetNodeOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Node == nil {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *Dbv0037AssociationMaxTresPer) HasNode() bool {
	if o != nil && o.Node != nil {
		return true
	}

	return false
}

// SetNode gets a reference to the given []map[string]interface{} and assigns it to the Node field.
func (o *Dbv0037AssociationMaxTresPer) SetNode(v []map[string]interface{}) {
	o.Node = &v
}

func (o Dbv0037AssociationMaxTresPer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Job != nil {
		toSerialize["job"] = o.Job
	}
	if o.Node != nil {
		toSerialize["node"] = o.Node
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037AssociationMaxTresPer struct {
	value *Dbv0037AssociationMaxTresPer
	isSet bool
}

func (v NullableDbv0037AssociationMaxTresPer) Get() *Dbv0037AssociationMaxTresPer {
	return v.value
}

func (v *NullableDbv0037AssociationMaxTresPer) Set(val *Dbv0037AssociationMaxTresPer) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037AssociationMaxTresPer) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037AssociationMaxTresPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037AssociationMaxTresPer(val *Dbv0037AssociationMaxTresPer) *NullableDbv0037AssociationMaxTresPer {
	return &NullableDbv0037AssociationMaxTresPer{value: val, isSet: true}
}

func (v NullableDbv0037AssociationMaxTresPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037AssociationMaxTresPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


