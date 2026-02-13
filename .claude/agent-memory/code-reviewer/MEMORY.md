# goaes Code Review Memory

## Project Overview
- AES-256-GCM file encryption CLI tool in Go 1.26
- Envelope encryption: passphrase -> Argon2id -> KEK -> wraps random DEK -> DEK encrypts data
- Dependencies: urfave/cli v3.6.2, golang.org/x/crypto v0.47.0
- File format: magic "GOAES" + version byte + 32-byte salt + 60-byte wrapped DEK + ciphertext
- Build tag `goexperiment.runtimesecret` for runtime/secret support

## Architecture
- `internal/` package: crypto primitives, packaging, types
- `cmd/goaes/` CLI layer using urfave/cli v3
- Passphrase read from env var `GOAES_PASSPHRASE` (base64-encoded)
- SecretDo abstraction: no-op fallback vs runtime/secret.Do

## Fixed Since 2026-02-09 Review
- UnpackagePayload now copies slices (no longer aliases input buffer)
- PackagePayload now validates salt/edek lengths via validPayloadLengths
- Tests added for wrong-key, tampered ciphertext, empty plaintext, wrong passphrase
- generate command now clears raw key bytes (but encoded string still leaks)
- NewKEKFromEnvB64 now defers Clear(passphrase) on decoded bytes

## Open Issues (2026-02-13 review)
- CRITICAL: Clear() uses subtle.XORBytes(b,b,b) - may be optimized away; use builtin clear()
- CRITICAL: Passphrase passed as string (immutable, cannot be zeroed) throughout API
- MAJOR: AAD does not bind salt/version/edek to data ciphertext
- MAJOR: No file size limit on ReadFile (OOM risk)
- MAJOR: generate leaks base64-encoded key as string outside SecretDo
- MAJOR: No check for source==destination (data destruction risk)
- MAJOR: Plaintext not zeroed after encrypt; decrypted plaintext not zeroed after write
- MINOR: validAESKeyLen accepts 16/24 but only 32 is ever used
- MINOR: NewSalt() uses variable named "key" and keyLen instead of saltLen
- MINOR: wrappedDEKLen=60 should be computed from components

## Documentation Issues (2026-02-09, likely still open)
- THREAT_MODEL.md says AAD should bind version/salt but code AAD is only domain tags
- README version possibly stale
