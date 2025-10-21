# Go Best Practices

Essential coding standards and conventions for writing maintainable Go code.

## Table of Contents

- [Code Organization](#code-organization)
- [Naming Conventions](#naming-conventions)
- [Error Handling](#error-handling)
- [Code Formatting](#code-formatting)
- [Package Design](#package-design)
- [Documentation](#documentation)
- [Performance](#performance)
- [Security](#security)

---

## Code Organization

### Project Structure

```
project/
├── cmd/                    # Application entry points
│   └── myapp/
│       └── main.go
├── internal/               # Private application code
│   ├── service/
│   └── repository/
├── pkg/                    # Public library code
│   └── api/
├── go.mod                  # Module definition
├── go.sum                  # Dependency checksums
└── README.md
```

**Guidelines**:
- Use `cmd/` for main applications
- Use `internal/` for private code (cannot be imported by other projects)
- Use `pkg/` for public library code
- Keep `main` packages small - delegate to packages

### File Organization

- One package per directory
- Group related functionality in same package
- Keep files focused and reasonably sized (< 500 lines)
- Use `_test.go` suffix for test files

**Example**:
```
user/
├── user.go          # Core types and logic
├── user_test.go     # Tests
├── handler.go       # HTTP handlers
├── repository.go    # Data access
└── service.go       # Business logic
```

---

## Naming Conventions

### General Principles

- **Short and concise** - Go favors brevity
- **Descriptive** - Clear purpose from name alone
- **Context-aware** - Use context to keep names short

### Variables

**Good**:
```go
// Short-lived variables
for i, v := range values {
    fmt.Println(i, v)
}

// Single letter for receivers
func (u *User) Save() error {}

// Descriptive for longer scope
var userRepository UserRepository
```

**Bad**:
```go
// Too verbose
var userDataFromDatabaseRepository UserRepository

// Unclear single letters in large scope
var u UserRepository  // What is 'u'?
```

### Functions

**Good**:
```go
// Starts with verb
func FetchUser(id int) (*User, error) {}
func SaveOrder(order Order) error {}

// Getter doesn't use Get prefix
func (u *User) Name() string {}  // Not GetName()
```

**Bad**:
```go
func User(id int) (*User, error) {}  // Noun, not verb
func GetUserDataFromDatabase(id int) {}  // Too verbose
```

### Constants

```go
// Use MixedCaps or mixedCaps, not underscores
const MaxRetries = 3
const defaultTimeout = 10 * time.Second

// Group related constants
const (
    StatusPending  Status = "pending"
    StatusActive   Status = "active"
    StatusInactive Status = "inactive"
)
```

### Packages

```go
// Short, lowercase, no underscores
package http
package json
package user

// Not:
package httpServer  // No mixedCaps
package user_service  // No underscores
```

### Interfaces

```go
// Single-method interfaces end in -er
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

// Multi-method interfaces use descriptive names
type UserRepository interface {
    FindByID(id int) (*User, error)
    Save(user *User) error
    Delete(id int) error
}
```

---

## Error Handling

### Return Errors

```go
// Good: Return errors, don't panic
func FetchUser(id int) (*User, error) {
    if id <= 0 {
        return nil, errors.New("invalid user ID")
    }
    // ... fetch logic
    return user, nil
}

// Bad: Using panic for normal errors
func FetchUser(id int) *User {
    if id <= 0 {
        panic("invalid user ID")  // Don't do this
    }
    return user
}
```

### Error as Last Return Value

```go
// Good: Error is the last return value
func Process() (result Data, err error) {}

// Bad: Error not last
func Process() (err error, result Data) {}  // Don't do this
```

### Check Errors Immediately

```go
// Good: Check right away
user, err := FetchUser(id)
if err != nil {
    return fmt.Errorf("fetch user: %w", err)
}
// Use user...

// Bad: Deferred checking
user, err := FetchUser(id)
// ... other code
if err != nil {  // Too late, may have used invalid user
    return err
}
```

### Add Context to Errors

```go
// Good: Wrap with context
if err := SaveUser(user); err != nil {
    return fmt.Errorf("save user %d: %w", user.ID, err)
}

// Bad: Losing context
if err := SaveUser(user); err != nil {
    return err  // No context about what failed
}
```

### Don't Ignore Errors

```go
// Good: Handle explicitly
if err := file.Close(); err != nil {
    log.Printf("failed to close file: %v", err)
}

// Bad: Ignoring errors
file.Close()  // Error ignored

// If you must ignore, be explicit
_ = file.Close()  // Shows intentional ignore
```

---

## Code Formatting

### Use gofmt

```bash
# Format all Go files
gofmt -w .

# Check if files are formatted
gofmt -l .
```

**Rule**: All Go code must be formatted with `gofmt`. No exceptions.

### Line Length

- No strict limit, but prefer < 120 characters
- Break long lines logically

```go
// Good: Logical breaks
user, err := userRepository.FindByEmailAndStatus(
    email,
    StatusActive,
)

// Bad: One very long line
user, err := userRepository.FindByEmailAndStatus(email, StatusActive, WithTimeout(10*time.Second), IncludeDeleted(false))
```

### Indentation

- Use tabs for indentation (handled by gofmt)
- Don't mix tabs and spaces

---

## Package Design

### Package Scope

- **Small and focused** - Do one thing well
- **Clear purpose** - Obvious from name what it does
- **Minimal dependencies** - Avoid circular dependencies

### Exported vs Unexported

```go
// Exported (public) - starts with uppercase
type User struct {
    ID   int
    Name string
}

// Unexported (private) - starts with lowercase
type userCache struct {
    data map[int]*User
}

func FetchUser(id int) (*User, error) {}  // Exported
func validateUser(u *User) error {}      // Unexported
```

**Principle**: Export only what's necessary. Start unexported, export when needed.

### Avoid Package-Level State

```go
// Bad: Package-level mutable state
package cache

var data = make(map[string]interface{})  // Don't do this

func Get(key string) interface{} {
    return data[key]
}

// Good: Explicit instance
type Cache struct {
    data map[string]interface{}
}

func New() *Cache {
    return &Cache{
        data: make(map[string]interface{}),
    }
}

func (c *Cache) Get(key string) interface{} {
    return c.data[key]
}
```

---

## Documentation

### Doc Comments

```go
// Package comment appears before package declaration
// Package user provides user management functionality.
package user

// Type comments describe the type's purpose
// User represents a system user with credentials and profile.
type User struct {
    ID   int
    Name string
}

// Function comments start with function name
// FetchUser retrieves a user by ID from the database.
// It returns an error if the user is not found.
func FetchUser(id int) (*User, error) {
    // Implementation...
}
```

**Rules**:
- Start with the name of the thing being documented
- Write complete sentences
- Explain what, not how (code shows how)
- Document exported identifiers

### Package Documentation

```go
// For short package docs, use package comment:
// Package math provides basic math operations.
package math

// For long package docs, create doc.go:
// doc.go
/*
Package user provides comprehensive user management.

This package handles user authentication, authorization,
and profile management. It supports multiple authentication
methods including OAuth, SAML, and local credentials.

Basic usage:

    repo := user.NewRepository(db)
    svc := user.NewService(repo)
    user, err := svc.Authenticate(email, password)

For more details, see the individual type documentation.
*/
package user
```

---

## Performance

### Prefer Value Types

```go
// Good: Value types for small structs
type Point struct {
    X, Y int
}

func Translate(p Point, dx, dy int) Point {
    return Point{X: p.X + dx, Y: p.Y + dy}
}

// Good: Pointers for large structs or when mutations needed
type User struct {
    // ... many fields
}

func (u *User) Save() error {  // Pointer receiver
    // ... modify user
}
```

### Preallocate Slices

```go
// Good: Preallocate when size is known
users := make([]User, 0, 100)  // Capacity 100
for i := 0; i < 100; i++ {
    users = append(users, User{ID: i})
}

// Bad: Repeated reallocation
var users []User
for i := 0; i < 100; i++ {
    users = append(users, User{ID: i})  // Reallocates multiple times
}
```

### Use sync.Pool for Temporary Objects

```go
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func ProcessData(data []byte) {
    buf := bufferPool.Get().(*bytes.Buffer)
    buf.Reset()
    defer bufferPool.Put(buf)

    // Use buf...
}
```

### Avoid String Concatenation in Loops

```go
// Good: Use strings.Builder
var b strings.Builder
for _, s := range items {
    b.WriteString(s)
}
result := b.String()

// Bad: String concatenation
var result string
for _, s := range items {
    result += s  // Creates new string each iteration
}
```

---

## Security

### Validate Input

```go
// Good: Validate all external input
func CreateUser(email, password string) (*User, error) {
    if !isValidEmail(email) {
        return nil, errors.New("invalid email")
    }
    if len(password) < 8 {
        return nil, errors.New("password too short")
    }
    // ... create user
}
```

### Use Crypto Properly

```go
// Good: Use crypto/rand for random numbers
import "crypto/rand"

token := make([]byte, 32)
if _, err := rand.Read(token); err != nil {
    return err
}

// Bad: Using math/rand for security
import "math/rand"
token := rand.Int()  // Don't use for security!
```

### Avoid SQL Injection

```go
// Good: Use parameterized queries
result, err := db.Query(
    "SELECT * FROM users WHERE email = ?",
    email,
)

// Bad: String concatenation
query := "SELECT * FROM users WHERE email = '" + email + "'"  // SQL injection risk!
result, err := db.Query(query)
```

### Don't Log Sensitive Data

```go
// Good: Redact sensitive info
log.Printf("user authenticated: %s", user.Email)

// Bad: Logging passwords
log.Printf("login attempt: %s/%s", email, password)  // Don't log passwords!
```

---

## General Best Practices

### Accept Interfaces, Return Structs

```go
// Good: Accept interface (flexible)
func Save(w io.Writer, data []byte) error {
    _, err := w.Write(data)
    return err
}

// Good: Return concrete type (clear)
func NewUser(name string) *User {
    return &User{Name: name}
}
```

### Use defer for Cleanup

```go
// Good: Use defer for cleanup
func ProcessFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // Always closes, even on error

    // Process file...
    return nil
}
```

### Initialize Maps and Slices

```go
// Good: Initialize with make
m := make(map[string]int)
m["key"] = 1

// Bad: Nil map
var m map[string]int
m["key"] = 1  // Panic! Nil map
```

### Use Constants for Fixed Values

```go
// Good: Named constants
const (
    MaxRetries     = 3
    DefaultTimeout = 10 * time.Second
)

// Bad: Magic numbers
time.Sleep(10 * time.Second)  // What's 10?
```

---

## References

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Google Go Style Guide](https://google.github.io/styleguide/go/)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
