---
name: project-updater
description: Update the entire project including main repository and all dependent git repositories in ./refs, running composer install for PHP projects. Use when user asks to update project, refresh dependencies, or sync all repos.
---

# Project Updater

You are a project maintenance specialist. Your task is to update the entire project by pulling latest changes from git repositories and updating dependencies.

## What This Skill Does

This skill performs a complete project update in the following order:
1. Updates the main project repository (git pull)
2. Updates all dependent git repositories in `./refs/` directory
3. Detects PHP projects (by presence of `composer.json`)
4. Runs `composer install` for each PHP project

## Update Process

### Step 1: Update Main Project

1. Check if the current directory is a git repository
2. Show current git status
3. Run `git pull` to update the main project
4. Report the result (files updated, commits pulled, or already up-to-date)

### Step 2: Discover Dependent Repositories

1. Scan the `./refs/` directory for subdirectories
2. Identify which subdirectories contain git repositories (have `.git` directory)
3. List all found repositories

### Step 3: Update Each Dependent Repository

For each repository found in `./refs/`:

1. Navigate to the repository directory
2. Show current branch and status
3. Run `git pull` to update
4. Report the result (success, conflicts, or errors)
5. If any errors occur, stop the process and report which repository failed

### Step 4: Update PHP Project Dependencies

For each repository in `./refs/`:

1. Check if `composer.json` exists in the repository
2. If found, it's a PHP project - run `composer install`
3. Show composer output (packages installed, updated, or already up-to-date)
4. If `composer install` fails, **stop immediately** and report:
   - Which project failed
   - The error message
   - The exit code

### Step 5: Generate Update Report

After all updates complete successfully:

1. Summarize what was updated:
   - Main project status
   - Number of repos updated in `./refs/`
   - Which repos had new commits
   - Which PHP projects had dependencies updated
2. Report total time taken
3. Confirm all updates completed successfully

## Error Handling

### Git Pull Failures

If `git pull` fails for any repository:
- Stop the update process
- Report which repository failed
- Show the error message
- Suggest possible solutions (conflicts, network issues, etc.)
- Do NOT continue to other repositories

### Composer Install Failures

If `composer install` fails for any PHP project:
- Stop the update process immediately
- Report which project failed
- Show the complete error output
- Suggest checking `composer.json` or running `composer diagnose`
- Do NOT continue to other projects

### Missing Directories

If `./refs/` directory doesn't exist:
- Report that no dependent repositories were found
- Continue with just the main project update

## Commands to Execute

### Main Project Update
```bash
git pull
```

### Find Git Repos in ./refs
```bash
find ./refs -maxdepth 1 -type d -name ".git" | sed 's|/.git||'
```

### For Each Repo in ./refs
```bash
cd ./refs/[repo-name] && git pull
```

### Check for PHP Projects
```bash
find ./refs -maxdepth 2 -name "composer.json" -type f
```

### Run Composer Install
```bash
cd ./refs/[repo-name] && composer install
```

## Output Format

Provide a clear, structured report:

```
# Project Update Report

## Main Project
Repository: [current directory]
Branch: [branch name]
Status: ✓ Updated successfully | Already up-to-date | ✗ Failed
Changes: [X files changed, Y commits pulled]

## Dependent Repositories in ./refs

### [repo-name-1]
Branch: [branch name]
Status: ✓ Updated successfully | Already up-to-date | ✗ Failed
Changes: [X files changed, Y commits pulled]
PHP Project: [Yes/No]
Composer: [✓ Dependencies updated | Already up-to-date | N/A]

### [repo-name-2]
[Same format]

## Summary
- Main project: [status]
- Dependent repos found: [count]
- Repos updated: [count]
- PHP projects found: [count]
- Composer updates: [count]
- Total time: [duration]

✓ All updates completed successfully
```

## Examples

### Example 1: Successful Update with PHP Projects

**User Request**: "Update project"

**Process**:
1. Run `git pull` in main project → 3 files updated
2. Find repos in `./refs/` → Found: addressvalidator
3. Update `./refs/addressvalidator` → 5 files updated
4. Check for `composer.json` → Found in addressvalidator
5. Run `composer install` in addressvalidator → 2 packages updated

**Output**:
```
# Project Update Report

## Main Project
Repository: /Users/marijnnieuwenhuis/Docker/devenv/apps/addressvalidator-go
Branch: master
Status: ✓ Updated successfully
Changes: 3 files changed, 2 commits pulled

## Dependent Repositories in ./refs

### addressvalidator
Branch: main
Status: ✓ Updated successfully
Changes: 5 files changed, 3 commits pulled
PHP Project: Yes
Composer: ✓ Dependencies updated (2 packages updated)

## Summary
- Main project: ✓ Updated
- Dependent repos found: 1
- Repos updated: 1
- PHP projects found: 1
- Composer updates: 1
- Total time: 12 seconds

✓ All updates completed successfully
```

### Example 2: Already Up-to-Date

**User Request**: "Refresh all dependencies"

**Process**:
1. Run `git pull` in main project → Already up-to-date
2. Find repos in `./refs/` → Found: addressvalidator
3. Update `./refs/addressvalidator` → Already up-to-date
4. Run `composer install` → Dependencies already installed

**Output**:
```
# Project Update Report

## Main Project
Repository: /Users/marijnnieuwenhuis/Docker/devenv/apps/addressvalidator-go
Branch: master
Status: Already up-to-date
Changes: None

## Dependent Repositories in ./refs

### addressvalidator
Branch: main
Status: Already up-to-date
Changes: None
PHP Project: Yes
Composer: Already up-to-date

## Summary
- Main project: Already up-to-date
- Dependent repos found: 1
- Repos updated: 0
- PHP projects found: 1
- Composer updates: 0
- Total time: 4 seconds

✓ All systems up-to-date
```

### Example 3: Composer Install Failure

**User Request**: "Update the project"

**Process**:
1. Run `git pull` in main project → Updated successfully
2. Update `./refs/addressvalidator` → Updated successfully
3. Run `composer install` → **FAILED**

**Output**:
```
# Project Update Report

## Main Project
Repository: /Users/marijnnieuwenhuis/Docker/devenv/apps/addressvalidator-go
Branch: master
Status: ✓ Updated successfully
Changes: 2 files changed, 1 commit pulled

## Dependent Repositories in ./refs

### addressvalidator
Branch: main
Status: ✓ Updated successfully
Changes: 3 files changed, 2 commits pulled
PHP Project: Yes
Composer: ✗ FAILED

## Error Details

Project: ./refs/addressvalidator
Command: composer install
Exit Code: 1

Error Output:
Loading composer repositories with package information
Installing dependencies (including require-dev) from lock file
Your requirements could not be resolved to an installable set of packages.

Problem 1
  - Installation request for vendor/package ^2.0 -> satisfiable by vendor/package[2.0.0].
  - vendor/package 2.0.0 requires php ^8.1 -> your PHP version (8.0.28) does not satisfy that requirement.

## Troubleshooting Steps

1. Check PHP version: php --version
2. Review composer.json requirements
3. Run: composer diagnose
4. Update composer: composer self-update
5. Clear cache: composer clear-cache

✗ Update process stopped due to error
```

### Example 4: Git Conflict

**User Request**: "Update project"

**Process**:
1. Run `git pull` in main project → Updated successfully
2. Update `./refs/addressvalidator` → **CONFLICT**

**Output**:
```
# Project Update Report

## Main Project
Repository: /Users/marijnnieuwenhuis/Docker/devenv/apps/addressvalidator-go
Branch: master
Status: ✓ Updated successfully
Changes: 1 file changed, 1 commit pulled

## Dependent Repositories in ./refs

### addressvalidator
Branch: main
Status: ✗ FAILED - Merge conflict
Error: Cannot pull, you have uncommitted changes

Git Status:
M  src/Validator.php
?? temp.txt

## Error Details

Repository: ./refs/addressvalidator
Command: git pull
Exit Code: 1

Error Output:
error: Your local changes to the following files would be overwritten by merge:
    src/Validator.php
Please commit your changes or stash them before you merge.
Aborting

## Troubleshooting Steps

1. Navigate to repository: cd ./refs/addressvalidator
2. View changes: git status
3. Options:
   - Commit changes: git add . && git commit -m "message"
   - Stash changes: git stash
   - Discard changes: git checkout -- .
4. Then retry update

✗ Update process stopped due to git conflict
```

## Success Criteria

A successful update means:
- ✓ Main project git pull completed without errors
- ✓ All repositories in `./refs/` git pull completed without errors
- ✓ All PHP projects composer install completed without errors
- ✓ No conflicts or merge issues
- ✓ Clear report generated

## When to Use This Skill

Use this skill when the user says:
- "Update project"
- "Update the project"
- "Refresh dependencies"
- "Pull all repos"
- "Sync everything"
- "Update all dependencies"
- "Get latest changes"
- "Update main and refs"

## Important Notes

1. **Stop on Error**: If any step fails, stop immediately and report
2. **Clear Output**: Always show what's happening at each step
3. **Status Reporting**: Report status of each repository individually
4. **Time Tracking**: Note how long the update takes
5. **No Assumptions**: Don't assume paths or branch names
6. **Safety First**: Always show git status before pulling
7. **Verify Composer**: Check that composer is available before running install

## Pre-Update Checks

Before starting the update:
1. Verify git is installed: `git --version`
2. Verify composer is installed (if PHP projects exist): `composer --version`
3. Check for uncommitted changes in main project
4. Warn user if there are uncommitted changes
