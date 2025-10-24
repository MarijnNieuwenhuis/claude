# Production Code Issues Report Template

Use this template when you identify issues in production code that need to be fixed by the user.

---

# Production Code Issues Report

**Package**: `{package_name}`
**Date**: `{date}`
**Analyzed By**: Go Unit Test Specialist
**Scope**: `{files_analyzed}`

## Executive Summary

- **Total Issues Found**: {number}
- **Critical**: {number} (Prevent testing or cause bugs)
- **High**: {number} (Impact testability significantly)
- **Medium**: {number} (Reduce code quality)
- **Low**: {number} (Minor improvements)

**Overall Impact**: {brief_description}

---

## Critical Issues

### Issue #{number}: {Short Description}

**File**: `{file_path}`
**Location**: `{file_path}:{line_number}`
**Severity**: 🔴 Critical
**Category**: {Bug | Security | Untestable Code | Nil Pointer Risk}

#### Problem
{Detailed description of what's wrong}

#### Current Code
```go
// Line {line_number}
{current_code}
```

#### Why This is Critical
- {Reason 1: e.g., Can cause nil pointer panic}
- {Reason 2: e.g., Prevents achieving code coverage}
- {Reason 3: e.g., Security vulnerability}

#### Impact on Testing
{How this affects your ability to write tests or achieve coverage}

#### Recommended Fix
```go
// Suggested replacement
{fixed_code}
```

#### Explanation
{Why this fix solves the problem}

#### Test Coverage Impact
- Current: Cannot test {specific scenario}
- After fix: Can achieve 100% coverage

---

## High Priority Issues

### Issue #{number}: {Short Description}

**File**: `{file_path}`
**Location**: `{file_path}:{line_number}`
**Severity**: 🟠 High
**Category**: {Error Handling | Race Condition | Resource Leak | Code Smell}

#### Problem
{Description}

#### Current Code
```go
{current_code}
```

#### Issues
1. {Issue 1}
2. {Issue 2}
3. {Issue 3}

#### Recommended Fix
```go
{fixed_code}
```

#### Impact
- **Functionality**: {Impact on behavior}
- **Testing**: {Impact on testability}
- **Maintenance**: {Impact on code maintainability}

---

## Medium Priority Issues

### Issue #{number}: {Short Description}

**File**: `{file_path}`
**Location**: `{file_path}:{line_number}`
**Severity**: 🟡 Medium
**Category**: {Best Practices | Code Quality | Documentation}

#### Problem
{Description}

#### Current Code
```go
{current_code}
```

#### Recommended Improvement
```go
{improved_code}
```

#### Benefits
- {Benefit 1}
- {Benefit 2}

---

## Low Priority Issues

### Issue #{number}: {Short Description}

**File**: `{file_path}`
**Location**: `{file_path}:{line_number}`
**Severity**: 🟢 Low
**Category**: {Style | Naming | Documentation}

#### Suggestion
{Brief description}

#### Example
```go
// Current
{current}

// Suggested
{suggested}
```

---

## Detailed Issue Categories

### 1. Untestable Code

Issues that make code impossible or very difficult to test:

#### Example: Direct os.Exit() Call
```go
// Problem: Cannot test error path
func Run() {
    if err := setup(); err != nil {
        fmt.Println(err)
        os.Exit(1)  // Untestable!
    }
}

// Fix: Return error instead
func Run() error {
    if err := setup(); err != nil {
        return fmt.Errorf("setup failed: %w", err)
    }
    return nil
}
```

#### Example: Global State Dependency
```go
// Problem: Cannot isolate tests
var globalCache = make(map[string]string)

func GetValue(key string) string {
    return globalCache[key]  // Untestable with different data
}

// Fix: Inject dependency
type Cache struct {
    data map[string]string
}

func (c *Cache) GetValue(key string) string {
    return c.data[key]  // Can test with mock cache
}
```

### 2. Missing Error Checks

```go
// Problem: Ignoring errors
result, _ := MayFail()
process(result)  // May use invalid result!

// Fix: Check errors
result, err := MayFail()
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}
process(result)
```

### 3. Nil Pointer Risks

```go
// Problem: No nil check
func ProcessUser(u *User) string {
    return u.Name  // Panic if u is nil!
}

// Fix: Add nil check
func ProcessUser(u *User) (string, error) {
    if u == nil {
        return "", errors.New("user is nil")
    }
    return u.Name, nil
}
```

### 4. Race Conditions

```go
// Problem: Concurrent access without synchronization
type Counter struct {
    count int
}

func (c *Counter) Increment() {
    c.count++  // Race condition!
}

// Fix: Add mutex
type Counter struct {
    mu    sync.Mutex
    count int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}
```

### 5. Resource Leaks

```go
// Problem: File not closed on error
func ReadFile(name string) ([]byte, error) {
    f, err := os.Open(name)
    if err != nil {
        return nil, err
    }
    data := make([]byte, 100)
    if _, err := f.Read(data); err != nil {
        return nil, err  // File not closed!
    }
    f.Close()
    return data, nil
}

// Fix: Use defer
func ReadFile(name string) ([]byte, error) {
    f, err := os.Open(name)
    if err != nil {
        return nil, err
    }
    defer f.Close()  // Always closes

    data := make([]byte, 100)
    if _, err := f.Read(data); err != nil {
        return nil, err
    }
    return data, nil
}
```

---

## Test Coverage Analysis

### Current Coverage
| Package | Coverage | Target | Gap |
|---------|----------|--------|-----|
| {package1} | {X}% | 100% | {gap}% |
| {package2} | {X}% | 100% | {gap}% |

### Coverage Blockers

#### Blocker 1: {Description}
- **Lines affected**: {file}:{line_range}
- **Cannot test because**: {reason}
- **Fix required**: {issue_reference}

#### Blocker 2: {Description}
- **Lines affected**: {file}:{line_range}
- **Cannot test because**: {reason}
- **Fix required**: {issue_reference}

### Estimated Coverage After Fixes
| Package | Current | After Fixes | Improvement |
|---------|---------|-------------|-------------|
| {package1} | {X}% | {Y}% | +{Z}% |

---

## Recommendations

### Immediate Actions (Critical)
1. ⚠️ Fix Issue #{number}: {description}
2. ⚠️ Fix Issue #{number}: {description}
3. ⚠️ Fix Issue #{number}: {description}

### Short-term Actions (High Priority)
1. 🔶 Address Issue #{number}: {description}
2. 🔶 Address Issue #{number}: {description}

### Long-term Improvements (Medium/Low Priority)
1. 📋 Consider Issue #{number}: {description}
2. 📋 Consider Issue #{number}: {description}

---

## Implementation Priority

### Phase 1: Critical Fixes (Do First)
These must be fixed before achieving proper test coverage:
- [ ] Issue #{number}: {description}
- [ ] Issue #{number}: {description}

**Timeline**: Immediate
**Effort**: {estimate}

### Phase 2: High Priority (Do Soon)
These significantly improve testability and code quality:
- [ ] Issue #{number}: {description}
- [ ] Issue #{number}: {description}

**Timeline**: Within 1 week
**Effort**: {estimate}

### Phase 3: Improvements (Optional)
Nice-to-have improvements:
- [ ] Issue #{number}: {description}
- [ ] Issue #{number}: {description}

**Timeline**: As time permits
**Effort**: {estimate}

---

## Testing Impact Summary

### Before Fixes
- ❌ Cannot achieve 100% coverage
- ❌ Cannot test {specific scenarios}
- ❌ Tests may produce false positives
- ❌ Race conditions in concurrent tests

### After Fixes
- ✅ Can achieve 100% coverage
- ✅ All scenarios testable
- ✅ Reliable test results
- ✅ Thread-safe tests possible

---

## Code Quality Metrics

### Current State
| Metric | Value | Target | Status |
|--------|-------|--------|--------|
| Test Coverage | {X}% | 100% | 🔴 Below |
| Cyclomatic Complexity | {Y} | <10 | {status} |
| Code Smells | {Z} | 0 | {status} |
| Critical Issues | {N} | 0 | 🔴 Needs Fix |

### After Fixes
| Metric | Value | Target | Status |
|--------|-------|--------|--------|
| Test Coverage | 100% | 100% | ✅ Met |
| Cyclomatic Complexity | {Y} | <10 | {status} |
| Code Smells | 0 | 0 | ✅ Met |
| Critical Issues | 0 | 0 | ✅ Met |

---

## Additional Resources

### Relevant Documentation
- [Error Handling Best Practices](.claude/refs/go/error-handling.md)
- [Idiomatic Go](.claude/refs/go/idiomatic-go.md)
- [Testing Practices](.claude/refs/go/testing-practices.md)

### Related Issues
- Similar issue in {other_package}
- Related pattern in {other_file}

### Code Review Guidelines
When fixing these issues, ensure:
- [ ] All error paths are checked
- [ ] No nil pointer dereferences possible
- [ ] Resources are properly cleaned up (use defer)
- [ ] Concurrent access is synchronized
- [ ] Functions are testable (no global state)
- [ ] Side effects are minimized

---

## Appendix: Full File Analysis

### File: `{file_path}`

#### Overall Assessment
- **Lines of Code**: {count}
- **Functions**: {count}
- **Complexity**: {average}
- **Issues Found**: {count}
- **Test Coverage**: {percentage}%

#### Issue Summary
| Line | Severity | Category | Description |
|------|----------|----------|-------------|
| {line} | Critical | {category} | {brief} |
| {line} | High | {category} | {brief} |
| {line} | Medium | {category} | {brief} |

#### Detailed Analysis
{Per-function analysis if needed}

---

## Sign-off

**Report Prepared By**: Go Unit Test Specialist
**Date**: {date}
**Next Review**: After fixes implemented

**User Action Required**:
1. Review all Critical and High priority issues
2. Apply recommended fixes
3. Re-run tests to verify coverage improvements
4. Request test suite completion after fixes

---

## Example: Complete Report

# Production Code Issues Report

**Package**: `internal/validator`
**Date**: 2025-01-24
**Analyzed By**: Go Unit Test Specialist
**Scope**: `validator.go`, `config.go`, `rules.go`

## Executive Summary

- **Total Issues Found**: 5
- **Critical**: 2 (Prevent testing)
- **High**: 2 (Impact testability)
- **Medium**: 1 (Code quality)
- **Low**: 0

**Overall Impact**: Cannot achieve 100% test coverage due to 2 critical issues. High priority issues reduce test reliability.

---

## Critical Issues

### Issue #1: Direct os.Exit in Error Handler

**File**: `validator.go`
**Location**: `validator.go:45`
**Severity**: 🔴 Critical
**Category**: Untestable Code

#### Problem
The `Initialize()` function calls `os.Exit(1)` on configuration error, making the error path completely untestable.

#### Current Code
```go
// Line 45-48
func Initialize(configPath string) *Validator {
    cfg, err := loadConfig(configPath)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to load config: %v\n", err)
        os.Exit(1)  // Cannot test this path
    }
    return &Validator{config: cfg}
}
```

#### Why This is Critical
- Tests cannot verify error handling behavior
- Calling this function in tests would terminate the test process
- Cannot achieve code coverage for error path
- Violates testability principles

#### Impact on Testing
Cannot test the configuration error scenario. Coverage shows line 45-48 as uncovered (red in HTML report).

#### Recommended Fix
```go
func Initialize(configPath string) (*Validator, error) {
    cfg, err := loadConfig(configPath)
    if err != nil {
        return nil, fmt.Errorf("failed to load config: %w", err)
    }
    return &Validator{config: cfg}, nil
}
```

#### Explanation
Return the error to the caller instead of exiting. This allows:
- Testing the error path
- Proper error propagation
- Graceful error handling
- 100% code coverage

#### Test Coverage Impact
- **Current**: Lines 45-48 uncovered (0%)
- **After fix**: Can achieve 100% coverage with error test case

---

### Issue #2: Nil Pointer Dereference

**File**: `validator.go`
**Location**: `validator.go:78`
**Severity**: 🔴 Critical
**Category**: Nil Pointer Risk

#### Problem
Function dereferences pointer without nil check, causing potential panic.

#### Current Code
```go
// Line 78
func (v *Validator) Validate(req *ValidationRequest) ValidationResult {
    currency := v.config.Currencies[req.CurrencyCode]  // Panic if v.config is nil!
    // ...
}
```

#### Why This is Critical
- Runtime panic if `v.config` is nil
- Production bug risk
- Cannot safely test with nil validator

#### Recommended Fix
```go
func (v *Validator) Validate(req *ValidationRequest) (ValidationResult, error) {
    if v == nil || v.config == nil {
        return ValidationResult{}, errors.New("validator not initialized")
    }
    if req == nil {
        return ValidationResult{}, errors.New("nil validation request")
    }

    currency, ok := v.config.Currencies[req.CurrencyCode]
    if !ok {
        return ValidationResult{Valid: false}, nil
    }
    // ...
}
```

---

## User Action Required

Please apply the recommended fixes for Issues #1 and #2 before test suite completion. These issues prevent achieving 100% code coverage and introduce potential bugs.

After fixes are applied, I can write comprehensive tests achieving full coverage.
