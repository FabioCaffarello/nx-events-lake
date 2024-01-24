import re
from typing import Dict
from pylog.log import setup_logging

logger = setup_logging(__name__)


def get_target_input_data_by_regex_pattern(html: str, target_pattern: Dict[str, any]) -> Dict[str, str]:
    """
    Extracts target input data from HTML using specified regex patterns.

    Args:
        html (str): The HTML content to search for patterns.
        target_pattern (Dict[str, any]): A dictionary containing variable names as keys
            and regex patterns as values.

    Returns:
        Dict[str, str]: A dictionary mapping variable names to extracted data.

    Raises:
        None
            If the data for any variable is not found in the HTML.
    """
    result_search = dict()
    for var_name, _target_pattern in target_pattern.items():
        fetch_pattern = re.findall(_target_pattern, html)
        if fetch_pattern == []:
            return None
        result_search[var_name] = fetch_pattern[0].decode()
    return result_search
