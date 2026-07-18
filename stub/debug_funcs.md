## Usage

Sources:

- [Pinned extension control file](https://github.com/MasahikoSawada/debug_funcs/blob/8508ef191735ef48375a2e626be1986af4d35dd8/debug_funcs.control)
- [Pinned SQL object definitions](https://github.com/MasahikoSawada/debug_funcs/blob/8508ef191735ef48375a2e626be1986af4d35dd8/debug_funcs--1.0.sql)

`debug_funcs` exposes low-level C functions for experimenting with PostgreSQL buffer cleanup locks, buffer lock modes, extension/relation locks, and lock timing. The SQL surface includes `pg_LockBufferForCleanup(regclass,bigint,int)`, `pg_LockBuffer(regclass,bigint,text,int)`, `pg_lockforextension(regclass)`, `replock()`, `show_define_variables()`, `rel_lock(regclass,bool)`, `rel_unlock(regclass)`, `rel_lock_unlock(regclass,int)`, and `extlock_bench(regclass,int)`.

```sql
CREATE EXTENSION debug_funcs;
SELECT show_define_variables();
```

The lock functions accept sleep durations and, for `pg_LockBuffer`, the modes `share` or `exclusive`. They deliberately acquire internal locks and can block concurrent sessions or hold resources for the requested interval. Do not call them in application traffic, connection-pool health checks, or automated production diagnostics.

The reviewed repository has no user guide or compatibility matrix; version `1.0` directly uses PostgreSQL internal buffer and relation-locking APIs. Build only against the exact target server headers, run on disposable test data, impose statement timeouts, and keep a separate session available for cancellation or backend termination.
