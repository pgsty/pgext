## 用法

来源：

- [TencentDB for PostgreSQL 进程监控](https://cloud.tencent.com/document/product/409/96894)
- [TencentDB 支持的扩展版本](https://cloud.tencent.com/document/product/409/75121)

`tencentdb_system_stat` 1.0 是 TencentDB for PostgreSQL 提供的进程资源采样监控扩展。它提供 `tencentdb_process_system_usage` 视图，其中包含后端身份、当前查询文本、CPU 使用率和内存字节数。当前供应商矩阵为 PostgreSQL 11 至 18 提供 1.0 版，而 PostgreSQL 10 没有可用版本。

### 检查采样的后端资源使用情况

```sql
CREATE EXTENSION tencentdb_system_stat;

SELECT pid, username, datname, backend_type, query,
       cpu_usage, memory_bytes
FROM tencentdb_process_system_usage;

SHOW tencentdb_system_stat.sampling_interval;
```

`cpu_usage` 覆盖当前采样周期。`tencentdb_system_stat.sampling_interval` 设置以毫秒为单位，默认值为 1000。`memory_bytes` 以字节为单位报告。应把这些值视为采样观测，而不是累计记账总量。

### 供应商与可见性边界

该扩展由腾讯云在 TencentDB 内提供并维护版本，不是面向自建 PostgreSQL 的可移植源码包。将视图写入监控自动化之前，应确认实际实例引擎版本是否可用，并容忍供应商升级、采样缺失和新增后端类型。

该视图包含用户名、数据库名、进程标识符和当前 SQL 文本。应依据租户与查询保密要求限制访问，避免把原始查询文本导出到广泛可读的遥测系统。缩短采样间隔会增加监控工作；只有在托管实例上测量开销并遵循供应商参数控制后才应修改。
