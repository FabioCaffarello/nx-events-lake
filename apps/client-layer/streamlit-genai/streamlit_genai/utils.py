
class BaseLogger:
    """
    A simple base logger class.

    This class is used for logging information.

    Methods:
        __init__(self) -> None:
            Initializes the BaseLogger instance.

    Attributes:
        info (callable): A callable method for logging information.

    """
    def __init__(self) -> None:
        self.info = print
