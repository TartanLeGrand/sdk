/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.131
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct SchemaPatch {
    /// The json schema
    #[serde(rename = "data")]
    pub data: serde_json::Value,
    /// The user defined schema name
    #[serde(rename = "name")]
    pub name: String,
}

impl SchemaPatch {
    pub fn new(data: serde_json::Value, name: String) -> SchemaPatch {
        SchemaPatch {
            data,
            name,
        }
    }
}


