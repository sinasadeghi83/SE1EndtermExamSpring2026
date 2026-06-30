# AGENTS.md

Guidance for AI coding agents (Claude Code, Cursor, Copilot, etc.) working in this repository.

## Project Overview

**NeoBank Core** (Go module `redbank`) is the practical-section submission for the
Software Engineering 1 end-term exam, Spring 2026. It is a digital banking backend
built on the **Ahu framework**, a Hexagonal/Ports-and-Adapters architecture for Go
microservices.

Current scope is intentionally a **skeleton**: domain entities, interfaces, database
integration structures, and design patterns for three domains — Account, Transaction,
Loan — are in place, but most business-logic method bodies are stubs (they return zero
values). Do not silently fill in real business logic; if asked to move past the
skeleton phase, confirm scope with the user first.

Framework reference docs live at the repo root and should be read before adding new
layers or services:
- `ahu.md` — architecture and per-layer conventions (Edge/Chain/Service/Data/Infrastructure)
- `ahum.md` — `ahum` CLI command reference (project/service/edge scaffolding)

## Architecture

```
Edge (edge/connect)  →  Chain (chain/{domain})  →  Service (service/{domain})  →  Data (data/{domain})  ←  Infrastructure (infrastructure/...)
```

Dependencies point inward. `data/{domain}` defines entities and repository
interfaces with zero external dependencies; `infrastructure/postgres/{domain}`
implements those interfaces with GORM.

| Layer | Path | Responsibility |
|---|---|---|
| Edge | `edge/connect/{domain}/v1` | Connect/gRPC-Web handlers, proto definitions, generated code under `edge/connect/gen` |
| Chain | `chain/{domain}` | DI composition (`chain.go`), request validation, proto/service DTO mapping |
| Service | `service/{domain}` | Business rules, depends only on `data` interfaces |
| Data | `data/{domain}` | Entities (`entity.go`), repo interfaces (`repo.go`), domain errors (`error.go`) |
| Infrastructure | `infrastructure/postgres/{domain}` | GORM models + repo implementations |
| Events | `data/event` | `Event`/`Publisher`/`Subscriber` contracts and domain events; `infrastructure/eventbus` has the Redis-backed publisher |
| Test doubles | `infrastructure/{domain}test`, `infrastructure/eventtest` | In-memory repo/publisher mocks for unit tests |

Domains currently implemented: `account`, `transaction`, `loan`. `service/loan/strategy.go`
holds the Strategy-pattern loan-evaluation contract (`EvaluationStrategy`) with
`CreditScoreStrategy`, `IncomeBasedStrategy`, and `CompositeStrategy`.

## Adding a New Feature

Prefer the `ahum` CLI over hand-writing edge boilerplate — it keeps the `// @ahum:`
marker comments (injection points in `connect.go`, `config.go`, `registrar.go`, etc.)
intact so future codegen doesn't clobber manual edits.

```bash
ahum connect service add <domain>
ahum connect service version add v1 <domain>
ahum connect service method add <Method> <domain> v1
ahum connect gen   # regenerate protobuf/Connect code after editing a .proto
```

After scaffolding the edge, hand-write the `data/{domain}`, `service/{domain}`,
`chain/{domain}`, and `infrastructure/postgres/{domain}` packages following the
existing `account`/`transaction`/`loan` packages as templates.

## Build, Test, Lint

There is no Makefile; use the Go toolchain directly.

```bash
go build ./...
go vet ./...
go test ./...
gofmt -l .          # should print nothing; gofmt -w . to fix
```

Run these four before considering any change complete. `go.mod` targets Go 1.25+.

## Running with Docker

`docker compose up --build` runs the app plus Postgres and Redis (the app serves
Connect/gRPC-Web on `localhost:8989`). `docker compose down -v` tears it down and
wipes the data volumes. The `Dockerfile` is a multi-stage build (`golang:1.25-alpine`
→ `alpine:3.21`, non-root user). Compose overrides `config/config.json` via env vars
(`INFRAS_POSTGRES_HOST=postgres`, `INFRAS_REDIS_HOST=redis`, etc.) using viper's
`.`→`_` key replacement set up in `config/config.go`.

## Conventions

- Money fields (`Account.Balance`, `Transaction.Amount`, `Loan.Principal`) are
  integers in the smallest currency unit (cents) — never use float for money.
- Repository methods return concrete entity values/slices, not pointers; only the
  entity's `ID` field is a pointer (nil before persistence).
- Service-layer and chain-layer each define their own DTOs (`entity.go` in each
  package) — don't reuse `data` entities directly in proto-facing code.
- Errors are sentinel `errors.New` values per layer (`data/{domain}/error.go`,
  `service/{domain}/error.go`); wrap with `connect.NewError` only at the edge layer.
- Tests follow Arrange-Act-Assert with explicit `// Arrange` / `// Act` / `// Assert`
  comments, using the `infrastructure/*test` in-memory mocks — see
  `service/account/handler_test.go` for the pattern.

## Out of Scope (for now)

Per the exam's practical-section scope: no full business-logic implementations
beyond what's already stubbed, no DevOps/CI pipelines, no Docker/Kubernetes, no
frontend/mobile client code. Ask before expanding scope in these directions.
