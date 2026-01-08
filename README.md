
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
3. Run the `goaes encrypt` command.

```bash
export GOAES_PASSPHRASE=XGfpiNUvKJy8k7KeUEyhev4jkTIajb1s9CMJP9xH/7A=

goaes encrypt ./input.txt
```

```bash
hexdump -C ./input.txt.goaes

00000000  3e 7f 03 01 01 14 45 6e  63 72 79 70 74 65 64 44  |>.....EncryptedD|
00000010  61 74 61 50 61 79 6c 6f  61 64 01 ff 80 00 01 03  |ataPayload......|
00000020  01 03 44 45 4b 01 0a 00  01 04 53 61 6c 74 01 0a  |..DEK.....Salt..|
00000030  00 01 07 50 61 79 6c 6f  61 64 01 0a 00 00 00 ff  |...Payload......|
00000040  8d ff 80 01 3c a6 d4 3d  5d ab b3 49 74 dd 5d 0f  |....<..=]..It.].|
00000050  1d bc 93 01 78 c9 5a 39  37 53 8f 09 40 56 00 a0  |....x.Z97S..@V..|
00000060  5a a2 03 7c 71 ae 2f 54  f4 fc d9 0d f4 35 b5 df  |Z..|q./T.....5..|
00000070  21 e0 18 ef 54 60 3b 61  38 f5 3b 79 be 08 c4 a5  |!...T`;a8.;y....|
00000080  c4 01 20 49 2d 4d 72 02  d0 43 a3 e6 2c 30 6f ba  |.. I-Mr..C..,0o.|
00000090  66 ed cd b4 13 d6 24 8f  e4 8c 07 5a 09 0a a2 e8  |f.....$....Z....|
000000a0  75 08 8f 01 28 b5 16 a9  f4 98 6d 55 32 86 57 01  |u...(.....mU2.W.|
000000b0  09 24 4f 82 72 ba 0f ee  88 6d 07 b8 e3 ff af 16  |.$O.r....m......|
000000c0  89 45 bb 87 50 e0 a2 82  ee 25 88 63 7d 00        |.E..P....%.c}.|
000000ce
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

## inspiration

- https://github.com/FiloSottile/age
- https://github.com/restic/restic
