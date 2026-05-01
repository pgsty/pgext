## 用法

来源：[README](https://github.com/viggy28/pg_savior/blob/v0.1.0/README.md), [release 0.1.0](https://github.com/viggy28/pg_savior/releases/tag/v0.1.0), [PGXN 0.1.0](https://pgxn.org/dist/pg_savior/0.1.0/), [SQL file](https://github.com/viggy28/pg_savior/blob/v0.1.0/pg_savior--0.1.0.sql), [C source](https://github.com/viggy28/pg_savior/blob/v0.1.0/pg_savior.c), [pg_savior.control](https://github.com/viggy28/pg_savior/blob/v0.1.0/pg_savior.control)

`pg_savior` 是一个 PostgreSQL 安全扩展，用于在执行前阻止特定高风险 DML 和 DDL 语句。版本 `0.1.0` 是有意发布到 PGXN 的版本，并且相对 `0.0.1` 进行了重大重写；README 仍将其标记为 pre-1.0，且未准备好用于生产。

### 激活

仅执行 `CREATE EXTENSION` 不会激活检查。SQL 文件只说明保护逻辑位于已加载的 shared library 中，因此每个 backend 都必须通过一种上游支持的路径加载 `pg_savior`：

集群级激活使用 `shared_preload_libraries`，并需要重启 PostgreSQL：

```conf
shared_preload_libraries = 'pg_savior'
```

新连接的会话级激活可以在 config reload 后使用 `session_preload_libraries`：

```conf
session_preload_libraries = 'pg_savior'
```

开发或测试会话可以手动加载 library：

```sql
LOAD 'pg_savior';
CREATE EXTENSION pg_savior;
```

library 加载后，`_PG_init` 会为该 backend 安装 `post_parse_analyze_hook`、`ExecutorStart_hook` 和 `ProcessUtility_hook`。

### DML 保护

`pg_savior` 会阻止没有 `WHERE` 子句的 `DELETE` 和 `UPDATE` 语句。parser hook 检查分析后的 `Query` tree 并抛出 `ERROR`，因此事务会中止，应用会看到失败。

```sql
CREATE TABLE emp (id int);
INSERT INTO emp VALUES (1), (2), (3);

DELETE FROM emp;
-- ERROR: pg_savior: DELETE without WHERE clause is blocked

UPDATE emp SET id = id + 1;
-- ERROR: pg_savior: UPDATE without WHERE clause is blocked

DELETE FROM emp WHERE id = 1;
-- allowed
```

可选的行数保护适用于 planner estimate 超过 `pg_savior.max_rows_affected` 的 `DELETE` 和 `UPDATE` 语句。它从 `ExecutorStart_hook` 运行，在规划之后、触碰 tuple 之前生效。

```sql
SET pg_savior.max_rows_affected = 100;

DELETE FROM emp WHERE id > 0;
-- blocked if the planner estimate is greater than 100 rows
```

### DDL 保护

`ProcessUtility_hook` 只保护上游列出的 DDL 操作：

- 没有 `CONCURRENTLY` 的 `CREATE INDEX` 总是被阻止。
- `DROP DATABASE` 总是被阻止。
- 当目标表大于 `pg_savior.large_table_threshold_rows` 时，`ALTER TABLE ADD COLUMN ... DEFAULT` 被阻止。
- 大表上的 `ALTER TABLE ALTER COLUMN TYPE` 被阻止。
- 当任一目标表是大表时，`TRUNCATE` 被阻止。
- 当任一目标表是大表时，`DROP TABLE` 被阻止。

大表检查使用 `pg_class.reltuples > pg_savior.large_table_threshold_rows`。

```sql
CREATE INDEX emp_idx ON emp (id);
-- ERROR: pg_savior: CREATE INDEX without CONCURRENTLY is blocked

CREATE INDEX CONCURRENTLY emp_idx ON emp (id);
-- allowed by this guard

ALTER TABLE big_emp ADD COLUMN status text DEFAULT 'active';
-- blocked when big_emp is over pg_savior.large_table_threshold_rows

TRUNCATE big_emp;
-- blocked when big_emp is over pg_savior.large_table_threshold_rows
```

### 配置

所有已文档化的 GUC 都是 session-scoped `USERSET` 变量：

| GUC | Default | Effect |
|---|---:|---|
| `pg_savior.enabled` | `on` | 总开关；为 `off` 时不运行检查。 |
| `pg_savior.bypass` | `off` | 允许当前 session 跳过保护。 |
| `pg_savior.max_rows_affected` | `0` | 当估计的 `DELETE`/`UPDATE` 行数高于该值时阻止；`0` 禁用该检查。 |
| `pg_savior.large_table_threshold_rows` | `1000000` | 为受保护的大表 DDL 操作定义 "large"。 |

有意绕过一个事务时使用 `SET LOCAL`：

```sql
BEGIN;
SET LOCAL pg_savior.bypass = on;
DELETE FROM staging_table;
COMMIT;
```

### 注意事项

- 保护生效前，library 必须先在 backend 中加载；`CREATE EXTENSION pg_savior` 只注册扩展元数据。
- 行数保护和大表保护依赖 planner/catalog estimate。近期变更导致 estimate 过期时，应运行 `ANALYZE`。
- `UPDATE` 覆盖范围限于无保护的 `UPDATE` 和可选 planner-estimate 阈值；README 未声称会语义审查每个 `WHERE` predicate。
- DDL 覆盖范围限于列出的 `ProcessUtility_hook` case。不要假设其他 schema 变更会被阻止。
- `ADD COLUMN ... DEFAULT` 保护较保守，会阻止大表上的任何 default，包括较新 PostgreSQL 版本可能无需 full table rewrite 处理的 non-volatile default。
