## Usage

Sources:

- [Upstream README](https://github.com/charmitro/pg_session_trace/blob/efb2a0bc68d3d44723e5259d7792755987f452d2/README.md)
- [Version 1.0 install SQL](https://github.com/charmitro/pg_session_trace/blob/efb2a0bc68d3d44723e5259d7792755987f452d2/sql/pg_session_trace--1.0.sql)
- [Session-management implementation](https://github.com/charmitro/pg_session_trace/blob/efb2a0bc68d3d44723e5259d7792755987f452d2/src/pg_session_trace_session.c)

pg_session_trace records detailed events from selected PostgreSQL sessions into per-session binary trace files. It is designed for PostgreSQL 15 and later and can capture statements, execution timing, plans, waits, and—at the highest level—bind values.

### Server configuration

Preload the module and restart PostgreSQL before installing its SQL objects:

```ini
shared_preload_libraries = 'pg_session_trace'
pg_session_trace.enabled = on
pg_session_trace.output_directory = 'pg_session_trace'
pg_session_trace.level = 4
```

Enable tracing for a backend, inspect active trace sessions, adjust detail, and stop tracing with:

```sql
CREATE EXTENSION pg_session_trace;

SELECT pg_session_trace_enable_session(pg_backend_pid());
SELECT pg_session_trace_set_level(pg_backend_pid(), 4);
SELECT * FROM pg_session_trace_list_sessions();
SELECT pg_session_trace_disable_session(pg_backend_pid());
```

Level 0 records basic statements, level 1 adds execution details, and level 4 includes bind values. Use the lowest level that answers the diagnostic question.

### Caveats

- Session activation and control are privileged operations. Keep execute privileges narrow and record who starts each trace.
- Level 4 can capture passwords, tokens, personal data, and application payloads supplied as bind parameters. Trace only approved sessions, protect the output directory, and delete files promptly.
- Trace generation adds CPU, memory, storage, and I/O overhead. Binary ring buffers can drop events under pressure; monitor the reported events_dropped count.
- The output files are server-local operational artifacts, not an audit log. File modes reduce accidental access but do not replace host access control, encryption, and retention management.
- Startup settings such as maximum session capacity require a restart. Test decoding tools and extension/server version compatibility before an incident.
