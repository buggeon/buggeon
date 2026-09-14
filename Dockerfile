FROM golang:1.26.5-alpine AS builder

RUN apk add --no-cache \
    git \
    make \
    ca-certificates \
    tzdata

WORKDIR /app
COPY go.mod go.sum ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s -X main.version=$(cat VERSION 2>/dev/null || echo 'dev')" \
    -o /app/bin/buggeon \
    ./cmd/main.go

FROM alpine:latest

RUN apk add --no-cache \
    ca-certificates \
    tzdata \
    curl \
    && cp /usr/share/zoneinfo/Europe/Moscow /etc/localtime \
    && echo "Europe/Moscow" > /etc/timezone

RUN adduser -D -g '' appuser

COPY --from=builder /app/bin/buggeon /app/buggeon

USER appuser

WORKDIR /app

EXPOSE 9090

CMD [ "./buggeon" ]