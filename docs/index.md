# Getting Started

Required Installations:

- [Node.js 18.x](https://nodejs.org/en/download/)
- [Golang 1.20](https://golang.google.cn/)
  - [Protobuf](https://grpc.io/docs/languages/go/quickstart/#prerequisites)
    ```shell
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
    ```
  - [Wire](https://pkg.go.dev/github.com/google/wire)
    ```shell
    go install github.com/google/wire/cmd/wire@latest 
    ```
- [Python 3.10](https://www.python.org/downloads/)
  - [Poetry](https://pypi.org/project/poetry/1.2.0b3/)
    ```shell
    pip install poetry==1.2.0b3
    ```
- [docker](https://www.docker.com/)
- [docker-compose](https://www.docker.com/)


## Install dependencies

```shell
npm install
```

```shell
poetry install
```

## Terminal virtual environment

```shell
poetry shell
```
