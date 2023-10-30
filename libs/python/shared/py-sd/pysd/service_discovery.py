import os
from dataclasses import dataclass


class UnrecoverableError(Exception):
    def __init__(self, *args: object) -> None:
        """
        Initializes an UnrecoverableError.

        Args:
            args (object): Any additional arguments for the exception.
        """
        super().__init__(*args)


class ServiceUnavailableError(Exception):
    def __init__(self, *args: object) -> None:
        """
        Initializes a ServiceUnavailableError.

        Args:
            args (object): Any additional arguments for the exception.
        """
        super().__init__(*args)


@dataclass
class ServiceVars:
    rabbitmq: str = "RABBITMQ"
    configHandler: str = "SERVICES_CONFIG_HANDLER"
    fileCatalogHandler: str = "SERVICES_FILE_CATALOG_HANDLER"
    schemasHandler: str = "SERVICES_SCHEMAS_HANDLER"
    minio: str = "MINIO"
    services_rabbitmq_exchange: str = "services"



class ServiceDiscovery:
    def __init__(self, envvars):
        """
        Initializes a ServiceDiscovery instance.

        Args:
            envvars (dict): A dictionary of environment variables.

        Raises:
            UnrecoverableError: If environment variables are not set.
        """
        if envvars is None:
            raise UnrecoverableError('Environment variables not set')
        self._vars = envvars
        self._service_vars = ServiceVars()

    def _get_endpoint(self, var_name: str, service_name: str, protocol: str = "http") -> str:
        """
        Gets the endpoint for a service.

        Args:
            var_name (str): The name of the environment variable containing the service endpoint.
            service_name (str): The name of the service.
            protocol (str): The protocol to use (default is "http").

        Returns:
            str: The service endpoint.

        Raises:
            ServiceUnavailableError: If the environment variable is not set.
        """
        if var_name not in self._vars:
            raise ServiceUnavailableError(f'Environment variable {var_name} not set')
        tcp_addr = self._vars[var_name]
        gt_host = self._get_gateway_host(service_name)
        return tcp_addr.replace("tcp", protocol).replace("gateway_host", gt_host)

    def _get_gateway_host(self, service_name: str) -> str:
        """
        Gets the gateway host for a service.

        Args:
            service_name (str): The name of the service.

        Returns:
            str: The gateway host.

        Notes:
            If the 'GATEWAY_ENVIRONMENT' environment variable is not set, 'localhost' is returned.
        """
        if os.getenv('GATEWAY_ENVIRONMENT') is None:
            return 'localhost'
        return os.getenv(f'{service_name}_GATEWAY_HOST')

    def rabbitmq_endpoint(self) -> str:
        """
        Gets the RabbitMQ endpoint.

        Returns:
            str: The RabbitMQ endpoint in 'amqp' protocol.
        """
        service_name = self._service_vars.rabbitmq
        return self._get_endpoint("RABBITMQ_PORT_6572_TCP", service_name, protocol="amqp")

    def services_rabbitmq_exchange(self) -> str:
        """
        Gets the services RabbitMQ exchange.

        Returns:
            str: The name of the services RabbitMQ exchange.
        """
        return self._service_vars.services_rabbitmq_exchange

    def services_config_handler_endpoint(self):
        """
        Gets the services config handler endpoint.

        Returns:
            str: The services config handler endpoint.
        """
        service_name = self._service_vars.configHandler
        endpoint = self._get_endpoint("SERVICES_CONFIG_HANDLER_PORT_8000_TCP", service_name)
        if "localhost" in endpoint:
            endpoint = endpoint.replace("8000", "8002")
        return endpoint

    def services_schemas_handler_endpoint(self):
        """
        Gets the services schemas handler endpoint.

        Returns:
            str: The services schemas handler endpoint.
        """
        service_name = self._service_vars.schemasHandler
        endpoint = self._get_endpoint("SERVICES_SCHEMAS_HANDLER_PORT_8000_TCP", service_name)
        if "localhost" in endpoint:
            endpoint = endpoint.replace("8000", "8003")
        return endpoint

    def services_file_catalog_handler_endpoint(self):
        """
        Gets the services file catalog handler endpoint.

        Returns:
            str: The services file catalog handler endpoint.
        """
        service_name = self._service_vars.fileCatalogHandler
        endpoint = self._get_endpoint("SERVICES_FILE_CATALOG_HANDLER", service_name)
        if "localhost" in endpoint:
            endpoint = endpoint.replace("8000", "8004")
        return endpoint

    def minio_endpoint(self):
        """
        Gets the Minio endpoint.

        Returns:
            str: The Minio endpoint.
        """
        service_name = self._service_vars.minio
        endpoint = self._get_endpoint("MINIO_PORT_9000_TCP", service_name)
        return endpoint

    def minio_access_key(self):
        """
        Gets the Minio access key.

        Returns:
            str: The Minio access key.
        """
        return os.getenv("MINIO_ACCESS_KEY")

    def minio_secret_key(self):
        """
        Gets the Minio secret key.

        Returns:
            str: The Minio secret key.
        """
        return os.getenv("MINIO_SECRET_KEY")

def new_from_env():
    """
    Creates a ServiceDiscovery instance using environment variables.

    Returns:
        ServiceDiscovery: A new ServiceDiscovery instance.
    """
    return ServiceDiscovery(os.environ)
