# Skill Creation Best Practices Checklist

## Pre-Creation Phase

### Scope Validation
- [ ] Skill has single, well-defined purpose
- [ ] Scope is not too broad (consider splitting if needed)
- [ ] Use cases are clear and specific
- [ ] Target users are identified

### Requirements Gathering
- [ ] Understand what the skill should do
- [ ] Know when the skill should be triggered
- [ ] Identify needed resources (templates, scripts, data)
- [ ] Clarify expected outputs

## Creation Phase

### Directory Structure
- [ ] Directory created at `.claude/skills/{skill-name}/`
- [ ] Directory name uses kebab-case
- [ ] SKILL.md file created
- [ ] Subdirectories created as needed (templates/, scripts/, data/, resources/)

### YAML Frontmatter
- [ ] Opening delimiter `---` present
- [ ] `name:` field present (kebab-case, matches directory name)
- [ ] `description:` field present and comprehensive
- [ ] Closing delimiter `---` present
- [ ] No syntax errors in YAML

### Description Quality
- [ ] Describes WHAT the skill does (clear action/capability)
- [ ] Describes WHEN to use it (triggers/scenarios)
- [ ] Specific enough for Claude to discover appropriately
- [ ] Not too vague (avoid "helps with X" without specifics)
- [ ] Keywords match likely user requests
- [ ] Length: 1-3 sentences, focused and clear

## Content Quality

### Instructions
- [ ] Clear and actionable
- [ ] Step-by-step format when appropriate
- [ ] Concise - only essential information
- [ ] No redundant general knowledge
- [ ] References bundled files when needed
- [ ] Uses progressive disclosure (common info first)
- [ ] Organized with clear headings

### Examples
- [ ] At least 2 examples provided
- [ ] Examples show input and expected output
- [ ] Examples cover common use cases
- [ ] Examples demonstrate success criteria
- [ ] Examples are realistic and practical
- [ ] Examples show different scenarios/variations

### Progressive Disclosure
- [ ] Information organized hierarchically
- [ ] Most common/important info in early sections
- [ ] Advanced features in later sections
- [ ] References to files instead of embedding content
- [ ] Large data in separate files, not in SKILL.md
- [ ] Context-efficient structure

## Resources and Files

### Template Files
- [ ] Stored in `templates/` directory
- [ ] Use clear placeholder syntax (e.g., `{variable_name}`)
- [ ] Include usage comments/instructions
- [ ] Focused and reusable
- [ ] Referenced in SKILL.md

### Scripts
- [ ] Stored in `scripts/` directory
- [ ] Include clear documentation/comments
- [ ] Have error handling
- [ ] Are self-contained
- [ ] Dependencies documented
- [ ] Executable permissions set if needed
- [ ] Referenced in SKILL.md

### Data Files
- [ ] Stored in `data/` or `resources/` directory
- [ ] Organized logically
- [ ] Referenced clearly in SKILL.md
- [ ] Used with relative paths
- [ ] Not duplicating information Claude already has

## Security

### Sensitive Information
- [ ] No hardcoded API keys
- [ ] No passwords or secrets
- [ ] No personal identifiable information (PII)
- [ ] Secrets use environment variables or user input
- [ ] Security requirements documented

### Scripts and Code
- [ ] Scripts are from trusted sources or self-authored
- [ ] Code dependencies documented
- [ ] No malicious code patterns
- [ ] Input validation present where needed
- [ ] Follows principle of least privilege

## Validation

### Functionality
- [ ] Skill accomplishes its stated purpose
- [ ] Instructions are complete and correct
- [ ] Examples match actual behavior
- [ ] Resources are accessible and correct
- [ ] No broken references or links

### Best Practices Compliance
- [ ] Follows official Anthropic guidelines
- [ ] Uses recommended patterns
- [ ] Efficient use of context window
- [ ] Appropriate complexity level
- [ ] Professional quality

### Testing
- [ ] Tested with realistic scenarios
- [ ] Works with example inputs from SKILL.md
- [ ] Claude discovers skill appropriately
- [ ] Skill loads and executes correctly
- [ ] Resources load properly

## Post-Creation

### Documentation
- [ ] SKILL.md is complete and clear
- [ ] Resources are documented
- [ ] Usage examples are sufficient
- [ ] Any special requirements noted

### Iteration Plan
- [ ] Monitoring plan for skill usage
- [ ] Feedback collection method
- [ ] Update/improvement process defined
- [ ] Version tracking if needed

## Common Issues to Avoid

### Description Problems
- [ ] ❌ Too vague: "Helps with code"
- [ ] ✅ Specific: "Review Python code for PEP 8 style compliance"

### Scope Problems
- [ ] ❌ Too broad: One skill for entire web stack
- [ ] ✅ Focused: Separate skills for React, CSS, testing

### Content Problems
- [ ] ❌ Missing examples
- [ ] ✅ At least 2 clear examples with I/O

### Organization Problems
- [ ] ❌ All information at same level
- [ ] ✅ Hierarchical with progressive disclosure

### Security Problems
- [ ] ❌ Hardcoded secrets in files
- [ ] ✅ Environment variables or user-provided

### Resource Problems
- [ ] ❌ Large datasets embedded in SKILL.md
- [ ] ✅ Data in separate files with references

## Quality Gates

Before considering the skill complete, ensure:

1. **Discoverable**: Claude will find it when appropriate
2. **Effective**: It accomplishes its purpose well
3. **Efficient**: Minimal context window usage
4. **Maintainable**: Easy to update and improve
5. **Secure**: No security vulnerabilities
6. **Professional**: High-quality, production-ready

## Sign-off Checklist

Final review before deployment:

- [ ] All sections of this checklist completed
- [ ] Skill tested successfully
- [ ] Documentation is complete
- [ ] No security issues identified
- [ ] Ready for production use
- [ ] User informed of skill capabilities and usage

---

## Quick Reference

**Minimum Required**:
- Directory: `.claude/skills/{name}/`
- File: `SKILL.md` with YAML frontmatter
- Fields: `name` and `description`
- Content: Instructions and examples

**Recommended**:
- At least 2 examples
- Progressive disclosure structure
- Bundled resources in subdirectories
- Clear, specific description
- Security review completed

**Best Practice**:
- Focused single purpose
- Context-efficient
- Well-documented
- Thoroughly tested
- Iteratively improved
