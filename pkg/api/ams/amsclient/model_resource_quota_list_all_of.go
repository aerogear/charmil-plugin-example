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

// ResourceQuotaListAllOf struct for ResourceQuotaListAllOf
type ResourceQuotaListAllOf struct {
	Items *[]ResourceQuota `json:"items,omitempty"`
}

// NewResourceQuotaListAllOf instantiates a new ResourceQuotaListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceQuotaListAllOf() *ResourceQuotaListAllOf {
	this := ResourceQuotaListAllOf{}
	return &this
}

// NewResourceQuotaListAllOfWithDefaults instantiates a new ResourceQuotaListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceQuotaListAllOfWithDefaults() *ResourceQuotaListAllOf {
	this := ResourceQuotaListAllOf{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ResourceQuotaListAllOf) GetItems() []ResourceQuota {
	if o == nil || o.Items == nil {
		var ret []ResourceQuota
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceQuotaListAllOf) GetItemsOk() (*[]ResourceQuota, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ResourceQuotaListAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ResourceQuota and assigns it to the Items field.
func (o *ResourceQuotaListAllOf) SetItems(v []ResourceQuota) {
	o.Items = &v
}

func (o ResourceQuotaListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableResourceQuotaListAllOf struct {
	value *ResourceQuotaListAllOf
	isSet bool
}

func (v NullableResourceQuotaListAllOf) Get() *ResourceQuotaListAllOf {
	return v.value
}

func (v *NullableResourceQuotaListAllOf) Set(val *ResourceQuotaListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceQuotaListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceQuotaListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceQuotaListAllOf(val *ResourceQuotaListAllOf) *NullableResourceQuotaListAllOf {
	return &NullableResourceQuotaListAllOf{value: val, isSet: true}
}

func (v NullableResourceQuotaListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceQuotaListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
