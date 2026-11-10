


## 用法

来源：

- [PGXN 元数据](https://api.pgxn.org/dist/pg_orca/1.1.0/META.json)
- [官方 README](https://api.pgxn.org/src/pg_orca/pg_orca-1.1.0/README.md)
- [控制文件](https://api.pgxn.org/src/pg_orca/pg_orca-1.1.0/pg_orca.control)
- [1.0 到 1.1 升级 SQL](https://api.pgxn.org/src/pg_orca/pg_orca-1.1.0/pg_orca--1.0.0--1.1.0.sql)

`pg_orca` 将 Greenplum/Apache Cloudberry 体系中的 ORCA 代价优化器接入标准 PostgreSQL 18/19 服务器。上游 README 将该项目描述为以 PostgreSQL 18 为主，并提供 PostgreSQL 19 预览支持。

### 在会话中启用 ORCA

`CREATE EXTENSION` 会在当前会话中加载该库，因此 `pg_orca.*` GUC 与 planner hook 会立即可用：

```sql
CREATE EXTENSION pg_orca;

SET pg_orca.enable_orca = on;

EXPLAIN
SELECT *
FROM orders
WHERE customer_id = 42
  AND created_at >= now() - interval '30 days';
```

如果 ORCA 无法处理某条查询，README 说明它会自动回退到标准 PostgreSQL planner。验证工作负载覆盖情况时，可以打开回退日志：

```sql
SET pg_orca.trace_fallback = on;
```

### 为新连接预加载

为了让后续会话自动加载 planner hook，上游建议使用 `session_preload_libraries`，而不是 `shared_preload_libraries`：

```sql
ALTER DATABASE mydb SET session_preload_libraries = 'pg_orca';
ALTER DATABASE mydb SET pg_orca.enable_orca = on;
```

已有会话不会受影响，必须重新连接后才会生效。如果已经配置了其他 session preload library，需要显式写出所有值：

```sql
ALTER DATABASE mydb
SET session_preload_libraries = 'pg_orca,pg_stat_statements';
```

按角色和按集群设置也有效：

```sql
ALTER ROLE bench SET session_preload_libraries = 'pg_orca';

ALTER SYSTEM SET session_preload_libraries = 'pg_orca';
SELECT pg_reload_conf();
```

### 调优与诊断

README 记录了以下主要 GUC：

- `pg_orca.enable_orca`：启用 ORCA；默认 `off`。
- `pg_orca.trace_fallback`：记录回退到标准 planner 的情况；默认 `off`。
- `optimizer_segments`：用于代价估算的 segment 数量；默认 `1`。
- `optimizer_sort_factor`：排序代价缩放因子；默认 `1.0`。
- `optimizer_metadata_caching`：缓存关系元数据；默认 `on`。
- `optimizer_mdcache_size`：元数据缓存大小，单位 KB；默认 `16384`。
- `optimizer_search_strategy_path`：可选的自定义搜索策略 XML 路径。

入口源代码还定义了额外的 ORCA 调优和调试 GUC，例如 `optimizer_join_order`、`pg_orca.join_order_dynamic_threshold`、`pg_orca.cost_model` 和 `optimizer_print_*`。这些更适合作为工作负载或调试旋钮；在把它们写入持久化数据库设置前，应先验证实际执行计划。

### 注意事项

- PostgreSQL 18 是主目标；PostgreSQL 19 属于预览支持。
- 新会话自动加载时使用 `session_preload_libraries = 'pg_orca'`。
- 加载后 ORCA 默认仍是禁用状态；需要设置 `pg_orca.enable_orca = on`。
- 对不支持的查询回退到 PostgreSQL planner 是预期行为；检查覆盖范围时可启用 `pg_orca.trace_fallback`。

### PostgreSQL 19 预览

1.1.0 通过版本条件编译支持 PostgreSQL 19 beta/RC，并覆盖新的窗口函数语法。服务器正式发布前应将 PostgreSQL 19 视为预览版本；大范围启用 ORCA 前请验证执行计划变化。
