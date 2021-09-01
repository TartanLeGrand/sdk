/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.18
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// IdentityCredentials : Credentials represents a specific credential type



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct IdentityCredentials {
    #[serde(rename = "config", skip_serializing_if = "Option::is_none")]
    pub config: Option<serde_json::Value>,
    /// CreatedAt is a helper struct field for gobuffalo.pop.
    #[serde(rename = "created_at", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<String>,
    /// Identifiers represents a list of unique identifiers this credential type matches.
    #[serde(rename = "identifiers", skip_serializing_if = "Option::is_none")]
    pub identifiers: Option<Vec<String>>,
    #[serde(rename = "type", skip_serializing_if = "Option::is_none")]
    pub _type: Option<crate::models::IdentityCredentialsType>,
    /// UpdatedAt is a helper struct field for gobuffalo.pop.
    #[serde(rename = "updated_at", skip_serializing_if = "Option::is_none")]
    pub updated_at: Option<String>,
}

impl IdentityCredentials {
    /// Credentials represents a specific credential type
    pub fn new() -> IdentityCredentials {
        IdentityCredentials {
            config: None,
            created_at: None,
            identifiers: None,
            _type: None,
            updated_at: None,
        }
    }
}


