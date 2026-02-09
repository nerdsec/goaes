
# goaes

`goaes` is a very simple tool for encrypting files with AES-256-GCM.

## about

I wanted to write some Go and encryption fascinates me, so I wrote `goaes`.

### what it is

- it's fun
- it's fast
- it's simple

### what it isn't

- it's not meant for production
- it's not meant for sensitive data
- it has not been formally audited

## how it works

`goaes` uses envelope encryption with AES-256-GCM:

1. A random salt is generated per file.
2. Your secret is run through [Argon2id](https://en.wikipedia.org/wiki/Argon2) to derive a Key Encryption Key (KEK):
	- time: `3`
	- memory: `256 MiB`
	- threads: `4`
	- key length: `32`
3. A random 32-byte Data Encryption Key (DEK) is generated per file.
4. The DEK is [wrapped](https://en.wikipedia.org/wiki/Key_wrap) (encrypted) with the KEK using AES-256-GCM.
5. File contents are encrypted with the DEK using [AES-256-GCM](https://en.wikipedia.org/wiki/Galois/Counter_Mode).
6. Random nonces are generated for each encryption operation.

### file format

```
[Magic "GOAES" (5 bytes)] [Version (1 byte)] [Salt (32 bytes)] [Wrapped DEK (60 bytes)] [Ciphertext (variable)]
```

## getting started

1. Generate a new key.

```bash
goaes generate
```

Don't use this one. This one is mine.

```
XGfpiNUvKJy8k7KeUEyhev4jkTIajb1s9CMJP9xH/7A=
```

2. Set `GOAES_PASSPHRASE` to the key.
3. Encrypt and decrypt files.

```bash
export GOAES_PASSPHRASE=XGfpiNUvKJy8k7KeUEyhev4jkTIajb1s9CMJP9xH/7A=

# encrypt (destination defaults to source + ".goaes")
goaes encrypt ./input.txt
goaes encrypt ./input.txt ./custom-output.enc

# decrypt (destination is required)
goaes decrypt ./input.txt.goaes ./input.txt
```

### usage

```bash
NAME:
   goaes - Simple AES encryption built with Go

USAGE:
   goaes [global options] [command [command options]]

COMMANDS:
   generate, g  Generate a base64 encoded key
   encrypt, e   Encrypt a file
   decrypt, d   Decrypt a file
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## reference material

- https://en.wikipedia.org/wiki/Key_wrap
- https://en.wikipedia.org/wiki/Advanced_Encryption_Standard
- https://en.wikipedia.org/wiki/Galois/Counter_Mode
- https://en.wikipedia.org/wiki/Authenticated_encryption
- https://en.wikipedia.org/wiki/Argon2
- https://en.wikipedia.org/wiki/Key_derivation_function
- https://en.wikipedia.org/wiki/Salt_(cryptography)
- https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html

## inspiration

- https://github.com/FiloSottile/age
- https://github.com/restic/restic
