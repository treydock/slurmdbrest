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

// Dbv0037Wckey struct for Dbv0037Wckey
type Dbv0037Wckey struct {
	// List of assigned accounts
	Accounts *[]string `json:"accounts,omitempty"`
	// Cluster name
	Cluster *string `json:"cluster,omitempty"`
	// wckey database unique id
	Id *int32 `json:"id,omitempty"`
	// wckey name
	Name *string `json:"name,omitempty"`
	// wckey user
	User *string `json:"user,omitempty"`
	// List of properties of wckey
	Flags *[]string `json:"flags,omitempty"`
}

// NewDbv0037Wckey instantiates a new Dbv0037Wckey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037Wckey() *Dbv0037Wckey {
	this := Dbv0037Wckey{}
	return &this
}

// NewDbv0037WckeyWithDefaults instantiates a new Dbv0037Wckey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037WckeyWithDefaults() *Dbv0037Wckey {
	this := Dbv0037Wckey{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *Dbv0037Wckey) GetAccounts() []string {
	if o == nil || o.Accounts == nil {
		var ret []string
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037Wckey) GetAccountsOk() (*[]string, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *Dbv0037Wckey) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []string and assigns it to the Accounts field.
func (o *Dbv0037Wckey) SetAccounts(v []string) {
	o.Accounts = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *Dbv0037Wckey) GetCluster() string {
	if o == nil || o.Cluster == nil {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037Wckey) GetClusterOk() (*string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *Dbv0037Wckey) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *Dbv0037Wckey) SetCluster(v string) {
	o.Cluster = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Dbv0037Wckey) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037Wckey) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Dbv0037Wckey) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Dbv0037Wckey) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dbv0037Wckey) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037Wckey) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dbv0037Wckey) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Dbv0037Wckey) SetName(v string) {
	o.Name = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Dbv0037Wckey) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037Wckey) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Dbv0037Wckey) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *Dbv0037Wckey) SetUser(v string) {
	o.User = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *Dbv0037Wckey) GetFlags() []string {
	if o == nil || o.Flags == nil {
		var ret []string
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037Wckey) GetFlagsOk() (*[]string, bool) {
	if o == nil || o.Flags == nil {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *Dbv0037Wckey) HasFlags() bool {
	if o != nil && o.Flags != nil {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *Dbv0037Wckey) SetFlags(v []string) {
	o.Flags = &v
}

func (o Dbv0037Wckey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Flags != nil {
		toSerialize["flags"] = o.Flags
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037Wckey struct {
	value *Dbv0037Wckey
	isSet bool
}

func (v NullableDbv0037Wckey) Get() *Dbv0037Wckey {
	return v.value
}

func (v *NullableDbv0037Wckey) Set(val *Dbv0037Wckey) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037Wckey) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037Wckey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037Wckey(val *Dbv0037Wckey) *NullableDbv0037Wckey {
	return &NullableDbv0037Wckey{value: val, isSet: true}
}

func (v NullableDbv0037Wckey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037Wckey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


