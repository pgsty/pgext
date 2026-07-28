## 用法

来源：

- [官方上游 README](https://github.com/kou/pg-copy-parquet/blob/7da367ea81d8964f5045fe0b1514a798d4ecbbc7/README.md)
- [官方扩展控制文件 (copy_parquet.control)](https://github.com/kou/pg-copy-parquet/blob/7da367ea81d8964f5045fe0b1514a798d4ecbbc7/copy_parquet.control)
- [官方扩展 SQL (copy_parquet--0.0.1.sql)](https://github.com/kou/pg-copy-parquet/blob/7da367ea81d8964f5045fe0b1514a798d4ecbbc7/copy_parquet--0.0.1.sql)

`copy_parquet` — PoC 以支持 PostgreSQL COPY 与 Apache Parquet。在移动、转换或集成相应数据到 PostgreSQL 时使用。上游将其描述为一个概念验证。

### 核心工作流

```sql
CREATE EXTENSION copy_parquet;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `copy_to_parquet(pg_catalog.regclass)` 是一个扩展函数，返回 `bytea`。
- `format_parquet(bytea)` 是一个扩展函数，返回 `text`。
- `scan_to_parquet(pg_catalog.regclass)` 是一个扩展函数，返回 `bytea`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.0.1`。
- 控制文件标记该扩展为可重定位。
- 上游将该项目描述为一个概念验证。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
