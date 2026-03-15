# Plo - 提示词流水线解析工具

通过解析 Draw.io 流程图来构建提示词流水线。

## 目录结构

```
plo/
├── main.go
├── go.mod
├── internal/
│   ├── core/              # 核心逻辑层
│   ├── adapter/           # 适配层 (Draw.io)
│   └── exporter/          # 导出层 (Markdown)
└── example.drawio
```

## 使用

```bash
go run main.go -input example.drawio -output output.md
```

## 三层架构

1. **Adapter 层**: 将 Draw.io XML 转换为内部 Pipeline 结构
2. **Core 层**: 处理图逻辑和校验
3. **Exporter 层**: 将 Pipeline 导出为 Markdown
