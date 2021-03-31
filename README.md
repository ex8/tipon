# Tipon
Tipon is a backend infrastructure that allows users to tip other users using a micro-services architecture.

## Technologies
- Go
- gRPC
- MongoDB
- Docker

## Services
- `tips` - Tips is the tips gRPC server for internal service communication.
- `tips-api` - Tips API is the public facing Tips API.

## Build & Serve
You can use `docker-compose` to quickly spin up the entire system.
```
docker-compose up -d
```
