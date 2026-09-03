## 用法

来源：

- [Clone Schema 官方设置指南](https://www.enterprisedb.com/docs/epas/latest/database_administration/14_edb_clone_schema/setting_up_edb_clone_schema/)
- [Clone Schema 官方文档](https://www.enterprisedb.com/docs/epas/latest/database_administration/14_edb_clone_schema/)
- [EDB 官方扩展目录](https://www.enterprisedb.com/docs/pg_extensions/)

`parallel_clone` 是 `edb_cloneschema` 使用的 EPAS worker 扩展，通过并行 background process 复制 schema object。

### 启用

安装匹配的 EPAS 软件包，与 `edb_job_scheduler` 一起预加载 `parallel_clone`，重启后在每个 source 或 target database 中创建：

```ini
shared_preload_libraries = 'parallel_clone,edb_job_scheduler'
```

```sql
CREATE EXTENSION parallel_clone;
```

该 provider-only module 只支持 EDB Postgres Advanced Server，通常作为 Clone Schema 工作流的一部分安装。

### 与 Clone Schema 的关系

只有在 `parallel_clone`、`edb_job_scheduler`、`dbms_job`、`postgres_fdw`、`dblink` 与 PL/Perl prerequisite 就绪后，才创建 `edb_cloneschema`。

```sql
CREATE EXTENSION edb_cloneschema;
SELECT edb_util.create_clone_log_dir();
```

### 运维边界

Parallel clone worker 会消耗 `max_worker_processes`、内存、WAL、lock、checkpoint 与日志空间。大型 schema copy 可能在一个事务中运行，需要有意调优。应把 `parallel_clone` 视为内部 runtime dependency：通过文档化的 `edb_cloneschema` API 使用、监控状态日志，不要直接调用未公开的 worker function。

