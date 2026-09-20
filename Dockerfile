FROM node:20-alpine AS web-builder

WORKDIR /app/web

COPY apps/web/package*.json ./
RUN npm ci

COPY apps/web/ ./
RUN npm run build

# -------------------------------------------------------------#

FROM golang:1.26.5-alpine AS builder

RUN apk add --no-cache \
    git \
    make \
    ca-certificates \
    tzdata

WORKDIR /app

COPY apps/backend/go.mod apps/backend/go.sum ./
RUN go mod download

COPY apps/backend/ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s -X main.version=$(cat VERSION 2>/dev/null || echo 'dev')" \
    -o /app/bin/buggeon \
    ./cmd/main.go

COPY --from=web-builder /app/web/dist /app/bin/static

#--------------------------------------------------------------#

FROM alpine:latest

RUN apk add --no-cache \
    ca-certificates \
    tzdata \
    curl \
    && cp /usr/share/zoneinfo/Europe/Moscow /etc/localtime \
    && echo "Europe/Moscow" > /etc/timezone

WORKDIR /app

COPY --from=go-builder /app/bin/buggeon /app/buggeon
COPY --from=go-builder /app/bin/static  /app/static

USER appuser

EXPOSE 9090

CMD [ "./buggeon" ]