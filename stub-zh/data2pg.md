## 用法

来源：

- [官方上游 README](https://github.com/dalibo/data2pg/blob/b4a2bd14b926b82553ccc5cc4e18fe5f140977ca/README.md)
- [官方扩展控制文件 (data2pg.control)](https://github.com/dalibo/data2pg/blob/b4a2bd14b926b82553ccc5cc4e18fe5f140977ca/ext/data2pg.control)
- [官方扩展 SQL (data2pg--0.3.sql)](https://github.com/dalibo/data2pg/blob/b4a2bd14b926b82553ccc5cc4e18fe5f140977ca/ext/data2pg--0.3.sql)

`data2pg` — 迁移框架，用于通过外部数据包装器发现、复制和比较非 PostgreSQL 数据库内容。在从 PostgreSQL 移动、转换或集成相应数据时使用它。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION data2pg;
```

在目标数据库中安装扩展，当可用时运行上游最小示例，并在将其集成到应用程序 SQL 之前验证安装的版本和返回值。

### 重要对象

- `add_step_parent(p_batchName TEXT, p_step TEXT, p_parent_step TEXT)` 是一个扩展函数，返回 `INT`。
- `assign_fkey_checks_to_batch(p_batchName TEXT, p_schema TEXT, p_table TEXT, p_fkey TEXT DEFAULT NULL)` 是一个扩展函数，返回 `INT`。
- `assign_index_to_batch(p_batchName TEXT, p_schema TEXT, p_table TEXT, p_object TEXT)` 是一个扩展函数，返回 `INTEGER`。
- `assign_sequence_to_batch(p_batchName TEXT, p_schema TEXT, p_sequence TEXT)` 是一个扩展函数，返回 `INTEGER`。
- `assign_sequences_to_batch(p_batchName TEXT, p_schema TEXT, p_sequencesToInclude TEXT, p_sequencesToExclude TEXT)` 是一个扩展函数，返回 `INT`。
- `assign_table_checks_to_batch(p_batchName TEXT, p_schema TEXT, p_table TEXT)` 是一个扩展函数，返回 `INTEGER`。
- `assign_table_part_to_batch(p_batchName TEXT, p_schema TEXT, p_table TEXT, p_partId TEXT)` 是一个扩展函数，返回 `INTEGER`。
- `assign_table_part_to_batch(p_batchName TEXT, p_schema TEXT, p_table TEXT, p_partNum INTEGER)` 是一个扩展函数，返回 `INTEGER`。
- `assign_table_to_batch(p_batchName TEXT, p_schema TEXT, p_table TEXT)` 是一个扩展函数，返回 `INTEGER`。
- `assign_tables_checks_to_batch(p_batchName TEXT, p_schema TEXT, p_tablesToInclude TEXT, p_tablesToExclude TEXT)` 是一个扩展函数，返回 `INTEGER`。
- `assign_tables_to_batch(p_batchName TEXT, p_schema TEXT, p_tablesToInclude TEXT, p_tablesToExclude TEXT)` 是一个扩展函数，返回 `INTEGER`。
- `check_schema(p_schema TEXT)` 是一个扩展函数，返回 `void`。
- `complete_migration_configuration(p_migration TEXT)` 是一个扩展函数，返回 `INT`。
- `create_migration(p_migration TEXT, p_sourceDbms TEXT, p_extension TEXT, p_serverOptions TEXT, p_userMappingOptions TEXT, p_userHasPrivileges BOOLEAN DEFAULT false, p_importSchemaOptions TEXT DEFAULT NULL)` 是一个扩展函数，返回 `INTEGER`。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `0.7`。
- 先安装并验证确认的扩展依赖项：`dblink`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
