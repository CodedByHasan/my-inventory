# My Inventory

![Github Workflow](https://github.com/codedByHasan/my-inventory/actions/workflows/ci.yml/badge.svg)

## Getting started

Copy the [`example.env`](example.env) file to `.env` and update variables:

```bash
cp example.env .env
```

| Variable      | Description                                |
| ------------- | ------------------------------------------ |
| `DB_USER`     | Username used to access mysql database.    |
| `DB_PASSWORD` | Password used to access mysql database.    |
| `DB_NAME `    | Name of the database e.g., `my-inventory`. |
| `DB_PORT`     | TCP port that database is running on.      |
| `APP_PORT`    | Port that the server will use.             |

Export `ENVPATH` and set it to absolute path of `.env` file.

```bash
export ENVPATH=/path/to/env/file
```

## Running the Database

To run the database:

```bash
docker compose up
```

## Running the server

To run the server:

```bash
go install # install dependencies
go build -o my-inventory cmd/api-server/main.go # build binary
go run /path/to/binary
```

## Endpoints

The available endpoints are as specified in `handleRoutes()` in [`app.go`](./app.go#L142).
The API endpoints are follows:

| HTTP Verb | Endpoint      | Description                                                                      |
| --------- | ------------- | -------------------------------------------------------------------------------- |
| GET       | `/products`   | Get a list of all products.                                                      |
| GET       | `/product/id` | Get information about the respective product.                                    |
| POST      | `/product`    | Create new product based on given information from user and save it to database. |
| PUT       | `/product/id` | Update the respective product with given information from user.                  |
| DELETE    | `/product/id` | Delete the respective product.                                                   |

## Acknowledgements

This project was created as part of KodeKloud's [Advanced Golang](https://learn.kodekloud.com/user/courses/advanced-golang) course.
