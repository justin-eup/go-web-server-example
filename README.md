# Go Web Server Example

This project is a sample Go web server built using the Gin framework, following the [Golang Standard Project Layout](https://github.com/golang-standards/project-layout) and incorporating frequently used [design](https://github.com/tmrts/go-patterns) [patterns](https://refactoring.guru/design-patterns/go). The application is designed to be modular, scalable, and production-ready.

## Features

- **Gin Framework**: Fast and lightweight HTTP web framework.
- **Standard Project Layout**: Organized for scalability and maintainability.
- **Health Check API**: Simple endpoint to check server health.
- **Hello API**: Dynamic response to demonstrate routing.

## Project Structure

```
.
├── cmd
│   └── server
│       └── main.go
├── config
│   └── config.go
├── internal
│   ├── handlers
│   │   ├── health.go
│   │   └── hello.go
│   └── routes
│       └── routes.go
├── pkg
│   └── server
│       └── server.go
├── README.md
├── go.mod
└── go.sum
```

## Run

```
go run cmd/server/main.go
```