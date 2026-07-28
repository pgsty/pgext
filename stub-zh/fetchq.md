## 用法

来源：

- [官方上游 README](https://github.com/fetchq/pg-extension/blob/b9c3f62e226401c94635709fec32e48fd85f754a/README.md)
- [官方扩展控制文件 (fetchq.control)](https://github.com/fetchq/pg-extension/blob/b9c3f62e226401c94635709fec32e48fd85f754a/src/fetchq.control)

`fetchq` — Postgres 扩展，可启用 FetchQ 功能。当应用程序需要此特定数据库功能时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION fetchq;
```

在目标数据库中安装扩展，如果有可用示例，请运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `4.0.2`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
