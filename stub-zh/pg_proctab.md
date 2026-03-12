


## 用法

> [pg_proctab: 从 PostgreSQL 访问操作系统进程表](https://github.com/markwkm/pg_proctab)

pg_proctab 允许在 PostgreSQL 内部通过 SQL 函数查询操作系统的进程、CPU、内存、磁盘和负载统计信息。

### 函数

**进程信息** (`pg_proctab()`)：

```sql
-- 所有 PostgreSQL 进程信息
SELECT pid, comm, state, utime, stime, vsize, rss, reads, writes
FROM pg_proctab();

-- 与 pg_stat_activity 关联获取每个查询的资源使用情况
SELECT a.pid, a.query, p.utime, p.stime, p.vsize, p.rss
FROM pg_stat_activity a
JOIN pg_proctab() p ON a.pid = p.pid;
```

返回每个进程的信息：`pid`、`comm`、`fullcomm`、`state`、`ppid`、`utime`、`stime`、`priority`、`nice`、`num_threads`、`vsize`、`rss`、`processor`、`uid`、`username`、`rchar`、`wchar`、`reads`、`writes` 等。

**CPU 时间** (`pg_cputime()`)：

```sql
SELECT "user", nice, system, idle, iowait FROM pg_cputime();
```

**负载均值** (`pg_loadavg()`)：

```sql
SELECT load1, load5, load15, last_pid FROM pg_loadavg();
```

**内存使用** (`pg_memusage()`)：

```sql
SELECT memused, memfree, memshared, membuffers, memcached,
       swapused, swapfree, swapcached
FROM pg_memusage();
```

**磁盘使用** (`pg_diskusage()`)：

```sql
SELECT devname, reads_completed, writes_completed,
       sectors_read, sectors_written
FROM pg_diskusage();
```
