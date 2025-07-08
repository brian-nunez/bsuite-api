# BSuite API Server

This is the core HTTP API for the BSuite project.

It serves as a central, versioned entry point for handling requests, managing tasks, and orchestrating internal services. Built with Go and Echo, the API is lightweight, and fast.

## 🔧 Features

- Versioned routing (`/api/v1`)
- Middleware: Logger, Recovery, CORS, Request ID
- JSON error responses
- Graceful shutdown
- Dockerized for both development and production

## 🚀 Run Locally

```bash
go run ./cmd/main.go
```

The server listens on port `8080` by default. Set the `PORT` environment
variable to override this when running locally.

Or use the dev container:

```bash
make
```

## 📦 Production Build

```bash
make prod-run
```

or create binaries for Linux and MacOS (amd64 and arm64):

```bash
make build-binary
```

## ✅ Health Check

```
GET /api/v1/health
```
