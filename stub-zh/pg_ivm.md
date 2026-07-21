## 用法

来源：

- [官方v1.15 README](https://github.com/sraoss/pg_ivm/blob/v1.15/README.md)
- [v1.15 发行说明](https://github.com/sraoss/pg_ivm/releases/tag/v1.15)
- [从v1.14到v1.15的升级SQL](https://github.com/sraoss/pg_ivm/blob/v1.15/pg_ivm--1.14--1.15.sql)
- [pg_ivm_dump_metadata 工具](https://github.com/sraoss/pg_ivm/blob/v1.15/scripts/pg_ivm_dump_metadata)

`pg_ivm` 为 PostgreSQL 提供了即时增量视图维护功能。增量可维护材料化视图（IMMV）以表的形式存储在 `pgivm` 模式中，带有触发器和元数据；基表的变化会在同一个事务中更新 IMMV 而不是重新计算整个查询。

### 启用并创建一个 IMMV

为可以修改 IMMV 基表的每个会话加载库。集群级设置需要重启：

```conf
shared_preload_libraries = 'pg_ivm'
```

`session_preload_libraries = 'pg_ivm'` 也可以在所有相关会话中一致管理时使用。

```sql
CREATE EXTENSION pg_ivm;

SELECT pgivm.create_immv(
    'account_totals',
    'SELECT branch_id, count(*) AS accounts, sum(balance) AS balance
     FROM accounts
     GROUP BY branch_id'
);

UPDATE accounts
SET balance = balance + 100
WHERE account_id = 42;

SELECT * FROM account_totals;
```

### 管理和检查 IMMV

- `pgivm.create_immv(name, query)`: 创建并填充一个 IMMV，返回其行数。
- `pgivm.refresh_immv(name, with_data)`: 完全重建 IMMV；`false` 在后续填充刷新之前禁用维护。
- `pgivm.get_immv_def(regclass)`: 返回存储的视图定义。
- `pgivm.restore_immv(name, query, populate)`: 1.15 版本函数，用于重构现有 IMMV 表的元数据、触发器和索引。
- `pgivm.get_create_immv_commands()` 和 `pgivm.get_restore_immv_commands()`: 发出重建 IMMV 或恢复其元数据所需的 SQL。

1.15 版包括一个用于导出或 `pg_upgrade` 工作流的辅助工具：

```shell
pg_ivm_dump_metadata -d application > pg_ivm_metadata.sql
```

该脚本发出 `pgivm.restore_immv()` 调用。先恢复表数据，然后执行保存的元数据 SQL 以使增量维护继续进行而不必重新创建表。

### 限制和操作注意事项

- 支持的定义包括选择性连接、`DISTINCT`、简单的子查询/CTE 和内置 `count`、`sum`、`avg`、`min` 和 `max` 聚合。不支持的结构包括 `HAVING`、窗口函数、`ORDER BY`、`LIMIT/OFFSET`、集合操作、`DISTINCT ON` 和用户定义的聚合。
- 高效维护依赖于合适的唯一索引。当定义提供可用分组、去重或基表主键列时，`create_immv()` 会自动创建一个。
- 创建和刷新需要 `AccessExclusiveLock`。上游警告在 `REPEATABLE READ` 或 `SERIALIZABLE` 下创建的一致性风险；使用 `READ COMMITTED` 或者在之后进行刷新。
- 当关系已注册或其表定义与提供的查询不符时，`restore_immv()` 会失败。
- 1.15 版还修复了多次触发驱动修改后的不正确维护和 v1.14 的外连接维护崩溃问题。
