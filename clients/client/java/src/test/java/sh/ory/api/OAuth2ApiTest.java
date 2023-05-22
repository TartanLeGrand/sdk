/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.1.33
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.api;

import sh.ory.ApiException;
import sh.ory.model.AcceptOAuth2ConsentRequest;
import sh.ory.model.AcceptOAuth2LoginRequest;
import sh.ory.model.ErrorOAuth2;
import sh.ory.model.GenericError;
import sh.ory.model.IntrospectedOAuth2Token;
import sh.ory.model.JsonPatch;
import sh.ory.model.OAuth2Client;
import sh.ory.model.OAuth2ClientTokenLifespans;
import sh.ory.model.OAuth2ConsentRequest;
import sh.ory.model.OAuth2ConsentSession;
import sh.ory.model.OAuth2LoginRequest;
import sh.ory.model.OAuth2LogoutRequest;
import sh.ory.model.OAuth2RedirectTo;
import sh.ory.model.OAuth2TokenExchange;
import sh.ory.model.RejectOAuth2Request;
import sh.ory.model.TrustOAuth2JwtGrantIssuer;
import sh.ory.model.TrustedOAuth2JwtGrantIssuer;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for OAuth2Api
 */
@Disabled
public class OAuth2ApiTest {

    private final OAuth2Api api = new OAuth2Api();

    /**
     * Accept OAuth 2.0 Consent Request
     *
     * When an authorization code, hybrid, or implicit OAuth 2.0 Flow is initiated, Ory asks the login provider to authenticate the subject and then tell Ory now about it. If the subject authenticated, he/she must now be asked if the OAuth 2.0 Client which initiated the flow should be allowed to access the resources on the subject&#39;s behalf.  The consent challenge is appended to the consent provider&#39;s URL to which the subject&#39;s user-agent (browser) is redirected to. The consent provider uses that challenge to fetch information on the OAuth2 request and then tells Ory if the subject accepted or rejected the request.  This endpoint tells Ory that the subject has authorized the OAuth 2.0 client to access resources on his/her behalf. The consent provider includes additional information, such as session data for access and ID tokens, and if the consent request should be used as basis for future requests.  The response contains a redirect URL which the consent provider should redirect the user-agent to.  The default consent provider is available via the Ory Managed Account Experience. To customize the consent provider, please head over to the OAuth 2.0 documentation.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void acceptOAuth2ConsentRequestTest() throws ApiException {
        String consentChallenge = null;
        AcceptOAuth2ConsentRequest acceptOAuth2ConsentRequest = null;
        OAuth2RedirectTo response = api.acceptOAuth2ConsentRequest(consentChallenge, acceptOAuth2ConsentRequest);
        // TODO: test validations
    }

    /**
     * Accept OAuth 2.0 Login Request
     *
     * When an authorization code, hybrid, or implicit OAuth 2.0 Flow is initiated, Ory asks the login provider to authenticate the subject and then tell the Ory OAuth2 Service about it.  The authentication challenge is appended to the login provider URL to which the subject&#39;s user-agent (browser) is redirected to. The login provider uses that challenge to fetch information on the OAuth2 request and then accept or reject the requested authentication process.  This endpoint tells Ory that the subject has successfully authenticated and includes additional information such as the subject&#39;s ID and if Ory should remember the subject&#39;s subject agent for future authentication attempts by setting a cookie.  The response contains a redirect URL which the login provider should redirect the user-agent to.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void acceptOAuth2LoginRequestTest() throws ApiException {
        String loginChallenge = null;
        AcceptOAuth2LoginRequest acceptOAuth2LoginRequest = null;
        OAuth2RedirectTo response = api.acceptOAuth2LoginRequest(loginChallenge, acceptOAuth2LoginRequest);
        // TODO: test validations
    }

    /**
     * Accept OAuth 2.0 Session Logout Request
     *
     * When a user or an application requests Ory OAuth 2.0 to remove the session state of a subject, this endpoint is used to confirm that logout request.  The response contains a redirect URL which the consent provider should redirect the user-agent to.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void acceptOAuth2LogoutRequestTest() throws ApiException {
        String logoutChallenge = null;
        OAuth2RedirectTo response = api.acceptOAuth2LogoutRequest(logoutChallenge);
        // TODO: test validations
    }

    /**
     * Create OAuth 2.0 Client
     *
     * Create a new OAuth 2.0 client. If you pass &#x60;client_secret&#x60; the secret is used, otherwise a random secret is generated. The secret is echoed in the response. It is not possible to retrieve it later on.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void createOAuth2ClientTest() throws ApiException {
        OAuth2Client oauth2Client = null;
        OAuth2Client response = api.createOAuth2Client(oauth2Client);
        // TODO: test validations
    }

    /**
     * Delete OAuth 2.0 Client
     *
     * Delete an existing OAuth 2.0 Client by its ID.  OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.  Make sure that this endpoint is well protected and only callable by first-party components.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void deleteOAuth2ClientTest() throws ApiException {
        String id = null;
        api.deleteOAuth2Client(id);
        // TODO: test validations
    }

    /**
     * Delete OAuth 2.0 Access Tokens from specific OAuth 2.0 Client
     *
     * This endpoint deletes OAuth2 access tokens issued to an OAuth 2.0 Client from the database.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void deleteOAuth2TokenTest() throws ApiException {
        String clientId = null;
        api.deleteOAuth2Token(clientId);
        // TODO: test validations
    }

    /**
     * Delete Trusted OAuth2 JWT Bearer Grant Type Issuer
     *
     * Use this endpoint to delete trusted JWT Bearer Grant Type Issuer. The ID is the one returned when you created the trust relationship.  Once deleted, the associated issuer will no longer be able to perform the JSON Web Token (JWT) Profile for OAuth 2.0 Client Authentication and Authorization Grant.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void deleteTrustedOAuth2JwtGrantIssuerTest() throws ApiException {
        String id = null;
        api.deleteTrustedOAuth2JwtGrantIssuer(id);
        // TODO: test validations
    }

    /**
     * Get an OAuth 2.0 Client
     *
     * Get an OAuth 2.0 client by its ID. This endpoint never returns the client secret.  OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getOAuth2ClientTest() throws ApiException {
        String id = null;
        OAuth2Client response = api.getOAuth2Client(id);
        // TODO: test validations
    }

    /**
     * Get OAuth 2.0 Consent Request
     *
     * When an authorization code, hybrid, or implicit OAuth 2.0 Flow is initiated, Ory asks the login provider to authenticate the subject and then tell Ory now about it. If the subject authenticated, he/she must now be asked if the OAuth 2.0 Client which initiated the flow should be allowed to access the resources on the subject&#39;s behalf.  The consent challenge is appended to the consent provider&#39;s URL to which the subject&#39;s user-agent (browser) is redirected to. The consent provider uses that challenge to fetch information on the OAuth2 request and then tells Ory if the subject accepted or rejected the request.  The default consent provider is available via the Ory Managed Account Experience. To customize the consent provider, please head over to the OAuth 2.0 documentation.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getOAuth2ConsentRequestTest() throws ApiException {
        String consentChallenge = null;
        OAuth2ConsentRequest response = api.getOAuth2ConsentRequest(consentChallenge);
        // TODO: test validations
    }

    /**
     * Get OAuth 2.0 Login Request
     *
     * When an authorization code, hybrid, or implicit OAuth 2.0 Flow is initiated, Ory asks the login provider to authenticate the subject and then tell the Ory OAuth2 Service about it.  Per default, the login provider is Ory itself. You may use a different login provider which needs to be a web-app you write and host, and it must be able to authenticate (\&quot;show the subject a login screen\&quot;) a subject (in OAuth2 the proper name for subject is \&quot;resource owner\&quot;).  The authentication challenge is appended to the login provider URL to which the subject&#39;s user-agent (browser) is redirected to. The login provider uses that challenge to fetch information on the OAuth2 request and then accept or reject the requested authentication process.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getOAuth2LoginRequestTest() throws ApiException {
        String loginChallenge = null;
        OAuth2LoginRequest response = api.getOAuth2LoginRequest(loginChallenge);
        // TODO: test validations
    }

    /**
     * Get OAuth 2.0 Session Logout Request
     *
     * Use this endpoint to fetch an Ory OAuth 2.0 logout request.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getOAuth2LogoutRequestTest() throws ApiException {
        String logoutChallenge = null;
        OAuth2LogoutRequest response = api.getOAuth2LogoutRequest(logoutChallenge);
        // TODO: test validations
    }

    /**
     * Get Trusted OAuth2 JWT Bearer Grant Type Issuer
     *
     * Use this endpoint to get a trusted JWT Bearer Grant Type Issuer. The ID is the one returned when you created the trust relationship.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getTrustedOAuth2JwtGrantIssuerTest() throws ApiException {
        String id = null;
        TrustedOAuth2JwtGrantIssuer response = api.getTrustedOAuth2JwtGrantIssuer(id);
        // TODO: test validations
    }

    /**
     * Introspect OAuth2 Access and Refresh Tokens
     *
     * The introspection endpoint allows to check if a token (both refresh and access) is active or not. An active token is neither expired nor revoked. If a token is active, additional information on the token will be included. You can set additional data for a token by setting &#x60;session.access_token&#x60; during the consent flow.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void introspectOAuth2TokenTest() throws ApiException {
        String token = null;
        String scope = null;
        IntrospectedOAuth2Token response = api.introspectOAuth2Token(token, scope);
        // TODO: test validations
    }

    /**
     * List OAuth 2.0 Clients
     *
     * This endpoint lists all clients in the database, and never returns client secrets. As a default it lists the first 100 clients.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listOAuth2ClientsTest() throws ApiException {
        Long pageSize = null;
        String pageToken = null;
        String clientName = null;
        String owner = null;
        List<OAuth2Client> response = api.listOAuth2Clients(pageSize, pageToken, clientName, owner);
        // TODO: test validations
    }

    /**
     * List OAuth 2.0 Consent Sessions of a Subject
     *
     * This endpoint lists all subject&#39;s granted consent sessions, including client and granted scope. If the subject is unknown or has not granted any consent sessions yet, the endpoint returns an empty JSON array with status code 200 OK.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listOAuth2ConsentSessionsTest() throws ApiException {
        String subject = null;
        Long pageSize = null;
        String pageToken = null;
        String loginSessionId = null;
        List<OAuth2ConsentSession> response = api.listOAuth2ConsentSessions(subject, pageSize, pageToken, loginSessionId);
        // TODO: test validations
    }

    /**
     * List Trusted OAuth2 JWT Bearer Grant Type Issuers
     *
     * Use this endpoint to list all trusted JWT Bearer Grant Type Issuers.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listTrustedOAuth2JwtGrantIssuersTest() throws ApiException {
        Long maxItems = null;
        Long defaultItems = null;
        String issuer = null;
        List<TrustedOAuth2JwtGrantIssuer> response = api.listTrustedOAuth2JwtGrantIssuers(maxItems, defaultItems, issuer);
        // TODO: test validations
    }

    /**
     * OAuth 2.0 Authorize Endpoint
     *
     * Use open source libraries to perform OAuth 2.0 and OpenID Connect available for any programming language. You can find a list of libraries at https://oauth.net/code/  The Ory SDK is not yet able to this endpoint properly.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void oAuth2AuthorizeTest() throws ApiException {
        ErrorOAuth2 response = api.oAuth2Authorize();
        // TODO: test validations
    }

    /**
     * The OAuth 2.0 Token Endpoint
     *
     * Use open source libraries to perform OAuth 2.0 and OpenID Connect available for any programming language. You can find a list of libraries here https://oauth.net/code/  The Ory SDK is not yet able to this endpoint properly.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void oauth2TokenExchangeTest() throws ApiException {
        String grantType = null;
        String clientId = null;
        String code = null;
        String redirectUri = null;
        String refreshToken = null;
        OAuth2TokenExchange response = api.oauth2TokenExchange(grantType, clientId, code, redirectUri, refreshToken);
        // TODO: test validations
    }

    /**
     * Patch OAuth 2.0 Client
     *
     * Patch an existing OAuth 2.0 Client using JSON Patch. If you pass &#x60;client_secret&#x60; the secret will be updated and returned via the API. This is the only time you will be able to retrieve the client secret, so write it down and keep it safe.  OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void patchOAuth2ClientTest() throws ApiException {
        String id = null;
        List<JsonPatch> jsonPatch = null;
        OAuth2Client response = api.patchOAuth2Client(id, jsonPatch);
        // TODO: test validations
    }

    /**
     * Reject OAuth 2.0 Consent Request
     *
     * When an authorization code, hybrid, or implicit OAuth 2.0 Flow is initiated, Ory asks the login provider to authenticate the subject and then tell Ory now about it. If the subject authenticated, he/she must now be asked if the OAuth 2.0 Client which initiated the flow should be allowed to access the resources on the subject&#39;s behalf.  The consent challenge is appended to the consent provider&#39;s URL to which the subject&#39;s user-agent (browser) is redirected to. The consent provider uses that challenge to fetch information on the OAuth2 request and then tells Ory if the subject accepted or rejected the request.  This endpoint tells Ory that the subject has not authorized the OAuth 2.0 client to access resources on his/her behalf. The consent provider must include a reason why the consent was not granted.  The response contains a redirect URL which the consent provider should redirect the user-agent to.  The default consent provider is available via the Ory Managed Account Experience. To customize the consent provider, please head over to the OAuth 2.0 documentation.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void rejectOAuth2ConsentRequestTest() throws ApiException {
        String consentChallenge = null;
        RejectOAuth2Request rejectOAuth2Request = null;
        OAuth2RedirectTo response = api.rejectOAuth2ConsentRequest(consentChallenge, rejectOAuth2Request);
        // TODO: test validations
    }

    /**
     * Reject OAuth 2.0 Login Request
     *
     * When an authorization code, hybrid, or implicit OAuth 2.0 Flow is initiated, Ory asks the login provider to authenticate the subject and then tell the Ory OAuth2 Service about it.  The authentication challenge is appended to the login provider URL to which the subject&#39;s user-agent (browser) is redirected to. The login provider uses that challenge to fetch information on the OAuth2 request and then accept or reject the requested authentication process.  This endpoint tells Ory that the subject has not authenticated and includes a reason why the authentication was denied.  The response contains a redirect URL which the login provider should redirect the user-agent to.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void rejectOAuth2LoginRequestTest() throws ApiException {
        String loginChallenge = null;
        RejectOAuth2Request rejectOAuth2Request = null;
        OAuth2RedirectTo response = api.rejectOAuth2LoginRequest(loginChallenge, rejectOAuth2Request);
        // TODO: test validations
    }

    /**
     * Reject OAuth 2.0 Session Logout Request
     *
     * When a user or an application requests Ory OAuth 2.0 to remove the session state of a subject, this endpoint is used to deny that logout request. No HTTP request body is required.  The response is empty as the logout provider has to chose what action to perform next.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void rejectOAuth2LogoutRequestTest() throws ApiException {
        String logoutChallenge = null;
        api.rejectOAuth2LogoutRequest(logoutChallenge);
        // TODO: test validations
    }

    /**
     * Revoke OAuth 2.0 Consent Sessions of a Subject
     *
     * This endpoint revokes a subject&#39;s granted consent sessions and invalidates all associated OAuth 2.0 Access Tokens. You may also only revoke sessions for a specific OAuth 2.0 Client ID.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void revokeOAuth2ConsentSessionsTest() throws ApiException {
        String subject = null;
        String client = null;
        Boolean all = null;
        api.revokeOAuth2ConsentSessions(subject, client, all);
        // TODO: test validations
    }

    /**
     * Revokes OAuth 2.0 Login Sessions by either a Subject or a SessionID
     *
     * This endpoint invalidates authentication sessions. After revoking the authentication session(s), the subject has to re-authenticate at the Ory OAuth2 Provider. This endpoint does not invalidate any tokens.  If you send the subject in a query param, all authentication sessions that belong to that subject are revoked. No OpennID Connect Front- or Back-channel logout is performed in this case.  Alternatively, you can send a SessionID via &#x60;sid&#x60; query param, in which case, only the session that is connected to that SessionID is revoked. OpenID Connect Back-channel logout is performed in this case.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void revokeOAuth2LoginSessionsTest() throws ApiException {
        String subject = null;
        String sid = null;
        api.revokeOAuth2LoginSessions(subject, sid);
        // TODO: test validations
    }

    /**
     * Revoke OAuth 2.0 Access or Refresh Token
     *
     * Revoking a token (both access and refresh) means that the tokens will be invalid. A revoked access token can no longer be used to make access requests, and a revoked refresh token can no longer be used to refresh an access token. Revoking a refresh token also invalidates the access token that was created with it. A token may only be revoked by the client the token was generated for.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void revokeOAuth2TokenTest() throws ApiException {
        String token = null;
        String clientId = null;
        String clientSecret = null;
        api.revokeOAuth2Token(token, clientId, clientSecret);
        // TODO: test validations
    }

    /**
     * Set OAuth 2.0 Client
     *
     * Replaces an existing OAuth 2.0 Client with the payload you send. If you pass &#x60;client_secret&#x60; the secret is used, otherwise the existing secret is used.  If set, the secret is echoed in the response. It is not possible to retrieve it later on.  OAuth 2.0 Clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void setOAuth2ClientTest() throws ApiException {
        String id = null;
        OAuth2Client oauth2Client = null;
        OAuth2Client response = api.setOAuth2Client(id, oauth2Client);
        // TODO: test validations
    }

    /**
     * Set OAuth2 Client Token Lifespans
     *
     * Set lifespans of different token types issued for this OAuth 2.0 client. Does not modify other fields.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void setOAuth2ClientLifespansTest() throws ApiException {
        String id = null;
        OAuth2ClientTokenLifespans oauth2ClientTokenLifespans = null;
        OAuth2Client response = api.setOAuth2ClientLifespans(id, oauth2ClientTokenLifespans);
        // TODO: test validations
    }

    /**
     * Trust OAuth2 JWT Bearer Grant Type Issuer
     *
     * Use this endpoint to establish a trust relationship for a JWT issuer to perform JSON Web Token (JWT) Profile for OAuth 2.0 Client Authentication and Authorization Grants [RFC7523](https://datatracker.ietf.org/doc/html/rfc7523).
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void trustOAuth2JwtGrantIssuerTest() throws ApiException {
        TrustOAuth2JwtGrantIssuer trustOAuth2JwtGrantIssuer = null;
        TrustedOAuth2JwtGrantIssuer response = api.trustOAuth2JwtGrantIssuer(trustOAuth2JwtGrantIssuer);
        // TODO: test validations
    }

}
