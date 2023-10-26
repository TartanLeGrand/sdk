/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.2.14
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateRecoveryFlowWithCodeMethod Update Recovery Flow with Code Method
type UpdateRecoveryFlowWithCodeMethod struct {
	// Code from the recovery email  If you want to submit a code, use this field, but make sure to _not_ include the email field, as well.
	Code *string `json:"code,omitempty"`
	// Sending the anti-csrf token is only required for browser login flows.
	CsrfToken *string `json:"csrf_token,omitempty"`
	// The email address of the account to recover  If the email belongs to a valid account, a recovery email will be sent.  If you want to notify the email address if the account does not exist, see the [notify_unknown_recipients flag](https://www.ory.sh/docs/kratos/self-service/flows/account-recovery-password-reset#attempted-recovery-notifications)  If a code was already sent, including this field in the payload will invalidate the sent code and re-send a new code.  format: email
	Email *string `json:"email,omitempty"`
	// Method is the method that should be used for this recovery flow  Allowed values are `link` and `code`. link RecoveryStrategyLink code RecoveryStrategyCode
	Method string `json:"method"`
	AdditionalProperties map[string]interface{}
}

type _UpdateRecoveryFlowWithCodeMethod UpdateRecoveryFlowWithCodeMethod

// NewUpdateRecoveryFlowWithCodeMethod instantiates a new UpdateRecoveryFlowWithCodeMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRecoveryFlowWithCodeMethod(method string) *UpdateRecoveryFlowWithCodeMethod {
	this := UpdateRecoveryFlowWithCodeMethod{}
	this.Method = method
	return &this
}

// NewUpdateRecoveryFlowWithCodeMethodWithDefaults instantiates a new UpdateRecoveryFlowWithCodeMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRecoveryFlowWithCodeMethodWithDefaults() *UpdateRecoveryFlowWithCodeMethod {
	this := UpdateRecoveryFlowWithCodeMethod{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UpdateRecoveryFlowWithCodeMethod) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRecoveryFlowWithCodeMethod) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UpdateRecoveryFlowWithCodeMethod) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UpdateRecoveryFlowWithCodeMethod) SetCode(v string) {
	o.Code = &v
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *UpdateRecoveryFlowWithCodeMethod) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRecoveryFlowWithCodeMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *UpdateRecoveryFlowWithCodeMethod) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *UpdateRecoveryFlowWithCodeMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateRecoveryFlowWithCodeMethod) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRecoveryFlowWithCodeMethod) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateRecoveryFlowWithCodeMethod) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UpdateRecoveryFlowWithCodeMethod) SetEmail(v string) {
	o.Email = &v
}

// GetMethod returns the Method field value
func (o *UpdateRecoveryFlowWithCodeMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *UpdateRecoveryFlowWithCodeMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *UpdateRecoveryFlowWithCodeMethod) SetMethod(v string) {
	o.Method = v
}

func (o UpdateRecoveryFlowWithCodeMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["method"] = o.Method
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateRecoveryFlowWithCodeMethod) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateRecoveryFlowWithCodeMethod := _UpdateRecoveryFlowWithCodeMethod{}

	if err = json.Unmarshal(bytes, &varUpdateRecoveryFlowWithCodeMethod); err == nil {
		*o = UpdateRecoveryFlowWithCodeMethod(varUpdateRecoveryFlowWithCodeMethod)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "csrf_token")
		delete(additionalProperties, "email")
		delete(additionalProperties, "method")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateRecoveryFlowWithCodeMethod struct {
	value *UpdateRecoveryFlowWithCodeMethod
	isSet bool
}

func (v NullableUpdateRecoveryFlowWithCodeMethod) Get() *UpdateRecoveryFlowWithCodeMethod {
	return v.value
}

func (v *NullableUpdateRecoveryFlowWithCodeMethod) Set(val *UpdateRecoveryFlowWithCodeMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRecoveryFlowWithCodeMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRecoveryFlowWithCodeMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRecoveryFlowWithCodeMethod(val *UpdateRecoveryFlowWithCodeMethod) *NullableUpdateRecoveryFlowWithCodeMethod {
	return &NullableUpdateRecoveryFlowWithCodeMethod{value: val, isSet: true}
}

func (v NullableUpdateRecoveryFlowWithCodeMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRecoveryFlowWithCodeMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


