/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.2.5
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateProjectMemberInviteBody Create Project MemberInvite Request Body
type CreateProjectMemberInviteBody struct {
	// A email to invite
	InviteeEmail *string `json:"invitee_email,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateProjectMemberInviteBody CreateProjectMemberInviteBody

// NewCreateProjectMemberInviteBody instantiates a new CreateProjectMemberInviteBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectMemberInviteBody() *CreateProjectMemberInviteBody {
	this := CreateProjectMemberInviteBody{}
	return &this
}

// NewCreateProjectMemberInviteBodyWithDefaults instantiates a new CreateProjectMemberInviteBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectMemberInviteBodyWithDefaults() *CreateProjectMemberInviteBody {
	this := CreateProjectMemberInviteBody{}
	return &this
}

// GetInviteeEmail returns the InviteeEmail field value if set, zero value otherwise.
func (o *CreateProjectMemberInviteBody) GetInviteeEmail() string {
	if o == nil || o.InviteeEmail == nil {
		var ret string
		return ret
	}
	return *o.InviteeEmail
}

// GetInviteeEmailOk returns a tuple with the InviteeEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectMemberInviteBody) GetInviteeEmailOk() (*string, bool) {
	if o == nil || o.InviteeEmail == nil {
		return nil, false
	}
	return o.InviteeEmail, true
}

// HasInviteeEmail returns a boolean if a field has been set.
func (o *CreateProjectMemberInviteBody) HasInviteeEmail() bool {
	if o != nil && o.InviteeEmail != nil {
		return true
	}

	return false
}

// SetInviteeEmail gets a reference to the given string and assigns it to the InviteeEmail field.
func (o *CreateProjectMemberInviteBody) SetInviteeEmail(v string) {
	o.InviteeEmail = &v
}

func (o CreateProjectMemberInviteBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InviteeEmail != nil {
		toSerialize["invitee_email"] = o.InviteeEmail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateProjectMemberInviteBody) UnmarshalJSON(bytes []byte) (err error) {
	varCreateProjectMemberInviteBody := _CreateProjectMemberInviteBody{}

	if err = json.Unmarshal(bytes, &varCreateProjectMemberInviteBody); err == nil {
		*o = CreateProjectMemberInviteBody(varCreateProjectMemberInviteBody)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "invitee_email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateProjectMemberInviteBody struct {
	value *CreateProjectMemberInviteBody
	isSet bool
}

func (v NullableCreateProjectMemberInviteBody) Get() *CreateProjectMemberInviteBody {
	return v.value
}

func (v *NullableCreateProjectMemberInviteBody) Set(val *CreateProjectMemberInviteBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectMemberInviteBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectMemberInviteBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectMemberInviteBody(val *CreateProjectMemberInviteBody) *NullableCreateProjectMemberInviteBody {
	return &NullableCreateProjectMemberInviteBody{value: val, isSet: true}
}

func (v NullableCreateProjectMemberInviteBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectMemberInviteBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


