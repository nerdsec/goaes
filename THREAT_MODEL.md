# Threat Model

## Summary

`goaes` is a CLI tool for encrypting files at rest before storage using an AEAD scheme (e.g., AES-256-GCM). The intended users are the author and other CLI-savvy technical users. The key material is provided via a base64-encoded value from an environment variable.

This threat model focuses on offline compromise scenarios (stolen encrypted files) and does not attempt to defend against a compromised runtime environment.

## Goals

### Security goals

- **Confidentiality:** An attacker who obtains encrypted files should not be able to recover plaintext without the secret key.
- **Integrity:** Any modification to ciphertext should be detected at decryption time (fail closed).
- **Practical usability:** The tool should remain simple to use for technical users and not require complex operational dependencies.

### Non-goals

- **Multi-user key sharing / access control**
- **Compliance alignment** (e.g., FIPS, SOC 2, HIPAA)
- **Key rotation / crypto-erasure** as a first-class feature (not a current priority)

## System overview

### Primary workflow

1. User supplies a base64-encoded secret via environment variable (the "KEK" or master key).
2. `goaes` encrypts file content and writes an encrypted output file.
3. `goaes` decrypts files given the same environment-supplied key.

### Data stored alongside ciphertext

- File format version identifier (cleartext)
- Salt (cleartext) used for Argon2id key derivation

Note: cleartext metadata is permitted as it does not reveal the key or plaintext.

## Assets

- **Plaintext file contents** (primary asset)
- **Encryption key material** (base64 env var secret; decoded bytes in memory during runtime)
- **Integrity of decrypted output** (avoid silent corruption or malleability)

## Trust boundaries

- **Trusted:** Local machine *only insofar as* it is not compromised at runtime.
- **Untrusted:** Storage locations where ciphertext may reside (cloud object storage, backups, shared filesystems), and any party who can read encrypted files.

## Adversaries (in scope)

1. **Offline file thief**
   - Obtains encrypted files (e.g., stolen laptop disk image, copied backup set, exfiltrated object storage bucket data).
   - Has unlimited time to attempt offline attacks.

2. **Curious cloud/storage provider**
   - Can read stored ciphertext and associated metadata.

3. **Malicious insider with file access**
   - Has access to encrypted files via shared storage, backups, or internal systems.
   - Does **not** have authorized access to the environment variable secret.

## Adversaries (explicitly out of scope)

If an attacker can obtain the environment variable secret, the primary confidentiality guarantee is already lost. Therefore the following are out of scope:

- **Compromised host / root attacker**
- **Memory scraping / debugger attachment**
- **Process inspection (e.g., reading `/proc`, environment, shell history)**
- **Malware with user-level access capable of reading env vars**
- **Keylogging / clipboard capture used to obtain the secret**
- **Any attacker who can read the secret from CI logs, shell config, or secrets management mistakes**

This is a deliberate scoping decision: `goaes` is not attempting to be a complete endpoint hardening solution.

## Assumptions

- Cryptographic primitives are implemented correctly by Go’s standard library.
- AES-256-GCM (or equivalent AEAD) is used correctly:
  - Unique nonce per encryption under the same key
  - Authentication failures cause decryption to fail closed
- The environment-supplied secret is high entropy:
  - Prefer **32 random bytes** (256 bits) encoded as base64
  - Not a human-memorable password

Argon2id key derivation provides additional hardening against offline guessing, but a high-entropy secret remains the primary defense.

## Key management model

### Current model

- Single secret (provided via environment variable) used to encrypt/decrypt files.

### Design intent

- The security of `goaes` depends primarily on **keeping this secret unknown to attackers** and **ensuring it has sufficient entropy**.

### KDF position (current)

- All secrets are run through Argon2id (time=3, memory=256 MiB, threads=4, keyLen=32) with a random per-file salt to derive the KEK.
- This hardens against offline brute-force even if the supplied secret has less-than-ideal entropy.

## Threats and mitigations

### T1: Offline brute force / guessing

**Threat:** Attacker steals encrypted files and attempts to guess the key.
**Mitigation:** Use a high-entropy secret (32 random bytes). All secrets are additionally hardened by Argon2id key derivation with a random per-file salt.

### T2: Ciphertext tampering / corruption

**Threat:** Attacker modifies ciphertext to cause malicious or silent changes to plaintext.
**Mitigation:** Use AEAD (AES-GCM) and fail closed on authentication error.

### T3: Nonce reuse under the same key

**Threat:** Reusing a nonce with AES-GCM under the same key can catastrophically weaken confidentiality and integrity.
**Mitigation:** Ensure nonce uniqueness per encryption operation (cryptographically random nonces or a robust deterministic strategy). Treat nonce generation as a critical security requirement.

### T4: Metadata manipulation (version/salt)

**Threat:** Attacker edits metadata to influence parsing or key derivation logic.
**Mitigation:** Domain-separation strings are bound into AAD (`"wrap:dek:v1"` and `"data:msg:v1"`), preventing cross-context confusion attacks. The salt is authenticated indirectly: tampering with it causes Argon2id to derive the wrong KEK, which fails at DEK unwrap (AEAD authentication error).

### T5: Key disclosure via operational mistakes

**Threat:** Secret leaks through shell history, logs, CI output, process inspection, or misconfigured environment handling.

**Mitigation (operational guidance):**
- Do not pass secrets via CLI args (they show in process lists and shell history).
- Avoid logging the secret or decoded bytes.
- Encourage secret injection via a secrets manager or `.env` with correct permissions.
- Prefer per-host secrets and limit reuse where feasible.

(These mitigations reduce risk but do not change the out-of-scope boundary: if the secret is exposed, confidentiality is lost.)

## Security requirements (implementation checklist)

- Use AEAD encryption (AES-256-GCM or equivalent).
- Enforce:
  - Nonce uniqueness per encryption under the same key
  - Authentication check on decrypt (fail closed)
- File format includes:
  - Magic bytes + version
  - Nonce
  - Ciphertext (including auth tag)
  - Optional salt (cleartext), if used
- Bind into AAD:
  - `"wrap:dek:v1"` for DEK wrapping operations (domain-separates key wrapping from data encryption)
  - `"data:msg:v1"` for data encryption operations
- Zero/overwrite sensitive buffers where practical (best-effort; Go does not guarantee).
- Provide a deterministic test vector suite to prevent accidental format/crypto regressions.

## Open design decisions (explicit)

- **Key rotation / crypto-erasure:** Not currently a priority. The existing KEK/DEK envelope scheme provides a foundation for future key rotation support with header metadata and versioned key slots.
- **Threat boundary:** The tool does not attempt to protect against runtime compromise or secret exfiltration from the host environment.

## "Secure by default" guidance for users

- Generate the secret as **32 random bytes**, then base64 encode it.
- Store it in a secure secret store or protected env file; do not paste it into shell history.
- Treat encrypted files as safe to store in untrusted locations, but treat the secret as highly sensitive.
