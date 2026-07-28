## 用法

来源：

- [官方上游 README](https://github.com/pierreforstmann/pg_wsl/blob/5444e39f6cf9adc378107b4f70e1d5c1000ad8fc/README.md)
- [官方扩展控制文件 (pg_wsl.control)](https://github.com/pierreforstmann/pg_wsl/blob/5444e39f6cf9adc378107b4f70e1d5c1000ad8fc/pg_wsl.control)
- [官方扩展 SQL (pg_wsl--1.0.sql)](https://github.com/pierreforstmann/pg_wsl/blob/5444e39f6cf9adc378107b4f70e1d5c1000ad8fc/pg_wsl--1.0.sql)

`pg_wsl` — PostgreSQL 扩展，用于从主节点写入备用日志。使用此扩展进行相应的 SQL 或数据库实用程序工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_wsl;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小代码，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
