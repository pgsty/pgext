## 用法

来源：

- [EDB Stat Monitor 官方文档](https://www.enterprisedb.com/docs/pg_extensions/edb_stat_monitor/)
- [官方配置指南](https://www.enterprisedb.com/docs/pg_extensions/edb_stat_monitor/configuring/)
- [官方使用指南](https://www.enterprisedb.com/docs/pg_extensions/edb_stat_monitor/using/)

`edb_stat_monitor` 为 PostgreSQL 与 EDB Postgres 记录按 bucket 划分的 query execution statistic，并在 pg_stat_monitor 模型上增加 EDB 支持的分析能力。

### 启用

安装 EDB Stat Monitor 2.1.0，预加载、重启并创建扩展：

```ini
shared_preload_libraries = 'edb_stat_monitor'
```

```sql
CREATE EXTENSION edb_stat_monitor;
```

每个需要查询统计信息的数据库都必须创建该扩展。

### 查询采集结果

```sql
SELECT application_name,
       userid::regrole AS user_name,
       datname,
       calls,
       query
FROM edb_stat_monitor
ORDER BY calls DESC;
```

应使用文档中的 bucket、plan、histogram、client 与 error 字段区分 workload interval，避免盲目聚合不同执行。

### 运维边界

采集使用 shared memory，并增加逐语句统计开销。大范围启用前应规划 retention、bucket、query text、plan capture 与 histogram 配置，并限制访问，因为规范化文本、client address、username 与 plan 可能暴露敏感 workload 细节。扩展升级不会自动改写依赖 view column 的应用查询；必须先阅读 release note。

