/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.2.0-alpha.18
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */


use std::fmt::Display;

use num_traits;
use reqwest;

use crate::apis::ResponseContent;
use super::{Error, configuration};


trait NumVecJoin {
    fn join(&self, sep: &str) -> String;
}

impl <T: Display + num_traits::Num> NumVecJoin for Vec<T> {
    fn join(&self, sep: &str) -> String {
        self.iter()
            .map(ToString::to_string)
            .collect::<Vec<String>>()
            .join(sep)
    }
}


/// struct for typed errors of method `update_o_auth2_client_lifespans`
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum UpdateOAuth2ClientLifespansError {
    DefaultResponse(crate::models::GenericError),
    UnknownValue(serde_json::Value),
}


/// UpdateLifespans an existing OAuth 2.0 client's token lifespan configuration. This client configuration takes precedence over the instance-wide token lifespan configuration.
pub async fn update_o_auth2_client_lifespans(configuration: &configuration::Configuration, id: &str, update_o_auth2_client_lifespans: Option<&crate::models::UpdateOAuth2ClientLifespans>) -> Result<crate::models::OAuth2Client, Error<UpdateOAuth2ClientLifespansError>> {

    let local_var_client = &configuration.client;

    let local_var_uri_str = format!("{}/admin/clients/{id}/lifespans", configuration.base_path, id=crate::apis::urlencode(id));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::PUT, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    local_var_req_builder = local_var_req_builder.json(&update_o_auth2_client_lifespans);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<UpdateOAuth2ClientLifespansError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

