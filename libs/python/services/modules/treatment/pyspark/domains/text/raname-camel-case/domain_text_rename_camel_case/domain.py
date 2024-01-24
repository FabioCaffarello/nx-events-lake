from pyspark.sql import DataFrame
from shared_python_features_text.rename import rename_with_camel_case

class Domain:
    def __init__(self, dataframe: DataFrame):
        """
        Initializes a Domain instance with a DataFrame.

        Args:
            dataframe (DataFrame): The input DataFrame for the domain.
        """
        self.dataframe = dataframe

    def transform(self) -> DataFrame:
        """
        Transforms the columns of the DataFrame using CamelCase renaming.

        This method applies the 'rename_with_camel_case' function to each column
        of the DataFrame and returns a new DataFrame with columns renamed in CamelCase.

        Returns:
            DataFrame: The transformed DataFrame with CamelCase column names.
        """
        rename_columns = list(map(lambda col: rename_with_camel_case(col), self.dataframe.columns))
        return self.dataframe.toDF(*rename_columns)
