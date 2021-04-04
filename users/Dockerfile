FROM golang:1.16 as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o users-server

FROM alpine:3
COPY --from=builder /app/users-server /
ENV GIN_MODE=release
CMD ["/users-server"]