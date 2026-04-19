## Usage

Source: [README](https://github.com/EnterpriseDB/system_stats/blob/master/README.md), [Release v4.0](https://github.com/EnterpriseDB/system_stats/releases/tag/v4.0), [SQL install script](https://github.com/EnterpriseDB/system_stats/blob/master/system_stats--4.0.sql)

`system_stats` exposes operating-system metrics through SQL functions. It supports Linux, macOS, and Windows, returning `NULL` for fields that are not meaningful on the current platform.

### Main functions

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

These functions cover OS identity, CPU inventory and usage, memory, block-device I/O, disks, load average, process counts, network interfaces, and per-process CPU and memory usage.

### Access control

```sql
GRANT monitor_system_stats TO nagios;
GRANT EXECUTE ON FUNCTION pg_sys_os_info() TO pg_monitor;
```

- The extension creates a `monitor_system_stats` role and grants execution on the shipped functions to that role.
- Functions are revoked from `PUBLIC`.

### Caveats

- The `monitor_system_stats` role is not dropped automatically when the extension is removed.
- macOS cannot expose full per-process details for processes owned by other users; those rows may contain only PID and process name.
- Current v4.0 upstream docs keep the same user-facing function family and security model; the refresh here mainly aligns names, privileges, and platform notes with the current README and SQL script.
