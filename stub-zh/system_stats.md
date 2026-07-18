


## 用法

来源：[README](https://github.com/EnterpriseDB/system_stats/blob/master/README.md), [Release v4.0](https://github.com/EnterpriseDB/system_stats/releases/tag/v4.0), [SQL install script](https://github.com/EnterpriseDB/system_stats/blob/master/system_stats--4.0.sql)

`system_stats` 通过 SQL 函数暴露操作系统指标。它支持 Linux、macOS 和 Windows；对于当前平台没有意义的字段，会返回 `NULL`。

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

这些函数覆盖操作系统身份、CPU 清单与使用率、内存、块设备 I/O、磁盘、负载均值、进程数量、网络接口，以及各进程的 CPU 与内存使用情况。

### 访问控制

```sql
GRANT monitor_system_stats TO nagios;
GRANT EXECUTE ON FUNCTION pg_sys_os_info() TO pg_monitor;
```

- 扩展会创建 `monitor_system_stats` 角色，并把随扩展提供的函数执行权授予该角色。
- 函数会从 `PUBLIC` 撤销权限。

### 注意事项

- 删除扩展时，`monitor_system_stats` 角色不会自动删除。
- macOS 无法暴露其他用户所拥有进程的完整详情；这些行可能只包含 PID 和进程名。
- Pigsty 元数据跟踪 `system_stats` 4.0，覆盖 PostgreSQL 14-18；RPM 和 DEB 包名不同（`system_stats_$v` 与 `postgresql-$v-system-stats`）。
- 当前 v4.0 上游文档保留了相同的用户侧函数家族和安全模型；本次刷新主要是让名称、权限和平台说明与当前 README 和 SQL 脚本对齐。
