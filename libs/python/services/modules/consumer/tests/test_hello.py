"""Hello unit test module."""

from mod_consumer.hello import hello


def test_hello():
    """Test the hello function."""
    assert hello() == "Hello python-services-modules-consumer"
