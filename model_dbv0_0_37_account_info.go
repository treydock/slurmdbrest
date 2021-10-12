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

// Dbv0037AccountInfo struct for Dbv0037AccountInfo
type Dbv0037AccountInfo struct {
	// Slurm errors
	Errors *[]Dbv0037Error `json:"errors,omitempty"`
	// List of accounts
	Accounts *[]Dbv0037Account `json:"accounts,omitempty"`
}

// NewDbv0037AccountInfo instantiates a new Dbv0037AccountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037AccountInfo() *Dbv0037AccountInfo {
	this := Dbv0037AccountInfo{}
	return &this
}

// NewDbv0037AccountInfoWithDefaults instantiates a new Dbv0037AccountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037AccountInfoWithDefaults() *Dbv0037AccountInfo {
	this := Dbv0037AccountInfo{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0037AccountInfo) GetErrors() []Dbv0037Error {
	if o == nil || o.Errors == nil {
		var ret []Dbv0037Error
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037AccountInfo) GetErrorsOk() (*[]Dbv0037Error, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0037AccountInfo) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0037Error and assigns it to the Errors field.
func (o *Dbv0037AccountInfo) SetErrors(v []Dbv0037Error) {
	o.Errors = &v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *Dbv0037AccountInfo) GetAccounts() []Dbv0037Account {
	if o == nil || o.Accounts == nil {
		var ret []Dbv0037Account
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037AccountInfo) GetAccountsOk() (*[]Dbv0037Account, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *Dbv0037AccountInfo) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []Dbv0037Account and assigns it to the Accounts field.
func (o *Dbv0037AccountInfo) SetAccounts(v []Dbv0037Account) {
	o.Accounts = &v
}

func (o Dbv0037AccountInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037AccountInfo struct {
	value *Dbv0037AccountInfo
	isSet bool
}

func (v NullableDbv0037AccountInfo) Get() *Dbv0037AccountInfo {
	return v.value
}

func (v *NullableDbv0037AccountInfo) Set(val *Dbv0037AccountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037AccountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037AccountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037AccountInfo(val *Dbv0037AccountInfo) *NullableDbv0037AccountInfo {
	return &NullableDbv0037AccountInfo{value: val, isSet: true}
}

func (v NullableDbv0037AccountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037AccountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


