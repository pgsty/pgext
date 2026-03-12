

## Usage

> [system_stats: system-level statistics for PostgreSQL](https://github.com/EnterpriseDB/system_stats)

system_stats provides SQL functions to access OS-level statistics for monitoring. Supports Linux, macOS, and Windows.

### Functions

```sql
-- Operating system info (name, version, hostname, uptime, etc.)
SELECT * FROM pg_sys_os_info();

-- CPU info (vendor, model, cores, clock speed, cache sizes)
SELECT * FROM pg_sys_cpu_info();

-- CPU usage (user, system, idle, iowait percentages)
SELECT * FROM pg_sys_cpu_usage_info();

-- Memory info (total, used, free, swap, cache in bytes)
SELECT * FROM pg_sys_memory_info();

-- Disk I/O analysis (reads, writes, bytes, time per device)
SELECT * FROM pg_sys_io_analysis_info();

-- Disk info (filesystem, mount point, total/used/available space)
SELECT * FROM pg_sys_disk_info();

-- Load average (1, 5, 10, 15 minute intervals)
SELECT * FROM pg_sys_load_avg_info();

-- Process info (total, running, sleeping, stopped, zombie counts)
SELECT * FROM pg_sys_process_info();

-- Network info (interface, IP, bytes/packets sent/received, errors)
SELECT * FROM pg_sys_network_info();

-- Per-process CPU and memory usage
SELECT * FROM pg_sys_cpu_memory_by_process();
```

### Security

Access is restricted to superusers and members of the `monitor_system_stats` role:

```sql
-- Grant access to a monitoring user
GRANT monitor_system_stats TO nagios;

-- Or grant per-function to pg_monitor
GRANT EXECUTE ON FUNCTION pg_sys_os_info() TO pg_monitor;
```

### Example: System Overview

```sql
SELECT * FROM pg_sys_load_avg_info();  -- load averages
SELECT memfree, memused, swapfree FROM pg_sys_memory_info();  -- memory
SELECT * FROM pg_sys_cpu_usage_info();  -- CPU usage breakdown
```
