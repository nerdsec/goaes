# goaes Code Review Memory

## Project Overview
- AES-256-GCM file encryption CLI tool in Go 1.25
- Envelope encryption: passphrase -> Argon2id -> KEK -> wraps random DEK -> DEK encrypts data
- Dependencies: urfave/cli v3.6.2, golang.org/x/crypto v0.47.0
- File format: magic "GOAES" + version byte + 32-byte salt + 60-byte wrapped DEK + ciphertext

## Architecture
- `internal/` package: crypto primitives, packaging, types
- `cmd/goaes/` CLI layer using urfave/cli v3
- Passphrase read from env var `GOAES_PASSPHRASE` (base64-encoded)

## Key Findings (2026-02-09 review)
- Passphrase not zeroed after use in NewKEKFromEnvB64
- Argon2id output decoded from passphrase is not zeroed
- UnpackagePayload returns slices aliasing the input buffer (no copy)
- Entire file read into memory - no streaming support
- wrappedDEKLength hardcoded to 60 (12 nonce + 32 key + 16 tag) - correct but fragile
- PackagePayload assumes salt is exactly 32 bytes with no validation
- clear() function is custom rather than using crypto/subtle or Go 1.21+ clear builtin
- No test for wrong-key decryption, empty input, or tampered ciphertext
- generate command does not zero key after printing

## Documentation Review Findings (2026-02-09)
- README "memory: 256mb" should be "256 MiB" (code: 256*1024 KiB)
- README says "not secure" but code has solid crypto; contradicts THREAT_MODEL.md
- THREAT_MODEL.md says Argon2id passphrase mode doesn't exist yet -- but code ALWAYS uses Argon2id
- THREAT_MODEL.md says AAD should bind version/salt but code AAD is only domain tags
- README version 0.7.0 is stale; main.go defaults to "dev" via ldflags
- decrypt requires destination arg; README only shows encrypt usage
- Bug report template uses web-app fields (browser, smartphone) not CLI fields
- README encrypt example doesn't show destination arg (optional for encrypt)
- No documentation of file format structure anywhere user-facing
