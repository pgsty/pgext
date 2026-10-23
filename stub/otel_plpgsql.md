## Usage

Sources:

- [Official source](https://github.com/ringerc/postgres_otel_api/blob/4192d644f04a9f52285bae55e565649e19b7cef2/otel_plpgsql/otel_plpgsql.control)
- [Official source](https://github.com/ringerc/postgres_otel_api/blob/4192d644f04a9f52285bae55e565649e19b7cef2/otel_plpgsql/otel_plpgsql--0.1.sql)
- [Official source](https://github.com/ringerc/postgres_otel_api/blob/4192d644f04a9f52285bae55e565649e19b7cef2/otel_plpgsql/otel_plpgsql.c)
- [Official source](https://github.com/ringerc/postgres_otel_api/blob/4192d644f04a9f52285bae55e565649e19b7cef2/README.md)

`otel_plpgsql` 0.1 instruments PL/pgSQL function, procedure and statement execution through plugin hooks. It belongs to the experimental postgres_otel_api family. Its versioned SQL script only records installation; loading the library activates tracing.

### Enablement

Load the provider and a suitable exporter, then this consumer, in the intended backend sessions. For an administrative test session:

```sql
LOAD 'otel_api';
LOAD 'otel_plpgsql';
SHOW otel_plpgsql.enabled;
SHOW otel_plpgsql.trace_statements;
```

An exporter is also necessary to retain spans; these commands alone do not configure one. Session preloading is an alternative. Adding libraries to server-wide preload configuration requires a restart. Use the privileges required for loading native libraries and restrict exporter destinations.

### Span Behavior

`otel_plpgsql.enabled` controls instrumentation and `otel_plpgsql.trace_statements` controls statement spans. Both default to enabled in the reviewed source. Function spans nest under the active span or propagated root context; otherwise an outer call can start a new trace. Errors and transaction boundaries affect span lifetime, so validate traces for exception-handling and transaction-controlling procedures.

The project documents an unpatched PostgreSQL 14+ build path, while advanced protocol propagation requires its companion core patches. Existing PL/pgSQL instrumentation can share or compete for plugin hooks; test coexistence. Tracing can expose code names and execution information and adds work to each instrumented call.

The separately built stub variant compiles out telemetry calls for overhead comparison. It is not an exporter or a way to obtain traces. This extension has no SQL query API and does not require a fabricated SQL activation statement.
