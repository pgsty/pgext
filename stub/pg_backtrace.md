## Usage

Sources:

- [Official README](https://github.com/postgrespro/pg_backtrace/blob/de245a5f6b42ba870259ba68b4fbdd0c8b7bae69/README)
- [Official extension SQL](https://github.com/postgrespro/pg_backtrace/blob/de245a5f6b42ba870259ba68b4fbdd0c8b7bae69/pg_backtrace--1.0.sql)
- [Official C implementation](https://github.com/postgrespro/pg_backtrace/blob/de245a5f6b42ba870259ba68b4fbdd0c8b7bae69/pg_backtrace.c)

`pg_backtrace` is a low-level diagnostic extension that adds native stack traces to selected PostgreSQL errors and logs traces for certain signals. It is intended for investigating backend failures or hangs when a debugger or core dump is unavailable, not for routine application observability.

### Initialization and Use

Creating the extension only installs two functions. The library and its hooks are initialized when a function is first called in that backend:

```sql
CREATE EXTENSION pg_backtrace;
SELECT pg_backtrace_init();

SET pg_backtrace.level = 'ERROR';
SELECT 1 / 0;
```

`pg_backtrace.level` selects the error severity at which the executor hook attaches a backtrace; the default is `ERROR`. `pg_backtrace_init()` exists to force library loading. `pg_backtrace_sigsegv()` deliberately raises a segmentation fault and is only a destructive diagnostic test—never call it on a backend that carries useful work.

After initialization, the extension installs handlers for `SIGSEGV`, `SIGBUS`, `SIGFPE`, and `SIGINT`. Sending `SIGINT` to a backend can log its current native stack before the previous signal handler continues.

### Caveats

This code replaces signal handlers and the executor hook within each initialized backend. Test interactions with other hook- or signal-based extensions before use. Stack symbol quality is platform and linker dependent: PostgreSQL is commonly built without `-rdynamic`, static symbols are unavailable, and unresolved addresses may need `addr2line` or matching debug binaries. Native addresses and function names can expose sensitive implementation details to clients and logs, so restrict initialization and log access. The implementation targets Unix-like systems providing `backtrace()`.
