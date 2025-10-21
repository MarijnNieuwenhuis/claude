# Go Testing Best Practices

Comprehensive guide to writing effective tests in Go.

## Test Organization

```
package/
├── user.go
├── user_test.go        # Unit tests
├── user_integration_test.go
└── testdata/           # Test fixtures
    └── users.json
```

## Table-Driven Tests

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -1, -2},
        {"zero", 0, 5, 5},
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

## Test Helpers

```go
func assertEqual(t *testing.T, got, want interface{}) {
    t.Helper()  // Reports caller's line on failure
    if got != want {
        t.Errorf("got %v, want %v", got, want)
    }
}

func assertNoError(t *testing.T, err error) {
    t.Helper()
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
}
```

## Mocking with Interfaces

```go
type UserRepository interface {
    FindByID(id int) (*User, error)
}

type MockUserRepository struct {
    users map[int]*User
}

func (m *MockUserRepository) FindByID(id int) (*User, error) {
    if user, ok := m.users[id]; ok {
        return user, nil
    }
    return nil, errors.New("not found")
}

func TestService(t *testing.T) {
    repo := &MockUserRepository{
        users: map[int]*User{
            1: {ID: 1, Name: "Alice"},
        },
    }
    service := NewService(repo)
    // Test with mock...
}
```

## Test Fixtures

```go
func TestLoadConfig(t *testing.T) {
    // Use testdata/ directory
    data, err := os.ReadFile("testdata/config.json")
    assertNoError(t, err)

    var config Config
    err = json.Unmarshal(data, &config)
    assertNoError(t, err)

    assertEqual(t, config.Port, 8080)
}
```

## Parallel Tests

```go
func TestParallel(t *testing.T) {
    t.Parallel()  // Run in parallel with other parallel tests

    // Test code...
}
```

## Subtests

```go
func TestUser(t *testing.T) {
    t.Run("Valid", func(t *testing.T) {
        user := NewUser("alice@example.com")
        assertEqual(t, user.Email, "alice@example.com")
    })

    t.Run("Invalid", func(t *testing.T) {
        user := NewUser("invalid")
        if user != nil {
            t.Error("expected nil for invalid email")
        }
    })
}
```

## Benchmarks

```go
func BenchmarkFibonacci(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Fibonacci(20)
    }
}

// Run: go test -bench=.
```

## Examples (Testable Documentation)

```go
func ExampleAdd() {
    result := Add(2, 3)
    fmt.Println(result)
    // Output: 5
}
```

## Test Coverage

```bash
# Run tests with coverage
go test -cover

# Generate coverage profile
go test -coverprofile=coverage.out

# View coverage in browser
go tool cover -html=coverage.out
```

## Integration Tests

```go
// +build integration

package user_test

import "testing"

func TestDatabaseIntegration(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping integration test")
    }

    db := setupTestDatabase(t)
    defer db.Close()

    // Integration test...
}
```

Run integration tests:
```bash
go test -tags=integration
```

## HTTP Testing

```go
func TestHTTPHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/users/1", nil)
    w := httptest.NewRecorder()

    handler := UserHandler()
    handler.ServeHTTP(w, req)

    assertEqual(t, w.Code, http.StatusOK)

    var user User
    json.NewDecoder(w.Body).Decode(&user)
    assertEqual(t, user.ID, 1)
}
```

## Best Practices

1. **Use table-driven tests** for multiple cases
2. **Use t.Helper()** in test utilities
3. **Run tests in parallel** when possible
4. **Mock external dependencies** via interfaces
5. **Keep tests fast** (unit tests < 1s)
6. **Test behavior, not implementation**
7. **Use testdata/** for fixtures
8. **Write examples** for documentation
9. **Aim for 80%+ coverage** on critical code
10. **Test error cases** thoroughly

## Testing Anti-Patterns

- Don't test private functions directly
- Don't use sleep for synchronization
- Don't rely on execution order
- Don't share state between tests
- Don't write brittle tests (test behavior, not implementation)

## References

- [Testing in Go](https://go.dev/doc/tutorial/add-a-test)
- [Table Driven Tests](https://github.com/golang/go/wiki/TableDrivenTests)
- [Advanced Testing](https://go.dev/blog/subtests)
