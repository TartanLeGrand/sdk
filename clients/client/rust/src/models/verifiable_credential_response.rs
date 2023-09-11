/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.2.5
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct VerifiableCredentialResponse {
    #[serde(rename = "credential_draft_00", skip_serializing_if = "Option::is_none")]
    pub credential_draft_00: Option<String>,
    #[serde(rename = "format", skip_serializing_if = "Option::is_none")]
    pub format: Option<String>,
}

impl Default for VerifiableCredentialResponse {
    fn default() -> Self {
        Self::new()
    }
}

impl VerifiableCredentialResponse {
    pub fn new() -> VerifiableCredentialResponse {
        VerifiableCredentialResponse {
                credential_draft_00: None,
                format: None,
        }
    }
}


