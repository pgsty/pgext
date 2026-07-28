## 用法

来源：

- [官方上游 README](https://github.com/ohdsi/trex/blob/c9e2b35ff8ebaf2320b060c9d3d32f0007045511/README.md)
- [官方扩展控制文件 (pg_trex.control)](https://github.com/ohdsi/trex/blob/c9e2b35ff8ebaf2320b060c9d3d32f0007045511/plugins/pg_trex/pg_trex.control)
- [官方扩展 SQL (pg_trex--0.1.0.sql)](https://github.com/ohdsi/trex/blob/c9e2b35ff8ebaf2320b060c9d3d32f0007045511/plugins/pg_trex/sql/pg_trex--0.1.0.sql)

`pg_trex` — pg trex: PostgreSQL 扩展，用于分布式 trexsql。使用它来实现相应的分析或存储工作流。在目标 PostgreSQL 构建上使用上述链接的上游固定版本作为 API 边界并进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_trex;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.1.0`。
- 控制文件标记该扩展为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
