# Build stage
FROM golang:1.20-alpine3.17 AS builder
WORKDIR /app/server
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.17
WORKDIR /app/server
COPY --from=builder /app/server/main .
COPY server.env .

EXPOSE 8080
ENTRYPOINT [ "/app/server/main" ]