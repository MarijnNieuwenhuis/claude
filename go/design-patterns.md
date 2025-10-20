# Go Design Patterns

Common design patterns and their idiomatic Go implementations.

## Table of Contents

- [Creational Patterns](#creational-patterns)
- [Structural Patterns](#structural-patterns)
- [Behavioral Patterns](#behavioral-patterns)
- [Concurrency Patterns](#concurrency-patterns)
- [Go-Specific Patterns](#go-specific-patterns)

---

## Creational Patterns

### Builder Pattern

Build complex objects step by step.

```go
type Server struct {
    host    string
    port    int
    timeout time.Duration
    maxConn int
    tls     *tls.Config
}

type ServerBuilder struct {
    server *Server
}

func NewServerBuilder(host string) *ServerBuilder {
    return &ServerBuilder{
        server: &Server{
            host:    host,
            port:    8080,      // Defaults
            timeout: 30 * time.Second,
            maxConn: 100,
        },
    }
}

func (b *ServerBuilder) Port(port int) *ServerBuilder {
    b.server.port = port
    return b
}

func (b *ServerBuilder) Timeout(d time.Duration) *ServerBuilder {
    b.server.timeout = d
    return b
}

func (b *ServerBuilder) MaxConnections(n int) *ServerBuilder {
    b.server.maxConn = n
    return b
}

func (b *ServerBuilder) Build() *Server {
    return b.server
}

// Usage
server := NewServerBuilder("localhost").
    Port(9000).
    Timeout(60 * time.Second).
    MaxConnections(200).
    Build()
```

**Go Idiom**: Use functional options instead:

```go
type Option func(*Server)

func WithPort(port int) Option {
    return func(s *Server) {
        s.port = port
    }
}

func WithTimeout(d time.Duration) Option {
    return func(s *Server) {
        s.timeout = d
    }
}

func NewServer(host string, opts ...Option) *Server {
    s := &Server{
        host:    host,
        port:    8080,
        timeout: 30 * time.Second,
    }
    for _, opt := range opts {
        opt(s)
    }
    return s
}

// Usage
server := NewServer("localhost",
    WithPort(9000),
    WithTimeout(60*time.Second),
)
```

### Factory Pattern

Create objects without specifying exact class.

```go
type Storage interface {
    Save(key string, value []byte) error
    Load(key string) ([]byte, error)
}

type MemoryStorage struct {
    data map[string][]byte
}

func (m *MemoryStorage) Save(key string, value []byte) error {
    m.data[key] = value
    return nil
}

func (m *MemoryStorage) Load(key string) ([]byte, error) {
    return m.data[key], nil
}

type FileStorage struct {
    basePath string
}

func (f *FileStorage) Save(key string, value []byte) error {
    return os.WriteFile(filepath.Join(f.basePath, key), value, 0644)
}

func (f *FileStorage) Load(key string) ([]byte, error) {
    return os.ReadFile(filepath.Join(f.basePath, key))
}

// Factory function
func NewStorage(storageType string) Storage {
    switch storageType {
    case "memory":
        return &MemoryStorage{data: make(map[string][]byte)}
    case "file":
        return &FileStorage{basePath: "/tmp/storage"}
    default:
        return &MemoryStorage{data: make(map[string][]byte)}
    }
}
```

### Singleton Pattern

Ensure only one instance exists.

```go
type Database struct {
    conn *sql.DB
}

var (
    instance *Database
    once     sync.Once
)

func GetDatabase() *Database {
    once.Do(func() {
        conn, _ := sql.Open("postgres", "connection-string")
        instance = &Database{conn: conn}
    })
    return instance
}
```

**Note**: Singletons are often an anti-pattern. Prefer dependency injection:

```go
// Better: Pass dependencies explicitly
type Service struct {
    db *Database
}

func NewService(db *Database) *Service {
    return &Service{db: db}
}
```

### Object Pool Pattern

Reuse expensive objects.

```go
import "sync"

type Connection struct {
    // ... connection fields
}

var connPool = sync.Pool{
    New: func() interface{} {
        return &Connection{}  // Create new connection
    },
}

func GetConnection() *Connection {
    return connPool.Get().(*Connection)
}

func PutConnection(conn *Connection) {
    conn.Reset()  // Clean up
    connPool.Put(conn)
}

// Usage
conn := GetConnection()
defer PutConnection(conn)
// Use connection...
```

---

## Structural Patterns

### Adapter Pattern

Convert one interface to another.

```go
// Legacy interface
type LegacyPrinter interface {
    Print(s string) error
}

// New interface
type ModernPrinter interface {
    PrintLine(s string) error
}

// Adapter
type PrinterAdapter struct {
    modernPrinter ModernPrinter
}

func (a *PrinterAdapter) Print(s string) error {
    return a.modernPrinter.PrintLine(s)
}

// Usage
modernPrinter := &SomeModernPrinter{}
legacyPrinter := &PrinterAdapter{modernPrinter: modernPrinter}
```

### Decorator Pattern

Add behavior to objects dynamically.

```go
type Handler interface {
    Handle(req *Request) (*Response, error)
}

// Base handler
type BaseHandler struct{}

func (h *BaseHandler) Handle(req *Request) (*Response, error) {
    // Handle request
    return &Response{}, nil
}

// Logging decorator
type LoggingHandler struct {
    handler Handler
}

func (h *LoggingHandler) Handle(req *Request) (*Response, error) {
    log.Printf("handling request: %v", req)
    resp, err := h.handler.Handle(req)
    log.Printf("response: %v, error: %v", resp, err)
    return resp, err
}

// Timing decorator
type TimingHandler struct {
    handler Handler
}

func (h *TimingHandler) Handle(req *Request) (*Response, error) {
    start := time.Now()
    resp, err := h.handler.Handle(req)
    log.Printf("request took: %v", time.Since(start))
    return resp, err
}

// Usage: Stack decorators
handler := &TimingHandler{
    handler: &LoggingHandler{
        handler: &BaseHandler{},
    },
}
```

### Proxy Pattern

Control access to an object.

```go
type Database interface {
    Query(sql string) ([]Row, error)
}

// Real database
type RealDatabase struct {
    conn *sql.DB
}

func (db *RealDatabase) Query(sql string) ([]Row, error) {
    // Execute query
    return nil, nil
}

// Caching proxy
type CachingDatabaseProxy struct {
    db    Database
    cache map[string][]Row
    mu    sync.RWMutex
}

func (p *CachingDatabaseProxy) Query(sql string) ([]Row, error) {
    p.mu.RLock()
    if rows, ok := p.cache[sql]; ok {
        p.mu.RUnlock()
        return rows, nil
    }
    p.mu.RUnlock()

    rows, err := p.db.Query(sql)
    if err != nil {
        return nil, err
    }

    p.mu.Lock()
    p.cache[sql] = rows
    p.mu.Unlock()

    return rows, nil
}
```

---

## Behavioral Patterns

### Strategy Pattern

Define a family of algorithms.

```go
type CompressionStrategy interface {
    Compress(data []byte) ([]byte, error)
}

type GzipCompression struct{}

func (g *GzipCompression) Compress(data []byte) ([]byte, error) {
    var buf bytes.Buffer
    w := gzip.NewWriter(&buf)
    _, err := w.Write(data)
    w.Close()
    return buf.Bytes(), err
}

type ZlibCompression struct{}

func (z *ZlibCompression) Compress(data []byte) ([]byte, error) {
    var buf bytes.Buffer
    w := zlib.NewWriter(&buf)
    _, err := w.Write(data)
    w.Close()
    return buf.Bytes(), err
}

type Compressor struct {
    strategy CompressionStrategy
}

func (c *Compressor) Compress(data []byte) ([]byte, error) {
    return c.strategy.Compress(data)
}

// Usage
compressor := &Compressor{strategy: &GzipCompression{}}
compressed, _ := compressor.Compress(data)

// Change strategy
compressor.strategy = &ZlibCompression{}
```

### Observer Pattern

Notify multiple objects of state changes.

```go
type Event struct {
    Type string
    Data interface{}
}

type Observer interface {
    Notify(event Event)
}

type Subject struct {
    observers []Observer
}

func (s *Subject) Register(observer Observer) {
    s.observers = append(s.observers, observer)
}

func (s *Subject) NotifyAll(event Event) {
    for _, observer := range s.observers {
        observer.Notify(event)
    }
}

// Concrete observer
type Logger struct{}

func (l *Logger) Notify(event Event) {
    log.Printf("Event: %s, Data: %v", event.Type, event.Data)
}

// Usage
subject := &Subject{}
subject.Register(&Logger{})
subject.NotifyAll(Event{Type: "user.created", Data: user})
```

**Go Idiom**: Use channels instead:

```go
type EventBus struct {
    subscribers []chan Event
}

func (eb *EventBus) Subscribe() chan Event {
    ch := make(chan Event, 10)
    eb.subscribers = append(eb.subscribers, ch)
    return ch
}

func (eb *EventBus) Publish(event Event) {
    for _, ch := range eb.subscribers {
        ch <- event
    }
}

// Usage
bus := &EventBus{}
events := bus.Subscribe()

go func() {
    for event := range events {
        log.Printf("Received: %v", event)
    }
}()

bus.Publish(Event{Type: "user.created"})
```

### Command Pattern

Encapsulate requests as objects.

```go
type Command interface {
    Execute() error
    Undo() error
}

type CreateUserCommand struct {
    db   *Database
    user *User
}

func (c *CreateUserCommand) Execute() error {
    return c.db.Insert(c.user)
}

func (c *CreateUserCommand) Undo() error {
    return c.db.Delete(c.user.ID)
}

type CommandQueue struct {
    commands []Command
}

func (q *CommandQueue) Add(cmd Command) {
    q.commands = append(q.commands, cmd)
}

func (q *CommandQueue) ExecuteAll() error {
    for _, cmd := range q.commands {
        if err := cmd.Execute(); err != nil {
            return err
        }
    }
    return nil
}
```

---

## Concurrency Patterns

### Worker Pool

Process tasks concurrently with limited workers.

```go
func WorkerPool(numWorkers int, jobs <-chan Job, results chan<- Result) {
    var wg sync.WaitGroup

    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for job := range jobs {
                result := processJob(job)
                results <- result
            }
        }()
    }

    wg.Wait()
    close(results)
}

// Usage
jobs := make(chan Job, 100)
results := make(chan Result, 100)

go WorkerPool(5, jobs, results)

// Send jobs
for _, job := range allJobs {
    jobs <- job
}
close(jobs)

// Collect results
for result := range results {
    fmt.Println(result)
}
```

### Pipeline

Chain processing stages.

```go
func generator(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func sum(in <-chan int) int {
    total := 0
    for n := range in {
        total += n
    }
    return total
}

// Usage
nums := generator(1, 2, 3, 4, 5)
squared := square(nums)
total := sum(squared)
fmt.Println(total)  // 55
```

### Fan-Out/Fan-In

Distribute work, then collect results.

```go
func fanOut(in <-chan int, n int) []<-chan int {
    outs := make([]<-chan int, n)
    for i := 0; i < n; i++ {
        outs[i] = worker(in)
    }
    return outs
}

func fanIn(channels ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup

    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for n := range c {
                out <- n
            }
        }(ch)
    }

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}

// Usage
in := generator(1, 2, 3, 4, 5)
workers := fanOut(in, 3)  // 3 workers
results := fanIn(workers...)
```

---

## Go-Specific Patterns

### Context Pattern

Carry deadlines, cancellation, and values.

```go
func DoWork(ctx context.Context, data Data) error {
    // Check for cancellation
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
    }

    // Do work with timeout
    resultCh := make(chan Result)
    go func() {
        resultCh <- processData(data)
    }()

    select {
    case <-ctx.Done():
        return ctx.Err()
    case result := <-resultCh:
        return handleResult(result)
    }
}

// Usage
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

if err := DoWork(ctx, data); err != nil {
    log.Printf("work failed: %v", err)
}
```

### Embedding Pattern

Compose types via embedding.

```go
// Base functionality
type Logger struct {
    prefix string
}

func (l *Logger) Log(msg string) {
    fmt.Printf("[%s] %s\n", l.prefix, msg)
}

// Extend with embedding
type Service struct {
    Logger  // Embedded
    db *Database
}

// Service inherits Log method
service := &Service{
    Logger: Logger{prefix: "SERVICE"},
}
service.Log("started")  // Uses embedded Logger's method
```

### Options Pattern

Configure objects flexibly (mentioned in Builder, worth emphasizing).

```go
type Config struct {
    host    string
    port    int
    timeout time.Duration
}

type Option func(*Config)

func WithHost(host string) Option {
    return func(c *Config) {
        c.host = host
    }
}

func WithPort(port int) Option {
    return func(c *Config) {
        c.port = port
    }
}

func WithTimeout(d time.Duration) Option {
    return func(c *Config) {
        c.timeout = d
    }
}

func NewConfig(opts ...Option) *Config {
    cfg := &Config{
        host:    "localhost",
        port:    8080,
        timeout: 30 * time.Second,
    }
    for _, opt := range opts {
        opt(cfg)
    }
    return cfg
}
```

---

## When NOT to Use Patterns

Go favors simplicity. Don't use patterns if:
- A simple function suffices
- The pattern adds unnecessary complexity
- Built-in language features solve the problem
- You're "patternizing" for the sake of it

**Remember**: In Go, simple and clear beats clever and abstract.

---

## References

- [Go Patterns](https://github.com/tmrts/go-patterns)
- [Design Patterns in Go](https://refactoring.guru/design-patterns/go)
- [Go Wiki: Common Mistakes](https://github.com/golang/go/wiki/CommonMistakes)
