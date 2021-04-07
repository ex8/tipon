# Tipon
Tipon is a backend infrastructure that allows users to tip other users. It utilizes a distributed service architecture.

## Technologies
- Go
- gRPC
- MongoDB
- Docker

## Micro-Services
- `core` - Tipon abstraction library in use by each micro-service.
- `tips` - Tips gRPC server for internal service communication.
- `tips-api` - Public facing Tips API.
- `users` - Users gRPC server for internal service communication.
- `users-api` - Public facing User API.
- `auth` - Auth gRPC server for internal tokens.
- `auth-api` - Public facing Auth API for generating external tokens.

## Build & Serve
You can use `docker-compose` to quickly spin up the entire system.
```
docker-compose up -d
```
