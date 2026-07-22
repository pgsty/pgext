## Usage

Sources:

- [Official README](https://github.com/sangli00/nanomsgtopdb/blob/87df402495454362ccb4ed6ab58c33e7683e7401/README.md)
- [Background-worker implementation](https://github.com/sangli00/nanomsgtopdb/blob/87df402495454362ccb4ed6ab58c33e7683e7401/nanomsgtopdb.c)
- [Extension SQL](https://github.com/sangli00/nanomsgtopdb/blob/87df402495454362ccb4ed6ab58c33e7683e7401/nanomsgtopdb--1.0.sql)

`nanomsgtopdb` 1.0 is a historical PipelineDB background-worker extension. It receives text datagrams over nanomsg TCP and inserts them into the fixed PipelineDB stream `public.generic_stream`. It depends on old PipelineDB internals and is not a general-purpose ingestion extension for current PostgreSQL.

### Configuration and Startup

Install the native nanomsg dependency, place the module in `shared_preload_libraries`, set its postmaster-time GUCs, and restart:

```conf
shared_preload_libraries = 'pipelinedb,nanomsgtopdb'
nanomsg_work.ip_address = '127.0.0.1'
nanomsg_work.port = 9999
nanomsg_work.database = 'pipeline'
nanomsg_work.isrec = true
nanomsg_work.nanomsg_works = 1
nanomsg_work.nanomsg_max_tuples_a_transaction = 1000
```

Then create the extension in the configured database:

```sql
CREATE EXTENSION nanomsgtopdb;

CREATE CONTINUOUS VIEW message_count AS
SELECT count(*) FROM generic_stream;
```

The installation SQL creates `public.generic_stream(data text)`. A main worker binds `tcp://IP:port`, forwards messages over a local IPC endpoint, and one or more writer workers batch the received text into transactions. `nanomsg_work.format` exists but defaults to `text`, and the implementation's stream tuple descriptor is text-only.

### Operational and Security Boundaries

All GUCs are `PGC_POSTMASTER`, so changes require restart. The receiver starts only when `nanomsg_work.isrec=true`; `nanomsg_work.nanomsg_works` accepts 1–10 workers, and the transaction batch limit accepts 1–10,000 tuples.

The listener has no authentication, encryption, message acknowledgement, or application-level validation. The upstream default bind address is `0.0.0.0`; bind to a trusted interface and enforce firewall rules. Delivery is not atomic with the sender, and malformed or hostile text reaches a database stream. The code targets the abandoned PipelineDB extension and old background-worker APIs, so expect source incompatibilities on modern PostgreSQL and do not use it for production without a maintained port and end-to-end failure testing.
