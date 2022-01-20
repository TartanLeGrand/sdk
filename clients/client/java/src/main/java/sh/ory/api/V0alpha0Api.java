/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.53
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.api;

import sh.ory.ApiCallback;
import sh.ory.ApiClient;
import sh.ory.ApiException;
import sh.ory.ApiResponse;
import sh.ory.Configuration;
import sh.ory.Pair;
import sh.ory.ProgressRequestBody;
import sh.ory.ProgressResponseBody;

import com.google.gson.reflect.TypeToken;

import java.io.IOException;


import sh.ory.model.GenericError;
import sh.ory.model.Project;
import sh.ory.model.ProjectPatch;

import java.lang.reflect.Type;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class V0alpha0Api {
    private ApiClient localVarApiClient;

    public V0alpha0Api() {
        this(Configuration.getDefaultApiClient());
    }

    public V0alpha0Api(ApiClient apiClient) {
        this.localVarApiClient = apiClient;
    }

    public ApiClient getApiClient() {
        return localVarApiClient;
    }

    public void setApiClient(ApiClient apiClient) {
        this.localVarApiClient = apiClient;
    }

    /**
     * Build call for createProject
     * @param projectPatch  (optional)
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 201 </td><td> project </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 0 </td><td> genericError </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call createProjectCall(ProjectPatch projectPatch, final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = projectPatch;

        // create path and map variables
        String localVarPath = "/backoffice/public/projects";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            "application/json"
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] { "oryAccessToken" };
        return localVarApiClient.buildCall(localVarPath, "POST", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call createProjectValidateBeforeCall(ProjectPatch projectPatch, final ApiCallback _callback) throws ApiException {
        

        okhttp3.Call localVarCall = createProjectCall(projectPatch, _callback);
        return localVarCall;

    }

    /**
     * Create a Project
     * Creates a new project.
     * @param projectPatch  (optional)
     * @return Project
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 201 </td><td> project </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 0 </td><td> genericError </td><td>  -  </td></tr>
     </table>
     */
    public Project createProject(ProjectPatch projectPatch) throws ApiException {
        ApiResponse<Project> localVarResp = createProjectWithHttpInfo(projectPatch);
        return localVarResp.getData();
    }

    /**
     * Create a Project
     * Creates a new project.
     * @param projectPatch  (optional)
     * @return ApiResponse&lt;Project&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 201 </td><td> project </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 0 </td><td> genericError </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<Project> createProjectWithHttpInfo(ProjectPatch projectPatch) throws ApiException {
        okhttp3.Call localVarCall = createProjectValidateBeforeCall(projectPatch, null);
        Type localVarReturnType = new TypeToken<Project>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     * Create a Project (asynchronously)
     * Creates a new project.
     * @param projectPatch  (optional)
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 201 </td><td> project </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 0 </td><td> genericError </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call createProjectAsync(ProjectPatch projectPatch, final ApiCallback<Project> _callback) throws ApiException {

        okhttp3.Call localVarCall = createProjectValidateBeforeCall(projectPatch, _callback);
        Type localVarReturnType = new TypeToken<Project>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
    /**
     * Build call for getProject
     * @param projectId Project ID  The project&#39;s ID. (required)
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> project </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 0 </td><td> genericError </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call getProjectCall(String projectId, final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = null;

        // create path and map variables
        String localVarPath = "/backoffice/public/projects/{project_id}"
            .replaceAll("\\{" + "project_id" + "\\}", localVarApiClient.escapeString(projectId.toString()));

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] { "oryAccessToken" };
        return localVarApiClient.buildCall(localVarPath, "GET", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call getProjectValidateBeforeCall(String projectId, final ApiCallback _callback) throws ApiException {
        
        // verify the required parameter 'projectId' is set
        if (projectId == null) {
            throw new ApiException("Missing the required parameter 'projectId' when calling getProject(Async)");
        }
        

        okhttp3.Call localVarCall = getProjectCall(projectId, _callback);
        return localVarCall;

    }

    /**
     * Get a Project
     * Get a projects you have access to by its ID.
     * @param projectId Project ID  The project&#39;s ID. (required)
     * @return Project
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> project </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 0 </td><td> genericError </td><td>  -  </td></tr>
     </table>
     */
    public Project getProject(String projectId) throws ApiException {
        ApiResponse<Project> localVarResp = getProjectWithHttpInfo(projectId);
        return localVarResp.getData();
    }

    /**
     * Get a Project
     * Get a projects you have access to by its ID.
     * @param projectId Project ID  The project&#39;s ID. (required)
     * @return ApiResponse&lt;Project&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> project </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 0 </td><td> genericError </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<Project> getProjectWithHttpInfo(String projectId) throws ApiException {
        okhttp3.Call localVarCall = getProjectValidateBeforeCall(projectId, null);
        Type localVarReturnType = new TypeToken<Project>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     * Get a Project (asynchronously)
     * Get a projects you have access to by its ID.
     * @param projectId Project ID  The project&#39;s ID. (required)
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> project </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 0 </td><td> genericError </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call getProjectAsync(String projectId, final ApiCallback<Project> _callback) throws ApiException {

        okhttp3.Call localVarCall = getProjectValidateBeforeCall(projectId, _callback);
        Type localVarReturnType = new TypeToken<Project>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
    /**
     * Build call for listProjects
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> projects </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 0 </td><td> genericError </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call listProjectsCall(final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = null;

        // create path and map variables
        String localVarPath = "/backoffice/public/projects";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] { "oryAccessToken" };
        return localVarApiClient.buildCall(localVarPath, "GET", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call listProjectsValidateBeforeCall(final ApiCallback _callback) throws ApiException {
        

        okhttp3.Call localVarCall = listProjectsCall(_callback);
        return localVarCall;

    }

    /**
     * List All Projects
     * Lists all projects you have access to.
     * @return List&lt;Project&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> projects </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 0 </td><td> genericError </td><td>  -  </td></tr>
     </table>
     */
    public List<Project> listProjects() throws ApiException {
        ApiResponse<List<Project>> localVarResp = listProjectsWithHttpInfo();
        return localVarResp.getData();
    }

    /**
     * List All Projects
     * Lists all projects you have access to.
     * @return ApiResponse&lt;List&lt;Project&gt;&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> projects </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 0 </td><td> genericError </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<List<Project>> listProjectsWithHttpInfo() throws ApiException {
        okhttp3.Call localVarCall = listProjectsValidateBeforeCall(null);
        Type localVarReturnType = new TypeToken<List<Project>>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     * List All Projects (asynchronously)
     * Lists all projects you have access to.
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> projects </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 0 </td><td> genericError </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call listProjectsAsync(final ApiCallback<List<Project>> _callback) throws ApiException {

        okhttp3.Call localVarCall = listProjectsValidateBeforeCall(_callback);
        Type localVarReturnType = new TypeToken<List<Project>>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
    /**
     * Build call for updateProject
     * @param projectId Project ID  The project&#39;s ID. (required)
     * @param projectPatch  (optional)
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> project </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 0 </td><td> genericError </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call updateProjectCall(String projectId, ProjectPatch projectPatch, final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = projectPatch;

        // create path and map variables
        String localVarPath = "/backoffice/public/projects/{project_id}"
            .replaceAll("\\{" + "project_id" + "\\}", localVarApiClient.escapeString(projectId.toString()));

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            "application/json"
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] { "oryAccessToken" };
        return localVarApiClient.buildCall(localVarPath, "PUT", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call updateProjectValidateBeforeCall(String projectId, ProjectPatch projectPatch, final ApiCallback _callback) throws ApiException {
        
        // verify the required parameter 'projectId' is set
        if (projectId == null) {
            throw new ApiException("Missing the required parameter 'projectId' when calling updateProject(Async)");
        }
        

        okhttp3.Call localVarCall = updateProjectCall(projectId, projectPatch, _callback);
        return localVarCall;

    }

    /**
     * Update a Project
     * Creates a new configuration revision for a project.
     * @param projectId Project ID  The project&#39;s ID. (required)
     * @param projectPatch  (optional)
     * @return Project
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> project </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 0 </td><td> genericError </td><td>  -  </td></tr>
     </table>
     */
    public Project updateProject(String projectId, ProjectPatch projectPatch) throws ApiException {
        ApiResponse<Project> localVarResp = updateProjectWithHttpInfo(projectId, projectPatch);
        return localVarResp.getData();
    }

    /**
     * Update a Project
     * Creates a new configuration revision for a project.
     * @param projectId Project ID  The project&#39;s ID. (required)
     * @param projectPatch  (optional)
     * @return ApiResponse&lt;Project&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> project </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 0 </td><td> genericError </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<Project> updateProjectWithHttpInfo(String projectId, ProjectPatch projectPatch) throws ApiException {
        okhttp3.Call localVarCall = updateProjectValidateBeforeCall(projectId, projectPatch, null);
        Type localVarReturnType = new TypeToken<Project>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     * Update a Project (asynchronously)
     * Creates a new configuration revision for a project.
     * @param projectId Project ID  The project&#39;s ID. (required)
     * @param projectPatch  (optional)
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> project </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> genericError </td><td>  -  </td></tr>
        <tr><td> 0 </td><td> genericError </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call updateProjectAsync(String projectId, ProjectPatch projectPatch, final ApiCallback<Project> _callback) throws ApiException {

        okhttp3.Call localVarCall = updateProjectValidateBeforeCall(projectId, projectPatch, _callback);
        Type localVarReturnType = new TypeToken<Project>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
}
