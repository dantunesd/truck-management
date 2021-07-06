# Trucks Management

This application is responsible for managing trucks

---

# Pre Requisites

Before running this application, you must to install:

`Docker version 20.10.5`

`Docker-compose version 1.28.6`

---

# Start

First of all, create a network, it will be required for running application and tests:

```bash
make create-network
```

Now run the following command to start application and its dependencies:

```bash
make start
```

Run the following bellow to see the application's logs:

```bash
make logs
```

---

# Unit Tests

Run the following command to run all unit tests:

```bash
make unit-test
```

---

# Integration Tests

Run the following command to run all integration tests:

```bash
make integration-test
```

---

# API Documentation

Here you can find an [OpenAPI spec file](https://github.com/loadsmart-recruiting/dantunesd/blob/main/docs/trucks-management-swagger.yaml) with all resources. Try it out in [editor.swagger](https://editor.swagger.io/)

Also a postman collection can be [found here](https://github.com/loadsmart-recruiting/dantunesd/blob/main/docs/trucks-management.postman_collection.json). Download it and import into your Postman.

---

# Folder Structure

## Top level directory structure

    .
    ├── .github                   # Github actions
    ├── docker                    # Docker stuff
    ├── docs                      # Documentation 
    ├── integration-tests         # Automated integration tests
    ├── truck-management          # source files 
    └── main.go                   # entrypoint of the application

## Truck Management directory structure

    .
    ├── ...
    ├── truck-management        # Source Files
    │   ├── api                 # Routers, handlers, middlewares etc
    │   ├── application         # Services related to application flow, etc 
    │   ├── domain              # Entities, domain services, etc
    │   └── infrastructure      # Repositories, drivers, configs, factories, etc
    └── ...

---

# Logs

All logs are logged in the stdout following the [ECS Logging standards](https://www.elastic.co/guide/en/ecs-logging/overview/current/intro.html).

---

# Github Actions

[Here you can find the CI](https://github.com/loadsmart-recruiting/dantunesd/actions/workflows/ci.yaml) of this application, running unit tests and integration tests for every commit done

---