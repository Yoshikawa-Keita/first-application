# Build stage
FROM golang:1.20-alpine3.17 AS builder
WORKDIR /app/bff
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.17
WORKDIR /app/bff
COPY --from=builder /app/bff/main .
COPY bff.env .

EXPOSE 8081
ENTRYPOINT [ "/app/bff/main" ]