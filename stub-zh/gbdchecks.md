## 用法

来源：

- [Pinned official README](https://github.com/huseynsnmz/gbdchecks/blob/ff7090d1a2a7064df6d43deeab4f10cf7c09ff0d/README.md)
- [Pinned extension SQL](https://github.com/huseynsnmz/gbdchecks/blob/ff7090d1a2a7064df6d43deeab4f10cf7c09ff0d/sql/gbdchecks--1.1.sql)

`gbdchecks` 安装一组诊断视图和权限报告函数，覆盖查询活动、膨胀、锁、缓存命中率、索引、顺序扫描、清理状态及有效权限。它是面向快照的排障工具，不是调度器或告警服务。

### 核心流程

扩展使用固定的 `gbdchecks` 模式，并依赖 `pg_stat_statements`：

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION gbdchecks;

SELECT * FROM gbdchecks.check_blocked_statements;
SELECT * FROM gbdchecks.check_unused_indexes;
SELECT * FROM gbdchecks.all_privs('app_user');
```

若要使用 `gbdchecks.check_queries`，需把 `pg_stat_statements` 配置到 `shared_preload_libraries`，重启 PostgreSQL 后再收集统计信息。没有该查询统计视图时，其他视图仍可使用。

### 诊断视图

- 工作负载与等待：`check_queries`、`check_blocked_statements`、`check_locks`。
- 存储与维护：`check_bloat`、`check_index_sizes`、`check_vacuum_stats`。
- 访问模式：`check_hit_ratio`、`check_index_usage`、`check_seq_scans`、`check_unused_indexes`。

这些视图展示实时的目录和统计数据。较低的扫描次数、估算的膨胀比例或预期触发 autovacuum 都只是进一步调查的证据，不能单独作为删除索引或重写表的指令。

### 权限报告

函数以角色名 `text` 为参数：`table_privs`、`database_privs`、`tablespace_privs`、`fdw_privs`、`fsrv_privs`、`language_privs`、`schema_privs`、`view_privs`、`sequence_privs`。`all_privs(text)` 会合并它们的输出。

### 安全与兼容性

安装过程会撤销公众对扩展模式、函数和视图的访问，再向创建者授予使用、执行和查询权限。只应向其他角色授予所需权限，例如：

```sql
GRANT USAGE ON SCHEMA gbdchecks TO monitor;
GRANT SELECT ON ALL TABLES IN SCHEMA gbdchecks TO monitor;
GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA gbdchecks TO monitor;
```

该项目已经归档，其 SQL 引用了历史上的 `pg_stat_statements` 列名，例如 `total_time` 和 `mean_time`。这些名称在较新的 PostgreSQL/扩展版本中已经变化，因此投入运维使用前，必须在目标服务端版本上验证扩展创建以及每个视图。
