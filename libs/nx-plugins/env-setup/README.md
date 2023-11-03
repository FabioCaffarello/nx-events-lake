# env-setup

The `env-setup` Nx plugin provides three executors to manage your local environment configuration, including creating a bucket, inserting configurations, and inserting schemas. This README will guide you through the usage of each executor and provide essential information about the plugin.

## Table of Contents
- [Create Bucket Executor](#create-bucket-executor)
- [Insert Configurations Executor](#insert-configurations-executor)
- [Insert Schemas Executor](#insert-schemas-executor)

## Create Bucket Executor

The Create Bucket Executor allows you to create a bucket using the [Minio](https://min.io/) client. To use this executor, follow these steps:

### Executor Configuration

Add a target in `targets` of the project `project.json`:

```json
...
"create-bucket": {
      "executor": "@nx-plugins/env-setup:create-bucket"
    }
...

```

### Executor Usage

Run the Create Bucket Executor using the following command:

```bash
npx nx create-bucket your-project-name --name=your-bucket-name
```

## Insert Configurations Executor

The Insert Configurations Executor allows you to insert configuration files into a specified endpoint. To use this executor, follow these steps:

### Executor Configuration

Add a target in `targets` of the project `project.json`:

```json
...
"insert-configs": {
      "executor": "@nx-plugins/env-setup:insert-configs"
    }
...

```

### Executor Usage

Run the Insert Configurations Executor using the following command:

```bash
npx nx insert-configs your-project-name --source=your-bucket-name
```

## Insert Schemas Executor

The Insert Schemas Executor allows you to insert schema files into a specified endpoint. To use this executor, follow these steps:

### Executor Configuration

Add a target in `targets` of the project `project.json`:

```json
...
"insert-schemas": {
      "executor": "@nx-plugins/env-setup:insert-schemas"
    }
...

```

### Executor Usage

Run the Insert Schemas Executor using the following command:

```bash
npx nx insert-schemas your-project-name --source=your-bucket-name
```


---

Feel free to customize the provided executor configurations and adapt them to your project's needs. For more information on Nx and the available options, please refer to the [Nx documentation](https://nx.dev/).