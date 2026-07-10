


## Usage

> Sources: [pg_stat_backtrace upstream README](https://github.com/Nickyoung0/pg_stat_backtrace), [upstream changelog](https://github.com/Nickyoung0/pg_stat_backtrace/blob/main/CHANGELOG.md), local source tarball `pg_stat_backtrace-1.0.0.tar.gz`.

`pg_stat_backtrace` captures or logs the C-level stack backtrace of a PostgreSQL backend or auxiliary process on the same Linux host. It uses `ptrace(PTRACE_SEIZE)` plus `libunwind`; it does not use `shared_preload_libraries` and does not send `SIGSTOP` to the target.

```sql
CREATE EXTENSION pg_stat_backtrace;
```

### Capture A Backtrace

Find a target process from PostgreSQL views, then call `pg_get_backtrace(pid)`:

```sql
SELECT pid, backend_type, state, wait_event, query
FROM pg_stat_activity
WHERE backend_type = 'autovacuum worker';

SELECT pg_get_backtrace(123456);
```

The returned text uses a `pstack(1)`-style format:

```text
#0  0x00007f5e6c1a4d9e in __epoll_wait+0x4e
#1  0x000055f1a8c2a213 in WaitEventSetWaitBlock+0x83
#2  0x000055f1a8c2a04e in WaitEventSetWait+0xfe
```

### Write To The Server Log

Use `pg_log_backtrace(pid)` when the result should go through the normal PostgreSQL log pipeline:

```sql
SELECT pg_log_backtrace(123456);

SELECT pid, pg_log_backtrace(pid)
FROM pg_stat_activity
WHERE backend_type = 'walsender';
```

The function returns `true` on success.

### Permissions

By default, execute privilege is revoked from `PUBLIC` for both functions. Grant access only to trusted monitoring roles:

```sql
GRANT EXECUTE ON FUNCTION pg_get_backtrace(int) TO observability;
GRANT EXECUTE ON FUNCTION pg_log_backtrace(int) TO observability;
```

The C code still enforces target checks:

- Superusers may target any PostgreSQL process in the instance, including auxiliary processes such as `walwriter`, `checkpointer`, `walsender`, autovacuum workers, startup, and archiver processes.
- Non-superusers may target only regular backends owned by roles they are members of.
- Auxiliary processes have no role ownership and are rejected for non-superusers.
- A non-superuser may not target a superuser-owned backend, even with role membership.

### Input And Error Behavior

Both functions are `VOLATILE STRICT PARALLEL RESTRICTED`.

```sql
SELECT pg_get_backtrace(NULL::int);  -- NULL, no ptrace attempt
SELECT pg_log_backtrace(NULL::int);  -- NULL, no ptrace attempt
SELECT pg_get_backtrace(0);          -- WARNING, then NULL
SELECT pg_log_backtrace(-1);         -- WARNING, then false
```

Self-targeting is rejected because a Linux process cannot ptrace itself:

```sql
SELECT pg_get_backtrace(pg_backend_pid());
```

### Operational Caveats

- Version 1.0.0 supports PostgreSQL 14-18. Upstream 1.0.0 also advertises PostgreSQL 19 compatibility.
- The extension is Linux-only and depends on `libunwind` / `libunwind-ptrace` at build and runtime.
- On hosts with Yama ptrace restrictions, backend-to-backend capture may require `kernel.yama.ptrace_scope = 0`.
- The target process is briefly paused while the stack is unwound. Avoid tight loops against critical processes such as `walwriter`, `checkpointer`, or synchronous-replication `walsender` on busy primaries.
- Linux permits only one tracer per target process. Concurrent calls against the same PID can fail with `EPERM`; retry after the in-flight call finishes.
