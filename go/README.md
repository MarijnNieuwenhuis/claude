# Go Resources

Comprehensive Go programming resources for writing idiomatic, maintainable, and performant Go code.

## Purpose

This directory contains curated Go best practices, patterns, and guidelines that can be referenced from Claude skills. These resources help ensure consistent, high-quality Go code across all projects.

## Contents

### Core Resources

1. **[Best Practices](best-practices.md)** - Essential Go coding standards and conventions
2. **[Idiomatic Go](idiomatic-go.md)** - Writing Go code the "Go way"
3. **[Design Patterns](design-patterns.md)** - Common design patterns in Go
4. **[Common Mistakes](common-mistakes.md)** - 100 Go mistakes and how to avoid them
5. **[Testing Practices](testing-practices.md)** - Testing strategies and patterns
6. **[Error Handling](error-handling.md)** - Proper error handling techniques
7. **[Concurrency Patterns](concurrency-patterns.md)** - Goroutines, channels, and synchronization

## Quick Reference

### Core Principles

1. **Simplicity** - Favor simple, clear code over clever solutions
2. **Readability** - Code is read more often than written
3. **Composition** - Prefer composition over inheritance
4. **Explicit** - Be explicit rather than implicit
5. **Pragmatic** - Focus on solving real problems

### Essential Tools

- `gofmt` - Automatic code formatting
- `go vet` - Static analysis for bugs
- `golint` - Linting for style issues
- `go test` - Built-in testing framework
- `go mod` - Dependency management

## Official Resources

- **Effective Go**: https://go.dev/doc/effective_go
- **Go Documentation**: https://go.dev/doc/
- **Go Blog**: https://go.dev/blog/
- **Go Spec**: https://go.dev/ref/spec

## Community Resources

- **100 Go Mistakes**: https://100go.co/
- **Go Patterns**: https://github.com/tmrts/go-patterns
- **Go by Example**: https://gobyexample.com/
- **Google Go Style Guide**: https://google.github.io/styleguide/go/

## How to Use These Resources

### In Skills

Reference specific sections in your SKILL.md files:

```markdown
## Go Best Practices

Follow the guidelines in `.claude/go/best-practices.md`, specifically:
- Naming conventions
- Error handling patterns
- Code organization
```

### For Code Reviews

Use these resources as a checklist:
- Check against common mistakes
- Verify idiomatic patterns are used
- Ensure proper error handling
- Validate testing coverage

### For Learning

Read these documents in order:
1. Start with **Idiomatic Go** to understand Go philosophy
2. Move to **Best Practices** for practical guidelines
3. Study **Design Patterns** for architectural patterns
4. Review **Common Mistakes** to avoid pitfalls
5. Master **Concurrency Patterns** for advanced topics

## Contributing

These resources are maintained as part of the Claude Code project structure. Updates should:
- Be based on official Go documentation
- Include practical examples
- Reference authoritative sources
- Follow the existing markdown structure

## Version Information

- **Go Version**: 1.21+ (adjust based on project requirements)
- **Last Updated**: 2025-10-20
- **Maintained By**: Development Team

---

For questions or suggestions, refer to the main project documentation.
