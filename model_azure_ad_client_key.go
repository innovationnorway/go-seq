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

// AzureADClientKey struct for AzureADClientKey
type AzureADClientKey struct {
	Id string `json:"Id"`
	Name string `json:"Name"`
	Value *string `json:"Value,omitempty"`
}

// NewAzureADClientKey instantiates a new AzureADClientKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureADClientKey(id string, name string, ) *AzureADClientKey {
	this := AzureADClientKey{}
	this.Id = id
	this.Name = name
	return &this
}

// NewAzureADClientKeyWithDefaults instantiates a new AzureADClientKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureADClientKeyWithDefaults() *AzureADClientKey {
	this := AzureADClientKey{}
	return &this
}

// GetId returns the Id field value
func (o *AzureADClientKey) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AzureADClientKey) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AzureADClientKey) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AzureADClientKey) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AzureADClientKey) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AzureADClientKey) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AzureADClientKey) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureADClientKey) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AzureADClientKey) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AzureADClientKey) SetValue(v string) {
	o.Value = &v
}

func (o AzureADClientKey) MarshalJSON() ([]byte, error) {
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

type NullableAzureADClientKey struct {
	value *AzureADClientKey
	isSet bool
}

func (v NullableAzureADClientKey) Get() *AzureADClientKey {
	return v.value
}

func (v *NullableAzureADClientKey) Set(val *AzureADClientKey) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureADClientKey) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureADClientKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureADClientKey(val *AzureADClientKey) *NullableAzureADClientKey {
	return &NullableAzureADClientKey{value: val, isSet: true}
}

func (v NullableAzureADClientKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureADClientKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


