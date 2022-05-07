# Crono

## Requirements

You will require the following tools to run this project:
- `golang` (>= 1.16.6) - https://go.dev/doc/install
- `git` - https://git-scm.com/downloads

## Getting Started

1. Clone the repository
```bash
> git clone https://github.com/BarendVanDerBerg/crono
> cd crono
```

2. Install the modules
```bash
> go mod tidy
```

3. Build the application
```bash
> go build -o crono
```

4. Run your expression and command
```
> ./crono "*/15 0 1,15 * 1-5 /usr/bin/find -r passwords.txt"
```

5. Celebrate the results! 🥳

## Testing

1. Run all unit tests in all subdirectories
```bash
> go test ./...
```

2. Run all tests with coverage
```bash
> go test ./... -coverprofile=coverage.out
> go tool cover -html=coverage.out
```

3. Run tests to obtain their benchmarks
```bash
> go test ./... -bench=.
```