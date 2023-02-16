/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.11
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type OidcApi interface {

	/*
	CreateOidcDynamicClient Register OAuth2 Client using OpenID Dynamic Client Registration

	This endpoint behaves like the administrative counterpart (`createOAuth2Client`) but is capable of facing the
public internet directly and can be used in self-service. It implements the OpenID Connect
Dynamic Client Registration Protocol. This feature needs to be enabled in the configuration. This endpoint
is disabled by default. It can be enabled by an administrator.

Please note that using this endpoint you are not able to choose the `client_secret` nor the `client_id` as those
values will be server generated when specifying `token_endpoint_auth_method` as `client_secret_basic` or
`client_secret_post`.

The `client_secret` will be returned in the response and you will not be able to retrieve it later on.
Write the secret down and keep it somewhere safe.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OidcApiCreateOidcDynamicClientRequest
	*/
	CreateOidcDynamicClient(ctx context.Context) OidcApiCreateOidcDynamicClientRequest

	// CreateOidcDynamicClientExecute executes the request
	//  @return OAuth2Client
	CreateOidcDynamicClientExecute(r OidcApiCreateOidcDynamicClientRequest) (*OAuth2Client, *http.Response, error)

	/*
	DeleteOidcDynamicClient Delete OAuth 2.0 Client using the OpenID Dynamic Client Registration Management Protocol

	This endpoint behaves like the administrative counterpart (`deleteOAuth2Client`) but is capable of facing the
public internet directly and can be used in self-service. It implements the OpenID Connect
Dynamic Client Registration Protocol. This feature needs to be enabled in the configuration. This endpoint
is disabled by default. It can be enabled by an administrator.

To use this endpoint, you will need to present the client's authentication credentials. If the OAuth2 Client
uses the Token Endpoint Authentication Method `client_secret_post`, you need to present the client secret in the URL query.
If it uses `client_secret_basic`, present the Client ID and the Client Secret in the Authorization header.

OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are
generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The id of the OAuth 2.0 Client.
	@return OidcApiDeleteOidcDynamicClientRequest
	*/
	DeleteOidcDynamicClient(ctx context.Context, id string) OidcApiDeleteOidcDynamicClientRequest

	// DeleteOidcDynamicClientExecute executes the request
	DeleteOidcDynamicClientExecute(r OidcApiDeleteOidcDynamicClientRequest) (*http.Response, error)

	/*
	DiscoverOidcConfiguration OpenID Connect Discovery

	A mechanism for an OpenID Connect Relying Party to discover the End-User's OpenID Provider and obtain information needed to interact with it, including its OAuth 2.0 endpoint locations.

Popular libraries for OpenID Connect clients include oidc-client-js (JavaScript), go-oidc (Golang), and others.
For a full list of clients go here: https://openid.net/developers/certified/

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OidcApiDiscoverOidcConfigurationRequest
	*/
	DiscoverOidcConfiguration(ctx context.Context) OidcApiDiscoverOidcConfigurationRequest

	// DiscoverOidcConfigurationExecute executes the request
	//  @return OidcConfiguration
	DiscoverOidcConfigurationExecute(r OidcApiDiscoverOidcConfigurationRequest) (*OidcConfiguration, *http.Response, error)

	/*
	GetOidcDynamicClient Get OAuth2 Client using OpenID Dynamic Client Registration

	This endpoint behaves like the administrative counterpart (`getOAuth2Client`) but is capable of facing the
public internet directly and can be used in self-service. It implements the OpenID Connect
Dynamic Client Registration Protocol.

To use this endpoint, you will need to present the client's authentication credentials. If the OAuth2 Client
uses the Token Endpoint Authentication Method `client_secret_post`, you need to present the client secret in the URL query.
If it uses `client_secret_basic`, present the Client ID and the Client Secret in the Authorization header.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The id of the OAuth 2.0 Client.
	@return OidcApiGetOidcDynamicClientRequest
	*/
	GetOidcDynamicClient(ctx context.Context, id string) OidcApiGetOidcDynamicClientRequest

	// GetOidcDynamicClientExecute executes the request
	//  @return OAuth2Client
	GetOidcDynamicClientExecute(r OidcApiGetOidcDynamicClientRequest) (*OAuth2Client, *http.Response, error)

	/*
	GetOidcUserInfo OpenID Connect Userinfo

	This endpoint returns the payload of the ID Token, including `session.id_token` values, of
the provided OAuth 2.0 Access Token's consent request.

In the case of authentication error, a WWW-Authenticate header might be set in the response
with more information about the error. See [the spec](https://datatracker.ietf.org/doc/html/rfc6750#section-3)
for more details about header format.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OidcApiGetOidcUserInfoRequest
	*/
	GetOidcUserInfo(ctx context.Context) OidcApiGetOidcUserInfoRequest

	// GetOidcUserInfoExecute executes the request
	//  @return OidcUserInfo
	GetOidcUserInfoExecute(r OidcApiGetOidcUserInfoRequest) (*OidcUserInfo, *http.Response, error)

	/*
	RevokeOidcSession OpenID Connect Front- and Back-channel Enabled Logout

	This endpoint initiates and completes user logout at the Ory OAuth2 & OpenID provider and initiates OpenID Connect Front- / Back-channel logout:

https://openid.net/specs/openid-connect-frontchannel-1_0.html
https://openid.net/specs/openid-connect-backchannel-1_0.html

Back-channel logout is performed asynchronously and does not affect logout flow.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OidcApiRevokeOidcSessionRequest
	*/
	RevokeOidcSession(ctx context.Context) OidcApiRevokeOidcSessionRequest

	// RevokeOidcSessionExecute executes the request
	RevokeOidcSessionExecute(r OidcApiRevokeOidcSessionRequest) (*http.Response, error)

	/*
	SetOidcDynamicClient Set OAuth2 Client using OpenID Dynamic Client Registration

	This endpoint behaves like the administrative counterpart (`setOAuth2Client`) but is capable of facing the
public internet directly to be used by third parties. It implements the OpenID Connect
Dynamic Client Registration Protocol.

This feature is disabled per default. It can be enabled by a system administrator.

If you pass `client_secret` the secret is used, otherwise the existing secret is used. If set, the secret is echoed in the response.
It is not possible to retrieve it later on.

To use this endpoint, you will need to present the client's authentication credentials. If the OAuth2 Client
uses the Token Endpoint Authentication Method `client_secret_post`, you need to present the client secret in the URL query.
If it uses `client_secret_basic`, present the Client ID and the Client Secret in the Authorization header.

OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are
generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id OAuth 2.0 Client ID
	@return OidcApiSetOidcDynamicClientRequest
	*/
	SetOidcDynamicClient(ctx context.Context, id string) OidcApiSetOidcDynamicClientRequest

	// SetOidcDynamicClientExecute executes the request
	//  @return OAuth2Client
	SetOidcDynamicClientExecute(r OidcApiSetOidcDynamicClientRequest) (*OAuth2Client, *http.Response, error)
}

// OidcApiService OidcApi service
type OidcApiService service

type OidcApiCreateOidcDynamicClientRequest struct {
	ctx context.Context
	ApiService OidcApi
	oAuth2Client *OAuth2Client
}

// Dynamic Client Registration Request Body
func (r OidcApiCreateOidcDynamicClientRequest) OAuth2Client(oAuth2Client OAuth2Client) OidcApiCreateOidcDynamicClientRequest {
	r.oAuth2Client = &oAuth2Client
	return r
}

func (r OidcApiCreateOidcDynamicClientRequest) Execute() (*OAuth2Client, *http.Response, error) {
	return r.ApiService.CreateOidcDynamicClientExecute(r)
}

/*
CreateOidcDynamicClient Register OAuth2 Client using OpenID Dynamic Client Registration

This endpoint behaves like the administrative counterpart (`createOAuth2Client`) but is capable of facing the
public internet directly and can be used in self-service. It implements the OpenID Connect
Dynamic Client Registration Protocol. This feature needs to be enabled in the configuration. This endpoint
is disabled by default. It can be enabled by an administrator.

Please note that using this endpoint you are not able to choose the `client_secret` nor the `client_id` as those
values will be server generated when specifying `token_endpoint_auth_method` as `client_secret_basic` or
`client_secret_post`.

The `client_secret` will be returned in the response and you will not be able to retrieve it later on.
Write the secret down and keep it somewhere safe.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OidcApiCreateOidcDynamicClientRequest
*/
func (a *OidcApiService) CreateOidcDynamicClient(ctx context.Context) OidcApiCreateOidcDynamicClientRequest {
	return OidcApiCreateOidcDynamicClientRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OAuth2Client
func (a *OidcApiService) CreateOidcDynamicClientExecute(r OidcApiCreateOidcDynamicClientRequest) (*OAuth2Client, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OAuth2Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OidcApiService.CreateOidcDynamicClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/register"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oAuth2Client == nil {
		return localVarReturnValue, nil, reportError("oAuth2Client is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.oAuth2Client
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorOAuth2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorOAuth2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OidcApiDeleteOidcDynamicClientRequest struct {
	ctx context.Context
	ApiService OidcApi
	id string
}

func (r OidcApiDeleteOidcDynamicClientRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOidcDynamicClientExecute(r)
}

/*
DeleteOidcDynamicClient Delete OAuth 2.0 Client using the OpenID Dynamic Client Registration Management Protocol

This endpoint behaves like the administrative counterpart (`deleteOAuth2Client`) but is capable of facing the
public internet directly and can be used in self-service. It implements the OpenID Connect
Dynamic Client Registration Protocol. This feature needs to be enabled in the configuration. This endpoint
is disabled by default. It can be enabled by an administrator.

To use this endpoint, you will need to present the client's authentication credentials. If the OAuth2 Client
uses the Token Endpoint Authentication Method `client_secret_post`, you need to present the client secret in the URL query.
If it uses `client_secret_basic`, present the Client ID and the Client Secret in the Authorization header.

OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are
generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The id of the OAuth 2.0 Client.
 @return OidcApiDeleteOidcDynamicClientRequest
*/
func (a *OidcApiService) DeleteOidcDynamicClient(ctx context.Context, id string) OidcApiDeleteOidcDynamicClientRequest {
	return OidcApiDeleteOidcDynamicClientRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *OidcApiService) DeleteOidcDynamicClientExecute(r OidcApiDeleteOidcDynamicClientRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OidcApiService.DeleteOidcDynamicClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/register/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type OidcApiDiscoverOidcConfigurationRequest struct {
	ctx context.Context
	ApiService OidcApi
}

func (r OidcApiDiscoverOidcConfigurationRequest) Execute() (*OidcConfiguration, *http.Response, error) {
	return r.ApiService.DiscoverOidcConfigurationExecute(r)
}

/*
DiscoverOidcConfiguration OpenID Connect Discovery

A mechanism for an OpenID Connect Relying Party to discover the End-User's OpenID Provider and obtain information needed to interact with it, including its OAuth 2.0 endpoint locations.

Popular libraries for OpenID Connect clients include oidc-client-js (JavaScript), go-oidc (Golang), and others.
For a full list of clients go here: https://openid.net/developers/certified/

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OidcApiDiscoverOidcConfigurationRequest
*/
func (a *OidcApiService) DiscoverOidcConfiguration(ctx context.Context) OidcApiDiscoverOidcConfigurationRequest {
	return OidcApiDiscoverOidcConfigurationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OidcConfiguration
func (a *OidcApiService) DiscoverOidcConfigurationExecute(r OidcApiDiscoverOidcConfigurationRequest) (*OidcConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OidcConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OidcApiService.DiscoverOidcConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/.well-known/openid-configuration"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorOAuth2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OidcApiGetOidcDynamicClientRequest struct {
	ctx context.Context
	ApiService OidcApi
	id string
}

func (r OidcApiGetOidcDynamicClientRequest) Execute() (*OAuth2Client, *http.Response, error) {
	return r.ApiService.GetOidcDynamicClientExecute(r)
}

/*
GetOidcDynamicClient Get OAuth2 Client using OpenID Dynamic Client Registration

This endpoint behaves like the administrative counterpart (`getOAuth2Client`) but is capable of facing the
public internet directly and can be used in self-service. It implements the OpenID Connect
Dynamic Client Registration Protocol.

To use this endpoint, you will need to present the client's authentication credentials. If the OAuth2 Client
uses the Token Endpoint Authentication Method `client_secret_post`, you need to present the client secret in the URL query.
If it uses `client_secret_basic`, present the Client ID and the Client Secret in the Authorization header.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The id of the OAuth 2.0 Client.
 @return OidcApiGetOidcDynamicClientRequest
*/
func (a *OidcApiService) GetOidcDynamicClient(ctx context.Context, id string) OidcApiGetOidcDynamicClientRequest {
	return OidcApiGetOidcDynamicClientRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return OAuth2Client
func (a *OidcApiService) GetOidcDynamicClientExecute(r OidcApiGetOidcDynamicClientRequest) (*OAuth2Client, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OAuth2Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OidcApiService.GetOidcDynamicClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/register/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorOAuth2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OidcApiGetOidcUserInfoRequest struct {
	ctx context.Context
	ApiService OidcApi
}

func (r OidcApiGetOidcUserInfoRequest) Execute() (*OidcUserInfo, *http.Response, error) {
	return r.ApiService.GetOidcUserInfoExecute(r)
}

/*
GetOidcUserInfo OpenID Connect Userinfo

This endpoint returns the payload of the ID Token, including `session.id_token` values, of
the provided OAuth 2.0 Access Token's consent request.

In the case of authentication error, a WWW-Authenticate header might be set in the response
with more information about the error. See [the spec](https://datatracker.ietf.org/doc/html/rfc6750#section-3)
for more details about header format.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OidcApiGetOidcUserInfoRequest
*/
func (a *OidcApiService) GetOidcUserInfo(ctx context.Context) OidcApiGetOidcUserInfoRequest {
	return OidcApiGetOidcUserInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OidcUserInfo
func (a *OidcApiService) GetOidcUserInfoExecute(r OidcApiGetOidcUserInfoRequest) (*OidcUserInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OidcUserInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OidcApiService.GetOidcUserInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/userinfo"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorOAuth2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OidcApiRevokeOidcSessionRequest struct {
	ctx context.Context
	ApiService OidcApi
}

func (r OidcApiRevokeOidcSessionRequest) Execute() (*http.Response, error) {
	return r.ApiService.RevokeOidcSessionExecute(r)
}

/*
RevokeOidcSession OpenID Connect Front- and Back-channel Enabled Logout

This endpoint initiates and completes user logout at the Ory OAuth2 & OpenID provider and initiates OpenID Connect Front- / Back-channel logout:

https://openid.net/specs/openid-connect-frontchannel-1_0.html
https://openid.net/specs/openid-connect-backchannel-1_0.html

Back-channel logout is performed asynchronously and does not affect logout flow.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OidcApiRevokeOidcSessionRequest
*/
func (a *OidcApiService) RevokeOidcSession(ctx context.Context) OidcApiRevokeOidcSessionRequest {
	return OidcApiRevokeOidcSessionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *OidcApiService) RevokeOidcSessionExecute(r OidcApiRevokeOidcSessionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OidcApiService.RevokeOidcSession")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/sessions/logout"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type OidcApiSetOidcDynamicClientRequest struct {
	ctx context.Context
	ApiService OidcApi
	id string
	oAuth2Client *OAuth2Client
}

// OAuth 2.0 Client Request Body
func (r OidcApiSetOidcDynamicClientRequest) OAuth2Client(oAuth2Client OAuth2Client) OidcApiSetOidcDynamicClientRequest {
	r.oAuth2Client = &oAuth2Client
	return r
}

func (r OidcApiSetOidcDynamicClientRequest) Execute() (*OAuth2Client, *http.Response, error) {
	return r.ApiService.SetOidcDynamicClientExecute(r)
}

/*
SetOidcDynamicClient Set OAuth2 Client using OpenID Dynamic Client Registration

This endpoint behaves like the administrative counterpart (`setOAuth2Client`) but is capable of facing the
public internet directly to be used by third parties. It implements the OpenID Connect
Dynamic Client Registration Protocol.

This feature is disabled per default. It can be enabled by a system administrator.

If you pass `client_secret` the secret is used, otherwise the existing secret is used. If set, the secret is echoed in the response.
It is not possible to retrieve it later on.

To use this endpoint, you will need to present the client's authentication credentials. If the OAuth2 Client
uses the Token Endpoint Authentication Method `client_secret_post`, you need to present the client secret in the URL query.
If it uses `client_secret_basic`, present the Client ID and the Client Secret in the Authorization header.

OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are
generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id OAuth 2.0 Client ID
 @return OidcApiSetOidcDynamicClientRequest
*/
func (a *OidcApiService) SetOidcDynamicClient(ctx context.Context, id string) OidcApiSetOidcDynamicClientRequest {
	return OidcApiSetOidcDynamicClientRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return OAuth2Client
func (a *OidcApiService) SetOidcDynamicClientExecute(r OidcApiSetOidcDynamicClientRequest) (*OAuth2Client, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OAuth2Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OidcApiService.SetOidcDynamicClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/register/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oAuth2Client == nil {
		return localVarReturnValue, nil, reportError("oAuth2Client is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.oAuth2Client
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorOAuth2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorOAuth2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
