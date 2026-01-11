# Contributing to Goove

Thank you for your interest in contributing to Goove! This document provides guidelines and instructions for contributing.

## Code of Conduct

By participating in this project, you agree to abide by our [Code of Conduct](CODE_OF_CONDUCT.md).

## How to Contribute

### Reporting Bugs

If you find a bug, please open an issue on GitHub with:

- A clear, descriptive title
- Steps to reproduce the issue
- Expected behavior vs actual behavior
- Your environment (OS, Go version, browser if applicable)
- Any relevant logs or screenshots

### Suggesting Features

Feature requests are welcome! Please open an issue with:

- A clear description of the feature
- The problem it solves or use case it addresses
- Any implementation ideas you have

### Pull Requests

1. **Fork the repository** and create your branch from `main`
2. **Set up the development environment** (see below)
3. **Make your changes** with clear, descriptive commits
4. **Test your changes** thoroughly
5. **Submit a pull request** with a clear description of the changes

## Development Setup

### Prerequisites

- Go 1.25 or later
- [Task](https://taskfile.dev/) (task runner)
- [Templ CLI](https://templ.guide/)

### Getting Started

1. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/goove.git
   cd goove
   ```

2. Install development dependencies:
   ```bash
   task install-dev
   ```

3. Copy the environment file:
   ```bash
   cp .env.example .env
   ```

4. Start the development server:
   ```bash
   task dev
   ```

### Project Structure

```
cmd/goove/       # Application entry point
internal/        # Private application code
  config/        # Configuration management
  database/      # Database setup and migrations
  handlers/      # HTTP request handlers
  http/          # Routing and middleware
  models/        # Data models
view/            # Templ templates
static/          # Static assets (CSS, JS)
```

### Development Commands

| Command | Description |
|---------|-------------|
| `task dev` | Start development server with hot reload |
| `task build` | Build the application |
| `task test` | Run tests |
| `task vet` | Run go fmt and go vet |
| `task clean` | Clean build artifacts |

## Coding Standards

### Go

- Follow standard Go conventions and idioms
- Run `task vet` before committing (runs both fmt and vet)
- Write tests for new functionality

### Templates (Templ)

- Keep templates focused and composable
- Use components for reusable UI elements
- Follow the existing naming conventions

### CSS (Tailwind)

- Use Tailwind utility classes
- Keep custom CSS minimal
- Maintain dark mode support

### Commits

- Use clear, descriptive commit messages
- Keep commits focused on a single change
- Reference issue numbers where applicable

## Questions?

If you have questions, feel free to open an issue or start a discussion on GitHub.

Thank you for contributing!
