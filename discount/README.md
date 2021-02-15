# Discount

This server is used as gRPC server and it's running at port :50051 by default.
It receives a request ID, product ID and user ID. It calculates the discount rules according to the received information and returns a discount percentage.

## Technologies

- Node.JS
- Postgres 9.6
- gRPC

## Business Requirements

Service 1: Individual product discount
- This service received a product and an user ID and returns a discout
- The discount application rules are:
  - If it's user's birthday, the product will have a discount of 5%
  - If it's blackfriday (25/11 for this test), the product will have a discount of 10%
  - The discount can't exceed 10%

## Usage

Download the dependencies:

```bash
npm install
```

Define the environment variables for database (`POSTGRESQL_URL`) and gRPC (`GRPC_HOST` and `GRPC_PORT`), for instance:
```
POSTGRESQL_URL=postgres://larien:lauren123@database:5432/product?sslmode=disable
GRPC_PORT=:50051
GRPC_HOST=0.0.0.0
```

Start the application:
```bash
npm start
```