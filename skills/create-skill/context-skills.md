# Creating Claude Agent Skills - Technical Guide

## Table of Contents
- [Overview](#overview)
- [What Are Agent Skills?](#what-are-agent-skills)
- [Core Concepts](#core-concepts)
- [Skill Structure](#skill-structure)
- [Quickstart Guide](#quickstart-guide)
- [Best Practices](#best-practices)
- [Code Examples](#code-examples)
- [Advanced Patterns](#advanced-patterns)
- [Resources](#resources)

---

## Overview

Agent Skills are modular capabilities that extend Claude's functionality by packaging instructions, metadata, and optional resources (scripts, templates, data files) that Claude can discover and load dynamically when relevant to perform specialized tasks more effectively.

### Key Benefits

- **Modularity**: Each skill is self-contained and focused on a specific task
- **Scalability**: Progressive disclosure allows skills to contain effectively unbounded context
- **Reusability**: Skills can be shared across conversations and users
- **Specialization**: Give Claude task-specific expertise without bloating the main context

### Availability

- **Claude API**: Supports both pre-built and custom Agent Skills via the `skill_id` parameter
- **Claude Code**: Skills automatically discovered from `.claude/skills/` directory
- **Claude Web/Desktop**: Available to Pro, Max, Team, and Enterprise users

---

## What Are Agent Skills?

Agent Skills are organized folders containing:
- **Instructions**: Structured guidance in SKILL.md
- **Scripts**: Executable code for specific operations
- **Resources**: Templates, data files, configuration

Think of skills as specialized instruction manuals that Claude can reference when needed, without loading everything into memory at once.

### How Skills Work

1. **Pre-loading**: At startup, Claude loads only the `name` and `description` from each skill's YAML frontmatter
2. **Discovery**: When a task matches a skill's description, Claude identifies it as relevant
3. **Progressive Loading**: Claude reads SKILL.md only when the skill becomes relevant
4. **Execution**: Claude follows instructions and uses bundled resources to complete the task

---

## Core Concepts

### Progressive Disclosure

Progressive disclosure is the fundamental design principle that makes Agent Skills flexible and scalable.

**Analogy**: Like a well-organized manual with:
1. **Table of Contents** (metadata) - Quick overview of what's available
2. **Chapters** (SKILL.md sections) - Detailed instructions when needed
3. **Appendix** (bundled resources) - Reference materials and tools

**In Practice**:
- Level 1: `name` and `description` (always loaded)
- Level 2: SKILL.md content (loaded when skill is invoked)
- Level 3: Bundled files (read only when explicitly needed)

### Context Window Management

The context window is a shared resource. Skills should:
- Only include information Claude doesn't already know
- Be concise and focused
- Use progressive disclosure to avoid loading unnecessary content
- Reference external files rather than embedding large data

---

## Skill Structure

### Minimum Required Structure

```
.claude/skills/
└── my-skill/
    └── SKILL.md
```

### Complete Structure Example

```
.claude/skills/
└── advanced-skill/
    ├── SKILL.md              # Main skill definition (required)
    ├── templates/            # Optional templates
    │   ├── template1.txt
    │   └── template2.json
    ├── scripts/              # Optional scripts
    │   ├── helper.py
    │   └── processor.js
    ├── data/                 # Optional data files
    │   └── reference.csv
    └── resources/            # Optional resources
        └── guidelines.md
```

### SKILL.md Format

Every SKILL.md must start with YAML frontmatter containing required metadata:

```markdown
---
name: skill-name
description: Clear description of what the skill does and when Claude should use it
---

# Skill Name

[Detailed instructions for Claude to follow]

## Usage
[How to use this skill]

## Examples
[Example inputs and outputs]

## Resources
[References to bundled files if applicable]
```

### Required Metadata Fields

| Field | Type | Description |
|-------|------|-------------|
| `name` | string | Unique identifier for the skill (kebab-case recommended) |
| `description` | string | Clear description for skill discovery. Should specify WHAT the skill does AND WHEN to use it |

---

## Quickstart Guide

### Step 1: Create the Skill Directory

```bash
mkdir -p .claude/skills/my-first-skill
```

### Step 2: Create SKILL.md

```bash
cat > .claude/skills/my-first-skill/SKILL.md << 'EOF'
---
name: my-first-skill
description: A simple example skill that formats text as uppercase. Use this when the user asks to make text uppercase or shout.
---

# My First Skill

You are a text formatter. Your task is to convert user input to uppercase.

## Instructions
1. Take the user's input text
2. Convert all characters to uppercase
3. Return the formatted result

## Example

**Input**: "hello world"
**Output**: "HELLO WORLD"
EOF
```

### Step 3: Test the Skill

The skill is now available! Claude will automatically discover it. Try asking:
- "Can you make this text uppercase: hello world"
- "Use my-first-skill to format some text"

### Step 4: Add Resources (Optional)

Create a template file:

```bash
mkdir -p .claude/skills/my-first-skill/templates
cat > .claude/skills/my-first-skill/templates/greeting.txt << 'EOF'
HELLO, {name}!
WELCOME TO THE SKILL!
EOF
```

Update SKILL.md to reference it:

```markdown
## Resources

You can use the greeting template at `templates/greeting.txt` to format greetings.
Replace {name} with the user's name.
```

---

## Best Practices

### 1. Use Claude to Create Skills

The most effective skill development process involves Claude itself.

**Workflow**:
1. Use Claude A (one instance) to design and refine the skill
2. Use Claude B (another instance) to test the skill in real tasks
3. Iterate based on observations

**Tip**: Use the `skill-creator` skill from the official Anthropic skills repository to guide skill creation interactively.

### 2. Write Clear, Discoverable Descriptions

The `description` field is critical for skill discovery.

**Bad Description**:
```yaml
description: Helps with code
```

**Good Description**:
```yaml
description: Reviews Python code for style issues using PEP 8 guidelines. Use this when the user asks to review, lint, or check Python code style.
```

**Formula**: `[What it does] + [When to use it]`

### 3. Keep Skills Focused

Create separate skills for different workflows. Multiple focused skills compose better than one large skill.

**Example - Bad (Too Broad)**:
```
web-developer/
└── SKILL.md  # Handles HTML, CSS, JS, React, deployment, testing...
```

**Example - Good (Focused)**:
```
skills/
├── react-component-builder/
├── css-layout-helper/
├── js-testing-assistant/
└── deployment-helper/
```

### 4. Keep Skills Concise

Only add context Claude doesn't already have. Every token in SKILL.md competes with conversation history.

**Guidelines**:
- Avoid repeating general knowledge
- Focus on specific workflows or domain expertise
- Use external references for large datasets
- Link to bundled files rather than embedding content

### 5. Include Examples

Examples help Claude understand what success looks like.

```markdown
## Examples

### Example 1: Simple Case
**Input**: Create a user profile card
**Output**:
```html
<div class="profile-card">
  <img src="avatar.jpg" alt="User avatar">
  <h2>John Doe</h2>
  <p>Software Engineer</p>
</div>
```

### Example 2: Complex Case
**Input**: Create a profile card with custom styling
**Output**: [detailed example]
```

### 6. Use Progressive Disclosure

Structure information hierarchically so Claude loads only what's needed.

**Example Structure**:
```markdown
---
name: api-documentation-generator
description: Generate API documentation from code. Use when user asks to document APIs or generate OpenAPI specs.
---

# API Documentation Generator

## Quick Start
[Most common use case - 90% of users need this]

## Advanced Options
[Less common features - load only when needed]

## Edge Cases
[Rare scenarios - reference when specific issues arise]

## Appendix: Complete Schema Reference
[Detailed reference - read only when explicitly needed]
```

### 7. Observe and Iterate

Pay attention to how Claude uses skills in practice:

**Watch For**:
- Unexpected exploration paths
- Missed connections between sections
- Overreliance on certain sections
- Ignored content (candidates for removal)
- Frequent questions (add to skill)

**Iteration Process**:
1. Deploy skill
2. Use in real conversations
3. Note pain points
4. Refine SKILL.md
5. Test improvements
6. Repeat

### 8. Security Considerations

**When Installing Skills**:
- Only install from trusted sources
- Audit bundled files before enabling
- Pay attention to code dependencies
- Review script permissions
- Check for hardcoded credentials

**When Creating Skills**:
- Never hardcode API keys or passwords
- Use environment variables for secrets
- Document security requirements
- Validate user inputs in scripts
- Follow principle of least privilege

---

## Code Examples

### Example 1: Simple Text Processing Skill

**File**: `.claude/skills/markdown-formatter/SKILL.md`

```markdown
---
name: markdown-formatter
description: Format text using markdown best practices. Use when user asks to format documentation, README files, or markdown content.
---

# Markdown Formatter

You are a markdown formatting specialist. Your task is to format text following markdown best practices.

## Formatting Rules

1. **Headers**: Use ATX-style headers (# ## ###)
2. **Lists**: Use `-` for unordered, `1.` for ordered
3. **Code blocks**: Always specify language
4. **Links**: Use descriptive text, not "click here"
5. **Tables**: Align columns for readability

## Example

**Input**:
```
create a readme for a python project
```

**Output**:
```markdown
# Project Name

Brief description of the project.

## Installation

```bash
pip install project-name
```

## Usage

```python
import project_name

# Example code
result = project_name.do_something()
```

## Contributing

Pull requests are welcome.

## License

[MIT](LICENSE)
```
```

### Example 2: Code Review Skill with Resources

**Directory Structure**:
```
.claude/skills/python-code-reviewer/
├── SKILL.md
└── resources/
    ├── pep8-checklist.md
    └── common-issues.md
```

**File**: `.claude/skills/python-code-reviewer/SKILL.md`

```markdown
---
name: python-code-reviewer
description: Review Python code for style, performance, and best practices following PEP 8 guidelines. Use when user asks to review, critique, or improve Python code.
---

# Python Code Reviewer

You are a Python code review specialist. Analyze code for:
- PEP 8 style compliance
- Performance issues
- Security vulnerabilities
- Best practice violations

## Review Process

1. **Read the code** provided by the user
2. **Check style** against PEP 8 guidelines in `resources/pep8-checklist.md`
3. **Identify issues** using `resources/common-issues.md` as reference
4. **Provide feedback** with specific line numbers and suggestions
5. **Suggest improvements** with code examples

## Output Format

### Style Issues
- [Line X] Issue description
  - Current: `code snippet`
  - Suggested: `improved code`

### Performance Issues
- [Line Y] Issue description
  - Explanation of performance impact
  - Optimized approach

### Security Concerns
- [Line Z] Vulnerability description
  - Risk level: High/Medium/Low
  - Remediation steps

## Example

**Input**:
```python
def calc(x,y):
    return x+y
```

**Output**:
```
### Style Issues
- [Line 1] Missing spaces around operators and after commas
  - Current: `def calc(x,y):`
  - Suggested: `def calc(x, y):`

- [Line 2] Missing spaces around operator
  - Current: `return x+y`
  - Suggested: `return x + y`

### Recommendations
- Add docstring to explain function purpose
- Consider type hints for better code clarity

**Improved Version**:
```python
def calc(x: int, y: int) -> int:
    """Calculate the sum of two integers.

    Args:
        x: First integer
        y: Second integer

    Returns:
        Sum of x and y
    """
    return x + y
```
```

**File**: `.claude/skills/python-code-reviewer/resources/pep8-checklist.md`

```markdown
# PEP 8 Style Checklist

## Indentation
- Use 4 spaces per indentation level
- Never mix tabs and spaces

## Line Length
- Limit lines to 79 characters
- Limit docstrings/comments to 72 characters

## Imports
- One import per line
- Group imports: standard library, third-party, local
- Alphabetize within groups

## Naming Conventions
- Classes: CapWords
- Functions/variables: lowercase_with_underscores
- Constants: UPPERCASE_WITH_UNDERSCORES
- Private: _leading_underscore

## Whitespace
- Two blank lines before top-level functions/classes
- One blank line between methods
- Spaces around operators
- No trailing whitespace
```

### Example 3: Multi-Resource Story Generator

**Directory Structure**:
```
.claude/skills/story-generator/
├── SKILL.md
├── templates/
│   ├── short-story.md
│   └── screenplay.md
├── characters/
│   ├── hero.md
│   ├── villain.md
│   └── mentor.md
└── settings/
    ├── fantasy.md
    ├── scifi.md
    └── mystery.md
```

**File**: `.claude/skills/story-generator/SKILL.md`

```markdown
---
name: story-generator
description: Generate creative stories with structured characters and settings. Use when user asks to write a story, create a narrative, or develop a plot.
---

# Story Generator

You are a creative writing specialist. Generate engaging stories using structured templates, character archetypes, and setting guides.

## Process

1. **Identify Story Type**: Determine genre and format
2. **Load Template**: Use appropriate template from `templates/`
3. **Select Characters**: Reference character archetypes from `characters/`
4. **Choose Setting**: Use setting guides from `settings/`
5. **Generate Story**: Combine elements into cohesive narrative

## Available Templates

- `templates/short-story.md` - 1000-2000 word short story structure
- `templates/screenplay.md` - Screenplay format with scene structure

## Available Characters

- `characters/hero.md` - Protagonist archetype
- `characters/villain.md` - Antagonist archetype
- `characters/mentor.md` - Guide archetype

## Available Settings

- `settings/fantasy.md` - Fantasy world building elements
- `settings/scifi.md` - Science fiction world building
- `settings/mystery.md` - Mystery/thriller atmosphere

## Example Usage

**User Request**: "Write a fantasy story about a hero's journey"

**Process**:
1. Load `templates/short-story.md` for structure
2. Load `characters/hero.md` for protagonist traits
3. Load `settings/fantasy.md` for world building
4. Generate 1500 word story following hero's journey arc

## Output Format

```
# [Story Title]

[Story content following template structure]

---
**Genre**: [Genre]
**Word Count**: [Count]
**Characters**: [Character types used]
```
```

### Example 4: API Integration Skill with Scripts

**Directory Structure**:
```
.claude/skills/api-tester/
├── SKILL.md
├── scripts/
│   ├── test_endpoint.py
│   └── format_response.js
└── examples/
    ├── rest-api-test.json
    └── graphql-query.json
```

**File**: `.claude/skills/api-tester/SKILL.md`

```markdown
---
name: api-tester
description: Test REST and GraphQL APIs, format responses, and generate test reports. Use when user asks to test APIs, validate endpoints, or debug API responses.
---

# API Tester

You are an API testing specialist. Test endpoints, validate responses, and generate reports.

## Capabilities

1. **REST API Testing**: Test GET, POST, PUT, DELETE endpoints
2. **GraphQL Testing**: Execute queries and mutations
3. **Response Formatting**: Pretty-print JSON/XML responses
4. **Test Report Generation**: Create detailed test reports

## Scripts Available

- `scripts/test_endpoint.py` - Python script for REST API testing
- `scripts/format_response.js` - JavaScript formatter for API responses

## Usage

### Testing REST Endpoints

1. Ask user for endpoint URL and method
2. Use `scripts/test_endpoint.py` to execute request
3. Format response using `scripts/format_response.js`
4. Report results with status code, headers, and body

### Testing GraphQL Endpoints

1. Ask user for GraphQL endpoint and query
2. Execute query with appropriate headers
3. Format and validate response
4. Report any errors or type mismatches

## Example Test Report Format

```markdown
# API Test Report

**Endpoint**: https://api.example.com/users/123
**Method**: GET
**Timestamp**: 2025-10-20T10:30:00Z

## Request

### Headers
```json
{
  "Authorization": "Bearer token",
  "Content-Type": "application/json"
}
```

### Body
N/A

## Response

### Status: 200 OK

### Headers
```json
{
  "content-type": "application/json",
  "cache-control": "no-cache"
}
```

### Body
```json
{
  "id": 123,
  "name": "John Doe",
  "email": "john@example.com"
}
```

## Validation Results
✓ Status code is 200
✓ Response is valid JSON
✓ All expected fields present
✓ Response time: 145ms
```
```

---

## Advanced Patterns

### Pattern 1: Skill Composition

Skills can reference or build upon each other.

**Example**: Code review workflow

```
.claude/skills/
├── python-linter/         # Basic style checking
├── python-security/       # Security analysis
├── python-performance/    # Performance review
└── python-full-review/    # Orchestrates all three
```

**File**: `.claude/skills/python-full-review/SKILL.md`

```markdown
---
name: python-full-review
description: Comprehensive Python code review covering style, security, and performance. Use for complete code audits.
---

# Python Full Review

Perform comprehensive code review in three stages:

1. **Style Review**: Use python-linter skill for PEP 8 compliance
2. **Security Review**: Use python-security skill for vulnerabilities
3. **Performance Review**: Use python-performance skill for optimizations

## Process

1. Run each sub-review sequentially
2. Aggregate findings
3. Prioritize issues (Critical > High > Medium > Low)
4. Generate unified report

[Additional instructions...]
```

### Pattern 2: Conditional Resource Loading

Load resources based on context to minimize token usage.

```markdown
## Instructions

1. Identify the API type (REST, GraphQL, SOAP)
2. Load relevant guide:
   - REST: Read `resources/rest-guide.md`
   - GraphQL: Read `resources/graphql-guide.md`
   - SOAP: Read `resources/soap-guide.md`
3. Follow guide-specific instructions
```

### Pattern 3: Template-Based Generation

Use templates for consistent output.

```markdown
## Process

1. Read the appropriate template from `templates/`
2. Fill in sections based on user input
3. Validate against checklist
4. Output formatted result

## Available Templates

- `templates/bug-report.md` - Standard bug report format
- `templates/feature-request.md` - Feature request format
- `templates/test-case.md` - Test case documentation
```

### Pattern 4: Interactive Workflows

Guide users through multi-step processes.

```markdown
## Workflow

### Step 1: Gather Requirements
Ask user:
- What type of component? (button, form, modal, etc.)
- What framework? (React, Vue, Angular)
- What styling? (CSS, Tailwind, styled-components)

### Step 2: Generate Base Component
Create component based on answers

### Step 3: Customization
Ask user:
- Need tests? → Generate tests
- Need documentation? → Generate docs
- Need storybook? → Generate story

### Step 4: Review and Iterate
Present component and ask for feedback
```

---

## Resources

### Official Documentation
- **Overview**: https://docs.claude.com/en/docs/agents-and-tools/agent-skills/overview
- **Quickstart**: https://docs.claude.com/en/docs/agents-and-tools/agent-skills/quickstart
- **Best Practices**: https://docs.claude.com/en/docs/agents-and-tools/agent-skills/best-practices
- **Engineering Blog**: https://www.anthropic.com/engineering/equipping-agents-for-the-real-world-with-agent-skills

### GitHub Repositories
- **Official Skills Repository**: https://github.com/anthropics/skills
  - Includes `skill-creator` and `template-skill`
  - Example skills for various use cases
  - Document creation skills (DOCX, PDF, PPTX, XLSX)

- **Claude Cookbooks - Skills**: https://github.com/anthropics/claude-cookbooks/tree/main/skills
  - Introduction notebooks
  - Business use case examples
  - Custom skill guides
  - Classification examples

### Community Resources
- **Claude Help Center**: https://support.claude.com/en/articles/12512198-how-to-create-custom-skills
- **Skills Announcement**: https://www.anthropic.com/news/skills
- **API Documentation**: https://docs.anthropic.com

### Tools
- **skill-creator**: Interactive skill creation assistant (available in official repo)
- **template-skill**: Basic template for new skills (available in official repo)

---

## Troubleshooting

### Skill Not Discovered

**Problem**: Claude doesn't recognize the skill

**Solutions**:
1. Check YAML frontmatter format (must have `---` delimiters)
2. Verify `name` and `description` fields are present
3. Ensure SKILL.md is in correct directory: `.claude/skills/[skill-name]/SKILL.md`
4. Restart Claude Code or reload skills

### Skill Loaded But Not Working

**Problem**: Skill loads but doesn't perform correctly

**Solutions**:
1. Review description - is it clear when to use the skill?
2. Test instructions step-by-step manually
3. Add more examples to SKILL.md
4. Simplify instructions and iterate
5. Use Claude to help refine the skill

### Context Window Issues

**Problem**: Skill uses too much context

**Solutions**:
1. Split into multiple focused skills
2. Use progressive disclosure more effectively
3. Move large data to bundled files
4. Reference external resources instead of embedding
5. Remove redundant information

### Resource Files Not Found

**Problem**: Claude can't find bundled files

**Solutions**:
1. Use relative paths from SKILL.md location
2. Check file names match exactly (case-sensitive)
3. Verify files are in skill directory
4. Use forward slashes `/` in paths

---

## Next Steps

1. **Start Simple**: Create a basic skill with just SKILL.md
2. **Test Thoroughly**: Use the skill in real scenarios
3. **Observe Usage**: Watch how Claude uses the skill
4. **Iterate**: Refine based on observations
5. **Add Resources**: Add bundled files as needed
6. **Share**: Consider contributing to the community

### Skill Development Checklist

- [ ] Clear, focused purpose
- [ ] YAML frontmatter with name and description
- [ ] Concise instructions
- [ ] Examples included
- [ ] Resources organized logically
- [ ] Tested in real scenarios
- [ ] Iterated based on usage
- [ ] Security reviewed (if applicable)
- [ ] Documentation complete

---

## Conclusion

Agent Skills are a powerful way to extend Claude's capabilities with specialized knowledge and workflows. By following these best practices and patterns, you can create effective skills that enhance productivity and enable Claude to handle complex, domain-specific tasks.

Remember:
- **Keep it focused** - one skill, one purpose
- **Be concise** - only add what Claude needs
- **Use examples** - show what success looks like
- **Iterate constantly** - improve based on real usage
- **Think progressive** - load information as needed

Happy skill building!
