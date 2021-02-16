# Product

This service was designed in [Golang 1.15](https://golang.org/) and communicates with the database with [PostgreSQL 9.6](https://www.postgresql.org/) and with Discount service via [gRPC](https://grpc.io/).

## Table of Contents

  - [Business Requirements](#business-requirements)
  - [Folder Structure](#folder-structure)
  - [Usage](#usage)
  - [Endpoints](#endpoints)
    - [GET Products](#get-products)
  - [Questions](#questions)
  - [Database](#database)
  - [Testing](#testing)
    - [Unit Testing](#unit-testing)
    - [Integration Testing](#integration-testing)
    - [Mocking](#mocking)
  - [Linting](#linting)
  - [Logging](#logging)

## Business Requirements

Service 2: Products listing

- It exposes a HTTP route in such a way that `GET /product` returns a JSON with a products list
- This route must optionally receive the user ID via `X-USER-ID` header
- To obtain a custom discount, this service must use service 1 (discount)
- In case service 1 returns an error, the products list still needs to be returned, but with the errored product without discount
- If discount service (1) is down, listing service (2) must keep working and returning the list, but it won't apply the discounts

## Folder Structure

- controller: contains the business logic inside the application and communicates only with product and discount repositories
- drivers: utility packages that wrap external dependencies and other packages
- entity: structures that defines the objects manipulated by the application, they are used by every layer
- handler: the application entrypoint that defines the router and the endpoints to be acessed via HTTP
- repository: contains the methods to access external dependencies, like discount service and database


## Usage

To run it locally, download the dependencies:
```bash
go mod tidy
```

Export the necessary environment variables (checkout `drivers/config/config.go`). A few to get you started (with a local database with the product database created without password):
```
DB_USER=larien
DB_HOST=localhost
DB_PORT=5432
DB_NAME=product
LOG_LEVEL=1
PORT=:8080
GRPC_PORT=:50051
GRPC_HOST=0.0.0.0
```

And then, just run the application:
```bash
go run main.go
```

## Endpoints

### GET Products

Obtains a list of products with each discount if applicable.

**GET /v1/product**

_Optional_:

- Header:
  - X-USER-ID: user ID as UUID V4

Possible answers:

- **200** Success
```json
// with user ID
[
    {
        "id": "9e277f2f-3edb-464d-82bf-5e41973fe668",
        "price_in_cents": 1000,
        "title": "Product 1",
        "description": "Product 1 description",
        "discount": {
            "percentage": 10,
            "value_in_cents": 100
        }
    }
]
// without user ID
[
    {
        "id": "9e277f2f-3edb-464d-82bf-5e41973fe668",
        "price_in_cents": 1000,
        "title": "Product 1",
        "description": "Product 1 description",
        "discount": {
            "percentage": 0,
            "value_in_cents": 0
        }
    }
]
```
- **204** No Content
No product was found.
```json
null
```

- **400** Bad Request
```json
{
    "error": "invalid product ID"
}
```

- **500** Internal Server Error
```json
{
    "error": "an error occurred when listing the products"
}
```

**GET /status**

Service healthcheck

Possible answers:

- **200** Success
```json
null
```

## Questions

- Why do we need to send the user ID for every product if the user ID is always the same? The system performance would probably be way better if the discount was obtained once by service 1 and then applied to the products in service 2.
- Do we really need to send the product ID every time since we are not using it in discount service? Isn't it an early implementation for a future requirement?
- If the user ID is optionally sent, should we even try to get the discounts if there is no discount related to product in service 2?

## Database

This service implements the following technologies:
- [PostgreSQL 9.6](https://www.postgresql.org/): open source relational database
- [golang-migrate](https://github.com/golang-migrate/migrate): performs the migration on schema changes via CLI for this project

## Testing

### Unit Testing

The tests implement the following packages:
- [testify](https://github.com/stretchr/testify): provides assertion methods for cleaner outputs
- [mockery](https://github.com/vektra/mockery): generates mocks from existing data structures to be used by unit tests

To run unit tests only:
```
go test ./... -v -short
```

### Integration Testing

The tests implement the following packages:
- [testify](https://github.com/stretchr/testify): provides assertion methods and a testsuite to setup and teardown tests
- [faker](https://github.com/bxcodec/faker): to generate random data so that tests are more consistent
- [dockertest](https://github.com/ory/dockertest): to run each suite in a docker container that is purged after the tests are run

Every testsuite creates its own container and their tests are executed there. Each test runs its test in a clean environment because the related table is created before each test and it's dropped after each test so that tests are isolated. The connection is also verified before and after each test.

To run all tests:
```
go test ./... -v -p 1 -cover
```

### Mocking

Mocks were generated with [mockery](https://github.com/vektra/mockery) with the following command:
```bash
mockery --name=<input-interface-name> --structname=<output-struct-name> --inpackage
```

## Linting

This service uses [golangci-lint](https://golangci-lint.run/) configured with `.golangci.yml` YAML file.
You can start in your machine and run the linter with:

```bash
golangci-lint run
```

## Logging

The log wrapper was written on top of [zerolog](github.com/rs/zerolog) and it's used only in the requests.

The logging level is configure in `LOG_LEVEL` integer environment variable with the following values:

- 0 - Debug
- 1 - Info
- 2 - Warn
- 3 - Error
- 4 - Fatal
- 5 - Panic
- 6 - No level
- 7 - Disabled
- 8 - Trace level
