/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.18
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// HandledOAuth2LogoutRequest struct for HandledOAuth2LogoutRequest
type HandledOAuth2LogoutRequest struct {
	// Original request URL to which you should redirect the user if request was already handled.
	RedirectTo string `json:"redirect_to"`
}

// NewHandledOAuth2LogoutRequest instantiates a new HandledOAuth2LogoutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandledOAuth2LogoutRequest(redirectTo string) *HandledOAuth2LogoutRequest {
	this := HandledOAuth2LogoutRequest{}
	this.RedirectTo = redirectTo
	return &this
}

// NewHandledOAuth2LogoutRequestWithDefaults instantiates a new HandledOAuth2LogoutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandledOAuth2LogoutRequestWithDefaults() *HandledOAuth2LogoutRequest {
	this := HandledOAuth2LogoutRequest{}
	return &this
}

// GetRedirectTo returns the RedirectTo field value
func (o *HandledOAuth2LogoutRequest) GetRedirectTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectTo
}

// GetRedirectToOk returns a tuple with the RedirectTo field value
// and a boolean to check if the value has been set.
func (o *HandledOAuth2LogoutRequest) GetRedirectToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectTo, true
}

// SetRedirectTo sets field value
func (o *HandledOAuth2LogoutRequest) SetRedirectTo(v string) {
	o.RedirectTo = v
}

func (o HandledOAuth2LogoutRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["redirect_to"] = o.RedirectTo
	}
	return json.Marshal(toSerialize)
}

type NullableHandledOAuth2LogoutRequest struct {
	value *HandledOAuth2LogoutRequest
	isSet bool
}

func (v NullableHandledOAuth2LogoutRequest) Get() *HandledOAuth2LogoutRequest {
	return v.value
}

func (v *NullableHandledOAuth2LogoutRequest) Set(val *HandledOAuth2LogoutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHandledOAuth2LogoutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHandledOAuth2LogoutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandledOAuth2LogoutRequest(val *HandledOAuth2LogoutRequest) *NullableHandledOAuth2LogoutRequest {
	return &NullableHandledOAuth2LogoutRequest{value: val, isSet: true}
}

func (v NullableHandledOAuth2LogoutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandledOAuth2LogoutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


