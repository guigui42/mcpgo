# Sample Octodemo MCP Server in Go

This is a sample [**Model Context Protocol (MCP)**](https://modelcontextprotocol.io/) server implementation called **MCP server Octodemo**. It demonstrates how to build an MCP with basic operations and provides a foundation for adding more operations for your use case.

## Features

- **Sample Operations**:
  - `get_services`: Retrieve all services.
  - `get_service_by_id`: Retrieve a specific service by its ID.

These operations are defined in [`main.go`](./cmd/main.go) and use the service schema and data defined in [`operations/service.ts`](./data/services,json).

- **Extensible**: You can easily add more operations to this server to suit your needs. The server is built using the MCP SDK, making it straightforward to define new tools and their corresponding request handlers.

## Installation

Follow these steps to set up and run the server:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/octodemo/mcp-octodemo
   cd mcp-octodemo
   ```

2. **Install Dependencies**:
   Make sure you have Node.js installed (version 16 or later is recommended). Then, run:
   ```bash
   npm mcp-octodemo
   ```

3. **Build the Project**:
   Compile the TypeScript code into JavaScript:
   ```bash
   npm run build
   ```

4. **Run and inspect the Server**:

You can run and inspect the server using the following command:
```bash
make run                                                        130 ↵
```

This will start the server and open the Model Context Protocol Inspector in your browser, allowing you to interact with the server and test its operations.

See the [MCP Inspector documentation](https://github.com/modelcontextprotocol/inspector?tab=readme-ov-file#mcp-inspector) for more details on how to use the inspector.
   

## Adding More Operations (TO DO)

The Utils class in [`operations/utils.ts`](./operations/utils.ts) operations that are not used in the server, this allows you to add more operations easily. 

You can use the prompt `add_tool` to add a new operation.


## Usage in your MCP Client

To use this MCP server in your own MCP client, you can reference it as a tool in your MCP configuration. Below is an example of how to configure the MCP server in your client:

```json
{
    "mcp": {
        "servers": {
            "localhttmMCP": {
                "type": "sse",
                "url": "http://localhost:8080/mcp/sse"
            }
        }
    }
}
```

Add this configration in  .vscode/mcp.json file of your client project. 

## Example Usage

You can use the following prompt in GitHu Copilot Chat to use the MCP server:

> List all the services

> Can you give the the owner and git url of the service 5 ?

> Can you give me the list of all the databases used by octodemo services?

Using the tool `octodemo` directly in the chat: 

> #get_service_by_id  10

_Note: the service is using a JSON list of service located in `operations/service.ts` file. Thie services are not real and are only used for demonstration purposes._

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.


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