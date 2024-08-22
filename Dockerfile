# Указываем базовый образ
FROM golang:1.23.0-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app cmd/wallet/main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app .
CMD ["./app"]
