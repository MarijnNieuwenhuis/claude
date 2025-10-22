# Common Go Mistakes

Based on "100 Go Mistakes and How to Avoid Them" and community experience.

## Table of Contents

- [Code Organization](#code-organization)
- [Data Types](#data-types)
- [Control Structures](#control-structures)
- [Strings](#strings)
- [Functions & Methods](#functions--methods)
- [Error Handling](#error-handling)
- [Concurrency](#concurrency)
- [Standard Library](#standard-library)
- [Testing](#testing)

---

## Code Organization

### Mistake: Unintended Variable Shadowing

```go
// Bad: Shadowing with :=
var client *http.Client
if tracing {
    client, err := createClientWithTracing()  // New 'client' variable!
    if err != nil {
        return err
    }
    // ... use client
} else {
    client, err := createDefaultClient()  // Another new 'client'!
    if err != nil {
        return err
    }
}
// client is still nil here!

// Good: Assign to existing variable
var client *http.Client
var err error
if tracing {
    client, err = createClientWithTracing()  // Assigns to outer client
    if err != nil {
        return err
    }
} else {
    client, err = createDefaultClient()
    if err != nil {
        return err
    }
}
```

### Mistake: Unnecessary Nested Code

```go
// Bad: Deep nesting
func process(data []byte) error {
    if len(data) > 0 {
        if isValid(data) {
            result, err := transform(data)
            if err == nil {
                return save(result)
            } else {
                return err
            }
        } else {
            return errors.New("invalid data")
        }
    } else {
        return errors.New("empty data")
    }
}

// Good: Early returns
func process(data []byte) error {
    if len(data) == 0 {
        return errors.New("empty data")
    }
    if !isValid(data) {
        return errors.New("invalid data")
    }
    result, err := transform(data)
    if err != nil {
        return err
    }
    return save(result)
}
```

### Mistake: Init Functions Side Effects

```go
// Bad: init with side effects
var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("postgres", "...")  // What if this fails?
    if err != nil {
        panic(err)  // Application crashes at startup
    }
}

// Good: Explicit initialization
type Service struct {
    db *sql.DB
}

func NewService(dbURL string) (*Service, error) {
    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        return nil, fmt.Errorf("open database: %w", err)
    }
    return &Service{db: db}, nil
}
```

---

## Data Types

### Mistake: Not Understanding Slice Length vs Capacity

```go
// Slice length vs capacity
s1 := make([]int, 3, 5)  // length=3, capacity=5
fmt.Println(len(s1))  // 3
fmt.Println(cap(s1))  // 5
fmt.Println(s1)       // [0 0 0]

s2 := s1[1:3]  // Shares same underlying array!
s2[0] = 10
fmt.Println(s1)  // [0 10 0] - s1 changed too!

// To avoid sharing:
s3 := make([]int, len(s1))
copy(s3, s1)  // Deep copy
```

### Mistake: Slice Append Gotcha

```go
// Bad: Appending to subslice
s := []int{1, 2, 3, 4, 5}
s1 := s[1:3]  // [2, 3]
s1 = append(s1, 10)  // May overwrite s[3]!
fmt.Println(s)  // [1, 2, 3, 10, 5] - Original modified!

// Good: Full slice expression
s1 := s[1:3:3]  // Length 2, capacity 2
s1 = append(s1, 10)  // Forces new allocation
fmt.Println(s)  // [1, 2, 3, 4, 5] - Original unchanged
```

### Mistake: Map Iteration Order

```go
// Bad: Assuming map order
m := map[string]int{
    "a": 1,
    "b": 2,
    "c": 3,
}
for k, v := range m {
    fmt.Println(k, v)  // Order is random!
}

// Good: Sort keys if order matters
keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Strings(keys)
for _, k := range keys {
    fmt.Println(k, m[k])  // Deterministic order
}
```

### Mistake: Comparing Slices/Maps Directly

```go
// Bad: Can't use ==
s1 := []int{1, 2, 3}
s2 := []int{1, 2, 3}
// if s1 == s2 {}  // Compilation error!

// Good: Use reflect or manual comparison
if reflect.DeepEqual(s1, s2) {
    // Equal
}

// Better: Manual comparison
func equalSlices(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}
```

---

## Control Structures

### Mistake: Range Loop Variable Capture

```go
// Bad: Capturing loop variable
var funcs []func()
for _, v := range []int{1, 2, 3} {
    funcs = append(funcs, func() {
        fmt.Println(v)  // Captures loop variable!
    })
}
for _, f := range funcs {
    f()  // Prints: 3, 3, 3 (all print last value!)
}

// Good: Create new variable
for _, v := range []int{1, 2, 3} {
    v := v  // New variable for each iteration
    funcs = append(funcs, func() {
        fmt.Println(v)
    })
}
```

### Mistake: Break in Switch

```go
// Bad: break breaks switch, not loop
for i := 0; i < 10; i++ {
    switch i {
    case 5:
        break  // Only breaks switch, not loop!
    }
}

// Good: Use label
Loop:
for i := 0; i < 10; i++ {
    switch i {
    case 5:
        break Loop  // Breaks out of loop
    }
}

// Or: Use function return
func process() {
    for i := 0; i < 10; i++ {
        switch i {
        case 5:
            return  // Exits function
        }
    }
}
```

---

## Strings

### Mistake: Inefficient String Concatenation

```go
// Bad: String concatenation in loop
var s string
for i := 0; i < 1000; i++ {
    s += "a"  // Creates new string each time!
}

// Good: Use strings.Builder
var b strings.Builder
for i := 0; i < 1000; i++ {
    b.WriteString("a")
}
s := b.String()
```

### Mistake: String vs []byte

```go
// String is immutable, []byte is mutable
s := "hello"
// s[0] = 'H'  // Compilation error

b := []byte("hello")
b[0] = 'H'  // OK
s = string(b)  // Convert back

// Conversion copies data - be aware of performance
data := make([]byte, 1000000)
s := string(data)  // Copies 1MB
```

### Mistake: UTF-8 Rune Handling

```go
s := "hello 世界"
fmt.Println(len(s))  // 12 bytes, not 8 characters!

// Bad: Indexing assumes ASCII
fmt.Println(s[0])  // 104 ('h')
fmt.Println(s[6])  // 228 (part of '世', not complete rune!)

// Good: Use runes
runes := []rune(s)
fmt.Println(len(runes))  // 8 runes
fmt.Println(runes[6])    // 19990 ('世')

// Or: Range over string (returns runes)
for i, r := range s {
    fmt.Printf("%d: %c\n", i, r)
}
```

---

## Functions & Methods

### Mistake: Pointer vs Value Receivers

```go
type Counter struct {
    count int
}

// Bad: Value receiver doesn't modify original
func (c Counter) Increment() {
    c.count++  // Modifies copy!
}

counter := Counter{}
counter.Increment()
fmt.Println(counter.count)  // 0 (unchanged)

// Good: Pointer receiver modifies original
func (c *Counter) Increment() {
    c.count++
}

counter := Counter{}
counter.Increment()
fmt.Println(counter.count)  // 1
```

**When to use pointer receivers:**
- Method modifies the receiver
- Receiver is large (avoid copying)
- Consistency (if some methods use pointers, all should)

### Mistake: Named Return Values Confusion

```go
// Bad: Shadowing named return
func process() (result int, err error) {
    if someCondition {
        result, err := doSomething()  // Shadows! Creates new variables
        if err != nil {
            return  // Returns zero values, not the error!
        }
    }
    return
}

// Good: Assign to named returns
func process() (result int, err error) {
    if someCondition {
        result, err = doSomething()  // Assigns to named returns
        if err != nil {
            return
        }
    }
    return
}
```

---

## Error Handling

### Mistake: Not Wrapping Errors

```go
// Bad: Losing error context
if err := saveUser(user); err != nil {
    return err  // Which user? What failed?
}

// Good: Add context
if err := saveUser(user); err != nil {
    return fmt.Errorf("save user %d: %w", user.ID, err)
}
```

### Mistake: Checking Error Type Incorrectly

```go
// Bad: String comparison
if err.Error() == "not found" {  // Brittle!
    // ...
}

// Good: Use errors.Is for sentinel errors
if errors.Is(err, sql.ErrNoRows) {
    // ...
}

// Good: Use errors.As for type checking
var pathErr *os.PathError
if errors.As(err, &pathErr) {
    fmt.Println("Path:", pathErr.Path)
}
```

### Mistake: Ignoring Errors

```go
// Bad: Ignoring error
result, _ := doSomething()  // What if it failed?

// Good: Handle it
result, err := doSomething()
if err != nil {
    return fmt.Errorf("do something: %w", err)
}

// If you must ignore, be explicit
result, err := doSomething()
_ = err  // Explicitly ignored
```

---

## Concurrency

### Mistake: Not Protecting Shared Data

```go
// Bad: Race condition
type Counter struct {
    count int
}

func (c *Counter) Increment() {
    c.count++  // NOT safe for concurrent use!
}

// Good: Use mutex
type Counter struct {
    mu    sync.Mutex
    count int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

// Or: Use atomic
type Counter struct {
    count atomic.Int64
}

func (c *Counter) Increment() {
    c.count.Add(1)
}
```

### Mistake: Copying sync Types

```go
// Bad: Copying mutex
type Counter struct {
    mu    sync.Mutex
    count int
}

func (c Counter) Increment() {  // Value receiver copies mutex!
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

// Good: Always use pointer receiver with sync types
func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}
```

### Mistake: Goroutine Leaks

```go
// Bad: Goroutine never exits
func leak() {
    ch := make(chan int)
    go func() {
        val := <-ch  // Blocks forever if nothing sends!
        fmt.Println(val)
    }()
    // Function returns, goroutine leaked
}

// Good: Use context for cancellation
func noLeak(ctx context.Context) {
    ch := make(chan int)
    go func() {
        select {
        case val := <-ch:
            fmt.Println(val)
        case <-ctx.Done():
            return  // Goroutine exits
        }
    }()
}
```

### Mistake: Not Closing Channels

```go
// Bad: Range waits forever
ch := make(chan int)
go func() {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    // Never closes!
}()

for v := range ch {  // Blocks forever after 10 values
    fmt.Println(v)
}

// Good: Close channel
go func() {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch)  // Signal completion
}()
```

---

## Standard Library

### Mistake: Not Checking JSON Unmarshal Errors

```go
// Bad: Ignoring error
var data MyStruct
json.Unmarshal(bytes, &data)
// What if unmarshal failed?

// Good: Check error
var data MyStruct
if err := json.Unmarshal(bytes, &data); err != nil {
    return fmt.Errorf("unmarshal: %w", err)
}
```

### Mistake: HTTP Client Timeouts

```go
// Bad: No timeout
client := &http.Client{}
resp, err := client.Get(url)  // May hang forever!

// Good: Set timeout
client := &http.Client{
    Timeout: 10 * time.Second,
}
resp, err := client.Get(url)
```

### Mistake: Not Closing Response Body

```go
// Bad: Resource leak
resp, err := http.Get(url)
if err != nil {
    return err
}
// Forgot to close body!

// Good: Always close
resp, err := http.Get(url)
if err != nil {
    return err
}
defer resp.Body.Close()
```

---

## Testing

### Mistake: Not Using t.Helper()

```go
// Bad: Error reports wrong line
func assertEqual(t *testing.T, got, want int) {
    if got != want {
        t.Errorf("got %d, want %d", got, want)  // Reports line in helper
    }
}

// Good: Use t.Helper()
func assertEqual(t *testing.T, got, want int) {
    t.Helper()  // Reports line in caller
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}
```

### Mistake: Not Running Tests in Parallel

```go
// Tests can run in parallel
func TestSomething(t *testing.T) {
    t.Parallel()  // Mark as parallel

    // Test code...
}
```

---

## Quick Reference: Common Gotchas

1. **Loop variables**: `v := v` in closures
2. **Slices**: Understand length vs capacity
3. **Maps**: Iteration order is random
4. **Defer**: Executes in LIFO order
5. **Range**: Copies values (use indices for large structs)
6. **Goroutines**: Always ensure they can exit
7. **Channels**: Close from sender side only
8. **Mutex**: Don't copy, use pointers
9. **Errors**: Wrap with context
10. **HTTP**: Close response bodies

---

## Resources

- [100 Go Mistakes](https://100go.co/)
- [Common Go Mistakes](https://github.com/golang/go/wiki/CommonMistakes)
- [Go Vet](https://pkg.go.dev/cmd/vet) - Catches many of these
- [golangci-lint](https://golangci-lint.run/) - Comprehensive linting
