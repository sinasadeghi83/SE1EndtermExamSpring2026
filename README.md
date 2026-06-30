# NeoBank Core

**SE1 Endterm Exam — Spring 2026 — Practical Section**

A digital banking backend skeleton built in Go on the **Ahu framework**, a
Hexagonal (Ports & Adapters) architecture for microservices. This repository
covers the domain model, repository contracts, business-logic skeletons, and
Connect/gRPC-Web edges for three banking domains: **Account**, **Transaction**,
and **Loan**.

> **Scope note:** this is intentionally a skeleton. Domain entities,
> interfaces, database integration, and design patterns are implemented;
> most service/chain method bodies are stubs. See [`AGENTS.md`](AGENTS.md)
> for the exam's scoping rules.

---

## Architecture

Dependencies flow inward only — `data/{domain}` has zero external dependencies
and is implemented by `infrastructure/postgres/{domain}`.

```
┌─────────────────────────────────────────────────────────────┐
│  EDGE        edge/connect/{domain}/v1   (Connect / gRPC-Web) │
└───────────────────────────┬───────────────────────────────────┘
                             ▼
┌─────────────────────────────────────────────────────────────┐
│  CHAIN       chain/{domain}     (DI, validation, DTO mapping) │
└───────────────────────────┬───────────────────────────────────┘
                             ▼
┌─────────────────────────────────────────────────────────────┐
│  SERVICE     service/{domain}   (business rules)              │
└───────────────────────────┬───────────────────────────────────┘
                             ▼
┌─────────────────────────────────────────────────────────────┐
│  DATA        data/{domain}      (entities, repo interfaces)   │
└───────────────────────────▲───────────────────────────────────┘
                             │ implements
┌─────────────────────────────────────────────────────────────┐
│  INFRASTRUCTURE  infrastructure/postgres/{domain}, eventbus   │
└─────────────────────────────────────────────────────────────┘
```

| Layer | Path | Responsibility |
|---|---|---|
| Edge | `edge/connect/{domain}/v1` | Connect/gRPC-Web handlers + `.proto` definitions |
| Chain | `chain/{domain}` | Dependency injection, request validation, DTO mapping |
| Service | `service/{domain}` | Business rules, depends only on `data` interfaces |
| Data | `data/{domain}` | Entities, repository interfaces, domain errors |
| Events | `data/event` | `Event` / `Publisher` / `Subscriber` contracts + domain events |
| Infrastructure | `infrastructure/postgres/{domain}`, `infrastructure/eventbus` | GORM repositories, Redis-backed event publisher |
| Test doubles | `infrastructure/{domain}test`, `infrastructure/eventtest` | In-memory mocks for unit tests |

Domains implemented: **`account`**, **`transaction`**, **`loan`**. Loan
underwriting is modeled with the Strategy pattern
(`service/loan/strategy.go`): `CreditScoreStrategy`, `IncomeBasedStrategy`,
and `CompositeStrategy`.

For full framework conventions, see [`ahu.md`](ahu.md) (architecture
reference) and [`ahum.md`](ahum.md) (CLI reference for the `ahum` scaffolding
tool).

## Tech Stack

| Component | Technology |
|---|---|
| Language | Go 1.25+ |
| API Protocol | [Connect](https://connectrpc.com/) (gRPC-Web compatible) |
| Database | PostgreSQL via [GORM](https://gorm.io) |
| Cache / Event bus | Redis |
| Configuration | [Viper](https://github.com/spf13/viper) (JSON file + env var overrides) |
| Scaffolding | [`ahum`](https://github.com/Ahu-Tools/ahum) CLI |

## Getting Started

### Run with Docker (recommended)

Starts the app together with Postgres and Redis:

```bash
docker compose up --build
```

The Connect/gRPC-Web server is available at `localhost:8989`. Tear everything
down (including data volumes) with:

```bash
docker compose down -v
```

### Run locally

Requires Go 1.25+, a running PostgreSQL instance, and a running Redis instance.

```bash
go mod download
go run main.go
```

Configuration is read from `config/config.json` and can be overridden with
environment variables (nested keys use `_` in place of `.`, e.g.
`infras.postgres.host` → `INFRAS_POSTGRES_HOST`):

| Env var | Default | Purpose |
|---|---|---|
| `APP_SECRET_KEY` | _(see config.json)_ | Application secret |
| `EDGES_CONNECT_SERVER_HOST` | `0.0.0.0` | Connect server bind host |
| `EDGES_CONNECT_SERVER_PORT` | `8989` | Connect server port |
| `INFRAS_POSTGRES_HOST` | `127.0.0.1` | Postgres host |
| `INFRAS_POSTGRES_PORT` | `5432` | Postgres port |
| `INFRAS_POSTGRES_USER` | `postgres` | Postgres user |
| `INFRAS_POSTGRES_PASSWORD` | `postgres` | Postgres password |
| `INFRAS_POSTGRES_DB_NAME` | `redbank` | Postgres database name |
| `INFRAS_REDIS_HOST` | `127.0.0.1` | Redis host |
| `INFRAS_REDIS_PORT` | `6379` | Redis port |

### Verify a running instance

The Connect protocol speaks plain JSON over HTTP/1.1 or HTTP/2, so a health
check needs nothing more than `curl`:

```bash
curl -X POST http://localhost:8989/hello_v_1.Service/Health \
  -H "Content-Type: application/json" -d '{}'
# {"message":"UP"}
```

## Development

```bash
go build ./...     # compile everything
go vet ./...        # static analysis
go test ./...        # run the test suite
gofmt -l .            # check formatting (should print nothing)
```

Adding a new domain or RPC method should go through the `ahum` CLI so the
`// @ahum:` injection-point markers stay intact:

```bash
ahum connect service add <domain>
ahum connect service version add v1 <domain>
ahum connect service method add <Method> <domain> v1
ahum connect gen
```

See [`AGENTS.md`](AGENTS.md) and [`CLAUDE.md`](CLAUDE.md) for conventions
followed by AI coding agents working in this repository.

## License

MIT — see [`LICENSE`](LICENSE).
