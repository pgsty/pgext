## Usage

Sources:

- [Official README](https://github.com/askyx/pg_tpch/blob/5684779c89b29b4674604c447a30ec07aaebf307/README.md)
- [SQL API and metadata](https://github.com/askyx/pg_tpch/blob/5684779c89b29b4674604c447a30ec07aaebf307/src/lib.rs)
- [Raw heap loader](https://github.com/askyx/pg_tpch/blob/5684779c89b29b4674604c447a30ec07aaebf307/src/loader.rs)
- [Table schema helpers](https://github.com/askyx/pg_tpch/blob/5684779c89b29b4674604c447a30ec07aaebf307/src/schema.rs)
- [Extension control file](https://github.com/askyx/pg_tpch/blob/5684779c89b29b4674604c447a30ec07aaebf307/pg_tpch.control)

`pg_tpch` 1.0.0 is a pgrx extension for creating the eight TPC-H tables, generating data in-process with `tpchgen-rs`, and retrieving the 22 concrete benchmark queries. Its speed comes from a specialized raw heap-file loader that bypasses normal PostgreSQL tuple insertion and later rebuilds indexes. Use it only in disposable benchmark databases, not on production relations.

### Create and Load a Dataset

Run the workflow in a dedicated schema. Passing true creates the extension's secondary indexes in addition to the primary keys embedded in the table definitions.

```sql
CREATE EXTENSION pg_tpch;

CREATE SCHEMA benchmark;
SET search_path = benchmark, public;

SELECT create_tpch_tables(true);
SELECT * FROM tpch_dbgen(0.01);

SELECT query AS q1 FROM tpch_queries(1) \gset
:q1

SELECT cleanup_tpch_data();
```

If `dblink` is installed, `tpch_dbgen(double precision)` can open local connections and load tables in parallel; otherwise it loads serially. The caller and local authentication configuration must permit those connections. Individual functions `generate_region`, `generate_nation`, `generate_part`, `generate_supplier`, `generate_customer`, `generate_partsupp`, `generate_orders`, and `generate_lineitem` load one existing table and report rows, heap time, and reindex time.

### Management and Metadata

- `create_tpch_tables(boolean)`: create standard tables in the current schema; the default omits secondary indexes.
- `create_tpch_indexes()`: create index definitions stored in metadata.
- `cleanup_tpch_data()`: truncate all eight tables while retaining definitions and indexes.
- `drop_tpch_tables()`: drop all eight tables.
- `tpch_queries(integer)`: return one query, or all queries when the argument is NULL.
- `tpch.tpch_table_metadata`: editable complete table DDL and secondary-index arrays used by future creation calls.
- `tpch.tpch_query_metadata`: the built-in concrete query texts.

Changing metadata executes stored DDL in the target schema. Restrict updates and extension functions to a benchmark owner; do not let untrusted users inject definitions.

### Raw-Loader Safety Boundary

The loader takes an `AccessExclusiveLock`, constructs heap pages itself, opens the relation's physical segment files with the server process, and writes pages directly. It records only a limited new-page WAL image path, bypasses normal buffer-manager insertion and index maintenance, then issues `REINDEX TABLE`. It also chooses aggressive session-local reindex settings from host memory/CPU, including up to one third of detected RAM for `maintenance_work_mem` and `synchronous_commit = off`.

These internal APIs and physical layouts are major-version-specific. A backend crash, operating-system crash, standby, backup, checksum, storage, or index interaction may differ from ordinary SQL/COPY semantics. Never point a generator at existing valuable tables even if their names and columns match TPC-H. Use an exact PostgreSQL-major artifact, isolate the database and tablespace, take no concurrent workload, verify checksums and recovery on a clone, and discard the cluster after benchmarking. Performance numbers are observations from the upstream test environment, not guarantees against correctly configured `COPY` on another host.
