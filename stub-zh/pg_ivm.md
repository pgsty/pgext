## 用法

来源：

- [官方 v1.17 README](https://github.com/sraoss/pg_ivm/blob/v1.17/README.md)
- [v1.17 发行说明](https://github.com/sraoss/pg_ivm/releases/tag/v1.17)
- [1.16 to 1.17 upgrade SQL](https://github.com/sraoss/pg_ivm/blob/v1.17/pg_ivm--1.16--1.17.sql)
- [pg_ivm_dump_metadata 工具](https://github.com/sraoss/pg_ivm/blob/v1.17/scripts/pg_ivm_dump_metadata)

`pg_ivm` 为 PostgreSQL 提供了即时增量视图维护功能。增量可维护材料化视图（IMMV）以表的形式存储在 `pgivm` 模式中，带有触发器和元数据；基表的变化会在同一个事务中更新 IMMV 而不是重新计算整个查询。

### 启用并创建一个 IMMV

为可以修改 IMMV 基表的每个会话加载库。集群级设置需要重启：

```ini
shared_preload_libraries = 'pg_ivm'
```

`session_preload_libraries = 'pg_ivm'` 也可以在所有相关会话中一致管理时使用。

```sql
CREATE EXTENSION pg_ivm;

CREATE TABLE accounts (
    account_id bigint PRIMARY KEY,
    branch_id integer NOT NULL,
    balance numeric NOT NULL
);
INSERT INTO accounts VALUES (42, 1, 1000);

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
- `pgivm.restore_immv(name, query, populate, replace)`: 重建现有 IMMV 表的元数据、触发器和索引；1.17 新增可选的替换标志。
- `pgivm.get_create_immv_commands()` 和 `pgivm.get_restore_immv_commands()`: 发出重建 IMMV 或恢复其元数据所需的 SQL。

1.15 版包括一个用于导出或 `pg_upgrade` 工作流的辅助工具：

```bash
pg_ivm_dump_metadata -d application > pg_ivm_metadata.sql
```

该脚本发出 `pgivm.restore_immv()` 调用。先恢复表数据，然后执行保存的元数据 SQL 以使增量维护继续进行而不必重新创建表。

### 限制和操作注意事项

- 支持的定义包括选择性连接、`DISTINCT`、简单的子查询/CTE 和内置 `count`、`sum`、`avg`、`min` 和 `max` 聚合。不支持的结构包括 `HAVING`、窗口函数、`ORDER BY`、`LIMIT/OFFSET`、集合操作、`DISTINCT ON` 和用户定义的聚合。
- 高效维护依赖于合适的唯一索引。当定义提供可用分组、去重或基表主键列时，`create_immv()` 会自动创建一个。
- 创建和刷新需要 `AccessExclusiveLock`。上游警告在 `REPEATABLE READ` 或 `SERIALIZABLE` 下创建的一致性风险；使用 `READ COMMITTED` 或者在之后进行刷新。
- 关系已注册时，`restore_immv()` 默认失败；设置 `replace = true` 可以替换，但要求超级用户。未启用替换时，调用者须为 IMMV 所有者。表定义不匹配时仍会失败。
- 1.15 版还修复了多次触发驱动修改后的不正确维护和 v1.14 的外连接维护崩溃问题。

### 1.16 升级

1.16 增加 PostgreSQL 19 支持，修复同一事务内维护后删除 IMMV 时的崩溃、无默认等值操作符列导致的维护失败、大 OID 处理以及增量维护锁释放时序问题。安装匹配的新文件并确保相关会话加载新动态库后，执行：

```sql
ALTER EXTENSION pg_ivm UPDATE TO '1.16';
```

### 升级至 1.17

上游支持 PostgreSQL 13–19。1.17 修复恢复后 IMMV 的填充状态，避免后续刷新创建重复维护触发器；同时改进自动唯一索引推导、拒绝目标列表中的集合返回函数，并修复 LATERAL 子查询的索引创建崩溃。

安装匹配文件，确保修改数据的会话使用新库，然后逐库更新：

```sql
ALTER EXTENSION pg_ivm UPDATE TO '1.17';
```

执行 PostgreSQL 大版本升级前导出元数据时，把新增选项放在首个参数位置：

```bash
pg_ivm_dump_metadata --pg-upgrade application > pg_ivm_metadata.sql
```

生成的恢复调用会替换大版本升级后残留的旧目录记录。完成升级后，以超级用户执行保存的 SQL。IMMV 表及其数据必须已经存在；替换元数据不能代替数据恢复。
