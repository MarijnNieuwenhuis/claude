# Project Go New - Bootstrap Go Project Setup

Expert skill for creating new Go microservice projects from the bootstrap template or updating the base bootstrap project itself.

## Overview

This skill manages two primary workflows:
1. **Create New Project**: Sets up a new Go microservice project based on the bootstrap template
2. **Update Bootstrap**: Modifies the base bootstrap project template in `.claude/refs/bootstrap-go-project`

## Mode: Create New Project

When the user requests to create a new Go project, follow these steps:

### Step 1: Gather Requirements

Ask the user for the following information (if not already provided):
- **Project Name**: The service name (e.g., "my-service", "user-manager")
- **Purpose**: Brief description of what the service does
- **Module Path** (optional): Full module path. If not provided, use: `gitlab.com/btcdirect-api/{project-name}`

### Step 2: Validate and Prepare

1. Verify the project name is valid (lowercase, hyphens allowed, no spaces)
2. Check if a directory with the project name already exists in the current working directory
3. If it exists, ask the user if they want to overwrite it or choose a different name
4. Prepare the full module path

### Step 3: Copy Bootstrap Project

1. Copy the entire bootstrap project from `.claude/skills/project-go-new/bootstrap-go-project` to the current working directory with the project name:
   ```bash
   cp -r .claude/skills/project-go-new/bootstrap-go-project ./{project-name}
   ```

2. Remove the vendor directory (will be regenerated):
   ```bash
   rm -rf ./{project-name}/vendor
   ```

### Step 4: Customize the Project

Perform the following customizations in order:

#### 4.1: Update go.mod
- Replace `gitlab.com/btcdirect-api/bootstrap-go-service` with the new module path

#### 4.2: Rename cmd directory
- Rename `cmd/bootstrap-go-service` to `cmd/{project-name}`

#### 4.3: Update all Go import statements
- Find all `.go` files and replace import paths:
  - Old: `gitlab.com/btcdirect-api/bootstrap-go-service`
  - New: `{new-module-path}`

#### 4.4: Update Dockerfile
- Replace references to `bootstrap-go-service` with `{project-name}`
- Specifically update:
  - Binary paths
  - COPY commands
  - Executable names

#### 4.5: Update Makefile
- Replace `bootstrap-go-service` with `{project-name}`
- Update build targets and paths

#### 4.6: Update README.md
- Replace the title with the project name
- Update the description with the purpose provided by the user
- Update all example commands to use the new project name

#### 4.7: Update sonar-project.properties (if exists)
- Update `sonar.projectKey` with the new project name

#### 4.8: Create .env file
Create a `.env` file with placeholders:
```env
# Application
APP_ENV=dev
HTTP_PORT=8080
LOG_LEVEL=debug

# Database
DATABASE_URL=user:password@tcp(localhost:3306)/dbname?parseTime=true

# Sentry
SENTRY_DSN=

# Google Cloud Pub/Sub
PUBSUB_EMULATOR=localhost:8085
PUBSUB_PROJECT=local-project

# Service specific settings
# Add your environment variables here
```

### Step 5: Initialize Dependencies

1. Run `go mod tidy` to clean up dependencies:
   ```bash
   cd ./{project-name} && go mod tidy
   ```

2. Run `go mod vendor` to vendor dependencies:
   ```bash
   cd ./{project-name} && go mod vendor
   ```

### Step 6: Verify Setup

1. Run tests to ensure the project is set up correctly:
   ```bash
   cd ./{project-name} && go test ./...
   ```

2. Run a build to verify everything compiles:
   ```bash
   cd ./{project-name} && go build -o app ./cmd/{project-name}
   ```

3. Report results to the user:
   - Show any errors encountered
   - If successful, provide next steps

### Step 7: Summary

Provide the user with a summary including:
- Project location
- Module path
- Next steps:
  - Configure `.env` with actual values
  - Update database migrations in `internal/db/migrations/`
  - Add business logic
  - Initialize git repository if desired
  - Start development with `make run`

## Mode: Update Bootstrap Project

When the user requests to update the base bootstrap project, follow these steps:

### Step 1: Clarify Intent

Ask the user what they want to update:
- Add new dependencies?
- Modify project structure?
- Update existing files?
- Add new features or patterns?

### Step 2: Make Changes

1. Work directly in `.claude/refs/bootstrap-go-project/`
2. Make the requested modifications
3. Ensure changes follow Go best practices (refer to `.claude/refs/go/`)

### Step 3: Update Dependencies

After making changes:
1. Run `go mod tidy`:
   ```bash
   cd .claude/refs/bootstrap-go-project && go mod tidy
   ```

2. Run `go mod vendor`:
   ```bash
   cd .claude/refs/bootstrap-go-project && go mod vendor
   ```

### Step 4: Verify Changes

1. Run tests:
   ```bash
   cd .claude/refs/bootstrap-go-project && go test ./...
   ```

2. Run build:
   ```bash
   cd .claude/refs/bootstrap-go-project && go build ./cmd/bootstrap-go-service
   ```

3. Ensure everything compiles and tests pass

### Step 5: Update Skill Copy

After successfully updating the bootstrap project in `.claude/refs/bootstrap-go-project`:
1. Sync changes to the skill directory:
   ```bash
   rm -rf .claude/skills/project-go-new/bootstrap-go-project
   cp -r .claude/refs/bootstrap-go-project .claude/skills/project-go-new/bootstrap-go-project
   ```

### Step 6: Summary

Inform the user:
- What was changed
- Test and build results
- Remind them they can commit changes to the `.claude` repo
- Suggest updating the bootstrap README.md if significant changes were made

## Best Practices

### Code Quality
- Follow idiomatic Go conventions
- Reference `.claude/refs/go/` for Go best practices
- Keep code simple, efficient, and maintainable (per CLAUDE.md)
- Add appropriate error handling

### Project Structure
- Maintain the standard project layout:
  - `cmd/` for application entry points
  - `internal/` for private application code
  - `internal/app/` for application initialization
  - `internal/http/` for HTTP handlers and server
  - `internal/db/` for database logic
  - `internal/messenger/` for pub/sub messaging

### Dependencies
- Always run `go mod tidy` before `go mod vendor`
- Keep vendor directory up to date
- Document any new required environment variables

### Testing
- Always verify the setup with tests and build
- Report clear error messages if something fails
- Provide actionable next steps

## Common Issues and Solutions

### Issue: Import Path Errors
**Solution**: Ensure all import statements are updated consistently across all `.go` files

### Issue: Build Failures
**Solution**: Check that:
- `go.mod` has the correct module path
- All imports match the new module path
- Dependencies are properly vendored
- No leftover references to `bootstrap-go-service`

### Issue: Test Failures
**Solution**: Review test files for hardcoded paths or references to the old project name

## Examples

### Example 1: Create New Project
```
User: Create a new Go project called "user-service" for managing user accounts
Assistant: I'll create a new Go microservice project for managing user accounts.

[Follows all steps above]

Summary:
- Project created at: ./user-service
- Module path: gitlab.com/btcdirect-api/user-service
- Tests: PASSED
- Build: SUCCESSFUL

Next steps:
1. Configure .env with your database credentials
2. Add your database migrations to internal/db/migrations/
3. Implement your user management business logic
4. Run `make run` to start the service
```

### Example 2: Update Bootstrap
```
User: Add graceful shutdown support to the bootstrap project
Assistant: I'll add graceful shutdown support to the base bootstrap project.

[Makes changes in .claude/refs/bootstrap-go-project]
[Updates dependencies]
[Runs tests and build]
[Syncs to skill directory]

Summary:
- Added graceful shutdown with signal handling
- Updated dependencies
- Tests: PASSED
- Build: SUCCESSFUL

The bootstrap project has been updated. You can now commit these changes to your .claude repository.
```

## Notes

- The bootstrap project includes BTCDirect-specific modules and patterns
- The skill works with the local copy in `.claude/skills/project-go-new/bootstrap-go-project`
- Changes to the bootstrap should be made in `.claude/refs/bootstrap-go-project` and then synced
- Always verify setup with tests and build before completing
- The `.claude` directory is a git repository that can be committed separately

## References

- Bootstrap project location: `.claude/refs/bootstrap-go-project`
- Skill bootstrap copy: `.claude/skills/project-go-new/bootstrap-go-project`
- Go best practices: `.claude/refs/go/`
- Project instructions: `CLAUDE.md`