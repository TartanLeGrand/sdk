"""
    Ory APIs

    Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.   # noqa: E501

    The version of the OpenAPI document: v1.2.14
    Contact: support@ory.sh
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import ory_client
from ory_client.model.create_identity_body import CreateIdentityBody
globals()['CreateIdentityBody'] = CreateIdentityBody
from ory_client.model.identity_patch import IdentityPatch


class TestIdentityPatch(unittest.TestCase):
    """IdentityPatch unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testIdentityPatch(self):
        """Test IdentityPatch"""
        # FIXME: construct object with mandatory attributes with example values
        # model = IdentityPatch()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
