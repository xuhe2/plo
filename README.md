<div align="center">

# Plo

**Visual Prompt Pipeline Builder** – Design your LLM workflows with flowcharts, not code.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/xuhe2/plo)](https://goreportcard.com/report/github.com/xuhe2/plo)
[![Release](https://img.shields.io/github/release/xuhe2/plo.svg)](https://github.com/xuhe2/plo/releases/latest)

</div>

## 🚀 What is Plo?

Plo transforms visual flowcharts into structured prompt pipelines for LLMs. Design complex AI workflows using Draw.io diagrams and export them to executable formats with a single command.

- **Visual Design**: Drag-and-drop flowchart interface – no coding required
- **Instant Export**: Convert diagrams to production-ready prompt pipelines
- **Extensible**: Plug in custom data sources and output formats
- **Developer Friendly**: Clean architecture that's easy to extend

## 💡 Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/xuhe2/plo.git
cd plo

# Build from source
make build
```

### Usage

```bash
# Convert a Draw.io flowchart to Markdown
plo -input your-workflow.drawio -output pipeline.md

# Or output to stdout
plo -input your-workflow.drawio
```

## 🎯 Features

### Core Functionality

- **Flowchart Parsing**: Parse Draw.io diagrams into structured pipelines
- **Validation**: Automatic validation of pipeline integrity and connections
- **Markdown Export**: Clean, human-readable prompt pipeline format
- **CLI Interface**: Simple command-line tool for quick conversions

### Extensibility

Plo is built with an extensible architecture that allows you to add:

- **Custom Data Sources**: Not just Draw.io – add support for Mermaid, Lucidchart, etc.
- **Custom Outputs**: Export to JSON, YAML, Python code, or your own format
- **Pipeline Processors**: Add validation, transformation, or optimization layers

## 📖 Examples

### Input: Draw.io Flowchart

```
[Start] → [Analyze Request] → {Has Data?}
                               ├─ Yes → [Process Data] → [Generate Response]
                               └─ No  → [Ask for Info] → [Generate Response]
```

### Output: Structured Pipeline

```markdown
# Prompt Pipeline

## Workflow Topology

| Node ID | Name          | Condition | Next Node |
|---------|---------------|-----------|-----------|
| Node001 | Analyze Request | - | Node002 |
| Node002 | Has Data?     | Yes       | Node003   |
| Node002 | Has Data?     | No        | Node004   |
```

Check out the [example.drawio](./example.drawio) and [output.md](./output.md) for a complete example.

## 🔌 Extending Plo

### Adding a New Data Source

Implement the `Parser` interface to support new input formats:

```go
type Parser interface {
    Parse(data []byte) (*core.Pipeline, error)
}
```

### Adding a New Output Format

Implement the `Exporter` interface for custom outputs:

```go
type Exporter interface {
    Export(pipeline *core.Pipeline) ([]byte, error)
}
```

## 🛠️ Development

```bash
# Run tests
make test

# Build the project
make build

# Run with example
make run
```

## 🤝 Contributing

Contributions are welcome! Whether it's:

- Adding new data source adapters
- Creating new export formats
- Improving documentation
- Reporting bugs

Feel free to open an issue or submit a PR!

## 📄 License

MIT License – see [LICENSE](./LICENSE) for details.

---

<div align="center">
Made with by the community
</div>
