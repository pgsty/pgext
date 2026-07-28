## 用法

来源：

- [官方上游 README](https://github.com/obartunov/pathprobe/blob/ae53ef0c7d9d266bd8c97c23b35a3aa6ff034143/README.md)
- [官方扩展控制文件 (pathprobe.control)](https://github.com/obartunov/pathprobe/blob/ae53ef0c7d9d266bd8c97c23b35a3aa6ff034143/pathprobe.control)
- [官方扩展 SQL (pathprobe--1.1.sql)](https://github.com/obartunov/pathprobe/blob/ae53ef0c7d9d266bd8c97c23b35a3aa6ff034143/pathprobe--1.1.sql)

`pathprobe` — pathprobe 是一个用于 PostgreSQL 计划器工作的诊断扩展。它不仅显示 EXPLAIN 选择的最终计划，还展示了路径级别的决策过程：哪些路径被跳过、拒绝、接受、替换或存活下来。在收集或解释相应的 PostgreSQL 统计信息时使用它。请使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pathprobe;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `pathprobe(query text)` 是一个扩展函数，返回 `text`。
- `pathprobe_json(query text)` 是一个扩展函数，返回 `text`。
- `pathprobe_propose(query text, spec text)` 是一个扩展函数，返回 `text`。
- `pathprobe_propose_json(query text, spec text)` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
