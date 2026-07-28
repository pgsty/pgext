## Usage

Sources:

- [Official upstream README](https://github.com/mstryoda/pgtrace/blob/77c2efdef80c0c50627589a8ad05f1058cef1ac5/README.md)
- [Official extension control file (pg_otel_tracer.control)](https://github.com/mstryoda/pgtrace/blob/77c2efdef80c0c50627589a8ad05f1058cef1ac5/pg_otel_tracer.control)
- [Official extension SQL (pg_otel_tracer--0.1.0.sql)](https://github.com/mstryoda/pgtrace/blob/77c2efdef80c0c50627589a8ad05f1058cef1ac5/sql/pg_otel_tracer--0.1.0.sql)

`pg_otel_tracer` — OpenTelemetry tracing extension for PostgreSQL. Extracts W3C traceparent from SQL comments and exports query lifecycle spans via OTLP/HTTP. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_otel_tracer;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
