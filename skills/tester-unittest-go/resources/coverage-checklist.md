# Test Coverage Checklist

Complete checklist for achieving 100% test coverage in Go projects.

## Pre-Testing Analysis

### 1. Understand the Code
- [ ] Read and understand the function/package being tested
- [ ] Identify all public functions
- [ ] Identify all exported types and methods
- [ ] Document all possible input types
- [ ] List all possible output types
- [ ] Map all error conditions
- [ ] Note all dependencies (interfaces, databases, external services)

### 2. Identify Test Scenarios
- [ ] Happy path (normal operation)
- [ ] Edge cases (boundaries, limits)
- [ ] Error cases (invalid input, failures)
- [ ] Nil/zero value handling
- [ ] Concurrent access (if applicable)
- [ ] State transitions (if stateful)

## Coverage Categories

### Line Coverage
- [ ] Every line of code executed at least once
- [ ] All variable declarations tested
- [ ] All assignments tested
- [ ] All function calls tested
- [ ] All return statements reached

### Branch Coverage
- [ ] All `if` conditions tested (true and false paths)
- [ ] All `else` branches tested
- [ ] All `else if` combinations tested
- [ ] All `switch` cases tested
- [ ] All `case` statements including `default`
- [ ] All `for` loop entry and exit conditions
- [ ] All `range` loops with empty and non-empty collections

### Function Coverage
- [ ] Every function has at least one test
- [ ] Every method has at least one test
- [ ] Every constructor tested
- [ ] Every getter/setter tested
- [ ] Every helper function tested

### Error Coverage
- [ ] All error returns tested
- [ ] All panic recovery tested (if applicable)
- [ ] All error wrapping tested
- [ ] All error types verified
- [ ] All sentinel errors tested

## Test Case Categories

### 1. Happy Path Tests
```go
- [ ] Valid input returns expected output
- [ ] Default values work correctly
- [ ] Typical use cases succeed
```

**Example**:
```go
{
    name: "valid user creates successfully",
    input: User{Name: "John", Email: "john@example.com"},
    want: success,
}
```

### 2. Edge Case Tests
```go
- [ ] Empty input
- [ ] Nil input
- [ ] Zero values
- [ ] Maximum values
- [ ] Minimum values
- [ ] Boundary values
- [ ] Single element
- [ ] Large collections
```

**Example**:
```go
{name: "empty string", input: "", wantErr: true},
{name: "nil pointer", input: nil, wantErr: true},
{name: "max int", input: math.MaxInt64, want: expected},
{name: "single element", input: []int{1}, want: result},
```

### 3. Error Case Tests
```go
- [ ] Invalid input types
- [ ] Out of range values
- [ ] Malformed data
- [ ] Missing required fields
- [ ] Duplicate values
- [ ] Conflicting values
```

**Example**:
```go
{name: "invalid email", input: "not-an-email", wantErr: ErrInvalidEmail},
{name: "negative age", input: -1, wantErr: ErrInvalidAge},
{name: "missing required field", input: User{}, wantErr: ErrMissingField},
```

### 4. State Transition Tests
```go
- [ ] Initial state
- [ ] State changes
- [ ] Final state
- [ ] Invalid transitions
- [ ] State rollback
```

**Example**:
```go
{name: "pending to active", from: Pending, to: Active, want: success},
{name: "active to pending not allowed", from: Active, to: Pending, wantErr: true},
```

### 5. Concurrency Tests
```go
- [ ] Concurrent reads
- [ ] Concurrent writes
- [ ] Race conditions
- [ ] Deadlock prevention
- [ ] Goroutine cleanup
```

**Example**:
```go
func TestConcurrentAccess(t *testing.T) {
    t.Parallel()
    cache := NewCache()

    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            cache.Set(id, id)
        }(i)
    }
    wg.Wait()

    assertEqual(t, cache.Len(), 100)
}
```

## Code Pattern Coverage

### If/Else Statements
```go
if condition {
    // [ ] Test with condition = true
} else {
    // [ ] Test with condition = false
}
```

### Switch Statements
```go
switch value {
case A:  // [ ] Test case A
case B:  // [ ] Test case B
case C:  // [ ] Test case C
default: // [ ] Test default case
}
```

### For Loops
```go
for i := 0; i < n; i++ {
    // [ ] Test with n = 0 (no iterations)
    // [ ] Test with n = 1 (single iteration)
    // [ ] Test with n > 1 (multiple iterations)
}

for k, v := range collection {
    // [ ] Test with empty collection
    // [ ] Test with single element
    // [ ] Test with multiple elements
}
```

### Error Handling
```go
result, err := Function()
if err != nil {
    // [ ] Test error path
    return err
}
// [ ] Test success path
```

### Defer/Panic/Recover
```go
defer cleanup()  // [ ] Test cleanup is called

func() {
    defer func() {
        if r := recover(); r != nil {
            // [ ] Test panic recovery
        }
    }()
    // Code that might panic
}()
```

## Interface Coverage

### Mock Implementation
```go
- [ ] All interface methods implemented in mock
- [ ] Mock returns expected values for happy path
- [ ] Mock returns errors for error cases
- [ ] Mock tracks calls (if needed)
- [ ] Mock verifies arguments (if needed)
```

**Example**:
```go
type MockRepository struct {
    FindFunc func(id int) (*User, error)
    calls    int
}

func (m *MockRepository) Find(id int) (*User, error) {
    m.calls++
    if m.FindFunc != nil {
        return m.FindFunc(id)
    }
    return nil, errors.New("not implemented")
}
```

### Interface Compliance
```go
- [ ] Verify type implements interface at compile time
```

**Example**:
```go
var _ Repository = (*DatabaseRepository)(nil)  // Compile-time check
```

## File and Package Coverage

### Per-File Checklist
- [ ] All exported functions tested
- [ ] All exported types tested
- [ ] All exported constants verified
- [ ] All exported variables checked
- [ ] All unexported functions tested (if complex)

### Per-Package Checklist
- [ ] Package-level tests exist
- [ ] Integration tests (if applicable)
- [ ] Example tests for documentation
- [ ] Benchmark tests for performance-critical code

## Coverage Validation

### Running Coverage
```bash
# [ ] Run tests with coverage
make test

# [ ] Generate coverage report
go test -coverprofile=coverage.out ./...

# [ ] View coverage in terminal
go tool cover -func=coverage.out

# [ ] View HTML coverage report
go tool cover -html=coverage.out

# [ ] Check coverage percentage
go test -cover ./...
```

### Coverage Analysis
- [ ] Overall coverage ≥ 100% (or ≥ 95% with justification)
- [ ] Per-package coverage ≥ 95%
- [ ] Per-file coverage ≥ 95%
- [ ] No untested exported functions
- [ ] All critical paths covered

### Coverage Gaps
For any uncovered code:
- [ ] Identify why it's uncovered
- [ ] Add tests to cover it, OR
- [ ] Document why it's acceptable (unreachable code, logging, etc.)

## Test Quality Checklist

### Test Structure
- [ ] Tests follow naming convention `Test{FunctionName}`
- [ ] Table-driven tests used for multiple cases
- [ ] Each test has descriptive name
- [ ] Test cases are independent
- [ ] Tests can run in any order
- [ ] Tests can run in parallel (where safe)

### Test Helpers
- [ ] Helpers use `t.Helper()`
- [ ] Helpers have clear names
- [ ] Helpers are reusable
- [ ] Helpers don't hide important logic

### Assertions
- [ ] Use clear assertion messages
- [ ] Verify exact values (not just non-nil)
- [ ] Check error types, not just presence
- [ ] Compare expected vs actual consistently

### Test Data
- [ ] Test data in `testdata/` directory
- [ ] Mock data is realistic
- [ ] Edge cases represented
- [ ] Test data is maintainable

## Common Coverage Gaps

### Gap 1: Error Paths Not Tested
```go
// Production code
result, err := Operation()
if err != nil {
    return fmt.Errorf("operation failed: %w", err)  // Often not tested
}

// Solution: Add error test
func TestOperation_Error(t *testing.T) {
    _, err := OperationWithError()
    assertError(t, err)
}
```

### Gap 2: Early Returns Not Tested
```go
// Production code
if input == nil {
    return ErrNilInput  // Sometimes missed
}

// Solution: Add nil test
{name: "nil input", input: nil, wantErr: ErrNilInput}
```

### Gap 3: Default Case Not Tested
```go
// Production code
switch status {
case Active: return "active"
case Inactive: return "inactive"
default: return "unknown"  // Often not tested
}

// Solution: Add default case test
{name: "unknown status", status: 999, want: "unknown"}
```

### Gap 4: Loop Edge Cases
```go
// Production code
for _, item := range items {
    process(item)
}

// Test with:
- [ ] Empty slice
- [ ] Single item
- [ ] Multiple items
```

### Gap 5: Concurrent Edge Cases
```go
// Production code (thread-safe)
func (c *Cache) Set(key, value) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = value
}

// Test with:
- [ ] Single goroutine
- [ ] Multiple concurrent goroutines
- [ ] Race detector enabled: go test -race
```

## Coverage Exemptions

### Acceptable Uncovered Code
Document why these are acceptable:

1. **Logging Statements**
   ```go
   log.Printf("debug: %v", data)  // OK to not test
   ```

2. **Defensive Programming**
   ```go
   if impossible_condition {  // OK if truly impossible
       panic("this should never happen")
   }
   ```

3. **Main Function**
   ```go
   func main() {  // Application entry point, tested via integration tests
       // ...
   }
   ```

4. **Generated Code**
   ```go
   // Code generated by tool; DO NOT EDIT.
   // OK to exclude from coverage
   ```

5. **Platform-Specific Code**
   ```go
   // +build !windows
   // OK if platform not tested
   ```

## Final Validation

### Before Submitting
- [ ] All tests pass: `go test ./...`
- [ ] Coverage report generated: `make test`
- [ ] Coverage ≥ 100% (or 95%+ with justification)
- [ ] No test files modified in production code
- [ ] All new tests follow Go conventions
- [ ] Test names are descriptive
- [ ] No commented-out tests
- [ ] No `t.Skip()` without good reason

### Coverage Report Review
- [ ] Open HTML coverage report
- [ ] Review all red (uncovered) lines
- [ ] Add tests for uncovered lines OR document exemption
- [ ] Verify all branches covered (green)
- [ ] Check coverage percentage per file

### Mutation Testing
- [ ] Run mutation tests (if implemented)
- [ ] Document mutation score
- [ ] Verify critical paths kill mutations
- [ ] Justify any survived mutations

## Coverage Improvement Strategies

### Strategy 1: Start with Happy Path
1. Write tests for normal operation
2. Verify basic functionality works
3. Build confidence before edge cases

### Strategy 2: Add Error Cases
1. Test each error return
2. Test invalid inputs
3. Test boundary conditions

### Strategy 3: Review Coverage Report
1. Run `make test` to generate report
2. Open HTML report
3. Click on files with < 100% coverage
4. Add tests for red lines

### Strategy 4: Use Coverage-Driven Development
1. Write test
2. Run coverage
3. Add test for uncovered line
4. Repeat until 100%

### Strategy 5: Refactor Untestable Code
If code is hard to test:
1. Extract interfaces
2. Inject dependencies
3. Split large functions
4. Remove direct I/O in business logic
**Note**: Report this to user; don't fix yourself!

## Example: Complete Coverage Workflow

### Step 1: Analyze Function
```go
// validator.go
func ValidateEmail(email string) error {
    if email == "" {
        return ErrEmptyEmail
    }
    if !strings.Contains(email, "@") {
        return ErrNoAtSign
    }
    parts := strings.Split(email, "@")
    if len(parts) != 2 {
        return ErrInvalidFormat
    }
    if parts[0] == "" || parts[1] == "" {
        return ErrMissingParts
    }
    return nil
}
```

### Step 2: Identify Test Cases
- [ ] Empty string → ErrEmptyEmail
- [ ] No @ sign → ErrNoAtSign
- [ ] Multiple @ signs → ErrInvalidFormat
- [ ] Empty user part → ErrMissingParts
- [ ] Empty domain part → ErrMissingParts
- [ ] Valid email → nil error

### Step 3: Write Tests
```go
func TestValidateEmail(t *testing.T) {
    tests := []struct {
        name    string
        email   string
        wantErr error
    }{
        {"empty email", "", ErrEmptyEmail},
        {"no at sign", "userexample.com", ErrNoAtSign},
        {"multiple at signs", "user@@example.com", ErrInvalidFormat},
        {"empty user", "@example.com", ErrMissingParts},
        {"empty domain", "user@", ErrMissingParts},
        {"valid email", "user@example.com", nil},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := ValidateEmail(tt.email)
            if !errors.Is(err, tt.wantErr) {
                t.Errorf("got %v, want %v", err, tt.wantErr)
            }
        })
    }
}
```

### Step 4: Run Coverage
```bash
make test
# Coverage: 100%
```

### Step 5: Verify in HTML Report
- All lines green ✓
- All branches covered ✓
- No red lines ✓

## Summary

Achieving 100% coverage requires:
1. ✅ Systematic approach to test case identification
2. ✅ Testing all paths: happy, error, edge cases
3. ✅ Using coverage tools to find gaps
4. ✅ Iterating until all lines are green
5. ✅ Documenting any acceptable exemptions

**Remember**: Coverage is necessary but not sufficient. Combine with mutation testing to ensure test quality!
