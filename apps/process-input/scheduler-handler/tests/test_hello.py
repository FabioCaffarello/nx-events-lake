"""Hello unit test module."""

from scheduler_handler.hello import hello


def test_hello():
    """Test the hello function."""
    assert hello() == "Hello process-input-scheduler-handler"
