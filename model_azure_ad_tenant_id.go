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

// AzureADTenantId struct for AzureADTenantId
type AzureADTenantId struct {
	Id *string `json:"Id,omitempty"`
	Name *string `json:"Name,omitempty"`
	Value *string `json:"Value,omitempty"`
}

// NewAzureADTenantId instantiates a new AzureADTenantId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureADTenantId() *AzureADTenantId {
	this := AzureADTenantId{}
	return &this
}

// NewAzureADTenantIdWithDefaults instantiates a new AzureADTenantId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureADTenantIdWithDefaults() *AzureADTenantId {
	this := AzureADTenantId{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzureADTenantId) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureADTenantId) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzureADTenantId) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzureADTenantId) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AzureADTenantId) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureADTenantId) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AzureADTenantId) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AzureADTenantId) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AzureADTenantId) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureADTenantId) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AzureADTenantId) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AzureADTenantId) SetValue(v string) {
	o.Value = &v
}

func (o AzureADTenantId) MarshalJSON() ([]byte, error) {
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

type NullableAzureADTenantId struct {
	value *AzureADTenantId
	isSet bool
}

func (v NullableAzureADTenantId) Get() *AzureADTenantId {
	return v.value
}

func (v *NullableAzureADTenantId) Set(val *AzureADTenantId) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureADTenantId) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureADTenantId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureADTenantId(val *AzureADTenantId) *NullableAzureADTenantId {
	return &NullableAzureADTenantId{value: val, isSet: true}
}

func (v NullableAzureADTenantId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureADTenantId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

