/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.7.3-alpha.2
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ServiceUpdateResponse ServiceUpdateResponse service update response
type ServiceUpdateResponse struct {
	// Optional warning messages
	Warnings []string `json:"Warnings,omitempty"`
}

// NewServiceUpdateResponse instantiates a new ServiceUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceUpdateResponse() *ServiceUpdateResponse {
	this := ServiceUpdateResponse{}
	return &this
}

// NewServiceUpdateResponseWithDefaults instantiates a new ServiceUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceUpdateResponseWithDefaults() *ServiceUpdateResponse {
	this := ServiceUpdateResponse{}
	return &this
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ServiceUpdateResponse) GetWarnings() []string {
	if o == nil || o.Warnings == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUpdateResponse) GetWarningsOk() ([]string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ServiceUpdateResponse) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *ServiceUpdateResponse) SetWarnings(v []string) {
	o.Warnings = v
}

func (o ServiceUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Warnings != nil {
		toSerialize["Warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableServiceUpdateResponse struct {
	value *ServiceUpdateResponse
	isSet bool
}

func (v NullableServiceUpdateResponse) Get() *ServiceUpdateResponse {
	return v.value
}

func (v *NullableServiceUpdateResponse) Set(val *ServiceUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceUpdateResponse(val *ServiceUpdateResponse) *NullableServiceUpdateResponse {
	return &NullableServiceUpdateResponse{value: val, isSet: true}
}

func (v NullableServiceUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


