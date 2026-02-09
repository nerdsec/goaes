# Contributing

All contributions are welcome.

## Workflow

1. Please first create an issue so we can discuss the changes.
2. Create a fork of the repo.
3. Create a feature branch.
4. Create a pull request when you're ready.

## Development

Build and test:

```bash
go build ./...
go test ./...
```

Run the linter:

```bash
golangci-lint run
```

## Guidelines

- Security-sensitive changes (anything in `internal/`) require careful review.
- Run all tests before submitting a PR.
- Follow existing code style and patterns.
