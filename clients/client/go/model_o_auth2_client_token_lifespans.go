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

// OAuth2ClientTokenLifespans Lifespans of different token types issued for this OAuth 2.0 Client.
type OAuth2ClientTokenLifespans struct {
	AuthorizationCodeGrantAccessTokenLifespan NullableString `json:"authorization_code_grant_access_token_lifespan,omitempty"`
	AuthorizationCodeGrantIdTokenLifespan NullableString `json:"authorization_code_grant_id_token_lifespan,omitempty"`
	AuthorizationCodeGrantRefreshTokenLifespan NullableString `json:"authorization_code_grant_refresh_token_lifespan,omitempty"`
	ClientCredentialsGrantAccessTokenLifespan NullableString `json:"client_credentials_grant_access_token_lifespan,omitempty"`
	ImplicitGrantAccessTokenLifespan NullableString `json:"implicit_grant_access_token_lifespan,omitempty"`
	ImplicitGrantIdTokenLifespan NullableString `json:"implicit_grant_id_token_lifespan,omitempty"`
	JwtBearerGrantAccessTokenLifespan NullableString `json:"jwt_bearer_grant_access_token_lifespan,omitempty"`
	RefreshTokenGrantAccessTokenLifespan NullableString `json:"refresh_token_grant_access_token_lifespan,omitempty"`
	RefreshTokenGrantIdTokenLifespan NullableString `json:"refresh_token_grant_id_token_lifespan,omitempty"`
	RefreshTokenGrantRefreshTokenLifespan NullableString `json:"refresh_token_grant_refresh_token_lifespan,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ClientTokenLifespans OAuth2ClientTokenLifespans

// NewOAuth2ClientTokenLifespans instantiates a new OAuth2ClientTokenLifespans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientTokenLifespans() *OAuth2ClientTokenLifespans {
	this := OAuth2ClientTokenLifespans{}
	return &this
}

// NewOAuth2ClientTokenLifespansWithDefaults instantiates a new OAuth2ClientTokenLifespans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientTokenLifespansWithDefaults() *OAuth2ClientTokenLifespans {
	this := OAuth2ClientTokenLifespans{}
	return &this
}

// GetAuthorizationCodeGrantAccessTokenLifespan returns the AuthorizationCodeGrantAccessTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ClientTokenLifespans) GetAuthorizationCodeGrantAccessTokenLifespan() string {
	if o == nil || o.AuthorizationCodeGrantAccessTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationCodeGrantAccessTokenLifespan.Get()
}

// GetAuthorizationCodeGrantAccessTokenLifespanOk returns a tuple with the AuthorizationCodeGrantAccessTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ClientTokenLifespans) GetAuthorizationCodeGrantAccessTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizationCodeGrantAccessTokenLifespan.Get(), o.AuthorizationCodeGrantAccessTokenLifespan.IsSet()
}

// HasAuthorizationCodeGrantAccessTokenLifespan returns a boolean if a field has been set.
func (o *OAuth2ClientTokenLifespans) HasAuthorizationCodeGrantAccessTokenLifespan() bool {
	if o != nil && o.AuthorizationCodeGrantAccessTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetAuthorizationCodeGrantAccessTokenLifespan gets a reference to the given NullableString and assigns it to the AuthorizationCodeGrantAccessTokenLifespan field.
func (o *OAuth2ClientTokenLifespans) SetAuthorizationCodeGrantAccessTokenLifespan(v string) {
	o.AuthorizationCodeGrantAccessTokenLifespan.Set(&v)
}
// SetAuthorizationCodeGrantAccessTokenLifespanNil sets the value for AuthorizationCodeGrantAccessTokenLifespan to be an explicit nil
func (o *OAuth2ClientTokenLifespans) SetAuthorizationCodeGrantAccessTokenLifespanNil() {
	o.AuthorizationCodeGrantAccessTokenLifespan.Set(nil)
}

// UnsetAuthorizationCodeGrantAccessTokenLifespan ensures that no value is present for AuthorizationCodeGrantAccessTokenLifespan, not even an explicit nil
func (o *OAuth2ClientTokenLifespans) UnsetAuthorizationCodeGrantAccessTokenLifespan() {
	o.AuthorizationCodeGrantAccessTokenLifespan.Unset()
}

// GetAuthorizationCodeGrantIdTokenLifespan returns the AuthorizationCodeGrantIdTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ClientTokenLifespans) GetAuthorizationCodeGrantIdTokenLifespan() string {
	if o == nil || o.AuthorizationCodeGrantIdTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationCodeGrantIdTokenLifespan.Get()
}

// GetAuthorizationCodeGrantIdTokenLifespanOk returns a tuple with the AuthorizationCodeGrantIdTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ClientTokenLifespans) GetAuthorizationCodeGrantIdTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizationCodeGrantIdTokenLifespan.Get(), o.AuthorizationCodeGrantIdTokenLifespan.IsSet()
}

// HasAuthorizationCodeGrantIdTokenLifespan returns a boolean if a field has been set.
func (o *OAuth2ClientTokenLifespans) HasAuthorizationCodeGrantIdTokenLifespan() bool {
	if o != nil && o.AuthorizationCodeGrantIdTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetAuthorizationCodeGrantIdTokenLifespan gets a reference to the given NullableString and assigns it to the AuthorizationCodeGrantIdTokenLifespan field.
func (o *OAuth2ClientTokenLifespans) SetAuthorizationCodeGrantIdTokenLifespan(v string) {
	o.AuthorizationCodeGrantIdTokenLifespan.Set(&v)
}
// SetAuthorizationCodeGrantIdTokenLifespanNil sets the value for AuthorizationCodeGrantIdTokenLifespan to be an explicit nil
func (o *OAuth2ClientTokenLifespans) SetAuthorizationCodeGrantIdTokenLifespanNil() {
	o.AuthorizationCodeGrantIdTokenLifespan.Set(nil)
}

// UnsetAuthorizationCodeGrantIdTokenLifespan ensures that no value is present for AuthorizationCodeGrantIdTokenLifespan, not even an explicit nil
func (o *OAuth2ClientTokenLifespans) UnsetAuthorizationCodeGrantIdTokenLifespan() {
	o.AuthorizationCodeGrantIdTokenLifespan.Unset()
}

// GetAuthorizationCodeGrantRefreshTokenLifespan returns the AuthorizationCodeGrantRefreshTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ClientTokenLifespans) GetAuthorizationCodeGrantRefreshTokenLifespan() string {
	if o == nil || o.AuthorizationCodeGrantRefreshTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationCodeGrantRefreshTokenLifespan.Get()
}

// GetAuthorizationCodeGrantRefreshTokenLifespanOk returns a tuple with the AuthorizationCodeGrantRefreshTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ClientTokenLifespans) GetAuthorizationCodeGrantRefreshTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizationCodeGrantRefreshTokenLifespan.Get(), o.AuthorizationCodeGrantRefreshTokenLifespan.IsSet()
}

// HasAuthorizationCodeGrantRefreshTokenLifespan returns a boolean if a field has been set.
func (o *OAuth2ClientTokenLifespans) HasAuthorizationCodeGrantRefreshTokenLifespan() bool {
	if o != nil && o.AuthorizationCodeGrantRefreshTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetAuthorizationCodeGrantRefreshTokenLifespan gets a reference to the given NullableString and assigns it to the AuthorizationCodeGrantRefreshTokenLifespan field.
func (o *OAuth2ClientTokenLifespans) SetAuthorizationCodeGrantRefreshTokenLifespan(v string) {
	o.AuthorizationCodeGrantRefreshTokenLifespan.Set(&v)
}
// SetAuthorizationCodeGrantRefreshTokenLifespanNil sets the value for AuthorizationCodeGrantRefreshTokenLifespan to be an explicit nil
func (o *OAuth2ClientTokenLifespans) SetAuthorizationCodeGrantRefreshTokenLifespanNil() {
	o.AuthorizationCodeGrantRefreshTokenLifespan.Set(nil)
}

// UnsetAuthorizationCodeGrantRefreshTokenLifespan ensures that no value is present for AuthorizationCodeGrantRefreshTokenLifespan, not even an explicit nil
func (o *OAuth2ClientTokenLifespans) UnsetAuthorizationCodeGrantRefreshTokenLifespan() {
	o.AuthorizationCodeGrantRefreshTokenLifespan.Unset()
}

// GetClientCredentialsGrantAccessTokenLifespan returns the ClientCredentialsGrantAccessTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ClientTokenLifespans) GetClientCredentialsGrantAccessTokenLifespan() string {
	if o == nil || o.ClientCredentialsGrantAccessTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientCredentialsGrantAccessTokenLifespan.Get()
}

// GetClientCredentialsGrantAccessTokenLifespanOk returns a tuple with the ClientCredentialsGrantAccessTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ClientTokenLifespans) GetClientCredentialsGrantAccessTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientCredentialsGrantAccessTokenLifespan.Get(), o.ClientCredentialsGrantAccessTokenLifespan.IsSet()
}

// HasClientCredentialsGrantAccessTokenLifespan returns a boolean if a field has been set.
func (o *OAuth2ClientTokenLifespans) HasClientCredentialsGrantAccessTokenLifespan() bool {
	if o != nil && o.ClientCredentialsGrantAccessTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetClientCredentialsGrantAccessTokenLifespan gets a reference to the given NullableString and assigns it to the ClientCredentialsGrantAccessTokenLifespan field.
func (o *OAuth2ClientTokenLifespans) SetClientCredentialsGrantAccessTokenLifespan(v string) {
	o.ClientCredentialsGrantAccessTokenLifespan.Set(&v)
}
// SetClientCredentialsGrantAccessTokenLifespanNil sets the value for ClientCredentialsGrantAccessTokenLifespan to be an explicit nil
func (o *OAuth2ClientTokenLifespans) SetClientCredentialsGrantAccessTokenLifespanNil() {
	o.ClientCredentialsGrantAccessTokenLifespan.Set(nil)
}

// UnsetClientCredentialsGrantAccessTokenLifespan ensures that no value is present for ClientCredentialsGrantAccessTokenLifespan, not even an explicit nil
func (o *OAuth2ClientTokenLifespans) UnsetClientCredentialsGrantAccessTokenLifespan() {
	o.ClientCredentialsGrantAccessTokenLifespan.Unset()
}

// GetImplicitGrantAccessTokenLifespan returns the ImplicitGrantAccessTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ClientTokenLifespans) GetImplicitGrantAccessTokenLifespan() string {
	if o == nil || o.ImplicitGrantAccessTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.ImplicitGrantAccessTokenLifespan.Get()
}

// GetImplicitGrantAccessTokenLifespanOk returns a tuple with the ImplicitGrantAccessTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ClientTokenLifespans) GetImplicitGrantAccessTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImplicitGrantAccessTokenLifespan.Get(), o.ImplicitGrantAccessTokenLifespan.IsSet()
}

// HasImplicitGrantAccessTokenLifespan returns a boolean if a field has been set.
func (o *OAuth2ClientTokenLifespans) HasImplicitGrantAccessTokenLifespan() bool {
	if o != nil && o.ImplicitGrantAccessTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetImplicitGrantAccessTokenLifespan gets a reference to the given NullableString and assigns it to the ImplicitGrantAccessTokenLifespan field.
func (o *OAuth2ClientTokenLifespans) SetImplicitGrantAccessTokenLifespan(v string) {
	o.ImplicitGrantAccessTokenLifespan.Set(&v)
}
// SetImplicitGrantAccessTokenLifespanNil sets the value for ImplicitGrantAccessTokenLifespan to be an explicit nil
func (o *OAuth2ClientTokenLifespans) SetImplicitGrantAccessTokenLifespanNil() {
	o.ImplicitGrantAccessTokenLifespan.Set(nil)
}

// UnsetImplicitGrantAccessTokenLifespan ensures that no value is present for ImplicitGrantAccessTokenLifespan, not even an explicit nil
func (o *OAuth2ClientTokenLifespans) UnsetImplicitGrantAccessTokenLifespan() {
	o.ImplicitGrantAccessTokenLifespan.Unset()
}

// GetImplicitGrantIdTokenLifespan returns the ImplicitGrantIdTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ClientTokenLifespans) GetImplicitGrantIdTokenLifespan() string {
	if o == nil || o.ImplicitGrantIdTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.ImplicitGrantIdTokenLifespan.Get()
}

// GetImplicitGrantIdTokenLifespanOk returns a tuple with the ImplicitGrantIdTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ClientTokenLifespans) GetImplicitGrantIdTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImplicitGrantIdTokenLifespan.Get(), o.ImplicitGrantIdTokenLifespan.IsSet()
}

// HasImplicitGrantIdTokenLifespan returns a boolean if a field has been set.
func (o *OAuth2ClientTokenLifespans) HasImplicitGrantIdTokenLifespan() bool {
	if o != nil && o.ImplicitGrantIdTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetImplicitGrantIdTokenLifespan gets a reference to the given NullableString and assigns it to the ImplicitGrantIdTokenLifespan field.
func (o *OAuth2ClientTokenLifespans) SetImplicitGrantIdTokenLifespan(v string) {
	o.ImplicitGrantIdTokenLifespan.Set(&v)
}
// SetImplicitGrantIdTokenLifespanNil sets the value for ImplicitGrantIdTokenLifespan to be an explicit nil
func (o *OAuth2ClientTokenLifespans) SetImplicitGrantIdTokenLifespanNil() {
	o.ImplicitGrantIdTokenLifespan.Set(nil)
}

// UnsetImplicitGrantIdTokenLifespan ensures that no value is present for ImplicitGrantIdTokenLifespan, not even an explicit nil
func (o *OAuth2ClientTokenLifespans) UnsetImplicitGrantIdTokenLifespan() {
	o.ImplicitGrantIdTokenLifespan.Unset()
}

// GetJwtBearerGrantAccessTokenLifespan returns the JwtBearerGrantAccessTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ClientTokenLifespans) GetJwtBearerGrantAccessTokenLifespan() string {
	if o == nil || o.JwtBearerGrantAccessTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.JwtBearerGrantAccessTokenLifespan.Get()
}

// GetJwtBearerGrantAccessTokenLifespanOk returns a tuple with the JwtBearerGrantAccessTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ClientTokenLifespans) GetJwtBearerGrantAccessTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JwtBearerGrantAccessTokenLifespan.Get(), o.JwtBearerGrantAccessTokenLifespan.IsSet()
}

// HasJwtBearerGrantAccessTokenLifespan returns a boolean if a field has been set.
func (o *OAuth2ClientTokenLifespans) HasJwtBearerGrantAccessTokenLifespan() bool {
	if o != nil && o.JwtBearerGrantAccessTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetJwtBearerGrantAccessTokenLifespan gets a reference to the given NullableString and assigns it to the JwtBearerGrantAccessTokenLifespan field.
func (o *OAuth2ClientTokenLifespans) SetJwtBearerGrantAccessTokenLifespan(v string) {
	o.JwtBearerGrantAccessTokenLifespan.Set(&v)
}
// SetJwtBearerGrantAccessTokenLifespanNil sets the value for JwtBearerGrantAccessTokenLifespan to be an explicit nil
func (o *OAuth2ClientTokenLifespans) SetJwtBearerGrantAccessTokenLifespanNil() {
	o.JwtBearerGrantAccessTokenLifespan.Set(nil)
}

// UnsetJwtBearerGrantAccessTokenLifespan ensures that no value is present for JwtBearerGrantAccessTokenLifespan, not even an explicit nil
func (o *OAuth2ClientTokenLifespans) UnsetJwtBearerGrantAccessTokenLifespan() {
	o.JwtBearerGrantAccessTokenLifespan.Unset()
}

// GetRefreshTokenGrantAccessTokenLifespan returns the RefreshTokenGrantAccessTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ClientTokenLifespans) GetRefreshTokenGrantAccessTokenLifespan() string {
	if o == nil || o.RefreshTokenGrantAccessTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.RefreshTokenGrantAccessTokenLifespan.Get()
}

// GetRefreshTokenGrantAccessTokenLifespanOk returns a tuple with the RefreshTokenGrantAccessTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ClientTokenLifespans) GetRefreshTokenGrantAccessTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshTokenGrantAccessTokenLifespan.Get(), o.RefreshTokenGrantAccessTokenLifespan.IsSet()
}

// HasRefreshTokenGrantAccessTokenLifespan returns a boolean if a field has been set.
func (o *OAuth2ClientTokenLifespans) HasRefreshTokenGrantAccessTokenLifespan() bool {
	if o != nil && o.RefreshTokenGrantAccessTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetRefreshTokenGrantAccessTokenLifespan gets a reference to the given NullableString and assigns it to the RefreshTokenGrantAccessTokenLifespan field.
func (o *OAuth2ClientTokenLifespans) SetRefreshTokenGrantAccessTokenLifespan(v string) {
	o.RefreshTokenGrantAccessTokenLifespan.Set(&v)
}
// SetRefreshTokenGrantAccessTokenLifespanNil sets the value for RefreshTokenGrantAccessTokenLifespan to be an explicit nil
func (o *OAuth2ClientTokenLifespans) SetRefreshTokenGrantAccessTokenLifespanNil() {
	o.RefreshTokenGrantAccessTokenLifespan.Set(nil)
}

// UnsetRefreshTokenGrantAccessTokenLifespan ensures that no value is present for RefreshTokenGrantAccessTokenLifespan, not even an explicit nil
func (o *OAuth2ClientTokenLifespans) UnsetRefreshTokenGrantAccessTokenLifespan() {
	o.RefreshTokenGrantAccessTokenLifespan.Unset()
}

// GetRefreshTokenGrantIdTokenLifespan returns the RefreshTokenGrantIdTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ClientTokenLifespans) GetRefreshTokenGrantIdTokenLifespan() string {
	if o == nil || o.RefreshTokenGrantIdTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.RefreshTokenGrantIdTokenLifespan.Get()
}

// GetRefreshTokenGrantIdTokenLifespanOk returns a tuple with the RefreshTokenGrantIdTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ClientTokenLifespans) GetRefreshTokenGrantIdTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshTokenGrantIdTokenLifespan.Get(), o.RefreshTokenGrantIdTokenLifespan.IsSet()
}

// HasRefreshTokenGrantIdTokenLifespan returns a boolean if a field has been set.
func (o *OAuth2ClientTokenLifespans) HasRefreshTokenGrantIdTokenLifespan() bool {
	if o != nil && o.RefreshTokenGrantIdTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetRefreshTokenGrantIdTokenLifespan gets a reference to the given NullableString and assigns it to the RefreshTokenGrantIdTokenLifespan field.
func (o *OAuth2ClientTokenLifespans) SetRefreshTokenGrantIdTokenLifespan(v string) {
	o.RefreshTokenGrantIdTokenLifespan.Set(&v)
}
// SetRefreshTokenGrantIdTokenLifespanNil sets the value for RefreshTokenGrantIdTokenLifespan to be an explicit nil
func (o *OAuth2ClientTokenLifespans) SetRefreshTokenGrantIdTokenLifespanNil() {
	o.RefreshTokenGrantIdTokenLifespan.Set(nil)
}

// UnsetRefreshTokenGrantIdTokenLifespan ensures that no value is present for RefreshTokenGrantIdTokenLifespan, not even an explicit nil
func (o *OAuth2ClientTokenLifespans) UnsetRefreshTokenGrantIdTokenLifespan() {
	o.RefreshTokenGrantIdTokenLifespan.Unset()
}

// GetRefreshTokenGrantRefreshTokenLifespan returns the RefreshTokenGrantRefreshTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ClientTokenLifespans) GetRefreshTokenGrantRefreshTokenLifespan() string {
	if o == nil || o.RefreshTokenGrantRefreshTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.RefreshTokenGrantRefreshTokenLifespan.Get()
}

// GetRefreshTokenGrantRefreshTokenLifespanOk returns a tuple with the RefreshTokenGrantRefreshTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ClientTokenLifespans) GetRefreshTokenGrantRefreshTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshTokenGrantRefreshTokenLifespan.Get(), o.RefreshTokenGrantRefreshTokenLifespan.IsSet()
}

// HasRefreshTokenGrantRefreshTokenLifespan returns a boolean if a field has been set.
func (o *OAuth2ClientTokenLifespans) HasRefreshTokenGrantRefreshTokenLifespan() bool {
	if o != nil && o.RefreshTokenGrantRefreshTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetRefreshTokenGrantRefreshTokenLifespan gets a reference to the given NullableString and assigns it to the RefreshTokenGrantRefreshTokenLifespan field.
func (o *OAuth2ClientTokenLifespans) SetRefreshTokenGrantRefreshTokenLifespan(v string) {
	o.RefreshTokenGrantRefreshTokenLifespan.Set(&v)
}
// SetRefreshTokenGrantRefreshTokenLifespanNil sets the value for RefreshTokenGrantRefreshTokenLifespan to be an explicit nil
func (o *OAuth2ClientTokenLifespans) SetRefreshTokenGrantRefreshTokenLifespanNil() {
	o.RefreshTokenGrantRefreshTokenLifespan.Set(nil)
}

// UnsetRefreshTokenGrantRefreshTokenLifespan ensures that no value is present for RefreshTokenGrantRefreshTokenLifespan, not even an explicit nil
func (o *OAuth2ClientTokenLifespans) UnsetRefreshTokenGrantRefreshTokenLifespan() {
	o.RefreshTokenGrantRefreshTokenLifespan.Unset()
}

func (o OAuth2ClientTokenLifespans) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthorizationCodeGrantAccessTokenLifespan.IsSet() {
		toSerialize["authorization_code_grant_access_token_lifespan"] = o.AuthorizationCodeGrantAccessTokenLifespan.Get()
	}
	if o.AuthorizationCodeGrantIdTokenLifespan.IsSet() {
		toSerialize["authorization_code_grant_id_token_lifespan"] = o.AuthorizationCodeGrantIdTokenLifespan.Get()
	}
	if o.AuthorizationCodeGrantRefreshTokenLifespan.IsSet() {
		toSerialize["authorization_code_grant_refresh_token_lifespan"] = o.AuthorizationCodeGrantRefreshTokenLifespan.Get()
	}
	if o.ClientCredentialsGrantAccessTokenLifespan.IsSet() {
		toSerialize["client_credentials_grant_access_token_lifespan"] = o.ClientCredentialsGrantAccessTokenLifespan.Get()
	}
	if o.ImplicitGrantAccessTokenLifespan.IsSet() {
		toSerialize["implicit_grant_access_token_lifespan"] = o.ImplicitGrantAccessTokenLifespan.Get()
	}
	if o.ImplicitGrantIdTokenLifespan.IsSet() {
		toSerialize["implicit_grant_id_token_lifespan"] = o.ImplicitGrantIdTokenLifespan.Get()
	}
	if o.JwtBearerGrantAccessTokenLifespan.IsSet() {
		toSerialize["jwt_bearer_grant_access_token_lifespan"] = o.JwtBearerGrantAccessTokenLifespan.Get()
	}
	if o.RefreshTokenGrantAccessTokenLifespan.IsSet() {
		toSerialize["refresh_token_grant_access_token_lifespan"] = o.RefreshTokenGrantAccessTokenLifespan.Get()
	}
	if o.RefreshTokenGrantIdTokenLifespan.IsSet() {
		toSerialize["refresh_token_grant_id_token_lifespan"] = o.RefreshTokenGrantIdTokenLifespan.Get()
	}
	if o.RefreshTokenGrantRefreshTokenLifespan.IsSet() {
		toSerialize["refresh_token_grant_refresh_token_lifespan"] = o.RefreshTokenGrantRefreshTokenLifespan.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuth2ClientTokenLifespans) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2ClientTokenLifespans := _OAuth2ClientTokenLifespans{}

	if err = json.Unmarshal(bytes, &varOAuth2ClientTokenLifespans); err == nil {
		*o = OAuth2ClientTokenLifespans(varOAuth2ClientTokenLifespans)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "authorization_code_grant_access_token_lifespan")
		delete(additionalProperties, "authorization_code_grant_id_token_lifespan")
		delete(additionalProperties, "authorization_code_grant_refresh_token_lifespan")
		delete(additionalProperties, "client_credentials_grant_access_token_lifespan")
		delete(additionalProperties, "implicit_grant_access_token_lifespan")
		delete(additionalProperties, "implicit_grant_id_token_lifespan")
		delete(additionalProperties, "jwt_bearer_grant_access_token_lifespan")
		delete(additionalProperties, "refresh_token_grant_access_token_lifespan")
		delete(additionalProperties, "refresh_token_grant_id_token_lifespan")
		delete(additionalProperties, "refresh_token_grant_refresh_token_lifespan")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2ClientTokenLifespans struct {
	value *OAuth2ClientTokenLifespans
	isSet bool
}

func (v NullableOAuth2ClientTokenLifespans) Get() *OAuth2ClientTokenLifespans {
	return v.value
}

func (v *NullableOAuth2ClientTokenLifespans) Set(val *OAuth2ClientTokenLifespans) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientTokenLifespans) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientTokenLifespans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientTokenLifespans(val *OAuth2ClientTokenLifespans) *NullableOAuth2ClientTokenLifespans {
	return &NullableOAuth2ClientTokenLifespans{value: val, isSet: true}
}

func (v NullableOAuth2ClientTokenLifespans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientTokenLifespans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


