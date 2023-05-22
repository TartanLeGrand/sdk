"""
    Ory APIs

    Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.   # noqa: E501

    The version of the OpenAPI document: v1.1.33
    Contact: support@ory.sh
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import ory_client
from ory_client.model.ui_container import UiContainer
from ory_client.model.verification_flow_state import VerificationFlowState
globals()['UiContainer'] = UiContainer
globals()['VerificationFlowState'] = VerificationFlowState
from ory_client.model.verification_flow import VerificationFlow


class TestVerificationFlow(unittest.TestCase):
    """VerificationFlow unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testVerificationFlow(self):
        """Test VerificationFlow"""
        # FIXME: construct object with mandatory attributes with example values
        # model = VerificationFlow()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
