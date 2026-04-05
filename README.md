# Go-AI-Agent

A complex, intelligent AI bot framework in Go.

## Features

- **ReAct Loop Reasoning**: The agent intelligently determines its next action.
- **Provider Abstraction**: Supporting various LLM providers with a common interface.
- **Memory Management**: Keeps context and tracks conversational states.
- **Plugin-System**: Easily register new tools (e.g., calculator, system functions).

## Getting Started

1.  **Project Setup**: Run `go mod tidy` (if Go is installed).
2.  **Execute the Bot**: Run `go run main.go`.

## Granular Commit History (How we did it)

Instead of a single `git add .`, we built this project in three distinct, logical commits:

### Commit 1: Framework and Types
We established the core interfaces and message types:
```powershell
git add go.mod pkg/ai/*.go
git commit -m "feat: Define core interfaces and framework"
```

### Commit 2: Logic and Intelligence
We added the Agent's reasoning engine and the memory system:
```powershell
git add internal/memory/history.go internal/agent/reasoner.go
git commit -m "feat: Implement Agent Reasoner and Memory history"
```

### Commit 3: Tools and Entry Point
Finally, we implemented specialized tools, the main execution program, and full documentation:
```powershell
git add internal/tools/calc.go main.go docs/ README.md
git commit -m "feat: Implement Calculator tool, Main execution loop, and Documentation"
```
