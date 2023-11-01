# services-config-handler

This is a Go API project structured following the clean architecture pattern with RabbitMQ integration for asynchronous messaging.

## Overview

This project is designed to provide a scalable and maintainable architecture for building APIs. It separates concerns into distinct layers and integrates RabbitMQ for message queuing, ensuring the decoupling of components.

## Project Structure

- `main.go`: Entry point of the application, where all components are initialized and the API server is started.
- `envs`: Configuration files for different environments.
- `configs`: Configuration Loading.
- `internal`: The core application code.
  - `entity`: Defines the application's entities.
  - `event`: Handles application events.
  - `infra`: Infrastructure code, such as database and web.
    - `web`: Handles HTTP server, routing, and API endpoints.
  - `usecase`: Contains business logic and use cases.

## Getting Started

1. **Setup Configuration**:

   - Create a `.env.{ENVIRONMENT}` file based on the example provided in the `envs` folder.

2. **Build The Application**:

```sh
npx nx image services-orchestration-services-config-handler --env=<ENVIRONMENT>
```

3. **Run the Application**:

```sh
docker-compose up -d
```

   The API server will start on the specified port, and RabbitMQ will be used for asynchronous messaging.


## Configuration

The project uses a configuration system that loads environment-specific settings from the `envs` folder. Ensure that you provide the necessary environment variables or configuration files for your specific deployment.

## Dependency Management

This project uses Go modules for dependency management. You can use the `nx` commands to add or update dependencies as needed.

```sh
npx nx go-tidy services-orchestration-services-config-handler
```

## Testing

Unit tests and integration tests can be added to the respective packages in the `internal` directory. You can use Go's built-in testing framework for this purpose and you can run it with `nx`:

```sh
npx nx test services-orchestration-services-config-handler
```

## Contributing

Contributions to this project are welcome. If you find issues or have suggestions for improvements, feel free to create an issue or submit a pull request.


## Acknowledgments

- This project follows the clean architecture pattern as described by Robert C. Martin.
- RabbitMQ is used for asynchronous messaging.

Feel free to adapt this README file to your specific project's needs and add more details as necessary.
