# Code Summarizer

![CI](https://github.com/Qyroxen/Code-Summarizer/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Code-Summarizer/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Code-Summarizer?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Code-Summarizer)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Code-Summarizer)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Code-Summarizer?style=social)](https://github.com/Qyroxen/Code-Summarizer/stargazers)

## What is it?

Code Summarizer is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Code-Summarizer.git
cd Code-Summarizer
go build -o codesummarizer .

# Run
./codesummarizer --help
```

## CLI Usage

```bash
# Basic usage
./codesummarizer

# With flags
./codesummarizer --verbose --output json

# Get help
./codesummarizer --help
```

## Examples

```bash
# Example 1
./codesummarizer example1

# Example 2
./codesummarizer example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o codesummarizer .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Code-Summarizer/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Code-Summarizer?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Code-Summarizer/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Code-Summarizer?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Code-Summarizer/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Code-Summarizer" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Code-Summarizer/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Code-Summarizer" alt="Pull Requests">
  </a>
</p>
