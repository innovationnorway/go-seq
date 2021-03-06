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

// MinimumPasswordLength struct for MinimumPasswordLength
type MinimumPasswordLength struct {
	Id string `json:"Id"`
	Name string `json:"Name"`
	Value *int32 `json:"Value,omitempty"`
}

// NewMinimumPasswordLength instantiates a new MinimumPasswordLength object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinimumPasswordLength(id string, name string, ) *MinimumPasswordLength {
	this := MinimumPasswordLength{}
	this.Id = id
	this.Name = name
	return &this
}

// NewMinimumPasswordLengthWithDefaults instantiates a new MinimumPasswordLength object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinimumPasswordLengthWithDefaults() *MinimumPasswordLength {
	this := MinimumPasswordLength{}
	return &this
}

// GetId returns the Id field value
func (o *MinimumPasswordLength) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MinimumPasswordLength) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MinimumPasswordLength) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *MinimumPasswordLength) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MinimumPasswordLength) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MinimumPasswordLength) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MinimumPasswordLength) GetValue() int32 {
	if o == nil || o.Value == nil {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimumPasswordLength) GetValueOk() (*int32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MinimumPasswordLength) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *MinimumPasswordLength) SetValue(v int32) {
	o.Value = &v
}

func (o MinimumPasswordLength) MarshalJSON() ([]byte, error) {
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

type NullableMinimumPasswordLength struct {
	value *MinimumPasswordLength
	isSet bool
}

func (v NullableMinimumPasswordLength) Get() *MinimumPasswordLength {
	return v.value
}

func (v *NullableMinimumPasswordLength) Set(val *MinimumPasswordLength) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimumPasswordLength) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimumPasswordLength) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimumPasswordLength(val *MinimumPasswordLength) *NullableMinimumPasswordLength {
	return &NullableMinimumPasswordLength{value: val, isSet: true}
}

func (v NullableMinimumPasswordLength) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimumPasswordLength) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


