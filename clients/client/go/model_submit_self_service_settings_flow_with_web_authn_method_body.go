/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.111
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody struct for SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody
type SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody struct {
	// CSRFToken is the anti-CSRF token
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method  Should be set to \"webauthn\" when trying to add, update, or remove a webAuthn pairing.
	Method string `json:"method"`
	// Register a WebAuthn Security Key  It is expected that the JSON returned by the WebAuthn registration process is included here.
	WebauthnRegister *string `json:"webauthn_register,omitempty"`
	// Name of the WebAuthn Security Key to be Added  A human-readable name for the security key which will be added.
	WebauthnRegisterDisplayname *string `json:"webauthn_register_displayname,omitempty"`
	// Remove a WebAuthn Security Key  This must contain the ID of the WebAuthN connection.
	WebauthnRemove *string `json:"webauthn_remove,omitempty"`
}

// NewSubmitSelfServiceSettingsFlowWithWebAuthnMethodBody instantiates a new SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitSelfServiceSettingsFlowWithWebAuthnMethodBody(method string) *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody {
	this := SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody{}
	this.Method = method
	return &this
}

// NewSubmitSelfServiceSettingsFlowWithWebAuthnMethodBodyWithDefaults instantiates a new SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitSelfServiceSettingsFlowWithWebAuthnMethodBodyWithDefaults() *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody {
	this := SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) SetMethod(v string) {
	o.Method = v
}

// GetWebauthnRegister returns the WebauthnRegister field value if set, zero value otherwise.
func (o *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) GetWebauthnRegister() string {
	if o == nil || o.WebauthnRegister == nil {
		var ret string
		return ret
	}
	return *o.WebauthnRegister
}

// GetWebauthnRegisterOk returns a tuple with the WebauthnRegister field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) GetWebauthnRegisterOk() (*string, bool) {
	if o == nil || o.WebauthnRegister == nil {
		return nil, false
	}
	return o.WebauthnRegister, true
}

// HasWebauthnRegister returns a boolean if a field has been set.
func (o *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) HasWebauthnRegister() bool {
	if o != nil && o.WebauthnRegister != nil {
		return true
	}

	return false
}

// SetWebauthnRegister gets a reference to the given string and assigns it to the WebauthnRegister field.
func (o *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) SetWebauthnRegister(v string) {
	o.WebauthnRegister = &v
}

// GetWebauthnRegisterDisplayname returns the WebauthnRegisterDisplayname field value if set, zero value otherwise.
func (o *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) GetWebauthnRegisterDisplayname() string {
	if o == nil || o.WebauthnRegisterDisplayname == nil {
		var ret string
		return ret
	}
	return *o.WebauthnRegisterDisplayname
}

// GetWebauthnRegisterDisplaynameOk returns a tuple with the WebauthnRegisterDisplayname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) GetWebauthnRegisterDisplaynameOk() (*string, bool) {
	if o == nil || o.WebauthnRegisterDisplayname == nil {
		return nil, false
	}
	return o.WebauthnRegisterDisplayname, true
}

// HasWebauthnRegisterDisplayname returns a boolean if a field has been set.
func (o *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) HasWebauthnRegisterDisplayname() bool {
	if o != nil && o.WebauthnRegisterDisplayname != nil {
		return true
	}

	return false
}

// SetWebauthnRegisterDisplayname gets a reference to the given string and assigns it to the WebauthnRegisterDisplayname field.
func (o *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) SetWebauthnRegisterDisplayname(v string) {
	o.WebauthnRegisterDisplayname = &v
}

// GetWebauthnRemove returns the WebauthnRemove field value if set, zero value otherwise.
func (o *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) GetWebauthnRemove() string {
	if o == nil || o.WebauthnRemove == nil {
		var ret string
		return ret
	}
	return *o.WebauthnRemove
}

// GetWebauthnRemoveOk returns a tuple with the WebauthnRemove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) GetWebauthnRemoveOk() (*string, bool) {
	if o == nil || o.WebauthnRemove == nil {
		return nil, false
	}
	return o.WebauthnRemove, true
}

// HasWebauthnRemove returns a boolean if a field has been set.
func (o *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) HasWebauthnRemove() bool {
	if o != nil && o.WebauthnRemove != nil {
		return true
	}

	return false
}

// SetWebauthnRemove gets a reference to the given string and assigns it to the WebauthnRemove field.
func (o *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) SetWebauthnRemove(v string) {
	o.WebauthnRemove = &v
}

func (o SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if o.WebauthnRegister != nil {
		toSerialize["webauthn_register"] = o.WebauthnRegister
	}
	if o.WebauthnRegisterDisplayname != nil {
		toSerialize["webauthn_register_displayname"] = o.WebauthnRegisterDisplayname
	}
	if o.WebauthnRemove != nil {
		toSerialize["webauthn_remove"] = o.WebauthnRemove
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitSelfServiceSettingsFlowWithWebAuthnMethodBody struct {
	value *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody
	isSet bool
}

func (v NullableSubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) Get() *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody {
	return v.value
}

func (v *NullableSubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) Set(val *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceSettingsFlowWithWebAuthnMethodBody(val *SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) *NullableSubmitSelfServiceSettingsFlowWithWebAuthnMethodBody {
	return &NullableSubmitSelfServiceSettingsFlowWithWebAuthnMethodBody{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceSettingsFlowWithWebAuthnMethodBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


