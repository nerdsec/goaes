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

- it uses [PBKDF2](https://en.wikipedia.org/wiki/PBKDF2)
- it uses [key wrapping](https://en.wikipedia.org/wiki/Key_wrap)

## getting started

1. Generate a new passphrase.
2. Generate a new salt.

```bash
goaes generate
```

Don't use these. These are mine.

```bash
XGfpiNUvKJy8k7KeUEyhev4jkTIajb1s9CMJP9xH/7A=
tjrXXdH+0/NXSsaDOYuGtEM2zXxWFNPXSPXoli5++iE=
```

3. Set `GOAES_PASSPHRASE` to the passphrase.
4. Set `GOAES_SALT` to the salt.
5. Run the `goaes` command.

```bash
export GOAES_PASSPHRASE=XGfpiNUvKJy8k7KeUEyhev4jkTIajb1s9CMJP9xH/7A=
export GOAES_SALT=tjrXXdH+0/NXSsaDOYuGtEM2zXxWFNPXSPXoli5++iE=

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
