# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

NeoBank Core (Go module `redbank`) — the practical-section submission for the
Software Engineering 1 end-term exam, Spring 2026. A digital banking backend built
on the **Ahu framework**, a Hexagonal/Ports-and-Adapters architecture for Go
microservices, scaffolded with the `ahum` CLI (`~/go/bin/ahum`).

Current scope is deliberately a **skeleton**: domain entities, interfaces, DB
integration, and design patterns for three domains (Account, Transaction, Loan) are
implemented, but most service/chain method bodies are stubs returning zero values.
Don't fill in real business logic unless the user explicitly asks to move past the
skeleton phase.

Read `ahu.md` (architecture/per-layer conventions) and `ahum.md` (CLI command
reference) before adding new layers or services — they document the framework this
codebase follows. `AGENTS.md` has the day-to-day conventions for extending this repo.

## Commands

```bash
go build ./...
go vet ./...
go test ./...
gofmt -l .                          # lists unformatted files; should be empty
gofmt -w .                          # fix formatting
go test ./service/account/...       # single package
go test ./service/loan/... -run TestCreditScoreStrategy_Evaluate -v   # single test
```

Run the whole stack (app + Postgres + Redis) locally with:

```bash
docker compose up --build      # serves Connect/gRPC-Web on localhost:8989
docker compose down -v         # stop and wipe the postgres/redis volumes
```

Config is loaded from `config/config.json` by viper, with nested keys overridable
via env vars using `.` → `_` (e.g. `infras.postgres.host` → `INFRAS_POSTGRES_HOST`,
set in `config.go`'s `SetEnvKeyReplacer`) — `docker-compose.yml` uses this to point
the app at the `postgres`/`redis` service names instead of the `config.json`
localhost defaults.

No Makefile, linter config, or Dockerfile exists — these four Go-toolchain commands
are the full verification loop. `go.mod` targets Go 1.25+.

Scaffolding new Connect/gRPC services or methods should go through the `ahum` CLI,
not be hand-written, because it keeps the `// @ahum:` marker comments (injection
points in `connect.go`, `config.go`, `registrar.go`) intact for future codegen:

```bash
ahum connect service add <domain>
ahum connect service version add v1 <domain>
ahum connect service method add <Method> <domain> v1
ahum connect gen        # regenerate protobuf/Connect code after editing a .proto
```

## Architecture

Dependency flow (inward only): `Edge → Chain → Service → Data ← Infrastructure`.
`data/{domain}` has zero external dependencies — it defines entities and repository
*interfaces*; `infrastructure/postgres/{domain}` implements those interfaces with GORM.

| Layer | Path | Responsibility |
|---|---|---|
| Edge | `edge/connect/{domain}/v1` | Connect/gRPC-Web handlers + `.proto`; generated code under `edge/connect/gen` |
| Chain | `chain/{domain}` | DI composition (`chain.go`), request validation, DTO mapping between proto and service types |
| Service | `service/{domain}` | Business rules; depends only on `data` interfaces |
| Data | `data/{domain}` | Entities (`entity.go`), repo interfaces (`repo.go`), domain errors (`error.go`) |
| Infrastructure | `infrastructure/postgres/{domain}` | GORM models + repo implementations |
| Events | `data/event` | `Event`/`Publisher`/`Subscriber` contracts + domain events; `infrastructure/eventbus` is the Redis-backed publisher |
| Test doubles | `infrastructure/{domain}test`, `infrastructure/eventtest` | In-memory repo/publisher mocks used by `service/*/handler_test.go` |

Domains implemented: `account`, `transaction`, `loan`. `service/loan/strategy.go`
holds the Strategy-pattern loan-evaluation contract (`EvaluationStrategy`) with
`CreditScoreStrategy`, `IncomeBasedStrategy`, and `CompositeStrategy` implementations.

Each domain's chain/service packages share the same internal package name as their
`data/{domain}` counterpart (e.g. `chain/account`, `service/account`, `data/account`
are all package `account`) — when one needs to import another, it's aliased
(`accountData "redbank/data/account"`, `accountSvc "redbank/service/account"`, etc.).
This is intentional and compiles fine; don't "fix" it by renaming packages.

## Conventions

- Money fields (`Account.Balance`, `Transaction.Amount`, `Loan.Principal`) are
  integers in the smallest currency unit (cents) — never float for money.
- Repository methods return concrete entity values/slices, not pointers; only the
  entity's `ID` field is a pointer (nil before persistence, set on `Create`).
- Service-layer and chain-layer each define their own DTOs in their own `entity.go`
  — don't pass `data` entities directly through to proto-facing edge code.
- Errors are sentinel `errors.New` values per layer (`data/{domain}/error.go`,
  `service/{domain}/error.go`); only the edge layer wraps them with `connect.NewError`.
- Tests follow Arrange-Act-Assert with explicit `// Arrange` / `// Act` / `// Assert`
  comments, built against the `infrastructure/*test` in-memory mocks — see
  `service/account/handler_test.go` for the pattern to follow for new domains.

## Out of Scope

Per the exam's practical-section scope: no full business-logic implementations
beyond what's already stubbed, no DevOps/CI pipelines, no Docker/Kubernetes, no
frontend/mobile client code. Confirm with the user before expanding into these areas.
