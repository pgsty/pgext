

## Usage

> [pgnodemx: SQL functions to access node OS metrics from PostgreSQL](https://github.com/CrunchyData/pgnodemx)

pgnodemx provides SQL access to OS-level metrics including cgroup statistics, `/proc` filesystem data, and system information. Requires `pg_monitor` role membership.

### cgroup Functions

```sql
-- Check cgroup support status
SELECT current_setting('pgnodemx.cgroup_enabled');
SELECT cgroup_mode();  -- 'legacy', 'unified', 'hybrid', or 'disabled'

-- Read cgroup scalar values
SELECT cgroup_scalar_bigint('memory.current');
SELECT cgroup_scalar_float8('cpu.uclamp.max');
SELECT cgroup_scalar_text('cgroup.type');

-- Read cgroup key-value files
SELECT * FROM cgroup_setof_kv('memory.stat');
SELECT * FROM cgroup_setof_kv('cpu.stat');

-- Read cgroup nested key-value files
SELECT * FROM cgroup_setof_nkv('memory.pressure');
SELECT * FROM cgroup_setof_nkv('cpu.pressure');

-- Get cgroup paths and process count
SELECT * FROM cgroup_path();
SELECT cgroup_process_count();
```

### /proc Functions

```sql
SELECT * FROM proc_diskstats();       -- /proc/diskstats
SELECT * FROM proc_mountinfo();       -- /proc/self/mountinfo
SELECT * FROM proc_meminfo();         -- /proc/meminfo
SELECT * FROM proc_network_stats();   -- /proc/self/net/dev
SELECT * FROM proc_pid_io();          -- /proc/<pid>/io for all PG processes
SELECT * FROM proc_pid_cmdline();     -- command line for all PG processes
SELECT * FROM proc_pid_stat();        -- /proc/<pid>/stat for all PG processes
SELECT * FROM proc_cputime();         -- first line of /proc/stat
SELECT * FROM proc_loadavg();         -- /proc/loadavg
```

### System Information

```sql
SELECT * FROM fsinfo('/pgdata');      -- filesystem info for a path
SELECT fips_mode();                   -- OpenSSL FIPS mode status
SELECT openssl_version();             -- OpenSSL version
SELECT exec_path();                   -- current PostgreSQL executable path
SELECT kpages_to_bytes(1024);         -- convert kernel pages to bytes
SELECT * FROM stat_file('/path');     -- file uid, gid, mode info
```

### Environment Variables

```sql
SELECT envvar_text('PGDATA');
SELECT envvar_bigint('PGPORT');
```

### Kubernetes DownwardAPI

```sql
SELECT * FROM kdapi_setof_kv('labels');
SELECT kdapi_scalar_bigint('cpu_limit');
```

### Configuration

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pgnodemx.cgroup_enabled` | `on` | Enable cgroup functions |
| `pgnodemx.containerized` | `off` | Force containerized cgroup paths |
| `pgnodemx.cgrouproot` | `/sys/fs/cgroup` | cgroup mount location |
| `pgnodemx.kdapi_enabled` | `on` | Enable Kubernetes DownwardAPI |
| `pgnodemx.kdapi_path` | `/etc/podinfo` | DownwardAPI files location |
