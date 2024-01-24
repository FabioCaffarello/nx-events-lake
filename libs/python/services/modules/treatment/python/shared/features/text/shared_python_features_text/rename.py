from shared_python_features_text.normalize import apply_utf8_normalization_without_accents

def rename_with_camel_case(column: str) -> str:
    """
    Renames a column using CamelCase convention.

    This function takes a column name as input, applies UTF-8 normalization
    without accents using the 'apply_utf8_normalization_without_accents' function
    from the 'shared_python_features_text.normalize' module. It then converts the
    normalized string to CamelCase by removing non-alphanumeric characters and
    capitalizing the first letter of each word.

    Args:
        column (str): The input column name to be renamed.

    Returns:
        str: The column name in CamelCase.

    Example:
        >>> rename_with_camel_case("First Name")
        'firstName'
        >>> rename_with_camel_case("customer_id")
        'customerId'
    """
    norm_column = apply_utf8_normalization_without_accents(column)
    column_clean = "".join(filter(lambda x: x.isalnum() or x.isdigit(), norm_column.title().replace(" ", "")))
    return column_clean[0].lower() + column_clean[1:]
