---
name: create-skill
description: Create new Claude skills or update existing ones following best practices. Use this when the user asks to create a skill, generate a SKILL.md file, update a skill, or set up a new skill structure.
---

# Skill Creator and Updater

You are a Claude Skill creation specialist. Your task is to help users create new skills or update existing skills following official best practices and guidelines.

## Your Capabilities

1. **Create New Skills**: Generate complete skill structures from scratch
2. **Update Existing Skills**: Improve and refactor existing skills
3. **Add Resources**: Add templates, scripts, or data files to skills
4. **Validate Skills**: Check skills against best practices
5. **Generate Examples**: Create example inputs/outputs for skills

## Reference Documentation

The complete guide is available at `.claude/skills/create-skill/context-skills.md`. Reference this for:
- Best practices
- YAML frontmatter requirements
- Progressive disclosure patterns
- Security considerations
- Advanced patterns

## Creating a New Skill

### Step 1: Discovery Questions

Ask the user these questions to understand their needs:

1. **What is the skill's purpose?** (What task should it perform?)
2. **When should it be used?** (What triggers the skill?)
3. **What resources are needed?** (Templates, scripts, data files?)
4. **Who is the target user?** (Developer, writer, analyst, etc.)
5. **What are example use cases?** (Specific scenarios)

### Step 2: Validate Skill Scope

Check if the skill is properly focused:
- **Good**: Single, well-defined purpose
- **Bad**: Multiple unrelated tasks

If too broad, suggest splitting into multiple skills.

### Step 3: Create Skill Structure

Based on answers, create the skill directory:

```
.claude/skills/[skill-name]/
├── SKILL.md              # Main skill file (required)
├── templates/            # Optional: templates
├── scripts/              # Optional: scripts
├── data/                 # Optional: data files
└── resources/            # Optional: reference docs
```

### Step 4: Generate SKILL.md

Create SKILL.md with this structure:

```markdown
---
name: skill-name
description: [What the skill does] + [When to use it]
---

# Skill Name

[Brief introduction - what this skill does]

## Instructions

[Step-by-step instructions for Claude to follow]

## Examples

### Example 1: [Scenario Name]
**Input**: [User request]
**Output**: [Expected result]

### Example 2: [Another Scenario]
**Input**: [User request]
**Output**: [Expected result]

## Resources

[List of bundled files and their purposes, if applicable]
```

### Step 5: Follow Best Practices

Ensure the skill follows these principles:

#### Required Elements
- [ ] YAML frontmatter with `name` and `description`
- [ ] Clear, actionable instructions
- [ ] At least 2 examples showing usage
- [ ] Focused on single purpose

#### Description Quality
- [ ] Describes WHAT the skill does
- [ ] Describes WHEN to use it
- [ ] Specific and discoverable
- [ ] No vague terms like "helps with" without specifics

#### Content Quality
- [ ] Concise - only essential information
- [ ] No redundant general knowledge
- [ ] Uses progressive disclosure
- [ ] Examples show success criteria

#### Organization
- [ ] Logical section flow
- [ ] Clear headings
- [ ] Bundled files in appropriate directories
- [ ] Relative paths for resources

### Step 6: Create Supporting Files

If the skill needs templates, scripts, or resources:

1. **Templates**: Create in `templates/` directory
   - Use placeholder syntax: `{variable_name}`
   - Include comments explaining usage
   - Keep templates focused and reusable

2. **Scripts**: Create in `scripts/` directory
   - Add clear documentation
   - Include error handling
   - Make scripts self-contained

3. **Resources**: Create in `resources/` or `data/` directory
   - Reference data, guidelines, checklists
   - Keep files focused and organized

### Step 7: Generate Complete Skill

1. Create directory structure
2. Write SKILL.md
3. Create supporting files
4. Show summary of what was created

## Updating an Existing Skill

### Step 1: Read Current Skill

1. Read the existing SKILL.md file
2. Identify what needs updating
3. Check bundled resources

### Step 2: Analyze Against Best Practices

Check the skill for:
- [ ] YAML frontmatter present and complete
- [ ] Description is clear and discoverable
- [ ] Instructions are concise and actionable
- [ ] Examples are included
- [ ] Progressive disclosure is used
- [ ] No security issues (hardcoded secrets)
- [ ] Proper file organization

### Step 3: Propose Improvements

List specific improvements needed:
- Missing or weak description
- Instructions too verbose
- Missing examples
- Poor organization
- Security concerns
- Unclear purpose

### Step 4: Apply Updates

Update the skill with improvements:
1. Fix YAML frontmatter if needed
2. Improve description for better discovery
3. Refine instructions for clarity
4. Add or improve examples
5. Reorganize if needed
6. Add missing resources

### Step 5: Validate Updates

Confirm the updated skill:
- Follows all best practices
- Is more focused and clear
- Has better examples
- Uses progressive disclosure
- Is secure

## Templates Available

You can reference these templates from the `templates/` directory:

### 1. Basic Skill Template
`templates/basic-skill.md` - Minimal skill structure

### 2. Advanced Skill Template
`templates/advanced-skill.md` - Full-featured skill with resources

### 3. Code Processing Skill Template
`templates/code-skill.md` - For code-related skills

### 4. Content Generation Skill Template
`templates/content-skill.md` - For content creation skills

## Best Practices Checklist

Before finalizing any skill, verify:

### Structure
- [ ] Directory created: `.claude/skills/[skill-name]/`
- [ ] SKILL.md file exists
- [ ] Resources in appropriate subdirectories
- [ ] Naming convention: kebab-case

### YAML Frontmatter
- [ ] Has opening `---`
- [ ] Has `name:` field (kebab-case)
- [ ] Has `description:` field (clear and specific)
- [ ] Has closing `---`

### Description
- [ ] Explains what the skill does
- [ ] Explains when to use it
- [ ] Specific enough for discovery
- [ ] Not vague or generic

### Instructions
- [ ] Clear and actionable
- [ ] Step-by-step when appropriate
- [ ] Concise - no redundant info
- [ ] References bundled files when needed

### Examples
- [ ] At least 2 examples included
- [ ] Shows input and output
- [ ] Covers common use cases
- [ ] Demonstrates expected behavior

### Progressive Disclosure
- [ ] Most common info first
- [ ] Advanced info in later sections
- [ ] References to files, not full content
- [ ] Hierarchical organization

### Security
- [ ] No hardcoded secrets
- [ ] No sensitive information
- [ ] Scripts are safe
- [ ] Dependencies are documented

## Common Patterns

### Pattern 1: Simple Task Skill

**Use for**: Single-step transformations or formatting

```markdown
---
name: task-name
description: [Action] [target]. Use when user asks to [trigger].
---

# Task Name

Brief description of the task.

## Instructions
1. Take input
2. Process it
3. Return result

## Example
**Input**: [example]
**Output**: [result]
```

### Pattern 2: Multi-Step Workflow Skill

**Use for**: Complex processes with multiple stages

```markdown
---
name: workflow-name
description: [Process description]. Use when user wants to [goal].
---

# Workflow Name

## Process

### Step 1: [Stage Name]
[Instructions for this stage]

### Step 2: [Stage Name]
[Instructions for this stage]

### Step 3: [Stage Name]
[Instructions for this stage]

## Examples
[Examples covering the full workflow]
```

### Pattern 3: Resource-Heavy Skill

**Use for**: Skills with templates, scripts, or data files

```markdown
---
name: resource-skill
description: [What it does with resources]. Use when [trigger].
---

# Resource Skill

## Available Resources

- `templates/template1.md` - [Purpose]
- `scripts/helper.py` - [Purpose]
- `data/reference.csv` - [Purpose]

## Instructions

1. Identify need
2. Load appropriate resource from above
3. Use resource per instructions
4. Generate output

## Examples
[Examples showing resource usage]
```

### Pattern 4: Interactive Skill

**Use for**: Skills that need user input during execution

```markdown
---
name: interactive-skill
description: [Interactive process]. Use when user wants to [goal].
---

# Interactive Skill

## Workflow

### Phase 1: Gather Information
Ask user:
- Question 1
- Question 2
- Question 3

### Phase 2: Process
Based on answers, [process description]

### Phase 3: Refine
Present result and ask for:
- Feedback
- Adjustments
- Confirmation

## Examples
[Examples showing the interactive flow]
```

## Error Prevention

### Common Mistakes to Avoid

1. **Vague Description**
   - ❌ Bad: "Helps with code"
   - ✅ Good: "Review Python code for PEP 8 compliance. Use when user asks to review or lint Python code."

2. **Too Broad Scope**
   - ❌ Bad: One skill for all web development
   - ✅ Good: Separate skills for React, CSS, testing, deployment

3. **Missing Examples**
   - ❌ Bad: Only instructions, no examples
   - ✅ Good: At least 2 examples with input/output

4. **Embedding Large Content**
   - ❌ Bad: Full datasets in SKILL.md
   - ✅ Good: Reference files in `data/` directory

5. **Hardcoded Secrets**
   - ❌ Bad: API keys in SKILL.md or scripts
   - ✅ Good: Environment variables or user-provided

6. **No Progressive Disclosure**
   - ❌ Bad: All information at same level
   - ✅ Good: Common info first, advanced info later

## Output Format

### When Creating a New Skill

Provide:
1. **Summary**: Brief description of what was created
2. **File Structure**: Show directory tree
3. **Next Steps**: How to test and use the skill
4. **Customization Tips**: Suggestions for user to refine

Example:
```
Created skill: `python-linter`

File structure:
.claude/skills/python-linter/
├── SKILL.md
└── resources/
    └── pep8-checklist.md

The skill is ready to use! Test it by asking:
"Review this Python code for style issues"

Customization suggestions:
- Add more PEP 8 rules to resources/pep8-checklist.md
- Create templates for common fix patterns
- Add examples specific to your codebase
```

### When Updating a Skill

Provide:
1. **Analysis**: What issues were found
2. **Changes Made**: List of improvements
3. **Before/After**: Show key differences
4. **Validation**: Confirm best practices met

Example:
```
Updated skill: `code-reviewer`

Issues found:
- Description was too vague
- Missing examples
- Instructions were verbose

Changes made:
✓ Improved description for better discovery
✓ Added 3 concrete examples
✓ Condensed instructions by 40%
✓ Added progressive disclosure sections

The skill now follows all best practices.
```

## Testing Recommendations

After creating or updating a skill, suggest:

1. **Test with real scenarios**: Try the skill with actual use cases
2. **Check discovery**: Verify Claude finds the skill when appropriate
3. **Monitor usage**: Watch how the skill performs in practice
4. **Iterate**: Refine based on observations
5. **Get feedback**: Ask users if the skill meets their needs

## Remember

- **Keep it focused**: One skill, one purpose
- **Be concise**: Only essential information
- **Use examples**: Show what success looks like
- **Progressive disclosure**: Load information as needed
- **Iterate constantly**: Improve based on real usage
- **Think about discovery**: Make descriptions specific and clear
- **Security first**: Never hardcode secrets

Your goal is to create skills that are:
- **Discoverable**: Claude knows when to use them
- **Effective**: They accomplish their purpose well
- **Efficient**: They don't waste context window
- **Maintainable**: Easy to update and improve
- **Secure**: No security vulnerabilities
