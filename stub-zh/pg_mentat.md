## 用法

来源：

- [官方上游 README](https://github.com/gburd/pg_mentat/blob/134015ebd3121e1a74eeff2de6a800143b33cb4e/pg_mentat/README.md)
- [官方扩展控制文件 (pg_mentat.control)](https://github.com/gburd/pg_mentat/blob/134015ebd3121e1a74eeff2de6a800143b33cb4e/pg_mentat/pg_mentat.control)
- [官方扩展 SQL (pg_mentat--1.0.0.sql)](https://github.com/gburd/pg_mentat/blob/134015ebd3121e1a74eeff2de6a800143b33cb4e/pg_mentat/sql/pg_mentat--1.0.0.sql)

`pg_mentat` — PostgreSQL 扩展，提供与 Datomic 兼容的 Datalog 查询引擎，并包含原生 EDN 数据类型。在移植或模拟相应数据库 API 时使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_mentat;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `mentat.allocate_entid(partition_name TEXT)` 是一个扩展函数，返回 `BIGINT`。
- `mentat.fulltext_update_trigger()` 是一个扩展函数，返回 `trigger`。
- `mentat.resolve_ident(keyword TEXT)` 是一个扩展函数，返回 `BIGINT`。
- `mentat.cardinality_type` 是一个扩展定义的类型。
- `mentat.EdnValue` 是一个扩展定义的类型。
- `mentat.unique_type` 是一个扩展定义的类型。
- `mentat.value_type` 是一个扩展定义的类型。
- `mentat.datoms` 是一个由扩展安装或管理的表。
- `mentat.datoms_bool` 是一个由扩展安装或管理的表。
- `mentat.datoms_bytes` 是一个由扩展安装或管理的表。
- `mentat.datoms_default` 是一个由扩展安装或管理的表。
- `mentat.datoms_double` 是一个由扩展安装或管理的表。
- `mentat.datoms_instant` 是一个由扩展安装或管理的表。
- `mentat.datoms_keyword` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 控制文件声明默认版本为 `1.5.7`。
- 控制文件标记该扩展为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况与固定源代码中的信息一致。
