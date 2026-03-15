<div align="center">

# Plo

**可视化提示词流水线构建工具** – 用流程图设计你的 LLM 工作流，无需编写代码。

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/xuhe2/plo)](https://goreportcard.com/report/github.com/xuhe2/plo)
[![Release](https://img.shields.io/github/release/xuhe2/plo.svg)](https://github.com/xuhe2/plo/releases/latest)

[English](./README.md) | 中文

</div>

## 🚀 Plo 是什么？

Plo 将可视化流程图转换为 LLM 可用的结构化提示词流水线。使用 Draw.io 图表设计复杂的 AI 工作流，只需一条命令即可导出为可执行格式。

- **可视化设计**：拖拽式流程图界面 – 无需编码
- **即时导出**：将图表转换为生产就绪的提示词流水线
- **可扩展**：支持接入自定义数据源和输出方案
- **开发者友好**：清晰的架构，易于扩展

## 💡 快速开始

### 安装

```bash
# 克隆仓库
git clone https://github.com/xuhe2/plo.git
cd plo

# 从源码构建
make build
```

### 使用

```bash
# 尝试示例
plo -input examples/example.drawio -output examples/output.md

# 将你自己的 Draw.io 流程图转换为 Markdown
plo -input your-workflow.drawio -output pipeline.md

# 或输出到标准输出
plo -input your-workflow.drawio
```

## 🎯 功能特性

### 核心功能

- **流程图解析**：将 Draw.io 图表解析为结构化流水线
- **校验机制**：自动验证流水线完整性和连接关系
- **Markdown 导出**：清晰、可读的提示词流水线格式
- **命令行界面**：简单易用的命令行工具，快速转换

### 可扩展性

Plo 采用可扩展架构，允许你添加：

- **自定义数据源**：不仅支持 Draw.io – 可添加 Mermaid、Lucidchart 等支持
- **自定义输出**：导出为 JSON、YAML、Python 代码或你自己的格式
- **流水线处理器**：添加验证、转换或优化层

## 📖 示例

### 输入：Draw.io 流程图

```
[开始] → [分析请求] → {有数据吗?}
                        ├─ 是 → [处理数据] → [生成响应]
                        └─ 否 → [询问信息] → [生成响应]
```

### 输出：结构化流水线

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

查看 [examples/example.drawio](./examples/example.drawio) 和 [examples/output.md](./examples/output.md) 获取完整示例。

## 🔌 扩展 Plo

### 添加新的数据源

实现 `Parser` 接口以支持新的输入格式：

```go
type Parser interface {
    Parse(data []byte) (*core.Pipeline, error)
}
```

### 添加新的输出格式

实现 `Exporter` 接口以支持自定义输出：

```go
type Exporter interface {
    Export(pipeline *core.Pipeline) ([]byte, error)
}
```

## 🛠️ 开发

```bash
# 运行测试
make test

# 构建项目
make build

# 运行示例
make run
```

## 🤝 贡献

欢迎贡献！无论是：

- 添加新的数据源适配器
- 创建新的导出格式
- 改进文档
- 报告 Bug

随时提交 Issue 或 PR！

## 📄 许可证

MIT License – 详见 [LICENSE](./LICENSE)。

---

<div align="center">
由社区用心打造
</div>
