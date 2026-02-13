# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

goaes is a CLI tool for encrypting/decrypting files using AES-256-GCM with envelope encryption and Argon2id key derivation. Written in Go 1.26. **Not intended for production use.**

## Commands

```bash
go build ./...              # Build all
go build ./cmd/goaes        # Build CLI binary
go test ./...               # Run tests
go test -race -v ./...      # Run tests with race detector (CI uses this)
golangci-lint run           # Lint (v2, config in .golangci.yml)
go test -run TestName ./internal/  # Run a single test
```

## Architecture

**Two packages:**

- `cmd/goaes/` — CLI entry point using urfave/cli/v3. Commands live in `cmd/goaes/commands/` (encrypt, decrypt, generate, features). The passphrase is read from `GOAES_PASSPHRASE` env var (base64-encoded).

- `internal/` — Core cryptography. Uses envelope encryption:
  1. Derive KEK from passphrase via Argon2id (time=3, mem=256MiB, threads=4)
  2. Generate random DEK, wrap it with KEK using AES-256-GCM
  3. Encrypt data with DEK using AES-256-GCM
  4. Domain-separated AAD: `"wrap:dek:v1"` for key wrapping, `"data:msg:v1"` for data encryption

**File format:** `[Magic "GOAES" 5B] [Version 1B] [Salt 32B] [Wrapped DEK 60B] [Ciphertext variable]`

**Build tags:** `secret_enabled.go` / `secret_fallback.go` use `GOEXPERIMENT=runtimesecret` (Go 1.26) for memory protection. GoReleaser enables this for release builds.

## CI

GitHub Actions runs: `go test -race -v ./...`, `golangci-lint run`, `govulncheck`, and GoReleaser on tags. Go version: 1.26.

## Key Types (internal package)

`KEK`, `DEK`, `WrappedDEK`, `Ciphertext`, `Salt` — all `[]byte` aliases with validation. `Clear()` methods zero memory.
