# Cebalrai

This is a simple project in go for using the PostgreSQL driver (pgx)

## How to run

### With docker compose

If you have docker compose installed, you can run the following command:

```bash
$ docker-compose up
```

This will start a PostgreSQL database, pgAdmin and the application.

- You can see the pgAdmin interface at http://localhost:5050
- You can see the application at http://localhost:8080

### IF you have a PostgreSQL database running

You can run the application directly:

```bash
$ go run main.go
```

## How to use

You can use the following endpoints:

- GET /users
- GET /users?id={your_id}

## Environment variables

See the [.env.example](.env.example) file for more information.

## Dependencies

- [pgx](https://pkg.go.dev/github.com/jackc/pgx/v5)
- [mux](https://pkg.go.dev/github.com/gorilla/mux)
