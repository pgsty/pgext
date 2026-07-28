## 用法

来源：

- [官方上游 README](https://github.com/postgraphdb/postgraph/blob/72a080d3367aea6e0ffffaed54a9d0b025caee32/README.md)
- [官方扩展控制文件 (postgraph.control)](https://github.com/postgraphdb/postgraph/blob/72a080d3367aea6e0ffffaed54a9d0b025caee32/postgraph.control)
- [官方实现源代码](https://github.com/postgraphdb/postgraph/blob/72a080d3367aea6e0ffffaed54a9d0b025caee32/src/backend/postgraph.c)

`postgraph` — PostGraph 是一个基于 Postgres 的多模型、以图为中心的查询引擎。PostGraph 设计用于快速处理 OLTP、OLAP 和 AI 应用程序。当应用程序需要这种特定的数据库功能时，请使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION postgraph;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以匹配固定源代码。
