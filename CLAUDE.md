# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

This is a Go project using standard Go tooling:

- `go build` - Build the project
- `go run main.go` - Run the application
- `go test ./...` - Run all tests
- `go mod tidy` - Clean up module dependencies
- `go fmt ./...` - Format all Go files
- `go vet ./...` - Run Go vet for static analysis

## Architecture Overview

This project demonstrates a Clean Architecture pattern in Go with the following structure:

- **domain/ports/**: Contains interface definitions that define the contracts between layers
  - `CoursePriceUseCase` interface in courseprice.go
  - `Prices` interface in prices.go

- **application/**: Contains use case implementations (business logic layer)
  - `courseprice/usecase.go` - CoursePrice struct implementing business logic
  - `prices/usecase.go` - Prices struct implementing business logic

- **di/**: Dependency injection layer that wires up dependencies
  - `course_price.go` - Factory function for CoursePriceUseCase
  - `prices.go` - Factory function for PriceUseCase

The project appears to be set up as a demonstration or learning project for Clean Architecture patterns, with minimal implementations in each layer. The main.go file is currently empty, suggesting this is either a work-in-progress or specifically designed to illustrate architectural concepts.

## Module Information

- Module name: `cycle-imports`
- Go version: 1.24.4
- No external dependencies currently used