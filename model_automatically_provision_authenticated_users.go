/*
 * Seq API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package seq

import (
	"encoding/json"
)

// AutomaticallyProvisionAuthenticatedUsers struct for AutomaticallyProvisionAuthenticatedUsers
type AutomaticallyProvisionAuthenticatedUsers struct {
	Id string `json:"Id"`
	Name string `json:"Name"`
	Value *bool `json:"Value,omitempty"`
}

// NewAutomaticallyProvisionAuthenticatedUsers instantiates a new AutomaticallyProvisionAuthenticatedUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomaticallyProvisionAuthenticatedUsers(id string, name string, ) *AutomaticallyProvisionAuthenticatedUsers {
	this := AutomaticallyProvisionAuthenticatedUsers{}
	this.Id = id
	this.Name = name
	return &this
}

// NewAutomaticallyProvisionAuthenticatedUsersWithDefaults instantiates a new AutomaticallyProvisionAuthenticatedUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomaticallyProvisionAuthenticatedUsersWithDefaults() *AutomaticallyProvisionAuthenticatedUsers {
	this := AutomaticallyProvisionAuthenticatedUsers{}
	return &this
}

// GetId returns the Id field value
func (o *AutomaticallyProvisionAuthenticatedUsers) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AutomaticallyProvisionAuthenticatedUsers) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AutomaticallyProvisionAuthenticatedUsers) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AutomaticallyProvisionAuthenticatedUsers) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutomaticallyProvisionAuthenticatedUsers) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AutomaticallyProvisionAuthenticatedUsers) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AutomaticallyProvisionAuthenticatedUsers) GetValue() bool {
	if o == nil || o.Value == nil {
		var ret bool
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticallyProvisionAuthenticatedUsers) GetValueOk() (*bool, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AutomaticallyProvisionAuthenticatedUsers) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given bool and assigns it to the Value field.
func (o *AutomaticallyProvisionAuthenticatedUsers) SetValue(v bool) {
	o.Value = &v
}

func (o AutomaticallyProvisionAuthenticatedUsers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Id"] = o.Id
	}
	if true {
		toSerialize["Name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableAutomaticallyProvisionAuthenticatedUsers struct {
	value *AutomaticallyProvisionAuthenticatedUsers
	isSet bool
}

func (v NullableAutomaticallyProvisionAuthenticatedUsers) Get() *AutomaticallyProvisionAuthenticatedUsers {
	return v.value
}

func (v *NullableAutomaticallyProvisionAuthenticatedUsers) Set(val *AutomaticallyProvisionAuthenticatedUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomaticallyProvisionAuthenticatedUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomaticallyProvisionAuthenticatedUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomaticallyProvisionAuthenticatedUsers(val *AutomaticallyProvisionAuthenticatedUsers) *NullableAutomaticallyProvisionAuthenticatedUsers {
	return &NullableAutomaticallyProvisionAuthenticatedUsers{value: val, isSet: true}
}

func (v NullableAutomaticallyProvisionAuthenticatedUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutomaticallyProvisionAuthenticatedUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


