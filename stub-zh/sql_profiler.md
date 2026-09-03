## 用法

来源：

- [SQL Profiler 官方文档](https://www.enterprisedb.com/docs/pg_extensions/sqlprofiler/)
- [官方安装指南](https://www.enterprisedb.com/docs/pg_extensions/sqlprofiler/installing/)
- [官方使用指南](https://www.enterprisedb.com/docs/pg_extensions/sqlprofiler/using/)

`sql_profiler` 为选定 user 与 database 捕获有界 query trace，其中包含 timing 与 explain 数据，可加载到临时表进行分析。

### 启用

安装 SQL Profiler 4.2.0，预加载带连字符的库名，重启并创建下划线命名的扩展：

```ini
shared_preload_libraries = 'sql-profiler'
```

```sql
CREATE EXTENSION sql_profiler;
GRANT sql_profiler_admin TO profiler_user;
```

`sql_profiler_admin` 角色允许非超级用户调用 profiler function；外围 PEM 集成还有额外 account 要求。

### 捕获并读取 Trace

`sp_activate` 返回 trace ID。应设置限制，避免遗忘的 trace 无界增长。

```sql
SELECT sp_activate(
  'checkout investigation',
  ''::oidvector,
  ''::oidvector,
  256,
  100,
  interval '15 minutes'
);

SELECT * FROM sp_active_traces();
SELECT sp_deactivate(1);
SELECT sp_load_trace(1, true);
```

加载后的行位于 `_sp_tmp_tbl_sql_profiler`；`sp_traces_list` 列出已完成 trace，`sp_cleanup` 删除 trace 数据。

### 运维边界

Trace 文件可能包含 query text、parameter、plan、username 与 database identifier。应限制 duration、最短语句时间和最大大小，保护 profiler role 与文件系统，并在分析后清理。即使没有活跃 trace，预加载仍会在整个集群安装 hook，因此应在代表性 workload 上验证开销。

