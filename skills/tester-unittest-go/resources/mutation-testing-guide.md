# Mutation Testing Guide for Go

Mutation testing is a technique to evaluate the quality of your tests by introducing intentional bugs (mutations) and verifying that tests catch them.

## What is Mutation Testing?

**Concept**: If you introduce a bug in your code, your tests should fail. If they don't, your tests are inadequate.

**Mutation Score** = (Killed Mutations / Total Mutations) × 100%

- **Killed**: Mutation caused test to fail ✓
- **Survived**: Mutation didn't break any test ✗
- **Target Score**: 80%+ indicates good test quality

## Mutation Testing Tools

### 1. go-mutesting (Recommended)

```bash
# Install
go install github.com/zimmski/go-mutesting/cmd/go-mutesting@latest

# Run on a package
go-mutesting ./internal/validator/

# Run with verbose output
go-mutesting --verbose ./internal/validator/

# Run on specific file
go-mutesting ./internal/validator/validator.go

# Run with specific mutation types
go-mutesting --exec-timeout 5s ./internal/validator/
```

**Example Output**:
```
PASS   internal/validator/validator.go:45:2 - Mutation: changed == to !=
FAIL   internal/validator/validator.go:67:5 - Mutation: removed error check
Score: 18/20 (90%)
```

### 2. Manual Mutation Testing

For critical code, manually verify mutations:

```go
// Original code
func ValidateEmail(email string) error {
    if len(email) == 0 {
        return ErrEmptyEmail
    }
    if !strings.Contains(email, "@") {
        return ErrInvalidEmail
    }
    return nil
}

// Mutation 1: Change == to !=
func ValidateEmail(email string) error {
    if len(email) != 0 {  // MUTATION
        return ErrEmptyEmail
    }
    // ...
}
// Expected: Test should fail for empty string

// Mutation 2: Change ! to nothing
func ValidateEmail(email string) error {
    if len(email) == 0 {
        return ErrEmptyEmail
    }
    if strings.Contains(email, "@") {  // MUTATION
        return ErrInvalidEmail
    }
    return nil
}
// Expected: Test should fail for valid emails
```

## Common Mutation Types

### 1. Conditional Boundary Mutations

**Original → Mutation**:
- `>` → `>=`
- `<` → `<=`
- `==` → `!=`
- `>=` → `>`
- `<=` → `<`

**Example**:
```go
// Original
if age >= 18 {
    return true
}

// Mutation
if age > 18 {  // Should be caught by test with age=18
    return true
}
```

**Test to catch**:
```go
{
    name: "exactly 18 years old",
    age:  18,
    want: true,  // This will fail if mutated to >
}
```

### 2. Arithmetic Operator Mutations

**Original → Mutation**:
- `+` → `-`
- `-` → `+`
- `*` → `/`
- `/` → `*`

**Example**:
```go
// Original
total := price + tax

// Mutation
total := price - tax  // Should be caught by test

// Test to catch
assertEqual(t, Calculate(100, 10), 110)  // Fails if mutation applied
```

### 3. Logical Operator Mutations

**Original → Mutation**:
- `&&` → `||`
- `||` → `&&`
- `!x` → `x`

**Example**:
```go
// Original
if user != nil && user.IsActive {
    return true
}

// Mutation
if user != nil || user.IsActive {  // Should panic on nil user
    return true
}

// Test to catch
func TestNilUser(t *testing.T) {
    result := CheckUser(nil)  // Should not panic, mutation causes panic
    assertEqual(t, result, false)
}
```

### 4. Return Value Mutations

**Original → Mutation**:
- `return true` → `return false`
- `return nil` → `return err`
- `return 0` → `return 1`

**Example**:
```go
// Original
func IsValid() bool {
    return true
}

// Mutation
func IsValid() bool {
    return false  // Should be caught if tested
}
```

### 5. Constant Mutations

**Original → Mutation**:
- `0` → `1`
- `1` → `0`
- `""` → `"x"`

**Example**:
```go
// Original
const MaxRetries = 3

// Mutation
const MaxRetries = 4  // Should be caught by test

// Test to catch
func TestRetries(t *testing.T) {
    counter := 0
    for i := 0; i < MaxRetries; i++ {
        counter++
    }
    assertEqual(t, counter, 3)  // Fails if MaxRetries mutated
}
```

### 6. Statement Deletion Mutations

**Original → Mutation**:
- Remove error check
- Remove validation
- Remove function call

**Example**:
```go
// Original
if err := Validate(input); err != nil {
    return err
}

// Mutation (remove check)
// [deleted]

// Test to catch
func TestInvalidInput(t *testing.T) {
    err := Process(invalidInput)
    assertError(t, err, ErrInvalidInput)  // Fails if check removed
}
```

## Mutation Testing Workflow

### Step 1: Write Your Tests

First, write comprehensive unit tests:

```go
func TestValidateEmail(t *testing.T) {
    tests := []struct {
        name    string
        email   string
        wantErr bool
    }{
        {"valid email", "user@example.com", false},
        {"empty email", "", true},
        {"no @", "userexample.com", true},
        {"no domain", "user@", true},
        {"no user", "@example.com", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := ValidateEmail(tt.email)
            if (err != nil) != tt.wantErr {
                t.Errorf("got error %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

### Step 2: Achieve 100% Code Coverage

Ensure all lines are covered:

```bash
make test
# Check HTML report for any red (uncovered) lines
```

### Step 3: Run Mutation Testing

```bash
go-mutesting --verbose ./internal/validator/validator.go
```

### Step 4: Analyze Results

Look for survived mutations:

```
SURVIVE internal/validator/validator.go:23:5 - changed == to !=
SURVIVE internal/validator/validator.go:45:2 - removed if statement
```

### Step 5: Add Tests for Survived Mutations

For each survived mutation, add a test:

```go
// Mutation survived: changed len(email) == 0 to len(email) != 0
// Add this test:
{
    name:    "empty string should error",
    email:   "",
    wantErr: true,
},

// Mutation survived: removed @ validation
// Add this test:
{
    name:    "missing @ should error",
    email:   "userexample.com",
    wantErr: true,
},
```

### Step 6: Rerun and Verify

```bash
# Run tests
go test -v ./...

# Rerun mutation testing
go-mutesting ./internal/validator/validator.go

# Verify all mutations killed
```

## Mutation Testing Report Template

```markdown
## Mutation Testing Report

### Package: `internal/validator`
**Date**: 2025-01-24
**Coverage**: 100%

### Summary
| Metric | Value |
|--------|-------|
| Total Mutations | 32 |
| Killed | 30 |
| Survived | 2 |
| Mutation Score | 93.75% |
| Target Score | 80% |
| **Status** | ✅ PASS |

### Mutations by Type
| Type | Total | Killed | Survived |
|------|-------|--------|----------|
| Conditional Boundary | 12 | 12 | 0 |
| Arithmetic Operators | 5 | 5 | 0 |
| Logical Operators | 8 | 7 | 1 |
| Return Values | 4 | 4 | 0 |
| Statement Deletion | 3 | 2 | 1 |

### Survived Mutations

#### 1. Logical Operator Mutation (Line 45)
**Mutation**: Changed `&&` to `||` in validation condition
**Location**: `validator.go:45`
**Code**:
```go
// Original
if user != nil && user.IsActive {

// Mutated
if user != nil || user.IsActive {
```

**Why Survived**: Missing test for nil user with inactive status
**Action Taken**: Added test case
```go
{
    name: "nil user returns false",
    user: nil,
    want: false,
}
```

**Retest Result**: ✅ Killed

#### 2. Statement Deletion (Line 67)
**Mutation**: Removed error logging statement
**Location**: `validator.go:67`
**Code**:
```go
// Original
if err != nil {
    log.Error(err)  // This line deleted
    return err
}
```

**Why Survived**: Error logging not verified in tests (acceptable)
**Action**: No action - logging is side effect, not business logic
**Justification**: Tests verify return values, not logging

### Final Mutation Score: 96.8% (31/32)

### Recommendations
- ✅ Excellent mutation score above 80% target
- ✅ All business logic mutations caught
- ✅ Only logging side effect not verified (acceptable)

### Tests Added
- `TestValidatorNilUser`
- `TestValidatorInactiveUser`
- `TestValidatorNilAndInactive`

### Conclusion
Test suite is robust and catches all meaningful bugs. Mutation testing validates that our 100% code coverage translates to high-quality tests.
```

## Best Practices

### 1. Don't Aim for 100% Mutation Score

Some mutations are acceptable to survive:
- Logging statements
- Debug output
- Performance optimizations
- Defensive programming checks

**Target**: 80-95% mutation score

### 2. Focus on Critical Code

Prioritize mutation testing for:
- Business logic
- Validation functions
- Error handling
- Security-critical code
- Complex algorithms

### 3. Use Mutation Testing to Find Test Gaps

If a mutation survives:
1. Analyze why it wasn't caught
2. Add a specific test case
3. Verify the new test kills the mutation

### 4. Combine with Code Coverage

```
High Coverage + High Mutation Score = Excellent Tests
High Coverage + Low Mutation Score = Weak Tests
Low Coverage + Any Mutation Score = Inadequate Tests
```

### 5. Automate in CI/CD

Add to your pipeline:

```bash
#!/bin/bash
# ci/mutation-test.sh

echo "Running mutation tests..."
go-mutesting ./internal/... > mutation-report.txt

# Extract score
SCORE=$(grep "Score:" mutation-report.txt | awk '{print $2}')

# Fail if below threshold
THRESHOLD=80
if [ "$SCORE" -lt "$THRESHOLD" ]; then
    echo "Mutation score $SCORE% is below threshold $THRESHOLD%"
    exit 1
fi

echo "Mutation testing passed: $SCORE%"
```

## Common Patterns for Mutation-Resistant Tests

### Pattern 1: Test Exact Values

```go
// Weak - mutation could survive
func TestCalculate(t *testing.T) {
    result := Calculate(5, 3)
    if result <= 0 {  // Too lenient
        t.Error("expected positive result")
    }
}

// Strong - mutations will be caught
func TestCalculate(t *testing.T) {
    result := Calculate(5, 3)
    assertEqual(t, result, 15)  // Exact value
}
```

### Pattern 2: Test Boundaries

```go
// Test boundary conditions explicitly
tests := []struct{
    age  int
    want bool
}{
    {17, false},  // Just below boundary
    {18, true},   // Exact boundary
    {19, true},   // Just above boundary
}
```

### Pattern 3: Test Error Messages

```go
// Weak
if err == nil {
    t.Error("expected error")
}

// Strong
expectedErr := "invalid email format"
if err == nil {
    t.Error("expected error")
}
if err.Error() != expectedErr {
    t.Errorf("got %q, want %q", err.Error(), expectedErr)
}
```

### Pattern 4: Test All Branches

```go
func Classify(age int) string {
    if age < 13 {
        return "child"
    } else if age < 18 {
        return "teen"
    } else if age < 65 {
        return "adult"
    }
    return "senior"
}

// Test ALL branches
tests := []struct{
    age  int
    want string
}{
    {10, "child"},   // First branch
    {15, "teen"},    // Second branch
    {30, "adult"},   // Third branch
    {70, "senior"},  // Default branch
    {12, "child"},   // Boundary
    {13, "teen"},    // Boundary
    {17, "teen"},    // Boundary
    {18, "adult"},   // Boundary
}
```

## Troubleshooting

### Issue: Mutation Score Too Low

**Solution**:
1. Review survived mutations
2. Add specific test cases for each
3. Test edge cases and boundaries
4. Verify all branches are tested

### Issue: Mutation Testing Too Slow

**Solution**:
```bash
# Use timeout
go-mutesting --exec-timeout 5s ./...

# Test specific packages
go-mutesting ./internal/validator/

# Run in parallel (if tool supports)
go-mutesting --parallel 4 ./...
```

### Issue: Too Many False Positives

**Solution**:
1. Identify acceptable survivors (logging, debug)
2. Document why they're acceptable
3. Focus on business logic mutations
4. Set realistic threshold (80%, not 100%)

## Integration with Development Workflow

### Pre-commit Hook

```bash
#!/bin/bash
# .git/hooks/pre-commit

# Run tests
go test ./... || exit 1

# Quick mutation test on changed files
CHANGED_FILES=$(git diff --cached --name-only --diff-filter=AM | grep '.go$' | grep -v '_test.go$')

if [ -n "$CHANGED_FILES" ]; then
    for file in $CHANGED_FILES; do
        echo "Mutation testing $file..."
        go-mutesting --exec-timeout 10s "$file"
    done
fi
```

### Code Review Checklist

- [ ] Code coverage 100% (or 95%+)
- [ ] Mutation score 80%+
- [ ] All survived mutations justified
- [ ] Critical paths have mutation tests
- [ ] Boundary conditions tested

## Resources

- [go-mutesting GitHub](https://github.com/zimmski/go-mutesting)
- [Mutation Testing Concepts](https://en.wikipedia.org/wiki/Mutation_testing)
- [Go Testing Best Practices](https://go.dev/doc/tutorial/add-a-test)
