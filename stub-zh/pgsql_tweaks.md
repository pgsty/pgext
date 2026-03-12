

## 用法

> [pgsql_tweaks: 为 DBA 日常工作提供的 PostgreSQL 视图与函数](https://codeberg.org/pgsql_tweaks/pgsql_tweaks)

所有对象创建在 `pgsql_tweaks` 模式中。完整文档：[rtfm.pgsql-tweaks.org](https://rtfm.pgsql-tweaks.org)

### 数据类型检查函数

```sql
SELECT pgsql_tweaks.is_date('2024-01-15');       -- true
SELECT pgsql_tweaks.is_integer('42');             -- true
SELECT pgsql_tweaks.is_numeric('3.14');           -- true
SELECT pgsql_tweaks.is_json('{"a":1}');           -- true
SELECT pgsql_tweaks.is_jsonb('{"a":1}');          -- true
SELECT pgsql_tweaks.is_boolean('true');            -- true
SELECT pgsql_tweaks.is_timestamp('2024-01-15 10:30:00');  -- true
SELECT pgsql_tweaks.is_hex('FF');                 -- true
```

### 系统信息函数

```sql
SELECT pgsql_tweaks.pg_schema_size('public');     -- 模式大小（字节）
```

### 聚合函数

- `gap_fill` -- 填充时间序列中的空缺
- `array_min`、`array_max`、`array_avg`、`array_sum` -- 数组聚合

### 转换函数

```sql
SELECT pgsql_tweaks.to_unix_timestamp(now());
SELECT pgsql_tweaks.hex2bigint('FF');
```

### 工具函数

```sql
SELECT pgsql_tweaks.is_empty('');                 -- true
SELECT pgsql_tweaks.array_trim(ARRAY['a','','b']);
```

### 系统信息视图

- `pg_db_views`、`pg_foreign_keys`、`pg_functions`、`pg_active_locks`
- `pg_table_matview_infos`、`pg_object_ownership`、`pg_unused_indexes`
- `pg_bloat_info`、`pg_missing_indexes`、`pg_role_permissions`

### 监控视图

- `monitoring_wal`、`monitoring_active_locks`、`monitoring_replication`
- `monitoring_database_conflicts`、`monitoring_vacuum`
- `statistics_top_ten_query_times`、`statistics_query_activity`
