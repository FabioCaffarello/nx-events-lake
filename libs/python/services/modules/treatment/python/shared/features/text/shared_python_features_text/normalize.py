import unicodedata

def apply_utf8_normalization_without_accents(text: str) -> str:
    """
    Applies UTF-8 normalization to the input text, removing accents.

    This function uses Unicode normalization (NFKD) to decompose characters
    with diacritic marks and then encodes the resulting string to ASCII, ignoring
    any characters that cannot be represented in ASCII. Finally, the string is
    decoded back to UTF-8.

    Args:
        text (str): The input text to be normalized.

    Returns:
        str: The normalized text with accents removed.

    Example:
        >>> apply_utf8_normalization_without_accents("héllö")
        'hello'
    """
    return str(unicodedata.normalize('NFKD', text).encode('ASCII', 'ignore').decode('utf-8'))
