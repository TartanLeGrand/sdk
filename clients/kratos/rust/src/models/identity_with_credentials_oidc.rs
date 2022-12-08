/*
 * Ory Identities API
 *
 * This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 
 *
 * The version of the OpenAPI document: v0.11.0
 * Contact: office@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// IdentityWithCredentialsOidc : Create Identity and Import Social Sign In Credentials



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct IdentityWithCredentialsOidc {
    #[serde(rename = "config", skip_serializing_if = "Option::is_none")]
    pub config: Option<Box<crate::models::IdentityWithCredentialsOidcConfig>>,
}

impl Default for IdentityWithCredentialsOidc {
    fn default() -> Self {
        Self::new()
    }
}

impl IdentityWithCredentialsOidc {
    /// Create Identity and Import Social Sign In Credentials
    pub fn new() -> IdentityWithCredentialsOidc {
        IdentityWithCredentialsOidc {
                config: None,
        }
    }
}

