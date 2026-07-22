## Usage

Sources:

- [Official README](https://github.com/RyanMarcus/pg_session_stats/blob/8d114d5b653304a3039741c5cad8fee751bea1b3/README.md)
- [Extension control file](https://github.com/RyanMarcus/pg_session_stats/blob/8d114d5b653304a3039741c5cad8fee751bea1b3/pg_session_stats.control)
- [Empty extension SQL](https://github.com/RyanMarcus/pg_session_stats/blob/8d114d5b653304a3039741c5cad8fee751bea1b3/pg_session_stats--0.0.1.sql)
- [Executor hook and SQLite writer](https://github.com/RyanMarcus/pg_session_stats/blob/8d114d5b653304a3039741c5cad8fee751bea1b3/main.c)

`pg_session_stats` is a Linux/PostgreSQL 12 experiment that records backend CPU and `/proc` snapshots in an external SQLite database. It has no SQL API: the extension script is empty, and the only effective activation path is server preload plus restart.

### Core Workflow

Install a binary linked with SQLite, choose a PostgreSQL-owned file location, preload the library, and restart. Protect the containing directory because the file includes process status and I/O information.

```conf
shared_preload_libraries = 'pg_session_stats'
pg_session_stats.path = '/var/lib/postgresql/pg_session_stats.sqlite'
```

The SQLite database contains a `log` table with `master_pid`, `my_pid`, `usage`, `procstatus`, and `procio`. The README predates the `procio` column found in the implementation. Each executor completion appends one row; `usage` is cumulative process CPU from `clock()`, not the CPU delta for that query.

### Reliability and Resource Caveats

Despite its name, the hook writes at every `ExecutorEnd`, not only when a session ends. SQLite writes are synchronous in the query backend and use a five-second busy timeout, so lock contention or filesystem latency can delay application queries. The external write is not part of the PostgreSQL transaction and is not rolled back.

The pinned code reads Linux `/proc`, assumes PostgreSQL 12 worker command-line behavior, and uses a fragile parent-PID heuristic. More seriously, the `/proc` readers do not close their files, leaking descriptors on every logged executor completion; a long-lived backend can eventually exhaust file descriptors. Error paths also lack robust cleanup. Do not use `pg_session_stats` on production systems. If evaluating it in an isolated lab, bound session lifetime, monitor file descriptors and SQLite growth, restrict file access, and independently validate worker attribution and CPU meaning.
