

## Usage

> [pg_proctab: access OS process table from PostgreSQL](https://github.com/markwkm/pg_proctab)

pg_proctab allows querying operating system process, CPU, memory, disk, and load statistics from within PostgreSQL via SQL functions.

### Functions

**Process information** (`pg_proctab()`):

```sql
-- All PostgreSQL process info
SELECT pid, comm, state, utime, stime, vsize, rss, reads, writes
FROM pg_proctab();

-- Join with pg_stat_activity for per-query resource usage
SELECT a.pid, a.query, p.utime, p.stime, p.vsize, p.rss
FROM pg_stat_activity a
JOIN pg_proctab() p ON a.pid = p.pid;
```

Returns per-process: `pid`, `comm`, `fullcomm`, `state`, `ppid`, `utime`, `stime`, `priority`, `nice`, `num_threads`, `vsize`, `rss`, `processor`, `uid`, `username`, `rchar`, `wchar`, `reads`, `writes`, and more.

**CPU time** (`pg_cputime()`):

```sql
SELECT "user", nice, system, idle, iowait FROM pg_cputime();
```

**Load average** (`pg_loadavg()`):

```sql
SELECT load1, load5, load15, last_pid FROM pg_loadavg();
```

**Memory usage** (`pg_memusage()`):

```sql
SELECT memused, memfree, memshared, membuffers, memcached,
       swapused, swapfree, swapcached
FROM pg_memusage();
```

**Disk usage** (`pg_diskusage()`):

```sql
SELECT devname, reads_completed, writes_completed,
       sectors_read, sectors_written
FROM pg_diskusage();
```
