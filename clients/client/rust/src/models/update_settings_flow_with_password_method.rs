/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.2.4
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// UpdateSettingsFlowWithPasswordMethod : Update Settings Flow with Password Method



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct UpdateSettingsFlowWithPasswordMethod {
    /// CSRFToken is the anti-CSRF token
    #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
    pub csrf_token: Option<String>,
    /// Method  Should be set to password when trying to update a password.
    #[serde(rename = "method")]
    pub method: String,
    /// Password is the updated password
    #[serde(rename = "password")]
    pub password: String,
}


impl UpdateSettingsFlowWithPasswordMethod {
    /// Update Settings Flow with Password Method
    pub fn new(method: String, password: String) -> UpdateSettingsFlowWithPasswordMethod {
        UpdateSettingsFlowWithPasswordMethod {
                csrf_token: None,
                method,
                password,
        }
    }
}


