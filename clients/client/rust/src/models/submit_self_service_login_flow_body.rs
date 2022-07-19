/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.1.0-alpha.3
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
#[serde(tag = "method")]
pub enum SubmitSelfServiceLoginFlowBody {
    #[serde(rename="lookup_secret")]
    SubmitSelfServiceLoginFlowWithLookupSecretMethodBody {
        /// Sending the anti-csrf token is only required for browser login flows.
        #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
        csrf_token: Option<String>,
        /// The lookup secret.
        #[serde(rename = "lookup_secret")]
        lookup_secret: String,
    },
    #[serde(rename="oidc")]
    SubmitSelfServiceLoginFlowWithOidcMethodBody {
        /// The CSRF Token
        #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
        csrf_token: Option<String>,
        /// The provider to register with
        #[serde(rename = "provider")]
        provider: String,
        /// The identity traits. This is a placeholder for the registration flow.
        #[serde(rename = "traits", skip_serializing_if = "Option::is_none")]
        traits: Option<serde_json::Value>,
    },
    #[serde(rename="password")]
    SubmitSelfServiceLoginFlowWithPasswordMethodBody {
        /// Sending the anti-csrf token is only required for browser login flows.
        #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
        csrf_token: Option<String>,
        /// Identifier is the email or username of the user trying to log in.
        #[serde(rename = "identifier")]
        identifier: String,
        /// The user's password.
        #[serde(rename = "password")]
        password: String,
        /// Identifier is the email or username of the user trying to log in. This field is deprecated!
        #[serde(rename = "password_identifier", skip_serializing_if = "Option::is_none")]
        password_identifier: Option<String>,
    },
    #[serde(rename="totp")]
    SubmitSelfServiceLoginFlowWithTotpMethodBody {
        /// Sending the anti-csrf token is only required for browser login flows.
        #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
        csrf_token: Option<String>,
        /// The TOTP code.
        #[serde(rename = "totp_code")]
        totp_code: String,
    },
    #[serde(rename="webauthn")]
    SubmitSelfServiceLoginFlowWithWebAuthnMethodBody {
        /// Sending the anti-csrf token is only required for browser login flows.
        #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
        csrf_token: Option<String>,
        /// Identifier is the email or username of the user trying to log in. This field is only required when using WebAuthn for passwordless login. When using WebAuthn for multi-factor authentication, it is not needed.
        #[serde(rename = "identifier", skip_serializing_if = "Option::is_none")]
        identifier: Option<String>,
        /// Login a WebAuthn Security Key  This must contain the ID of the WebAuthN connection.
        #[serde(rename = "webauthn_login", skip_serializing_if = "Option::is_none")]
        webauthn_login: Option<String>,
    },
}




