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


## Sample Workflow
```
name: Build and Release Binary

on:
  push:
    tags:
      - 'v*' # Trigger workflow on version tags (e.g., v1.0.0)

jobs:
  build:
    name: Build and Release
    runs-on: ubuntu-latest

    strategy:
      matrix:
        os: [windows, linux, darwin]
        arch: [amd64, arm64] # You can add more architectures if needed (e.g., arm64)

    steps:
      # Checkout the code
      - name: Checkout code
        uses: actions/checkout@v3

      # Set up Go environment
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: 1.23 # Specify the Go version you want to use

      # Get the commit hash
      - name: Get commit hash
        id: get-commit
        run: echo "::set-output name=commit::$(git rev-parse --short HEAD)"

      # Build the binary
      - name: Build binary
        run: |
          GOOS=${{ matrix.os }} GOARCH=${{ matrix.arch }} go build \
            -ldflags "-X main.version=${{ github.ref_name }} -X main.codeHash=${{ steps.get-commit.outputs.commit }} -X main.goVersion=$(go version | awk '{print $3}')" \
            -o binary-${{ matrix.os }}-${{ matrix.arch }} main.go

      # Upload the binary as an artifact
      - name: Upload binary artifact
        uses: actions/upload-artifact@v4
        with:
          name: binary-${{ matrix.os }}-${{ matrix.arch }}
          path: binary-${{ matrix.os }}-${{ matrix.arch }}

  release:
    name: Create Release
    needs: build
    runs-on: ubuntu-latest

    steps:
      # Checkout the code
      - name: Checkout code
        uses: actions/checkout@v3

      # Create a release
      - name: Create GitHub Release
        uses: softprops/action-gh-release@v1
        with:
          files: |
            binary-windows-amd64
            binary-linux-amd64
            binary-macos-amd64
            binary-windows-arm64
            binary-linux-arm64
            binary-macos-arm64
        env:
          TOKEN: ${{ secrets.AUTH_TOKEN }}
```