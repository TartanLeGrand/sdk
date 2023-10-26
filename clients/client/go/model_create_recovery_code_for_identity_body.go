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

// CreateRecoveryCodeForIdentityBody Create Recovery Code for Identity Request Body
type CreateRecoveryCodeForIdentityBody struct {
	// Code Expires In  The recovery code will expire after that amount of time has passed. Defaults to the configuration value of `selfservice.methods.code.config.lifespan`.
	ExpiresIn *string `json:"expires_in,omitempty"`
	// Identity to Recover  The identity's ID you wish to recover.
	IdentityId string `json:"identity_id"`
	AdditionalProperties map[string]interface{}
}

type _CreateRecoveryCodeForIdentityBody CreateRecoveryCodeForIdentityBody

// NewCreateRecoveryCodeForIdentityBody instantiates a new CreateRecoveryCodeForIdentityBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRecoveryCodeForIdentityBody(identityId string) *CreateRecoveryCodeForIdentityBody {
	this := CreateRecoveryCodeForIdentityBody{}
	this.IdentityId = identityId
	return &this
}

// NewCreateRecoveryCodeForIdentityBodyWithDefaults instantiates a new CreateRecoveryCodeForIdentityBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRecoveryCodeForIdentityBodyWithDefaults() *CreateRecoveryCodeForIdentityBody {
	this := CreateRecoveryCodeForIdentityBody{}
	return &this
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *CreateRecoveryCodeForIdentityBody) GetExpiresIn() string {
	if o == nil || o.ExpiresIn == nil {
		var ret string
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecoveryCodeForIdentityBody) GetExpiresInOk() (*string, bool) {
	if o == nil || o.ExpiresIn == nil {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *CreateRecoveryCodeForIdentityBody) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn != nil {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given string and assigns it to the ExpiresIn field.
func (o *CreateRecoveryCodeForIdentityBody) SetExpiresIn(v string) {
	o.ExpiresIn = &v
}

// GetIdentityId returns the IdentityId field value
func (o *CreateRecoveryCodeForIdentityBody) GetIdentityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value
// and a boolean to check if the value has been set.
func (o *CreateRecoveryCodeForIdentityBody) GetIdentityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityId, true
}

// SetIdentityId sets field value
func (o *CreateRecoveryCodeForIdentityBody) SetIdentityId(v string) {
	o.IdentityId = v
}

func (o CreateRecoveryCodeForIdentityBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiresIn != nil {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if true {
		toSerialize["identity_id"] = o.IdentityId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateRecoveryCodeForIdentityBody) UnmarshalJSON(bytes []byte) (err error) {
	varCreateRecoveryCodeForIdentityBody := _CreateRecoveryCodeForIdentityBody{}

	if err = json.Unmarshal(bytes, &varCreateRecoveryCodeForIdentityBody); err == nil {
		*o = CreateRecoveryCodeForIdentityBody(varCreateRecoveryCodeForIdentityBody)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "expires_in")
		delete(additionalProperties, "identity_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateRecoveryCodeForIdentityBody struct {
	value *CreateRecoveryCodeForIdentityBody
	isSet bool
}

func (v NullableCreateRecoveryCodeForIdentityBody) Get() *CreateRecoveryCodeForIdentityBody {
	return v.value
}

func (v *NullableCreateRecoveryCodeForIdentityBody) Set(val *CreateRecoveryCodeForIdentityBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRecoveryCodeForIdentityBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRecoveryCodeForIdentityBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRecoveryCodeForIdentityBody(val *CreateRecoveryCodeForIdentityBody) *NullableCreateRecoveryCodeForIdentityBody {
	return &NullableCreateRecoveryCodeForIdentityBody{value: val, isSet: true}
}

func (v NullableCreateRecoveryCodeForIdentityBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRecoveryCodeForIdentityBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


