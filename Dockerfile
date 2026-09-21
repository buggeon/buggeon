FROM node:20-alpine AS web-builder

WORKDIR /app/web

COPY apps/web/package*.json ./
RUN npm ci

COPY apps/web/ ./
RUN ls
RUN npm run build

# -------------------------------------------------------------#

FROM golang:1.26.5-alpine AS go-builder

RUN apk add --no-cache \
    git \
    make \
    ca-certificates \
    tzdata

WORKDIR /app

COPY apps/backend/go.mod apps/backend/go.sum ./
RUN go mod download

COPY apps/backend/ ./

RUN ls

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s -X main.version=$(cat VERSION 2>/dev/null || echo 'dev')" \
    -o /app/bin/buggeon \
    ./cmd/main.go

COPY --from=web-builder /app/web/dist /app/bin/static

RUN ls

#--------------------------------------------------------------#

FROM alpine:latest

RUN apk add --no-cache \
    ca-certificates \
    tzdata \
    curl \
    && cp /usr/share/zoneinfo/Europe/Moscow /etc/localtime \
    && echo "Europe/Moscow" > /etc/timezone

WORKDIR /app

COPY --from=go-builder /app/bin/buggeon /app/
COPY --from=go-builder /app/bin/static  /app/static

RUN ls

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 9187

CMD [ "./buggeon" ]