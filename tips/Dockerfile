FROM golang:1.16 as builder
# TODO: Install 'protoc' to compile protobufs
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o tips-server

FROM alpine:3
COPY --from=builder /app/tips-server /
CMD ["/tips-server"]