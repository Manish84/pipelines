# Copyright 2022 The Kubeflow Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import abc
from typing import Optional

from kubernetes.client import configuration


class TokenCredentialsBase(abc.ABC):

    @abc.abstractmethod
    def refresh_api_key_hook(self, config: configuration.Configuration) -> None:
        """Refreshes the api key.

        This is a helper function for registering token refresh with swagger
        generated clients.

        All classes that inherit from TokenCredentialsBase must implement this
        method to refresh the credentials.

        Args:
            config (kubernetes.client.configuration.Configuration):
                The configuration object that the client uses.

                The Configuration object of the kubernetes client's is the same
                with kfp_server_api.configuration.Configuration.
        """
        raise NotImplementedError()


def read_token_from_file(path: Optional[str] = None) -> str:
    """Reads a token found in some file.

    Args:
        path (Optional[str], optional): Path of file containing auth token. Defaults to None.

    Returns:
        str: The token.
    """
    token = None
    with open(path) as f:
        token = f.read().strip()
    return token
