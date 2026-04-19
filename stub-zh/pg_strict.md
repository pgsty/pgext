## 用法

来源：[README](https://github.com/spa5k/pg_strict/blob/master/README.md)，[Release v1.0.5](https://github.com/spa5k/pg_strict/releases/tag/v1.0.5)，[API source](https://github.com/spa5k/pg_strict/blob/master/src/api.rs)

`pg_strict` 会阻止或警告缺少 `WHERE` 子句的 `UPDATE` 与 `DELETE` 语句。它通过 `post_parse_analyze_hook` 工作，因此必须从 `shared_preload_libraries` 加载。

### 所需设置

```sql
-- postgresql.conf
shared_preload_libraries = 'pg_strict'

CREATE EXTENSION pg_strict;
```

### GUCs

- `pg_strict.require_where_on_update`
- `pg_strict.require_where_on_delete`

每个设置都支持 `off`、`warn` 与 `on`。

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

- `pg_strict_set_update_mode(mode)` 与 `pg_strict_set_delete_mode(mode)` 提供通用模式设置器。
- `SET LOCAL` 可用于事务中的一次性 bulk operation。

### 注意事项

- 它检查的是 `WHERE` 的存在性，而不是语义意图：任何非空 `WHERE` 子句都会满足规则。
- 仅检查 `UPDATE` 与 `DELETE`。
- 当前上游版本是 `1.0.5`；Pigsty 关于 `pgrx` 0.17.0 的说明属于打包或构建元数据，不是文档化的用户功能变化。
