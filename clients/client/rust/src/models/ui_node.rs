/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.23
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// UiNode : Nodes are represented as HTML elements or their native UI equivalents. For example, a node can be an `<img>` tag, or an `<input element>` but also `some plain text`.



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct UiNode {
    #[serde(rename = "attributes")]
    pub attributes: Box<crate::models::UiNodeAttributes>,
    #[serde(rename = "group")]
    pub group: String,
    #[serde(rename = "messages")]
    pub messages: Vec<crate::models::UiText>,
    #[serde(rename = "meta")]
    pub meta: Box<crate::models::UiNodeMeta>,
    #[serde(rename = "type")]
    pub _type: String,
}

impl UiNode {
    /// Nodes are represented as HTML elements or their native UI equivalents. For example, a node can be an `<img>` tag, or an `<input element>` but also `some plain text`.
    pub fn new(attributes: crate::models::UiNodeAttributes, group: String, messages: Vec<crate::models::UiText>, meta: crate::models::UiNodeMeta, _type: String) -> UiNode {
        UiNode {
            attributes: Box::new(attributes),
            group,
            messages,
            meta: Box::new(meta),
            _type,
        }
    }
}


