# Project Go New Skill

Bootstrap new Go microservice projects or update the base bootstrap template.

## Usage

Invoke this skill when the user wants to:
- Create a new Go microservice project
- Update the base bootstrap project template

## Examples

### Creating a New Project
```
User: Create a new Go project called payment-service
User: Set up a new microservice for handling orders
User: I need a new Go service
```

### Updating Bootstrap
```
User: Add logging middleware to the bootstrap project
User: Update the bootstrap template with better error handling
User: Modify the base Go project template
```

## What This Skill Does

### For New Projects:
1. Copies the bootstrap template
2. Customizes the project (module path, names, etc.)
3. Creates environment configuration
4. Updates dependencies with `go mod vendor`
5. Verifies with tests and build

### For Bootstrap Updates:
1. Makes changes to `.claude/refs/bootstrap-go-project`
2. Updates dependencies
3. Runs tests and build to verify
4. Syncs changes to the skill's bootstrap copy

## Files

- `SKILL.md` - Main skill instructions for Claude
- `bootstrap-go-project/` - Local copy of the bootstrap template
- `README.md` - This file

## Notes

- The skill uses the bootstrap project from `.claude/skills/project-go-new/bootstrap-go-project`
- Bootstrap updates are made in `.claude/refs/bootstrap-go-project` first, then synced
- All projects are created in the current working directory
- Module paths follow the pattern: `gitlab.com/btcdirect-api/{project-name}`