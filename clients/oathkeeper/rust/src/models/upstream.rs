/*
 * ORY Oathkeeper
 *
 * ORY Oathkeeper is a reverse proxy that checks the HTTP Authorization for validity against a set of rules. This service uses Hydra to validate access tokens and policies.
 *
 * The version of the OpenAPI document: v0.0.0-alpha.62
 * Contact: hi@ory.am
 * Generated by: https://openapi-generator.tech
 */

/// Upstream : Upstream Upstream Upstream Upstream Upstream Upstream Upstream upstream



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct Upstream {
    /// PreserveHost, if false (the default), tells ORY Oathkeeper to set the upstream request's Host header to the hostname of the API's upstream's URL. Setting this flag to true instructs ORY Oathkeeper not to do so.
    #[serde(rename = "preserve_host", skip_serializing_if = "Option::is_none")]
    pub preserve_host: Option<bool>,
    /// StripPath if set, replaces the provided path prefix when forwarding the requested URL to the upstream URL.
    #[serde(rename = "strip_path", skip_serializing_if = "Option::is_none")]
    pub strip_path: Option<String>,
    /// URL is the URL the request will be proxied to.
    #[serde(rename = "url", skip_serializing_if = "Option::is_none")]
    pub url: Option<String>,
}

impl Upstream {
    /// Upstream Upstream Upstream Upstream Upstream Upstream Upstream upstream
    pub fn new() -> Upstream {
        Upstream {
            preserve_host: None,
            strip_path: None,
            url: None,
        }
    }
}


