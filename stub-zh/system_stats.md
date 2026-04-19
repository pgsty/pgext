## 用法

来源：[README](https://github.com/EnterpriseDB/system_stats/blob/master/README.md)，[Release v4.0](https://github.com/EnterpriseDB/system_stats/releases/tag/v4.0)，[SQL install script](https://github.com/EnterpriseDB/system_stats/blob/master/system_stats--4.0.sql)

`system_stats` 通过 SQL 函数暴露操作系统指标。它支持 Linux、macOS 与 Windows，对于当前平台无意义的字段会返回 `NULL`。

### 主要函数

```sql
CREATE EXTENSION system_stats;

SELECT * FROM pg_sys_os_info();
SELECT * FROM pg_sys_cpu_info();
SELECT * FROM pg_sys_cpu_usage_info();
SELECT * FROM pg_sys_memory_info();
SELECT * FROM pg_sys_io_analysis_info();
SELECT * FROM pg_sys_disk_info();
SELECT * FROM pg_sys_load_avg_info();
SELECT * FROM pg_sys_process_info();
SELECT * FROM pg_sys_network_info();
SELECT * FROM pg_sys_cpu_memory_by_process();
```

这些函数覆盖 OS identity、CPU inventory 与 usage、memory、block-device I/O、disks、load average、process counts、network interfaces，以及 per-process CPU 与 memory usage。

### 访问控制

```sql
GRANT monitor_system_stats TO nagios;
GRANT EXECUTE ON FUNCTION pg_sys_os_info() TO pg_monitor;
```

- 该扩展会创建 `monitor_system_stats` role，并将内置函数的执行权限授予该 role。
- 所有函数都会从 `PUBLIC` 撤销。

### 注意事项

- 删除扩展时不会自动删除 `monitor_system_stats` role。
- 在 macOS 上，系统无法暴露其他用户拥有进程的完整 per-process 细节；这些行可能只包含 PID 与进程名。
- 当前 v4.0 上游文档保持相同的用户函数族和安全模型；这次刷新主要是使名称、权限和平台说明与当前 README 和 SQL script 对齐。
