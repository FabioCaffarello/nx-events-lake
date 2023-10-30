import os
import unittest
from pydotenv.loader import DotEnvLoader
from pathlib import Path

def _get_env_path():
    return Path(__file__).parent.joinpath("reference_files")

class TestDotEnvLoader(unittest.TestCase):

    def setUp(self):
        path = _get_env_path()
        environment= "test"
        self.env_loader = DotEnvLoader(environment, path)

    def test_load(self):
        self.env_loader.load()
        self.assertEqual(os.getenv('SECRET_KEY'), 'your_secret_key')

    def test_get_variable_existing(self):
        value = self.env_loader.get_variable('SECRET_KEY')
        self.assertEqual(value, 'your_secret_key')

    def test_get_variable_nonexistent(self):
        value = self.env_loader.get_variable('NON_EXISTENT_KEY')
        self.assertEqual(value, "")

if __name__ == '__main__':
    unittest.main()
