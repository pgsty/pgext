
## 用法

> [system_stats: PostgreSQL 系统级统计信息](https://github.com/EnterpriseDB/system_stats)

`system_stats` 提供 SQL 函数，用于访问操作系统级统计信息并支持监控。支持 Linux、macOS 和 Windows。

### 函数

```sql
-- 操作系统信息（名称、版本、主机名、运行时间等）
SELECT * FROM pg_sys_os_info();

-- CPU 信息（厂商、型号、核心数、时钟频率、缓存大小）
SELECT * FROM pg_sys_cpu_info();

-- CPU 使用率（用户态、系统态、空闲、iowait 百分比）
SELECT * FROM pg_sys_cpu_usage_info();

-- 内存信息（总量、已用、空闲、交换区、缓存，单位字节）
SELECT * FROM pg_sys_memory_info();

-- 磁盘 I/O 分析（每个设备的读写次数、字节数、耗时）
SELECT * FROM pg_sys_io_analysis_info();

-- 磁盘信息（文件系统、挂载点、总量/已用/可用空间）
SELECT * FROM pg_sys_disk_info();

-- 负载均值（1、5、10、15 分钟区间）
SELECT * FROM pg_sys_load_avg_info();

-- 进程信息（总数、运行中、睡眠、停止、僵尸进程数）
SELECT * FROM pg_sys_process_info();

-- 网络信息（接口、IP、发送/接收的字节数/包数、错误数）
SELECT * FROM pg_sys_network_info();

-- 每个进程的 CPU 和内存使用情况
SELECT * FROM pg_sys_cpu_memory_by_process();
```

### 安全性

访问限制为超级用户和 `monitor_system_stats` 角色的成员：

```sql
-- 授予监控用户访问权限
GRANT monitor_system_stats TO nagios;

-- 或按函数授予 pg_monitor 权限
GRANT EXECUTE ON FUNCTION pg_sys_os_info() TO pg_monitor;
```

### 示例：系统概览

```sql
SELECT * FROM pg_sys_load_avg_info();  -- 负载均值
SELECT memfree, memused, swapfree FROM pg_sys_memory_info();  -- 内存
SELECT * FROM pg_sys_cpu_usage_info();  -- CPU 使用率明细
```
