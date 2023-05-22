=begin
#Ory APIs

#Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

The version of the OpenAPI document: v1.1.33
Contact: support@ory.sh
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.0.1

=end

require 'cgi'

module OryClient
  class ProjectApi
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end
    # Create a Project
    # Creates a new project.
    # @param [Hash] opts the optional parameters
    # @option opts [CreateProjectBody] :create_project_body 
    # @return [Project]
    def create_project(opts = {})
      data, _status_code, _headers = create_project_with_http_info(opts)
      data
    end

    # Create a Project
    # Creates a new project.
    # @param [Hash] opts the optional parameters
    # @option opts [CreateProjectBody] :create_project_body 
    # @return [Array<(Project, Integer, Hash)>] Project data, response status code and response headers
    def create_project_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: ProjectApi.create_project ...'
      end
      # resource path
      local_var_path = '/projects'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json'])
      if !content_type.nil?
          header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(opts[:'create_project_body'])

      # return_type
      return_type = opts[:debug_return_type] || 'Project'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['oryAccessToken']

      new_options = opts.merge(
        :operation => :"ProjectApi.create_project",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: ProjectApi#create_project\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Create project API token
    # Create an API token for a project.
    # @param project [String] The Project ID or Project slug
    # @param [Hash] opts the optional parameters
    # @option opts [CreateProjectApiKeyRequest] :create_project_api_key_request 
    # @return [ProjectApiKey]
    def create_project_api_key(project, opts = {})
      data, _status_code, _headers = create_project_api_key_with_http_info(project, opts)
      data
    end

    # Create project API token
    # Create an API token for a project.
    # @param project [String] The Project ID or Project slug
    # @param [Hash] opts the optional parameters
    # @option opts [CreateProjectApiKeyRequest] :create_project_api_key_request 
    # @return [Array<(ProjectApiKey, Integer, Hash)>] ProjectApiKey data, response status code and response headers
    def create_project_api_key_with_http_info(project, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: ProjectApi.create_project_api_key ...'
      end
      # verify the required parameter 'project' is set
      if @api_client.config.client_side_validation && project.nil?
        fail ArgumentError, "Missing the required parameter 'project' when calling ProjectApi.create_project_api_key"
      end
      # resource path
      local_var_path = '/projects/{project}/tokens'.sub('{' + 'project' + '}', CGI.escape(project.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json'])
      if !content_type.nil?
          header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(opts[:'create_project_api_key_request'])

      # return_type
      return_type = opts[:debug_return_type] || 'ProjectApiKey'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['oryAccessToken']

      new_options = opts.merge(
        :operation => :"ProjectApi.create_project_api_key",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: ProjectApi#create_project_api_key\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Delete project API token
    # Deletes an API token and immediately removes it.
    # @param project [String] The Project ID or Project slug
    # @param token_id [String] The Token ID
    # @param [Hash] opts the optional parameters
    # @return [nil]
    def delete_project_api_key(project, token_id, opts = {})
      delete_project_api_key_with_http_info(project, token_id, opts)
      nil
    end

    # Delete project API token
    # Deletes an API token and immediately removes it.
    # @param project [String] The Project ID or Project slug
    # @param token_id [String] The Token ID
    # @param [Hash] opts the optional parameters
    # @return [Array<(nil, Integer, Hash)>] nil, response status code and response headers
    def delete_project_api_key_with_http_info(project, token_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: ProjectApi.delete_project_api_key ...'
      end
      # verify the required parameter 'project' is set
      if @api_client.config.client_side_validation && project.nil?
        fail ArgumentError, "Missing the required parameter 'project' when calling ProjectApi.delete_project_api_key"
      end
      # verify the required parameter 'token_id' is set
      if @api_client.config.client_side_validation && token_id.nil?
        fail ArgumentError, "Missing the required parameter 'token_id' when calling ProjectApi.delete_project_api_key"
      end
      # resource path
      local_var_path = '/projects/{project}/tokens/{token_id}'.sub('{' + 'project' + '}', CGI.escape(project.to_s)).sub('{' + 'token_id' + '}', CGI.escape(token_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type]

      # auth_names
      auth_names = opts[:debug_auth_names] || ['oryAccessToken']

      new_options = opts.merge(
        :operation => :"ProjectApi.delete_project_api_key",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:DELETE, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: ProjectApi#delete_project_api_key\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Returns the Ory Network Project selected in the Ory Network Console
    # Use this API to get your active project in the Ory Network Console UI.
    # @param [Hash] opts the optional parameters
    # @return [ActiveProjectInConsole]
    def get_active_project_in_console(opts = {})
      data, _status_code, _headers = get_active_project_in_console_with_http_info(opts)
      data
    end

    # Returns the Ory Network Project selected in the Ory Network Console
    # Use this API to get your active project in the Ory Network Console UI.
    # @param [Hash] opts the optional parameters
    # @return [Array<(ActiveProjectInConsole, Integer, Hash)>] ActiveProjectInConsole data, response status code and response headers
    def get_active_project_in_console_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: ProjectApi.get_active_project_in_console ...'
      end
      # resource path
      local_var_path = '/console/active/project'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'ActiveProjectInConsole'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['oryAccessToken']

      new_options = opts.merge(
        :operation => :"ProjectApi.get_active_project_in_console",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: ProjectApi#get_active_project_in_console\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Get a Project
    # Get a projects you have access to by its ID.
    # @param project_id [String] Project ID  The project&#39;s ID.
    # @param [Hash] opts the optional parameters
    # @return [Project]
    def get_project(project_id, opts = {})
      data, _status_code, _headers = get_project_with_http_info(project_id, opts)
      data
    end

    # Get a Project
    # Get a projects you have access to by its ID.
    # @param project_id [String] Project ID  The project&#39;s ID.
    # @param [Hash] opts the optional parameters
    # @return [Array<(Project, Integer, Hash)>] Project data, response status code and response headers
    def get_project_with_http_info(project_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: ProjectApi.get_project ...'
      end
      # verify the required parameter 'project_id' is set
      if @api_client.config.client_side_validation && project_id.nil?
        fail ArgumentError, "Missing the required parameter 'project_id' when calling ProjectApi.get_project"
      end
      # resource path
      local_var_path = '/projects/{project_id}'.sub('{' + 'project_id' + '}', CGI.escape(project_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Project'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['oryAccessToken']

      new_options = opts.merge(
        :operation => :"ProjectApi.get_project",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: ProjectApi#get_project\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Get all members associated with this project
    # This endpoint requires the user to be a member of the project with the role `OWNER` or `DEVELOPER`.
    # @param project_id [String] Project ID  The project&#39;s ID.
    # @param [Hash] opts the optional parameters
    # @return [Array<CloudAccount>]
    def get_project_members(project_id, opts = {})
      data, _status_code, _headers = get_project_members_with_http_info(project_id, opts)
      data
    end

    # Get all members associated with this project
    # This endpoint requires the user to be a member of the project with the role &#x60;OWNER&#x60; or &#x60;DEVELOPER&#x60;.
    # @param project_id [String] Project ID  The project&#39;s ID.
    # @param [Hash] opts the optional parameters
    # @return [Array<(Array<CloudAccount>, Integer, Hash)>] Array<CloudAccount> data, response status code and response headers
    def get_project_members_with_http_info(project_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: ProjectApi.get_project_members ...'
      end
      # verify the required parameter 'project_id' is set
      if @api_client.config.client_side_validation && project_id.nil?
        fail ArgumentError, "Missing the required parameter 'project_id' when calling ProjectApi.get_project_members"
      end
      # resource path
      local_var_path = '/projects/{project_id}/members'.sub('{' + 'project_id' + '}', CGI.escape(project_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Array<CloudAccount>'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['oryAccessToken']

      new_options = opts.merge(
        :operation => :"ProjectApi.get_project_members",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: ProjectApi#get_project_members\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # List a project's API Tokens
    # A list of all the project's API tokens.
    # @param project [String] The Project ID or Project slug
    # @param [Hash] opts the optional parameters
    # @return [Array<ProjectApiKey>]
    def list_project_api_keys(project, opts = {})
      data, _status_code, _headers = list_project_api_keys_with_http_info(project, opts)
      data
    end

    # List a project&#39;s API Tokens
    # A list of all the project&#39;s API tokens.
    # @param project [String] The Project ID or Project slug
    # @param [Hash] opts the optional parameters
    # @return [Array<(Array<ProjectApiKey>, Integer, Hash)>] Array<ProjectApiKey> data, response status code and response headers
    def list_project_api_keys_with_http_info(project, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: ProjectApi.list_project_api_keys ...'
      end
      # verify the required parameter 'project' is set
      if @api_client.config.client_side_validation && project.nil?
        fail ArgumentError, "Missing the required parameter 'project' when calling ProjectApi.list_project_api_keys"
      end
      # resource path
      local_var_path = '/projects/{project}/tokens'.sub('{' + 'project' + '}', CGI.escape(project.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Array<ProjectApiKey>'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['oryAccessToken']

      new_options = opts.merge(
        :operation => :"ProjectApi.list_project_api_keys",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: ProjectApi#list_project_api_keys\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # List All Projects
    # Lists all projects you have access to.
    # @param [Hash] opts the optional parameters
    # @return [Array<ProjectMetadata>]
    def list_projects(opts = {})
      data, _status_code, _headers = list_projects_with_http_info(opts)
      data
    end

    # List All Projects
    # Lists all projects you have access to.
    # @param [Hash] opts the optional parameters
    # @return [Array<(Array<ProjectMetadata>, Integer, Hash)>] Array<ProjectMetadata> data, response status code and response headers
    def list_projects_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: ProjectApi.list_projects ...'
      end
      # resource path
      local_var_path = '/projects'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Array<ProjectMetadata>'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['oryAccessToken']

      new_options = opts.merge(
        :operation => :"ProjectApi.list_projects",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: ProjectApi#list_projects\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Patch an Ory Network Project Configuration
    # Deprecated: Use the `patchProjectWithRevision` endpoint instead to specify the exact revision the patch was generated for.  This endpoints allows you to patch individual Ory Network project configuration keys for Ory's services (identity, permission, ...). The configuration format is fully compatible with the open source projects for the respective services (e.g. Ory Kratos for Identity, Ory Keto for Permissions).  This endpoint expects the `version` key to be set in the payload. If it is unset, it will try to import the config as if it is from the most recent version.  If you have an older version of a configuration, you should set the version key in the payload!  While this endpoint is able to process all configuration items related to features (e.g. password reset), it does not support operational configuration items (e.g. port, tracing, logging) otherwise available in the open source.  For configuration items that can not be translated to the Ory Network, this endpoint will return a list of warnings to help you understand which parts of your config could not be processed.
    # @param project_id [String] Project ID  The project&#39;s ID.
    # @param [Hash] opts the optional parameters
    # @option opts [Array<JsonPatch>] :json_patch 
    # @return [SuccessfulProjectUpdate]
    def patch_project(project_id, opts = {})
      data, _status_code, _headers = patch_project_with_http_info(project_id, opts)
      data
    end

    # Patch an Ory Network Project Configuration
    # Deprecated: Use the &#x60;patchProjectWithRevision&#x60; endpoint instead to specify the exact revision the patch was generated for.  This endpoints allows you to patch individual Ory Network project configuration keys for Ory&#39;s services (identity, permission, ...). The configuration format is fully compatible with the open source projects for the respective services (e.g. Ory Kratos for Identity, Ory Keto for Permissions).  This endpoint expects the &#x60;version&#x60; key to be set in the payload. If it is unset, it will try to import the config as if it is from the most recent version.  If you have an older version of a configuration, you should set the version key in the payload!  While this endpoint is able to process all configuration items related to features (e.g. password reset), it does not support operational configuration items (e.g. port, tracing, logging) otherwise available in the open source.  For configuration items that can not be translated to the Ory Network, this endpoint will return a list of warnings to help you understand which parts of your config could not be processed.
    # @param project_id [String] Project ID  The project&#39;s ID.
    # @param [Hash] opts the optional parameters
    # @option opts [Array<JsonPatch>] :json_patch 
    # @return [Array<(SuccessfulProjectUpdate, Integer, Hash)>] SuccessfulProjectUpdate data, response status code and response headers
    def patch_project_with_http_info(project_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: ProjectApi.patch_project ...'
      end
      # verify the required parameter 'project_id' is set
      if @api_client.config.client_side_validation && project_id.nil?
        fail ArgumentError, "Missing the required parameter 'project_id' when calling ProjectApi.patch_project"
      end
      # resource path
      local_var_path = '/projects/{project_id}'.sub('{' + 'project_id' + '}', CGI.escape(project_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json'])
      if !content_type.nil?
          header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(opts[:'json_patch'])

      # return_type
      return_type = opts[:debug_return_type] || 'SuccessfulProjectUpdate'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['oryAccessToken']

      new_options = opts.merge(
        :operation => :"ProjectApi.patch_project",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:PATCH, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: ProjectApi#patch_project\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Irrecoverably purge a project
    # !! Use with extreme caution !!  Using this API endpoint you can purge (completely delete) a project and its data. This action can not be undone and will delete ALL your data.  !! Use with extreme caution !!
    # @param project_id [String] Project ID  The project&#39;s ID.
    # @param [Hash] opts the optional parameters
    # @return [nil]
    def purge_project(project_id, opts = {})
      purge_project_with_http_info(project_id, opts)
      nil
    end

    # Irrecoverably purge a project
    # !! Use with extreme caution !!  Using this API endpoint you can purge (completely delete) a project and its data. This action can not be undone and will delete ALL your data.  !! Use with extreme caution !!
    # @param project_id [String] Project ID  The project&#39;s ID.
    # @param [Hash] opts the optional parameters
    # @return [Array<(nil, Integer, Hash)>] nil, response status code and response headers
    def purge_project_with_http_info(project_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: ProjectApi.purge_project ...'
      end
      # verify the required parameter 'project_id' is set
      if @api_client.config.client_side_validation && project_id.nil?
        fail ArgumentError, "Missing the required parameter 'project_id' when calling ProjectApi.purge_project"
      end
      # resource path
      local_var_path = '/projects/{project_id}'.sub('{' + 'project_id' + '}', CGI.escape(project_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type]

      # auth_names
      auth_names = opts[:debug_auth_names] || ['oryAccessToken']

      new_options = opts.merge(
        :operation => :"ProjectApi.purge_project",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:DELETE, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: ProjectApi#purge_project\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Remove a member associated with this project
    # This also sets their invite status to `REMOVED`. This endpoint requires the user to be a member of the project with the role `OWNER`.
    # @param project_id [String] Project ID  The project&#39;s ID.
    # @param member_id [String] Member ID
    # @param [Hash] opts the optional parameters
    # @return [nil]
    def remove_project_member(project_id, member_id, opts = {})
      remove_project_member_with_http_info(project_id, member_id, opts)
      nil
    end

    # Remove a member associated with this project
    # This also sets their invite status to &#x60;REMOVED&#x60;. This endpoint requires the user to be a member of the project with the role &#x60;OWNER&#x60;.
    # @param project_id [String] Project ID  The project&#39;s ID.
    # @param member_id [String] Member ID
    # @param [Hash] opts the optional parameters
    # @return [Array<(nil, Integer, Hash)>] nil, response status code and response headers
    def remove_project_member_with_http_info(project_id, member_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: ProjectApi.remove_project_member ...'
      end
      # verify the required parameter 'project_id' is set
      if @api_client.config.client_side_validation && project_id.nil?
        fail ArgumentError, "Missing the required parameter 'project_id' when calling ProjectApi.remove_project_member"
      end
      # verify the required parameter 'member_id' is set
      if @api_client.config.client_side_validation && member_id.nil?
        fail ArgumentError, "Missing the required parameter 'member_id' when calling ProjectApi.remove_project_member"
      end
      # resource path
      local_var_path = '/projects/{project_id}/members/{member_id}'.sub('{' + 'project_id' + '}', CGI.escape(project_id.to_s)).sub('{' + 'member_id' + '}', CGI.escape(member_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type]

      # auth_names
      auth_names = opts[:debug_auth_names] || ['oryAccessToken']

      new_options = opts.merge(
        :operation => :"ProjectApi.remove_project_member",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:DELETE, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: ProjectApi#remove_project_member\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Sets the Ory Network Project active in the Ory Network Console
    # Use this API to set your active project in the Ory Network Console UI.
    # @param [Hash] opts the optional parameters
    # @option opts [SetActiveProjectInConsoleBody] :set_active_project_in_console_body 
    # @return [nil]
    def set_active_project_in_console(opts = {})
      set_active_project_in_console_with_http_info(opts)
      nil
    end

    # Sets the Ory Network Project active in the Ory Network Console
    # Use this API to set your active project in the Ory Network Console UI.
    # @param [Hash] opts the optional parameters
    # @option opts [SetActiveProjectInConsoleBody] :set_active_project_in_console_body 
    # @return [Array<(nil, Integer, Hash)>] nil, response status code and response headers
    def set_active_project_in_console_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: ProjectApi.set_active_project_in_console ...'
      end
      # resource path
      local_var_path = '/console/active/project'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json'])
      if !content_type.nil?
          header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(opts[:'set_active_project_in_console_body'])

      # return_type
      return_type = opts[:debug_return_type]

      # auth_names
      auth_names = opts[:debug_auth_names] || ['oryAccessToken']

      new_options = opts.merge(
        :operation => :"ProjectApi.set_active_project_in_console",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:PUT, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: ProjectApi#set_active_project_in_console\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Update an Ory Network Project Configuration
    # This endpoints allows you to update the Ory Network project configuration for individual services (identity, permission, ...). The configuration is fully compatible with the open source projects for the respective services (e.g. Ory Kratos for Identity, Ory Keto for Permissions).  This endpoint expects the `version` key to be set in the payload. If it is unset, it will try to import the config as if it is from the most recent version.  If you have an older version of a configuration, you should set the version key in the payload!  While this endpoint is able to process all configuration items related to features (e.g. password reset), it does not support operational configuration items (e.g. port, tracing, logging) otherwise available in the open source.  For configuration items that can not be translated to the Ory Network, this endpoint will return a list of warnings to help you understand which parts of your config could not be processed.  Be aware that updating any service's configuration will completely override your current configuration for that service!
    # @param project_id [String] Project ID  The project&#39;s ID.
    # @param [Hash] opts the optional parameters
    # @option opts [SetProject] :set_project 
    # @return [SuccessfulProjectUpdate]
    def set_project(project_id, opts = {})
      data, _status_code, _headers = set_project_with_http_info(project_id, opts)
      data
    end

    # Update an Ory Network Project Configuration
    # This endpoints allows you to update the Ory Network project configuration for individual services (identity, permission, ...). The configuration is fully compatible with the open source projects for the respective services (e.g. Ory Kratos for Identity, Ory Keto for Permissions).  This endpoint expects the &#x60;version&#x60; key to be set in the payload. If it is unset, it will try to import the config as if it is from the most recent version.  If you have an older version of a configuration, you should set the version key in the payload!  While this endpoint is able to process all configuration items related to features (e.g. password reset), it does not support operational configuration items (e.g. port, tracing, logging) otherwise available in the open source.  For configuration items that can not be translated to the Ory Network, this endpoint will return a list of warnings to help you understand which parts of your config could not be processed.  Be aware that updating any service&#39;s configuration will completely override your current configuration for that service!
    # @param project_id [String] Project ID  The project&#39;s ID.
    # @param [Hash] opts the optional parameters
    # @option opts [SetProject] :set_project 
    # @return [Array<(SuccessfulProjectUpdate, Integer, Hash)>] SuccessfulProjectUpdate data, response status code and response headers
    def set_project_with_http_info(project_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: ProjectApi.set_project ...'
      end
      # verify the required parameter 'project_id' is set
      if @api_client.config.client_side_validation && project_id.nil?
        fail ArgumentError, "Missing the required parameter 'project_id' when calling ProjectApi.set_project"
      end
      # resource path
      local_var_path = '/projects/{project_id}'.sub('{' + 'project_id' + '}', CGI.escape(project_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json'])
      if !content_type.nil?
          header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(opts[:'set_project'])

      # return_type
      return_type = opts[:debug_return_type] || 'SuccessfulProjectUpdate'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['oryAccessToken']

      new_options = opts.merge(
        :operation => :"ProjectApi.set_project",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:PUT, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: ProjectApi#set_project\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end
  end
end
