## Usage

Sources:

- [README.md](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/README.md)
- [otel_api/otel_api--0.1.1.sql](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/otel_api/otel_api--0.1.1.sql)
- [otel_api/otel_api.c](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/otel_api/otel_api.c)
- [otel_api/otel_api.control](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/otel_api/otel_api.control)

`otel_api` provides common trace-context state and a C rendezvous API for instrumentation and exporters. It does not itself ship traces over OTLP or create query spans without an instrumentation consumer.

### Core Workflow

```ini
shared_preload_libraries = 'otel_api,otel_postgres_tracing'
```

```sql
CREATE EXTENSION otel_api;
BEGIN;
SET LOCAL otel_api.traceparent = '00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01';
SELECT otel_current_traceparent();
COMMIT;
```

### Operational Boundaries

Load the API at backend start. `otel_api.traceparent` and `otel_api.tracestate` carry context through the GUC fallback; `otel_current_traceparent()` is the optional SQL inspection helper installed by `CREATE EXTENSION`. Creating its C SQL function requires a superuser. Pair the library with `otel_postgres_tracing` for query instrumentation and an exporter to consume spans.

The project is pre-release: wire and file formats may change. Stock PostgreSQL 14+ supports the GUC/sqlcommenter fallback; protocol-header propagation and some error/parallel-worker behavior require the companion patched PostgreSQL. Do not assume those patches are present on an ordinary server. Append libraries to existing preload entries and restart for cluster-wide enablement. Protect traces because statement text and attributes may contain sensitive data.
