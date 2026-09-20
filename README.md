<div align="center">

  <img src="documentation/images/logo.png" width="200" alt="Buggeon logo" />

  <h5>Open source project managing platform</h5>

  <p>
    <a href="LICENSE"><img src="https://img.shields.io/github/license/buggeon/buggeon?style=flat-square" alt="License"></a>
    <a href="https://github.com/buggeon/buggeon/actions"><img src="https://img.shields.io/github/actions/workflow/status/buggeon/buggeon/ci.yml?style=flat-square" alt="CI"></a>
    <img src="https://img.shields.io/badge/Go-1.26-00ADD8?style=flat-square&logo=go" alt="Go">
    <img src="https://img.shields.io/badge/React-19-61DAFB?style=flat-square&logo=react" alt="React">
  </p>

</div>

---

**Buggeon** is an open source platform for managing projects, tasks, and team collaboration. It brings together issue tracking, visual planning (diagrams, boards), and team communication in a single interface.

## Features

- **Task management** — create, assign, and track statuses and priorities.
- **Visual planning** — built-in Excalidraw-based diagrams and drag-and-drop boards.
- **Collaboration** — real-time updates via WebSocket.
- **Analytics** — charts and progress metrics (Recharts).
- **Metrics** — Buggeon exposes Prometheus metrics out of the box.

## Architecture

The monorepo consists of two applications:

| Component | Path | Stack |
|-----------|------|-------|
| **Backend** | `apps/backend` | Go 1.26, Gin, gqlgen (GraphQL), MongoDB, S3 (AWS SDK), JWT, Prometheus, Swagger |
| **Frontend** | `apps/web` | React 19, TypeScript, Vite, Apollo Client, Redux Toolkit, Excalidraw, dnd-kit, Recharts, SCSS |

The backend exposes a GraphQL API (schemas defined in `apps/backend/internal/graph`); the frontend talks to it via Apollo Client.

## Quick Start

> [!IMPORTANT]
> Make sure Docker is installed on your system before proceeding.

Create a `.env` file in the project root:

```bash
DB_USER=buggeon
DB_PASSWORD=change_me
RUSTFS_ACCESS_KEY=change_me
RUSTFS_SECRET_KEY=change_me
ADMIN_LOGIN=admin
ADMIN_PASSWORD=change_me
ADMIN_NAME=Admin
ADMIN_EMAIL=admin@example.com
```

Then create `docker-compose.yaml`

```yaml
services:
  backend:
    image: ghcr.io/buggeon/backend:latest
    container_name: buggeon-backend
    ports:
      - "9090:9090"
    environment:
      PORT: 9090
      DB_HOST: mongo
      DB_PORT: 27017
      DB_USER: ${DB_USER}
      DB_PASSWORD: ${DB_PASSWORD}
      S3_ENDPOINT: http://rustfs:9000
      S3_ACCESS_KEY: ${RUSTFS_ACCESS_KEY}
      S3_SECRET_KEY: ${RUSTFS_SECRET_KEY}
      S3_REGION: us-east-1
      ADMIN_LOGIN: ${ADMIN_LOGIN}
      ADMIN_PASSWORD: ${ADMIN_PASSWORD}
      ADMIN_NAME: ${ADMIN_NAME}
      ADMIN_EMAIL: ${ADMIN_EMAIL}
    depends_on:
      - rustfs
    restart: unless-stopped

  mongo:
    image: mongo:7
    container_name: buggeon-mongo
    environment:
      MONGO_INITDB_ROOT_USERNAME: ${DB_USER}
      MONGO_INITDB_ROOT_PASSWORD: ${DB_PASSWORD}
    volumes:
      - ./mongo-data:/data/db
    restart: unless-stopped

  # Any S3-compatible storage can be used instead of RustFS
  rustfs:
    image: rustfs/rustfs:latest
    container_name: buggeon-rustfs
    ports:
      - "9000:9000"
      - "9001:9001"
    environment:
      RUSTFS_ACCESS_KEY: ${RUSTFS_ACCESS_KEY}
      RUSTFS_SECRET_KEY: ${RUSTFS_SECRET_KEY}
      RUSTFS_VOLUMES: /data
      RUSTFS_ADDRESS: "0.0.0.0:9000"
      RUSTFS_CONSOLE_ADDRESS: "0.0.0.0:9001"
    volumes:
      - ./rustfs-data:/data
    restart: unless-stopped
```

Run:

```
docker-compose up -d
```

Buggeon will be available at http://localhost:9090

## Project Structure

```
buggeon/
├── apps/
│   ├── backend/          # Go API (GraphQL, Gin, MongoDB)
│   │   ├── cmd/          # Entry point
│   │   ├── config/       # Configuration
│   │   ├── docs/         # Swagger documentation
│   │   ├── internal/     # Business logic and GraphQL schema
│   │   ├── metrics/      # Prometheus metrics
│   │   └── test/         # Tests
│   └── web/              # React SPA
│       ├── src/          # Source code
│       ├── public/       # Static assets
│       └── styles/       # SCSS
├── documentation/        # Docs and images
├── Dockerfile
├── Makefile
└── LICENSE               # AGPL-3.0
```

## License

Licensed under the **GNU Affero General Public License v3.0** (AGPL-3.0). See [LICENSE](LICENSE) for details.