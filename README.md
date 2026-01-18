
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
3. Run the `goaes encrypt` command.

```bash
export GOAES_PASSPHRASE=XGfpiNUvKJy8k7KeUEyhev4jkTIajb1s9CMJP9xH/7A=

goaes encrypt ./input.txt
```

```
hexdump -C ./input.txt.goaes

00000000  d0 17 66 f2 cb b9 14 af  56 14 75 26 9b 89 50 f1  |..f.....V.u&..P.|
00000010  4c a2 14 ed 58 58 82 a3  64 8d 98 e9 02 ff a0 e5  |L...XX..d.......|
00000020  40 a7 bc 2e 81 65 24 68  3c c4 e3 c3 c3 de 2b 55  |@....e$h<.....+U|
00000030  3c 1a d7 26 59 8f f7 cf  88 00 ac 06 4a b7 dc 25  |<..&Y.......J..%|
00000040  75 46 60 6a 15 a9 2e 3a  ff 15 6d 25 39 25 15 71  |uF`j...:..m%9%.q|
00000050  a1 10 55 28 cc 2d 7e 67  58 46 f6 8e 48 ce e3 7e  |..U(.-~gXF..H..~|
00000060  4c 4c 7a b3 c6 c9 ef 4d  3a a0 3f 41 03 7e 2f 3f  |LLz....M:.?A.~/?|
00000070  09 6d b3 74 4c d9 ae 5f  a5 ea 49 5d 8f f8 87 6f  |.m.tL.._..I]...o|
00000080  c3 4a 7d 67                                       |.J}g|
00000084
```

### usage

```bash
NAME:
   goaes - Simple AES encryption built with Go

USAGE:
   goaes [global options] [command [command options]]

VERSION:
   0.7.0

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
- https://words.filippo.io/2025-state/

## inspiration

- https://github.com/FiloSottile/age
- https://github.com/restic/restic
