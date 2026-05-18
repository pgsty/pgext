## 用法

来源：[README](https://github.com/spa5k/pg_strict/blob/master/README.md), [Release v1.0.5](https://github.com/spa5k/pg_strict/releases/tag/v1.0.5), [API source](https://github.com/spa5k/pg_strict/blob/master/src/api.rs)

`pg_strict` 会阻断或警告省略 `WHERE` 子句的 `UPDATE` 和 `DELETE` 语句。它安装 `post_parse_analyze_hook`，因此必须通过 `shared_preload_libraries` 加载。

### 必要设置

```sql
-- postgresql.conf
shared_preload_libraries = 'pg_strict'

CREATE EXTENSION pg_strict;
```

### GUC

- `pg_strict.require_where_on_update`
- `pg_strict.require_where_on_delete`

每个设置都支持 `off`、`warn` 和 `on`。

```sql
SET pg_strict.require_where_on_update = 'on';
SET pg_strict.require_where_on_delete = 'warn';
```

### 辅助函数

```sql
SELECT pg_strict_version();
SELECT * FROM pg_strict_config();

SELECT pg_strict_check_where_clause('DELETE FROM t', 'DELETE');
SELECT pg_strict_validate_update('UPDATE t SET x = 1 WHERE id = 42');
SELECT pg_strict_validate_delete('DELETE FROM t WHERE id = 42');

SELECT pg_strict_enable_update();
SELECT pg_strict_warn_delete();
SELECT pg_strict_disable_delete();
```

- `pg_strict_set_update_mode(mode)` 和 `pg_strict_set_delete_mode(mode)` 提供通用模式设置函数。
- `SET LOCAL` 可在事务内用于一次性批量操作。

### 注意事项

- Enforcement 只检查是否存在 `WHERE`，不判断意图：任何非空 `WHERE` 子句都满足规则。
- 只检查 `UPDATE` 和 `DELETE`。
- 当前上游 release 是 `1.0.5`；上游文档记录 PostgreSQL 13 到 18，而 `db/extension.csv` 中的 Pigsty package row 覆盖 PostgreSQL 14 到 18。
- Pigsty 关于 `pgrx` 0.17.0 的备注是 packaging/build metadata，不是文档化的用户侧功能变化。
