# binary

This repository contains a simple utility to generate binary file. It is designed for testing purposes, with values configured during the build process in GitHub workflows.

## Build

Clone the repository to get started:

```bash
git clone https://github.com/shachindra/binary.git
cd binary
go build -ldflags "-X main.version=1.0.0 -X main.codeHash=abc123 -X main.goVersion=go1.20" -o binary main.go
```

## Usage

Run the CLI commands:

```bash
./binary version --json
```
```
{"version": "1.0.0", "codehash": "crazzzy", "goVersion": "go1.29"}
```

The input values can also be set dynamically in GitHub workflows.

# Publish Release

Tag Your Release:

Create a tag for your release:
```
git tag v1.0.0
git push origin v1.0.0
```