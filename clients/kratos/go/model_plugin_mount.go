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

// PluginMount PluginMount plugin mount
type PluginMount struct {
	// description
	Description string `json:"Description"`
	// destination
	Destination string `json:"Destination"`
	// name
	Name string `json:"Name"`
	// options
	Options []string `json:"Options"`
	// settable
	Settable []string `json:"Settable"`
	// source
	Source string `json:"Source"`
	// type
	Type string `json:"Type"`
}

// NewPluginMount instantiates a new PluginMount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginMount(description string, destination string, name string, options []string, settable []string, source string, type_ string) *PluginMount {
	this := PluginMount{}
	this.Description = description
	this.Destination = destination
	this.Name = name
	this.Options = options
	this.Settable = settable
	this.Source = source
	this.Type = type_
	return &this
}

// NewPluginMountWithDefaults instantiates a new PluginMount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginMountWithDefaults() *PluginMount {
	this := PluginMount{}
	return &this
}

// GetDescription returns the Description field value
func (o *PluginMount) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PluginMount) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PluginMount) SetDescription(v string) {
	o.Description = v
}

// GetDestination returns the Destination field value
func (o *PluginMount) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *PluginMount) GetDestinationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *PluginMount) SetDestination(v string) {
	o.Destination = v
}

// GetName returns the Name field value
func (o *PluginMount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PluginMount) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PluginMount) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *PluginMount) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *PluginMount) GetOptionsOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *PluginMount) SetOptions(v []string) {
	o.Options = v
}

// GetSettable returns the Settable field value
func (o *PluginMount) GetSettable() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Settable
}

// GetSettableOk returns a tuple with the Settable field value
// and a boolean to check if the value has been set.
func (o *PluginMount) GetSettableOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Settable, true
}

// SetSettable sets field value
func (o *PluginMount) SetSettable(v []string) {
	o.Settable = v
}

// GetSource returns the Source field value
func (o *PluginMount) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *PluginMount) GetSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *PluginMount) SetSource(v string) {
	o.Source = v
}

// GetType returns the Type field value
func (o *PluginMount) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PluginMount) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PluginMount) SetType(v string) {
	o.Type = v
}

func (o PluginMount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Description"] = o.Description
	}
	if true {
		toSerialize["Destination"] = o.Destination
	}
	if true {
		toSerialize["Name"] = o.Name
	}
	if true {
		toSerialize["Options"] = o.Options
	}
	if true {
		toSerialize["Settable"] = o.Settable
	}
	if true {
		toSerialize["Source"] = o.Source
	}
	if true {
		toSerialize["Type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullablePluginMount struct {
	value *PluginMount
	isSet bool
}

func (v NullablePluginMount) Get() *PluginMount {
	return v.value
}

func (v *NullablePluginMount) Set(val *PluginMount) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginMount) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginMount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginMount(val *PluginMount) *NullablePluginMount {
	return &NullablePluginMount{value: val, isSet: true}
}

func (v NullablePluginMount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginMount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


