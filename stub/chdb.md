## Usage

Sources:

- [PGXN 0.1.2 README](https://pgxn.org/dist/chdb/0.1.2/README.html)
- [chdb 0.1 documentation](https://pgxn.org/dist/chdb/0.1.2/doc/chdb.html)
- [chdb control file](https://api.pgxn.org/src/chdb/chdb-0.1.2/chdb.control)
- [Apache 2.0 license](https://api.pgxn.org/src/chdb/chdb-0.1.2/LICENSE.md)
- [0.1.2 changelog](https://api.pgxn.org/src/chdb/chdb-0.1.2/CHANGELOG.md)

`chdb` executes chDB queries from PostgreSQL through an isolated helper process. Use it for occasional ClickHouse-format analytics or reads from chDB table functions when the returned row shape can be declared explicitly; it is not a persistent embedded ClickHouse database.

### Core Workflow

```sql
CREATE EXTENSION chdb;

SELECT *
FROM chdb_query('SELECT number, number * number FROM numbers(5)')
AS result(number bigint, square bigint);
```

`chdb_query(text)` returns `record`, so every call needs an `AS` column definition list matching the chDB result. No role receives `EXECUTE` by default; an administrator must grant the function deliberately.

```sql
GRANT EXECUTE ON FUNCTION chdb_query(text) TO analytics_role;
SELECT pgchdb_version();
```

### Resource Controls

- `chdb.max_memory` maps to chDB's maximum memory setting; `0` means unlimited.
- `chdb.max_threads` limits query-processing threads.
- `chdb.max_parsing_threads` limits parallel parsing for supported input formats.

These settings require superuser privileges. Set finite values before exposing `chdb_query(text)` to workloads that can issue large scans, because an unconstrained helper can compete with PostgreSQL for memory, CPU, disk, and network bandwidth.

### Architecture and Boundaries

Each call starts a helper-backed, temporary chDB database and removes it when the query completes. Objects created by one `chdb_query(text)` call are not visible to the next. A helper crash fails the initiating backend without sharing PostgreSQL memory, but it can still consume host resources and reach data sources allowed by chDB.

Version 0.1.2 requires PostgreSQL 15 or newer and libchdb 26.7.0 or newer on Linux or macOS. The control file is superuser-only, untrusted, relocatable, and exposes extension version `0.1` while the library/distribution version is `0.1.2`. Review remote credentials and URLs carefully: query text can read network or local sources supported by chDB, and results cross an explicit type-conversion boundary into PostgreSQL.

### Version 0.1.2

This is a binary-only update: the `chdb` control version remains `0.1`, with no extension SQL upgrade required. New mappings cover Tuple, Map, unflattened Nested and composite types, plus additional LowCardinality and parameterized JSON conversions. Arrays and composites have different shapes; declare matching destination types and test representative values.
