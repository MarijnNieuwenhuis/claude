# Idiomatic Go

Writing Go code the "Go way" - philosophy, patterns, and idioms that make code truly idiomatic.

## Core Philosophy

### Simplicity Over Cleverness

Go values clarity and maintainability over clever solutions.

```go
// Good: Simple and clear
func IsAdult(age int) bool {
    return age >= 18
}

// Bad: Clever but obscure
func IsAdult(age int) bool {
    return ^uint(age-18)>>63 == 0  // Bit manipulation trick
}
```

### Explicit Over Implicit

Go prefers explicit code that's easy to follow.

```go
// Good: Explicit error handling
result, err := DoSomething()
if err != nil {
    return fmt.Errorf("do something: %w", err)
}

// Bad (in other languages): Exceptions hiding flow
try {
    result = DoSomething()  // Where does this error go?
}
```

### Composition Over Inheritance

Go has no inheritance; use composition and interfaces.

```go
// Good: Composition
type Logger struct {
    output io.Writer
}

type Service struct {
    logger Logger
    db     Database
}

// Not possible in Go: Inheritance
// type Service extends BaseService {}
```

---

## Idiomatic Patterns

### Early Returns

Exit early to reduce nesting.

```go
// Good: Early returns, happy path clear
func ProcessUser(id int) error {
    if id <= 0 {
        return errors.New("invalid ID")
    }

    user, err := FetchUser(id)
    if err != nil {
        return err
    }

    if !user.IsActive {
        return errors.New("user not active")
    }

    return SaveUser(user)
}

// Bad: Nested if statements
func ProcessUser(id int) error {
    if id > 0 {
        user, err := FetchUser(id)
        if err == nil {
            if user.IsActive {
                return SaveUser(user)
            } else {
                return errors.New("user not active")
            }
        } else {
            return err
        }
    } else {
        return errors.New("invalid ID")
    }
}
```

### Zero Values

Design types so zero values are useful.

```go
// Good: Zero value is valid
type Buffer struct {
    data []byte
}

// Can use without initialization
var buf Buffer
buf.Write([]byte("hello"))  // Works!

// Also good: Constructor for complex initialization
func NewClient(url string) *Client {
    return &Client{
        url:     url,
        timeout: 30 * time.Second,  // Default timeout
        client:  &http.Client{},
    }
}
```

### Interfaces for Behavior

Use small interfaces for behavior, not data.

```go
// Good: Interface defines behavior
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Good: Small, focused interface
type Closer interface {
    Close() error
}

// Combine interfaces
type ReadCloser interface {
    Reader
    Closer
}

// Bad: Large interface
type DataStore interface {
    Get(key string) (interface{}, error)
    Set(key string, value interface{}) error
    Delete(key string) error
    List() ([]string, error)
    Count() int
    Clear() error
    // ... many more methods
}
```

### Accept Interfaces, Return Concrete Types

```go
// Good: Accept interface (flexible for callers)
func Save(w io.Writer, data []byte) error {
    _, err := w.Write(data)
    return err
}

// Good: Return concrete type (clear for callers)
func LoadConfig(filename string) (*Config, error) {
    // ... load config
    return &Config{}, nil
}

// Bad: Returning interface (adds unnecessary abstraction)
func LoadConfig(filename string) (Configurer, error) {
    return &Config{}, nil
}
```

### Short Variable Names

Use short names in small scopes, longer in larger scopes.

```go
// Good: Short names in small scope
for i, v := range items {
    fmt.Println(i, v)
}

// Good: Longer names in larger scope
func ProcessUserRegistration(email, password string) error {
    userRepository := NewUserRepository(db)
    validationService := NewValidationService()

    if err := validationService.ValidateEmail(email); err != nil {
        return err
    }
    // ... more code
}

// Good: Single-letter receivers
func (u *User) Save() error {}
func (c *Client) Get(url string) (*Response, error) {}
```

### Line of Sight

Keep happy path at minimal indentation.

```go
// Good: Happy path has least indentation
func Process(items []Item) error {
    if len(items) == 0 {
        return errors.New("no items")
    }

    for _, item := range items {
        if !item.IsValid() {
            continue
        }

        if err := item.Process(); err != nil {
            return err
        }

        item.MarkProcessed()
    }

    return nil
}

// Bad: Happy path deeply nested
func Process(items []Item) error {
    if len(items) > 0 {
        for _, item := range items {
            if item.IsValid() {
                if err := item.Process(); err == nil {
                    item.MarkProcessed()
                } else {
                    return err
                }
            }
        }
    } else {
        return errors.New("no items")
    }
    return nil
}
```

---

## Idiomatic Techniques

### Functional Options

Pattern for optional configuration.

```go
type Server struct {
    addr    string
    timeout time.Duration
    maxConn int
}

// Option is a functional option
type Option func(*Server)

func WithTimeout(d time.Duration) Option {
    return func(s *Server) {
        s.timeout = d
    }
}

func WithMaxConn(n int) Option {
    return func(s *Server) {
        s.maxConn = n
    }
}

// Constructor accepts options
func NewServer(addr string, opts ...Option) *Server {
    s := &Server{
        addr:    addr,
        timeout: 30 * time.Second,  // Defaults
        maxConn: 100,
    }

    for _, opt := range opts {
        opt(s)
    }

    return s
}

// Usage
server := NewServer("localhost:8080",
    WithTimeout(60*time.Second),
    WithMaxConn(200),
)
```

### Empty Interface for Generic Data

Use `interface{}` (or `any` in Go 1.18+) sparingly.

```go
// Good: Use when truly generic
func PrintAny(v interface{}) {
    fmt.Println(v)
}

// Better: Use generics (Go 1.18+)
func Print[T any](v T) {
    fmt.Println(v)
}

// Bad: Overuse loses type safety
type Cache map[string]interface{}  // What types are stored?

// Better: Specific types
type UserCache map[string]*User
type ConfigCache map[string]*Config
```

### Struct Embedding

Compose types by embedding.

```go
// Embed to inherit methods
type Reader struct {
    r io.Reader
}

func (r *Reader) Read(p []byte) (n int, err error) {
    return r.r.Read(p)
}

// Better: Direct embedding (delegates automatically)
type Reader struct {
    io.Reader  // Embeds the interface
}

// Now Reader automatically implements io.Reader!

// Example with structs
type Base struct {
    ID   int
    Name string
}

type User struct {
    Base         // Embedded
    Email string
}

// Can access Base fields directly
user := User{
    Base:  Base{ID: 1, Name: "Alice"},
    Email: "alice@example.com",
}
fmt.Println(user.ID)    // Direct access to embedded field
fmt.Println(user.Name)  // Direct access to embedded field
```

### Blank Identifier for Side Effects

Import for init side effects.

```go
// Import database driver for its side effects
import _ "github.com/lib/pq"

// Ignore values explicitly
_, err := fmt.Println("hello")
if err != nil {
    return err
}

// Compile-time interface check
var _ io.Reader = (*MyReader)(nil)  // Ensures MyReader implements io.Reader
```

### Type Switch

Handle different types idiomatically.

```go
func Handle(v interface{}) {
    switch v := v.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

---

## Idiomatic Error Handling

### Error Wrapping

Add context while preserving original error.

```go
// Good: Wrap errors with context
if err := saveToDatabase(user); err != nil {
    return fmt.Errorf("save user %d to database: %w", user.ID, err)
}

// Can unwrap later
if errors.Is(err, sql.ErrNoRows) {
    // Handle specific error
}
```

### Sentinel Errors

Define package-level errors for known cases.

```go
var (
    ErrNotFound   = errors.New("not found")
    ErrInvalidID  = errors.New("invalid ID")
    ErrUnauthorized = errors.New("unauthorized")
)

func FetchUser(id int) (*User, error) {
    if id <= 0 {
        return nil, ErrInvalidID
    }

    user, err := db.Query(...)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrNotFound
        }
        return nil, err
    }

    return user, nil
}

// Caller can check
user, err := FetchUser(id)
if errors.Is(err, ErrNotFound) {
    // Handle not found case
}
```

### Custom Error Types

For rich error information.

```go
type ValidationError struct {
    Field string
    Value interface{}
    Reason string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed for %s=%v: %s",
        e.Field, e.Value, e.Reason)
}

func Validate(user *User) error {
    if user.Age < 0 {
        return &ValidationError{
            Field:  "age",
            Value:  user.Age,
            Reason: "must be non-negative",
        }
    }
    return nil
}

// Caller can type assert
if err := Validate(user); err != nil {
    var ve *ValidationError
    if errors.As(err, &ve) {
        fmt.Printf("Field %s failed: %s\n", ve.Field, ve.Reason)
    }
}
```

---

## Idiomatic Concurrency

### Use Channels for Communication

Channels over shared memory.

```go
// Good: Communicate via channels
func Worker(jobs <-chan Job, results chan<- Result) {
    for job := range jobs {
        result := process(job)
        results <- result
    }
}

// Start workers
jobs := make(chan Job, 100)
results := make(chan Result, 100)

for w := 0; w < 5; w++ {
    go Worker(jobs, results)
}
```

### Don't Leak Goroutines

Ensure goroutines can exit.

```go
// Good: Context for cancellation
func Stream(ctx context.Context, results chan<- Data) {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ctx.Done():
            return  // Exit when cancelled
        case <-ticker.C:
            results <- fetchData()
        }
    }
}

// Usage
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

results := make(chan Data)
go Stream(ctx, results)

// Cancel when done
cancel()  // Goroutine exits
```

---

## Idiomatic Testing

### Table-Driven Tests

Test multiple cases efficiently.

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 1, 2, 3},
        {"negative", -1, -2, -3},
        {"zero", 0, 0, 0},
        {"mixed", -1, 2, 1},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("Add(%d, %d) = %d, want %d",
                    tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

### Test Helpers

Use `t.Helper()` for test utilities.

```go
func assertNoError(t *testing.T, err error) {
    t.Helper()  // Marks this as helper
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
}

func TestSomething(t *testing.T) {
    err := DoSomething()
    assertNoError(t, err)  // Failure reports correct line
}
```

---

## Summary: The Go Way

1. **Keep it simple** - Clarity over cleverness
2. **Be explicit** - No magic, clear control flow
3. **Use composition** - Build from small pieces
4. **Handle errors** - Explicit error returns
5. **Minimize nesting** - Early returns, flat structure
6. **Small interfaces** - Focused behavior contracts
7. **Useful zero values** - Design for initialization
8. **Communicate via channels** - Share by communicating
9. **Document** - Clear comments and names
10. **Format with gofmt** - Consistent style

---

## References

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Proverbs](https://go-proverbs.github.io/)
- [How to Write Go Code](https://go.dev/doc/code)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
