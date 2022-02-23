/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.107
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// GenericError : Error responses are sent when an error (e.g. unauthorized, bad request, ...) occurred.



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct GenericError {
    /// The status code
    #[serde(rename = "code", skip_serializing_if = "Option::is_none")]
    pub code: Option<i64>,
    /// Debug contains debug information. This is usually not available and has to be enabled.
    #[serde(rename = "debug", skip_serializing_if = "Option::is_none")]
    pub debug: Option<String>,
    /// Further error details
    #[serde(rename = "details", skip_serializing_if = "Option::is_none")]
    pub details: Option<serde_json::Value>,
    /// Name is the error name.
    #[serde(rename = "error", skip_serializing_if = "Option::is_none")]
    pub error: Option<String>,
    /// Description contains further information on the nature of the error.
    #[serde(rename = "error_description", skip_serializing_if = "Option::is_none")]
    pub error_description: Option<String>,
    /// The error ID  Useful when trying to identify various errors in application logic.
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    /// Message contains the error message.
    #[serde(rename = "message")]
    pub message: String,
    /// A human-readable reason for the error
    #[serde(rename = "reason", skip_serializing_if = "Option::is_none")]
    pub reason: Option<String>,
    /// The request ID  The request ID is often exposed internally in order to trace errors across service architectures. This is often a UUID.
    #[serde(rename = "request", skip_serializing_if = "Option::is_none")]
    pub request: Option<String>,
    /// The status description
    #[serde(rename = "status", skip_serializing_if = "Option::is_none")]
    pub status: Option<String>,
    /// Code represents the error status code (404, 403, 401, ...).
    #[serde(rename = "status_code", skip_serializing_if = "Option::is_none")]
    pub status_code: Option<i64>,
}

impl GenericError {
    /// Error responses are sent when an error (e.g. unauthorized, bad request, ...) occurred.
    pub fn new(message: String) -> GenericError {
        GenericError {
            code: None,
            debug: None,
            details: None,
            error: None,
            error_description: None,
            id: None,
            message,
            reason: None,
            request: None,
            status: None,
            status_code: None,
        }
    }
}


