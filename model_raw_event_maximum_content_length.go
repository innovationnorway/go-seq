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

// RawEventMaximumContentLength struct for RawEventMaximumContentLength
type RawEventMaximumContentLength struct {
	Id *string `json:"Id,omitempty"`
	Name *string `json:"Name,omitempty"`
	Value *int32 `json:"Value,omitempty"`
}

// NewRawEventMaximumContentLength instantiates a new RawEventMaximumContentLength object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRawEventMaximumContentLength() *RawEventMaximumContentLength {
	this := RawEventMaximumContentLength{}
	return &this
}

// NewRawEventMaximumContentLengthWithDefaults instantiates a new RawEventMaximumContentLength object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRawEventMaximumContentLengthWithDefaults() *RawEventMaximumContentLength {
	this := RawEventMaximumContentLength{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RawEventMaximumContentLength) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawEventMaximumContentLength) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RawEventMaximumContentLength) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RawEventMaximumContentLength) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RawEventMaximumContentLength) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawEventMaximumContentLength) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RawEventMaximumContentLength) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RawEventMaximumContentLength) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RawEventMaximumContentLength) GetValue() int32 {
	if o == nil || o.Value == nil {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawEventMaximumContentLength) GetValueOk() (*int32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RawEventMaximumContentLength) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *RawEventMaximumContentLength) SetValue(v int32) {
	o.Value = &v
}

func (o RawEventMaximumContentLength) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableRawEventMaximumContentLength struct {
	value *RawEventMaximumContentLength
	isSet bool
}

func (v NullableRawEventMaximumContentLength) Get() *RawEventMaximumContentLength {
	return v.value
}

func (v *NullableRawEventMaximumContentLength) Set(val *RawEventMaximumContentLength) {
	v.value = val
	v.isSet = true
}

func (v NullableRawEventMaximumContentLength) IsSet() bool {
	return v.isSet
}

func (v *NullableRawEventMaximumContentLength) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRawEventMaximumContentLength(val *RawEventMaximumContentLength) *NullableRawEventMaximumContentLength {
	return &NullableRawEventMaximumContentLength{value: val, isSet: true}
}

func (v NullableRawEventMaximumContentLength) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRawEventMaximumContentLength) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

