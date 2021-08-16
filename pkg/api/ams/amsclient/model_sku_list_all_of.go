/*
 * Account Management Service API
 *
 * Manage user subscriptions and clusters
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package amsclient

import (
	"encoding/json"
)

// SkuListAllOf struct for SkuListAllOf
type SkuListAllOf struct {
	Items *[]SKU `json:"items,omitempty"`
}

// NewSkuListAllOf instantiates a new SkuListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListAllOf() *SkuListAllOf {
	this := SkuListAllOf{}
	return &this
}

// NewSkuListAllOfWithDefaults instantiates a new SkuListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListAllOfWithDefaults() *SkuListAllOf {
	this := SkuListAllOf{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SkuListAllOf) GetItems() []SKU {
	if o == nil || o.Items == nil {
		var ret []SKU
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListAllOf) GetItemsOk() (*[]SKU, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SkuListAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SKU and assigns it to the Items field.
func (o *SkuListAllOf) SetItems(v []SKU) {
	o.Items = &v
}

func (o SkuListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableSkuListAllOf struct {
	value *SkuListAllOf
	isSet bool
}

func (v NullableSkuListAllOf) Get() *SkuListAllOf {
	return v.value
}

func (v *NullableSkuListAllOf) Set(val *SkuListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListAllOf(val *SkuListAllOf) *NullableSkuListAllOf {
	return &NullableSkuListAllOf{value: val, isSet: true}
}

func (v NullableSkuListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
