## 用法

来源：

- [官方上游 README](https://github.com/ai-db-uom/dbtune/blob/997916b3ed9aff19b15bd9a6a8379ddbfe52cdb0/README.md)
- [官方扩展控制文件 (dbtune_mab.control)](https://github.com/ai-db-uom/dbtune/blob/997916b3ed9aff19b15bd9a6a8379ddbfe52cdb0/dbtune_pg_mab_extension/dbtune_mab.control)
- [官方扩展 SQL (dbtune_mab--0.0.1.sql)](https://github.com/ai-db-uom/dbtune/blob/997916b3ed9aff19b15bd9a6a8379ddbfe52cdb0/dbtune_pg_mab_extension/dbtune_mab--0.0.1.sql)

`dbtune_mab` — DBTune MAB 顾问用于 PostgreSQL。当应用程序需要此特定数据库功能时，请使用它。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION dbtune_mab;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `dbtune_mab_tune(tablename TEXT, columns TEXT[])` 是一个扩展函数，返回 `TEXT`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.0.1`。
- 控制文件标记该扩展为不可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
