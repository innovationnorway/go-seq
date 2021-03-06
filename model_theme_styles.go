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

// ThemeStyles struct for ThemeStyles
type ThemeStyles struct {
	Id string `json:"Id"`
	Name string `json:"Name"`
	Value *string `json:"Value,omitempty"`
}

// NewThemeStyles instantiates a new ThemeStyles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThemeStyles(id string, name string, ) *ThemeStyles {
	this := ThemeStyles{}
	this.Id = id
	this.Name = name
	return &this
}

// NewThemeStylesWithDefaults instantiates a new ThemeStyles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThemeStylesWithDefaults() *ThemeStyles {
	this := ThemeStyles{}
	return &this
}

// GetId returns the Id field value
func (o *ThemeStyles) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ThemeStyles) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ThemeStyles) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ThemeStyles) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ThemeStyles) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ThemeStyles) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ThemeStyles) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemeStyles) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ThemeStyles) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ThemeStyles) SetValue(v string) {
	o.Value = &v
}

func (o ThemeStyles) MarshalJSON() ([]byte, error) {
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

type NullableThemeStyles struct {
	value *ThemeStyles
	isSet bool
}

func (v NullableThemeStyles) Get() *ThemeStyles {
	return v.value
}

func (v *NullableThemeStyles) Set(val *ThemeStyles) {
	v.value = val
	v.isSet = true
}

func (v NullableThemeStyles) IsSet() bool {
	return v.isSet
}

func (v *NullableThemeStyles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThemeStyles(val *ThemeStyles) *NullableThemeStyles {
	return &NullableThemeStyles{value: val, isSet: true}
}

func (v NullableThemeStyles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThemeStyles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


