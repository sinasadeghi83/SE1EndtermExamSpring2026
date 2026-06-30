# syntax=docker/dockerfile:1

FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/redbank .

FROM alpine:3.21

RUN apk add --no-cache ca-certificates && \
    addgroup -S app && adduser -S app -G app

WORKDIR /app

COPY --from=builder --chown=app:app /out/redbank ./redbank
COPY --chown=app:app config ./config

USER app

EXPOSE 8989

ENTRYPOINT ["./redbank"]
