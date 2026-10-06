## 用法

来源：

- [external/pg_cron_preload/pg_cron_preload--1.0.sql](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/fd9fc37f6a4d6fce58b61af39c18cb979fa15419/external/pg_cron_preload/pg_cron_preload--1.0.sql)
- [external/pg_cron_preload/invalid_cache.c](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/fd9fc37f6a4d6fce58b61af39c18cb979fa15419/external/pg_cron_preload/invalid_cache.c)
- [external/pg_cron_preload/Makefile](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/fd9fc37f6a4d6fce58b61af39c18cb979fa15419/external/pg_cron_preload/Makefile)
- [external/pg_cron_preload/pg_cron_preload.control](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/fd9fc37f6a4d6fce58b61af39c18cb979fa15419/external/pg_cron_preload/pg_cron_preload.control)

`pg_cron_preload` 是 PolarDB 中独立安装的 cron 目录组件，创建任务和运行记录表，并用 C 触发器使任务缓存失效，本身不实现调度器。

### 核心用法

```sql
CREATE EXTENSION pg_cron_preload;
SELECT jobid, schedule, database, username, active FROM cron.job;
SELECT jobid, status, start_time, end_time FROM cron.job_run_details;
```

### 运行边界

SQL 安装 `cron.job`、`cron.job_run_details`、序列、行安全策略及 `cron.job_cache_invalidate`。该组件中调度和取消调度函数的声明被注释掉，应配合对应的 PolarDB 调度集成使用，不能套用普通 `pg_cron` API。

仅应由超级用户在预期的 PolarDB 安装中创建。控制文件固定扩展模式，而 SQL 另行创建 `cron`，已有 cron 对象可能冲突。虽然名字含 preload，C 模块并没有预加载初始化钩子。此目录记录不宣称兼容原生 PostgreSQL，也不表示它是独立调度器。
