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
from ory_client.model.courier_message_status import CourierMessageStatus
from ory_client.model.courier_message_type import CourierMessageType
from ory_client.model.message_dispatch import MessageDispatch
globals()['CourierMessageStatus'] = CourierMessageStatus
globals()['CourierMessageType'] = CourierMessageType
globals()['MessageDispatch'] = MessageDispatch
from ory_client.model.message import Message


class TestMessage(unittest.TestCase):
    """Message unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testMessage(self):
        """Test Message"""
        # FIXME: construct object with mandatory attributes with example values
        # model = Message()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
