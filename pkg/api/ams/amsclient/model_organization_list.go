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

// OrganizationList struct for OrganizationList
type OrganizationList struct {
	Kind  string         `json:"kind"`
	Page  int32          `json:"page"`
	Size  int32          `json:"size"`
	Total int32          `json:"total"`
	Items []Organization `json:"items"`
}

// NewOrganizationList instantiates a new OrganizationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationList(kind string, page int32, size int32, total int32, items []Organization) *OrganizationList {
	this := OrganizationList{}
	this.Kind = kind
	this.Page = page
	this.Size = size
	this.Total = total
	this.Items = items
	return &this
}

// NewOrganizationListWithDefaults instantiates a new OrganizationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationListWithDefaults() *OrganizationList {
	this := OrganizationList{}
	return &this
}

// GetKind returns the Kind field value
func (o *OrganizationList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *OrganizationList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *OrganizationList) SetKind(v string) {
	o.Kind = v
}

// GetPage returns the Page field value
func (o *OrganizationList) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *OrganizationList) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *OrganizationList) SetPage(v int32) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *OrganizationList) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *OrganizationList) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *OrganizationList) SetSize(v int32) {
	o.Size = v
}

// GetTotal returns the Total field value
func (o *OrganizationList) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *OrganizationList) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *OrganizationList) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value
func (o *OrganizationList) GetItems() []Organization {
	if o == nil {
		var ret []Organization
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *OrganizationList) GetItemsOk() (*[]Organization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *OrganizationList) SetItems(v []Organization) {
	o.Items = v
}

func (o OrganizationList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationList struct {
	value *OrganizationList
	isSet bool
}

func (v NullableOrganizationList) Get() *OrganizationList {
	return v.value
}

func (v *NullableOrganizationList) Set(val *OrganizationList) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationList) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationList(val *OrganizationList) *NullableOrganizationList {
	return &NullableOrganizationList{value: val, isSet: true}
}

func (v NullableOrganizationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
