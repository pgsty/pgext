## Usage

Sources:

- [README.md](https://github.com/postgrespro/pg_uprobe/blob/336336c087070c8cc35a697216a6a51bd34905b4/README.md)
- [pg_uprobe--0.3.sql](https://github.com/postgrespro/pg_uprobe/blob/336336c087070c8cc35a697216a6a51bd34905b4/pg_uprobe--0.3.sql)
- [doc/trace_session.md](https://github.com/postgrespro/pg_uprobe/blob/336336c087070c8cc35a697216a6a51bd34905b4/doc/trace_session.md)
- [pg_uprobe.control](https://github.com/postgrespro/pg_uprobe/blob/336336c087070c8cc35a697216a6a51bd34905b4/pg_uprobe.control)

`pg_uprobe` captures query text, plans, per-node execution, waits and locks, and profiles native functions through Frida Gum. Upstream lists PostgreSQL 15–17, Linux and FreeBSD; ARM64 is not tested.

### Core Workflow

```ini
shared_preload_libraries = 'pg_uprobe'
```

```sql
CREATE EXTENSION pg_uprobe;
SELECT start_session_trace();
SELECT 1;
SELECT stop_session_trace();
```

### Operational Boundaries

Preload the library and restart before creating the extension as a superuser. `start_session_trace()` and `stop_session_trace()` delimit a trace. `set_uprobe()` installs TIME, HIST or MEM probes; `list_uprobes()` and `delete_uprobe()` manage them. Native targets must be exported, non-static, non-inline functions. `pg_uprobe.data_dir` defaults to the server data directory subfolder, and `pg_uprobe.trace_file_limit` limits trace files. Protect those files because they include query text and plans. Tracing adds overhead and is intended for focused investigations.
