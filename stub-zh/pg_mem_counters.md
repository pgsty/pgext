## 用法

来源：

- [固定提交的上游 README](https://github.com/topharley/pg_mem_counters/blob/3a6a6f4e2e0b62fa23ef1041a7ae357d016e600f/README.md)
- [固定提交的扩展 control 文件](https://github.com/topharley/pg_mem_counters/blob/3a6a6f4e2e0b62fa23ef1041a7ae357d016e600f/pg_mem_counters.control)

`pg_mem_counters` 在 PostgreSQL 共享内存中维护具名的累计命中数和每分钟请求数，适用于不要求持久化或事务记账的轻量进程内指标。

必须预加载该库、重启 PostgreSQL，然后才能创建扩展：

```conf
shared_preload_libraries = 'pg_mem_counters'
```

```sql
CREATE EXTENSION pg_mem_counters;

SELECT inc_mem_counter('api_requests', 1);
SELECT inc_mem_counter('api_requests');
SELECT get_mem_counter_rpm('api_requests');
SELECT * FROM mem_counters();
```

`inc_mem_counter(counter,increment)` 会增加并返回累计值；省略增量时用于读取累计值。`get_mem_counter_rpm(counter)` 返回近期每分钟计数，`mem_counters()` 则列出各名称及其累计值和 RPM 值。

状态保存在共享内存而不是持久数据库表中，因此不要把它用于计费、审计、配额，或任何必须跨重启保留、参与回滚的计数器。已复核 README 只记录 PostgreSQL 10 和 11，control 版本为 `1.0`；部署前应测试内存容量、名称基数、并发、重启行为和当前服务器 ABI。
