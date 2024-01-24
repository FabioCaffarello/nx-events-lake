"""Hello unit test module."""

from mod_pyspark_domain_core.hello import hello


def test_hello():
    """Test the hello function."""
    assert hello() == "Hello python-services-modules-treatment-pyspark-core"
