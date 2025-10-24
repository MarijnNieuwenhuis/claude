# Detailed Testing Workflow Steps

Complete step-by-step workflow for achieving 100% test coverage with progress tracking.

## Step 0: Initialize Coverage Tracking (First Time Only)

**When**: At the very beginning, before any testing work starts
**Condition**: `TEST-COVERAGE-TODO.md` does not exist in project root

### Actions

1. **Run baseline coverage**:
   ```bash
   make test
   ```

2. **Generate detailed coverage report**:
   ```bash
   go test -coverprofile=coverage.out ./internal/... ./pkg/...
   go tool cover -func=coverage.out > coverage-func.txt
   ```

3. **Analyze codebase structure**:
   ```bash
   # Find all Go files (excluding tests)
   find ./internal ./pkg -name "*.go" -not -name "*_test.go"

   # Count packages
   go list ./internal/... ./pkg/...
   ```

4. **Parse coverage data** to extract:
   - Overall coverage percentage
   - Per-package coverage
   - Per-file coverage and uncovered functions
   - Existing test files

5. **Create TEST-COVERAGE-TODO.md**:
   - Copy template from `.claude/skills/tester-unittest-go/templates/TEST-COVERAGE-TODO.md.template`
   - Replace all `{placeholders}` with actual data
   - List all packages, files, and functions
   - Calculate priorities:
     - **High**: Core business logic (validator, rules, chain, currency)
     - **Medium**: Configuration and supporting code
     - **Low**: Utilities and helpers

6. **Categorize files**:
   - **✅ Completed**: Coverage 100%
   - **🔄 In Progress**: Coverage 1-99%, has tests
   - **⏳ Not Started**: Coverage 0%, no tests

7. **Present to user**:
   ```
   Created TEST-COVERAGE-TODO.md

   📊 Initial Coverage Analysis:
   - Overall: 42.3%
   - Packages: 8 total
     - ✅ Complete: 1 (regex)
     - 🔄 In Progress: 3 (validator, config, chain)
     - ⏳ Not Started: 4 (currency, rules, utils, etc.)
   - Files: 23 total
     - ✅ Complete: 3
     - 🔄 In Progress: 8
     - ⏳ Not Started: 12
   - Functions: 156 total
     - ✅ Tested: 67
     - ⏳ Untested: 89

   📋 Recommended Testing Order:
   1. Complete validator.go (currently 60%)
   2. Complete config/loader.go (currently 80%)
   3. Start chain/registry.go (currently 0%)
   4. Start currency/registry.go (currently 0%)

   🎯 Next Action:
   Starting with internal/validator/validator.go
   Target: Reach 100% coverage
   ```

---

## Step 1: Check TODO and Select Work

**When**: At the start of EVERY testing session
**Condition**: `TEST-COVERAGE-TODO.md` exists

### Actions

1. **Read TEST-COVERAGE-TODO.md** to review:
   ```markdown
   Check these sections:
   - Summary: Overall progress
   - Package Status: What's done, in progress, pending
   - Detailed Function Coverage: Specific functions needing tests
   - Coverage Blockers: Known issues
   - Next Steps: Previous plan
   - Session Progress Log: Last session notes
   ```

2. **Determine current focus**:
   - **Priority 1**: Continue any 🔄 In Progress work
   - **Priority 2**: Start highest priority ⏳ Not Started work
   - **Priority 3**: Address blockers if user fixed production code

3. **Update TEST-COVERAGE-TODO.md status**:
   ```markdown
   Update "Next Steps" section:
   ### Immediate Actions
   1. ⚡ Working on `internal/validator/validator.go`
      - Target: 100% coverage (currently 60%)
      - Functions to test:
        - [x] New() - Already tested
        - [ ] Validate() - Needs error path tests
        - [ ] preValidate() - Not tested yet
        - [ ] postValidate() - Not tested yet
   ```

4. **Add session start note**:
   ```markdown
   Add to "Session Progress Log":
   ### Session {N} - {date} - IN PROGRESS
   **Focus**: internal/validator/validator.go
   **Starting Coverage**: 60%
   **Goal**: Achieve 100% coverage
   ```

---

## Step 2: Analyze Target Code

**When**: Before writing any tests for a file
**Condition**: Target file selected from TODO

### Actions

1. **Read the production file completely**:
   - Understand all functions and their purpose
   - Identify function parameters and return types
   - Note all error conditions
   - Understand dependencies (interfaces, external packages)
   - Identify business logic vs. infrastructure code

2. **Run coverage for this specific file**:
   ```bash
   go test -coverprofile=coverage.out ./internal/validator/
   go tool cover -html=coverage.out
   ```

3. **Analyze HTML coverage report**:
   - Open in browser
   - Identify RED lines (uncovered)
   - Note all branches (if/else, switch cases)
   - Check error returns
   - Look for early returns

4. **List uncovered functions**:
   ```
   In validator.go:
   - Validate() - 60% covered
     - Missing: nil request check (line 45)
     - Missing: error path (lines 67-70)
     - Missing: edge case (line 89)
   - preValidate() - 0% covered
     - No tests at all
   - postValidate() - 0% covered
     - No tests at all
   ```

5. **Identify test requirements**:
   - **Mocks needed**: List interfaces to mock
   - **Test data**: What inputs needed for tests
   - **Edge cases**: Boundary conditions, empty/nil values
   - **Error cases**: All error returns
   - **Concurrency**: Any goroutines or shared state

6. **Plan test structure**:
   ```
   Test Plan for validator.go:

   1. TestNew() - Constructor
      - Valid config paths
      - Invalid config paths
      - Missing files

   2. TestValidate() - Main function
      - Happy path: valid request
      - Error: nil request
      - Error: invalid currency
      - Error: validation fails
      - Edge: empty address
      - Edge: max length address

   3. TestPreValidate() - Helper
      - X-address detection
      - Valid formats
      - Invalid formats

   4. TestPostValidate() - Helper
      - Additional validations
      - Format cleanup
   ```

---

## Step 3: Write Comprehensive Tests

**When**: After analyzing code and planning tests
**Condition**: Test plan created

### Actions

1. **Create or open test file**:
   ```bash
   # If doesn't exist, create it
   touch internal/validator/validator_test.go
   ```

2. **Set up test package and imports**:
   ```go
   package validator

   import (
       "errors"
       "reflect"
       "testing"
   )
   ```

3. **Write table-driven tests** for each function:
   ```go
   func TestValidate(t *testing.T) {
       // Create validator instance for tests
       validator, err := New("../../config/chains.yaml", "../../config/currencies.yaml")
       if err != nil {
           t.Fatalf("Failed to create validator: %v", err)
       }

       tests := []struct {
           name      string
           request   ValidationRequest
           want      ValidationResult
           wantErr   bool
       }{
           {
               name: "valid BTC address",
               request: ValidationRequest{
                   CurrencyCode: "BTC",
                   Address:      "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
                   NetworkMode:  Mainnet,
               },
               want: ValidationResult{Valid: true},
               wantErr: false,
           },
           {
               name: "nil request",
               request: ValidationRequest{},
               want: ValidationResult{Valid: false},
               wantErr: true,
           },
           // ... more test cases
       }

       for _, tt := range tests {
           t.Run(tt.name, func(t *testing.T) {
               got, err := validator.Validate(tt.request)
               if (err != nil) != tt.wantErr {
                   t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
                   return
               }
               if !reflect.DeepEqual(got, tt.want) {
                   t.Errorf("Validate() = %v, want %v", got, tt.want)
               }
           })
       }
   }
   ```

4. **Create test helpers** for common operations:
   ```go
   func assertEqual(t *testing.T, got, want interface{}) {
       t.Helper()
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

5. **Create mocks** for interfaces (if needed):
   ```go
   type MockChainRegistry struct {
       GetChainFunc func(code string) (*Chain, error)
   }

   func (m *MockChainRegistry) GetChain(code string) (*Chain, error) {
       if m.GetChainFunc != nil {
           return m.GetChainFunc(code)
       }
       return nil, errors.New("not implemented")
   }
   ```

6. **Add test data** files (if needed):
   ```bash
   mkdir -p internal/validator/testdata
   # Add test YAML/JSON files
   ```

7. **Test all scenarios**:
   - ✅ Happy path (normal operation)
   - ✅ Error cases (all error returns)
   - ✅ Edge cases (boundaries, empty, nil, max values)
   - ✅ Concurrent access (if applicable)

---

## Step 4: Verify Coverage and Update TODO

**When**: After writing tests for a file or function
**Condition**: Tests are passing

### Actions

1. **Run tests**:
   ```bash
   go test -v ./internal/validator/
   ```

2. **Generate coverage**:
   ```bash
   make test
   ```

3. **Check coverage in terminal**:
   ```bash
   go tool cover -func=coverage.out | grep validator.go
   ```

4. **Open HTML coverage report**:
   ```bash
   go tool cover -html=coverage.out
   ```

5. **Analyze coverage**:
   - Identify remaining RED lines
   - Check if all branches are GREEN
   - Verify error paths covered
   - Confirm edge cases covered

6. **If coverage < 100%**:
   - List specific uncovered lines
   - Identify why they're uncovered
   - Write additional tests
   - Repeat steps 1-5

7. **If coverage = 100%**:
   - Celebrate! 🎉
   - Proceed to update TODO

8. **Update TEST-COVERAGE-TODO.md**:
   ```markdown
   Update these sections:

   ### Summary
   - Overall Coverage: 65.3% → 78.5% (+13.2%)
   - Files with Tests: 12/23
   - Functions Tested: 89/156 (+22)

   ### Package Status
   #### `internal/validator` - 100% ✅ (was 60% 🔄)
   - Files: 1/1 complete
   - Functions: 4/4 tested
   - Last Updated: 2025-01-24

   ### Detailed Function Coverage
   #### File: `validator.go` (100%) ✅
   | Function | Status | Coverage |
   |----------|--------|----------|
   | New() | ✅ Done | 100% |
   | Validate() | ✅ Done | 100% |
   | preValidate() | ✅ Done | 100% |
   | postValidate() | ✅ Done | 100% |

   ### Session Progress Log
   ### Session 3 - 2025-01-24 ✅ COMPLETED
   **Focus**: internal/validator/validator.go
   **Coverage Change**: 60% → 100% (+40%)

   **Tests Added**:
   - ✅ TestNew - 3 test cases
   - ✅ TestValidate - 12 test cases
   - ✅ TestPreValidate - 5 test cases
   - ✅ TestPostValidate - 4 test cases

   **Coverage Achieved**:
   - ✅ validator.go - 100%
   - ✅ internal/validator package - 100%

   **Next Session**:
   - Start internal/validator/config/loader.go
   ```

---

## Step 5: Implement Mutation Testing

**When**: After a file reaches 100% coverage
**Condition**: File coverage verified at 100%

### Actions

1. **Document mutation test plan** in TEST-COVERAGE-TODO.md:
   ```markdown
   ### Mutation Testing Status

   #### `internal/validator/validator.go` - Testing Mutations

   **Planned Mutations**:
   1. Line 45: Change `== nil` to `!= nil`
      - Expected: TestValidate nil check should fail
   2. Line 67: Change `>` to `>=`
      - Expected: TestValidate boundary should fail
   3. Line 89: Remove error check
      - Expected: TestValidate error path should fail
   ```

2. **Install mutation testing tool** (if not installed):
   ```bash
   go install github.com/zimmski/go-mutesting/cmd/go-mutesting@latest
   ```

3. **Run automated mutation testing**:
   ```bash
   go-mutesting ./internal/validator/validator.go
   ```

4. **Capture results**:
   ```
   Example output:
   PASS validator.go:45 - changed == to !=
   PASS validator.go:67 - changed > to >=
   FAIL validator.go:89 - removed error check  [SURVIVED]

   Mutation Score: 20/22 (90.9%)
   ```

5. **For survived mutations**:
   - Analyze why test didn't catch it
   - Add specific test case
   - Rerun mutation test
   - Verify mutation now killed

6. **Update TEST-COVERAGE-TODO.md**:
   ```markdown
   #### `internal/validator/validator.go` - ✅ Mutation Tested
   - **Mutation Score**: 95.5% (21/22 mutations killed)
   - **Date**: 2025-01-24
   - **Status**: ✅ Passed (>80% target)

   **Mutations Tested**:
   - ✅ Conditional boundaries (10/10)
   - ✅ Logical operators (5/5)
   - ✅ Return values (4/4)
   - ⚠️ Error check removal (1 survived, acceptable - logging only)

   **Action**: Marked as complete
   ```

---

## Step 6: Report Production Code Issues

**When**: When you discover code that cannot be tested
**Condition**: Found untestable code, bugs, or problems

### Actions

1. **Document in TEST-COVERAGE-TODO.md** "Coverage Blockers":
   ```markdown
   ### Coverage Blockers

   #### Issue #1: Untestable Code - Direct os.Exit()
   - **File**: `internal/app/init.go`
   - **Location**: Line 78
   - **Function**: `Initialize()`
   - **Problem**: Direct `os.Exit(1)` call prevents testing error path
   - **Impact**: Cannot achieve 100% coverage for this file
   - **Blocks**: 15% of file coverage
   - **Status**: ⚠️ Reported to user (2025-01-24)
   - **Priority**: High

   **Current Code**:
   ```go
   func Initialize() {
       if err := setup(); err != nil {
           fmt.Println(err)
           os.Exit(1)  // UNTESTABLE
       }
   }
   ```

   **Recommended Fix**:
   ```go
   func Initialize() error {
       if err := setup(); err != nil {
           return fmt.Errorf("setup failed: %w", err)
       }
       return nil
   }
   ```

   **After Fix**: Can write tests for error path
   ```

2. **Create detailed issue report** using template:
   ```markdown
   # Production Code Issues Report

   **Package**: `internal/app`
   **Date**: 2025-01-24
   **Severity**: 🔴 Critical

   ## Issue #1: Untestable Code - os.Exit() Call

   ### Problem
   The `Initialize()` function in `init.go:78` calls `os.Exit(1)` directly,
   making the error path completely untestable.

   ### Current Code
   [Show code block]

   ### Why This Blocks Testing
   - Tests cannot verify error handling
   - Calling this in tests would terminate test process
   - Coverage report shows uncovered lines

   ### Recommended Fix
   [Show fixed code]

   ### Impact
   - Prevents 15% file coverage
   - Prevents testing error scenarios
   - Blocks overall 100% goal

   ### User Action Required
   Please apply the recommended fix, then request test coverage for this file.
   ```

3. **Show report to user**

4. **Mark affected functions** in TODO:
   ```markdown
   | Function | Status | Coverage | Notes |
   |----------|--------|----------|-------|
   | Initialize() | ⚠️ Blocked | 0% | Untestable - Issue #1 |
   ```

5. **Continue with other testable code**

---

## Step 7: Session Completion

**When**: At the end of each testing session
**Condition**: Finishing work for now

### Actions

1. **Run final coverage check**:
   ```bash
   make test
   go tool cover -func=coverage.out | tail -1
   ```

2. **Update TEST-COVERAGE-TODO.md** session log:
   ```markdown
   ### Session 3 - 2025-01-24 ✅ COMPLETED
   **Focus**: internal/validator package
   **Duration**: ~2 hours
   **Starting Coverage**: 60.2%
   **Ending Coverage**: 78.5%
   **Improvement**: +18.3%

   **Tests Added**:
   - ✅ TestNew (3 cases)
   - �� TestValidate (12 cases)
   - ✅ TestPreValidate (5 cases)
   - ✅ TestPostValidate (4 cases)
   - ✅ TestValidatorNew (8 cases)

   **Files Completed** (100% coverage):
   - ✅ internal/validator/validator.go

   **Files Improved**:
   - 🔄 internal/validator/config/loader.go (80% → 90%)

   **Mutation Testing**:
   - ✅ validator.go - 95.5% mutation score

   **Issues Found**:
   - ⚠️ Issue #1: os.Exit() in init.go - Reported

   **Statistics**:
   - Test files created: 1
   - Test functions written: 4
   - Test cases added: 24
   - Lines of test code: 187

   **Next Session Plan**:
   1. Complete config/loader.go (need 10% more)
   2. Start chain/registry.go
   3. Target: Reach 85% overall coverage
   ```

3. **Update summary statistics**:
   ```markdown
   ## Summary

   | Metric | Current | Target | Progress |
   |--------|---------|--------|----------|
   | Overall Coverage | 78.5% | 100% | ████████░░ 78% |
   | Files with Tests | 13/23 | 23 | ██████░░░░ 57% |
   | Functions Tested | 95/156 | 156 | ██████░░░░ 61% |
   | Packages at 100% | 2/8 | 8 | ███░░░░░░░ 25% |
   ```

4. **Show user summary**:
   ```
   ✅ Session Complete!

   📊 Progress:
   - Coverage: 60.2% → 78.5% (+18.3%)
   - Files completed: 1 (validator.go)
   - Tests added: 24 test cases
   - Mutation score: 95.5%

   🎯 Achievements:
   - ✅ internal/validator package now at 100%
   - ✅ Mutation testing passed
   - ✅ 24 comprehensive test cases added

   ⚠️ Issues Found:
   - Issue #1: Untestable os.Exit() in init.go (reported)

   📋 Next Steps:
   1. Complete config/loader.go (90% → 100%)
   2. Start chain/registry.go (0% → target 100%)
   3. Continue toward 85% overall coverage

   See TEST-COVERAGE-TODO.md for full details.
   ```

---

## Step 8: Achieving 100% Coverage

**When**: When all packages reach 100% coverage
**Condition**: Overall coverage = 100%

### Actions

1. **Run final verification**:
   ```bash
   make test
   ```

2. **Check all packages**:
   ```bash
   go test -cover ./internal/... ./pkg/...
   ```

3. **Verify output**:
   ```
   ok      internal/validator              0.234s  coverage: 100.0% of statements
   ok      internal/validator/config       0.123s  coverage: 100.0% of statements
   ok      internal/validator/chain        0.156s  coverage: 100.0% of statements
   ok      internal/validator/currency     0.145s  coverage: 100.0% of statements
   ok      internal/validator/regex        0.089s  coverage: 100.0% of statements
   ok      pkg/utils                       0.067s  coverage: 100.0% of statements

   Overall coverage: 100.0%
   ```

4. **Run final mutation tests** on all packages:
   ```bash
   for pkg in $(go list ./internal/... ./pkg/...); do
       echo "Testing $pkg"
       go-mutesting ./$pkg
   done
   ```

5. **Update TEST-COVERAGE-TODO.md** final status:
   ```markdown
   # Test Coverage Progress Tracker

   **Status**: 🎉 **COMPLETED** - 100% Coverage Achieved!

   ## Summary

   | Metric | Current | Target | Progress |
   |--------|---------|--------|----------|
   | Overall Coverage | 100.0% | 100% | ██████████ 100% |
   | Files with Tests | 23/23 | 23 | ██████████ 100% |
   | Functions Tested | 156/156 | 156 | ██████████ 100% |
   | Packages at 100% | 8/8 | 8 | ██████████ 100% |

   ## Package Coverage Status

   ### ✅ All Packages Completed (100%)

   - ✅ internal/validator - 100% - Mutation Score: 95.5%
   - ✅ internal/validator/config - 100% - Mutation Score: 92.3%
   - ✅ internal/validator/chain - 100% - Mutation Score: 88.7%
   - ✅ internal/validator/currency - 100% - Mutation Score: 91.2%
   - ✅ internal/validator/regex - 100% - Mutation Score: 96.8%
   - ✅ internal/validator/rules - 100% - Mutation Score: 89.4%
   - ✅ pkg/utils - 100% - Mutation Score: 87.6%
   - ✅ pkg/helpers - 100% - Mutation Score: 90.1%

   ## Final Statistics

   - **Total Test Files**: 23
   - **Total Test Functions**: 187
   - **Total Test Cases**: 1,245
   - **Lines of Test Code**: 8,934
   - **Coverage Improvement**: 42.3% → 100.0% (+57.7%)
   - **Sessions**: 15
   - **Time**: ~30 hours

   ## Completion Summary

   🎉 **100% Test Coverage Achieved!**

   All packages, files, and functions now have comprehensive test coverage.
   Mutation testing shows high test quality (80%+ on all packages).

   ### Ready to Delete This File

   This TEST-COVERAGE-TODO.md file has served its purpose.
   You may now delete it.

   ### Recommendations

   1. ✅ Add coverage badge to README
   2. ✅ Set up CI/CD to maintain 100% coverage
   3. ✅ Run mutation tests in CI pipeline
   4. ✅ Enforce coverage requirements for new code

   **Date Completed**: 2025-01-24
   ```

6. **Show user final report**:
   ```
   🎉🎉🎉 CONGRATULATIONS! 🎉🎉🎉

   100% Test Coverage Achieved!

   📊 Final Statistics:
   - Overall Coverage: 100.0%
   - Packages: 8/8 at 100%
   - Files: 23/23 fully tested
   - Functions: 156/156 with tests
   - Test Cases: 1,245 comprehensive tests

   ✅ Quality Metrics:
   - All mutation scores > 80%
   - Average mutation score: 91.4%
   - No production code modified
   - All Go best practices followed

   📈 Improvement:
   - Starting: 42.3% coverage
   - Ending: 100.0% coverage
   - Improvement: +57.7%
   - Sessions: 15

   🗑️ Ready to Delete:
   TEST-COVERAGE-TODO.md has served its purpose and can be deleted.

   📋 Next Steps:
   1. Delete TEST-COVERAGE-TODO.md
   2. Add coverage badge to README
   3. Set up CI to maintain 100% coverage
   4. Celebrate! 🎉

   Excellent work on achieving complete test coverage!
   ```

7. **Await user confirmation** to delete TEST-COVERAGE-TODO.md

---

## Summary: The Complete Cycle

```
Step 0: Initialize
    ↓
[Create TEST-COVERAGE-TODO.md with baseline]
    ↓
Step 1: Check TODO
    ↓
[Select next file/function to test]
    ↓
Step 2: Analyze Code
    ↓
[Understand code, identify gaps]
    ↓
Step 3: Write Tests
    ↓
[Create comprehensive test cases]
    ↓
Step 4: Verify & Update
    ↓
[Check coverage, update TODO]
    ↓
Coverage < 100%? → Back to Step 3
Coverage = 100%? → Step 5
    ↓
Step 5: Mutation Testing
    ↓
[Run mutations, verify test quality]
    ↓
Step 6: Report Issues (if any)
    ↓
[Document blockers, report to user]
    ↓
Step 7: Session Complete
    ↓
[Update TODO, show summary]
    ↓
More work? → Back to Step 1
All done? → Step 8
    ↓
Step 8: 100% Complete!
    ↓
[Final verification, delete TODO]
    ↓
🎉 SUCCESS! 🎉
```
