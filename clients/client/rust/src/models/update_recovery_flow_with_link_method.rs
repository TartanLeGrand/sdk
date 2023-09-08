/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.2.4
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// UpdateRecoveryFlowWithLinkMethod : Update Recovery Flow with Link Method



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct UpdateRecoveryFlowWithLinkMethod {
    /// Sending the anti-csrf token is only required for browser login flows.
    #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
    pub csrf_token: Option<String>,
    /// Email to Recover  Needs to be set when initiating the flow. If the email is a registered recovery email, a recovery link will be sent. If the email is not known, a email with details on what happened will be sent instead.  format: email
    #[serde(rename = "email")]
    pub email: String,
    /// Method is the method that should be used for this recovery flow  Allowed values are `link` and `code` link RecoveryStrategyLink code RecoveryStrategyCode
    #[serde(rename = "method")]
    pub method: MethodEnum,
}


impl UpdateRecoveryFlowWithLinkMethod {
    /// Update Recovery Flow with Link Method
    pub fn new(email: String, method: MethodEnum) -> UpdateRecoveryFlowWithLinkMethod {
        UpdateRecoveryFlowWithLinkMethod {
                csrf_token: None,
                email,
                method,
        }
    }
}

/// Method is the method that should be used for this recovery flow  Allowed values are `link` and `code` link RecoveryStrategyLink code RecoveryStrategyCode
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum MethodEnum {
    #[serde(rename = "link")]
    Link,
    #[serde(rename = "code")]
    Code,
}

