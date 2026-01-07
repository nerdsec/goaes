
# goaes

`goaes` is a very simple tool for encrypting files with AES-256 GCM.

## about

I wanted to write some Go and encryption fascinates me, so I wrote `goaes`.

### what it is

- it's fun
- it's fast
- it's simple

### what it isn't

- it's not meant for production
- it's not meant for sensitive data
- it's not secure

## how it works

- it uses [Argon2id](https://en.wikipedia.org/wiki/Argon2)
	- time: `3`
	- memory: `256mb`
	- threads: `4`
	- key length: `32`
- it uses [key wrapping](https://en.wikipedia.org/wiki/Key_wrap)

## getting started

1. Generate a new passphrase.

```bash
goaes generate
```

Don't use this one. This one is mine.

```
XGfpiNUvKJy8k7KeUEyhev4jkTIajb1s9CMJP9xH/7A=
```

2. Set `GOAES_PASSPHRASE` to the passphrase.
3. Run the `goaes` command.

```bash
export GOAES_PASSPHRASE=XGfpiNUvKJy8k7KeUEyhev4jkTIajb1s9CMJP9xH/7A=

goaes encrypt -s ./input.txt -d ./output.enc
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
   --help, -h  show help
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
- https://words.filippo.io/2025-state/
