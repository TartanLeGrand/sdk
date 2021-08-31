/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.7.3-alpha.2
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// SubmitSelfServiceLoginFlowWithOidcMethodBody : SubmitSelfServiceLoginFlowWithOidcMethodBody is used to decode the login form payload when using the oidc method.



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct SubmitSelfServiceLoginFlowWithOidcMethodBody {
    /// The CSRF Token
    #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
    pub csrf_token: Option<String>,
    /// Method to use  This field must be set to `oidc` when using the oidc method.
    #[serde(rename = "method")]
    pub method: String,
    /// The provider to register with
    #[serde(rename = "traits")]
    pub traits: String,
}

impl SubmitSelfServiceLoginFlowWithOidcMethodBody {
    /// SubmitSelfServiceLoginFlowWithOidcMethodBody is used to decode the login form payload when using the oidc method.
    pub fn new(method: String, traits: String) -> SubmitSelfServiceLoginFlowWithOidcMethodBody {
        SubmitSelfServiceLoginFlowWithOidcMethodBody {
            csrf_token: None,
            method,
            traits,
        }
    }
}


