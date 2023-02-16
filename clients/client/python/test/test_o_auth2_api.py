"""
    Ory APIs

    Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.   # noqa: E501

    The version of the OpenAPI document: v1.1.11
    Contact: support@ory.sh
    Generated by: https://openapi-generator.tech
"""


import unittest

import ory_client
from ory_client.api.o_auth2_api import OAuth2Api  # noqa: E501


class TestOAuth2Api(unittest.TestCase):
    """OAuth2Api unit test stubs"""

    def setUp(self):
        self.api = OAuth2Api()  # noqa: E501

    def tearDown(self):
        pass

    def test_accept_o_auth2_consent_request(self):
        """Test case for accept_o_auth2_consent_request

        Accept OAuth 2.0 Consent Request  # noqa: E501
        """
        pass

    def test_accept_o_auth2_login_request(self):
        """Test case for accept_o_auth2_login_request

        Accept OAuth 2.0 Login Request  # noqa: E501
        """
        pass

    def test_accept_o_auth2_logout_request(self):
        """Test case for accept_o_auth2_logout_request

        Accept OAuth 2.0 Session Logout Request  # noqa: E501
        """
        pass

    def test_create_o_auth2_client(self):
        """Test case for create_o_auth2_client

        Create OAuth 2.0 Client  # noqa: E501
        """
        pass

    def test_delete_o_auth2_client(self):
        """Test case for delete_o_auth2_client

        Delete OAuth 2.0 Client  # noqa: E501
        """
        pass

    def test_delete_o_auth2_token(self):
        """Test case for delete_o_auth2_token

        Delete OAuth 2.0 Access Tokens from specific OAuth 2.0 Client  # noqa: E501
        """
        pass

    def test_delete_trusted_o_auth2_jwt_grant_issuer(self):
        """Test case for delete_trusted_o_auth2_jwt_grant_issuer

        Delete Trusted OAuth2 JWT Bearer Grant Type Issuer  # noqa: E501
        """
        pass

    def test_get_o_auth2_client(self):
        """Test case for get_o_auth2_client

        Get an OAuth 2.0 Client  # noqa: E501
        """
        pass

    def test_get_o_auth2_consent_request(self):
        """Test case for get_o_auth2_consent_request

        Get OAuth 2.0 Consent Request  # noqa: E501
        """
        pass

    def test_get_o_auth2_login_request(self):
        """Test case for get_o_auth2_login_request

        Get OAuth 2.0 Login Request  # noqa: E501
        """
        pass

    def test_get_o_auth2_logout_request(self):
        """Test case for get_o_auth2_logout_request

        Get OAuth 2.0 Session Logout Request  # noqa: E501
        """
        pass

    def test_get_trusted_o_auth2_jwt_grant_issuer(self):
        """Test case for get_trusted_o_auth2_jwt_grant_issuer

        Get Trusted OAuth2 JWT Bearer Grant Type Issuer  # noqa: E501
        """
        pass

    def test_introspect_o_auth2_token(self):
        """Test case for introspect_o_auth2_token

        Introspect OAuth2 Access and Refresh Tokens  # noqa: E501
        """
        pass

    def test_list_o_auth2_clients(self):
        """Test case for list_o_auth2_clients

        List OAuth 2.0 Clients  # noqa: E501
        """
        pass

    def test_list_o_auth2_consent_sessions(self):
        """Test case for list_o_auth2_consent_sessions

        List OAuth 2.0 Consent Sessions of a Subject  # noqa: E501
        """
        pass

    def test_list_trusted_o_auth2_jwt_grant_issuers(self):
        """Test case for list_trusted_o_auth2_jwt_grant_issuers

        List Trusted OAuth2 JWT Bearer Grant Type Issuers  # noqa: E501
        """
        pass

    def test_o_auth2_authorize(self):
        """Test case for o_auth2_authorize

        OAuth 2.0 Authorize Endpoint  # noqa: E501
        """
        pass

    def test_oauth2_token_exchange(self):
        """Test case for oauth2_token_exchange

        The OAuth 2.0 Token Endpoint  # noqa: E501
        """
        pass

    def test_patch_o_auth2_client(self):
        """Test case for patch_o_auth2_client

        Patch OAuth 2.0 Client  # noqa: E501
        """
        pass

    def test_reject_o_auth2_consent_request(self):
        """Test case for reject_o_auth2_consent_request

        Reject OAuth 2.0 Consent Request  # noqa: E501
        """
        pass

    def test_reject_o_auth2_login_request(self):
        """Test case for reject_o_auth2_login_request

        Reject OAuth 2.0 Login Request  # noqa: E501
        """
        pass

    def test_reject_o_auth2_logout_request(self):
        """Test case for reject_o_auth2_logout_request

        Reject OAuth 2.0 Session Logout Request  # noqa: E501
        """
        pass

    def test_revoke_o_auth2_consent_sessions(self):
        """Test case for revoke_o_auth2_consent_sessions

        Revoke OAuth 2.0 Consent Sessions of a Subject  # noqa: E501
        """
        pass

    def test_revoke_o_auth2_login_sessions(self):
        """Test case for revoke_o_auth2_login_sessions

        Revokes All OAuth 2.0 Login Sessions of a Subject  # noqa: E501
        """
        pass

    def test_revoke_o_auth2_token(self):
        """Test case for revoke_o_auth2_token

        Revoke OAuth 2.0 Access or Refresh Token  # noqa: E501
        """
        pass

    def test_set_o_auth2_client(self):
        """Test case for set_o_auth2_client

        Set OAuth 2.0 Client  # noqa: E501
        """
        pass

    def test_set_o_auth2_client_lifespans(self):
        """Test case for set_o_auth2_client_lifespans

        Set OAuth2 Client Token Lifespans  # noqa: E501
        """
        pass

    def test_trust_o_auth2_jwt_grant_issuer(self):
        """Test case for trust_o_auth2_jwt_grant_issuer

        Trust OAuth2 JWT Bearer Grant Type Issuer  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
