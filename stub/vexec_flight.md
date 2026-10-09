## Usage

Sources:

- [modules/vexec_flight/README.md](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_flight/README.md)
- [modules/vexec_flight/src/flight.h](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_flight/src/flight.h)
- [modules/vexec_flight/src/vexec_flight.c](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_flight/src/vexec_flight.c)
- [modules/vexec_flight/Makefile](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_flight/Makefile)
- [README.md](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/README.md)

`vexec_flight` exposes PostgreSQL 19 through Arrow Flight SQL while `vexec` runs vectorized execution. It is a preload library with C version 0.1.0 and has no SQL control file.

### Core Workflow

```ini
shared_preload_libraries = 'vexec,vexec_flight'
vexec.mode = auto
vexec_flight.listen_addresses = '127.0.0.1'
vexec_flight.port = 32010
```

### Operational Boundaries

Append both libraries to the existing preload list, configure a listener and restart. The listener is disabled when the address is empty; its default port is 32010. Statements require `vexec.mode` auto or force. There is no `CREATE EXTENSION vexec_flight`; bulk ingestion requires the optional `vexec` SQL objects and `vexec.ingest_stream`.

Connections use PostgreSQL roles, database privileges and pg_hba.conf. Configure TLS certificates and hostssl rules before exposing credentials outside a trusted host. `vexec_flight.max_sessions` defaults to 16 and consumes background-worker slots; `vexec_flight.idle_session_timeout` bounds idle sessions. The module implements queries, prepared statements, transactions, cancellation, metadata and ingestion, but not Substrait plans or DoExchange. It needs nghttp2, protobuf-c and OpenSSL.
