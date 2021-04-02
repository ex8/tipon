FROM golang:1.16 as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o tips-api

FROM alpine:3
COPY --from=builder /app/tips-api /
ENV GIN_MODE=release
CMD ["/tips-api"]