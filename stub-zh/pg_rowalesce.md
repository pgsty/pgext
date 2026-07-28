## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_rowalesce/pg_rowalesce-0.1.12/README.md)
- [官方扩展控制文件 (pg_rowalesce.control)](https://api.pgxn.org/src/pg_rowalesce/pg_rowalesce-0.1.12/pg_rowalesce.control)
- [官方扩展 SQL (pg_rowalesce--0.1.0.sql)](https://api.pgxn.org/src/pg_rowalesce/pg_rowalesce-0.1.12/sql/pg_rowalesce--0.1.0.sql)

`pg_rowalesce` — PostgreSQL 扩展 pg_rowalesce 的核心功能是 rowalesce() 函数。当 SQL 需要这些特殊函数或聚合时，请使用它。在安装扩展之前，必须先安装并验证其依赖项。

### 核心工作流

```sql
CREATE EXTENSION pg_rowalesce;

select rowalesce('{"my_attr_1": 3, "my_attr_2": "b"}'::jsonb, null::my.type)
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `insert_row(inout anyelement)` 是一个扩展函数，返回 `anyelement`。
- `pg_rowalesce_meta_pgxn()` 是一个扩展函数，返回 `jsonb`。
- `pg_rowalesce_readme()` 是一个扩展函数，返回 `text`。
- `table_defaults(pg_class$ regclass, include_columns$ hstore = null)` 是一个扩展函数，返回 `hstore`。
- `test__pg_rowalesce` 是一个扩展过程。
- `myrow` 是一个扩展定义的类型。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1.12`。
- 请先安装并验证确认的扩展依赖项：`hstore`。
- 控制文件将扩展标记为可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
