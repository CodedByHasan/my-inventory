# My Inventory

![Github Workflow](https://github.com/CodedByHasan/my-inventory/actions/workflows/main.yml/badge.svg)

## Getting started

Copy the `example.env` file to `.env` and update variables:

```bash
cp example.env .env
```

| Variable      | Description                               |
| ------------- | ----------------------------------------- |
| `DB_USER`     | Username used to access mysql database    |
| `DB_PASSWORD` | Password used to access mysql database    |
| `DB`          | Name of the database e.g., `my-inventory` |
| `APP_PORT`    | Port that the server will use.            |

You will need to install mysql on your host machine before running server.

## Running the server

To run the server:

```bash
go install # install dependencies
go build # build binary
go run /path/to/binary
```

## Endpoints

The available endpoints are as specified in `handleRoutes()` in [`app.go`](./app.go).
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
