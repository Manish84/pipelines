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
import os
import tempfile
import unittest

from kfp.client import token_credentials_base
from kubernetes.client import configuration


class TestTokenCredentialsBase(unittest.TestCase):

    def test_cannot_instantiate_abc(self):

        class MyTokenGetter(token_credentials_base.TokenCredentialsBase):

            def something(self) -> None:
                pass

        with self.assertRaisesRegex(TypeError,
                                    r"Can't instantiate abstract class"):
            MyTokenGetter()

    def test_can_instantiate_with_refresh_api_key_hook(self):

        class MyTokenGetter(token_credentials_base.TokenCredentialsBase):

            def refresh_api_key_hook(
                    self, config: configuration.Configuration) -> None:
                pass

        # does not raise Exception
        MyTokenGetter()


class TestReadTokenFromFile(unittest.TestCase):

    def test_none_path(self):
        with self.assertRaisesRegex(TypeError, r'^expected .* not NoneType'):
            token_credentials_base.read_token_from_file()

    def test_str_path(self):
        with tempfile.TemporaryDirectory() as tempdir:
            token_path = os.path.join(tempdir, 'token.txt')
            with open(token_path, 'w') as f:
                f.write('my_token')

            self.assertEqual(
                token_credentials_base.read_token_from_file(token_path),
                'my_token')

    def test_str_path_extra_spaces_in_token(self):
        with tempfile.TemporaryDirectory() as tempdir:
            token_path = os.path.join(tempdir, 'token.txt')
            with open(token_path, 'w') as f:
                f.write(' my_token ')

            self.assertEqual(
                token_credentials_base.read_token_from_file(token_path),
                'my_token')
