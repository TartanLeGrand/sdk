/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.7.3-alpha.2
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// ContainerUpdateOkBody : ContainerUpdateOKBody OK response to ContainerUpdate operation



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct ContainerUpdateOkBody {
    /// warnings
    #[serde(rename = "Warnings")]
    pub warnings: Vec<String>,
}

impl ContainerUpdateOkBody {
    /// ContainerUpdateOKBody OK response to ContainerUpdate operation
    pub fn new(warnings: Vec<String>) -> ContainerUpdateOkBody {
        ContainerUpdateOkBody {
            warnings,
        }
    }
}


