# Go Concurrency Patterns

Patterns for effective concurrent programming with goroutines and channels.

## Core Concepts

- **Goroutines**: Lightweight threads
- **Channels**: Communication between goroutines
- **Select**: Multiplex channel operations
- **Sync package**: Mutexes, WaitGroups, etc.

## Worker Pool

Process jobs concurrently with limited workers.

```go
func worker(id int, jobs <-chan Job, results chan<- Result) {
    for job := range jobs {
        fmt.Printf("worker %d processing job %d\n", id, job.ID)
        result := job.Process()
        results <- result
    }
}

func main() {
    jobs := make(chan Job, 100)
    results := make(chan Result, 100)

    // Start workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send jobs
    for j := 1; j <= 9; j++ {
        jobs <- Job{ID: j}
    }
    close(jobs)

    // Collect results
    for a := 1; a <= 9; a++ {
        <-results
    }
}
```

## Pipeline

Chain processing stages.

```go
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// Usage
c := gen(2, 3, 4)
out := sq(c)
for n := range out {
    fmt.Println(n)  // 4, 9, 16
}
```

## Fan-Out/Fan-In

Distribute work, collect results.

```go
func fanIn(channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    output := func(c <-chan int) {
        defer wg.Done()
        for n := range c {
            out <- n
        }
    }

    wg.Add(len(channels))
    for _, c := range channels {
        go output(c)
    }

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}
```

## Context for Cancellation

```go
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("cancelled")
            return
        default:
            // Do work
            time.Sleep(time.Second)
        }
    }
}

func main() {
    ctx, cancel := context.WithTimeout(
        context.Background(),
        5*time.Second,
    )
    defer cancel()

    go worker(ctx)
    time.Sleep(10 * time.Second)
}
```

## Rate Limiting

```go
func rateLimited(ctx context.Context, rate time.Duration) {
    ticker := time.NewTicker(rate)
    defer ticker.Stop()

    for {
        select {
        case <-ctx.Done():
            return
        case <-ticker.C:
            // Process one item per tick
            process()
        }
    }
}
```

## Timeout Pattern

```go
func fetchWithTimeout(url string, timeout time.Duration) ([]byte, error) {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    return io.ReadAll(resp.Body)
}
```

## Best Practices

1. **Always ensure goroutines can exit**
2. **Use context for cancellation**
3. **Close channels from sender side**
4. **Don't copy mutexes**
5. **Protect shared data with mutexes or channels**
6. **Use WaitGroup for synchronization**

## References

- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Advanced Go Concurrency Patterns](https://go.dev/blog/io2013-talk-concurrency)
