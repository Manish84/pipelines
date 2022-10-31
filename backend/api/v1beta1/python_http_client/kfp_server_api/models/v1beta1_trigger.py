# coding: utf-8

"""
    Kubeflow Pipelines API

    This file contains REST API specification for Kubeflow Pipelines. The file is autogenerated from the swagger definition.

    Contact: kubeflow-pipelines@google.com
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from kfp_server_api.configuration import Configuration


class V1beta1Trigger(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'cron_schedule': 'V1beta1CronSchedule',
        'periodic_schedule': 'V1beta1PeriodicSchedule'
    }

    attribute_map = {
        'cron_schedule': 'cron_schedule',
        'periodic_schedule': 'periodic_schedule'
    }

    def __init__(self, cron_schedule=None, periodic_schedule=None, local_vars_configuration=None):  # noqa: E501
        """V1beta1Trigger - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._cron_schedule = None
        self._periodic_schedule = None
        self.discriminator = None

        if cron_schedule is not None:
            self.cron_schedule = cron_schedule
        if periodic_schedule is not None:
            self.periodic_schedule = periodic_schedule

    @property
    def cron_schedule(self):
        """Gets the cron_schedule of this V1beta1Trigger.  # noqa: E501


        :return: The cron_schedule of this V1beta1Trigger.  # noqa: E501
        :rtype: V1beta1CronSchedule
        """
        return self._cron_schedule

    @cron_schedule.setter
    def cron_schedule(self, cron_schedule):
        """Sets the cron_schedule of this V1beta1Trigger.


        :param cron_schedule: The cron_schedule of this V1beta1Trigger.  # noqa: E501
        :type cron_schedule: V1beta1CronSchedule
        """

        self._cron_schedule = cron_schedule

    @property
    def periodic_schedule(self):
        """Gets the periodic_schedule of this V1beta1Trigger.  # noqa: E501


        :return: The periodic_schedule of this V1beta1Trigger.  # noqa: E501
        :rtype: V1beta1PeriodicSchedule
        """
        return self._periodic_schedule

    @periodic_schedule.setter
    def periodic_schedule(self, periodic_schedule):
        """Sets the periodic_schedule of this V1beta1Trigger.


        :param periodic_schedule: The periodic_schedule of this V1beta1Trigger.  # noqa: E501
        :type periodic_schedule: V1beta1PeriodicSchedule
        """

        self._periodic_schedule = periodic_schedule

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1beta1Trigger):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1beta1Trigger):
            return True

        return self.to_dict() != other.to_dict()
