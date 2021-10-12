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

// Dbv0037AssociationMaxPer Max per settings
type Dbv0037AssociationMaxPer struct {
	Account *Dbv0037AssociationMaxPerAccount `json:"account,omitempty"`
}

// NewDbv0037AssociationMaxPer instantiates a new Dbv0037AssociationMaxPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037AssociationMaxPer() *Dbv0037AssociationMaxPer {
	this := Dbv0037AssociationMaxPer{}
	return &this
}

// NewDbv0037AssociationMaxPerWithDefaults instantiates a new Dbv0037AssociationMaxPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037AssociationMaxPerWithDefaults() *Dbv0037AssociationMaxPer {
	this := Dbv0037AssociationMaxPer{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Dbv0037AssociationMaxPer) GetAccount() Dbv0037AssociationMaxPerAccount {
	if o == nil || o.Account == nil {
		var ret Dbv0037AssociationMaxPerAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037AssociationMaxPer) GetAccountOk() (*Dbv0037AssociationMaxPerAccount, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Dbv0037AssociationMaxPer) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given Dbv0037AssociationMaxPerAccount and assigns it to the Account field.
func (o *Dbv0037AssociationMaxPer) SetAccount(v Dbv0037AssociationMaxPerAccount) {
	o.Account = &v
}

func (o Dbv0037AssociationMaxPer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037AssociationMaxPer struct {
	value *Dbv0037AssociationMaxPer
	isSet bool
}

func (v NullableDbv0037AssociationMaxPer) Get() *Dbv0037AssociationMaxPer {
	return v.value
}

func (v *NullableDbv0037AssociationMaxPer) Set(val *Dbv0037AssociationMaxPer) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037AssociationMaxPer) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037AssociationMaxPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037AssociationMaxPer(val *Dbv0037AssociationMaxPer) *NullableDbv0037AssociationMaxPer {
	return &NullableDbv0037AssociationMaxPer{value: val, isSet: true}
}

func (v NullableDbv0037AssociationMaxPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037AssociationMaxPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


