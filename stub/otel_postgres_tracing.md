## Usage

Sources:

- [README.md](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/README.md)
- [otel_postgres_tracing/otel_postgres_tracing--0.1.1.sql](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/otel_postgres_tracing/otel_postgres_tracing--0.1.1.sql)
- [otel_postgres_tracing/otel_postgres_tracing.c](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/otel_postgres_tracing/otel_postgres_tracing.c)
- [otel_postgres_tracing/otel_postgres_tracing.control](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/otel_postgres_tracing/otel_postgres_tracing.control)

`otel_postgres_tracing` instruments query execution and utility statements through PostgreSQL hooks, using the shared `otel_api` provider. An exporter is needed to persist or send the produced spans.

### Core Workflow

```ini
shared_preload_libraries = 'otel_api,otel_postgres_tracing,otel_demo_exporter'
otel_demo_exporter.output_file = '/var/log/postgresql/otel_spans.jsonl'
```

### Operational Boundaries

The library installs hooks when loaded. Cluster preload is the recommended path; per-session loading is best-effort and affects only that backend. The versioned SQL contains no user-facing SQL objects, so `CREATE EXTENSION` is optional bookkeeping, not activation. Preload ordering with the API is not significant in this revision.

The project is pre-release: wire and file formats may change. Stock PostgreSQL 14+ supports the GUC/sqlcommenter fallback; protocol-header propagation and some error/parallel-worker behavior require the companion patched PostgreSQL. Do not assume those patches are present on an ordinary server. Append libraries to existing preload entries and restart for cluster-wide enablement. Protect traces because statement text and attributes may contain sensitive data.
