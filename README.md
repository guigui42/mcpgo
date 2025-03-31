# MCPGO

A Go project with standard structure.

## Project Structure

```
.
├── api/          # API definitions, OpenAPI/Swagger specs, JSON schema files, protocol definition files
├── build/        # Build and package scripts, CI configurations
├── cmd/          # Main applications for this project
│   └── mcpgo/    # Main application entry point
├── configs/      # Configuration file templates or default configs
├── deployments/  # IaaS, PaaS, system and container orchestration deployment configurations
├── docs/         # Design and user documents
├── internal/     # Private application and library code
├── pkg/          # Library code that's ok to use by external applications
├── scripts/      # Scripts to perform various build, install, analysis, etc operations
├── test/         # Additional external test apps and test data
└── web/          # Web application specific components
```

## Getting Started

### Prerequisites

- Go 1.24 or later
- Make (for using the Makefile commands)

### Using the Makefile

This project includes a Makefile to simplify common development tasks.

To see all available commands:

```bash
make help
```

#### Common Commands

| Command | Description |
|---------|-------------|
| `make build` | Builds the application binary in `./bin` |
| `make run` | Runs the application locally |
| `make test` | Runs all tests |
| `make fmt` | Formats the Go code |
| `make lint` | Runs the linter |
| `make deps` | Updates and tidies dependencies |
| `make clean` | Cleans build artifacts |

#### Docker Commands

| Command | Description |
|---------|-------------|
| `make docker-build` | Builds the Docker image |
| `make docker-run` | Runs the application in a Docker container |

### Running Manually

#### Running Locally

```bash
go run cmd/mcpgo/main.go
```

Or using make:

```bash
make run
```

#### Building

```bash
go build -o bin/mcpgo cmd/mcpgo/main.go
```

Or using make:

```bash
make build
```

#### Docker

Build the Docker image:

```bash
docker build -t mcpgo .
```

Or using make:

```bash
make docker-build
```

Run the Docker container:

```bash
docker run -p 8080:8080 mcpgo
```

Or using make:

```bash
make docker-run
```