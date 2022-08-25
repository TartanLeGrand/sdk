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

// SubmitSelfServiceRegistrationFlowWithPasswordMethodBody SubmitSelfServiceRegistrationFlowWithPasswordMethodBody is used to decode the registration form payload when using the password method.
type SubmitSelfServiceRegistrationFlowWithPasswordMethodBody struct {
	// The CSRF Token
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method to use  This field must be set to `password` when using the password method.
	Method string `json:"method"`
	// Password to sign the user up with
	Password string `json:"password"`
	// The identity's traits
	Traits map[string]interface{} `json:"traits"`
}

// NewSubmitSelfServiceRegistrationFlowWithPasswordMethodBody instantiates a new SubmitSelfServiceRegistrationFlowWithPasswordMethodBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitSelfServiceRegistrationFlowWithPasswordMethodBody(method string, password string, traits map[string]interface{}) *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody {
	this := SubmitSelfServiceRegistrationFlowWithPasswordMethodBody{}
	this.Method = method
	this.Password = password
	this.Traits = traits
	return &this
}

// NewSubmitSelfServiceRegistrationFlowWithPasswordMethodBodyWithDefaults instantiates a new SubmitSelfServiceRegistrationFlowWithPasswordMethodBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitSelfServiceRegistrationFlowWithPasswordMethodBodyWithDefaults() *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody {
	this := SubmitSelfServiceRegistrationFlowWithPasswordMethodBody{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody) SetMethod(v string) {
	o.Method = v
}

// GetPassword returns the Password field value
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody) SetPassword(v string) {
	o.Password = v
}

// GetTraits returns the Traits field value
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody) GetTraits() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody) GetTraitsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Traits, true
}

// SetTraits sets field value
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody) SetTraits(v map[string]interface{}) {
	o.Traits = v
}

func (o SubmitSelfServiceRegistrationFlowWithPasswordMethodBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["traits"] = o.Traits
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitSelfServiceRegistrationFlowWithPasswordMethodBody struct {
	value *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody
	isSet bool
}

func (v NullableSubmitSelfServiceRegistrationFlowWithPasswordMethodBody) Get() *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody {
	return v.value
}

func (v *NullableSubmitSelfServiceRegistrationFlowWithPasswordMethodBody) Set(val *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceRegistrationFlowWithPasswordMethodBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceRegistrationFlowWithPasswordMethodBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceRegistrationFlowWithPasswordMethodBody(val *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody) *NullableSubmitSelfServiceRegistrationFlowWithPasswordMethodBody {
	return &NullableSubmitSelfServiceRegistrationFlowWithPasswordMethodBody{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceRegistrationFlowWithPasswordMethodBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceRegistrationFlowWithPasswordMethodBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


