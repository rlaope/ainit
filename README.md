# ainit

**ainit** is a command-line tool that generates a complete Go project scaffold from natural language input.  
It automates the creation of the project directory, README, Go module initialization, and essential boilerplate files — all from a single command.

## Features

- Parses natural language descriptions to generate a meaningful `README.md`
- Creates a new project directory with the specified name
- Automatically generates:
  - `README.md`
  - `.gitignore`
  - `.gitattributes`
  - `Makefile`
  - `go.mod` (via `go mod init`)
  - `cmd/main.go` with a basic "Hello, world" program
- Prepares a clean, ready-to-code Go project structure

## Usage

```bash
go run main.go "A gRPC-based chat server in Go" chat-server
```

This will create the following structure:

```go
chat-server/
├── README.md
├── .gitignore
├── .gitattributes
├── Makefile
├── go.mod
└── cmd/
    └── main.go
```

## Requirements
Go 1.18 or higher

Git (optional)

GPT API key (optional, for enhanced README generation)

## Roadmap
- GPT API integration to dynamically generate README content
- Support for multiple project templates (REST API, CLI, gRPC, etc.)
- Auto-generation of .env, Dockerfile, CI/CD files
- Git repository initialization and remote push automation
- Optional inclusion of test and lint setups

## Installation
Currently, ainit runs via go run. In the future, it will support go install for global CLI usage.

## Contributing
Contributions are welcome. This project is in its early experimental stage, and feedback, issues, and pull requests are highly appreciated.

## LICENSE

```
If you’d like this README to include additional sections like **badge integrations**, **real terminal screenshot**, or **AI-generated README samples**, let me know — I can help expand it. Ready for the next step?
```