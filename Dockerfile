# Build stage
FROM golang:1.21-alpine3.18 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Building app
RUN cd cmd/api && CGO_ENABLED=0 GOOS=linux go build -o api

# Final stage
FROM alpine:3.18 as runner

ARG TELEGRAM_TOKEN
ARG TELEGRAM_CHAT_ID

ENV TELEGRAM_TOKEN $TELEGRAM_TOKEN
ENV TELEGRAM_CHAT_ID $TELEGRAM_CHAT_ID

ENV ENV_CONFIG_ONLY=true
ENV GIN_MODE=release
ENV HOST 0.0.0.0
ENV PORT=8080

# Tạo thư mục config nếu chưa tồn tại
RUN mkdir -p /app/config

COPY --from=builder /app/cmd/api/api /app/
COPY ./config/config_cloud.yml /app/config/config.cloud.yml

WORKDIR /app

EXPOSE 8080

# Run the web service on container startup.
CMD ["/app/api", "-e", "cloud"]
# CMD ["/app/api", "--config_url", "https://storage.cloud.google.com/fx_golang_server_buckets/config.yml"]
