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

// AutomaticAccessADGroup struct for AutomaticAccessADGroup
type AutomaticAccessADGroup struct {
	Id string `json:"Id"`
	Name string `json:"Name"`
	Value *string `json:"Value,omitempty"`
}

// NewAutomaticAccessADGroup instantiates a new AutomaticAccessADGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomaticAccessADGroup(id string, name string, ) *AutomaticAccessADGroup {
	this := AutomaticAccessADGroup{}
	this.Id = id
	this.Name = name
	return &this
}

// NewAutomaticAccessADGroupWithDefaults instantiates a new AutomaticAccessADGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomaticAccessADGroupWithDefaults() *AutomaticAccessADGroup {
	this := AutomaticAccessADGroup{}
	return &this
}

// GetId returns the Id field value
func (o *AutomaticAccessADGroup) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AutomaticAccessADGroup) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AutomaticAccessADGroup) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AutomaticAccessADGroup) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutomaticAccessADGroup) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AutomaticAccessADGroup) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AutomaticAccessADGroup) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticAccessADGroup) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AutomaticAccessADGroup) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AutomaticAccessADGroup) SetValue(v string) {
	o.Value = &v
}

func (o AutomaticAccessADGroup) MarshalJSON() ([]byte, error) {
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

type NullableAutomaticAccessADGroup struct {
	value *AutomaticAccessADGroup
	isSet bool
}

func (v NullableAutomaticAccessADGroup) Get() *AutomaticAccessADGroup {
	return v.value
}

func (v *NullableAutomaticAccessADGroup) Set(val *AutomaticAccessADGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomaticAccessADGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomaticAccessADGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomaticAccessADGroup(val *AutomaticAccessADGroup) *NullableAutomaticAccessADGroup {
	return &NullableAutomaticAccessADGroup{value: val, isSet: true}
}

func (v NullableAutomaticAccessADGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutomaticAccessADGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


