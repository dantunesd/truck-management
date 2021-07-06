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

# Documentation

A postman collection with all resources can be [found here](https://github.com/loadsmart-recruiting/dantunesd/blob/main/docs/trucks-management.postman_collection.json). Download it and import into your Postman.

---