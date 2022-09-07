/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.2.0-alpha.33
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct AdminCreateIdentityImportCredentialsPassword {
    #[serde(rename = "config", skip_serializing_if = "Option::is_none")]
    pub config: Option<Box<crate::models::AdminCreateIdentityImportCredentialsPasswordConfig>>,
}

impl Default for AdminCreateIdentityImportCredentialsPassword {
    fn default() -> Self {
        Self::new()
    }
}

impl AdminCreateIdentityImportCredentialsPassword {
    pub fn new() -> AdminCreateIdentityImportCredentialsPassword {
        AdminCreateIdentityImportCredentialsPassword {
                config: None,
        }
    }
}


