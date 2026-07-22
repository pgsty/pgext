## Usage

Sources:

- [DagDB engine README](https://github.com/norayr-m/dagdb-engine/blob/7cd3c50ccceca47f61b21724a6262518ceadb1e7/README.md)
- [`pg_dagdb` Rust implementation](https://github.com/norayr-m/dagdb-engine/blob/7cd3c50ccceca47f61b21724a6262518ceadb1e7/dagdb/pg_dagdb/src/lib.rs)
- [`pg_dagdb` control file](https://github.com/norayr-m/dagdb-engine/blob/7cd3c50ccceca47f61b21724a6262518ceadb1e7/dagdb/pg_dagdb/pg_dagdb.control)

`pg_dagdb` is a PostgreSQL client bridge to the experimental DagDB ranked directed-acyclic-graph engine. The SQL extension does not store or evaluate the graph itself: it sends DagDB DSL commands to a separate local daemon over a fixed Unix socket and reads row results from a fixed shared-memory file.

### Core Workflow

Install and start the matching `dagdb_daemon` on the PostgreSQL server host before querying the extension:

```sql
CREATE EXTENSION pg_dagdb;

SELECT dagdb_status();
SELECT dagdb_tick(100);

SELECT *
FROM dagdb_exec('NODES AT RANK 2 WHERE truth=1');
```

For commands that return node rows, `dagdb_exec` exposes `node_id`, `rank`, `truth`, and `node_type`; status-style commands return their daemon response in `message` instead.

### Object Index

- `dagdb_status() RETURNS text` sends `STATUS` and reports either the daemon reply or an `OFFLINE` message.
- `dagdb_tick(n integer DEFAULT 1) RETURNS text` asks the GPU engine to run ticks.
- `dagdb_exec(command text)` returns a table with nullable `node_id`, `rank`, `truth`, `node_type`, and `message` columns.

### Operational Notes

Catalog version `0.0.0` reflects a pre-release interface. The control file marks the extension non-relocatable, superuser-only, and untrusted; it declares no preload requirement. The implementation hard-codes `/tmp/dagdb.sock` and `/tmp/dagdb_shm_file`, so the daemon must run on the database server with compatible file permissions and binary layout. These paths also create a local trust boundary: restrict extension execution and protect the socket and shared-memory file from other host users.

Errors are returned as text rows rather than SQL exceptions, and a missing or malformed shared-memory result can produce only the daemon message or an empty result. The upstream project explicitly describes unvalidated limits and an Apple Silicon/Metal design. Pin the daemon and extension to the same commit, treat this as experimental, and do not infer transactional consistency or PostgreSQL durability from the bridge.
