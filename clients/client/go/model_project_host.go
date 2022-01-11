/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.44
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ProjectHost struct for ProjectHost
type ProjectHost struct {
	// The project's host.
	Host string `json:"host"`
	Id string `json:"id"`
	ProjectId string `json:"project_id"`
}

// NewProjectHost instantiates a new ProjectHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectHost(host string, id string, projectId string) *ProjectHost {
	this := ProjectHost{}
	this.Host = host
	this.Id = id
	this.ProjectId = projectId
	return &this
}

// NewProjectHostWithDefaults instantiates a new ProjectHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectHostWithDefaults() *ProjectHost {
	this := ProjectHost{}
	return &this
}

// GetHost returns the Host field value
func (o *ProjectHost) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *ProjectHost) GetHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *ProjectHost) SetHost(v string) {
	o.Host = v
}

// GetId returns the Id field value
func (o *ProjectHost) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectHost) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectHost) SetId(v string) {
	o.Id = v
}

// GetProjectId returns the ProjectId field value
func (o *ProjectHost) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ProjectHost) GetProjectIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ProjectHost) SetProjectId(v string) {
	o.ProjectId = v
}

func (o ProjectHost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["project_id"] = o.ProjectId
	}
	return json.Marshal(toSerialize)
}

type NullableProjectHost struct {
	value *ProjectHost
	isSet bool
}

func (v NullableProjectHost) Get() *ProjectHost {
	return v.value
}

func (v *NullableProjectHost) Set(val *ProjectHost) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectHost) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectHost(val *ProjectHost) *NullableProjectHost {
	return &NullableProjectHost{value: val, isSet: true}
}

func (v NullableProjectHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


