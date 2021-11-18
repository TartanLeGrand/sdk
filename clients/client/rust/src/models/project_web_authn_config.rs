/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.23
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct ProjectWebAuthnConfig {
    /// Set to true to enable the WebAuthn authentication method.
    #[serde(rename = "enabled", skip_serializing_if = "Option::is_none")]
    pub enabled: Option<bool>,
    /// The display name to show in the authenticator.
    #[serde(rename = "rp_display_name", skip_serializing_if = "Option::is_none")]
    pub rp_display_name: Option<String>,
    /// A URL to an icon which might be shown in the authenticator.
    #[serde(rename = "rp_icon", skip_serializing_if = "Option::is_none")]
    pub rp_icon: Option<String>,
    /// This must be the hostname of the login page.
    #[serde(rename = "rp_id", skip_serializing_if = "Option::is_none")]
    pub rp_id: Option<String>,
    /// This must be the scheme://hostname of the login page.
    #[serde(rename = "rp_origin", skip_serializing_if = "Option::is_none")]
    pub rp_origin: Option<String>,
}

impl ProjectWebAuthnConfig {
    pub fn new() -> ProjectWebAuthnConfig {
        ProjectWebAuthnConfig {
            enabled: None,
            rp_display_name: None,
            rp_icon: None,
            rp_id: None,
            rp_origin: None,
        }
    }
}


