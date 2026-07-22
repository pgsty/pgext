## Usage

Sources:

- [Official README](https://github.com/pipelinedb/pipelinedb/blob/5cc6ef58ab5b1c84bb7b4e932f99bf5c347f46d8/README.md)
- [Extension control file](https://github.com/pipelinedb/pipelinedb/blob/5cc6ef58ab5b1c84bb7b4e932f99bf5c347f46d8/pipelinedb.control)
- [Runtime configuration](https://github.com/pipelinedb/pipelinedb/blob/5cc6ef58ab5b1c84bb7b4e932f99bf5c347f46d8/src/config.c)
- [Anonymous update-check implementation](https://github.com/pipelinedb/pipelinedb/blob/5cc6ef58ab5b1c84bb7b4e932f99bf5c347f46d8/src/update.c)

`pipelinedb` 1.1.0 runs continuous SQL queries that incrementally aggregate events arriving on streams. A continuous view stores aggregate state rather than raw input rows, which is useful for high-throughput time-series summaries but changes durability, replay, and query semantics compared with ordinary tables.

### Enable the Runtime

The module must be loaded through `shared_preload_libraries` and PostgreSQL must be restarted. This archived version enables anonymous update checks by default; disable them before startup because the code sends aggregate installation/runtime data over cleartext HTTP to the obsolete `anonymous.pipelinedb.com` service.

```conf
shared_preload_libraries = 'pipelinedb'
pipelinedb.anonymous_update_checks = off
```

```sql
CREATE EXTENSION pipelinedb;
```

The final repository code supports PostgreSQL 10.1 through 10.5 and PostgreSQL 11.0 on 64-bit systems and requires ZeroMQ. Do not load this binary on an untested major.

### Stream and Continuous View

```sql
CREATE FOREIGN TABLE event_stream (
    key integer,
    value integer
)
SERVER pipelinedb;

CREATE VIEW event_counts
WITH (action = materialize) AS
SELECT key, count(*) AS count
FROM event_stream
GROUP BY key;

INSERT INTO event_stream (key, value)
SELECT (random() * 10)::integer, g
FROM generate_series(1, 100000) AS g;

SELECT *
FROM event_counts
ORDER BY count DESC;
```

The foreign table is a stream endpoint, not persistent row storage. Inserts emit events to continuous queries that read the stream; the raw rows are discarded unless the continuous query explicitly materializes the needed information elsewhere.

### Main Objects and Controls

- A view created with `action = materialize` is a continuous view with an internal materialization table.
- A view created with `action = transform` can filter or transform a stream and send output through functions such as `pipelinedb.insert_into_stream(...)`.
- `pipelinedb.get_views()` and `pipelinedb.get_streams()` inspect registered continuous objects.
- `pipelinedb.activate(text)`, `pipelinedb.deactivate(text)`, and `pipelinedb.truncate_continuous_view(text)` control a continuous query.
- `pipelinedb.stream_insert_level` chooses asynchronous or synchronous acknowledgement behavior; test its exact loss and latency tradeoff before ingestion.

Do not write internal materialization tables directly; `pipelinedb.matrels_writable` is off by default for this reason. Worker, combiner, queue, batch, commit-interval, and memory GUCs are cluster-level capacity controls and must be sized together.

### Lifecycle and Durability

PipelineDB joined Confluent and upstream announced no new releases beyond the old line; the repository is deprecated. Stream acceptance, worker processing, and materialized aggregate commit are distinct durability points. Keep a replayable source of truth outside the stream, test crash/restart and backup/restore behavior, monitor `pipelinedb.query_stats` and `pipelinedb.stream_stats`, and plan migration before adopting it for any retained system.
