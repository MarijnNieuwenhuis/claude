# Error Handling in Go

Comprehensive guide to idiomatic error handling in Go.

## Core Principles

1. **Errors are values** - Not exceptions
2. **Explicit handling** - Check every error
3. **Add context** - Wrap errors with information
4. **Preserve original** - Use `%w` to wrap

## Basic Error Handling

```go
// Return errors
func FetchUser(id int) (*User, error) {
    if id <= 0 {
        return nil, errors.New("invalid user ID")
    }
    // ...
    return user, nil
}

// Check immediately
user, err := FetchUser(id)
if err != nil {
    return fmt.Errorf("fetch user: %w", err)
}
```

## Error Wrapping

```go
// Wrap with fmt.Errorf and %w
if err := db.Save(user); err != nil {
    return fmt.Errorf("save user %d: %w", user.ID, err)
}

// Check wrapped errors
if errors.Is(err, sql.ErrNoRows) {
    // Handle not found
}
```

## Custom Errors

```go
type ValidationError struct {
    Field  string
    Reason string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Reason)
}

// Check type with errors.As
var ve *ValidationError
if errors.As(err, &ve) {
    fmt.Printf("Field %s failed\n", ve.Field)
}
```

## Sentinel Errors

```go
var (
    ErrNotFound = errors.New("not found")
    ErrInvalid  = errors.New("invalid")
)

func Find(id int) (*Item, error) {
    if id <= 0 {
        return nil, ErrInvalid
    }
    // ...
    return nil, ErrNotFound
}

// Check with errors.Is
if errors.Is(err, ErrNotFound) {
    // Handle
}
```

## Error Best Practices

1. **Don't panic** - Use for truly exceptional cases only
2. **Check immediately** - Don't defer error checking
3. **Add context** - Help debug issues
4. **Be specific** - Clear error messages
5. **Don't ignore** - Handle or explicitly discard

## References

- [Error handling in Go](https://go.dev/blog/error-handling-and-go)
- [Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)
