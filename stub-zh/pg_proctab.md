## 用法

来源：

- [pgnodemx 2.0.1 文档](https://github.com/pgnodemx/pgnodemx/blob/v2.0.1/README.md)
- [兼容层控制文件](https://github.com/pgnodemx/pgnodemx/blob/v2.0.1/pg_proctab.control)
- [兼容层 SQL](https://github.com/pgnodemx/pgnodemx/blob/v2.0.1/pg_proctab--0.0.10-compat.sql)

`pgnodemx` 包中的 `pg_proctab` 提供五个 SQL 包装函数，兼容旧的进程监控接口。它依赖 `pgnodemx`；软件包版本为 2.0.1，SQL 安装版本为 `0.0.10-compat`。

### 启用

保留 `shared_preload_libraries` 中已有的条目，加入 `pgnodemx`，然后重启 PostgreSQL：

```ini
shared_preload_libraries = 'pgnodemx'
```

使用超级用户将两个扩展安装到同一模式：

```sql
CREATE EXTENSION pgnodemx;
CREATE EXTENSION pg_proctab VERSION '0.0.10-compat';
```

主控制文件没有默认版本，因此必须显式指定版本。本轮核验的 PGDG RPM 缺少该主控制文件，执行此命令前需安装包含它的软件包。查询底层操作系统指标需要超级用户权限或 `pg_monitor` 成员资格。

### 查询

```sql
SELECT a.pid, a.query, p.utime, p.stime, p.vsize, p.rss
FROM pg_stat_activity a
JOIN pg_proctab() p ON a.pid = p.pid;

SELECT "user", nice, system, idle, iowait FROM pg_cputime();
SELECT load1, load5, load15, last_pid FROM pg_loadavg();
SELECT memused, memfree, memshared, membuffers, memcached,
       swapused, swapfree, swapcached
FROM pg_memusage();
SELECT devname, reads_completed, writes_completed,
       sectors_read, sectors_written
FROM pg_diskusage();
```

这些函数读取 PostgreSQL 服务端可见的 Linux 进程与系统统计信息，容器命名空间和操作系统权限会影响可见范围。在此兼容层中，`rss` 使用 KiB，内存汇总值使用 MiB，`memshared` 来自 `Shmem`。
