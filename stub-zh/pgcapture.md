## 用法

来源：

- [官方上游 README](https://github.com/replicase/pgcapture/blob/e9f4d88d4be1f12bddd72c86bbecfeddc9ea2c62/README.md)
- [官方扩展控制文件 (pgcapture.control)](https://github.com/replicase/pgcapture/blob/e9f4d88d4be1f12bddd72c86bbecfeddc9ea2c62/hack/postgres/extension/pgcapture.control)
- [官方扩展 SQL (pgcapture--0.1.sql)](https://github.com/replicase/pgcapture/blob/e9f4d88d4be1f12bddd72c86bbecfeddc9ea2c62/hack/postgres/extension/pgcapture--0.1.sql)

`pgcapture` — 一个可扩展的 Netflix DBLog 实现，适用于 PostgreSQL。在从 PostgreSQL 移动、转换或集成相应数据时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgcapture;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pgcapture.current_query()` 是一个扩展函数，返回 `TEXT`。
- `pgcapture.log_ddl()` 是一个扩展函数，返回 `event_trigger`。
- `pgcapture.sql_command_tags(p_sql TEXT)` 是一个扩展函数，返回 `TEXT[]`。
- `pgcapture.ddl_logs` 是一个由扩展安装或管理的表。
- `pgcapture.sources` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1`。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
