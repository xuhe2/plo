<div align="center">

# Plo

**Visual Prompt Pipeline Builder** – Design your LLM workflows with flowcharts, not code.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/xuhe2/plo)](https://goreportcard.com/report/github.com/xuhe2/plo)
[![Release](https://img.shields.io/github/release/xuhe2/plo.svg)](https://github.com/xuhe2/plo/releases/latest)

English | [中文](./README_CN.md)

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
# Try the example
plo -input examples/example.drawio -output examples/output.md

# Convert your own Draw.io flowchart to Markdown
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
# 🛠 提示词流水线执行技能 (Skill: Plo-Pipeline)

## 第一部分：全局流程拓扑 (Workflow Topology)
> [PROTOCOL]: 此部分仅用于确定逻辑流转，严禁在此执行具体指令。

| 当前节点 ID | 节点摘要 | 跳转条件 (Condition) | 下一个节点 (Target) |
| :--- | :--- | :--- | :--- |
| **Node001** | 检查我的本地文件夹(仓库) | [IF] "DEFAULT" | ➡️ **Node002** |
| **Node002** | 是否有文件 | [IF] "否" <br> [IF] "是" | ➡️ **Node003** <br> ➡️ **Node004** |

## 第二部分：节点指令详情 (Node Specifications)
> [PROTOCOL]: 确定当前节点后，请严格执行以下 [INSTRUCTION] 内容。

### 📍 Node001

**[INSTRUCTION]**
"""
检查我的本地文件夹(仓库)
"""
```

Check out the [examples/example.drawio](./examples/example.drawio) and [examples/output.md](./examples/output.md) for a complete example.

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
