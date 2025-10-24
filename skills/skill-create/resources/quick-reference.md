# Skill Creation Quick Reference

## Minimum Skill Structure

```
.claude/skills/skill-name/
└── SKILL.md
```

## SKILL.md Template

```markdown
---
name: skill-name
description: What it does and when to use it
---

# Skill Name

Instructions for Claude

## Examples

**Input**: example
**Output**: result
```

## Required YAML Fields

| Field | Format | Example |
|-------|--------|---------|
| `name` | kebab-case | `python-linter` |
| `description` | [What] + [When] | `Review Python code for PEP 8. Use when user asks to lint Python.` |

## Description Formula

```
[Action/Capability] + [Context/Domain] + [Trigger/When to use]
```

**Examples**:
- "Generate API documentation from code. Use when user asks to document APIs."
- "Review Python code for PEP 8 compliance. Use when user asks to review Python code."
- "Create React components with TypeScript. Use when user wants to generate React components."

## Common Patterns

### Simple Task
```markdown
---
name: task-name
description: [Action]. Use when [trigger].
---

# Task Name

## Instructions
1. Step 1
2. Step 2
3. Step 3

## Examples
**Input**: example
**Output**: result
```

### Multi-Step Workflow
```markdown
---
name: workflow-name
description: [Process]. Use when [goal].
---

# Workflow Name

## Phase 1: [Name]
Instructions

## Phase 2: [Name]
Instructions

## Examples
Complete workflow example
```

### Resource-Based
```markdown
---
name: resource-skill
description: [What with resources]. Use when [trigger].
---

# Resource Skill

## Resources
- `templates/file.md` - Purpose
- `scripts/script.py` - Purpose

## Instructions
1. Load appropriate resource
2. Use per instructions
3. Generate output

## Examples
Example showing resource usage
```

## Directory Structure Options

### Basic
```
skill-name/
└── SKILL.md
```

### With Templates
```
skill-name/
├── SKILL.md
└── templates/
    └── template.md
```

### With Scripts
```
skill-name/
├── SKILL.md
└── scripts/
    └── helper.py
```

### Full Featured
```
skill-name/
├── SKILL.md
├── templates/
│   └── template.md
├── scripts/
│   └── script.py
├── data/
│   └── reference.csv
└── resources/
    └── guide.md
```

## File Naming Conventions

- Skill directory: `kebab-case`
- SKILL.md: Always `SKILL.md` (uppercase)
- Templates: `descriptive-name.extension`
- Scripts: `descriptive-name.extension`
- Resources: `descriptive-name.extension`

## Progressive Disclosure Levels

**Level 1**: YAML metadata (always loaded)
```yaml
---
name: skill-name
description: Brief description
---
```

**Level 2**: Main instructions (loaded when skill used)
```markdown
## Instructions
Common use case info here
```

**Level 3**: Advanced sections (loaded when needed)
```markdown
## Advanced Features
Detailed/rare use case info here
```

**Level 4**: Bundled files (read when referenced)
```markdown
For complete reference, see `resources/full-guide.md`
```

## Common Mistakes

| ❌ Mistake | ✅ Correct |
|-----------|----------|
| Vague description | Specific what + when |
| Too broad scope | Focused single purpose |
| No examples | At least 2 examples |
| Embedded large data | Data in separate files |
| Hardcoded secrets | Environment variables |
| All info at one level | Progressive disclosure |

## Quality Checklist

Quick validation:
- [ ] YAML frontmatter correct
- [ ] Description has what + when
- [ ] At least 2 examples
- [ ] Instructions are clear
- [ ] No hardcoded secrets
- [ ] Tested successfully

## Example Descriptions

### Good ✅
- "Review Python code for PEP 8 style. Use when user asks to lint or review Python code."
- "Generate React components with TypeScript and tests. Use when user wants to create React components."
- "Create API documentation from OpenAPI specs. Use when user asks to document APIs."

### Bad ❌
- "Helps with code" (too vague)
- "Does Python things" (unclear)
- "Useful tool" (no specifics)

## File Path References

Always use relative paths from SKILL.md:

```markdown
Read the template at `templates/example.md`
Use the script at `scripts/helper.py`
Reference data at `data/reference.csv`
```

## Testing Your Skill

1. **Create** skill in `.claude/skills/`
2. **Restart** Claude Code (if needed)
3. **Try** example requests from SKILL.md
4. **Verify** Claude discovers and uses skill
5. **Iterate** based on results

## Resources

- Full guide: `refs/creating-skills/README.md`
- Official docs: https://docs.claude.com/en/docs/agents-and-tools/agent-skills
- Examples: https://github.com/anthropics/skills

## Next Steps

1. Start with basic template
2. Add instructions and examples
3. Test with real scenarios
4. Add resources as needed
5. Iterate and improve
