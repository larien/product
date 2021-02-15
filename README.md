# Product Service

Product Service is a implementation of services communication via gRPC and HTTP requests. Some technologies used here are Golang, Node.JS, gRPC, PostgreSQL, Docker and Docker Compose and many others.

These services were designed according to [Hash's hiring test](https://github.com/hashlab/hiring). To understand each service and its technical decisions, please access each service's README:
- [Product](product/README.md)
- [Discount](discount/README.md)

## Project structure

- **discount**: service 1 implementation, written in Node.JS
- **migrations**: contains the database migration scripts and it's used along golang-migrate
- **product**: service 2 implementation, written in Golang
- **protos**: contains the protobuf scripts and generated code for gRPC communication

## Installation

Install [**docker**](https://docs.docker.com/get-docker/) and [**docker-compose**](https://docs.docker.com/compose/) and build the images.

```bash
docker-compose build
```

## Usage

Export the necessary environment variables (hint: save them to a `.env` file and run `export $(xargs <.env)`). Below there are a few to get you started:

```bash
DB_USER=larien
DB_HOST=database
DB_PORT=5432
DB_NAME=product
DB_PASSWORD=lauren123
LOG_LEVEL=1
PORT=:8080
POSTGRESQL_URL=postgres://larien:lauren123@database:5432/product?sslmode=disable
GRPC_PORT=:50051
GRPC_HOST=0.0.0.0
```

Run the built images (include `-d` if you don't want to read the logs in your terminal)

```bash
docker-compose up --remove-orphans
```

## Tips and tricks

### Connecting to the database inside the container

```bash
docker exec -it database bash
su - postgres
psql -p 5432 -d <database-name> -U <user-name>
```

Happy inserting!

## Authors and acknowledgment

Lauren Ferreira - [larien.dev](larien.dev)

## Contributing
Pull requests are welcome, but there is no development planned for this project.

## License
[GNU General Public License v3.0](https://choosealicense.com/licenses/gpl-3.0/)