/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.1.11
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


/// struct for typed errors of method `create_identity`
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum CreateIdentityError {
    Status400(crate::models::ErrorGeneric),
    Status409(crate::models::ErrorGeneric),
    DefaultResponse(crate::models::ErrorGeneric),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method `create_recovery_code_for_identity`
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum CreateRecoveryCodeForIdentityError {
    Status400(crate::models::ErrorGeneric),
    Status404(crate::models::ErrorGeneric),
    DefaultResponse(crate::models::ErrorGeneric),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method `create_recovery_link_for_identity`
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum CreateRecoveryLinkForIdentityError {
    Status400(crate::models::ErrorGeneric),
    Status404(crate::models::ErrorGeneric),
    DefaultResponse(crate::models::ErrorGeneric),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method `delete_identity`
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum DeleteIdentityError {
    Status404(crate::models::ErrorGeneric),
    DefaultResponse(crate::models::ErrorGeneric),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method `delete_identity_credentials`
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum DeleteIdentityCredentialsError {
    Status404(crate::models::ErrorGeneric),
    DefaultResponse(crate::models::ErrorGeneric),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method `delete_identity_sessions`
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum DeleteIdentitySessionsError {
    Status400(crate::models::ErrorGeneric),
    Status401(crate::models::ErrorGeneric),
    Status404(crate::models::ErrorGeneric),
    DefaultResponse(crate::models::ErrorGeneric),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method `disable_session`
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum DisableSessionError {
    Status400(crate::models::ErrorGeneric),
    Status401(crate::models::ErrorGeneric),
    DefaultResponse(crate::models::ErrorGeneric),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method `extend_session`
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ExtendSessionError {
    Status400(crate::models::ErrorGeneric),
    Status404(crate::models::ErrorGeneric),
    DefaultResponse(crate::models::ErrorGeneric),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method `get_identity`
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetIdentityError {
    Status404(crate::models::ErrorGeneric),
    DefaultResponse(crate::models::ErrorGeneric),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method `get_identity_schema`
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetIdentitySchemaError {
    Status404(crate::models::ErrorGeneric),
    DefaultResponse(crate::models::ErrorGeneric),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method `get_session`
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetSessionError {
    Status400(crate::models::ErrorGeneric),
    DefaultResponse(crate::models::ErrorGeneric),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method `list_identities`
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ListIdentitiesError {
    DefaultResponse(crate::models::ErrorGeneric),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method `list_identity_schemas`
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ListIdentitySchemasError {
    DefaultResponse(crate::models::ErrorGeneric),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method `list_identity_sessions`
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ListIdentitySessionsError {
    Status400(crate::models::ErrorGeneric),
    Status404(crate::models::ErrorGeneric),
    DefaultResponse(crate::models::ErrorGeneric),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method `list_sessions`
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ListSessionsError {
    Status400(crate::models::ErrorGeneric),
    DefaultResponse(crate::models::ErrorGeneric),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method `patch_identity`
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum PatchIdentityError {
    Status400(crate::models::ErrorGeneric),
    Status404(crate::models::ErrorGeneric),
    Status409(crate::models::ErrorGeneric),
    DefaultResponse(crate::models::ErrorGeneric),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method `update_identity`
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum UpdateIdentityError {
    Status400(crate::models::ErrorGeneric),
    Status404(crate::models::ErrorGeneric),
    Status409(crate::models::ErrorGeneric),
    DefaultResponse(crate::models::ErrorGeneric),
    UnknownValue(serde_json::Value),
}


/// Create an [identity](https://www.ory.sh/docs/kratos/concepts/identity-user-model).  This endpoint can also be used to [import credentials](https://www.ory.sh/docs/kratos/manage-identities/import-user-accounts-identities) for instance passwords, social sign in configurations or multifactor methods.
pub async fn create_identity(configuration: &configuration::Configuration, create_identity_body: Option<&crate::models::CreateIdentityBody>) -> Result<crate::models::Identity, Error<CreateIdentityError>> {

    let local_var_client = &configuration.client;

    let local_var_uri_str = format!("{}/admin/identities", configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(ref local_var_token) = configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };
    local_var_req_builder = local_var_req_builder.json(&create_identity_body);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<CreateIdentityError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

/// This endpoint creates a recovery code which should be given to the user in order for them to recover (or activate) their account.
pub async fn create_recovery_code_for_identity(configuration: &configuration::Configuration, create_recovery_code_for_identity_body: Option<&crate::models::CreateRecoveryCodeForIdentityBody>) -> Result<crate::models::RecoveryCodeForIdentity, Error<CreateRecoveryCodeForIdentityError>> {

    let local_var_client = &configuration.client;

    let local_var_uri_str = format!("{}/admin/recovery/code", configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(ref local_var_token) = configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };
    local_var_req_builder = local_var_req_builder.json(&create_recovery_code_for_identity_body);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<CreateRecoveryCodeForIdentityError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

/// This endpoint creates a recovery link which should be given to the user in order for them to recover (or activate) their account.
pub async fn create_recovery_link_for_identity(configuration: &configuration::Configuration, create_recovery_link_for_identity_body: Option<&crate::models::CreateRecoveryLinkForIdentityBody>) -> Result<crate::models::RecoveryLinkForIdentity, Error<CreateRecoveryLinkForIdentityError>> {

    let local_var_client = &configuration.client;

    let local_var_uri_str = format!("{}/admin/recovery/link", configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(ref local_var_token) = configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };
    local_var_req_builder = local_var_req_builder.json(&create_recovery_link_for_identity_body);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<CreateRecoveryLinkForIdentityError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

/// Calling this endpoint irrecoverably and permanently deletes the [identity](https://www.ory.sh/docs/kratos/concepts/identity-user-model) given its ID. This action can not be undone. This endpoint returns 204 when the identity was deleted or when the identity was not found, in which case it is assumed that is has been deleted already.
pub async fn delete_identity(configuration: &configuration::Configuration, id: &str) -> Result<(), Error<DeleteIdentityError>> {

    let local_var_client = &configuration.client;

    let local_var_uri_str = format!("{}/admin/identities/{id}", configuration.base_path, id=crate::apis::urlencode(id));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::DELETE, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(ref local_var_token) = configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        Ok(())
    } else {
        let local_var_entity: Option<DeleteIdentityError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

/// Delete an [identity](https://www.ory.sh/docs/kratos/concepts/identity-user-model) credential by its type You can only delete second factor (aal2) credentials.
pub async fn delete_identity_credentials(configuration: &configuration::Configuration, id: &str, _type: &str) -> Result<crate::models::Identity, Error<DeleteIdentityCredentialsError>> {

    let local_var_client = &configuration.client;

    let local_var_uri_str = format!("{}/admin/identities/{id}/credentials/{type}", configuration.base_path, id=crate::apis::urlencode(id), type=crate::apis::urlencode(_type));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::DELETE, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(ref local_var_token) = configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<DeleteIdentityCredentialsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

/// Calling this endpoint irrecoverably and permanently deletes and invalidates all sessions that belong to the given Identity.
pub async fn delete_identity_sessions(configuration: &configuration::Configuration, id: &str) -> Result<(), Error<DeleteIdentitySessionsError>> {

    let local_var_client = &configuration.client;

    let local_var_uri_str = format!("{}/admin/identities/{id}/sessions", configuration.base_path, id=crate::apis::urlencode(id));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::DELETE, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(ref local_var_token) = configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        Ok(())
    } else {
        let local_var_entity: Option<DeleteIdentitySessionsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

/// Calling this endpoint deactivates the specified session. Session data is not deleted.
pub async fn disable_session(configuration: &configuration::Configuration, id: &str) -> Result<(), Error<DisableSessionError>> {

    let local_var_client = &configuration.client;

    let local_var_uri_str = format!("{}/admin/sessions/{id}", configuration.base_path, id=crate::apis::urlencode(id));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::DELETE, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(ref local_var_token) = configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        Ok(())
    } else {
        let local_var_entity: Option<DisableSessionError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

/// Calling this endpoint extends the given session ID. If `session.earliest_possible_extend` is set it will only extend the session after the specified time has passed.  Retrieve the session ID from the `/sessions/whoami` endpoint / `toSession` SDK method.
pub async fn extend_session(configuration: &configuration::Configuration, id: &str) -> Result<crate::models::Session, Error<ExtendSessionError>> {

    let local_var_client = &configuration.client;

    let local_var_uri_str = format!("{}/admin/sessions/{id}/extend", configuration.base_path, id=crate::apis::urlencode(id));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::PATCH, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(ref local_var_token) = configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<ExtendSessionError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

/// Return an [identity](https://www.ory.sh/docs/kratos/concepts/identity-user-model) by its ID. You can optionally include credentials (e.g. social sign in connections) in the response by using the `include_credential` query parameter.
pub async fn get_identity(configuration: &configuration::Configuration, id: &str, include_credential: Option<Vec<String>>) -> Result<crate::models::Identity, Error<GetIdentityError>> {

    let local_var_client = &configuration.client;

    let local_var_uri_str = format!("{}/admin/identities/{id}", configuration.base_path, id=crate::apis::urlencode(id));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = include_credential {
        local_var_req_builder = local_var_req_builder.query(&[("include_credential", local_var_str.iter().map(|p| p.to_string()).collect::<Vec<String>>().join(","))]);
    }
    if let Some(ref local_var_user_agent) = configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(ref local_var_token) = configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetIdentityError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

/// Return a specific identity schema.
pub async fn get_identity_schema(configuration: &configuration::Configuration, id: &str) -> Result<serde_json::Value, Error<GetIdentitySchemaError>> {

    let local_var_client = &configuration.client;

    let local_var_uri_str = format!("{}/schemas/{id}", configuration.base_path, id=crate::apis::urlencode(id));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetIdentitySchemaError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

/// This endpoint is useful for:  Getting a session object with all specified expandables that exist in an administrative context.
pub async fn get_session(configuration: &configuration::Configuration, id: &str, expand: Option<Vec<String>>) -> Result<crate::models::Session, Error<GetSessionError>> {

    let local_var_client = &configuration.client;

    let local_var_uri_str = format!("{}/admin/sessions/{id}", configuration.base_path, id=crate::apis::urlencode(id));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = expand {
        local_var_req_builder = local_var_req_builder.query(&[("expand", local_var_str.iter().map(|p| p.to_string()).collect::<Vec<String>>().join(","))]);
    }
    if let Some(ref local_var_user_agent) = configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(ref local_var_token) = configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetSessionError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

/// Lists all [identities](https://www.ory.sh/docs/kratos/concepts/identity-user-model) in the system.
pub async fn list_identities(configuration: &configuration::Configuration, per_page: Option<i64>, page: Option<i64>, identifier: Option<&str>) -> Result<Vec<crate::models::Identity>, Error<ListIdentitiesError>> {

    let local_var_client = &configuration.client;

    let local_var_uri_str = format!("{}/admin/identities", configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = per_page {
        local_var_req_builder = local_var_req_builder.query(&[("per_page", local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = page {
        local_var_req_builder = local_var_req_builder.query(&[("page", local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = identifier {
        local_var_req_builder = local_var_req_builder.query(&[("identifier", local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(ref local_var_token) = configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<ListIdentitiesError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

/// Returns a list of all identity schemas currently in use.
pub async fn list_identity_schemas(configuration: &configuration::Configuration, per_page: Option<i64>, page: Option<i64>) -> Result<Vec<crate::models::IdentitySchemaContainer>, Error<ListIdentitySchemasError>> {

    let local_var_client = &configuration.client;

    let local_var_uri_str = format!("{}/schemas", configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = per_page {
        local_var_req_builder = local_var_req_builder.query(&[("per_page", local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = page {
        local_var_req_builder = local_var_req_builder.query(&[("page", local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<ListIdentitySchemasError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

/// This endpoint returns all sessions that belong to the given Identity.
pub async fn list_identity_sessions(configuration: &configuration::Configuration, id: &str, per_page: Option<i64>, page: Option<i64>, active: Option<bool>) -> Result<Vec<crate::models::Session>, Error<ListIdentitySessionsError>> {

    let local_var_client = &configuration.client;

    let local_var_uri_str = format!("{}/admin/identities/{id}/sessions", configuration.base_path, id=crate::apis::urlencode(id));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = per_page {
        local_var_req_builder = local_var_req_builder.query(&[("per_page", local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = page {
        local_var_req_builder = local_var_req_builder.query(&[("page", local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = active {
        local_var_req_builder = local_var_req_builder.query(&[("active", local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(ref local_var_token) = configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<ListIdentitySessionsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

/// Listing all sessions that exist.
pub async fn list_sessions(configuration: &configuration::Configuration, page_size: Option<i64>, page_token: Option<&str>, active: Option<bool>, expand: Option<Vec<String>>) -> Result<Vec<crate::models::Session>, Error<ListSessionsError>> {

    let local_var_client = &configuration.client;

    let local_var_uri_str = format!("{}/admin/sessions", configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = page_size {
        local_var_req_builder = local_var_req_builder.query(&[("page_size", local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = page_token {
        local_var_req_builder = local_var_req_builder.query(&[("page_token", local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = active {
        local_var_req_builder = local_var_req_builder.query(&[("active", local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = expand {
        local_var_req_builder = local_var_req_builder.query(&[("expand", local_var_str.iter().map(|p| p.to_string()).collect::<Vec<String>>().join(","))]);
    }
    if let Some(ref local_var_user_agent) = configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(ref local_var_token) = configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<ListSessionsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

/// Partially updates an [identity's](https://www.ory.sh/docs/kratos/concepts/identity-user-model) field using [JSON Patch](https://jsonpatch.com/). The fields `id`, `stateChangedAt` and `credentials` can not be updated using this method.
pub async fn patch_identity(configuration: &configuration::Configuration, id: &str, json_patch: Option<Vec<crate::models::JsonPatch>>) -> Result<crate::models::Identity, Error<PatchIdentityError>> {

    let local_var_client = &configuration.client;

    let local_var_uri_str = format!("{}/admin/identities/{id}", configuration.base_path, id=crate::apis::urlencode(id));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::PATCH, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(ref local_var_token) = configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };
    local_var_req_builder = local_var_req_builder.json(&json_patch);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<PatchIdentityError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

/// This endpoint updates an [identity](https://www.ory.sh/docs/kratos/concepts/identity-user-model). The full identity payload (except credentials) is expected. It is possible to update the identity's credentials as well.
pub async fn update_identity(configuration: &configuration::Configuration, id: &str, update_identity_body: Option<&crate::models::UpdateIdentityBody>) -> Result<crate::models::Identity, Error<UpdateIdentityError>> {

    let local_var_client = &configuration.client;

    let local_var_uri_str = format!("{}/admin/identities/{id}", configuration.base_path, id=crate::apis::urlencode(id));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::PUT, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(ref local_var_token) = configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };
    local_var_req_builder = local_var_req_builder.json(&update_identity_body);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<UpdateIdentityError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

