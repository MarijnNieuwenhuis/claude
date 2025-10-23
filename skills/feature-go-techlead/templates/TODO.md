# Feature: {Feature Name} - Implementation Plan

**Status**: Not Started
**Created**: {Date}
**Last Updated**: {Date}
**Go Version**: {1.21+ or specific version}
**Feature Doc**: [FEATURE.md](./FEATURE.md)

## Overview

{1-2 paragraph summary of what will be built, why it's needed, and the high-level approach}

### Key Objectives

1. {Objective 1}
2. {Objective 2}
3. {Objective 3}

### Success Criteria

- [ ] {Criterion 1}
- [ ] {Criterion 2}
- [ ] {Criterion 3}

---

## Implementation Phases

### Phase 1: Foundation

Setup project structure, define types, and create interfaces.

- [ ] Task 1.1: {Task name}
- [ ] Task 1.2: {Task name}
- [ ] Task 1.3: {Task name}

### Phase 2: Core Implementation

Implement core business logic and functionality.

- [ ] Task 2.1: {Task name}
- [ ] Task 2.2: {Task name}
- [ ] Task 2.3: {Task name}

### Phase 3: Documentation & Deployment

Documentation, examples, and deployment preparation.

- [ ] Task 3.1: {Task name}
- [ ] Task 3.2: {Task name}
- [ ] Task 3.3: {Task name}

---

## Detailed Task Breakdown

### Phase 1: Foundation

#### Task 1.1: Initialize Go Module and Project Structure

**Dependencies**: None
**Status**: Not Started

**Description**:
Create the Go module and set up the project directory structure following Go best practices.

**Implementation Steps**:
1. Run `go mod init {module-path}`
2. Create directory structure:
   ```
   {project-name}/
   ├── cmd/
   │   └── {app-name}/
   │       └── main.go
   ├── pkg/
   │   └── {package}/
   ├── internal/
   │   └── {package}/
   ├── go.mod
   ├── go.sum
   ├── README.md
   └── Makefile
   ```
3. Create initial README.md with project overview

**Acceptance Criteria**:
- [ ] Go module initialized
- [ ] Directory structure created
- [ ] README.md exists with basic info
- [ ] Makefile with common targets (build, run)

**Files to Create**:
- `go.mod`
- `README.md`
- `Makefile`
- `cmd/{app-name}/main.go`

**Go Best Practices**:
- Reference: `.claude/go/best-practices.md` - Code Organization
- Use `cmd/` for applications
- Use `internal/` for private code
- Use `pkg/` for public libraries

---

#### Task 1.2: Define Core Types and Interfaces

**Dependencies**: Task 1.1
**Status**: Not Started

**Description**:
Define the core data types, interfaces, and constants that will be used throughout the application.

**Implementation Steps**:
1. Create `pkg/{package}/types.go`
2. Define main data structures
3. Create `pkg/{package}/interfaces.go`
4. Define key interfaces
5. Create `pkg/{package}/errors.go`
6. Define custom error types and sentinel errors

**Acceptance Criteria**:
- [ ] Core types defined with proper documentation
- [ ] Interfaces defined following single-responsibility principle
- [ ] Custom error types created
- [ ] Sentinel errors defined as package-level variables
- [ ] All types have godoc comments

**Files to Create**:
- `pkg/{package}/types.go`
- `pkg/{package}/interfaces.go`
- `pkg/{package}/errors.go`

**Go Best Practices**:
- Reference: `.claude/go/idiomatic-go.md` - Interfaces for Behavior
- Reference: `.claude/go/error-handling.md` - Custom Errors
- Keep interfaces small (single method if possible)
- Use value types for small structs, pointers for large ones
- Define errors at package level

**Example**:
```go
// types.go
type Config struct {
    Host    string
    Port    int
    Timeout time.Duration
}

// interfaces.go
type Service interface {
    Process(ctx context.Context, data Data) (Result, error)
}

// errors.go
var (
    ErrNotFound   = errors.New("not found")
    ErrInvalidInput = errors.New("invalid input")
)

type ValidationError struct {
    Field  string
    Reason string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed: %s: %s", e.Field, e.Reason)
}
```

---

### Phase 2: Core Implementation

#### Task 2.1: Implement {Core Component}

**Dependencies**: Task 1.2
**Status**: Not Started

**Description**:
{Detailed description of what this component does}

**Implementation Steps**:
1. {Step 1}
2. {Step 2}
3. {Step 3}

**Acceptance Criteria**:
- [ ] {Criterion 1}
- [ ] {Criterion 2}
- [ ] Error handling implemented
- [ ] Logging added at appropriate levels

**Files to Create/Modify**:
- `pkg/{package}/{file}.go`

**Go Best Practices**:
- Reference: `.claude/go/{relevant-doc}.md`
- {Specific practice}

---

#### Task 2.2: Implement {Another Component}

**Dependencies**: Task 2.1
**Status**: Not Started

**Description**:
{Description}

**Implementation Steps**:
1. {Step 1}
2. {Step 2}

**Acceptance Criteria**:
- [ ] {Criterion 1}
- [ ] {Criterion 2}

**Files to Create/Modify**:
- `{file path}`

**Go Best Practices**:
- Reference: `.claude/go/{relevant-doc}.md`

---

### Phase 3: Documentation & Deployment

#### Task 3.1: Add Package Documentation

**Dependencies**: Phase 2
**Status**: Not Started

**Description**:
Add comprehensive documentation to all packages, types, and functions.

**Implementation Steps**:
1. Add package comments to each package
2. Add godoc comments to all exported types
3. Add godoc comments to all exported functions
4. Add example tests for public APIs
5. Generate and review godoc locally

**Acceptance Criteria**:
- [ ] Every package has package comment
- [ ] Every exported type has comment
- [ ] Every exported function has comment
- [ ] Examples provided for key APIs
- [ ] Godoc renders correctly

**Files to Modify**:
- All `.go` files with exports
- Add `example_test.go` files

**Go Best Practices**:
- Reference: `.claude/go/best-practices.md` - Documentation
- Start comments with the name being documented
- Write complete sentences
- Provide examples in `Example` functions

**Example**:
```go
// Package validator provides cryptocurrency address validation.
//
// This package validates addresses for multiple cryptocurrencies
// using external validation libraries.
//
// Example usage:
//
//     v := validator.New()
//     valid, err := v.Validate("BTC", "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa")
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Println(valid) // true
//
package validator

// Validator validates cryptocurrency addresses.
type Validator struct {
    // ...
}

// New creates a new Validator instance.
func New() *Validator {
    // ...
}

// Validate checks if an address is valid for the given currency.
// It returns true if the address is valid, false otherwise.
// An error is returned if the currency is not supported.
func (v *Validator) Validate(currency, address string) (bool, error) {
    // ...
}
```

---

#### Task 3.2: Update README

**Dependencies**: Task 3.1
**Status**: Not Started

**Description**:
Create comprehensive README with installation, usage, and examples.

**Implementation Steps**:
1. Write project overview
2. Add installation instructions
3. Add usage examples
4. Document CLI flags/options
5. Add development setup instructions

**Acceptance Criteria**:
- [ ] README has clear overview
- [ ] Installation instructions provided
- [ ] Usage examples included
- [ ] Development setup documented
- [ ] Contributing guidelines (if applicable)

**Files to Modify**:
- `README.md`

**README Structure**:
```markdown
# {Project Name}

{Brief description}

## Features

- Feature 1
- Feature 2

## Installation

```bash
go install {module-path}/cmd/{app-name}@latest
```

## Usage

```bash
{app-name} [flags]
```

### Examples

...

## Development

### Prerequisites

- Go 1.21+

### Building

```bash
make build
```

## Contributing

...

## License

{License}
```

---

#### Task 3.3: Create Makefile

**Dependencies**: None (can be done early)
**Status**: Not Started

**Description**:
Create Makefile with common development tasks.

**Implementation Steps**:
1. Create Makefile
2. Add build target
3. Add lint target
4. Add clean target

**Acceptance Criteria**:
- [ ] `make build` compiles the project
- [ ] `make lint` runs linters
- [ ] `make clean` removes build artifacts
- [ ] `make help` shows available targets

**Files to Create**:
- `Makefile`

**Example**:
```makefile
.PHONY: build lint clean help

APP_NAME := {app-name}
BUILD_DIR := ./build

build: ## Build the application
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/$(APP_NAME)

lint: ## Run linters
	golangci-lint run

fmt: ## Format code
	gofmt -s -w .
	go mod tidy

clean: ## Clean build artifacts
	rm -rf $(BUILD_DIR)

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
```

---

## Go Best Practices to Follow

### Code Organization
- **Reference**: `.claude/go/best-practices.md` - Code Organization
- Use `cmd/` for main applications
- Use `internal/` for private packages
- Use `pkg/` for public libraries
- One package per directory

### Naming Conventions
- **Reference**: `.claude/go/best-practices.md` - Naming Conventions
- Short variable names in small scopes
- Descriptive names in larger scopes
- Interfaces end in `-er` for single methods
- No underscores in package names

### Error Handling
- **Reference**: `.claude/go/error-handling.md`
- Always check errors
- Wrap errors with context using `fmt.Errorf` with `%w`
- Define sentinel errors at package level
- Create custom error types for rich information

### Concurrency
- **Reference**: `.claude/go/concurrency-patterns.md`
- Use goroutines for concurrent work
- Use channels for communication
- Always ensure goroutines can exit
- Use context for cancellation

### Common Mistakes to Avoid
- **Reference**: `.claude/go/common-mistakes.md`
- Don't shadow variables with `:=`
- Don't ignore errors
- Don't leak goroutines
- Don't copy mutexes
- Close channels from sender side only

### Design Patterns
- **Reference**: `.claude/go/design-patterns.md`
- Use functional options for configuration
- Accept interfaces, return concrete types
- Use context for cancellation and timeouts
- Prefer composition over inheritance (embedding)

---

## Validation Checklist

Before marking implementation complete, verify:

### Code Quality
- [ ] All code formatted with `gofmt`
- [ ] No linter warnings with `golangci-lint`
- [ ] All functions have godoc comments
- [ ] Code follows Go idioms from `.claude/go/idiomatic-go.md`

### Error Handling
- [ ] All errors checked and handled
- [ ] Errors wrapped with context
- [ ] Custom errors for domain-specific cases
- [ ] Error messages are descriptive

### Concurrency
- [ ] Goroutines can be cancelled
- [ ] Channels closed properly
- [ ] Mutexes not copied

### Documentation
- [ ] README is complete and accurate
- [ ] Package documentation exists
- [ ] Public APIs documented
- [ ] Examples provided

### Performance
- [ ] No obvious performance issues
- [ ] Memory allocations reasonable

### Security
- [ ] Input validation implemented
- [ ] No hardcoded secrets
- [ ] Proper use of crypto/rand (if needed)
- [ ] SQL injection prevented (if applicable)

---

## Dependencies

### Internal Dependencies
- {Dependency 1}: {Why needed}
- {Dependency 2}: {Why needed}

### External Dependencies
- `{module}` {version}: {Purpose}
- `{module}` {version}: {Purpose}

To add external dependencies:
```bash
go get {module}@{version}
```

---

## Risks & Mitigations

| Risk | Impact | Likelihood | Mitigation |
|------|--------|------------|------------|
| {Risk 1} | {High/Med/Low} | {High/Med/Low} | {Mitigation strategy} |
| {Risk 2} | {High/Med/Low} | {High/Med/Low} | {Mitigation strategy} |

---

## Notes

### Technical Decisions

1. **{Decision 1}**: {Rationale and alternatives considered}
2. **{Decision 2}**: {Rationale and alternatives considered}

### Future Enhancements

- {Enhancement 1}
- {Enhancement 2}

### References

- Feature Documentation: [FEATURE.md](./FEATURE.md)
- Go Resources: `.claude/go/`
- {External resource 1}
- {External resource 2}

---

## Change Log

| Date | Author | Changes |
|------|--------|---------|
| {Date} | feature-go-techlead | Initial TODO.md created |
| {Date} | {Name} | {Changes} |
