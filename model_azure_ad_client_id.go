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

// AzureADClientId struct for AzureADClientId
type AzureADClientId struct {
	Id *string `json:"Id,omitempty"`
	Name *string `json:"Name,omitempty"`
	Value *string `json:"Value,omitempty"`
}

// NewAzureADClientId instantiates a new AzureADClientId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureADClientId() *AzureADClientId {
	this := AzureADClientId{}
	return &this
}

// NewAzureADClientIdWithDefaults instantiates a new AzureADClientId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureADClientIdWithDefaults() *AzureADClientId {
	this := AzureADClientId{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzureADClientId) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureADClientId) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzureADClientId) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzureADClientId) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AzureADClientId) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureADClientId) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AzureADClientId) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AzureADClientId) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AzureADClientId) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureADClientId) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AzureADClientId) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AzureADClientId) SetValue(v string) {
	o.Value = &v
}

func (o AzureADClientId) MarshalJSON() ([]byte, error) {
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

type NullableAzureADClientId struct {
	value *AzureADClientId
	isSet bool
}

func (v NullableAzureADClientId) Get() *AzureADClientId {
	return v.value
}

func (v *NullableAzureADClientId) Set(val *AzureADClientId) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureADClientId) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureADClientId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureADClientId(val *AzureADClientId) *NullableAzureADClientId {
	return &NullableAzureADClientId{value: val, isSet: true}
}

func (v NullableAzureADClientId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureADClientId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

