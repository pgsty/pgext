


## 用法

> [pgnodemx: 从 PostgreSQL 通过 SQL 访问节点操作系统指标](https://github.com/CrunchyData/pgnodemx)

pgnodemx 提供 SQL 访问操作系统级指标的能力，包括 cgroup 统计、`/proc` 文件系统数据和系统信息。需要 `pg_monitor` 角色成员身份。

### cgroup 函数

```sql
-- 检查 cgroup 支持状态
SELECT current_setting('pgnodemx.cgroup_enabled');
SELECT cgroup_mode();  -- 'legacy'、'unified'、'hybrid' 或 'disabled'

-- 读取 cgroup 标量值
SELECT cgroup_scalar_bigint('memory.current');
SELECT cgroup_scalar_float8('cpu.uclamp.max');
SELECT cgroup_scalar_text('cgroup.type');

-- 读取 cgroup 键值文件
SELECT * FROM cgroup_setof_kv('memory.stat');
SELECT * FROM cgroup_setof_kv('cpu.stat');

-- 读取 cgroup 嵌套键值文件
SELECT * FROM cgroup_setof_nkv('memory.pressure');
SELECT * FROM cgroup_setof_nkv('cpu.pressure');

-- 获取 cgroup 路径和进程数
SELECT * FROM cgroup_path();
SELECT cgroup_process_count();
```

### /proc 函数

```sql
SELECT * FROM proc_diskstats();       -- /proc/diskstats
SELECT * FROM proc_mountinfo();       -- /proc/self/mountinfo
SELECT * FROM proc_meminfo();         -- /proc/meminfo
SELECT * FROM proc_network_stats();   -- /proc/self/net/dev
SELECT * FROM proc_pid_io();          -- 所有 PG 进程的 /proc/<pid>/io
SELECT * FROM proc_pid_cmdline();     -- 所有 PG 进程的命令行
SELECT * FROM proc_pid_stat();        -- 所有 PG 进程的 /proc/<pid>/stat
SELECT * FROM proc_cputime();         -- /proc/stat 的第一行
SELECT * FROM proc_loadavg();         -- /proc/loadavg
```

### 系统信息

```sql
SELECT * FROM fsinfo('/pgdata');      -- 路径的文件系统信息
SELECT fips_mode();                   -- OpenSSL FIPS 模式状态
SELECT openssl_version();             -- OpenSSL 版本
SELECT exec_path();                   -- 当前 PostgreSQL 可执行文件路径
SELECT kpages_to_bytes(1024);         -- 将内核页转换为字节
SELECT * FROM stat_file('/path');     -- 文件的 uid、gid、mode 信息
```

### 环境变量

```sql
SELECT envvar_text('PGDATA');
SELECT envvar_bigint('PGPORT');
```

### Kubernetes DownwardAPI

```sql
SELECT * FROM kdapi_setof_kv('labels');
SELECT kdapi_scalar_bigint('cpu_limit');
```

### 配置

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `pgnodemx.cgroup_enabled` | `on` | 启用 cgroup 函数 |
| `pgnodemx.containerized` | `off` | 强制使用容器化 cgroup 路径 |
| `pgnodemx.cgrouproot` | `/sys/fs/cgroup` | cgroup 挂载位置 |
| `pgnodemx.kdapi_enabled` | `on` | 启用 Kubernetes DownwardAPI |
| `pgnodemx.kdapi_path` | `/etc/podinfo` | DownwardAPI 文件位置 |
