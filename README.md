# My Inventory

![Github Workflow](https://github.com/codedByHasan/my-inventory/actions/workflows/ci.yml/badge.svg)

## Getting started

Copy the [`example.env`](example.env) file to `.env` and update variables:

```bash
cp example.env .env
```

| Variable        | Description                                                                                                         |
| --------------- | ------------------------------------------------------------------------------------------------------------------- |
| `DB_USER`       | Username used to access mysql database.                                                                             |
| `DB_PASSWORD`   | Password used to access mysql database.                                                                             |
| `DB_NAME `      | Name of the database e.g., `my-inventory`.                                                                          |
| `DB_PORT`       | TCP port that database is running on.                                                                               |
| `DB_HOST`       | The host to deploy my-inventory to. Use `localhost` for local deployments and `mysql` for docker compose deployment |
| `APP_PORT`      | Port that the server will use.                                                                                      |
| `VITE_APP_PORT` | Equal to `APP_PORT`. Needs to match `APP_PORT` to allow frontend to connect to backend.                             |

## Running Using Docker Compose

To deploy my-inventory using docker compose, simply run the following:

```bash
docker compose up

# To rebuild image
docker compose up --build
```

## Running on Localhost

To run my-inventory on localhost:

```bash
go install # install dependencies
go build -o my-inventory cmd/api-server/main.go # build binary
./path/to/binary
```

## Running Frontend

```bash
npm --prefix frontend run dev
```

## Endpoints

The available endpoints are as specified in `handleRoutes()` in [`app.go`](/internal/app/app.go#L169).
The API endpoints are follows:

| HTTP Verb | Endpoint          | Description +                                                                    |
| --------- | ----------------- | -------------------------------------------------------------------------------- |
| GET       | `/api/products`   | Get a list of all products.                                                      |
| GET       | `/api/product/id` | Get information about the respective product.                                    |
| POST      | `/api/product`    | Create new product based on given information from user and save it to database. |
| PUT       | `/api/product/id` | Update the respective product with given information from user.                  |
| DELETE    | `/api/product/id` | Delete the respective product.                                                   |

## Acknowledgements

This project was created as part of KodeKloud's [Advanced Golang](https://learn.kodekloud.com/user/courses/advanced-golang) course.
